package builders

import (
	"fmt"

	"github.com/adil-faiyaz98/go-builder-kit/models"
)

// GeoLocationBuilder builds a GeoLocation model
type GeoLocationBuilder struct {
	geoLocation *models.GeoLocation
	// Custom validation functions
	validationFuncs []func(*models.GeoLocation) error
}

// NewGeoLocationBuilder creates a new GeoLocationBuilder
func NewGeoLocationBuilder() *GeoLocationBuilder {
	return &GeoLocationBuilder{
		geoLocation: &models.GeoLocation{
			Latitude:  0.0,
			Longitude: 0.0,
			Accuracy:  0.0,
		},
		validationFuncs: []func(*models.GeoLocation) error{},
	}
}

// NewGeoLocationBuilderWithDefaults creates a new GeoLocationBuilder with sensible defaults
func NewGeoLocationBuilderWithDefaults() *GeoLocationBuilder {
	builder := NewGeoLocationBuilder()
	// Add default values here if needed
	return builder
}

// WithLatitude sets the Latitude
func (b *GeoLocationBuilder) WithLatitude(latitude float64) *GeoLocationBuilder {
	if b == nil {
		return b
	}
	b.geoLocation.Latitude = latitude
	return b
}

// WithLongitude sets the Longitude
func (b *GeoLocationBuilder) WithLongitude(longitude float64) *GeoLocationBuilder {
	if b == nil {
		return b
	}
	b.geoLocation.Longitude = longitude
	return b
}

// WithAccuracy sets the Accuracy
func (b *GeoLocationBuilder) WithAccuracy(accuracy float64) *GeoLocationBuilder {
	if b == nil {
		return b
	}
	b.geoLocation.Accuracy = accuracy
	return b
}

// WithValidation adds a custom validation function
func (b *GeoLocationBuilder) WithValidation(validationFunc func(*models.GeoLocation) error) *GeoLocationBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the GeoLocation
func (b *GeoLocationBuilder) Build() interface{} {
	return b.geoLocation
}

// BuildPtr builds the GeoLocation and returns a pointer
func (b *GeoLocationBuilder) BuildPtr() *models.GeoLocation {
	return b.geoLocation
}

// BuildAndValidate builds the GeoLocation and validates it
func (b *GeoLocationBuilder) BuildAndValidate() (*models.GeoLocation, error) {
	if b == nil || b.geoLocation == nil {
		return nil, fmt.Errorf("builder is not properly initialized")
	}

	geoLocation := b.geoLocation

	// Run custom validation functions
	for i, validationFunc := range b.validationFuncs {
		if validationFunc == nil {
			continue // Skip nil validators
		}
		if err := validationFunc(geoLocation); err != nil {
			return nil, fmt.Errorf("custom validation failed at index %d: %w", i, err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := any(geoLocation).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return geoLocation, fmt.Errorf("model validation failed: %w", err)
		}
	}

	return geoLocation, nil
}

// MustBuild builds the GeoLocation and panics if validation fails
func (b *GeoLocationBuilder) MustBuild() *models.GeoLocation {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *GeoLocationBuilder) Clone() *GeoLocationBuilder {
	if b == nil || b.geoLocation == nil {
		return NewGeoLocationBuilder()
	}

	// Deep copy the struct
	clonedGeoLocation := *b.geoLocation

	// Create new builder with cloned data
	clonedBuilder := &GeoLocationBuilder{
		geoLocation:     &clonedGeoLocation,
		validationFuncs: make([]func(*models.GeoLocation) error, 0, len(b.validationFuncs)),
	}

	// Copy validation functions safely
	if len(b.validationFuncs) > 0 {
		clonedBuilder.validationFuncs = append(clonedBuilder.validationFuncs, b.validationFuncs...)
	}

	return clonedBuilder
}
