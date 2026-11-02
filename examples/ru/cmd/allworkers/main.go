// Команда allworkers запускает в одном процессе воркеры для всех примеров курса,
// у которых в тексте нет своего воркера. Очереди перечислены в examples/README.md.
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

// demoGateway — заглушка платёжного провайдера для модуля 4:
// токен карты "declined" — отказ банка, "flaky" — две временные ошибки, затем успех.
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
			return "", errors.New("503 от платёжного провайдера")
		}
	}
	return "pay-" + in.OrderID, nil
}

// slowReport — заглушка внешней системы для паттерна опроса из модуля 19:
// отчёт становится готов с третьей проверки.
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
		log.Fatalln("не удалось подключиться к Temporal:", err)
	}
	defer c.Close()

	var workers []worker.Worker
	add := func(queue string, register func(w worker.Worker)) {
		w := worker.New(c, queue, worker.Options{})
		register(w)
		workers = append(workers, w)
	}

	add("determinism", func(w worker.Worker) { // модуль 3
		w.RegisterWorkflow(determinism.ParallelWorkflow)
		w.RegisterWorkflow(determinism.LimitedWorkflow)
		w.RegisterActivity(determinism.Normalize)
	})
	add("shipping", func(w worker.Worker) { // модуль 3: эксперимент с недетерминизмом
		w.RegisterWorkflow(determinism.ShippingWorkflow)
		w.RegisterActivity(determinism.Pack)
		w.RegisterActivity(determinism.Ship)
		w.RegisterActivity(determinism.Notify)
	})
	add("payments", func(w worker.Worker) { // модуль 4
		w.RegisterWorkflow(payments.PaymentWorkflow)
		w.RegisterActivity(&payments.Activities{Gateway: &demoGateway{calls: map[string]int{}}})
	})
	add(heartbeat.TaskQueue, func(w worker.Worker) { // модуль 4: очередь "import"
		w.RegisterWorkflow(heartbeat.ImportWorkflow)
		w.RegisterActivity(heartbeat.ImportRows)
	})
	add("saga", func(w worker.Worker) { // модуль 5
		w.RegisterWorkflow(saga.OrderWorkflow)
		w.RegisterWorkflow(saga.HoldReservationWorkflow)
		w.RegisterActivity(saga.Charge)
		w.RegisterActivity(saga.Refund)
		w.RegisterActivity(saga.ReserveStock)
		w.RegisterActivity(saga.ReleaseStock)
		w.RegisterActivity(saga.Ship)
	})
	add("approval", func(w worker.Worker) { // модуль 7
		w.RegisterWorkflow(approval.ApprovalWorkflow)
		w.RegisterActivity(approval.Notify)
	})
	add(report.TaskQueue, func(w worker.Worker) { // модуль 7: очередь "reports"
		w.RegisterWorkflow(report.ReportWorkflow)
		w.RegisterActivity(report.BuildReport)
	})
	add("batch", func(w worker.Worker) { // модуль 8
		w.RegisterWorkflow(batch.BatchWorkflow)
		w.RegisterWorkflow(batch.ChunkWorkflow)
		w.RegisterActivity(batch.ProcessItem)
	})
	add("accounts", func(w worker.Worker) { // модуль 8
		w.RegisterWorkflow(entity.AccountWorkflow)
	})
	add(orders.TaskQueue, func(w worker.Worker) { // модуль 13: очередь "orders"
		w.RegisterWorkflow(orders.TrackedOrderWorkflow)
	})
	add("polling", func(w worker.Worker) { // модуль 19
		w.RegisterWorkflow(patterns.PollReportWorkflow)
		w.RegisterActivity(&patterns.PollActivities{API: &slowReport{calls: map[string]int{}}})
	})

	for _, w := range workers {
		if err := w.Start(); err != nil {
			log.Fatalln("воркер не запустился:", err)
		}
	}
	log.Printf("запущено воркеров: %d; Ctrl+C — остановка", len(workers))

	<-worker.InterruptCh()
	for _, w := range workers {
		w.Stop()
	}
}
