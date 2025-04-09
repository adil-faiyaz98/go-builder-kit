package models

import (
	"fmt"
	"strings"
	"time"
)

// Loan represents a financial loan
type Loan struct {
	ID               string
	Amount           float64
	InterestRate     float64
	StartDate        string
	EndDate          string
	MonthlyPayment   float64
	RemainingBalance float64
	Type             string // Mortgage, Auto, Personal, Student
	Collateral       string
	Status           string // Active, Paid, Defaulted
	Payments         []*Payment
}

// Validate validates the Loan model
func (l *Loan) Validate() error {
	var errors []string

	// Validate ID
	if l.ID == "" {
		errors = append(errors, "ID is required")
	}

	// Validate Amount
	if l.Amount <= 0 {
		errors = append(errors, "Amount must be greater than 0")
	}

	// Validate InterestRate
	if l.InterestRate < 0 {
		errors = append(errors, "InterestRate cannot be negative")
	}

	// Validate StartDate
	if l.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", l.StartDate)
		if err != nil {
			errors = append(errors, "StartDate must be in the format YYYY-MM-DD")
		} else {
			// Validate EndDate if provided
			if l.EndDate != "" {
				endDate, err := time.Parse("2006-01-02", l.EndDate)
				if err != nil {
					errors = append(errors, "EndDate must be in the format YYYY-MM-DD")
				} else if endDate.Before(startDate) {
					errors = append(errors, "EndDate cannot be before StartDate")
				}
			}
		}
	}

	// Validate MonthlyPayment
	if l.MonthlyPayment <= 0 {
		errors = append(errors, "MonthlyPayment must be greater than 0")
	}

	// Validate RemainingBalance
	if l.RemainingBalance < 0 {
		errors = append(errors, "RemainingBalance cannot be negative")
	}

	// Validate Type
	validTypes := []string{"mortgage", "auto", "personal", "student", "business", "other"}
	isValidType := false
	for _, validType := range validTypes {
		if strings.ToLower(l.Type) == validType {
			isValidType = true
			break
		}
	}
	if !isValidType {
		errors = append(errors, "Type must be one of: mortgage, auto, personal, student, business, other")
	}

	// Validate Status
	validStatuses := []string{"active", "paid", "defaulted", "refinanced", "closed"}
	isValidStatus := false
	for _, status := range validStatuses {
		if strings.ToLower(l.Status) == status {
			isValidStatus = true
			break
		}
	}
	if !isValidStatus {
		errors = append(errors, "Status must be one of: active, paid, defaulted, refinanced, closed")
	}

	// Validate Payments if provided
	for i, payment := range l.Payments {
		if payment != nil {
			if err := payment.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("Payment[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// Payment represents a loan payment
type Payment struct {
	Date        string
	Amount      float64
	Principal   float64
	Interest    float64
	Status      string // Scheduled, Paid, Late, Missed
	Transaction *Transaction
}

// Validate validates the Payment model
func (p *Payment) Validate() error {
	var errors []string

	// Validate Date
	if p.Date != "" {
		_, err := time.Parse("2006-01-02", p.Date)
		if err != nil {
			errors = append(errors, "Date must be in the format YYYY-MM-DD")
		}
	}

	// Validate Amount
	if p.Amount <= 0 {
		errors = append(errors, "Amount must be greater than 0")
	}

	// Validate Principal
	if p.Principal < 0 {
		errors = append(errors, "Principal cannot be negative")
	}

	// Validate Interest
	if p.Interest < 0 {
		errors = append(errors, "Interest cannot be negative")
	}

	// Validate Status
	validStatuses := []string{"scheduled", "paid", "late", "missed"}
	isValidStatus := false
	for _, status := range validStatuses {
		if strings.ToLower(p.Status) == status {
			isValidStatus = true
			break
		}
	}
	if !isValidStatus {
		errors = append(errors, "Status must be one of: scheduled, paid, late, missed")
	}

	// Validate Transaction if provided
	if p.Transaction != nil {
		if err := p.Transaction.Validate(); err != nil {
			errors = append(errors, fmt.Sprintf("Transaction validation failed: %s", err.Error()))
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}
