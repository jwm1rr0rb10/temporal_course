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

// Gateway — платёжный провайдер. В проде это HTTP-клиент, в тестах — заглушка.
type Gateway interface {
	Charge(ctx context.Context, idempotencyKey string, in ChargeInput) (paymentID string, err error)
}

// ErrDeclined возвращает провайдер, когда банк отклонил карту.
var ErrDeclined = errors.New("card declined")

// Activities держит зависимости: один экземпляр регистрируется в воркере.
type Activities struct {
	Gateway Gateway
}

func (a *Activities) Charge(ctx context.Context, in ChargeInput) (ChargeResult, error) {
	// Ключ идемпотентности привязан к заказу, а не к попытке: при повторе
	// провайдер вернёт уже созданный платёж вместо второго списания.
	paymentID, err := a.Gateway.Charge(ctx, "charge-"+in.OrderID, in)
	switch {
	case errors.Is(err, ErrDeclined):
		// Тип ошибки — контракт с workflow: по нему решают, повторять или нет.
		return ChargeResult{}, temporal.NewApplicationErrorWithCause("карта отклонена", "CardDeclined", err)
	case err != nil:
		return ChargeResult{}, err // сеть, 5xx, таймаут: повторится по retry policy
	}
	return ChargeResult{PaymentID: paymentID}, nil
}
