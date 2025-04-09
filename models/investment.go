package models

import (
	"fmt"
	"strings"
	"time"
)

// PerformanceRecord represents a performance record for an investment
type PerformanceRecord struct {
	Date  string
	Value float64
	Notes string
	ROI   float64
}

// Validate validates the PerformanceRecord model
func (p *PerformanceRecord) Validate() error {
	var errors []string

	// Validate Date
	if p.Date == "" {
		errors = append(errors, "Date is required")
	} else {
		// Parse date
		date, err := time.Parse("2006-01-02", p.Date)
		if err != nil {
			errors = append(errors, "Date must be in the format YYYY-MM-DD")
		} else {
			// Check if date is in the future
			if date.After(time.Now()) {
				errors = append(errors, "Date cannot be in the future")
			}
		}
	}

	// Validate Value
	if p.Value < 0 {
		errors = append(errors, "Value cannot be negative")
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// Investment represents an investment
type Investment struct {
	ID          string
	Name        string
	Type        string
	Value       float64
	StartDate   string
	EndDate     string
	Risk        string
	Portfolio   *Portfolio
	Performance []*PerformanceRecord
}

// Validate validates the Investment model
func (i *Investment) Validate() error {
	var errors []string

	// Validate ID
	if i.ID == "" {
		errors = append(errors, "ID is required")
	}

	// Validate Name
	if i.Name == "" {
		errors = append(errors, "Name is required")
	}

	// Validate Type
	validTypes := []string{"stock", "bond", "mutual fund", "etf", "real estate", "cryptocurrency", "other"}
	isValidType := false
	for _, t := range validTypes {
		if strings.ToLower(i.Type) == t {
			isValidType = true
			break
		}
	}
	if !isValidType {
		errors = append(errors, "Type must be one of: stock, bond, mutual fund, etf, real estate, cryptocurrency, other")
	}

	// Validate Value
	if i.Value < 0 {
		errors = append(errors, "Value cannot be negative")
	}

	// Validate StartDate
	if i.StartDate != "" {
		// Parse date
		date, err := time.Parse("2006-01-02", i.StartDate)
		if err != nil {
			errors = append(errors, "StartDate must be in the format YYYY-MM-DD")
		} else {
			// Check if date is in the future
			if date.After(time.Now()) {
				errors = append(errors, "StartDate cannot be in the future")
			}

			// Validate EndDate if provided
			if i.EndDate != "" {
				endDate, err := time.Parse("2006-01-02", i.EndDate)
				if err != nil {
					errors = append(errors, "EndDate must be in the format YYYY-MM-DD")
				} else if endDate.Before(date) {
					errors = append(errors, "EndDate cannot be before StartDate")
				}
			}
		}
	}

	// Validate Risk
	validRisks := []string{"low", "medium", "high"}
	isValidRisk := false
	for _, r := range validRisks {
		if strings.ToLower(i.Risk) == r {
			isValidRisk = true
			break
		}
	}
	if !isValidRisk {
		errors = append(errors, "Risk must be one of: low, medium, high")
	}

	// Validate Portfolio if provided
	if i.Portfolio != nil {
		if err := i.Portfolio.Validate(); err != nil {
			errors = append(errors, fmt.Sprintf("Portfolio validation failed: %s", err.Error()))
		}
	}

	// Validate Performance records if provided
	for j, record := range i.Performance {
		if record != nil {
			if err := record.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("PerformanceRecord[%d] validation failed: %s", j, err.Error()))
			}
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}
