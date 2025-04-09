package validation

import (
	"fmt"
	"strings"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidationErrors represents multiple validation errors
type ValidationErrors []ValidationError

func (e ValidationErrors) Error() string {
	if len(e) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.WriteString("Validation failed:\n")
	for _, err := range e {
		sb.WriteString(fmt.Sprintf("- %s\n", err.Error()))
	}
	return sb.String()
}

// Add adds a validation error
func (e *ValidationErrors) Add(field, message string) {
	*e = append(*e, ValidationError{Field: field, Message: message})
}

// HasErrors returns true if there are validation errors
func (e ValidationErrors) HasErrors() bool {
	return len(e) > 0
}

// Validator is the interface that all validatable models should implement
type Validator interface {
	Validate() error
}

// ValidateAll validates a slice of validators
func ValidateAll(validators ...Validator) error {
	var allErrors ValidationErrors

	for _, v := range validators {
		if v == nil {
			continue
		}

		err := v.Validate()
		if err == nil {
			continue
		}

		switch e := err.(type) {
		case ValidationError:
			allErrors = append(allErrors, e)
		case ValidationErrors:
			allErrors = append(allErrors, e...)
		default:
			allErrors.Add("unknown", err.Error())
		}
	}

	if !allErrors.HasErrors() {
		return nil
	}

	return allErrors
}
