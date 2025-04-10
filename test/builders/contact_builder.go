package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/test/models"
)

// ContactBuilder builds a Contact model
type ContactBuilder struct {
	contact *models.Contact
	// Custom validation functions
	validationFuncs []func(*models.Contact) error
}

// NewContactBuilder creates a new ContactBuilder
func NewContactBuilder() *ContactBuilder {
	return &ContactBuilder{
		contact: &models.Contact{
			Email: "",
			Phone: "",
			Address: nil,
			Alternative: nil,
		},
		validationFuncs: []func(*models.Contact) error{},
	}
}

// NewContactBuilderWithDefaults creates a new ContactBuilder with sensible defaults
func NewContactBuilderWithDefaults() *ContactBuilder {
	builder := NewContactBuilder()
	// Add default values here if needed
	return builder
}
// WithEmail sets the Email
func (b *ContactBuilder) WithEmail(email string) *ContactBuilder {
	b.contact.Email = email
	return b
}

// WithPhone sets the Phone
func (b *ContactBuilder) WithPhone(phone string) *ContactBuilder {
	b.contact.Phone = phone
	return b
}

// WithAddress sets the Address
func (b *ContactBuilder) WithAddress(address *AddressBuilder) *ContactBuilder {
	// Handle nested pointer
	b.contact.Address = address.BuildPtr()
	return b
}

// WithAlternative sets the Alternative
func (b *ContactBuilder) WithAlternative(alternative *ContactBuilder) *ContactBuilder {
	// Handle nested pointer
	b.contact.Alternative = alternative.BuildPtr()
	return b
}


// WithValidation adds a custom validation function
func (b *ContactBuilder) WithValidation(validationFunc func(*models.Contact) error) *ContactBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Contact
func (b *ContactBuilder) Build() interface{} {
	return b.contact
}

// BuildPtr builds the Contact and returns a pointer
func (b *ContactBuilder) BuildPtr() *models.Contact {
	return b.contact
}

// BuildAndValidate builds the Contact and validates it
func (b *ContactBuilder) BuildAndValidate() (*models.Contact, error) {
	contact := b.contact

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(contact); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(contact).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return contact, err
		}
	}

	return contact, nil
}

// MustBuild builds the Contact and panics if validation fails
func (b *ContactBuilder) MustBuild() *models.Contact {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *ContactBuilder) Clone() *ContactBuilder {
	clonedContact := *b.contact
	return &ContactBuilder{
		contact: &clonedContact,
		validationFuncs: append([]func(*models.Contact) error{}, b.validationFuncs...),
	}
}
