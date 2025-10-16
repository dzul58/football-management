package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// FormatValidationError formats validation errors to a readable string
func FormatValidationError(err error) string {
	var errors []string

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errors = append(errors, formatFieldError(e))
		}
	}

	if len(errors) > 0 {
		return strings.Join(errors, "; ")
	}

	return err.Error()
}

// formatFieldError formats a single field error
func formatFieldError(err validator.FieldError) string {
	field := strings.ToLower(err.Field())

	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s wajib diisi", field)
	case "min":
		return fmt.Sprintf("%s minimal %s", field, err.Param())
	case "max":
		return fmt.Sprintf("%s maksimal %s", field, err.Param())
	case "email":
		return fmt.Sprintf("%s harus berformat email yang valid", field)
	case "url":
		return fmt.Sprintf("%s harus berformat URL yang valid", field)
	default:
		return fmt.Sprintf("%s tidak valid", field)
	}
}

// ValidateStruct validates a struct using validator
func ValidateStruct(s interface{}) error {
	validate := validator.New()
	return validate.Struct(s)
}
