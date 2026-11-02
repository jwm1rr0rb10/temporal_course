package main

import (
	"log"

	"go.temporal.io/sdk/worker"

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

	w := worker.New(c, secure.TaskQueue, worker.Options{})
	w.RegisterWorkflow(hello.GreetingWorkflow)
	w.RegisterActivity(hello.Greet)
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln(err)
	}
}
