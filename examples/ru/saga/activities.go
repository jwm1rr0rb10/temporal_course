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

// Заглушки: в реальном сервисе здесь вызовы платёжки, склада и доставки.
// Все действия идемпотентны по ID заказа.

func Charge(ctx context.Context, o Order) error {
	activity.GetLogger(ctx).Info("списание", "order", o.ID)
	return nil
}

func Refund(ctx context.Context, o Order) error {
	activity.GetLogger(ctx).Info("возврат", "order", o.ID)
	return nil
}

func ReserveStock(ctx context.Context, o Order) error {
	activity.GetLogger(ctx).Info("резерв", "order", o.ID)
	return nil
}

func ReleaseStock(ctx context.Context, o Order) error {
	activity.GetLogger(ctx).Info("снятие резерва", "order", o.ID)
	return nil
}

func Ship(ctx context.Context, o Order) error {
	activity.GetLogger(ctx).Info("отправка", "order", o.ID)
	return nil
}
