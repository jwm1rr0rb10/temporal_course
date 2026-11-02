package saga

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// HoldReservationWorkflow держит резерв 15 минут и снимает его.
// Если workflow отменили раньше, резерв снимается сразу.
func HoldReservationWorkflow(ctx workflow.Context, o Order) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	})

	if err := workflow.ExecuteActivity(ctx, ReserveStock, o).Get(ctx, nil); err != nil {
		return err
	}

	err := workflow.Sleep(ctx, 15*time.Minute)
	if temporal.IsCanceledError(err) {
		// ctx уже отменён, и activity в нём сразу упадёт. Для уборки нужен
		// «отсоединённый» контекст: он не наследует отмену, но хранит опции.
		cleanupCtx, _ := workflow.NewDisconnectedContext(ctx)
		if cerr := workflow.ExecuteActivity(cleanupCtx, ReleaseStock, o).Get(cleanupCtx, nil); cerr != nil {
			return cerr
		}
		return err // выполнение закроется со статусом Canceled
	}
	if err != nil {
		return err
	}
	return workflow.ExecuteActivity(ctx, ReleaseStock, o).Get(ctx, nil)
}
