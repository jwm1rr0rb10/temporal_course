package versioning

import (
	"time"

	"go.temporal.io/sdk/workflow"

	"tcourse/determinism"
)

// ShippingWorkflow is a new version of ShippingWorkflow from Module 3.
// The function name is the same, so the workflow type is the same too.
func ShippingWorkflow(ctx workflow.Context, orderID string) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	})

	// New executions record a version 1 marker in the history and take the new branch.
	// Executions started before the change have no marker and get DefaultVersion.
	if v := workflow.GetVersion(ctx, "notify-before-pack", workflow.DefaultVersion, 1); v == 1 {
		if err := workflow.ExecuteActivity(ctx, determinism.Notify, orderID).Get(ctx, nil); err != nil {
			return err
		}
	}

	if err := workflow.ExecuteActivity(ctx, determinism.Pack, orderID).Get(ctx, nil); err != nil {
		return err
	}
	if err := workflow.Sleep(ctx, 30*time.Second); err != nil {
		return err
	}
	return workflow.ExecuteActivity(ctx, determinism.Ship, orderID).Get(ctx, nil)
}
