package saga

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// compensations is a stack of undo actions.
type compensations []any

func (c *compensations) add(activity any) { *c = append(*c, activity) }

// run executes the compensations in reverse order and returns the original error.
func (c compensations) run(ctx workflow.Context, o Order, cause error) error {
	// A disconnected context: compensation runs even after the workflow is canceled.
	// Its own options without MaximumAttempts: compensation must eventually succeed.
	ctx, _ = workflow.NewDisconnectedContext(ctx)
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	})
	for i := len(c) - 1; i >= 0; i-- {
		if err := workflow.ExecuteActivity(ctx, c[i], o).Get(ctx, nil); err != nil {
			workflow.GetLogger(ctx).Error("compensation failed", "error", err)
		}
	}
	return cause
}

func OrderWorkflow(ctx workflow.Context, o Order) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts:        5,
			NonRetryableErrorTypes: []string{"CardDeclined", "OutOfStock"},
		},
	})

	var comp compensations

	// The compensation is added BEFORE the step: if Charge times out after the
	// card has already been charged, the refund still happens. That's why Refund
	// must be idempotent: refunding a payment that doesn't exist does nothing.
	comp.add(Refund)
	if err := workflow.ExecuteActivity(ctx, Charge, o).Get(ctx, nil); err != nil {
		return comp.run(ctx, o, err)
	}

	comp.add(ReleaseStock)
	if err := workflow.ExecuteActivity(ctx, ReserveStock, o).Get(ctx, nil); err != nil {
		return comp.run(ctx, o, err)
	}

	if err := workflow.ExecuteActivity(ctx, Ship, o).Get(ctx, nil); err != nil {
		return comp.run(ctx, o, err)
	}
	return nil
}
