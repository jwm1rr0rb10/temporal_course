package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"tcourse/checkout"
)

func main() {
	c, err := client.Dial(client.Options{Namespace: "orders"})
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	w := worker.New(c, checkout.TaskQueue, worker.Options{})
	w.RegisterWorkflow(checkout.CheckoutWorkflow)
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln(err)
	}
}
