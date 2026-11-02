package saga

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/testsuite"
)

func record(env *testsuite.TestWorkflowEnvironment) *[]string {
	var started []string
	env.SetOnActivityStartedListener(func(info *activity.Info, ctx context.Context, args converter.EncodedValues) {
		started = append(started, info.ActivityType.Name)
	})
	return &started
}

func TestCompensationsRunInReverse(t *testing.T) {
	var s testsuite.WorkflowTestSuite
	env := s.NewTestWorkflowEnvironment()
	for _, a := range []any{Charge, Refund, ReserveStock, ReleaseStock, Ship} {
		env.RegisterActivity(a)
	}
	env.OnActivity(Ship, mock.Anything, mock.Anything).Return(temporal.NewNonRetryableApplicationError("нет курьеров", "NoCourier", nil))
	started := record(env)
	env.ExecuteWorkflow(OrderWorkflow, Order{ID: "1"})
	var appErr *temporal.ApplicationError
	if err := env.GetWorkflowError(); !errors.As(err, &appErr) || appErr.Type() != "NoCourier" {
		t.Fatal(err)
	}
	if got := strings.Join(*started, ","); got != "Charge,ReserveStock,Ship,ReleaseStock,Refund" {
		t.Fatalf("order: %s", got)
	}
}

func TestHoldReservationCancel(t *testing.T) {
	var s testsuite.WorkflowTestSuite
	env := s.NewTestWorkflowEnvironment()
	env.RegisterActivity(ReserveStock)
	env.RegisterActivity(ReleaseStock)
	started := record(env)
	env.RegisterDelayedCallback(func() { env.CancelWorkflow() }, time.Minute)
	env.ExecuteWorkflow(HoldReservationWorkflow, Order{ID: "1"})
	if err := env.GetWorkflowError(); !temporal.IsCanceledError(err) || strings.Join(*started, ",") != "ReserveStock,ReleaseStock" {
		t.Fatalf("err=%v started=%v", err, *started)
	}
}
