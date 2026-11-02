package saga

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// HoldReservationWorkflow holds a reservation for 15 minutes and releases it.
// If the workflow is canceled earlier, the reservation is released right away.
func HoldReservationWorkflow(ctx workflow.Context, o Order) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	})

	if err := workflow.ExecuteActivity(ctx, ReserveStock, o).Get(ctx, nil); err != nil {
		return err
	}

	err := workflow.Sleep(ctx, 15*time.Minute)
	if temporal.IsCanceledError(err) {
		// ctx is already canceled, and an activity in it would fail immediately.
		// Cleanup needs a "disconnected" context: it doesn't inherit the
		// cancellation but keeps the options.
		cleanupCtx, _ := workflow.NewDisconnectedContext(ctx)
		if cerr := workflow.ExecuteActivity(cleanupCtx, ReleaseStock, o).Get(cleanupCtx, nil); cerr != nil {
			return cerr
		}
		return err // the execution closes with the Canceled status
	}
	if err != nil {
		return err
	}
	return workflow.ExecuteActivity(ctx, ReleaseStock, o).Get(ctx, nil)
}
