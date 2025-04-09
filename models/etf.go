package models

import (
	"fmt"
	"strings"
)

// ETF represents an Exchange-Traded Fund
type ETF struct {
	ID            string
	Symbol        string
	TickerSymbol  string
	Name          string
	Provider      string
	ExpenseRatio  float64
	Category      string
	Index         string
	Holdings      map[string]float64
	Price         float64
	PurchasePrice float64
	CurrentPrice  float64
	PurchaseDate  string
	Quantity      int
	Shares        float64
	TotalValue    float64
	YieldPercent  float64
}

// Validate validates the ETF model
func (e *ETF) Validate() error {
	var errors []string

	// Validate Symbol
	if e.Symbol == "" {
		errors = append(errors, "Symbol is required")
	}

	// Validate Name
	if e.Name == "" {
		errors = append(errors, "Name is required")
	}

	// Validate ExpenseRatio
	if e.ExpenseRatio < 0 {
		errors = append(errors, "ExpenseRatio cannot be negative")
	}

	// Validate Category
	validCategories := []string{"equity", "bond", "commodity", "real estate", "mixed", "sector", "international", "other"}
	isValidCategory := false
	for _, category := range validCategories {
		if strings.ToLower(e.Category) == category {
			isValidCategory = true
			break
		}
	}
	if !isValidCategory {
		errors = append(errors, "Category must be one of: equity, bond, commodity, real estate, mixed, sector, international, other")
	}

	// Validate Holdings
	if e.Holdings != nil {
		total := 0.0
		for _, holding := range e.Holdings {
			if holding < 0 {
				errors = append(errors, "Holdings values cannot be negative")
				break
			}
			total += holding
		}
		if total > 100.01 || total < 99.99 {
			errors = append(errors, "Holdings values must sum to 100%")
		}
	}

	// Validate Price
	if e.Price < 0 {
		errors = append(errors, "Price cannot be negative")
	}

	// Validate Quantity
	if e.Quantity < 0 {
		errors = append(errors, "Quantity cannot be negative")
	}

	// Validate TotalValue
	if e.TotalValue < 0 {
		errors = append(errors, "TotalValue cannot be negative")
	}

	// Validate YieldPercent
	if e.YieldPercent < 0 {
		errors = append(errors, "YieldPercent cannot be negative")
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}
