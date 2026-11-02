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

// AccountWorkflow is an "eternal" account workflow. After maxPerRun deposits,
// or when the server suggests it, it continues as a new run with the same Workflow ID.
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

	// Drain the signals that have already arrived before Continue-As-New, or they'll be lost.
	for {
		var amount int64
		if !deposits.ReceiveAsync(&amount) {
			break
		}
		state.Balance += amount
	}
	// If the workflow has update handlers, wait for them to finish.
	if err := workflow.Await(ctx, func() bool { return workflow.AllHandlersFinished(ctx) }); err != nil {
		return err
	}
	return workflow.NewContinueAsNewError(ctx, AccountWorkflow, state, maxPerRun)
}
