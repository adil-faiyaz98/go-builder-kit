package models

import (
	"fmt"
	"strings"
	"time"
)

// Stock represents a stock investment
type Stock struct {
	Ticker        string
	CompanyName   string
	Units         int
	PurchasePrice float64
	CurrentPrice  float64
	PurchaseDate  string
	Sector        string
	DividendYield float64
	Trades        []*Trade
}

// Validate validates the Stock model
func (s *Stock) Validate() error {
	var errors []string

	// Validate Ticker
	if s.Ticker == "" {
		errors = append(errors, "Ticker is required")
	}

	// Validate CompanyName
	if s.CompanyName == "" {
		errors = append(errors, "CompanyName is required")
	}

	// Validate Units
	if s.Units < 0 {
		errors = append(errors, "Units cannot be negative")
	}

	// Validate PurchasePrice
	if s.PurchasePrice < 0 {
		errors = append(errors, "PurchasePrice cannot be negative")
	}

	// Validate CurrentPrice
	if s.CurrentPrice < 0 {
		errors = append(errors, "CurrentPrice cannot be negative")
	}

	// Validate Trades if provided
	for i, trade := range s.Trades {
		if trade != nil {
			if err := trade.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("Trade[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// Trade represents a stock trade
type Trade struct {
	Date       string
	Type       string // Buy, Sell
	Units      int
	Price      float64
	Fees       float64
	Total      float64
	Exchange   string
	ExecutedBy string
}

// Validate validates the Trade model
func (t *Trade) Validate() error {
	var errors []string

	// Validate Date
	if t.Date != "" {
		_, err := time.Parse("2006-01-02", t.Date)
		if err != nil {
			errors = append(errors, "Date must be in the format YYYY-MM-DD")
		}
	}

	// Validate Type
	validTypes := []string{"buy", "sell", "dividend", "split"}
	isValidType := false
	for _, validType := range validTypes {
		if strings.ToLower(t.Type) == validType {
			isValidType = true
			break
		}
	}
	if !isValidType {
		errors = append(errors, "Type must be one of: buy, sell, dividend, split")
	}

	// Validate Units
	if t.Units <= 0 {
		errors = append(errors, "Units must be greater than 0")
	}

	// Validate Price
	if t.Price < 0 {
		errors = append(errors, "Price cannot be negative")
	}

	// Validate Fees
	if t.Fees < 0 {
		errors = append(errors, "Fees cannot be negative")
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}
