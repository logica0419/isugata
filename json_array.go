package isugata

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/exp/constraints"
)

// JSONArrayValidateOpt is function type to validate JSON array body
type JSONArrayValidateOpt[V any] func(body []V) error

// WithJSONArrayValidation decodes JSON array body and validates it with opts
func WithJSONArrayValidation[V any](opts ...JSONArrayValidateOpt[V]) ValidateOpt {
	return func(res *http.Response) error {
		var body []V
		if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
			return fmt.Errorf("%w: %w", ErrUndecodableBody, err)
		}

		for _, o := range opts {
			if err := o(body); err != nil {
				return err
			}
		}

		return nil
	}
}

// JSONArrayLengthEquals validates if JSON array body length equals to length
func JSONArrayLengthEquals[V any](length int) JSONArrayValidateOpt[V] {
	return func(body []V) error {
		if len(body) != length {
			return fmt.Errorf("%w: invalid array length: expected: %v, actual: %v", ErrInvalidBody, length, len(body))
		}

		return nil
	}
}

// JSONArrayLengthRange validates if JSON array body length is in range [min, max]
func JSONArrayLengthRange[V any](min, max int) JSONArrayValidateOpt[V] {
	return func(body []V) error {
		if len(body) < min || len(body) > max {
			return fmt.Errorf("%w: invalid array length: expected: %v~%v actual: %v", ErrInvalidBody, min, max, len(body))
		}

		return nil
	}
}

// JSONArrayValidateEach validates each element of JSON array body with opts
func JSONArrayValidateEach[V any](opts ...JSONValidateOpt[V]) JSONArrayValidateOpt[V] {
	return func(body []V) error {
		for _, v := range body {
			for _, opt := range opts {
				if err := opt(v); err != nil {
					return err
				}
			}
		}

		return nil
	}
}

type order string

const (
	// Asc is ascending order
	Asc order = "asc"

	// Desc is descending order
	Desc order = "desc"

	minLength = 2
)

// JSONArrayValidateOrder validates if JSON array body is ordered by idxFunc in ord
func JSONArrayValidateOrder[V any, I constraints.Ordered](idxFunc func(v V) I, ord order) JSONArrayValidateOpt[V] {
	return func(body []V) error {
		if len(body) < minLength {
			return nil
		}

		idxList := make([]I, len(body))
		for i, v := range body {
			idxList[i] = idxFunc(v)
		}

		for idx := 0; idx < len(idxList)-1; idx++ {
			switch ord {
			case Asc:
				if idxList[idx] > idxList[idx+1] {
					return fmt.Errorf("%w: invalid array order", ErrInvalidBody)
				}

			case Desc:
				if idxList[idx] < idxList[idx+1] {
					return fmt.Errorf("%w: invalid array order", ErrInvalidBody)
				}
			}
		}

		return nil
	}
}
