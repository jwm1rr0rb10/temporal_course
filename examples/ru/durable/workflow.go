package durable

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

const TaskQueue = "durable"

func Step(ctx context.Context, name string) (string, error) {
	activity.GetLogger(ctx).Info("выполняю шаг", "step", name)
	return name + ": готово", nil
}

// TwoStepWorkflow делает шаг, спит 30 секунд и делает второй шаг.
func TwoStepWorkflow(ctx workflow.Context) ([]string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	})
	logger := workflow.GetLogger(ctx)

	var first, second string
	if err := workflow.ExecuteActivity(ctx, Step, "первый").Get(ctx, &first); err != nil {
		return nil, err
	}

	logger.Info("засыпаю на 30 секунд: самое время убить воркер")
	if err := workflow.Sleep(ctx, 30*time.Second); err != nil {
		return nil, err
	}

	if err := workflow.ExecuteActivity(ctx, Step, "второй").Get(ctx, &second); err != nil {
		return nil, err
	}
	return []string{first, second}, nil
}
