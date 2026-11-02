// Command allworkers runs, in one process, workers for every course example that
// doesn't come with its own worker in the text. The queues are listed in examples/README.md.
package main

import (
	"context"
	"errors"
	"log"
	"sync"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"tcourse/approval"
	"tcourse/batch"
	"tcourse/determinism"
	"tcourse/entity"
	"tcourse/heartbeat"
	"tcourse/orders"
	"tcourse/patterns"
	"tcourse/payments"
	"tcourse/report"
	"tcourse/saga"
)

// demoGateway is a payment provider stub for Module 4:
// card token "declined" is a bank decline, "flaky" fails twice and then succeeds.
type demoGateway struct {
	mu    sync.Mutex
	calls map[string]int
}

func (g *demoGateway) Charge(ctx context.Context, key string, in payments.ChargeInput) (string, error) {
	switch in.CardToken {
	case "declined":
		return "", payments.ErrDeclined
	case "flaky":
		g.mu.Lock()
		g.calls[key]++
		n := g.calls[key]
		g.mu.Unlock()
		if n <= 2 {
			return "", errors.New("503 from the payment provider")
		}
	}
	return "pay-" + in.OrderID, nil
}

// slowReport is an external system stub for the polling pattern from Module 19:
// the report becomes ready on the third check.
type slowReport struct {
	mu    sync.Mutex
	calls map[string]int
}

func (r *slowReport) IsReady(ctx context.Context, reportID string) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.calls[reportID]++
	return r.calls[reportID] >= 3, nil
}

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("could not connect to Temporal:", err)
	}
	defer c.Close()

	var workers []worker.Worker
	add := func(queue string, register func(w worker.Worker)) {
		w := worker.New(c, queue, worker.Options{})
		register(w)
		workers = append(workers, w)
	}

	add("determinism", func(w worker.Worker) { // Module 3
		w.RegisterWorkflow(determinism.ParallelWorkflow)
		w.RegisterWorkflow(determinism.LimitedWorkflow)
		w.RegisterActivity(determinism.Normalize)
	})
	add("shipping", func(w worker.Worker) { // Module 3: the nondeterminism experiment
		w.RegisterWorkflow(determinism.ShippingWorkflow)
		w.RegisterActivity(determinism.Pack)
		w.RegisterActivity(determinism.Ship)
		w.RegisterActivity(determinism.Notify)
	})
	add("payments", func(w worker.Worker) { // Module 4
		w.RegisterWorkflow(payments.PaymentWorkflow)
		w.RegisterActivity(&payments.Activities{Gateway: &demoGateway{calls: map[string]int{}}})
	})
	add(heartbeat.TaskQueue, func(w worker.Worker) { // Module 4: the "import" queue
		w.RegisterWorkflow(heartbeat.ImportWorkflow)
		w.RegisterActivity(heartbeat.ImportRows)
	})
	add("saga", func(w worker.Worker) { // Module 5
		w.RegisterWorkflow(saga.OrderWorkflow)
		w.RegisterWorkflow(saga.HoldReservationWorkflow)
		w.RegisterActivity(saga.Charge)
		w.RegisterActivity(saga.Refund)
		w.RegisterActivity(saga.ReserveStock)
		w.RegisterActivity(saga.ReleaseStock)
		w.RegisterActivity(saga.Ship)
	})
	add("approval", func(w worker.Worker) { // Module 7
		w.RegisterWorkflow(approval.ApprovalWorkflow)
		w.RegisterActivity(approval.Notify)
	})
	add(report.TaskQueue, func(w worker.Worker) { // Module 7: the "reports" queue
		w.RegisterWorkflow(report.ReportWorkflow)
		w.RegisterActivity(report.BuildReport)
	})
	add("batch", func(w worker.Worker) { // Module 8
		w.RegisterWorkflow(batch.BatchWorkflow)
		w.RegisterWorkflow(batch.ChunkWorkflow)
		w.RegisterActivity(batch.ProcessItem)
	})
	add("accounts", func(w worker.Worker) { // Module 8
		w.RegisterWorkflow(entity.AccountWorkflow)
	})
	add(orders.TaskQueue, func(w worker.Worker) { // Module 13: the "orders" queue
		w.RegisterWorkflow(orders.TrackedOrderWorkflow)
	})
	add("polling", func(w worker.Worker) { // Module 19
		w.RegisterWorkflow(patterns.PollReportWorkflow)
		w.RegisterActivity(&patterns.PollActivities{API: &slowReport{calls: map[string]int{}}})
	})

	for _, w := range workers {
		if err := w.Start(); err != nil {
			log.Fatalln("a worker failed to start:", err)
		}
	}
	log.Printf("workers started: %d; press Ctrl+C to stop", len(workers))

	<-worker.InterruptCh()
	for _, w := range workers {
		w.Stop()
	}
}
