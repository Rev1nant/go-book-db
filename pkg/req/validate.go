package req

import "github.com/go-playground/validator/v10"

func IsValid[T any](payload T) error {
	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		return err
	}

	return nil
}
