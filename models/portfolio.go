package models

import (
	"fmt"
	"strings"
)

// Portfolio represents an investment portfolio
type Portfolio struct {
	ID               string
	Name             string
	Description      string
	RiskLevel        string
	TotalValue       float64
	Allocation       map[string]float64
	Stocks           []*Stock
	Bonds            []*Bond
	ETFs             []*ETF
	MutualFunds      []*MutualFund
	Cryptocurrencies []*Cryptocurrency
}

// Validate validates the Portfolio model
func (p *Portfolio) Validate() error {
	var errors []string

	// Validate ID
	if p.ID == "" {
		errors = append(errors, "ID is required")
	}

	// Validate Name
	if p.Name == "" {
		errors = append(errors, "Name is required")
	}

	// Validate RiskLevel
	validRiskLevels := []string{"low", "medium", "high"}
	isValidRiskLevel := false
	for _, riskLevel := range validRiskLevels {
		if strings.ToLower(p.RiskLevel) == riskLevel {
			isValidRiskLevel = true
			break
		}
	}
	if !isValidRiskLevel {
		errors = append(errors, "RiskLevel must be one of: low, medium, high")
	}

	// Validate TotalValue
	if p.TotalValue < 0 {
		errors = append(errors, "TotalValue cannot be negative")
	}

	// Validate Allocation
	if p.Allocation != nil {
		total := 0.0
		for _, allocation := range p.Allocation {
			if allocation < 0 {
				errors = append(errors, "Allocation values cannot be negative")
				break
			}
			total += allocation
		}
		if total > 100.01 || total < 99.99 {
			errors = append(errors, "Allocation values must sum to 100%")
		}
	}

	// Validate Stocks if provided
	for i, stock := range p.Stocks {
		if stock != nil {
			if err := stock.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("Stock[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Validate Bonds if provided
	for i, bond := range p.Bonds {
		if bond != nil {
			if err := bond.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("Bond[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Validate ETFs if provided
	for i, etf := range p.ETFs {
		if etf != nil {
			if err := etf.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("ETF[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Validate MutualFunds if provided
	for i, mutualFund := range p.MutualFunds {
		if mutualFund != nil {
			if err := mutualFund.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("MutualFund[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Validate Cryptocurrencies if provided
	for i, crypto := range p.Cryptocurrencies {
		if crypto != nil {
			if err := crypto.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("Cryptocurrency[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}
