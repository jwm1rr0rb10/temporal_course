package main

import (
	"context"
	"log"
	"os"

	"go.temporal.io/sdk/client"

	"tcourse/hello"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("could not connect to Temporal:", err)
	}
	defer c.Close()

	name := "Temporal"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	run, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
		ID:        "greeting-" + name, // you choose the Workflow ID: it is the idempotency key
		TaskQueue: hello.TaskQueue,
	}, hello.GreetingWorkflow, name)
	if err != nil {
		log.Fatalln("could not start the workflow:", err)
	}
	log.Println("started", run.GetID(), run.GetRunID())

	var result string
	if err := run.Get(context.Background(), &result); err != nil { // wait for completion
		log.Fatalln("the workflow failed:", err)
	}
	log.Println("result:", result)
}
