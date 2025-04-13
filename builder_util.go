package builders

import (
	"fmt"
	"reflect"
)

// Builder is the interface that all builders must implement
type Builder interface {
	// Build builds the object and returns it
	Build() interface{}
}

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

// DeepCopy creates a deep copy of the given value
func DeepCopy(src interface{}) interface{} {
	if src == nil {
		return nil
	}

	// Get the value that the pointer src points to
	original := reflect.ValueOf(src)
	
	// If it's a pointer, get the value it points to
	if original.Kind() == reflect.Ptr {
		original = original.Elem()
	}
	
	// Create a new object of the same type as the original
	cpy := reflect.New(original.Type()).Elem()
	
	// Copy all the fields
	copyRecursive(original, cpy)
	
	// Return the address of the copy
	return cpy.Addr().Interface()
}

// copyRecursive does the actual copying of the values
func copyRecursive(original, cpy reflect.Value) {
	// Handle according to the kind of value
	switch original.Kind() {
	case reflect.Ptr:
		// Get the actual value being pointed to
		originalValue := original.Elem()
		
		// If it's a nil pointer, return
		if !originalValue.IsValid() {
			return
		}
		
		// Create a new object and set the pointer to it
		cpy.Set(reflect.New(originalValue.Type()))
		
		// Recursively copy the pointed-to value
		copyRecursive(originalValue, cpy.Elem())
		
	case reflect.Interface:
		// If it contains a nil value, return
		if original.IsNil() {
			return
		}
		
		// Get the value for the interface
		originalValue := original.Elem()
		
		// Create a copy of the value
		copyValue := reflect.New(originalValue.Type()).Elem()
		copyRecursive(originalValue, copyValue)
		
		// Set the copy value to the interface
		cpy.Set(copyValue)
		
	case reflect.Struct:
		// Copy each field
		for i := 0; i < original.NumField(); i++ {
			copyRecursive(original.Field(i), cpy.Field(i))
		}
		
	case reflect.Slice:
		// If it's a nil slice, return
		if original.IsNil() {
			return
		}
		
		// Create a new slice and copy each element
		cpy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
		for i := 0; i < original.Len(); i++ {
			copyRecursive(original.Index(i), cpy.Index(i))
		}
		
	case reflect.Map:
		// If it's a nil map, return
		if original.IsNil() {
			return
		}
		
		// Create a new map and copy each key-value pair
		cpy.Set(reflect.MakeMap(original.Type()))
		for _, key := range original.MapKeys() {
			originalValue := original.MapIndex(key)
			copyValue := reflect.New(originalValue.Type()).Elem()
			copyRecursive(originalValue, copyValue)
			cpy.SetMapIndex(key, copyValue)
		}
		
	default:
		// For primitive types, just set the value
		cpy.Set(original)
	}
}
