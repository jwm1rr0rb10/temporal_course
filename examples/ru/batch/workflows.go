package batch

import (
	"context"
	"fmt"
	"time"

	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

func ProcessItem(ctx context.Context, item string) error {
	activity.GetLogger(ctx).Debug("обработка", "item", item)
	return nil
}

// ChunkWorkflow обрабатывает одну пачку: своя история, свои лимиты.
func ChunkWorkflow(ctx workflow.Context, items []string) (int, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	})
	futures := make([]workflow.Future, len(items))
	for i, item := range items {
		futures[i] = workflow.ExecuteActivity(ctx, ProcessItem, item)
	}
	for _, f := range futures {
		if err := f.Get(ctx, nil); err != nil {
			return 0, err
		}
	}
	return len(items), nil
}

// BatchWorkflow режет список на пачки и обрабатывает каждую в child workflow.
func BatchWorkflow(ctx workflow.Context, items []string, chunkSize int) (int, error) {
	parentID := workflow.GetInfo(ctx).WorkflowExecution.ID

	var children []workflow.ChildWorkflowFuture
	for start := 0; start < len(items); start += chunkSize {
		end := min(start+chunkSize, len(items))
		cctx := workflow.WithChildOptions(ctx, workflow.ChildWorkflowOptions{
			WorkflowID:        fmt.Sprintf("%s-chunk-%d", parentID, start/chunkSize), // детерминированный ID
			ParentClosePolicy: enumspb.PARENT_CLOSE_POLICY_TERMINATE,
		})
		children = append(children, workflow.ExecuteChildWorkflow(cctx, ChunkWorkflow, items[start:end]))
	}

	total := 0
	for _, child := range children {
		var n int
		if err := child.Get(ctx, &n); err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}
