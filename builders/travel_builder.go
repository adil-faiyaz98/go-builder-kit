package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/v2/models"
	"github.com/adil-faiyaz98/go-builder-kit/v2/pkg/builder"
)

// TravelBuilder builds a Travel model
type TravelBuilder struct {
	travel *models.Travel
	// Custom validation functions
	validationFuncs []func(*models.Travel) error
}

// NewTravelBuilder creates a new TravelBuilder
func NewTravelBuilder() *TravelBuilder {
	return &TravelBuilder{
		travel: &models.Travel{
			Destination: models.Address{},
			StartDate: "",
			EndDate: "",
			Purpose: "",
			Accommodation: "",
			Transportation: "",
			Activities: []string{},
			Expenses: 0.0,
		},
		validationFuncs: []func(*models.Travel) error{},
	}
}

// NewTravelBuilderWithDefaults creates a new TravelBuilder with sensible defaults
func NewTravelBuilderWithDefaults() *TravelBuilder {
	builder := NewTravelBuilder()
	// Add default values here if needed
	return builder
}
// WithDestination sets the Destination
func (b *TravelBuilder) WithDestination(destination *AddressBuilder) *TravelBuilder {
	if b == nil {
		return b
	}
	builtValue := destination.Build().(*models.Address)
	b.travel.Destination = *builtValue
	return b
}

// WithStartDate sets the StartDate
func (b *TravelBuilder) WithStartDate(startDate string) *TravelBuilder {
	if b == nil {
		return b
	}
	b.travel.StartDate = builder.SanitizeString(startDate)
	return b
}

// WithEndDate sets the EndDate
func (b *TravelBuilder) WithEndDate(endDate string) *TravelBuilder {
	if b == nil {
		return b
	}
	b.travel.EndDate = builder.SanitizeString(endDate)
	return b
}

// WithPurpose sets the Purpose
func (b *TravelBuilder) WithPurpose(purpose string) *TravelBuilder {
	if b == nil {
		return b
	}
	b.travel.Purpose = builder.SanitizeString(purpose)
	return b
}

// WithAccommodation sets the Accommodation
func (b *TravelBuilder) WithAccommodation(accommodation string) *TravelBuilder {
	if b == nil {
		return b
	}
	b.travel.Accommodation = builder.SanitizeString(accommodation)
	return b
}

// WithTransportation sets the Transportation
func (b *TravelBuilder) WithTransportation(transportation string) *TravelBuilder {
	if b == nil {
		return b
	}
	b.travel.Transportation = builder.SanitizeString(transportation)
	return b
}

// WithActivities sets the Activities
func (b *TravelBuilder) WithActivities(activities []string) *TravelBuilder {
	if b == nil {
		return b
	}
	b.travel.Activities = append(b.travel.Activities, activities...)
	return b
}

// WithExpenses sets the Expenses
func (b *TravelBuilder) WithExpenses(expenses float64) *TravelBuilder {
	if b == nil {
		return b
	}
	b.travel.Expenses = expenses
	return b
}


// WithValidation adds a custom validation function
func (b *TravelBuilder) WithValidation(validationFunc func(*models.Travel) error) *TravelBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Travel
func (b *TravelBuilder) Build() interface{} {
	return b.travel
}

// BuildPtr builds the Travel and returns a pointer
func (b *TravelBuilder) BuildPtr() *models.Travel {
	return b.travel
}

// BuildAndValidate builds the Travel and validates it
func (b *TravelBuilder) BuildAndValidate() (*models.Travel, error) {
	if b == nil || b.travel == nil {
		return nil, fmt.Errorf("builder is not properly initialized")
	}

	travel := b.travel

	// Run custom validation functions
	for i, validationFunc := range b.validationFuncs {
		if validationFunc == nil {
			continue // Skip nil validators
		}
		if err := validationFunc(travel); err != nil {
			return nil, fmt.Errorf("custom validation failed at index %d: %w", i, err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := any(travel).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return travel, fmt.Errorf("model validation failed: %w", err)
		}
	}

	return travel, nil
}

// MustBuild builds the Travel and panics if validation fails
func (b *TravelBuilder) MustBuild() *models.Travel {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *TravelBuilder) Clone() *TravelBuilder {
	if b == nil || b.travel == nil {
		return NewTravelBuilder()
	}

	// Deep copy the struct
	clonedTravel := *b.travel

	// Create new builder with cloned data
	clonedBuilder := &TravelBuilder{
		travel: &clonedTravel,
		validationFuncs: make([]func(*models.Travel) error, 0, len(b.validationFuncs)),
	}

	// Copy validation functions safely
	if len(b.validationFuncs) > 0 {
		clonedBuilder.validationFuncs = append(clonedBuilder.validationFuncs, b.validationFuncs...)
	}

	return clonedBuilder
}
