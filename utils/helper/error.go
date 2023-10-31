package helper

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

func ValidationError( err error) error {
	validationError, ok := err.(validator.ValidationErrors)
	if ok {
		messages := make([]string, 0)
		for _, e := range validationError {
			messages = append(messages, fmt.Sprintf("Validation error on field %s, tag %s", e.Field(), e.Tag()))
		}

		return fmt.Errorf("Validation failed: %s", strings.Join(messages, "; "))
	}

	return nil
}