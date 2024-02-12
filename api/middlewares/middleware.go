package middlewares

import (
	"context"
	"log/slog"
	"net/http"
)

func Apply(
	ctx context.Context,
	logger *slog.Logger,
	handler http.HandlerFunc,
	middlewareList []func(context.Context, *slog.Logger, http.HandlerFunc) http.HandlerFunc,
) http.HandlerFunc {
	for _, middleware := range middlewareList {
		handler = middleware(ctx, logger, handler)
	}

	return handler
}
