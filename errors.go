package xfers

import (
	"errors"
	"fmt"
)

var (
	// ErrInternal is general internal error.
	ErrInternal = errors.New("internal error")
	// ErrSandboxOnly is error when calling sandbox feature only in prod env.
	ErrSandboxOnly = errors.New("sandbox only")
)

func errRequiredField(str string) error {
	return fmt.Errorf("required field %s", str)
}

func errInvalidValueField(str string) error {
	return fmt.Errorf("invalid %s value", str)
}

func errNumericField(str string) error {
	return fmt.Errorf("field %s must contain number only", str)
}

func errURLField(str string) error {
	return fmt.Errorf("field %s must be in URL format", str)
}

func errGTField(str, value string) error {
	return fmt.Errorf("field %s must be greater than %s", str, value)
}

func errGTEField(str, value string) error {
	return fmt.Errorf("field %s must be greater than or equal %s", str, value)
}

func errMaxField(str, value string) error {
	return fmt.Errorf("field %s max value is %s", str, value)
}
