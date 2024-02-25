package apperrors

type AppError struct {
	ErrCode
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func (e *AppError) Unwrap() error {
	return e.Err
}
