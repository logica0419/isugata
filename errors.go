package isugata

import "github.com/pkg/errors"

var (
	// ErrUndecodableBody is returned when the body cannot be decoded
	ErrUndecodableBody = errors.New("undecodable body")

	// ErrInvalidStatusCode is returned when the status code is invalid
	ErrInvalidStatusCode = errors.New("invalid status code")

	// ErrInvalidContentType is returned when the Content-Type header is invalid
	ErrInvalidContentType = errors.New("invalid Content-Type header")

	// ErrInvalidBody is returned when the body doesn't match the expected
	ErrInvalidBody = errors.New("invalid body")
)
