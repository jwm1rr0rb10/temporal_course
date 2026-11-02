package inventory

import "github.com/nexus-rpc/sdk-go/nexus"

const (
	ServiceName  = "inventory"
	EndpointName = "inventory" // имя endpoint, которое заводит платформа
)

type StockInput struct {
	SKU string
	Qty int
}

type StockOutput struct {
	Available bool
}

type ReserveInput struct {
	OrderID string
	SKU     string
	Qty     int
}

type ReserveOutput struct {
	ReservationID string
}

// Типизированные ссылки на операции: вызывающая сторона не ошибётся в имени и типах.
var (
	CheckStockOp = nexus.NewOperationReference[StockInput, StockOutput]("check-stock")
	ReserveOp    = nexus.NewOperationReference[ReserveInput, ReserveOutput]("reserve")
)
