package model

import (
	"errors"
	"reflect"
)

func BA[T any](list []T, err error) (T, error) {
	var zero T

	if err != nil {
		return zero, err
	}

	if len(list) > 1 {
		return zero, CreateTooManyError(zero, "Expect One Result")
	}

	if len(list) == 0 {
		return zero, CreateNotExistError(zero, "Expect One Result")
	}

	return list[0], nil
}

func BNA[T any](list []T, err error) (*T, error) {
	var zero *T

	if err != nil {
		return zero, err
	}

	if len(list) > 1 {
		return zero, CreateTooManyError(zero, "Expect One Result")
	}

	if len(list) == 0 {
		return nil, nil
	}

	return &list[0], nil
}

/*
 * NotExistError
 */
type NotExistError struct {
	StructName string
	error
}

func CreateNotExistError[T any](obj T, message string) *NotExistError {
	return NewNotExistError(reflect.TypeOf(obj).Name(), message)
}

func NewNotExistError(name string, message string) *NotExistError {
	return &NotExistError{
		StructName: name,
		error: errors.New(message),
	}
}

func (e NotExistError) Error() string {
	return e.error.Error() + ", struct_name: " + e.StructName
}

func (e NotExistError) Unwrap() error {
	return e.error
}

func (e NotExistError) HttpStatus() uint {
	return 400
}

/*
 * TooManyError
 */
type TooManyError struct {
	StructName string
	error
}

func CreateTooManyError[T any](obj T, message string) *TooManyError {
	return NewTooManyError(reflect.TypeOf(obj).Name(), message)
}

func NewTooManyError(name string, message string) *TooManyError {
	return &TooManyError{
		StructName: name,
		error: errors.New(message),
	}
}

func (e TooManyError) Error() string {
	return e.error.Error() + ", struct_name: " + e.StructName
}

func (e TooManyError) Unwrap() error {
	return e.error
}

func (e TooManyError) HttpStatus() uint {
	return 400
}
