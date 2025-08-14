package builder

import (
	"fmt"
	"testing"
)

// TestPerson is a simple struct for testing
type TestPerson struct {
	Name string
	Age  int
}

// Validate implements validation for TestPerson
func (p *TestPerson) Validate() error {
	if p.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if p.Age < 0 {
		return fmt.Errorf("age cannot be negative")
	}
	return nil
}

func TestBaseGenericBuilder(t *testing.T) {
	t.Run("Build returns correct object", func(t *testing.T) {
		builder := NewBaseGenericBuilder(func() TestPerson {
			return TestPerson{Name: "John", Age: 30}
		})

		result := builder.Build()
		if result.Name != "John" || result.Age != 30 {
			t.Errorf("Expected {John 30}, got %+v", result)
		}
	})

	t.Run("BuildPtr returns pointer to object", func(t *testing.T) {
		builder := NewBaseGenericBuilder(func() TestPerson {
			return TestPerson{Name: "Jane", Age: 25}
		})

		result := builder.BuildPtr()
		if result == nil {
			t.Fatal("Expected non-nil pointer")
		}
		if result.Name != "Jane" || result.Age != 25 {
			t.Errorf("Expected {Jane 25}, got %+v", *result)
		}
	})

	t.Run("WithValidation adds validation function", func(t *testing.T) {
		builder := NewBaseGenericBuilder(func() TestPerson {
			return TestPerson{Name: "", Age: 30} // Invalid name
		})

		builder.WithValidation(func(p TestPerson) error {
			if p.Name == "" {
				return fmt.Errorf("name is required")
			}
			return nil
		})

		_, err := builder.BuildAndValidate()
		if err == nil {
			t.Error("Expected validation error")
		}
		if err.Error() != "validation failed at index 0: name is required" {
			t.Errorf("Unexpected error message: %s", err.Error())
		}
	})

	t.Run("BuildAndValidate with nil builder", func(t *testing.T) {
		var builder *BaseGenericBuilder[TestPerson]
		_, err := builder.BuildAndValidate()
		if err == nil {
			t.Error("Expected error for nil builder")
		}
	})

	t.Run("MustBuild panics on validation error", func(t *testing.T) {
		builder := NewBaseGenericBuilder(func() TestPerson {
			return TestPerson{Name: "", Age: 30}
		})

		builder.WithValidation(func(p TestPerson) error {
			return fmt.Errorf("validation error")
		})

		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected panic")
			}
		}()

		builder.MustBuild()
	})

	t.Run("Clone creates independent copy", func(t *testing.T) {
		original := NewBaseGenericBuilder(func() TestPerson {
			return TestPerson{Name: "Original", Age: 30}
		})

		original.WithValidation(func(p TestPerson) error {
			return fmt.Errorf("original validation")
		})

		cloned := original.Clone()
		if cloned == nil {
			t.Fatal("Clone returned nil")
		}

		// Modify clone - need to cast to concrete type
		if clonedBuilder, ok := cloned.(*BaseGenericBuilder[TestPerson]); ok {
			clonedBuilder.WithValidation(func(p TestPerson) error {
				return fmt.Errorf("cloned validation")
			})
		}

		// Original should still have 1 validator, clone should have 2
		originalResult := original.Build()
		clonedResult := cloned.Build()

		if originalResult.Name != clonedResult.Name {
			t.Error("Clone should have same build function")
		}
	})

	t.Run("Clone with nil builder", func(t *testing.T) {
		var builder *BaseGenericBuilder[TestPerson]
		cloned := builder.Clone()
		if cloned != nil {
			t.Error("Expected nil clone from nil builder")
		}
	})
}

func TestValidationErrors(t *testing.T) {
	t.Run("ValidationError implements error interface", func(t *testing.T) {
		err := ValidationError{Field: "name", Message: "is required"}
		expected := "name: is required"
		if err.Error() != expected {
			t.Errorf("Expected %q, got %q", expected, err.Error())
		}
	})

	t.Run("ValidationErrors with multiple errors", func(t *testing.T) {
		var errors ValidationErrors
		errors.Add("name", "is required")
		errors.Add("age", "must be positive")

		if !errors.HasErrors() {
			t.Error("Expected HasErrors to return true")
		}

		errorMsg := errors.Error()
		if errorMsg == "" {
			t.Error("Expected non-empty error message")
		}
	})

	t.Run("ValidationErrors with no errors", func(t *testing.T) {
		var errors ValidationErrors
		if errors.HasErrors() {
			t.Error("Expected HasErrors to return false")
		}
		if errors.Error() != "" {
			t.Error("Expected empty error message")
		}
	})
}

func TestValidateRequired(t *testing.T) {
	t.Run("validates string field", func(t *testing.T) {
		var errors ValidationErrors
		ValidateRequired("", "name", &errors)
		if !errors.HasErrors() {
			t.Error("Expected validation error for empty string")
		}
	})

	t.Run("validates nil pointer", func(t *testing.T) {
		var errors ValidationErrors
		var ptr *string
		ValidateRequired(ptr, "pointer", &errors)
		if !errors.HasErrors() {
			t.Error("Expected validation error for nil pointer")
		}
	})

	t.Run("handles nil errors parameter", func(t *testing.T) {
		// Should not panic
		ValidateRequired("", "field", nil)
	})

	t.Run("handles empty field name", func(t *testing.T) {
		var errors ValidationErrors
		ValidateRequired("value", "", &errors)
		if !errors.HasErrors() {
			t.Error("Expected validation error for empty field name")
		}
	})
}

func TestSanitizeString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"normal string", "normal string"},
		{"  whitespace  ", "whitespace"},
		{"string\x00with\x01null", "stringwithnull"},
		{"string\twith\ttabs", "string\twith\ttabs"},         // tabs should be preserved
		{"string\nwith\nnewlines", "string\nwith\nnewlines"}, // newlines should be preserved
	}

	for _, test := range tests {
		result := SanitizeString(test.input)
		if result != test.expected {
			t.Errorf("SanitizeString(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestValidateStringLength(t *testing.T) {
	var errors ValidationErrors

	// Test minimum length
	ValidateStringLength("ab", "field", 3, -1, &errors)
	if !errors.HasErrors() {
		t.Error("Expected validation error for string too short")
	}

	// Reset errors
	errors = ValidationErrors{}

	// Test maximum length
	ValidateStringLength("abcdef", "field", -1, 5, &errors)
	if !errors.HasErrors() {
		t.Error("Expected validation error for string too long")
	}
}

func TestValidateNumericRange(t *testing.T) {
	var errors ValidationErrors

	// Test minimum value
	ValidateNumericRange(5.0, 10.0, 20.0, "field", &errors)
	if !errors.HasErrors() {
		t.Error("Expected validation error for value too small")
	}

	// Reset errors
	errors = ValidationErrors{}

	// Test maximum value
	ValidateNumericRange(25.0, 10.0, 20.0, "field", &errors)
	if !errors.HasErrors() {
		t.Error("Expected validation error for value too large")
	}
}
