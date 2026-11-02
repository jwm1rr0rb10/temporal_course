package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"tcourse/hello"
)

func main() {
	// Пустые Options: localhost:7233 и namespace "default".
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("не удалось подключиться к Temporal:", err)
	}
	defer c.Close()

	w := worker.New(c, hello.TaskQueue, worker.Options{})
	w.RegisterWorkflow(hello.GreetingWorkflow)
	w.RegisterActivity(hello.Greet)

	// Run блокируется до Ctrl+C и корректно останавливает воркер.
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln("воркер остановился с ошибкой:", err)
	}
}
