package apperrors

import (
	"encoding/json"
	"errors"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	var appErr *AppError
	if !errors.As(err, &appErr) {
		appErr = &AppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	var statusCode int
	switch appErr.ErrCode {
	case NAData, BadParam:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
