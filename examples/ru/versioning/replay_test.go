package versioning

import (
	"testing"

	"go.temporal.io/sdk/worker"
)

// История выполнения, начатого на старом коде, сохранена так:
//
//	temporal workflow show --workflow-id shipping-old --output json > versioning/testdata/shipping-v1.json
func TestReplayShippingV1History(t *testing.T) {
	replayer := worker.NewWorkflowReplayer()
	replayer.RegisterWorkflow(ShippingWorkflow)
	if err := replayer.ReplayWorkflowHistoryFromJSONFile(nil, "testdata/shipping-v1.json"); err != nil {
		t.Fatalf("новый код несовместим со старой историей: %v", err)
	}
}
