package db

import (
	"context"
	"time"
)

func Timeout() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	return ctx
}
