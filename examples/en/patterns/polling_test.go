package patterns

import (
	"context"
	"testing"

	"go.temporal.io/sdk/testsuite"
)

type fakeAPI struct{ calls, readyAfter int }

func (f *fakeAPI) IsReady(context.Context, string) (bool, error) {
	f.calls++
	return f.calls >= f.readyAfter, nil
}

func TestPollUntilReady(t *testing.T) {
	var s testsuite.WorkflowTestSuite
	env := s.NewTestWorkflowEnvironment()
	api := &fakeAPI{readyAfter: 4}
	env.RegisterActivity(&PollActivities{API: api})
	env.ExecuteWorkflow(PollReportWorkflow, "r-1")
	if err := env.GetWorkflowError(); err != nil || api.calls != 4 {
		t.Fatalf("err=%v calls=%d", err, api.calls)
	}
}
