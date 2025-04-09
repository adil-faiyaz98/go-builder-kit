package models

import (
	"fmt"
	"strings"
)

// MutualFund represents a mutual fund investment
type MutualFund struct {
	ID              string
	Name            string
	Symbol          string
	TickerSymbol    string
	Category        string
	FundManager     string
	FundFamily      string
	ExpenseRatio    float64
	NAV             float64
	AUM             float64
	YTDReturn       float64
	OneYearReturn   float64
	ThreeYearReturn float64
	FiveYearReturn  float64
	Risk            string
	Strategy        string
	Shares          float64
	PurchaseDate    string
	Holdings        []*FundHolding
}

// Validate validates the MutualFund model
func (m *MutualFund) Validate() error {
	var errors []string

	// Validate ID
	if m.ID == "" {
		errors = append(errors, "ID cannot be empty")
	}

	// Validate Name
	if m.Name == "" {
		errors = append(errors, "Name cannot be empty")
	}

	// Validate Symbol
	if m.Symbol == "" {
		errors = append(errors, "Symbol cannot be empty")
	}

	// Validate Category
	if m.Category == "" {
		errors = append(errors, "Category cannot be empty")
	}

	// Validate FundManager
	if m.FundManager == "" {
		errors = append(errors, "FundManager cannot be empty")
	}

	// Validate ExpenseRatio
	if m.ExpenseRatio < 0 {
		errors = append(errors, "ExpenseRatio cannot be negative")
	}

	// Validate NAV
	if m.NAV < 0 {
		errors = append(errors, "NAV cannot be negative")
	}

	// Validate AUM
	if m.AUM < 0 {
		errors = append(errors, "AUM cannot be negative")
	}

	// Validate Shares
	if m.Shares < 0 {
		errors = append(errors, "Shares cannot be negative")
	}

	// Validate Risk if provided
	if m.Risk != "" {
		validRisks := []string{"low", "moderate", "high", "very high"}
		isValidRisk := false
		for _, risk := range validRisks {
			if strings.ToLower(m.Risk) == risk {
				isValidRisk = true
				break
			}
		}
		if !isValidRisk {
			errors = append(errors, "Risk must be one of: low, moderate, high, very high")
		}
	}

	// Validate Holdings if provided
	for i, holding := range m.Holdings {
		if holding != nil {
			if err := holding.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("Holding[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}
