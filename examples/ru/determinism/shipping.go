package determinism

import (
	"context"
	"time"

	"go.temporal.io/sdk/workflow"
)

func Pack(ctx context.Context, orderID string) error   { return nil }
func Ship(ctx context.Context, orderID string) error   { return nil }
func Notify(ctx context.Context, orderID string) error { return nil }

func ShippingWorkflow(ctx workflow.Context, orderID string) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	})
	if err := workflow.ExecuteActivity(ctx, Pack, orderID).Get(ctx, nil); err != nil {
		return err
	}
	if err := workflow.Sleep(ctx, 30*time.Second); err != nil {
		return err
	}
	return workflow.ExecuteActivity(ctx, Ship, orderID).Get(ctx, nil)
}
