package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// FamilyMemberBuilder builds a FamilyMember model
type FamilyMemberBuilder struct {
	familyMember *models.FamilyMember
	// Custom validation functions
	validationFuncs []func(*models.FamilyMember) error
}

// NewFamilyMemberBuilder creates a new FamilyMemberBuilder
func NewFamilyMemberBuilder() *FamilyMemberBuilder {
	return &FamilyMemberBuilder{
		familyMember: &models.FamilyMember{
			Person: nil,
			Relationship: "",
		},
		validationFuncs: []func(*models.FamilyMember) error{},
	}
}

// NewFamilyMemberBuilderWithDefaults creates a new FamilyMemberBuilder with sensible defaults
func NewFamilyMemberBuilderWithDefaults() *FamilyMemberBuilder {
	builder := NewFamilyMemberBuilder()
	// Add default values here if needed
	return builder
}
// WithPerson sets the Person
func (b *FamilyMemberBuilder) WithPerson(person *PersonBuilder) *FamilyMemberBuilder {
	// Handle nested pointer
	b.familyMember.Person = person.BuildPtr()
	return b
}

// WithRelationship sets the Relationship
func (b *FamilyMemberBuilder) WithRelationship(relationship string) *FamilyMemberBuilder {
	b.familyMember.Relationship = relationship
	return b
}


// WithValidation adds a custom validation function
func (b *FamilyMemberBuilder) WithValidation(validationFunc func(*models.FamilyMember) error) *FamilyMemberBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the FamilyMember
func (b *FamilyMemberBuilder) Build() interface{} {
	return b.familyMember
}

// BuildPtr builds the FamilyMember and returns a pointer
func (b *FamilyMemberBuilder) BuildPtr() *models.FamilyMember {
	return b.familyMember
}

// BuildAndValidate builds the FamilyMember and validates it
func (b *FamilyMemberBuilder) BuildAndValidate() (*models.FamilyMember, error) {
	familyMember := b.familyMember

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(familyMember); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(familyMember).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return familyMember, err
		}
	}

	return familyMember, nil
}

// MustBuild builds the FamilyMember and panics if validation fails
func (b *FamilyMemberBuilder) MustBuild() *models.FamilyMember {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *FamilyMemberBuilder) Clone() *FamilyMemberBuilder {
	clonedFamilyMember := *b.familyMember
	return &FamilyMemberBuilder{
		familyMember: &clonedFamilyMember,
		validationFuncs: append([]func(*models.FamilyMember) error{}, b.validationFuncs...),
	}
}
