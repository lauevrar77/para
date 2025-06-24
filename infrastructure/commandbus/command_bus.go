package commandbus

import (
	"context"

	"github.com/go-dew/dew"
)

var CommandBus = dew.New()

type Bus interface {
	Register(handler any)
}

func NewContext(parent context.Context) context.Context {
	return dew.NewContext(parent, CommandBus)
}

func Dispatch[T dew.Action](ctx context.Context, action *T) (*T, error) {
	return dew.Dispatch(ctx, action)
}
