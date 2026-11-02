package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"

	"tcourse/hello"
	"tcourse/secure"
)

func main() {
	key, err := secure.KeyFromEnv("TEMPORAL_ENC_KEY")
	if err != nil {
		log.Fatalln(err)
	}
	c, err := secure.Dial("key-2026-11", key)
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	run, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
		ID:        "secure-1",
		TaskQueue: secure.TaskQueue,
	}, hello.GreetingWorkflow, "Anna Ivanova")
	if err != nil {
		log.Fatalln(err)
	}
	var res string
	if err := run.Get(context.Background(), &res); err != nil {
		log.Fatalln(err)
	}
	log.Println("result:", res)
}
