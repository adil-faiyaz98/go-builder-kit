package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

// AddressBuilder builds an Address model
type AddressBuilder struct {
	address         *models.Address
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

// WithStreet sets the Street
func (b *AddressBuilder) WithStreet(street string) *AddressBuilder {
	b.address.Street = street
	return b
}

// WithCity sets the City
func (b *AddressBuilder) WithCity(city string) *AddressBuilder {
	b.address.City = city
	return b
}

// WithState sets the State
func (b *AddressBuilder) WithState(state string) *AddressBuilder {
	b.address.State = state
	return b
}

// WithPostalCode sets the PostalCode
func (b *AddressBuilder) WithPostalCode(postalCode string) *AddressBuilder {
	b.address.PostalCode = postalCode
	return b
}

// WithCountry sets the Country
func (b *AddressBuilder) WithCountry(country string) *AddressBuilder {
	b.address.Country = country
	return b
}

// WithCoordinates sets the Coordinates
func (b *AddressBuilder) WithCoordinates(coordinates *GeoLocationBuilder) *AddressBuilder {
	builtValue := coordinates.Build().(*models.GeoLocation)
	b.address.Coordinates = builtValue
	return b
}

// WithType sets the Type
func (b *AddressBuilder) WithType(addressType string) *AddressBuilder {
	b.address.Type = addressType
	return b
}

// WithIsPrimary sets the IsPrimary
func (b *AddressBuilder) WithIsPrimary(isPrimary bool) *AddressBuilder {
	b.address.IsPrimary = isPrimary
	return b
}

// WithValidation adds a custom validation function
func (b *AddressBuilder) WithValidation(validationFunc func(*models.Address) error) *AddressBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Address
func (b *AddressBuilder) Build() any {
	return b.address
}

// BuildPtr builds the Address and returns a pointer
func (b *AddressBuilder) BuildPtr() *models.Address {
	return b.address
}

// BuildAndValidate builds the Address and validates it
func (b *AddressBuilder) BuildAndValidate() (*models.Address, error) {
	address := b.address

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(address); err != nil {
			return address, err
		}
	}

	// Run model's Validate method
	if err := address.Validate(); err != nil {
		return address, err
	}

	return address, nil
}

// MustBuild builds the Address and panics if validation fails
func (b *AddressBuilder) MustBuild() *models.Address {
	address, err := b.BuildAndValidate()
	if err != nil {
		panic(fmt.Sprintf("Address validation failed: %s", err.Error()))
	}
	return address
}

// Clone creates a deep copy of the AddressBuilder
func (b *AddressBuilder) Clone() *AddressBuilder {
	clonedAddress := *b.address
	return &AddressBuilder{
		address:         &clonedAddress,
		validationFuncs: append([]func(*models.Address) error{}, b.validationFuncs...),
	}
}
