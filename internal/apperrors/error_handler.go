package apperrors

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/exaream/go-api/internal/middlewares"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, logger *slog.Logger, err error) {
	var appErr *AppError
	if !errors.As(err, &appErr) {
		appErr = &AppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	ctx := r.Context()

	traceID := middlewares.GetTraceID(ctx)
	logger.ErrorContext(ctx, "error occurred",
		slog.Any("trace_id", traceID),
		slog.String("code", appErr.ErrCode.String()),
		slog.String("message", appErr.Message),
		slog.String("error", appErr.Error()))

	w.WriteHeader(appErr.ErrCode.HTTPStatusCode())

	if err := json.NewEncoder(w).Encode(appErr); err != nil {
		logger.ErrorContext(ctx, "error occurred",
			slog.Any("trace_id", traceID),
			slog.String("code", FailedToEncodeJSON.String()),
			slog.String("message", "failed to write response"),
			slog.String("error", err.Error()))
	}
}
