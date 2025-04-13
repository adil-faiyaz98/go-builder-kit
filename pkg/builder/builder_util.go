package builder

import (
	"fmt"
	"reflect"
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
func ValidateRequired(value interface{}, field string, errors *ValidationErrors) {
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
func ValidateMinValue(value interface{}, field string, minValue float64, errors *ValidationErrors) {
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
func ValidateMaxValue(value interface{}, field string, maxValue float64, errors *ValidationErrors) {
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
