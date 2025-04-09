package models

import (
	"fmt"
	"regexp"
	"strings"
)

// Address represents a physical location
type Address struct {
	Street      string
	City        string
	State       string
	PostalCode  string
	Country     string
	Coordinates *GeoLocation
	Type        string // Home, Work, etc.
	IsPrimary   bool
}

// Validate validates the Address model
func (a *Address) Validate() error {
	var errors []string

	// Validate required fields
	if a.Street == "" {
		errors = append(errors, "Street cannot be empty")
	}

	if a.City == "" {
		errors = append(errors, "City cannot be empty")
	}

	if a.Country == "" {
		errors = append(errors, "Country cannot be empty")
	}

	// Validate postal code format if provided
	if a.PostalCode != "" {
		// This is a simple regex that matches common postal code formats
		// You might want to use country-specific validation in a real application
		postalCodeRegex := regexp.MustCompile(`^[0-9A-Za-z\-\s]{3,10}$`)
		if !postalCodeRegex.MatchString(a.PostalCode) {
			errors = append(errors, "PostalCode format is invalid")
		}
	}

	// Validate address type if provided
	if a.Type != "" {
		validTypes := map[string]bool{
			"Home":    true,
			"Work":    true,
			"Mailing": true,
			"Billing": true,
			"Other":   true,
		}

		if !validTypes[a.Type] {
			errors = append(errors, "Type must be one of: Home, Work, Mailing, Billing, Other")
		}
	}

	// Validate coordinates if provided
	if a.Coordinates != nil {
		if err := a.Coordinates.Validate(); err != nil {
			errors = append(errors, fmt.Sprintf("Coordinates validation failed: %s", err.Error()))
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// GeoLocation represents geographical coordinates
type GeoLocation struct {
	Latitude  float64
	Longitude float64
	Accuracy  float64
}

// Validate validates the GeoLocation model
func (g *GeoLocation) Validate() error {
	var errors []string

	// Validate latitude range (-90 to 90)
	if g.Latitude < -90 || g.Latitude > 90 {
		errors = append(errors, "Latitude must be between -90 and 90")
	}

	// Validate longitude range (-180 to 180)
	if g.Longitude < -180 || g.Longitude > 180 {
		errors = append(errors, "Longitude must be between -180 and 180")
	}

	// Validate accuracy (must be positive)
	if g.Accuracy < 0 {
		errors = append(errors, "Accuracy cannot be negative")
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}
