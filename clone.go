package appctx

import (
	"context"
)

// Clone creates a new Context object but inherits values from the base context object.
func Clone(base context.Context) context.Context {
	return &clonedCtx{
		Context: context.Background(),
		base:    base,
	}
}

type clonedCtx struct {
	context.Context
	base context.Context
}

func (c *clonedCtx) Value(key interface{}) interface{} {
	if v := c.Context.Value(key); v != nil {
		return v
	}
	return c.base.Value(key)
}
