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

	// Slots follow the machine's CPU and memory load instead of a fixed number.
	tuner, err := worker.NewResourceBasedTuner(worker.ResourceBasedTunerOptions{
		TargetMem: 0.8,
		TargetCpu: 0.9,
	})
	if err != nil {
		log.Fatalln(err)
	}

	w := worker.New(c, hello.TaskQueue, worker.Options{
		Tuner: tuner,
		// The number of pollers grows and shrinks with the task flow.
		WorkflowTaskPollerBehavior: worker.NewPollerBehaviorAutoscaling(worker.PollerBehaviorAutoscalingOptions{}),
		ActivityTaskPollerBehavior: worker.NewPollerBehaviorAutoscaling(worker.PollerBehaviorAutoscalingOptions{
			MaximumNumberOfPollers: 50,
		}),
		// A limit for the whole queue, not one worker: the external API handles 50 calls per second.
		TaskQueueActivitiesPerSecond: 50,
		// How long to wait for activities to finish after SIGTERM.
		WorkerStopTimeout: 30 * time.Second,
	})
	w.RegisterWorkflow(hello.GreetingWorkflow)
	w.RegisterActivity(hello.Greet)

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln(err)
	}
}
