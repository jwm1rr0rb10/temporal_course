package versioning

import (
	"testing"

	"go.temporal.io/sdk/worker"
)

// The history of an execution started on the old code was saved like this:
//
//	temporal workflow show --workflow-id shipping-old --output json > versioning/testdata/shipping-v1.json
func TestReplayShippingV1History(t *testing.T) {
	replayer := worker.NewWorkflowReplayer()
	replayer.RegisterWorkflow(ShippingWorkflow)
	if err := replayer.ReplayWorkflowHistoryFromJSONFile(nil, "testdata/shipping-v1.json"); err != nil {
		t.Fatalf("the new code is incompatible with the old history: %v", err)
	}
}
