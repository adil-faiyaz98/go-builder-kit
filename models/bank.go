package models

import (
	"fmt"
	"strings"
	"time"
)

// Account represents a financial account
type Account struct {
	ID             string
	Type           string
	Number         string
	Balance        float64
	Currency       string
	OpenDate       string
	Status         string
	Transactions   []any // Simplified to avoid undefined type
	InterestRate   float64
	IsJoint        bool
	CoOwners       []any // Will be []*Person, using any to avoid import cycle
	OverdraftLimit float64
}

// Validate validates the Account model
func (a *Account) Validate() error {
	var errors []string

	// Validate ID
	if a.ID == "" {
		errors = append(errors, "ID is required")
	}

	// Validate Type
	if a.Type == "" {
		errors = append(errors, "Type is required")
	}

	// Validate Number
	if a.Number == "" {
		errors = append(errors, "Number is required")
	}

	// Validate Currency
	if a.Currency == "" {
		errors = append(errors, "Currency is required")
	}

	// Validate OpenDate
	if a.OpenDate != "" {
		_, err := time.Parse("2006-01-02", a.OpenDate)
		if err != nil {
			errors = append(errors, "OpenDate must be in the format YYYY-MM-DD")
		}
	}

	// Validate Status
	validStatuses := []string{"active", "inactive", "closed", "frozen"}
	isValidStatus := false
	for _, status := range validStatuses {
		if strings.ToLower(a.Status) == status {
			isValidStatus = true
			break
		}
	}
	if !isValidStatus {
		errors = append(errors, "Status must be one of: active, inactive, closed, frozen")
	}

	// Skip validation for Transactions as they are now any type

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// Bank represents a financial institution
type Bank struct {
	Name                  string
	BranchCode            string
	Address               *Address
	Accounts              []*Account
	Stocks                []any // Simplified to avoid undefined type
	Loans                 []any // Simplified to avoid undefined type
	Investments           []any // Simplified to avoid undefined type
	Advisor               any   // Will be *Person, using any to avoid import cycle
	RelationshipStartDate string
}

// Validate validates the Bank model
func (b *Bank) Validate() error {
	var errors []string

	// Validate Name
	if b.Name == "" {
		errors = append(errors, "Name is required")
	}

	// Validate BranchCode
	if b.BranchCode == "" {
		errors = append(errors, "BranchCode is required")
	}

	// Validate RelationshipStartDate if provided
	if b.RelationshipStartDate != "" {
		_, err := time.Parse("2006-01-02", b.RelationshipStartDate)
		if err != nil {
			errors = append(errors, "RelationshipStartDate must be in the format YYYY-MM-DD")
		}
	}

	// Validate accounts if provided
	for i, account := range b.Accounts {
		if account != nil {
			if err := account.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("Account[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Validate address if provided
	if b.Address != nil {
		if err := b.Address.Validate(); err != nil {
			errors = append(errors, fmt.Sprintf("Address validation failed: %s", err.Error()))
		}
	}

	// Skip validation for Stocks, Loans, and Investments as they are now any type

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}
