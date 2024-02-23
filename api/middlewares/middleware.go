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

func (m *Middleware) Apply(
	handler http.HandlerFunc,
	middlewareList []func(*slog.Logger, http.HandlerFunc) http.HandlerFunc,
) http.HandlerFunc {
	for _, middleware := range middlewareList {
		handler = middleware(m.logger, handler)
	}

	return handler
}
