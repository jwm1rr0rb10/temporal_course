package durable

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

const TaskQueue = "durable"

func Step(ctx context.Context, name string) (string, error) {
	activity.GetLogger(ctx).Info("running step", "step", name)
	return name + ": done", nil
}

// TwoStepWorkflow runs a step, sleeps for 30 seconds and runs a second step.
func TwoStepWorkflow(ctx workflow.Context) ([]string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	})
	logger := workflow.GetLogger(ctx)

	var first, second string
	if err := workflow.ExecuteActivity(ctx, Step, "first").Get(ctx, &first); err != nil {
		return nil, err
	}

	logger.Info("sleeping for 30 seconds: a good time to kill the worker")
	if err := workflow.Sleep(ctx, 30*time.Second); err != nil {
		return nil, err
	}

	if err := workflow.ExecuteActivity(ctx, Step, "second").Get(ctx, &second); err != nil {
		return nil, err
	}
	return []string{first, second}, nil
}
