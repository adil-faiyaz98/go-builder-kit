package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// MedicalRecordBuilder builds a MedicalRecord model
type MedicalRecordBuilder struct {
	medicalRecord *models.MedicalRecord
	// Custom validation functions
	validationFuncs []func(*models.MedicalRecord) error
}

// NewMedicalRecordBuilder creates a new MedicalRecordBuilder
func NewMedicalRecordBuilder() *MedicalRecordBuilder {
	return &MedicalRecordBuilder{
		medicalRecord: &models.MedicalRecord{
			Date: "",
			Type: "",
			Provider: "",
			Diagnosis: "",
			Treatment: "",
			Notes: "",
			Attachments: []string{},
			Location: nil,
			FollowUp: "",
		},
		validationFuncs: []func(*models.MedicalRecord) error{},
	}
}

// NewMedicalRecordBuilderWithDefaults creates a new MedicalRecordBuilder with sensible defaults
func NewMedicalRecordBuilderWithDefaults() *MedicalRecordBuilder {
	builder := NewMedicalRecordBuilder()
	// Add default values here if needed
	return builder
}
// WithDate sets the Date
func (b *MedicalRecordBuilder) WithDate(date string) *MedicalRecordBuilder {
	b.medicalRecord.Date = date
	return b
}

// WithType sets the Type
func (b *MedicalRecordBuilder) WithType(value string) *MedicalRecordBuilder {
	b.medicalRecord.Type = value
	return b
}

// WithProvider sets the Provider
func (b *MedicalRecordBuilder) WithProvider(provider string) *MedicalRecordBuilder {
	b.medicalRecord.Provider = provider
	return b
}

// WithDiagnosis sets the Diagnosis
func (b *MedicalRecordBuilder) WithDiagnosis(diagnosis string) *MedicalRecordBuilder {
	b.medicalRecord.Diagnosis = diagnosis
	return b
}

// WithTreatment sets the Treatment
func (b *MedicalRecordBuilder) WithTreatment(treatment string) *MedicalRecordBuilder {
	b.medicalRecord.Treatment = treatment
	return b
}

// WithNotes sets the Notes
func (b *MedicalRecordBuilder) WithNotes(notes string) *MedicalRecordBuilder {
	b.medicalRecord.Notes = notes
	return b
}

// WithAttachments sets the Attachments
func (b *MedicalRecordBuilder) WithAttachments(attachments string) *MedicalRecordBuilder {
	b.medicalRecord.Attachments = append(b.medicalRecord.Attachments, attachments)
	return b
}

// WithLocation sets the Location
func (b *MedicalRecordBuilder) WithLocation(location *AddressBuilder) *MedicalRecordBuilder {
	// Handle nested pointer
	b.medicalRecord.Location = location.BuildPtr()
	return b
}

// WithFollowUp sets the FollowUp
func (b *MedicalRecordBuilder) WithFollowUp(followUp string) *MedicalRecordBuilder {
	b.medicalRecord.FollowUp = followUp
	return b
}


// WithValidation adds a custom validation function
func (b *MedicalRecordBuilder) WithValidation(validationFunc func(*models.MedicalRecord) error) *MedicalRecordBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the MedicalRecord
func (b *MedicalRecordBuilder) Build() interface{} {
	return b.medicalRecord
}

// BuildPtr builds the MedicalRecord and returns a pointer
func (b *MedicalRecordBuilder) BuildPtr() *models.MedicalRecord {
	return b.medicalRecord
}

// BuildAndValidate builds the MedicalRecord and validates it
func (b *MedicalRecordBuilder) BuildAndValidate() (*models.MedicalRecord, error) {
	medicalRecord := b.medicalRecord

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(medicalRecord); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(medicalRecord).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return medicalRecord, err
		}
	}

	return medicalRecord, nil
}

// MustBuild builds the MedicalRecord and panics if validation fails
func (b *MedicalRecordBuilder) MustBuild() *models.MedicalRecord {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *MedicalRecordBuilder) Clone() *MedicalRecordBuilder {
	clonedMedicalRecord := *b.medicalRecord
	return &MedicalRecordBuilder{
		medicalRecord: &clonedMedicalRecord,
		validationFuncs: append([]func(*models.MedicalRecord) error{}, b.validationFuncs...),
	}
}
