package simple

import (
	"fmt"
)

// PersonBuilder builds a Person model
type PersonBuilder struct {
	person          *Person
	validationFuncs []func(*Person) error
}

// NewPersonBuilder creates a new PersonBuilder
func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{
		person: &Person{
			ID:       "",
			Name:     "",
			Age:      0,
			Email:    "",
			Address:  nil,
			Tags:     []string{},
			Metadata: map[string]string{},
		},
		validationFuncs: []func(*Person) error{},
	}
}

// WithID sets the ID
func (b *PersonBuilder) WithID(id string) *PersonBuilder {
	b.person.ID = id
	return b
}

// WithName sets the Name
func (b *PersonBuilder) WithName(name string) *PersonBuilder {
	b.person.Name = name
	return b
}

// WithAge sets the Age
func (b *PersonBuilder) WithAge(age int) *PersonBuilder {
	b.person.Age = age
	return b
}

// WithEmail sets the Email
func (b *PersonBuilder) WithEmail(email string) *PersonBuilder {
	b.person.Email = email
	return b
}

// WithAddress sets the Address
func (b *PersonBuilder) WithAddress(address *AddressBuilder) *PersonBuilder {
	builtValue := address.Build()
	b.person.Address = &builtValue
	return b
}

// WithTag adds a tag to the Tags slice
func (b *PersonBuilder) WithTag(tag string) *PersonBuilder {
	b.person.Tags = append(b.person.Tags, tag)
	return b
}

// WithMetadata adds a key-value pair to the Metadata map
func (b *PersonBuilder) WithMetadata(key, value string) *PersonBuilder {
	if b.person.Metadata == nil {
		b.person.Metadata = map[string]string{}
	}
	b.person.Metadata[key] = value
	return b
}

// WithValidation adds a custom validation function
func (b *PersonBuilder) WithValidation(validationFunc func(*Person) error) *PersonBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Person
func (b *PersonBuilder) Build() Person {
	return *b.person
}

// BuildPtr builds the Person and returns a pointer
func (b *PersonBuilder) BuildPtr() *Person {
	return b.person
}

// BuildAndValidate builds the Person and validates it
func (b *PersonBuilder) BuildAndValidate() (*Person, error) {
	person := b.person

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(person); err != nil {
			return person, err
		}
	}

	// Run model's Validate method
	if err := person.Validate(); err != nil {
		return person, err
	}

	return person, nil
}

// MustBuild builds the Person and panics if validation fails
func (b *PersonBuilder) MustBuild() *Person {
	person, err := b.BuildAndValidate()
	if err != nil {
		panic(fmt.Sprintf("Person validation failed: %s", err.Error()))
	}
	return person
}

// Clone creates a deep copy of the PersonBuilder
func (b *PersonBuilder) Clone() *PersonBuilder {
	clonedPerson := *b.person
	
	// Deep copy slices and maps
	if b.person.Tags != nil {
		clonedPerson.Tags = make([]string, len(b.person.Tags))
		copy(clonedPerson.Tags, b.person.Tags)
	}
	
	if b.person.Metadata != nil {
		clonedPerson.Metadata = make(map[string]string)
		for k, v := range b.person.Metadata {
			clonedPerson.Metadata[k] = v
		}
	}
	
	// Deep copy Address if it exists
	if b.person.Address != nil {
		addressCopy := *b.person.Address
		clonedPerson.Address = &addressCopy
	}
	
	return &PersonBuilder{
		person:          &clonedPerson,
		validationFuncs: append([]func(*Person) error{}, b.validationFuncs...),
	}
}

// AddressBuilder builds an Address model
type AddressBuilder struct {
	address         *Address
	validationFuncs []func(*Address) error
}

// NewAddressBuilder creates a new AddressBuilder
func NewAddressBuilder() *AddressBuilder {
	return &AddressBuilder{
		address: &Address{
			Street:     "",
			City:       "",
			State:      "",
			PostalCode: "",
			Country:    "",
		},
		validationFuncs: []func(*Address) error{},
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

// WithValidation adds a custom validation function
func (b *AddressBuilder) WithValidation(validationFunc func(*Address) error) *AddressBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Address
func (b *AddressBuilder) Build() Address {
	return *b.address
}

// BuildPtr builds the Address and returns a pointer
func (b *AddressBuilder) BuildPtr() *Address {
	return b.address
}

// BuildAndValidate builds the Address and validates it
func (b *AddressBuilder) BuildAndValidate() (*Address, error) {
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
func (b *AddressBuilder) MustBuild() *Address {
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
		validationFuncs: append([]func(*Address) error{}, b.validationFuncs...),
	}
}
