package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/v2/models"
	"github.com/adil-faiyaz98/go-builder-kit/v2/pkg/builder"
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
	if b == nil {
		return b
	}
	b.performanceRecord.Date = builder.SanitizeString(date)
	return b
}

// WithValue sets the Value
func (b *PerformanceRecordBuilder) WithValue(value float64) *PerformanceRecordBuilder {
	if b == nil {
		return b
	}
	b.performanceRecord.Value = value
	return b
}

// WithNotes sets the Notes
func (b *PerformanceRecordBuilder) WithNotes(notes string) *PerformanceRecordBuilder {
	if b == nil {
		return b
	}
	b.performanceRecord.Notes = builder.SanitizeString(notes)
	return b
}

// WithROI sets the ROI
func (b *PerformanceRecordBuilder) WithROI(rOI float64) *PerformanceRecordBuilder {
	if b == nil {
		return b
	}
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
	if b == nil || b.performanceRecord == nil {
		return nil, fmt.Errorf("builder is not properly initialized")
	}

	performanceRecord := b.performanceRecord

	// Run custom validation functions
	for i, validationFunc := range b.validationFuncs {
		if validationFunc == nil {
			continue // Skip nil validators
		}
		if err := validationFunc(performanceRecord); err != nil {
			return nil, fmt.Errorf("custom validation failed at index %d: %w", i, err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := any(performanceRecord).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return performanceRecord, fmt.Errorf("model validation failed: %w", err)
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
	if b == nil || b.performanceRecord == nil {
		return NewPerformanceRecordBuilder()
	}

	// Deep copy the struct
	clonedPerformanceRecord := *b.performanceRecord

	// Create new builder with cloned data
	clonedBuilder := &PerformanceRecordBuilder{
		performanceRecord: &clonedPerformanceRecord,
		validationFuncs: make([]func(*models.PerformanceRecord) error, 0, len(b.validationFuncs)),
	}

	// Copy validation functions safely
	if len(b.validationFuncs) > 0 {
		clonedBuilder.validationFuncs = append(clonedBuilder.validationFuncs, b.validationFuncs...)
	}

	return clonedBuilder
}
