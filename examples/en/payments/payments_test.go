package payments

import (
	"context"
	"errors"
	"testing"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/testsuite"
)

type fakeGateway struct {
	failures, calls int
	declined        bool
}

func (g *fakeGateway) Charge(ctx context.Context, key string, in ChargeInput) (string, error) {
	g.calls++
	if g.declined {
		return "", ErrDeclined
	}
	if g.calls <= g.failures {
		return "", errors.New("503 from provider")
	}
	return "pay-" + in.OrderID + "-" + key, nil
}

func TestTransientErrorsAreRetried(t *testing.T) {
	var s testsuite.WorkflowTestSuite
	env := s.NewTestWorkflowEnvironment()
	gw := &fakeGateway{failures: 2}
	env.RegisterActivity(&Activities{Gateway: gw})
	env.ExecuteWorkflow(PaymentWorkflow, ChargeInput{OrderID: "42", AmountCents: 1000})
	if err := env.GetWorkflowError(); err != nil {
		t.Fatal(err)
	}
	var res ChargeResult
	_ = env.GetWorkflowResult(&res)
	if res.PaymentID != "pay-42-charge-42" || gw.calls != 3 {
		t.Fatalf("res=%+v calls=%d", res, gw.calls)
	}
}

func TestDeclinedIsNotRetried(t *testing.T) {
	var s testsuite.WorkflowTestSuite
	env := s.NewTestWorkflowEnvironment()
	gw := &fakeGateway{declined: true}
	env.RegisterActivity(&Activities{Gateway: gw})
	env.ExecuteWorkflow(PaymentWorkflow, ChargeInput{OrderID: "42"})
	var appErr *temporal.ApplicationError
	if err := env.GetWorkflowError(); !errors.As(err, &appErr) || appErr.Type() != "CardDeclined" || gw.calls != 1 {
		t.Fatalf("err=%v calls=%d", err, gw.calls)
	}
}
