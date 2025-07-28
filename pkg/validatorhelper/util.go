package validatorhelper

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func FormatValidationErrors(err error) map[string]string {
	errors := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			field := strings.ToLower(e.Field())
			switch e.Tag() {
			case "required":
				errors[field] = fmt.Sprintf("the field %s is required", field)
			case "email":
				errors[field] = fmt.Sprintf("the field %s is not a email format", field)
			case "max":
				errors[field] = fmt.Sprintf("the length of %s should not be greater than %s", field, e.Param())
			case "min":
				errors[field] = fmt.Sprintf("the length of %s should not be lesser than %s", field, e.Param())
			default:
				errors[field] = fmt.Sprintf("the field %s is invalid", field)
			}
		}
	}
	return errors
}
