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

// TrackedOrderWorkflow обновляет статус заказа в Search Attributes на каждом шаге,
// чтобы заказы можно было искать по статусу.
func TrackedOrderWorkflow(ctx workflow.Context, orderID string) error {
	for _, status := range []string{"paid", "packed", "shipped"} {
		if err := workflow.Sleep(ctx, time.Second); err != nil { // здесь были бы настоящие шаги
			return err
		}
		if err := workflow.UpsertTypedSearchAttributes(ctx, OrderStatusKey.ValueSet(status)); err != nil {
			return err
		}
	}
	return nil
}
