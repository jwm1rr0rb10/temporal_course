package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"tcourse/durable"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	w := worker.New(c, durable.TaskQueue, worker.Options{})
	w.RegisterWorkflow(durable.TwoStepWorkflow)
	w.RegisterActivity(durable.Step)
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln(err)
	}
}
