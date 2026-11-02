package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"tcourse/hello"
)

func main() {
	// Empty Options: localhost:7233 and the "default" namespace.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("could not connect to Temporal:", err)
	}
	defer c.Close()

	w := worker.New(c, hello.TaskQueue, worker.Options{})
	w.RegisterWorkflow(hello.GreetingWorkflow)
	w.RegisterActivity(hello.Greet)

	// Run blocks until Ctrl+C and stops the worker gracefully.
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln("worker stopped with an error:", err)
	}
}
