package versioning

import (
	"time"

	"go.temporal.io/sdk/workflow"

	"tcourse/determinism"
)

// ShippingWorkflow — новая версия ShippingWorkflow из модуля 3.
// Имя функции то же, поэтому и тип workflow тот же.
func ShippingWorkflow(ctx workflow.Context, orderID string) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	})

	// Новые выполнения запишут в историю маркер версии 1 и пойдут по новой ветке.
	// Выполнения, начатые до изменения, маркера не имеют и получат DefaultVersion.
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
