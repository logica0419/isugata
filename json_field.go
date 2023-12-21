package isugata

import (
	"fmt"
	"reflect"
)

// JSONFieldValidate validates JSON filed with opts
func JSONFieldValidate[V, F any](name string, opts ...JSONValidateOpt[F]) JSONValidateOpt[V] {
	return func(body V) error {
		bodyV := reflect.ValueOf(body)
		fieldI := bodyV.FieldByName(name).Interface()

		field, ok := fieldI.(F)
		if !ok {
			return fmt.Errorf("%w: unable to get field %s", ErrInvalidBody, name)
		}

		for _, o := range opts {
			if err := o(field); err != nil {
				return err
			}
		}

		return nil
	}
}

// JSONArrayFieldValidate validates JSON array filed with opts
func JSONArrayFieldValidate[V, F any](name string, opts ...JSONArrayValidateOpt[F]) JSONValidateOpt[V] {
	return func(body V) error {
		bodyRv := reflect.ValueOf(body)
		fieldI := bodyRv.FieldByName(name).Interface()

		field, ok := fieldI.([]F)
		if !ok {
			return fmt.Errorf("%w: unable to get field %s", ErrInvalidBody, name)
		}

		for _, o := range opts {
			if err := o(field); err != nil {
				return err
			}
		}

		return nil
	}
}
