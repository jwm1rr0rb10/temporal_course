package determinism

import (
	"testing"

	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/workflow"
)

func TestParallelAndLimited(t *testing.T) {
	var s testsuite.WorkflowTestSuite
	for _, tc := range []struct {
		name string
		run  func(env *testsuite.TestWorkflowEnvironment)
	}{
		{"parallel", func(env *testsuite.TestWorkflowEnvironment) {
			env.ExecuteWorkflow(ParallelWorkflow, []string{" a", "b ", "c"})
		}},
		{"limited", func(env *testsuite.TestWorkflowEnvironment) {
			env.ExecuteWorkflow(LimitedWorkflow, []string{" a", "b ", "c"}, 2)
		}},
	} {
		env := s.NewTestWorkflowEnvironment()
		env.RegisterActivity(Normalize)
		tc.run(env)
		var res []string
		if err := env.GetWorkflowError(); err != nil {
			t.Fatal(tc.name, err)
		}
		_ = env.GetWorkflowResult(&res)
		if len(res) != 3 || res[0] != "A" || res[1] != "B" || res[2] != "C" {
			t.Fatalf("%s: %v", tc.name, res)
		}
	}
}

func idWorkflow(ctx workflow.Context) (string, error) { return NewOrderID(ctx) }

func TestNewOrderID(t *testing.T) {
	var s testsuite.WorkflowTestSuite
	env := s.NewTestWorkflowEnvironment()
	env.ExecuteWorkflow(idWorkflow)
	var id string
	if err := env.GetWorkflowResult(&id); err != nil || len(id) != 36 {
		t.Fatalf("id=%q err=%v", id, err)
	}
}
