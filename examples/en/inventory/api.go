package inventory

import "github.com/nexus-rpc/sdk-go/nexus"

const (
	ServiceName  = "inventory"
	EndpointName = "inventory" // the endpoint name the platform sets up
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

// Typed operation references: the caller can't get the name or the types wrong.
var (
	CheckStockOp = nexus.NewOperationReference[StockInput, StockOutput]("check-stock")
	ReserveOp    = nexus.NewOperationReference[ReserveInput, ReserveOutput]("reserve")
)
