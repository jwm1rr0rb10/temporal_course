package patterns

import (
	"context"
	"errors"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

var errNotReady = errors.New("отчёт ещё не готов")

// ReportAPI — внешняя система, которая долго готовит отчёт.
type ReportAPI interface {
	IsReady(ctx context.Context, reportID string) (bool, error)
}

type PollActivities struct {
	API ReportAPI
}

// WaitReport возвращает ошибку, пока отчёт не готов: повторами займётся retry policy.
func (a *PollActivities) WaitReport(ctx context.Context, reportID string) error {
	ready, err := a.API.IsReady(ctx, reportID)
	if err != nil {
		return err
	}
	if !ready {
		return errNotReady
	}
	return nil
}

// PollReportWorkflow ждёт отчёт до суток, проверяя раз в минуту. Повторы activity
// не добавляют событий в историю: она остаётся короткой, сколько бы ни длился опрос.
func PollReportWorkflow(ctx workflow.Context, reportID string) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout:    10 * time.Second,
		ScheduleToCloseTimeout: 24 * time.Hour,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Minute,
			BackoffCoefficient: 1.0, // ровный интервал вместо экспоненты
			MaximumInterval:    time.Minute,
		},
	})
	var a *PollActivities
	return workflow.ExecuteActivity(ctx, a.WaitReport, reportID).Get(ctx, nil)
}
