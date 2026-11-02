package main

import (
	"context"
	"log"

	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"

	"tcourse/cart"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()
	ctx := context.Background()
	workflowID := "cart-user-42"

	// Update-With-Start: добавить товар и создать корзину, если её ещё нет, — атомарно.
	addItem := func(item cart.Item) (int, error) {
		startOp := c.NewWithStartWorkflowOperation(client.StartWorkflowOptions{
			ID:                       workflowID,
			TaskQueue:                cart.TaskQueue,
			WorkflowIDConflictPolicy: enumspb.WORKFLOW_ID_CONFLICT_POLICY_USE_EXISTING,
		}, cart.CartWorkflow)

		handle, err := c.UpdateWithStartWorkflow(ctx, client.UpdateWithStartWorkflowOptions{
			StartWorkflowOperation: startOp,
			UpdateOptions: client.UpdateWorkflowOptions{
				UpdateName:   cart.UpdateAddItem,
				Args:         []any{item},
				WaitForStage: client.WorkflowUpdateStageCompleted,
			},
		})
		if err != nil {
			return 0, err
		}
		var distinct int
		err = handle.Get(ctx, &distinct)
		return distinct, err
	}

	n, err := addItem(cart.Item{SKU: "book", Qty: 1}) // корзины нет: создастся
	log.Println("book:", n, err)
	n, err = addItem(cart.Item{SKU: "pen", Qty: 3}) // корзина уже есть
	log.Println("pen:", n, err)
	_, err = addItem(cart.Item{SKU: "pen", Qty: 0}) // validator отклонит
	log.Println("pen с Qty=0:", err)

	// Query: прочитать состояние, ничего не меняя.
	resp, err := c.QueryWorkflow(ctx, workflowID, "", cart.QueryGetCart)
	if err != nil {
		log.Fatalln(err)
	}
	var state cart.Cart
	if err := resp.Get(&state); err != nil {
		log.Fatalln(err)
	}
	log.Println("в корзине:", state.Items)

	// Signal: оформить заказ. Ответа нет, только факт доставки.
	if err := c.SignalWorkflow(ctx, workflowID, "", cart.SignalCheckout, nil); err != nil {
		log.Fatalln(err)
	}

	var final cart.Cart
	if err := c.GetWorkflow(ctx, workflowID, "").Get(ctx, &final); err != nil {
		log.Fatalln(err)
	}
	log.Println("оформлено:", final.CheckedOut, final.Items)
}
