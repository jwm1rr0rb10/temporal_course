package payments

import (
	"context"
	"errors"

	"go.temporal.io/sdk/temporal"
)

type ChargeInput struct {
	OrderID     string
	AmountCents int64
	CardToken   string
}

type ChargeResult struct {
	PaymentID string
}

// Gateway is the payment provider: an HTTP client in production, a stub in tests.
type Gateway interface {
	Charge(ctx context.Context, idempotencyKey string, in ChargeInput) (paymentID string, err error)
}

// ErrDeclined is returned by the provider when the bank declines the card.
var ErrDeclined = errors.New("card declined")

// Activities holds the dependencies: one instance is registered in the worker.
type Activities struct {
	Gateway Gateway
}

func (a *Activities) Charge(ctx context.Context, in ChargeInput) (ChargeResult, error) {
	// The idempotency key is tied to the order, not the attempt: on a retry the
	// provider returns the payment it already created instead of charging twice.
	paymentID, err := a.Gateway.Charge(ctx, "charge-"+in.OrderID, in)
	switch {
	case errors.Is(err, ErrDeclined):
		// The error type is a contract with the workflow: it decides whether to retry.
		return ChargeResult{}, temporal.NewApplicationErrorWithCause("card declined", "CardDeclined", err)
	case err != nil:
		return ChargeResult{}, err // network, 5xx, timeout: retried by the retry policy
	}
	return ChargeResult{PaymentID: paymentID}, nil
}
