// Package isugata HTTP response validator for ISUCON benchmarker
package isugata

import (
	"io"
	"net/http"
)

// ValidateOpt is function type to validate http.Response
type ValidateOpt func(*http.Response) error

// Validate validates http.Response with opts
//
//	validate http.Response in order of opts, and return the first error
func Validate(res *http.Response, opts ...ValidateOpt) error {
	if res.Body != nil {
		// drain body to enable HTTP keep-alive
		defer func() {
			_, _ = io.Copy(io.Discard, res.Body)
			_ = res.Body.Close()
		}()
	}

	for _, opt := range opts {
		if err := opt(res); err != nil {
			return err
		}
	}

	return nil
}
