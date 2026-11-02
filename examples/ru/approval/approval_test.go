package approval

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"go.temporal.io/sdk/testsuite"
)

func TestAutoRejectAfterThreeDays(t *testing.T) {
	var s testsuite.WorkflowTestSuite
	env := s.NewTestWorkflowEnvironment()
	env.RegisterActivity(Notify)
	// Ровно два напоминания: после первых и после вторых суток.
	env.OnActivity(Notify, mock.Anything, mock.Anything).Return(nil).Times(2)

	env.ExecuteWorkflow(ApprovalWorkflow, "req-1") // трое суток пролетают за миллисекунды

	var d Decision
	if err := env.GetWorkflowResult(&d); err != nil {
		t.Fatal(err)
	}
	if d.Approved || d.By != "auto-timeout" {
		t.Fatalf("ожидали автоотказ, получили %+v", d)
	}
	env.AssertExpectations(t)
}

func TestApprovedOnSecondDay(t *testing.T) {
	var s testsuite.WorkflowTestSuite
	env := s.NewTestWorkflowEnvironment()
	env.RegisterActivity(Notify)
	env.OnActivity(Notify, mock.Anything, mock.Anything).Return(nil).Once()

	// Сигнал «придёт» через 30 часов виртуального времени.
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow(SignalDecision, Decision{Approved: true, By: "anna"})
	}, 30*time.Hour)

	env.ExecuteWorkflow(ApprovalWorkflow, "req-2")

	var d Decision
	if err := env.GetWorkflowResult(&d); err != nil {
		t.Fatal(err)
	}
	if !d.Approved || d.By != "anna" {
		t.Fatalf("ожидали одобрение от anna, получили %+v", d)
	}
	env.AssertExpectations(t)
}
