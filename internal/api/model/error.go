package model

type PocErrorType string

const (
	ErrInvalidInput PocErrorType = "invalid_input"
	ErrNotFound     PocErrorType = "not_found"
)

type PocError struct {
	err     error
	errType PocErrorType
}

func (e PocError) Error() string {
	return e.err.Error()
}

func IsError(err error, errType PocErrorType) bool {
	if pocErr, ok := err.(PocError); ok {
		return pocErr.errType == errType
	}

	return false
}

func NewPocError(err error, errType PocErrorType) PocError {
	return PocError{
		err:     err,
		errType: errType,
	}
}

func NewNotFoundError(err error) PocError {
	return NewPocError(err, ErrNotFound)
}

func IsNotFoundError(err error) bool {
	return IsError(err, ErrNotFound)
}

func NewInvalidInputError(err error) PocError {
	return NewPocError(err, ErrInvalidInput)
}

func IsInvalidInputError(err error) bool {
	return IsError(err, ErrInvalidInput)
}
