package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"

	"tcourse/orders"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	for _, o := range []struct {
		id, customer string
		amount       int64
	}{
		{"order-1", "c-42", 4990},
		{"order-2", "c-42", 15000},
		{"order-3", "c-7", 990},
	} {
		_, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
			ID:        o.id,
			TaskQueue: orders.TaskQueue,
			TypedSearchAttributes: temporal.NewSearchAttributes(
				orders.CustomerIDKey.ValueSet(o.customer),
				orders.OrderStatusKey.ValueSet("new"),
				orders.AmountKey.ValueSet(o.amount),
			),
		}, orders.TrackedOrderWorkflow, o.id)
		if err != nil {
			log.Fatalln(err)
		}
	}
	log.Println("started three orders")
}
