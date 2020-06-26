package db

import (
	"context"
	"time"
)

func Timeout() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	return ctx, cancel
}
