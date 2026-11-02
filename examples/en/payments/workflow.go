package payments

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func PaymentWorkflow(ctx workflow.Context, in ChargeInput) (ChargeResult, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout:    10 * time.Second, // one attempt
		ScheduleToCloseTimeout: 10 * time.Minute, // all attempts together
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:        time.Second,
			BackoffCoefficient:     2.0,
			MaximumInterval:        time.Minute,
			MaximumAttempts:        0, // 0 means unlimited; ScheduleToClose bounds the total time
			NonRetryableErrorTypes: []string{"CardDeclined"},
		},
	})

	var a *Activities // a nil pointer, only needed to reference the method
	var res ChargeResult
	err := workflow.ExecuteActivity(ctx, a.Charge, in).Get(ctx, &res)
	return res, err
}
