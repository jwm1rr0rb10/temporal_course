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

// Синхронная операция: быстрый ответ прямо из обработчика.
var checkStock = nexus.NewSyncOperation(CheckStockOp.Name(),
	func(ctx context.Context, in StockInput, _ nexus.StartOperationOptions) (StockOutput, error) {
		return StockOutput{Available: in.Qty <= 10}, nil // здесь был бы запрос к складу
	})

// Асинхронная операция: запускает workflow в namespace склада и отдаёт его результат.
var reserve = temporalnexus.NewWorkflowRunOperation(ReserveOp.Name(), ReserveWorkflow,
	func(ctx context.Context, in ReserveInput, _ nexus.StartOperationOptions) (client.StartWorkflowOptions, error) {
		// ID от заказа — ключ идемпотентности: повтор вызова не создаст второй резерв.
		return client.StartWorkflowOptions{ID: "reserve-" + in.OrderID}, nil
	})

func ReserveWorkflow(ctx workflow.Context, in ReserveInput) (ReserveOutput, error) {
	if err := workflow.Sleep(ctx, time.Second); err != nil { // здесь были бы шаги резервирования
		return ReserveOutput{}, err
	}
	return ReserveOutput{ReservationID: fmt.Sprintf("res-%s-%s-%d", in.OrderID, in.SKU, in.Qty)}, nil
}

// NewService собирает Nexus-сервис из операций.
func NewService() (*nexus.Service, error) {
	s := nexus.NewService(ServiceName)
	return s, s.Register(checkStock, reserve)
}
