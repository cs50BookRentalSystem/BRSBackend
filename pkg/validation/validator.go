package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(s interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func FormatErrors(errors []*ErrorResponse) string {
	var errorMsgs []string
	for _, err := range errors {
		errorMsgs = append(errorMsgs, fmt.Sprintf("field %s failed on the '%s' tag", err.FailedField, err.Tag))
	}
	return strings.Join(errorMsgs, ", ")
}
