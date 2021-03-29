package errors

import "strings"

// Silent holds a root cause error and allows to graceful exit
// with a specified code and optionally with a custom message.
type Silent interface {
	error
	// Code returns an exit code.
	Code() int
	// Message optionally returns an exit message.
	Message() (string, bool)
}

// NewSilent returns a new Silent error caused by the passed one.
// If the passed error is nil then the result is the same.
func NewSilent(cause error, code int, message ...string) Silent {
	if cause == nil {
		return nil
	}
	return silent{cause, code, strings.Join(message, ": ")}
}

type silent struct {
	error
	code    int
	message string
}

func (err silent) Cause() error  { return err.error }
func (err silent) Code() int     { return err.code }
func (err silent) Unwrap() error { return err.error }

func (err silent) Message() (string, bool) {
	return err.message, err.message != ""
}
