package hello

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

const TaskQueue = "hello"

// Greet is an activity: it may call the network, databases and files.
func Greet(ctx context.Context, name string) (string, error) {
	activity.GetLogger(ctx).Info("Greet", "name", name)
	return fmt.Sprintf("Hello, %s!", name), nil
}

// GreetingWorkflow is a workflow: orchestration only, no I/O.
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
