package saga

import (
	"context"

	"go.temporal.io/sdk/activity"
)

type Order struct {
	ID          string
	AmountCents int64
	SKU         string
	Qty         int
}

// Stubs: a real service would call the payment provider, the warehouse and the carrier here.
// Every action is idempotent by order ID.

func Charge(ctx context.Context, o Order) error {
	activity.GetLogger(ctx).Info("charge", "order", o.ID)
	return nil
}

func Refund(ctx context.Context, o Order) error {
	activity.GetLogger(ctx).Info("refund", "order", o.ID)
	return nil
}

func ReserveStock(ctx context.Context, o Order) error {
	activity.GetLogger(ctx).Info("reserve stock", "order", o.ID)
	return nil
}

func ReleaseStock(ctx context.Context, o Order) error {
	activity.GetLogger(ctx).Info("release stock", "order", o.ID)
	return nil
}

func Ship(ctx context.Context, o Order) error {
	activity.GetLogger(ctx).Info("ship", "order", o.ID)
	return nil
}
