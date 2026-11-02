package inventory

import (
	"context"
	"fmt"
	"time"

	"github.com/nexus-rpc/sdk-go/nexus"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporalnexus"
	"go.temporal.io/sdk/workflow"
)

const TaskQueue = "inventory-nexus"

// A synchronous operation: a fast answer right from the handler.
var checkStock = nexus.NewSyncOperation(CheckStockOp.Name(),
	func(ctx context.Context, in StockInput, _ nexus.StartOperationOptions) (StockOutput, error) {
		return StockOutput{Available: in.Qty <= 10}, nil // a real warehouse query would go here
	})

// An asynchronous operation: starts a workflow in the inventory namespace and returns its result.
var reserve = temporalnexus.NewWorkflowRunOperation(ReserveOp.Name(), ReserveWorkflow,
	func(ctx context.Context, in ReserveInput, _ nexus.StartOperationOptions) (client.StartWorkflowOptions, error) {
		// An ID derived from the order is the idempotency key: a repeated call won't create a second reservation.
		return client.StartWorkflowOptions{ID: "reserve-" + in.OrderID}, nil
	})

func ReserveWorkflow(ctx workflow.Context, in ReserveInput) (ReserveOutput, error) {
	if err := workflow.Sleep(ctx, time.Second); err != nil { // the real reservation steps would go here
		return ReserveOutput{}, err
	}
	return ReserveOutput{ReservationID: fmt.Sprintf("res-%s-%s-%d", in.OrderID, in.SKU, in.Qty)}, nil
}

// NewService assembles the Nexus service from its operations.
func NewService() (*nexus.Service, error) {
	s := nexus.NewService(ServiceName)
	return s, s.Register(checkStock, reserve)
}
