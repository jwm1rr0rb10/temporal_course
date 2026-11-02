package entity

import (
	"go.temporal.io/sdk/workflow"
)

const (
	SignalDeposit = "deposit"
	QueryBalance  = "balance"
)

type State struct {
	Balance int64
}

// AccountWorkflow — «вечный» workflow счёта. После maxPerRun пополнений
// или по подсказке сервера он продолжает себя новым run с тем же Workflow ID.
func AccountWorkflow(ctx workflow.Context, state State, maxPerRun int) error {
	err := workflow.SetQueryHandler(ctx, QueryBalance, func() (int64, error) {
		return state.Balance, nil
	})
	if err != nil {
		return err
	}
	deposits := workflow.GetSignalChannel(ctx, SignalDeposit)

	for processed := 0; processed < maxPerRun && !workflow.GetInfo(ctx).GetContinueAsNewSuggested(); processed++ {
		var amount int64
		deposits.Receive(ctx, &amount)
		state.Balance += amount
	}

	// Сигналы, которые уже пришли, забираем до Continue-As-New: иначе они пропадут.
	for {
		var amount int64
		if !deposits.ReceiveAsync(&amount) {
			break
		}
		state.Balance += amount
	}
	// Если у workflow есть обработчики updates, ждём их завершения.
	if err := workflow.Await(ctx, func() bool { return workflow.AllHandlersFinished(ctx) }); err != nil {
		return err
	}
	return workflow.NewContinueAsNewError(ctx, AccountWorkflow, state, maxPerRun)
}
