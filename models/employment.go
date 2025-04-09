package models

import (
	"fmt"
	"strings"
	"time"
)

// Employment represents a person's employment information
type Employment struct {
	Company      *Company
	Position     string
	Department   string
	StartDate    string
	EndDate      string
	Salary       float64
	IsCurrent    bool
	Supervisor   any   // Will be *Person, using any to avoid import cycle
	Subordinates []any // Will be []*Person, using any to avoid import cycle
}

// Validate validates the Employment model
func (e *Employment) Validate() error {
	var errors []string

	// Validate Position
	if e.Position == "" {
		errors = append(errors, "Position cannot be empty")
	}

	// Validate Department
	if e.Department == "" {
		errors = append(errors, "Department cannot be empty")
	}

	// Validate StartDate
	if e.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", e.StartDate)
		if err != nil {
			errors = append(errors, "StartDate must be in the format YYYY-MM-DD")
		} else {
			// Check if StartDate is in the future
			if startDate.After(time.Now()) {
				errors = append(errors, "StartDate cannot be in the future")
			}

			// Check if StartDate is too far in the past
			if time.Now().Year()-startDate.Year() > 100 {
				errors = append(errors, "StartDate cannot be more than 100 years ago")
			}

			// Validate EndDate if provided
			if e.EndDate != "" {
				endDate, err := time.Parse("2006-01-02", e.EndDate)
				if err != nil {
					errors = append(errors, "EndDate must be in the format YYYY-MM-DD")
				} else if endDate.Before(startDate) {
					errors = append(errors, "EndDate cannot be before StartDate")
				}
			}
		}
	}

	// Validate Salary
	if e.Salary < 0 {
		errors = append(errors, "Salary cannot be negative")
	}

	// Validate Company if provided
	if e.Company != nil {
		if err := e.Company.Validate(); err != nil {
			errors = append(errors, fmt.Sprintf("Company validation failed: %s", err.Error()))
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// Company represents a company
type Company struct {
	ID          string
	Name        string
	Industry    string
	Description string
	Founded     string
	Website     string
	Address     *Address
	Location    *Address
	Size        string
	Revenue     string
	Public      bool
	StockSymbol string
	Departments []*Department
}

// Validate validates the Company model
func (c *Company) Validate() error {
	var errors []string

	// Validate Name
	if c.Name == "" {
		errors = append(errors, "Name cannot be empty")
	}

	// Validate Industry
	if c.Industry == "" {
		errors = append(errors, "Industry cannot be empty")
	}

	// Validate Founded date if provided
	if c.Founded != "" {
		foundedDate, err := time.Parse("2006-01-02", c.Founded)
		if err != nil {
			errors = append(errors, "Founded must be in the format YYYY-MM-DD")
		} else {
			// Check if founded date is in the future
			if foundedDate.After(time.Now()) {
				errors = append(errors, "Founded cannot be in the future")
			}
		}
	}

	// Validate Website if provided
	if c.Website != "" {
		// Simple validation for website format
		if !strings.HasPrefix(c.Website, "http://") && !strings.HasPrefix(c.Website, "https://") {
			errors = append(errors, "Website must start with http:// or https://")
		}
	}

	// Validate Size if provided
	if c.Size != "" {
		validSizes := []string{"startup", "small", "medium", "large", "enterprise"}
		isValidSize := false
		for _, size := range validSizes {
			if strings.ToLower(c.Size) == size {
				isValidSize = true
				break
			}
		}
		if !isValidSize {
			errors = append(errors, "Size must be one of: startup, small, medium, large, enterprise")
		}
	}

	// Validate StockSymbol if public
	if c.Public && c.StockSymbol == "" {
		errors = append(errors, "StockSymbol is required for public companies")
	}

	// Validate Address if provided
	if c.Address != nil {
		if err := c.Address.Validate(); err != nil {
			errors = append(errors, fmt.Sprintf("Address validation failed: %s", err.Error()))
		}
	}

	// Validate Location if provided
	if c.Location != nil {
		if err := c.Location.Validate(); err != nil {
			errors = append(errors, fmt.Sprintf("Location validation failed: %s", err.Error()))
		}
	}

	// Validate Departments if provided
	for i, department := range c.Departments {
		if department != nil {
			if err := department.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("Department[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}
