package models

import (
	"fmt"
	"strings"
	"time"
)

// Transaction represents a financial transaction
type Transaction struct {
	ID          string
	Type        string
	Amount      float64
	Currency    string
	Date        string
	Description string
	Category    string
	Status      string
	FromAccount *Account
	ToAccount   *Account
	Reference   string
	Tags        []string
}

// Validate validates the Transaction model
func (t *Transaction) Validate() error {
	var errors []string

	// Validate ID
	if t.ID == "" {
		errors = append(errors, "ID is required")
	}

	// Validate Type
	validTypes := []string{"deposit", "withdrawal", "transfer", "payment", "refund", "fee", "interest", "other"}
	isValidType := false
	for _, validType := range validTypes {
		if strings.ToLower(t.Type) == validType {
			isValidType = true
			break
		}
	}
	if !isValidType {
		errors = append(errors, "Type must be one of: deposit, withdrawal, transfer, payment, refund, fee, interest, other")
	}

	// Validate Currency
	if t.Currency == "" {
		errors = append(errors, "Currency is required")
	}

	// Validate Date
	if t.Date != "" {
		_, err := time.Parse("2006-01-02", t.Date)
		if err != nil {
			errors = append(errors, "Date must be in the format YYYY-MM-DD")
		}
	}

	// Validate Status
	validStatuses := []string{"pending", "completed", "failed", "cancelled", "refunded"}
	isValidStatus := false
	for _, status := range validStatuses {
		if strings.ToLower(t.Status) == status {
			isValidStatus = true
			break
		}
	}
	if !isValidStatus {
		errors = append(errors, "Status must be one of: pending, completed, failed, cancelled, refunded")
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}
