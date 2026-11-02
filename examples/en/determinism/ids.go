package determinism

import (
	"github.com/google/uuid"
	"go.temporal.io/sdk/workflow"
)

// NewOrderID generates an ID once and records it in the history:
// during replay the same value comes back and the function isn't called.
func NewOrderID(ctx workflow.Context) (string, error) {
	var id string
	err := workflow.SideEffect(ctx, func(ctx workflow.Context) any {
		return uuid.NewString()
	}).Get(&id)
	return id, err
}
