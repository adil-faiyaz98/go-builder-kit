package simple

import (
	"fmt"
	"regexp"
	"strings"
)

// Person represents a simple person model
type Person struct {
	ID        string
	Name      string
	Age       int
	Email     string
	Address   *Address
	Tags      []string
	Metadata  map[string]string
}

// Validate validates the Person model
func (p *Person) Validate() error {
	var errors []string

	// Validate ID
	if p.ID == "" {
		errors = append(errors, "ID cannot be empty")
	}

	// Validate Name
	if p.Name == "" {
		errors = append(errors, "Name cannot be empty")
	}

	// Validate Age
	if p.Age < 0 {
		errors = append(errors, "Age cannot be negative")
	}

	// Validate Email if provided
	if p.Email != "" {
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(p.Email) {
			errors = append(errors, "Email is not valid")
		}
	}

	// Validate Address if provided
	if p.Address != nil {
		if err := p.Address.Validate(); err != nil {
			errors = append(errors, fmt.Sprintf("Address validation failed: %s", err.Error()))
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// Address represents a physical location
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
	Country    string
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

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}
