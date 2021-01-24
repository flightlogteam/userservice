package common

type UserErrorType int

type UserError struct {
	ErrorType UserErrorType
	Message   string
}

func NewUserError(message string, errorType UserErrorType) error {
	return &UserError{
		Message:   message,
		ErrorType: errorType,
	}
}

func (u *UserError) Error() string {
	return u.Message
}

const (
	VALIDATION_ERROR_TYPE     UserErrorType = 1
	SQL_ERROR_TYPE            UserErrorType = 2
	USER_EXISTS_TYPE          UserErrorType = 3
	USER_DOES_NOT_EXISTS_TYPE UserErrorType = 4
	BAD_CREDENTIALS_ERROR     UserErrorType = 5
)
