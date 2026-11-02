package checkout

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"

	"tcourse/inventory"
)

const TaskQueue = "checkout"

type Order struct {
	ID  string
	SKU string
	Qty int
}

// CheckoutWorkflow живёт в namespace команды заказов и вызывает склад через Nexus,
// не зная ни namespace склада, ни его task queue, ни внутренних workflow.
func CheckoutWorkflow(ctx workflow.Context, o Order) (string, error) {
	inv := workflow.NewNexusClient(inventory.EndpointName, inventory.ServiceName)
	opts := workflow.NexusOperationOptions{ScheduleToCloseTimeout: 5 * time.Minute}

	var stock inventory.StockOutput
	in := inventory.StockInput{SKU: o.SKU, Qty: o.Qty}
	if err := inv.ExecuteOperation(ctx, inventory.CheckStockOp, in, opts).Get(ctx, &stock); err != nil {
		return "", err
	}
	if !stock.Available {
		return "", temporal.NewNonRetryableApplicationError("товара нет в нужном количестве", "OutOfStock", nil)
	}

	var res inventory.ReserveOutput
	req := inventory.ReserveInput{OrderID: o.ID, SKU: o.SKU, Qty: o.Qty}
	if err := inv.ExecuteOperation(ctx, inventory.ReserveOp, req, opts).Get(ctx, &res); err != nil {
		return "", err
	}
	return res.ReservationID, nil
}
