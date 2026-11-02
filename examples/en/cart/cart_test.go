package cart

import (
	"strings"
	"testing"
	"time"

	"go.temporal.io/sdk/testsuite"
)

func TestCartUpdatesAndCheckout(t *testing.T) {
	var s testsuite.WorkflowTestSuite
	env := s.NewTestWorkflowEnvironment()

	env.RegisterDelayedCallback(func() {
		env.UpdateWorkflow(UpdateAddItem, "add-book", &testsuite.TestUpdateCallback{
			OnAccept: func() {},
			OnReject: func(err error) { t.Errorf("book rejected: %v", err) },
			OnComplete: func(_ any, err error) {
				if err != nil {
					t.Errorf("book: %v", err)
				}
			},
		}, Item{SKU: "book", Qty: 1})
	}, time.Second)

	env.RegisterDelayedCallback(func() {
		env.UpdateWorkflow(UpdateAddItem, "add-zero", &testsuite.TestUpdateCallback{
			OnAccept: func() { t.Error("an update with Qty=0 must be rejected") },
			OnReject: func(err error) {
				if !strings.Contains(err.Error(), "greater than zero") {
					t.Errorf("unexpected error: %v", err)
				}
			},
			OnComplete: func(any, error) {},
		}, Item{SKU: "pen", Qty: 0})
	}, 2*time.Second)

	env.RegisterDelayedCallback(func() { env.SignalWorkflow(SignalCheckout, nil) }, 3*time.Second)

	env.ExecuteWorkflow(CartWorkflow)

	var cart Cart
	if err := env.GetWorkflowResult(&cart); err != nil {
		t.Fatal(err)
	}
	if !cart.CheckedOut || cart.Items["book"] != 1 || len(cart.Items) != 1 {
		t.Fatalf("cart: %+v", cart)
	}
}
