package cart

import (
	"errors"
	"fmt"

	"go.temporal.io/sdk/workflow"
)

const (
	TaskQueue      = "cart"
	UpdateAddItem  = "add-item"
	QueryGetCart   = "get-cart"
	SignalCheckout = "checkout"
)

type Item struct {
	SKU string
	Qty int
}

type Cart struct {
	Items      map[string]int
	CheckedOut bool
}

func CartWorkflow(ctx workflow.Context) (Cart, error) {
	cart := Cart{Items: map[string]int{}}

	// Query: read-only, no blocking and no side effects.
	err := workflow.SetQueryHandler(ctx, QueryGetCart, func() (Cart, error) {
		return cart, nil
	})
	if err != nil {
		return cart, err
	}

	// Update: changes state and returns a result to the caller.
	err = workflow.SetUpdateHandlerWithOptions(ctx, UpdateAddItem,
		func(ctx workflow.Context, item Item) (int, error) {
			cart.Items[item.SKU] += item.Qty
			return len(cart.Items), nil // how many distinct items are in the cart
		},
		workflow.UpdateHandlerOptions{
			// Validator: the same signature as the handler, but it returns only an error.
			// A rejected update never reaches the history. Don't change state here.
			Validator: func(ctx workflow.Context, item Item) error {
				if cart.CheckedOut {
					return errors.New("the cart is already checked out")
				}
				if item.Qty <= 0 {
					return fmt.Errorf("quantity must be greater than zero, got %d", item.Qty)
				}
				return nil
			},
		})
	if err != nil {
		return cart, err
	}

	// Signal: wait for checkout. Updates and queries are handled while we wait.
	workflow.GetSignalChannel(ctx, SignalCheckout).Receive(ctx, nil)
	cart.CheckedOut = true

	// Don't complete until the updates that were already accepted have finished.
	err = workflow.Await(ctx, func() bool { return workflow.AllHandlersFinished(ctx) })
	return cart, err
}
