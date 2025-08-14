package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/v2/models"
	"github.com/adil-faiyaz98/go-builder-kit/v2/pkg/builder"
)

// AddressBuilder builds a Address model
type AddressBuilder struct {
	address *models.Address
	// Custom validation functions
	validationFuncs []func(*models.Address) error
}

// NewAddressBuilder creates a new AddressBuilder
func NewAddressBuilder() *AddressBuilder {
	return &AddressBuilder{
		address: &models.Address{
			Street:      "",
			City:        "",
			State:       "",
			PostalCode:  "",
			Country:     "",
			Coordinates: nil,
			Type:        "",
			IsPrimary:   false,
		},
		validationFuncs: []func(*models.Address) error{},
	}
}

// NewAddressBuilderWithDefaults creates a new AddressBuilder with sensible defaults
func NewAddressBuilderWithDefaults() *AddressBuilder {
	builder := NewAddressBuilder()
	// Add default values here if needed
	return builder
}

// WithStreet sets the Street
func (b *AddressBuilder) WithStreet(street string) *AddressBuilder {
	if b == nil {
		return b
	}
	b.address.Street = builder.SanitizeString(street)
	return b
}

// WithCity sets the City
func (b *AddressBuilder) WithCity(city string) *AddressBuilder {
	if b == nil {
		return b
	}
	b.address.City = builder.SanitizeString(city)
	return b
}

// WithState sets the State
func (b *AddressBuilder) WithState(state string) *AddressBuilder {
	if b == nil {
		return b
	}
	b.address.State = builder.SanitizeString(state)
	return b
}

// WithPostalCode sets the PostalCode
func (b *AddressBuilder) WithPostalCode(postalCode string) *AddressBuilder {
	if b == nil {
		return b
	}
	b.address.PostalCode = builder.SanitizeString(postalCode)
	return b
}

// WithCountry sets the Country
func (b *AddressBuilder) WithCountry(country string) *AddressBuilder {
	if b == nil {
		return b
	}
	b.address.Country = builder.SanitizeString(country)
	return b
}

// WithCoordinates sets the Coordinates
func (b *AddressBuilder) WithCoordinates(coordinates *GeoLocationBuilder) *AddressBuilder {
	if b == nil {
		return b
	}
	// Handle nested pointer
	b.address.Coordinates = coordinates.BuildPtr()
	return b
}

// WithType sets the Type
func (b *AddressBuilder) WithType(value string) *AddressBuilder {
	if b == nil {
		return b
	}
	b.address.Type = builder.SanitizeString(value)
	return b
}

// WithIsPrimary sets the IsPrimary
func (b *AddressBuilder) WithIsPrimary(isPrimary bool) *AddressBuilder {
	if b == nil {
		return b
	}
	b.address.IsPrimary = isPrimary
	return b
}

// WithValidation adds a custom validation function
func (b *AddressBuilder) WithValidation(validationFunc func(*models.Address) error) *AddressBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Address
func (b *AddressBuilder) Build() interface{} {
	return b.address
}

// BuildPtr builds the Address and returns a pointer
func (b *AddressBuilder) BuildPtr() *models.Address {
	return b.address
}

// BuildAndValidate builds the Address and validates it
func (b *AddressBuilder) BuildAndValidate() (*models.Address, error) {
	if b == nil || b.address == nil {
		return nil, fmt.Errorf("builder is not properly initialized")
	}

	address := b.address

	// Run custom validation functions
	for i, validationFunc := range b.validationFuncs {
		if validationFunc == nil {
			continue // Skip nil validators
		}
		if err := validationFunc(address); err != nil {
			return nil, fmt.Errorf("custom validation failed at index %d: %w", i, err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := any(address).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return address, fmt.Errorf("model validation failed: %w", err)
		}
	}

	return address, nil
}

// MustBuild builds the Address and panics if validation fails
func (b *AddressBuilder) MustBuild() *models.Address {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *AddressBuilder) Clone() *AddressBuilder {
	if b == nil || b.address == nil {
		return NewAddressBuilder()
	}

	// Deep copy the struct
	clonedAddress := *b.address

	// Create new builder with cloned data
	clonedBuilder := &AddressBuilder{
		address:         &clonedAddress,
		validationFuncs: make([]func(*models.Address) error, 0, len(b.validationFuncs)),
	}

	// Copy validation functions safely
	if len(b.validationFuncs) > 0 {
		clonedBuilder.validationFuncs = append(clonedBuilder.validationFuncs, b.validationFuncs...)
	}

	return clonedBuilder
}
