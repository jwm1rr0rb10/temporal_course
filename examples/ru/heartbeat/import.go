package heartbeat

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

const TaskQueue = "import"

// ImportRows обрабатывает строки по одной и после каждой отправляет heartbeat
// с номером обработанной строки. После падения воркера следующая попытка
// продолжит с этого места, а не с начала.
func ImportRows(ctx context.Context, total int) (int, error) {
	start := 0
	if activity.HasHeartbeatDetails(ctx) {
		var last int
		if err := activity.GetHeartbeatDetails(ctx, &last); err == nil {
			start = last + 1
		}
	}
	activity.GetLogger(ctx).Info("начинаю импорт", "from", start, "attempt", activity.GetInfo(ctx).Attempt)

	for i := start; i < total; i++ {
		select {
		case <-ctx.Done(): // отмена и таймауты приходят через heartbeat
			return i, ctx.Err()
		case <-time.After(200 * time.Millisecond): // имитация обработки строки
		}
		activity.RecordHeartbeat(ctx, i)
	}
	return total, nil
}

func ImportWorkflow(ctx workflow.Context, total int) (string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Minute, // вся работа одной попытки
		HeartbeatTimeout:    5 * time.Second,  // сколько можно молчать
	})

	var done int
	if err := workflow.ExecuteActivity(ctx, ImportRows, total).Get(ctx, &done); err != nil {
		return "", err
	}
	return fmt.Sprintf("импортировано строк: %d", done), nil
}
