package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// PersonBuilder builds a Person model
type PersonBuilder struct {
	person *models.Person
	// Custom validation functions
	validationFuncs []func(*models.Person) error
}

// NewPersonBuilder creates a new PersonBuilder
func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{
		person: &models.Person{
			ID: "",
			Name: "",
			Age: 0,
			Email: "",
			Phone: "",
			Birthdate: "",
			Gender: "",
			Nationality: "",
			MaritalStatus: "",
			Address: nil,
			Education: nil,
			Profile: nil,
			Bank: nil,
			Employment: nil,
			Friends: []*models.Person{},
			Family: []*models.FamilyMember{},
			Health: nil,
			Digital: nil,
			TravelHistory: []*models.Travel{},
			Preferences: nil,
			CreatedAt: "",
			UpdatedAt: "",
		},
		validationFuncs: []func(*models.Person) error{},
	}
}

// NewPersonBuilderWithDefaults creates a new PersonBuilder with sensible defaults
func NewPersonBuilderWithDefaults() *PersonBuilder {
	builder := NewPersonBuilder()
	// Add default values here if needed
	return builder
}
// WithID sets the ID
func (b *PersonBuilder) WithID(iD string) *PersonBuilder {
	b.person.ID = iD
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

// WithPhone sets the Phone
func (b *PersonBuilder) WithPhone(phone string) *PersonBuilder {
	b.person.Phone = phone
	return b
}

// WithBirthdate sets the Birthdate
func (b *PersonBuilder) WithBirthdate(birthdate string) *PersonBuilder {
	b.person.Birthdate = birthdate
	return b
}

// WithGender sets the Gender
func (b *PersonBuilder) WithGender(gender string) *PersonBuilder {
	b.person.Gender = gender
	return b
}

// WithNationality sets the Nationality
func (b *PersonBuilder) WithNationality(nationality string) *PersonBuilder {
	b.person.Nationality = nationality
	return b
}

// WithMaritalStatus sets the MaritalStatus
func (b *PersonBuilder) WithMaritalStatus(maritalStatus string) *PersonBuilder {
	b.person.MaritalStatus = maritalStatus
	return b
}

// WithAddress sets the Address
func (b *PersonBuilder) WithAddress(address *AddressBuilder) *PersonBuilder {
	// Handle nested pointer
	b.person.Address = address.BuildPtr()
	return b
}

// WithEducation sets the Education
func (b *PersonBuilder) WithEducation(education *EducationBuilder) *PersonBuilder {
	// Handle nested pointer
	b.person.Education = education.BuildPtr()
	return b
}

// WithProfile sets the Profile
func (b *PersonBuilder) WithProfile(profile *ProfileBuilder) *PersonBuilder {
	// Handle nested pointer
	b.person.Profile = profile.BuildPtr()
	return b
}

// WithBank sets the Bank
func (b *PersonBuilder) WithBank(bank *BankBuilder) *PersonBuilder {
	// Handle nested pointer
	b.person.Bank = bank.BuildPtr()
	return b
}

// WithEmployment sets the Employment
func (b *PersonBuilder) WithEmployment(employment *EmploymentBuilder) *PersonBuilder {
	// Handle nested pointer
	b.person.Employment = employment.BuildPtr()
	return b
}

// WithFriends sets the Friends
func (b *PersonBuilder) WithFriends(friends *PersonBuilder) *PersonBuilder {
	// Ensure the slice is initialized
	if b.person.Friends == nil {
		b.person.Friends = []*models.Person{}
	}
	// Handle nested slice element
	builtValue := friends.Build().(*models.Person)
	b.person.Friends = append(b.person.Friends, builtValue)
	return b
}

// WithFamily sets the Family
func (b *PersonBuilder) WithFamily(family *FamilyMemberBuilder) *PersonBuilder {
	// Ensure the slice is initialized
	if b.person.Family == nil {
		b.person.Family = []*models.FamilyMember{}
	}
	// Handle nested slice element
	builtValue := family.Build().(*models.FamilyMember)
	b.person.Family = append(b.person.Family, builtValue)
	return b
}

// WithHealth sets the Health
func (b *PersonBuilder) WithHealth(health *HealthProfileBuilder) *PersonBuilder {
	// Handle nested pointer
	b.person.Health = health.BuildPtr()
	return b
}

// WithDigital sets the Digital
func (b *PersonBuilder) WithDigital(digital *DigitalProfileBuilder) *PersonBuilder {
	// Handle nested pointer
	b.person.Digital = digital.BuildPtr()
	return b
}

// WithTravelHistory sets the TravelHistory
func (b *PersonBuilder) WithTravelHistory(travelHistory *TravelBuilder) *PersonBuilder {
	// Ensure the slice is initialized
	if b.person.TravelHistory == nil {
		b.person.TravelHistory = []*models.Travel{}
	}
	// Handle nested slice element
	builtValue := travelHistory.Build().(*models.Travel)
	b.person.TravelHistory = append(b.person.TravelHistory, builtValue)
	return b
}

// WithPreferences sets the Preferences
func (b *PersonBuilder) WithPreferences(preferences *PersonalPreferencesBuilder) *PersonBuilder {
	// Handle nested pointer
	b.person.Preferences = preferences.BuildPtr()
	return b
}

// WithCreatedAt sets the CreatedAt
func (b *PersonBuilder) WithCreatedAt(createdAt string) *PersonBuilder {
	b.person.CreatedAt = createdAt
	return b
}

// WithUpdatedAt sets the UpdatedAt
func (b *PersonBuilder) WithUpdatedAt(updatedAt string) *PersonBuilder {
	b.person.UpdatedAt = updatedAt
	return b
}


// WithValidation adds a custom validation function
func (b *PersonBuilder) WithValidation(validationFunc func(*models.Person) error) *PersonBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Person
func (b *PersonBuilder) Build() interface{} {
	return b.person
}

// BuildPtr builds the Person and returns a pointer
func (b *PersonBuilder) BuildPtr() *models.Person {
	return b.person
}

// BuildAndValidate builds the Person and validates it
func (b *PersonBuilder) BuildAndValidate() (*models.Person, error) {
	person := b.person

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(person); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(person).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return person, err
		}
	}

	return person, nil
}

// MustBuild builds the Person and panics if validation fails
func (b *PersonBuilder) MustBuild() *models.Person {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *PersonBuilder) Clone() *PersonBuilder {
	clonedPerson := *b.person
	return &PersonBuilder{
		person: &clonedPerson,
		validationFuncs: append([]func(*models.Person) error{}, b.validationFuncs...),
	}
}
