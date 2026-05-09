package core

import (
	"errors"
)

/*
 * UserInterruptError
 */
type UserInterruptError struct {
	error
}

func NewUserInterruptError(message string) UserInterruptError {
	return UserInterruptError{
		error: errors.New(message),
	}
}

func (e UserInterruptError) Error() string {
	return e.error.Error()
}

func (e UserInterruptError) Unwrap() error {
	return e.error
}

func (e UserInterruptError) HttpStatus() uint {
	return 400
}
