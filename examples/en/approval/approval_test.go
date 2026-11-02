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
	// Exactly two reminders: after the first and after the second day.
	env.OnActivity(Notify, mock.Anything, mock.Anything).Return(nil).Times(2)

	env.ExecuteWorkflow(ApprovalWorkflow, "req-1") // three days fly by in milliseconds

	var d Decision
	if err := env.GetWorkflowResult(&d); err != nil {
		t.Fatal(err)
	}
	if d.Approved || d.By != "auto-timeout" {
		t.Fatalf("expected an automatic rejection, got %+v", d)
	}
	env.AssertExpectations(t)
}

func TestApprovedOnSecondDay(t *testing.T) {
	var s testsuite.WorkflowTestSuite
	env := s.NewTestWorkflowEnvironment()
	env.RegisterActivity(Notify)
	env.OnActivity(Notify, mock.Anything, mock.Anything).Return(nil).Once()

	// The signal "arrives" after 30 hours of virtual time.
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow(SignalDecision, Decision{Approved: true, By: "anna"})
	}, 30*time.Hour)

	env.ExecuteWorkflow(ApprovalWorkflow, "req-2")

	var d Decision
	if err := env.GetWorkflowResult(&d); err != nil {
		t.Fatal(err)
	}
	if !d.Approved || d.By != "anna" {
		t.Fatalf("expected approval by anna, got %+v", d)
	}
	env.AssertExpectations(t)
}
