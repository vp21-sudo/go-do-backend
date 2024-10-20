package utils

import "github.com/go-playground/validator/v10"

// handlers/todo.go
func GetValidationErrorMessage(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "This field is required"
	default:
		return "Invalid value"
	}
}
