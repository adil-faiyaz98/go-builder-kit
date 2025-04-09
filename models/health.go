package models

import (
	"fmt"
	"strings"
	"time"
)

// HealthProfile represents a person's health information
type HealthProfile struct {
	Height            float64 // in cm
	Weight            float64 // in kg
	BloodType         string
	Allergies         []string
	ChronicConditions []string
	Medications       []*Medication
	MedicalHistory    []*MedicalRecord
	Insurance         *Insurance
	EmergencyContact  any // Will be *Person, using any to avoid import cycle
	PrimaryPhysician  any // Will be *Person, using any to avoid import cycle
}

// Validate validates the HealthProfile model
func (h *HealthProfile) Validate() error {
	var errors []string

	// Validate Height
	if h.Height < 0 {
		errors = append(errors, "Height cannot be negative")
	} else if h.Height > 300 {
		errors = append(errors, "Height cannot be greater than 300 cm")
	}

	// Validate Weight
	if h.Weight < 0 {
		errors = append(errors, "Weight cannot be negative")
	} else if h.Weight > 700 {
		errors = append(errors, "Weight cannot be greater than 700 kg")
	}

	// Validate BloodType
	if h.BloodType != "" {
		validBloodTypes := []string{"A+", "A-", "B+", "B-", "AB+", "AB-", "O+", "O-"}
		isValidBloodType := false
		for _, bloodType := range validBloodTypes {
			if h.BloodType == bloodType {
				isValidBloodType = true
				break
			}
		}
		if !isValidBloodType {
			errors = append(errors, "BloodType must be one of: A+, A-, B+, B-, AB+, AB-, O+, O-")
		}
	}

	// Validate Medications if provided
	for i, medication := range h.Medications {
		if medication != nil {
			if err := medication.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("Medication[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Validate MedicalHistory if provided
	for i, record := range h.MedicalHistory {
		if record != nil {
			if err := record.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("MedicalRecord[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Validate Insurance if provided
	if h.Insurance != nil {
		if err := h.Insurance.Validate(); err != nil {
			errors = append(errors, fmt.Sprintf("Insurance validation failed: %s", err.Error()))
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// Medication represents a prescribed medication
type Medication struct {
	Name         string
	Dosage       string
	Frequency    string
	StartDate    string
	EndDate      string
	PrescribedBy string
	Purpose      string
	SideEffects  []string
}

// Validate validates the Medication model
func (m *Medication) Validate() error {
	var errors []string

	// Validate Name
	if m.Name == "" {
		errors = append(errors, "Name cannot be empty")
	}

	// Validate Dosage
	if m.Dosage == "" {
		errors = append(errors, "Dosage cannot be empty")
	}

	// Validate Frequency
	if m.Frequency == "" {
		errors = append(errors, "Frequency cannot be empty")
	}

	// Validate StartDate if provided
	if m.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", m.StartDate)
		if err != nil {
			errors = append(errors, "StartDate must be in the format YYYY-MM-DD")
		} else {
			// Validate EndDate if provided
			if m.EndDate != "" {
				endDate, err := time.Parse("2006-01-02", m.EndDate)
				if err != nil {
					errors = append(errors, "EndDate must be in the format YYYY-MM-DD")
				} else if endDate.Before(startDate) {
					errors = append(errors, "EndDate cannot be before StartDate")
				}
			}
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// MedicalRecord represents a medical event or visit
type MedicalRecord struct {
	Date        string
	Type        string // Checkup, Emergency, Surgery, etc.
	Provider    string
	Diagnosis   string
	Treatment   string
	Notes       string
	Attachments []string // URLs to documents
	Location    *Address
	FollowUp    string
}

// Validate validates the MedicalRecord model
func (m *MedicalRecord) Validate() error {
	var errors []string

	// Validate Date
	if m.Date == "" {
		errors = append(errors, "Date cannot be empty")
	} else {
		_, err := time.Parse("2006-01-02", m.Date)
		if err != nil {
			errors = append(errors, "Date must be in the format YYYY-MM-DD")
		}
	}

	// Validate Type
	if m.Type == "" {
		errors = append(errors, "Type cannot be empty")
	}

	// Validate Provider
	if m.Provider == "" {
		errors = append(errors, "Provider cannot be empty")
	}

	// Validate Location if provided
	if m.Location != nil {
		if err := m.Location.Validate(); err != nil {
			errors = append(errors, fmt.Sprintf("Location validation failed: %s", err.Error()))
		}
	}

	// Validate FollowUp if provided
	if m.FollowUp != "" {
		_, err := time.Parse("2006-01-02", m.FollowUp)
		if err != nil {
			errors = append(errors, "FollowUp must be in the format YYYY-MM-DD")
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// Insurance represents health insurance coverage
type Insurance struct {
	Provider       string
	PolicyNumber   string
	Type           string // Health, Dental, Vision, Life
	StartDate      string
	EndDate        string
	Coverage       map[string]float64 // Type of service to coverage percentage
	Deductible     float64
	CoPay          float64
	Copay          float64
	OutOfPocketMax float64
	Beneficiaries  []any // Will be []*Person, using any to avoid import cycle
}

// Validate validates the Insurance model
func (i *Insurance) Validate() error {
	var errors []string

	// Validate Provider
	if i.Provider == "" {
		errors = append(errors, "Provider cannot be empty")
	}

	// Validate PolicyNumber
	if i.PolicyNumber == "" {
		errors = append(errors, "PolicyNumber cannot be empty")
	}

	// Validate Type
	if i.Type != "" {
		validTypes := []string{"health", "dental", "vision", "life"}
		isValidType := false
		for _, validType := range validTypes {
			if strings.ToLower(i.Type) == validType {
				isValidType = true
				break
			}
		}
		if !isValidType {
			errors = append(errors, "Type must be one of: health, dental, vision, life")
		}
	}

	// Validate StartDate if provided
	if i.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", i.StartDate)
		if err != nil {
			errors = append(errors, "StartDate must be in the format YYYY-MM-DD")
		} else {
			// Validate EndDate if provided
			if i.EndDate != "" {
				endDate, err := time.Parse("2006-01-02", i.EndDate)
				if err != nil {
					errors = append(errors, "EndDate must be in the format YYYY-MM-DD")
				} else if endDate.Before(startDate) {
					errors = append(errors, "EndDate cannot be before StartDate")
				}
			}
		}
	}

	// Validate Coverage if provided
	if i.Coverage != nil {
		for service, percentage := range i.Coverage {
			if percentage < 0 || percentage > 100 {
				errors = append(errors, fmt.Sprintf("Coverage percentage for %s must be between 0 and 100", service))
			}
		}
	}

	// Validate Deductible
	if i.Deductible < 0 {
		errors = append(errors, "Deductible cannot be negative")
	}

	// Validate CoPay
	if i.CoPay < 0 {
		errors = append(errors, "CoPay cannot be negative")
	}

	// Validate OutOfPocketMax
	if i.OutOfPocketMax < 0 {
		errors = append(errors, "OutOfPocketMax cannot be negative")
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}
