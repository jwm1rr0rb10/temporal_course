package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"tcourse/inventory"
)

func main() {
	c, err := client.Dial(client.Options{Namespace: "inventory"})
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	svc, err := inventory.NewService()
	if err != nil {
		log.Fatalln(err)
	}
	w := worker.New(c, inventory.TaskQueue, worker.Options{})
	w.RegisterWorkflow(inventory.ReserveWorkflow)
	w.RegisterNexusService(svc)
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln(err)
	}
}
