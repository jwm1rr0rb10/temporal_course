package patterns

import (
	"context"
	"errors"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

var errNotReady = errors.New("the report is not ready yet")

// ReportAPI is an external system that takes a long time to prepare a report.
type ReportAPI interface {
	IsReady(ctx context.Context, reportID string) (bool, error)
}

type PollActivities struct {
	API ReportAPI
}

// WaitReport returns an error until the report is ready: the retry policy handles the repeats.
func (a *PollActivities) WaitReport(ctx context.Context, reportID string) error {
	ready, err := a.API.IsReady(ctx, reportID)
	if err != nil {
		return err
	}
	if !ready {
		return errNotReady
	}
	return nil
}

// PollReportWorkflow waits up to a day for the report, checking once a minute. Activity
// retries add no events to the history: it stays short no matter how long the polling lasts.
func PollReportWorkflow(ctx workflow.Context, reportID string) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout:    10 * time.Second,
		ScheduleToCloseTimeout: 24 * time.Hour,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Minute,
			BackoffCoefficient: 1.0, // a fixed interval instead of exponential backoff
			MaximumInterval:    time.Minute,
		},
	})
	var a *PollActivities
	return workflow.ExecuteActivity(ctx, a.WaitReport, reportID).Get(ctx, nil)
}
