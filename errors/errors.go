package errors

import (
	"errors"
	"fmt"
)

type Error struct {
	Status
	cause error
}

func New(code int, reason, message string) *Error {
	return &Error{
		Status: Status{
			Code:    int32(code),
			Reason:  reason,
			Message: message,
		},
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("error: code = %d reason = %s message = %s metadata = %v cause = %v", e.Code, e.Reason, e.Message, e.Metadata, e.cause)
}

func (e *Error) Is(err error) bool {
	if se := new(Error); errors.As(err, &se) {
		return se.Code == e.Code && se.Reason == e.Reason
	}
	return false
}

func (e *Error) WithCause(err error) {
	e.cause = err
}

func (e *Error) WithMetadata(md map[string]string) {
	e.Metadata = md
}
