package middlewares

import (
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

func Logging(logger *slog.Logger, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		traceID := NewTraceID()
		ctx := SetTraceID(r.Context(), traceID)
		r = r.WithContext(ctx)

		logger.InfoContext(ctx, "starting process",
			slog.Any("trace_id", traceID), slog.String("http_method", r.Method), slog.String("request_uri", r.RequestURI))

		rlw := NewResponseLoggingWriter(w)
		next(rlw, r)

		logger.InfoContext(ctx, "ending process",
			slog.Any("trace_id", traceID), slog.Int("http_status_code", rlw.code))
	}
}
