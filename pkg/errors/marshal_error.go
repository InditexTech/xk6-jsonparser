package errors

import "github.com/pkg/errors"

type MarshalError struct {
	error
}

func WrapMarshalError(err error, format string, args ...interface{}) error {
	return &MarshalError{errors.Wrapf(err, format, args...)}
}
