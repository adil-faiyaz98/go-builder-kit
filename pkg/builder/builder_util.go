package builder

import (
	"fmt"
	"reflect"
	"strings"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

// Error returns the error message
func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidationErrors represents multiple validation errors
type ValidationErrors []ValidationError

// Error returns the error message
func (e ValidationErrors) Error() string {
	if len(e) == 0 {
		return ""
	}

	msg := "validation failed:"
	for _, err := range e {
		msg += fmt.Sprintf(" %s;", err.Error())
	}
	return msg
}

// Add adds a validation error
func (e *ValidationErrors) Add(field, message string) {
	*e = append(*e, ValidationError{Field: field, Message: message})
}

// HasErrors returns true if there are any validation errors
func (e ValidationErrors) HasErrors() bool {
	return len(e) > 0
}

// ValidateRequired validates that a field is not empty
func ValidateRequired(value any, field string, errors *ValidationErrors) {
	if errors == nil {
		return // Defensive programming - don't panic if errors is nil
	}

	if field == "" {
		errors.Add("unknown", "field name cannot be empty")
		return
	}

	if value == nil {
		errors.Add(field, "cannot be empty")
		return
	}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String:
		if v.String() == "" {
			errors.Add(field, "cannot be empty")
		}
	case reflect.Slice, reflect.Map:
		if v.Len() == 0 {
			errors.Add(field, "cannot be empty")
		}
	case reflect.Ptr:
		if v.IsNil() {
			errors.Add(field, "cannot be empty")
		}
	}
}

// ValidateStringLength validates string length constraints
func ValidateStringLength(value string, field string, minLen, maxLen int, errors *ValidationErrors) {
	if errors == nil {
		return
	}

	if field == "" {
		errors.Add("unknown", "field name cannot be empty")
		return
	}

	length := len(value)
	if minLen >= 0 && length < minLen {
		errors.Add(field, fmt.Sprintf("must be at least %d characters long", minLen))
	}
	if maxLen >= 0 && length > maxLen {
		errors.Add(field, fmt.Sprintf("must be at most %d characters long", maxLen))
	}
}

// ValidateNumericRange validates numeric values are within range
func ValidateNumericRange(value, min, max float64, field string, errors *ValidationErrors) {
	if errors == nil {
		return
	}

	if field == "" {
		errors.Add("unknown", "field name cannot be empty")
		return
	}

	if value < min {
		errors.Add(field, fmt.Sprintf("must be at least %g", min))
	}
	if value > max {
		errors.Add(field, fmt.Sprintf("must be at most %g", max))
	}
}

// SanitizeString removes potentially dangerous characters from strings
func SanitizeString(input string) string {
	// Remove null bytes and control characters except tab, newline, and carriage return
	sanitized := strings.Map(func(r rune) rune {
		if r == 0 || (r < 32 && r != 9 && r != 10 && r != 13) {
			return -1 // Remove character
		}
		return r
	}, input)

	// Trim whitespace
	return strings.TrimSpace(sanitized)
}

// ValidateMinLength validates that a string has a minimum length
func ValidateMinLength(value string, field string, minLength int, errors *ValidationErrors) {
	if len(value) < minLength {
		errors.Add(field, fmt.Sprintf("must be at least %d characters", minLength))
	}
}

// ValidateMaxLength validates that a string has a maximum length
func ValidateMaxLength(value string, field string, maxLength int, errors *ValidationErrors) {
	if len(value) > maxLength {
		errors.Add(field, fmt.Sprintf("must be at most %d characters", maxLength))
	}
}

// ValidateMinValue validates that a number is at least a minimum value
func ValidateMinValue(value any, field string, minValue float64, errors *ValidationErrors) {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if float64(v.Int()) < minValue {
			errors.Add(field, fmt.Sprintf("must be at least %v", minValue))
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if float64(v.Uint()) < minValue {
			errors.Add(field, fmt.Sprintf("must be at least %v", minValue))
		}
	case reflect.Float32, reflect.Float64:
		if v.Float() < minValue {
			errors.Add(field, fmt.Sprintf("must be at least %v", minValue))
		}
	}
}

// ValidateMaxValue validates that a number is at most a maximum value
func ValidateMaxValue(value any, field string, maxValue float64, errors *ValidationErrors) {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if float64(v.Int()) > maxValue {
			errors.Add(field, fmt.Sprintf("must be at most %v", maxValue))
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if float64(v.Uint()) > maxValue {
			errors.Add(field, fmt.Sprintf("must be at most %v", maxValue))
		}
	case reflect.Float32, reflect.Float64:
		if v.Float() > maxValue {
			errors.Add(field, fmt.Sprintf("must be at most %v", maxValue))
		}
	}
}

// ValidateEmail validates that a string is a valid email address
func ValidateEmail(value string, field string, errors *ValidationErrors) {
	// Simple email validation
	if value == "" {
		return
	}

	if !strings.Contains(value, "@") || !strings.Contains(value, ".") {
		errors.Add(field, "must be a valid email address")
	}
}

// ValidateURL validates that a string is a valid URL
func ValidateURL(value string, field string, errors *ValidationErrors) {
	// Simple URL validation
	if value == "" {
		return
	}

	if !strings.HasPrefix(value, "http://") && !strings.HasPrefix(value, "https://") {
		errors.Add(field, "must be a valid URL")
	}
}
