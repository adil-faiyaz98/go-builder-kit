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
	Manager     any   // Will be *Person, using any to avoid import cycle
	Employees   []any // Will be []*Person, using any to avoid import cycle
	Budget      float64
	HeadCount   int
	Projects    []*Project
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

	// Validate Projects if provided
	for i, project := range d.Projects {
		if project != nil {
			if err := project.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("Project[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

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
