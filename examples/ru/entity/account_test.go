package entity

import (
	"errors"
	"testing"
	"time"

	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/workflow"
)

func TestContinueAsNew(t *testing.T) {
	var s testsuite.WorkflowTestSuite
	env := s.NewTestWorkflowEnvironment()
	for i := 1; i <= 3; i++ {
		env.RegisterDelayedCallback(func() { env.SignalWorkflow(SignalDeposit, int64(100)) }, time.Duration(i)*time.Second)
	}
	var mid int64
	env.RegisterDelayedCallback(func() {
		v, err := env.QueryWorkflow(QueryBalance)
		if err == nil {
			_ = v.Get(&mid)
		}
	}, 2500*time.Millisecond)
	env.ExecuteWorkflow(AccountWorkflow, State{}, 3)
	var canErr *workflow.ContinueAsNewError
	if err := env.GetWorkflowError(); !errors.As(err, &canErr) || mid != 200 {
		t.Fatalf("err=%v mid=%d", err, mid)
	}
}
