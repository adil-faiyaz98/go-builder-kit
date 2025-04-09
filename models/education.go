package models

import (
	"fmt"
	"strings"
	"time"
)

// Education represents educational background
type Education struct {
	Degree      string
	Institution string
	Location    *Address
	StartDate   string
	EndDate     string
	GPA         float64
	Honors      []string
	Major       string
	Minor       string
	Courses     []*Course
	Activities  []string
}

// Validate validates the Education model
func (e *Education) Validate() error {
	var errors []string

	// Validate required fields
	if e.Degree == "" {
		errors = append(errors, "Degree cannot be empty")
	}

	if e.Institution == "" {
		errors = append(errors, "Institution cannot be empty")
	}

	// Validate dates if provided
	if e.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", e.StartDate)
		if err != nil {
			errors = append(errors, "StartDate must be in the format YYYY-MM-DD")
		} else {
			// Check if start date is in the future
			if startDate.After(time.Now()) {
				errors = append(errors, "StartDate cannot be in the future")
			}

			// Check if start date is too far in the past
			if time.Now().Year()-startDate.Year() > 100 {
				errors = append(errors, "StartDate cannot be more than 100 years ago")
			}

			// Check end date if provided
			if e.EndDate != "" {
				endDate, err := time.Parse("2006-01-02", e.EndDate)
				if err != nil {
					errors = append(errors, "EndDate must be in the format YYYY-MM-DD")
				} else {
					// Check if end date is before start date
					if endDate.Before(startDate) {
						errors = append(errors, "EndDate cannot be before StartDate")
					}
				}
			}
		}
	}

	// Validate GPA range
	if e.GPA < 0 {
		errors = append(errors, "GPA cannot be negative")
	} else if e.GPA > 4.0 {
		errors = append(errors, "GPA cannot be greater than 4.0")
	}

	// Validate Location if provided
	if e.Location != nil {
		if err := e.Location.Validate(); err != nil {
			errors = append(errors, fmt.Sprintf("Location validation failed: %s", err.Error()))
		}
	}

	// Validate Courses if provided
	for i, course := range e.Courses {
		if course != nil {
			if err := course.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("Course[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// Course represents an academic course
type Course struct {
	Code        string
	Name        string
	Description string
	Credits     float64
	Grade       string
	Semester    string
	Year        int
	Instructor  string
}

// Validate validates the Course model
func (c *Course) Validate() error {
	var errors []string

	// Validate required fields
	if c.Code == "" {
		errors = append(errors, "Code cannot be empty")
	}

	if c.Name == "" {
		errors = append(errors, "Name cannot be empty")
	}

	// Validate Credits
	if c.Credits < 0 {
		errors = append(errors, "Credits cannot be negative")
	}

	// Validate Year
	currentYear := time.Now().Year()
	if c.Year < 1900 {
		errors = append(errors, "Year cannot be before 1900")
	} else if c.Year > currentYear+1 { // Allow for next year's courses
		errors = append(errors, fmt.Sprintf("Year cannot be after %d", currentYear+1))
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}
