package models

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Person represents an individual with all their associated data
type Person struct {
	ID            string
	Name          string
	Age           int
	Email         string
	Phone         string
	Birthdate     string
	Gender        string
	Nationality   string
	MaritalStatus string
	Address       *Address
	Education     *Education
	Profile       any
	Bank          *Bank
	Employment    *Employment
	Friends       []*Person
	Family        []*FamilyMember
	Health        any
	Digital       any
	TravelHistory []*Travel
	Preferences   *PersonalPreferences
	CreatedAt     string
	UpdatedAt     string
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
	} else if len(p.Name) < 2 {
		errors = append(errors, "Name must be at least 2 characters long")
	}

	// Validate Age
	if p.Age < 0 {
		errors = append(errors, "Age cannot be negative")
	} else if p.Age > 150 {
		errors = append(errors, "Age cannot be greater than 150")
	}

	// Validate Email if provided
	if p.Email != "" {
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(p.Email) {
			errors = append(errors, "Email is not valid")
		}
	}

	// Validate Phone if provided
	if p.Phone != "" {
		// This is a simple regex that matches common phone number formats
		// You might want to use country-specific validation in a real application
		phoneRegex := regexp.MustCompile(`^[+]?[\d\s\(\)\-]{7,20}$`)
		if !phoneRegex.MatchString(p.Phone) {
			errors = append(errors, "Phone number is not valid")
		}
	}

	// Validate Birthdate if provided
	if p.Birthdate != "" {
		birthdate, err := time.Parse("2006-01-02", p.Birthdate)
		if err != nil {
			errors = append(errors, "Birthdate must be in the format YYYY-MM-DD")
		} else {
			// Check if birthdate is in the future
			if birthdate.After(time.Now()) {
				errors = append(errors, "Birthdate cannot be in the future")
			}

			// Check if birthdate is too far in the past
			if time.Now().Year()-birthdate.Year() > 150 {
				errors = append(errors, "Birthdate cannot be more than 150 years ago")
			}
		}
	}

	// Validate Gender if provided
	if p.Gender != "" {
		validGenders := []string{"male", "female", "non-binary", "other", "prefer not to say"}
		isValidGender := false
		for _, gender := range validGenders {
			if strings.ToLower(p.Gender) == gender {
				isValidGender = true
				break
			}
		}
		if !isValidGender {
			errors = append(errors, "Gender must be one of: male, female, non-binary, other, prefer not to say")
		}
	}

	// Validate MaritalStatus if provided
	if p.MaritalStatus != "" {
		validStatuses := []string{"single", "married", "divorced", "widowed", "separated", "other"}
		isValidStatus := false
		for _, status := range validStatuses {
			if strings.ToLower(p.MaritalStatus) == status {
				isValidStatus = true
				break
			}
		}
		if !isValidStatus {
			errors = append(errors, "MaritalStatus must be one of: single, married, divorced, widowed, separated, other")
		}
	}

	// Validate Address if provided
	if p.Address != nil {
		if err := p.Address.Validate(); err != nil {
			errors = append(errors, fmt.Sprintf("Address validation failed: %s", err.Error()))
		}
	}

	// Validate Education if provided
	if p.Education != nil {
		if err := p.Education.Validate(); err != nil {
			errors = append(errors, fmt.Sprintf("Education validation failed: %s", err.Error()))
		}
	}

	// Skip validation for Profile as it is now any type

	// Validate Bank if provided
	if p.Bank != nil {
		if err := p.Bank.Validate(); err != nil {
			errors = append(errors, fmt.Sprintf("Bank validation failed: %s", err.Error()))
		}
	}

	// Validate Employment if provided
	if p.Employment != nil {
		if err := p.Employment.Validate(); err != nil {
			errors = append(errors, fmt.Sprintf("Employment validation failed: %s", err.Error()))
		}
	}

	// Skip validation for Health as it is now any type

	// Skip validation for Digital as it is now any type

	// Validate Preferences if provided
	if p.Preferences != nil {
		if err := p.Preferences.Validate(); err != nil {
			errors = append(errors, fmt.Sprintf("Preferences validation failed: %s", err.Error()))
		}
	}

	// Validate TravelHistory if provided
	for i, travel := range p.TravelHistory {
		if travel != nil {
			if err := travel.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("Travel[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Validate Family if provided
	for i, member := range p.Family {
		if member != nil {
			if err := member.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("FamilyMember[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Validate CreatedAt if provided
	if p.CreatedAt != "" {
		_, err := time.Parse("2006-01-02T15:04:05Z", p.CreatedAt)
		if err != nil {
			errors = append(errors, "CreatedAt must be in the format YYYY-MM-DDThh:mm:ssZ")
		}
	}

	// Validate UpdatedAt if provided
	if p.UpdatedAt != "" {
		_, err := time.Parse("2006-01-02T15:04:05Z", p.UpdatedAt)
		if err != nil {
			errors = append(errors, "UpdatedAt must be in the format YYYY-MM-DDThh:mm:ssZ")
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// FamilyMember represents a family relationship
type FamilyMember struct {
	Person       *Person
	Relationship string
}

// Validate validates the FamilyMember model
func (f *FamilyMember) Validate() error {
	var errors []string

	// Validate Relationship
	if f.Relationship == "" {
		errors = append(errors, "Relationship cannot be empty")
	} else {
		validRelationships := []string{
			"spouse", "partner", "child", "parent", "sibling",
			"grandparent", "grandchild", "aunt", "uncle", "cousin",
			"niece", "nephew", "in-law", "other",
		}
		isValidRelationship := false
		for _, rel := range validRelationships {
			if strings.ToLower(f.Relationship) == rel {
				isValidRelationship = true
				break
			}
		}
		if !isValidRelationship {
			errors = append(errors, "Relationship must be one of: spouse, partner, child, parent, sibling, grandparent, grandchild, aunt, uncle, cousin, niece, nephew, in-law, other")
		}
	}

	// Validate Person if provided
	if f.Person != nil {
		if err := f.Person.Validate(); err != nil {
			errors = append(errors, fmt.Sprintf("Person validation failed: %s", err.Error()))
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// Travel represents a trip or journey
type Travel struct {
	Destination    Address
	StartDate      string
	EndDate        string
	Purpose        string
	Accommodation  string
	Transportation string
	Activities     []string
	Expenses       float64
}

// Validate validates the Travel model
func (t *Travel) Validate() error {
	var errors []string

	// Validate Destination
	if err := t.Destination.Validate(); err != nil {
		errors = append(errors, fmt.Sprintf("Destination validation failed: %s", err.Error()))
	}

	// Validate StartDate
	if t.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", t.StartDate)
		if err != nil {
			errors = append(errors, "StartDate must be in the format YYYY-MM-DD")
		} else {
			// Validate EndDate if provided
			if t.EndDate != "" {
				endDate, err := time.Parse("2006-01-02", t.EndDate)
				if err != nil {
					errors = append(errors, "EndDate must be in the format YYYY-MM-DD")
				} else if endDate.Before(startDate) {
					errors = append(errors, "EndDate cannot be before StartDate")
				}
			}
		}
	}

	// Validate Purpose
	if t.Purpose != "" {
		validPurposes := []string{"business", "leisure", "education", "medical", "family", "other"}
		isValidPurpose := false
		for _, purpose := range validPurposes {
			if strings.ToLower(t.Purpose) == purpose {
				isValidPurpose = true
				break
			}
		}
		if !isValidPurpose {
			errors = append(errors, "Purpose must be one of: business, leisure, education, medical, family, other")
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// PersonalPreferences represents a person's preferences
type PersonalPreferences struct {
	FavoriteColor       string
	FavoriteColors      []string
	FavoriteFood        string
	FavoriteFoods       []string
	FavoriteMusic       string
	MusicTastes         []string
	FavoriteMovie       string
	MovieGenres         []string
	FavoriteBook        string
	BookGenres          []string
	FavoriteSport       string
	FavoriteAnimal      string
	Hobbies             []string
	Interests           []string
	Languages           []string
	TravelPreferences   map[string]string
	ShoppingPreferences map[string]bool
}

// Validate validates the PersonalPreferences model
func (p *PersonalPreferences) Validate() error {
	// No validation rules for preferences as they are subjective
	return nil
}
