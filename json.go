package isugata

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

// JSONValidateOpt is function type to validate JSON body
type JSONValidateOpt[V any] func(body V) error

// WithJSONValidation decodes JSON body and validates it with opts
func WithJSONValidation[V any](opt ...JSONValidateOpt[V]) ValidateOpt {
	return func(res *http.Response) error {
		var body V
		if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
			return fmt.Errorf("%w: %w", ErrUndecodableBody, err)
		}

		for _, o := range opt {
			if err := o(body); err != nil {
				return err
			}
		}

		return nil
	}
}

// JSONEquals validates if JSON body deeply equals to v
func JSONEquals[V comparable](v V) JSONValidateOpt[V] {
	return func(body V) error {
		if !reflect.DeepEqual(v, body) {
			return fmt.Errorf("%w: invalid body content: expected: %v, actual: %v", ErrInvalidBody, v, body)
		}

		return nil
	}
}
