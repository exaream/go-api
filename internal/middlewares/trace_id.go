package middlewares

import (
	"context"

	"github.com/google/uuid"
)

type traceIDKey struct{}

func NewTraceID() uuid.UUID {
	return uuid.Must(uuid.NewV7())
}

func SetTraceID(ctx context.Context, traceID uuid.UUID) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

func GetTraceID(ctx context.Context) uuid.UUID {
	v := ctx.Value(traceIDKey{})
	if id, ok := v.(uuid.UUID); ok {
		return id
	}

	return uuid.Nil
}
