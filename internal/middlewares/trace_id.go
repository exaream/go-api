package middlewares

import (
	"context"
	"sync"
)

var (
	mu         sync.Mutex
	internalID int = 1
)

type traceIDKey struct{}

func NewTraceID() int {
	var traceID int

	mu.Lock()
	traceID = internalID
	internalID++
	mu.Unlock()

	return traceID
}

func SetTraceID(ctx context.Context, traceID int) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

func GetTraceID(ctx context.Context) int {
	v := ctx.Value(traceIDKey{})
	if id, ok := v.(int); ok {
		return id
	}

	return 0
}
