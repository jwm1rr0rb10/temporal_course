package heartbeat

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

const TaskQueue = "import"

// ImportRows processes rows one by one and sends a heartbeat with the number of
// the processed row after each one. After a worker crash the next attempt
// continues from that point instead of starting over.
func ImportRows(ctx context.Context, total int) (int, error) {
	start := 0
	if activity.HasHeartbeatDetails(ctx) {
		var last int
		if err := activity.GetHeartbeatDetails(ctx, &last); err == nil {
			start = last + 1
		}
	}
	activity.GetLogger(ctx).Info("starting import", "from", start, "attempt", activity.GetInfo(ctx).Attempt)

	for i := start; i < total; i++ {
		select {
		case <-ctx.Done(): // cancellation and timeouts arrive through heartbeats
			return i, ctx.Err()
		case <-time.After(200 * time.Millisecond): // simulate processing a row
		}
		activity.RecordHeartbeat(ctx, i)
	}
	return total, nil
}

func ImportWorkflow(ctx workflow.Context, total int) (string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Minute, // all the work of one attempt
		HeartbeatTimeout:    5 * time.Second,  // how long it may stay silent
	})

	var done int
	if err := workflow.ExecuteActivity(ctx, ImportRows, total).Get(ctx, &done); err != nil {
		return "", err
	}
	return fmt.Sprintf("rows imported: %d", done), nil
}
