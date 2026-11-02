package saga

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/testsuite"
)

func TestShipFailureCompensatesInReverse(t *testing.T) {
	var s testsuite.WorkflowTestSuite
	env := s.NewTestWorkflowEnvironment()
	for _, a := range []any{Charge, Refund, ReserveStock, ReleaseStock, Ship} {
		env.RegisterActivity(a)
	}
	// Ship fails for good: no couriers.
	env.OnActivity(Ship, mock.Anything, mock.Anything).
		Return(temporal.NewNonRetryableApplicationError("no couriers", "NoCourier", nil))

	var started []string
	env.SetOnActivityStartedListener(func(info *activity.Info, _ context.Context, _ converter.EncodedValues) {
		started = append(started, info.ActivityType.Name)
	})

	env.ExecuteWorkflow(OrderWorkflow, Order{ID: "o-1"})

	var appErr *temporal.ApplicationError
	if err := env.GetWorkflowError(); !errors.As(err, &appErr) || appErr.Type() != "NoCourier" {
		t.Fatalf("expected a NoCourier error, got %v", err)
	}
	if got := strings.Join(started, " → "); got != "Charge → ReserveStock → Ship → ReleaseStock → Refund" {
		t.Fatalf("activity order: %s", got)
	}
}
