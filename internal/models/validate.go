package models

import (
	"fmt"

	"github.com/go-playground/validator"
)

var validate = validator.New()

func ValidateStruct(s interface{}) (bool, []error) {
	err := validate.Struct(s)
	if err != nil {
		var errs []error
		for _, validationError := range err.(validator.ValidationErrors) {
			errs = append(errs, fmt.Errorf("invalid %s value for '%s'",
				validationError.Tag(),
				validationError.StructNamespace(),
			))
		}
		if len(errs) > 0 {
			return false, errs
		}
	}
	return true, nil
}