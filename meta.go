package isugata

import (
	"fmt"
	"net/http"
	"strings"
)

// WithStatusCode validates if the status code equals to code
func WithStatusCode(code int) ValidateOpt {
	return func(res *http.Response) error {
		if res.StatusCode != code {
			return fmt.Errorf("%w: expected: %d, actual: %d", ErrInvalidStatusCode, code, res.StatusCode)
		}

		return nil
	}
}

// WithContentType validates if the Content-Type header starts with contentType
func WithContentType(contentType string) ValidateOpt {
	return func(res *http.Response) error {
		actual := res.Header.Get("Content-Type")
		if !strings.HasPrefix(actual, contentType) {
			return fmt.Errorf("%w: expected: %s, actual: %s", ErrInvalidContentType, contentType, actual)
		}

		return nil
	}
}
