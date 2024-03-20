package middlewares

import (
	"log/slog"
	"net/http"
)

type Middleware struct {
	logger *slog.Logger
}

func NewMiddleware(logger *slog.Logger) *Middleware {
	return &Middleware{
		logger: logger,
	}
}

func (m *Middleware) Chain(
	handler http.HandlerFunc,
	middlewareList []func(*slog.Logger, http.HandlerFunc) http.HandlerFunc,
) http.HandlerFunc {
	for i := range middlewareList {
		handler = middlewareList[len(middlewareList)-1-i](m.logger, handler)
	}

	return handler
}
