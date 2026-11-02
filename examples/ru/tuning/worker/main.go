package main

import (
	"log"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"tcourse/hello"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	// Слоты подбираются по загрузке CPU и памяти машины, а не задаются числом.
	tuner, err := worker.NewResourceBasedTuner(worker.ResourceBasedTunerOptions{
		TargetMem: 0.8,
		TargetCpu: 0.9,
	})
	if err != nil {
		log.Fatalln(err)
	}

	w := worker.New(c, hello.TaskQueue, worker.Options{
		Tuner: tuner,
		// Число поллеров растёт и падает вместе с потоком задач.
		WorkflowTaskPollerBehavior: worker.NewPollerBehaviorAutoscaling(worker.PollerBehaviorAutoscalingOptions{}),
		ActivityTaskPollerBehavior: worker.NewPollerBehaviorAutoscaling(worker.PollerBehaviorAutoscalingOptions{
			MaximumNumberOfPollers: 50,
		}),
		// Лимит на всю очередь, а не на один воркер: внешний API выдерживает 50 вызовов в секунду.
		TaskQueueActivitiesPerSecond: 50,
		// Сколько ждать завершения activities после SIGTERM.
		WorkerStopTimeout: 30 * time.Second,
	})
	w.RegisterWorkflow(hello.GreetingWorkflow)
	w.RegisterActivity(hello.Greet)

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln(err)
	}
}
