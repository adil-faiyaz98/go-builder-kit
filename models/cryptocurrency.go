package models

import (
	"fmt"
	"strings"
)

// Cryptocurrency represents a cryptocurrency investment
type Cryptocurrency struct {
	ID            string
	Symbol        string
	Name          string
	Quantity      float64
	PurchasePrice float64
	CurrentPrice  float64
	Value         float64
	Wallet        string
	PurchaseDate  string
	Transactions  []*CryptoTransaction
}

// Validate validates the Cryptocurrency model
func (c *Cryptocurrency) Validate() error {
	var errors []string

	// Validate Symbol
	if c.Symbol == "" {
		errors = append(errors, "Symbol is required")
	}

	// Validate Name
	if c.Name == "" {
		errors = append(errors, "Name is required")
	}

	// Validate Quantity
	if c.Quantity < 0 {
		errors = append(errors, "Quantity cannot be negative")
	}

	// Validate PurchasePrice
	if c.PurchasePrice < 0 {
		errors = append(errors, "PurchasePrice cannot be negative")
	}

	// Validate CurrentPrice
	if c.CurrentPrice < 0 {
		errors = append(errors, "CurrentPrice cannot be negative")
	}

	// Validate Transactions if provided
	for i, transaction := range c.Transactions {
		if transaction != nil {
			if err := transaction.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("Transaction[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// CryptoTransaction represents a cryptocurrency transaction
type CryptoTransaction struct {
	Date       string
	Type       string // Buy, Sell, Transfer
	Quantity   float64
	Price      float64
	Fees       float64
	Total      float64
	Exchange   string
	Network    string
	TxHash     string
	WalletFrom string
	WalletTo   string
}

// Validate validates the CryptoTransaction model
func (t *CryptoTransaction) Validate() error {
	var errors []string

	// Validate Date
	if t.Date == "" {
		errors = append(errors, "Date is required")
	}

	// Validate Type
	validTypes := []string{"buy", "sell", "transfer", "stake", "unstake", "swap", "airdrop", "mining", "other"}
	isValidType := false
	for _, validType := range validTypes {
		if strings.ToLower(t.Type) == validType {
			isValidType = true
			break
		}
	}
	if !isValidType {
		errors = append(errors, "Type must be one of: buy, sell, transfer, stake, unstake, swap, airdrop, mining, other")
	}

	// Validate Quantity
	if t.Quantity <= 0 {
		errors = append(errors, "Quantity must be greater than 0")
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
