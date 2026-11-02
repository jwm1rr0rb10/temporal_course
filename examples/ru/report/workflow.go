package report

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

const TaskQueue = "reports"

func BuildReport(ctx context.Context) (string, error) {
	activity.GetLogger(ctx).Info("строю отчёт")
	return "report.csv", nil
}

func ReportWorkflow(ctx workflow.Context) (string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Minute,
	})
	var file string
	err := workflow.ExecuteActivity(ctx, BuildReport).Get(ctx, &file)
	return file, err
}
