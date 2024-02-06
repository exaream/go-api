package apperrors

type ErrCode string

const (
	Unknown ErrCode = "U000"

	InsertDataFailed ErrCode = "S001"
	GetDataFailed    ErrCode = "S002"
	NAData           ErrCode = "S003" // TODO: when selecting
	NoTargetData     ErrCode = "S004" // TODO: when updating or deleting
	UpdateDataFailed ErrCode = "S005"

	ReqBodyDecodeFailed ErrCode = "R001"
	BadParam            ErrCode = "R002"
)

func (code ErrCode) Wrap(err error, msg string) *AppError {
	return &AppError{
		ErrCode: code,
		Message: msg,
		Err:     err,
	}
}
