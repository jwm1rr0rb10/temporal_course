package hello

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

const TaskQueue = "hello"

// Greet — activity: здесь можно ходить в сеть, базу и файлы.
func Greet(ctx context.Context, name string) (string, error) {
	activity.GetLogger(ctx).Info("Greet", "name", name)
	return fmt.Sprintf("Привет, %s!", name), nil
}

// GreetingWorkflow — workflow: только оркестрация, никакого ввода-вывода.
func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	})

	var greeting string
	if err := workflow.ExecuteActivity(ctx, Greet, name).Get(ctx, &greeting); err != nil {
		return "", err
	}
	return greeting, nil
}
