package errors

import "github.com/pkg/errors"

type UnmarshalError struct {
	error
}

func WrapUnmarshalError(err error, format string, args ...interface{}) error {
	return &UnmarshalError{errors.Wrapf(err, format, args...)}
}
