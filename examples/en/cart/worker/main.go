package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"tcourse/cart"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	w := worker.New(c, cart.TaskQueue, worker.Options{})
	w.RegisterWorkflow(cart.CartWorkflow)
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln(err)
	}
}
