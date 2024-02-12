package middlewares

import (
	"context"
	"log/slog"
	"net/http"
)

type responseLoggingWriter struct {
	http.ResponseWriter
	code int
}

func NewResponseLoggingWriter(w http.ResponseWriter) *responseLoggingWriter {
	return &responseLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

func (rlw *responseLoggingWriter) WriteHeader(code int) {
	rlw.code = code
	rlw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(ctx context.Context, logger *slog.Logger, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.InfoContext(ctx, "starting logging middleware",
			slog.String("http_method", r.Method), slog.String("request_uri", r.RequestURI))

		rlw := NewResponseLoggingWriter(w)
		next(rlw, r)

		logger.InfoContext(ctx, "ending loggeing middleware", slog.Int("http_status_code", rlw.code))
	}
}
