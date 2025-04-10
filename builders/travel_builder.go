package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

// TravelBuilder builds a Travel model
type TravelBuilder struct {
	travel         *models.Travel
	validationFuncs []func(*models.Travel) error
}

// NewTravelBuilder creates a new TravelBuilder
func NewTravelBuilder() *TravelBuilder {
	return &TravelBuilder{
		travel: &models.Travel{
			Destination:    models.Address{},
			StartDate:      "",
			EndDate:        "",
			Purpose:        "",
			Accommodation:  "",
			Transportation: "",
			Activities:     []string{},
			Expenses:       0.0,
		},
		validationFuncs: []func(*models.Travel) error{},
	}
}

// WithDestination sets the Destination
func (b *TravelBuilder) WithDestination(destination *AddressBuilder) *TravelBuilder {
	address := destination.Build().(models.Address)
	b.travel.Destination = address
	return b
}

// WithStartDate sets the StartDate
func (b *TravelBuilder) WithStartDate(startDate string) *TravelBuilder {
	b.travel.StartDate = startDate
	return b
}

// WithEndDate sets the EndDate
func (b *TravelBuilder) WithEndDate(endDate string) *TravelBuilder {
	b.travel.EndDate = endDate
	return b
}

// WithPurpose sets the Purpose
func (b *TravelBuilder) WithPurpose(purpose string) *TravelBuilder {
	b.travel.Purpose = purpose
	return b
}

// WithAccommodation sets the Accommodation
func (b *TravelBuilder) WithAccommodation(accommodation string) *TravelBuilder {
	b.travel.Accommodation = accommodation
	return b
}

// WithTransportation sets the Transportation
func (b *TravelBuilder) WithTransportation(transportation string) *TravelBuilder {
	b.travel.Transportation = transportation
	return b
}

// WithActivities sets the Activities
func (b *TravelBuilder) WithActivities(activities []string) *TravelBuilder {
	b.travel.Activities = activities
	return b
}

// AddActivity adds an activity to the Activities slice
func (b *TravelBuilder) AddActivity(activity string) *TravelBuilder {
	b.travel.Activities = append(b.travel.Activities, activity)
	return b
}

// WithExpenses sets the Expenses
func (b *TravelBuilder) WithExpenses(expenses float64) *TravelBuilder {
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
	travel := *b.travel
	return &travel
}

// BuildAndValidate builds the Travel and validates it
func (b *TravelBuilder) BuildAndValidate() (*models.Travel, error) {
	travel := b.travel

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(travel); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method
	if err := travel.Validate(); err != nil {
		return travel, err
	}

	return travel, nil
}

// MustBuild builds the Travel and panics if validation fails
func (b *TravelBuilder) MustBuild() *models.Travel {
	travel, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return travel
}

// Clone creates a deep copy of the builder
func (b *TravelBuilder) Clone() *TravelBuilder {
	clonedTravel := *b.travel
	return &TravelBuilder{
		travel:         &clonedTravel,
		validationFuncs: append([]func(*models.Travel) error{}, b.validationFuncs...),
	}
}
