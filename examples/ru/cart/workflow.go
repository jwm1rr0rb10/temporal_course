package cart

import (
	"errors"
	"fmt"

	"go.temporal.io/sdk/workflow"
)

const (
	TaskQueue      = "cart"
	UpdateAddItem  = "add-item"
	QueryGetCart   = "get-cart"
	SignalCheckout = "checkout"
)

type Item struct {
	SKU string
	Qty int
}

type Cart struct {
	Items      map[string]int
	CheckedOut bool
}

func CartWorkflow(ctx workflow.Context) (Cart, error) {
	cart := Cart{Items: map[string]int{}}

	// Query: только чтение, без блокировок и побочных эффектов.
	err := workflow.SetQueryHandler(ctx, QueryGetCart, func() (Cart, error) {
		return cart, nil
	})
	if err != nil {
		return cart, err
	}

	// Update: меняет состояние и возвращает результат вызывающему.
	err = workflow.SetUpdateHandlerWithOptions(ctx, UpdateAddItem,
		func(ctx workflow.Context, item Item) (int, error) {
			cart.Items[item.SKU] += item.Qty
			return len(cart.Items), nil // сколько разных товаров в корзине
		},
		workflow.UpdateHandlerOptions{
			// Validator: та же сигнатура, что у обработчика, но возвращает только error.
			// Отклонённый update не попадает в историю. Менять состояние здесь нельзя.
			Validator: func(ctx workflow.Context, item Item) error {
				if cart.CheckedOut {
					return errors.New("корзина уже оформлена")
				}
				if item.Qty <= 0 {
					return fmt.Errorf("количество должно быть больше нуля, получено %d", item.Qty)
				}
				return nil
			},
		})
	if err != nil {
		return cart, err
	}

	// Signal: ждём оформления. Пока ждём, updates и queries обрабатываются.
	workflow.GetSignalChannel(ctx, SignalCheckout).Receive(ctx, nil)
	cart.CheckedOut = true

	// Не завершаемся, пока не отработали уже принятые updates.
	err = workflow.Await(ctx, func() bool { return workflow.AllHandlersFinished(ctx) })
	return cart, err
}
