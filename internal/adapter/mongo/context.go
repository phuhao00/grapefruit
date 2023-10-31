package mongo

import (
	"context"
	"grapefruit/config"
	"time"
)

// Config struct contains extra configuration properties for the mgm package.
type Config struct {
	// Set to 10 second (10*time.Second) for example.
	CtxTimeout time.Duration
}

// NewCtx function creates and returns a new context with the specified timeout.
func NewCtx(timeout time.Duration) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), timeout)

	return ctx
}

// Ctx function creates and returns a new context with a default timeout value.
func Ctx() context.Context {
	return ctx()
}

func ctx() context.Context {
	return NewCtx(time.Duration(config.GetMongoConfig().CtxTimeout))
}
