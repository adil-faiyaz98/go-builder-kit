package builders

import (
	"fmt"

	"github.com/adil-faiyaz98/go-builder-kit/models"
)

// FamilyMemberBuilder builds a FamilyMember model
type FamilyMemberBuilder struct {
	familyMember    *models.FamilyMember
	validationFuncs []func(*models.FamilyMember) error
}

// NewFamilyMemberBuilder creates a new FamilyMemberBuilder
func NewFamilyMemberBuilder() *FamilyMemberBuilder {
	return &FamilyMemberBuilder{
		familyMember: &models.FamilyMember{
			Relationship: "",
			Person:       nil,
		},
		validationFuncs: []func(*models.FamilyMember) error{},
	}
}

// WithRelationship sets the Relationship
func (b *FamilyMemberBuilder) WithRelationship(relationship string) *FamilyMemberBuilder {
	b.familyMember.Relationship = relationship
	return b
}

// WithPerson sets the Person
func (b *FamilyMemberBuilder) WithPerson(person *PersonBuilder) *FamilyMemberBuilder {
	builtValue := person.Build().(*models.Person)
	b.familyMember.Person = builtValue
	return b
}

// WithValidation adds a custom validation function
func (b *FamilyMemberBuilder) WithValidation(validationFunc func(*models.FamilyMember) error) *FamilyMemberBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the FamilyMember
func (b *FamilyMemberBuilder) Build() any {
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
			return familyMember, err
		}
	}

	// Run model's Validate method
	if err := familyMember.Validate(); err != nil {
		return familyMember, err
	}

	return familyMember, nil
}

// MustBuild builds the FamilyMember and panics if validation fails
func (b *FamilyMemberBuilder) MustBuild() *models.FamilyMember {
	familyMember, err := b.BuildAndValidate()
	if err != nil {
		panic(fmt.Sprintf("FamilyMember validation failed: %s", err.Error()))
	}
	return familyMember
}

// Clone creates a deep copy of the FamilyMemberBuilder
func (b *FamilyMemberBuilder) Clone() *FamilyMemberBuilder {
	clonedFamilyMember := *b.familyMember
	return &FamilyMemberBuilder{
		familyMember:    &clonedFamilyMember,
		validationFuncs: append([]func(*models.FamilyMember) error{}, b.validationFuncs...),
	}
}
