package models

import (
	"fmt"
	"strings"
	"time"
)

// FundHolding represents a holding in a mutual fund
type FundHolding struct {
	Name         string
	Type         string
	Percentage   float64
	Value        float64
	PurchaseDate string
	Performance  float64
}

// Validate validates the FundHolding model
func (f *FundHolding) Validate() error {
	var errors []string

	// Validate Name
	if f.Name == "" {
		errors = append(errors, "Name cannot be empty")
	}

	// Validate Type
	if f.Type == "" {
		errors = append(errors, "Type cannot be empty")
	}

	// Validate Percentage
	if f.Percentage < 0 || f.Percentage > 100 {
		errors = append(errors, "Percentage must be between 0 and 100")
	}

	// Validate Value
	if f.Value < 0 {
		errors = append(errors, "Value cannot be negative")
	}

	// Validate PurchaseDate if provided
	if f.PurchaseDate != "" {
		_, err := time.Parse("2006-01-02", f.PurchaseDate)
		if err != nil {
			errors = append(errors, "PurchaseDate must be in the format YYYY-MM-DD")
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}
