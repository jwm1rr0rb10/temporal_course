package approval

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

const SignalDecision = "decision"

type Decision struct {
	Approved bool
	By       string
}

func Notify(ctx context.Context, msg string) error {
	activity.GetLogger(ctx).Info(msg)
	return nil
}

// ApprovalWorkflow waits for a decision: it sends a reminder every day and,
// after three days without an answer, rejects the request automatically.
func ApprovalWorkflow(ctx workflow.Context, requestID string) (Decision, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	})
	decisions := workflow.GetSignalChannel(ctx, SignalDecision)

	for day := 1; day <= 3; day++ {
		var d Decision
		received := false

		timerCtx, cancelTimer := workflow.WithCancel(ctx)
		timer := workflow.NewTimer(timerCtx, 24*time.Hour)

		sel := workflow.NewSelector(ctx)
		sel.AddReceive(decisions, func(c workflow.ReceiveChannel, _ bool) {
			c.Receive(ctx, &d)
			received = true
		})
		sel.AddFuture(timer, func(workflow.Future) {})
		sel.Select(ctx)
		cancelTimer() // the timer is no longer needed: don't let it fire for nothing

		if received {
			return d, nil
		}
		if day < 3 {
			msg := fmt.Sprintf("request %s has been waiting for a decision for %d day(s)", requestID, day)
			if err := workflow.ExecuteActivity(ctx, Notify, msg).Get(ctx, nil); err != nil {
				return Decision{}, err
			}
		}
	}
	return Decision{Approved: false, By: "auto-timeout"}, nil
}
