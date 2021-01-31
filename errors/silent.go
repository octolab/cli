package errors

// Silent holds a root cause error and allows to exit
// with specified code and optionally with a custom message.
type Silent struct {
	error
	Code    int
	Message string
}
