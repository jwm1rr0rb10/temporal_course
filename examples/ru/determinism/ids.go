package determinism

import (
	"github.com/google/uuid"
	"go.temporal.io/sdk/workflow"
)

// NewOrderID генерирует ID один раз и записывает его в историю:
// при replay вернётся то же значение, а функция не вызовется.
func NewOrderID(ctx workflow.Context) (string, error) {
	var id string
	err := workflow.SideEffect(ctx, func(ctx workflow.Context) any {
		return uuid.NewString()
	}).Get(&id)
	return id, err
}
