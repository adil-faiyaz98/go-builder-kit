package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// HealthProfileBuilder builds a HealthProfile model
type HealthProfileBuilder struct {
	healthProfile *models.HealthProfile
	// Custom validation functions
	validationFuncs []func(*models.HealthProfile) error
}

// NewHealthProfileBuilder creates a new HealthProfileBuilder
func NewHealthProfileBuilder() *HealthProfileBuilder {
	return &HealthProfileBuilder{
		healthProfile: &models.HealthProfile{
			Height: 0.0,
			Weight: 0.0,
			BloodType: "",
			Allergies: []string{},
			ChronicConditions: []string{},
			Medications: []*models.Medication{},
			MedicalHistory: []*models.MedicalRecord{},
			Insurance: nil,
			EmergencyContact: any(0),
			PrimaryPhysician: any(0),
		},
		validationFuncs: []func(*models.HealthProfile) error{},
	}
}

// NewHealthProfileBuilderWithDefaults creates a new HealthProfileBuilder with sensible defaults
func NewHealthProfileBuilderWithDefaults() *HealthProfileBuilder {
	builder := NewHealthProfileBuilder()
	// Add default values here if needed
	return builder
}
// WithHeight sets the Height
func (b *HealthProfileBuilder) WithHeight(height float64) *HealthProfileBuilder {
	b.healthProfile.Height = height
	return b
}

// WithWeight sets the Weight
func (b *HealthProfileBuilder) WithWeight(weight float64) *HealthProfileBuilder {
	b.healthProfile.Weight = weight
	return b
}

// WithBloodType sets the BloodType
func (b *HealthProfileBuilder) WithBloodType(bloodType string) *HealthProfileBuilder {
	b.healthProfile.BloodType = bloodType
	return b
}

// WithAllergies sets the Allergies
func (b *HealthProfileBuilder) WithAllergies(allergies string) *HealthProfileBuilder {
	b.healthProfile.Allergies = append(b.healthProfile.Allergies, allergies)
	return b
}

// WithChronicConditions sets the ChronicConditions
func (b *HealthProfileBuilder) WithChronicConditions(chronicConditions string) *HealthProfileBuilder {
	b.healthProfile.ChronicConditions = append(b.healthProfile.ChronicConditions, chronicConditions)
	return b
}

// WithMedications sets the Medications
func (b *HealthProfileBuilder) WithMedications(medications *MedicationBuilder) *HealthProfileBuilder {
	// Ensure the slice is initialized
	if b.healthProfile.Medications == nil {
		b.healthProfile.Medications = []*models.Medication{}
	}
	// Handle nested slice element
	builtValue := medications.Build().(*models.Medication)
	b.healthProfile.Medications = append(b.healthProfile.Medications, builtValue)
	return b
}

// WithMedicalHistory sets the MedicalHistory
func (b *HealthProfileBuilder) WithMedicalHistory(medicalHistory *MedicalRecordBuilder) *HealthProfileBuilder {
	// Ensure the slice is initialized
	if b.healthProfile.MedicalHistory == nil {
		b.healthProfile.MedicalHistory = []*models.MedicalRecord{}
	}
	// Handle nested slice element
	builtValue := medicalHistory.Build().(*models.MedicalRecord)
	b.healthProfile.MedicalHistory = append(b.healthProfile.MedicalHistory, builtValue)
	return b
}

// WithInsurance sets the Insurance
func (b *HealthProfileBuilder) WithInsurance(insurance *InsuranceBuilder) *HealthProfileBuilder {
	// Handle nested pointer
	b.healthProfile.Insurance = insurance.BuildPtr()
	return b
}

// WithEmergencyContact sets the EmergencyContact
func (b *HealthProfileBuilder) WithEmergencyContact(emergencyContact any) *HealthProfileBuilder {
	b.healthProfile.EmergencyContact = emergencyContact
	return b
}

// WithPrimaryPhysician sets the PrimaryPhysician
func (b *HealthProfileBuilder) WithPrimaryPhysician(primaryPhysician any) *HealthProfileBuilder {
	b.healthProfile.PrimaryPhysician = primaryPhysician
	return b
}


// WithValidation adds a custom validation function
func (b *HealthProfileBuilder) WithValidation(validationFunc func(*models.HealthProfile) error) *HealthProfileBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the HealthProfile
func (b *HealthProfileBuilder) Build() interface{} {
	return b.healthProfile
}

// BuildPtr builds the HealthProfile and returns a pointer
func (b *HealthProfileBuilder) BuildPtr() *models.HealthProfile {
	return b.healthProfile
}

// BuildAndValidate builds the HealthProfile and validates it
func (b *HealthProfileBuilder) BuildAndValidate() (*models.HealthProfile, error) {
	healthProfile := b.healthProfile

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(healthProfile); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(healthProfile).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return healthProfile, err
		}
	}

	return healthProfile, nil
}

// MustBuild builds the HealthProfile and panics if validation fails
func (b *HealthProfileBuilder) MustBuild() *models.HealthProfile {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *HealthProfileBuilder) Clone() *HealthProfileBuilder {
	clonedHealthProfile := *b.healthProfile
	return &HealthProfileBuilder{
		healthProfile: &clonedHealthProfile,
		validationFuncs: append([]func(*models.HealthProfile) error{}, b.validationFuncs...),
	}
}
