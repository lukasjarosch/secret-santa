package models

import (
	"fmt"

	"github.com/go-playground/validator"
)

var validate = validator.New()

// Validate uses the global (cached) validator to analyze a struct's struct-tags and validate them accordingly.
// If errors are found, they are reformatted and returned as slice.
func Validate(s interface{}) (bool, []error) {
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