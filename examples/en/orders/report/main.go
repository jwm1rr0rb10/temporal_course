package main

import (
	"context"
	"log"

	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()
	ctx := context.Background()

	resp, err := c.ListWorkflow(ctx, &workflowservice.ListWorkflowExecutionsRequest{
		Query: `CustomerId = "c-42" AND AmountCents > 1000`,
	})
	if err != nil {
		log.Fatalln(err)
	}
	for _, e := range resp.GetExecutions() { // more than one page — use NextPageToken
		log.Println(e.GetExecution().GetWorkflowId(), e.GetStatus())
	}

	cnt, err := c.CountWorkflow(ctx, &workflowservice.CountWorkflowExecutionsRequest{
		Query: `OrderStatus = "shipped"`,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("orders shipped:", cnt.GetCount())
}
