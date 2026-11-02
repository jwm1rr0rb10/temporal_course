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
		log.Fatalln("не удалось подключиться к Temporal:", err)
	}
	defer c.Close()

	name := "Temporal"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	run, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
		ID:        "greeting-" + name, // Workflow ID задаёшь ты: это ключ идемпотентности
		TaskQueue: hello.TaskQueue,
	}, hello.GreetingWorkflow, name)
	if err != nil {
		log.Fatalln("не удалось запустить workflow:", err)
	}
	log.Println("запущен", run.GetID(), run.GetRunID())

	var result string
	if err := run.Get(context.Background(), &result); err != nil { // ждём завершения
		log.Fatalln("workflow завершился с ошибкой:", err)
	}
	log.Println("результат:", result)
}
