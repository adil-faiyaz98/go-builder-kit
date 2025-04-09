package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// MedicationBuilder builds a Medication model
type MedicationBuilder struct {
	medication *models.Medication
	// Custom validation functions
	validationFuncs []func(*models.Medication) error
}

// NewMedicationBuilder creates a new MedicationBuilder
func NewMedicationBuilder() *MedicationBuilder {
	return &MedicationBuilder{
		medication: &models.Medication{
			Name: "",
			Dosage: "",
			Frequency: "",
			StartDate: "",
			EndDate: "",
			PrescribedBy: "",
			Purpose: "",
			SideEffects: []string{},
		},
		validationFuncs: []func(*models.Medication) error{},
	}
}

// NewMedicationBuilderWithDefaults creates a new MedicationBuilder with sensible defaults
func NewMedicationBuilderWithDefaults() *MedicationBuilder {
	builder := NewMedicationBuilder()
	// Add default values here if needed
	return builder
}
// WithName sets the Name
func (b *MedicationBuilder) WithName(name string) *MedicationBuilder {
	b.medication.Name = name
	return b
}

// WithDosage sets the Dosage
func (b *MedicationBuilder) WithDosage(dosage string) *MedicationBuilder {
	b.medication.Dosage = dosage
	return b
}

// WithFrequency sets the Frequency
func (b *MedicationBuilder) WithFrequency(frequency string) *MedicationBuilder {
	b.medication.Frequency = frequency
	return b
}

// WithStartDate sets the StartDate
func (b *MedicationBuilder) WithStartDate(startDate string) *MedicationBuilder {
	b.medication.StartDate = startDate
	return b
}

// WithEndDate sets the EndDate
func (b *MedicationBuilder) WithEndDate(endDate string) *MedicationBuilder {
	b.medication.EndDate = endDate
	return b
}

// WithPrescribedBy sets the PrescribedBy
func (b *MedicationBuilder) WithPrescribedBy(prescribedBy string) *MedicationBuilder {
	b.medication.PrescribedBy = prescribedBy
	return b
}

// WithPurpose sets the Purpose
func (b *MedicationBuilder) WithPurpose(purpose string) *MedicationBuilder {
	b.medication.Purpose = purpose
	return b
}

// WithSideEffects sets the SideEffects
func (b *MedicationBuilder) WithSideEffects(sideEffects string) *MedicationBuilder {
	b.medication.SideEffects = append(b.medication.SideEffects, sideEffects)
	return b
}


// WithValidation adds a custom validation function
func (b *MedicationBuilder) WithValidation(validationFunc func(*models.Medication) error) *MedicationBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Medication
func (b *MedicationBuilder) Build() interface{} {
	return b.medication
}

// BuildPtr builds the Medication and returns a pointer
func (b *MedicationBuilder) BuildPtr() *models.Medication {
	return b.medication
}

// BuildAndValidate builds the Medication and validates it
func (b *MedicationBuilder) BuildAndValidate() (*models.Medication, error) {
	medication := b.medication

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(medication); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(medication).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return medication, err
		}
	}

	return medication, nil
}

// MustBuild builds the Medication and panics if validation fails
func (b *MedicationBuilder) MustBuild() *models.Medication {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *MedicationBuilder) Clone() *MedicationBuilder {
	clonedMedication := *b.medication
	return &MedicationBuilder{
		medication: &clonedMedication,
		validationFuncs: append([]func(*models.Medication) error{}, b.validationFuncs...),
	}
}
