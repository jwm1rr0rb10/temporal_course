package determinism

import (
	"context"
	"strings"
	"time"

	"go.temporal.io/sdk/workflow"
)

func Normalize(ctx context.Context, s string) (string, error) {
	return strings.ToUpper(strings.TrimSpace(s)), nil
}

// ParallelWorkflow runs an activity for every item in parallel
// and collects the results in the original order.
func ParallelWorkflow(ctx workflow.Context, items []string) ([]string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	})

	futures := make([]workflow.Future, len(items))
	for i, item := range items {
		futures[i] = workflow.ExecuteActivity(ctx, Normalize, item) // the command goes out immediately
	}

	results := make([]string, len(items))
	for i, f := range futures {
		if err := f.Get(ctx, &results[i]); err != nil {
			return nil, err
		}
	}
	return results, nil
}
