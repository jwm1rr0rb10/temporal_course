package payments

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func PaymentWorkflow(ctx workflow.Context, in ChargeInput) (ChargeResult, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout:    10 * time.Second, // одна попытка
		ScheduleToCloseTimeout: 10 * time.Minute, // все попытки вместе
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:        time.Second,
			BackoffCoefficient:     2.0,
			MaximumInterval:        time.Minute,
			MaximumAttempts:        0, // 0 — без ограничения, общий срок держит ScheduleToClose
			NonRetryableErrorTypes: []string{"CardDeclined"},
		},
	})

	var a *Activities // nil-указатель нужен только для ссылки на метод
	var res ChargeResult
	err := workflow.ExecuteActivity(ctx, a.Charge, in).Get(ctx, &res)
	return res, err
}
