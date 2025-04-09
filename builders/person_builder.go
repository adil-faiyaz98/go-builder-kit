package builders

import (
	"fmt"

	"github.com/adil-faiyaz98/go-builder-kit/models"
)

// PersonBuilder builds a Person model
type PersonBuilder struct {
	person          *models.Person
	validationFuncs []func(*models.Person) error
}

// NewPersonBuilder creates a new PersonBuilder
func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{
		person: &models.Person{
			ID:            "",
			Name:          "",
			Age:           0,
			Email:         "",
			Phone:         "",
			Birthdate:     "",
			Gender:        "",
			Nationality:   "",
			MaritalStatus: "",
			Address:       nil,
			Education:     nil,
			Profile:       nil,
			Bank:          nil,
			Employment:    nil,
			Friends:       []*models.Person{},
			Family:        []*models.FamilyMember{},
			Health:        nil,
			Digital:       nil,
			TravelHistory: []*models.Travel{},
			Preferences:   nil,
			CreatedAt:     "",
			UpdatedAt:     "",
		},
		validationFuncs: []func(*models.Person) error{},
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
	builtValue := address.Build().(*models.Address)
	b.person.Address = builtValue
	return b
}

// WithEducation sets the Education
func (b *PersonBuilder) WithEducation(education *EducationBuilder) *PersonBuilder {
	builtValue := education.Build().(*models.Education)
	b.person.Education = builtValue
	return b
}

// WithProfile sets the Profile
func (b *PersonBuilder) WithProfile(profile any) *PersonBuilder {
	b.person.Profile = profile
	return b
}

// WithBank sets the Bank
func (b *PersonBuilder) WithBank(bank *BankBuilder) *PersonBuilder {
	builtValue := bank.Build().(*models.Bank)
	b.person.Bank = builtValue
	return b
}

// WithEmployment sets the Employment
func (b *PersonBuilder) WithEmployment(employment *EmploymentBuilder) *PersonBuilder {
	builtValue := employment.Build().(*models.Employment)
	b.person.Employment = builtValue
	return b
}

// WithFriend adds a friend to the Friends slice
func (b *PersonBuilder) WithFriend(friend *PersonBuilder) *PersonBuilder {
	builtValue := friend.Build().(*models.Person)
	b.person.Friends = append(b.person.Friends, builtValue)
	return b
}

// WithFamilyMember adds a family member to the Family slice
func (b *PersonBuilder) WithFamilyMember(familyMember *FamilyMemberBuilder) *PersonBuilder {
	builtValue := familyMember.Build().(*models.FamilyMember)
	b.person.Family = append(b.person.Family, builtValue)
	return b
}

// WithHealth sets the Health
func (b *PersonBuilder) WithHealth(health any) *PersonBuilder {
	b.person.Health = health
	return b
}

// WithDigital sets the Digital
func (b *PersonBuilder) WithDigital(digital any) *PersonBuilder {
	b.person.Digital = digital
	return b
}

// WithTravel adds a travel to the TravelHistory slice
func (b *PersonBuilder) WithTravel(travel *models.Travel) *PersonBuilder {
	b.person.TravelHistory = append(b.person.TravelHistory, travel)
	return b
}

// WithPreferences sets the Preferences
func (b *PersonBuilder) WithPreferences(preferences *PersonalPreferencesBuilder) *PersonBuilder {
	builtValue := preferences.Build().(*models.PersonalPreferences)
	b.person.Preferences = builtValue
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
func (b *PersonBuilder) Build() any {
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
func (b *PersonBuilder) MustBuild() *models.Person {
	person, err := b.BuildAndValidate()
	if err != nil {
		panic(fmt.Sprintf("Person validation failed: %s", err.Error()))
	}
	return person
}

// Clone creates a deep copy of the PersonBuilder
func (b *PersonBuilder) Clone() *PersonBuilder {
	clonedPerson := *b.person
	return &PersonBuilder{
		person:          &clonedPerson,
		validationFuncs: append([]func(*models.Person) error{}, b.validationFuncs...),
	}
}
