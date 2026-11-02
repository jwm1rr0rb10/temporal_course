package determinism

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

// LimitedWorkflow обрабатывает элементы не больше чем по limit одновременно.
func LimitedWorkflow(ctx workflow.Context, items []string, limit int) ([]string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	})

	sem := workflow.NewSemaphore(ctx, int64(limit))
	wg := workflow.NewWaitGroup(ctx)
	results := make([]string, len(items))
	var firstErr error

	for i, item := range items {
		if err := sem.Acquire(ctx, 1); err != nil { // ждём свободный слот
			return nil, err
		}
		wg.Add(1)
		workflow.Go(ctx, func(ctx workflow.Context) {
			defer wg.Done()
			defer sem.Release(1)
			err := workflow.ExecuteActivity(ctx, Normalize, item).Get(ctx, &results[i])
			if err != nil && firstErr == nil {
				firstErr = err // корутины идут по очереди: гонки нет
			}
		})
	}
	wg.Wait(ctx)
	return results, firstErr
}
