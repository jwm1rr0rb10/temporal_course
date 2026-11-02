package saga

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// compensations — стек отменяющих действий.
type compensations []any

func (c *compensations) add(activity any) { *c = append(*c, activity) }

// run выполняет компенсации в обратном порядке и возвращает исходную ошибку.
func (c compensations) run(ctx workflow.Context, o Order, cause error) error {
	// Отсоединённый контекст: компенсации отработают и после отмены workflow.
	// Свои опции без MaximumAttempts: компенсация обязана в итоге пройти.
	ctx, _ = workflow.NewDisconnectedContext(ctx)
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	})
	for i := len(c) - 1; i >= 0; i-- {
		if err := workflow.ExecuteActivity(ctx, c[i], o).Get(ctx, nil); err != nil {
			workflow.GetLogger(ctx).Error("компенсация не выполнилась", "error", err)
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

	// Компенсацию добавляем ДО шага: если Charge упадёт по таймауту уже после
	// списания, возврат всё равно выполнится. Поэтому Refund обязан быть
	// идемпотентным: возврат несуществующего платежа ничего не делает.
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
