package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// PerformanceRecordBuilder builds a PerformanceRecord model
type PerformanceRecordBuilder struct {
	performanceRecord *models.PerformanceRecord
	// Custom validation functions
	validationFuncs []func(*models.PerformanceRecord) error
}

// NewPerformanceRecordBuilder creates a new PerformanceRecordBuilder
func NewPerformanceRecordBuilder() *PerformanceRecordBuilder {
	return &PerformanceRecordBuilder{
		performanceRecord: &models.PerformanceRecord{
			Date: "",
			Value: 0.0,
			Notes: "",
			ROI: 0.0,
		},
		validationFuncs: []func(*models.PerformanceRecord) error{},
	}
}

// NewPerformanceRecordBuilderWithDefaults creates a new PerformanceRecordBuilder with sensible defaults
func NewPerformanceRecordBuilderWithDefaults() *PerformanceRecordBuilder {
	builder := NewPerformanceRecordBuilder()
	// Add default values here if needed
	return builder
}
// WithDate sets the Date
func (b *PerformanceRecordBuilder) WithDate(date string) *PerformanceRecordBuilder {
	b.performanceRecord.Date = date
	return b
}

// WithValue sets the Value
func (b *PerformanceRecordBuilder) WithValue(value float64) *PerformanceRecordBuilder {
	b.performanceRecord.Value = value
	return b
}

// WithNotes sets the Notes
func (b *PerformanceRecordBuilder) WithNotes(notes string) *PerformanceRecordBuilder {
	b.performanceRecord.Notes = notes
	return b
}

// WithROI sets the ROI
func (b *PerformanceRecordBuilder) WithROI(rOI float64) *PerformanceRecordBuilder {
	b.performanceRecord.ROI = rOI
	return b
}


// WithValidation adds a custom validation function
func (b *PerformanceRecordBuilder) WithValidation(validationFunc func(*models.PerformanceRecord) error) *PerformanceRecordBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the PerformanceRecord
func (b *PerformanceRecordBuilder) Build() interface{} {
	return b.performanceRecord
}

// BuildPtr builds the PerformanceRecord and returns a pointer
func (b *PerformanceRecordBuilder) BuildPtr() *models.PerformanceRecord {
	return b.performanceRecord
}

// BuildAndValidate builds the PerformanceRecord and validates it
func (b *PerformanceRecordBuilder) BuildAndValidate() (*models.PerformanceRecord, error) {
	performanceRecord := b.performanceRecord

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(performanceRecord); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(performanceRecord).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return performanceRecord, err
		}
	}

	return performanceRecord, nil
}

// MustBuild builds the PerformanceRecord and panics if validation fails
func (b *PerformanceRecordBuilder) MustBuild() *models.PerformanceRecord {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *PerformanceRecordBuilder) Clone() *PerformanceRecordBuilder {
	clonedPerformanceRecord := *b.performanceRecord
	return &PerformanceRecordBuilder{
		performanceRecord: &clonedPerformanceRecord,
		validationFuncs: append([]func(*models.PerformanceRecord) error{}, b.validationFuncs...),
	}
}
