package errors

import (
	"errors"
	"fmt"
	"runtime/debug"

	"firebase.google.com/go/v4/auth"
	"github.com/go-playground/validator/v10"
	"github.com/yudai2929/kendai-navi/backend/db/ent"
)

type ErrorCode int

const (
	Unknown ErrorCode = iota
	InvalidArgument
	NotFound
	AlreadyExists
	Internal
	Unauthenticated
	PermissionDenied
	Unimplemented
	Unavailable
)

func (c ErrorCode) String() string {
	switch c {
	case Unknown:
		return "Unknown"
	case InvalidArgument:
		return "InvalidArgument"
	case NotFound:
		return "NotFound"
	case AlreadyExists:
		return "AlreadyExists"
	case Internal:
		return "Internal"
	case Unauthenticated:
		return "Unauthenticated"
	case PermissionDenied:
		return "PermissionDenied"
	case Unimplemented:
		return "Unimplemented"
	case Unavailable:
		return "Unavailable"
	default:
		return "Unknown"
	}
}

type customError struct {
	code   ErrorCode
	origin error
	stack  string
}

func (ce *customError) Error() string {
	return ce.code.String()
}

func New(code ErrorCode) error {
	return &customError{
		code:   code,
		origin: fmt.Errorf(code.String()),
		stack:  string(debug.Stack()),
	}
}

func Newf(code ErrorCode, format string, args ...interface{}) error {
	return &customError{
		code:   code,
		origin: fmt.Errorf(format, args...),
		stack:  string(debug.Stack()),
	}
}

func Wrap(err error) error {
	ce := &customError{}
	if errors.As(err, &ce) {
		return err
	}
	return ce.convert(err, err, ce.stack)
}

func Wrapf(err error, format string, args ...interface{}) error {
	ce := &customError{}
	if As(err, &ce) {
		return err
	}
	return ce.convert(err, fmt.Errorf(format, args...), ce.stack)
}

func (cs *customError) convert(err error, origin error, stack string) error {
	cs.origin = origin
	cs.stack = stack

	// Validator
	var ve validator.ValidationErrors
	if As(err, &ve) {
		cs.code = InvalidArgument
		return cs
	}

	// Firebase Auth
	if auth.IsEmailNotFound(err) {
		cs.code = NotFound
		return cs
	}
	if auth.IsEmailAlreadyExists(err) {
		cs.code = AlreadyExists
		return cs
	}
	if auth.IsUserNotFound(err) {
		cs.code = NotFound
		return cs
	}

	// Ent
	if ent.IsNotFound(err) {
		cs.code = NotFound
		return cs
	}
	if ent.IsConstraintError(err) {
		cs.code = AlreadyExists
		return cs
	}
	if ent.IsNotSingular(err) {
		cs.code = Internal
		return cs
	}
	if ent.IsNotLoaded(err) {
		cs.code = Internal
		return cs
	}
	if ent.IsValidationError(err) {
		cs.code = InvalidArgument
		return cs
	}

	// TODO: tx error

	cs.code = Internal
	return cs
}

func EqualCode(err error, code ErrorCode) bool {
	e, ok := err.(*customError)
	if !ok {
		return false
	}
	return e.code == code
}

func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}
