package orders

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

const TaskQueue = "orders"

var (
	CustomerIDKey  = temporal.NewSearchAttributeKeyKeyword("CustomerId")
	OrderStatusKey = temporal.NewSearchAttributeKeyKeyword("OrderStatus")
	AmountKey      = temporal.NewSearchAttributeKeyInt64("AmountCents")
)

// TrackedOrderWorkflow updates the order status in Search Attributes at every step,
// so orders can be searched by status.
func TrackedOrderWorkflow(ctx workflow.Context, orderID string) error {
	for _, status := range []string{"paid", "packed", "shipped"} {
		if err := workflow.Sleep(ctx, time.Second); err != nil { // the real steps would go here
			return err
		}
		if err := workflow.UpsertTypedSearchAttributes(ctx, OrderStatusKey.ValueSet(status)); err != nil {
			return err
		}
	}
	return nil
}
