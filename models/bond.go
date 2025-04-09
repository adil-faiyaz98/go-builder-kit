package models

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Bond represents a bond investment
type Bond struct {
	ID               string
	ISIN             string
	Name             string
	Issuer           string
	Type             string
	FaceValue        float64
	CouponRate       float64
	MaturityDate     string
	PurchaseDate     string
	PurchasePrice    float64
	CurrentPrice     float64
	Quantity         int
	Currency         string
	PaymentFrequency string
	Rating           string
	Yield            float64
}

// Validate validates the Bond model
func (b *Bond) Validate() error {
	var errors []string

	// Validate ISIN
	if b.ISIN == "" {
		errors = append(errors, "ISIN is required")
	} else {
		// ISIN is a 12-character alphanumeric code
		isinRegex := regexp.MustCompile(`^[A-Z]{2}[A-Z0-9]{9}[0-9]$`)
		if !isinRegex.MatchString(b.ISIN) {
			errors = append(errors, "ISIN format is invalid")
		}
	}

	// Validate Name
	if b.Name == "" {
		errors = append(errors, "Name is required")
	}

	// Validate Issuer
	if b.Issuer == "" {
		errors = append(errors, "Issuer is required")
	}

	// Validate Type
	validTypes := []string{"government", "corporate", "municipal", "treasury", "zero-coupon", "other"}
	isValidType := false
	for _, validType := range validTypes {
		if strings.ToLower(b.Type) == validType {
			isValidType = true
			break
		}
	}
	if !isValidType {
		errors = append(errors, "Type must be one of: government, corporate, municipal, treasury, zero-coupon, other")
	}

	// Validate FaceValue
	if b.FaceValue <= 0 {
		errors = append(errors, "FaceValue must be greater than 0")
	}

	// Validate CouponRate
	if b.CouponRate < 0 {
		errors = append(errors, "CouponRate cannot be negative")
	}

	// Validate MaturityDate
	if b.MaturityDate != "" {
		maturityDate, err := time.Parse("2006-01-02", b.MaturityDate)
		if err != nil {
			errors = append(errors, "MaturityDate must be in the format YYYY-MM-DD")
		} else {
			// Validate PurchaseDate if provided
			if b.PurchaseDate != "" {
				purchaseDate, err := time.Parse("2006-01-02", b.PurchaseDate)
				if err != nil {
					errors = append(errors, "PurchaseDate must be in the format YYYY-MM-DD")
				} else if maturityDate.Before(purchaseDate) {
					errors = append(errors, "MaturityDate cannot be before PurchaseDate")
				}
			}
		}
	}

	// Validate PurchasePrice
	if b.PurchasePrice <= 0 {
		errors = append(errors, "PurchasePrice must be greater than 0")
	}

	// Validate CurrentPrice
	if b.CurrentPrice <= 0 {
		errors = append(errors, "CurrentPrice must be greater than 0")
	}

	// Validate Quantity
	if b.Quantity <= 0 {
		errors = append(errors, "Quantity must be greater than 0")
	}

	// Validate Currency
	if b.Currency == "" {
		errors = append(errors, "Currency is required")
	}

	// Validate PaymentFrequency
	validFrequencies := []string{"annual", "semi-annual", "quarterly", "monthly", "none"}
	isValidFrequency := false
	for _, frequency := range validFrequencies {
		if strings.ToLower(b.PaymentFrequency) == frequency {
			isValidFrequency = true
			break
		}
	}
	if !isValidFrequency {
		errors = append(errors, "PaymentFrequency must be one of: annual, semi-annual, quarterly, monthly, none")
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}
