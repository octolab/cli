package errors

import "strings"

// Silent holds a root cause error and allows to exit
// with specified code and optionally with a custom message.
type Silent interface {
	error
	Code() int
	Message() (string, bool)
}

func NewSilent(cause error, code int, message ...string) Silent {
	return silent{cause, code, strings.Join(message, ": ")}
}

type silent struct {
	error
	code    int
	message string
}

func (err silent) Cause() error {
	return err.error
}

func (err silent) Code() int {
	return err.code
}

func (err silent) Message() (string, bool) {
	return err.message, err.message != ""
}

func (err silent) Unwrap() error {
	return err.error
}
