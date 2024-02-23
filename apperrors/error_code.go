package apperrors

import "net/http"

type ErrCode string

const (
	Unknown ErrCode = "U000"

	FailedToInsert   ErrCode = "S001"
	FailedToSelect   ErrCode = "S002"
	FailedToUpdate   ErrCode = "S003"
	FailedToDelete   ErrCode = "S004"
	NotFound         ErrCode = "S005"
	NotFoundToUpdate ErrCode = "S006"
	AlreadyUpdated   ErrCode = "S007"
	AlreadyDeleted   ErrCode = "S008"

	FailedToDecodeReq ErrCode = "R001"
	BadParam          ErrCode = "R002"
)

func (code ErrCode) Wrap(err error, msg string) *AppError {
	return &AppError{
		ErrCode: code,
		Message: msg,
		Err:     err,
	}
}

func (code ErrCode) String() string {
	return string(code)
}

func (code ErrCode) HTTPStatusCode() int {
	switch code {
	case BadParam:
		return http.StatusBadRequest
	case NotFound:
		return http.StatusNotFound
	case NotFoundToUpdate, FailedToDecodeReq:
		return http.StatusBadRequest
	case AlreadyUpdated, AlreadyDeleted:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
