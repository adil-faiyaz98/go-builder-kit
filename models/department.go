package models

import (
	"fmt"
	"strings"
)

// Department represents a department in an organization
type Department struct {
	Name        string
	Code        string
	Description string
	Manager     interface{}   // Will be *Person, using interface{} to avoid import cycle
	Employees   []interface{} // Will be []*Person, using interface{} to avoid import cycle
	Budget      float64
	HeadCount   int
	Projects    []interface{} // Simplified to avoid undefined type
	Location    *Address
}

// Validate validates the Department model
func (d *Department) Validate() error {
	var errors []string

	// Validate Name
	if d.Name == "" {
		errors = append(errors, "Name cannot be empty")
	}

	// Validate Code
	if d.Code == "" {
		errors = append(errors, "Code cannot be empty")
	}

	// Validate Budget
	if d.Budget < 0 {
		errors = append(errors, "Budget cannot be negative")
	}

	// Skip validation for Projects as they are now any type

	// Validate Location if provided
	if d.Location != nil {
		if err := d.Location.Validate(); err != nil {
			errors = append(errors, fmt.Sprintf("Location validation failed: %s", err.Error()))
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}
