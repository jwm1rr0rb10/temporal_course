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

// ParallelWorkflow запускает activity на каждый элемент параллельно
// и собирает результаты в исходном порядке.
func ParallelWorkflow(ctx workflow.Context, items []string) ([]string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	})

	futures := make([]workflow.Future, len(items))
	for i, item := range items {
		futures[i] = workflow.ExecuteActivity(ctx, Normalize, item) // команда уходит сразу
	}

	results := make([]string, len(items))
	for i, f := range futures {
		if err := f.Get(ctx, &results[i]); err != nil {
			return nil, err
		}
	}
	return results, nil
}
