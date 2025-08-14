package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/v2/models"
	"github.com/adil-faiyaz98/go-builder-kit/v2/pkg/builder"
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
			Profile: interface{}(0),
			Bank: nil,
			Employment: nil,
			Friends: []*models.Person{},
			Family: []*models.FamilyMember{},
			Health: interface{}(0),
			Digital: interface{}(0),
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
func (b *PersonBuilder) WithID(id string) *PersonBuilder {
	if b == nil {
		return b
	}
	b.person.ID = builder.SanitizeString(id)
	return b
}

// WithName sets the Name
func (b *PersonBuilder) WithName(name string) *PersonBuilder {
	if b == nil {
		return b
	}
	b.person.Name = builder.SanitizeString(name)
	return b
}

// WithAge sets the Age
func (b *PersonBuilder) WithAge(age int) *PersonBuilder {
	if b == nil {
		return b
	}
	b.person.Age = age
	return b
}

// WithEmail sets the Email
func (b *PersonBuilder) WithEmail(email string) *PersonBuilder {
	if b == nil {
		return b
	}
	b.person.Email = builder.SanitizeString(email)
	return b
}

// WithPhone sets the Phone
func (b *PersonBuilder) WithPhone(phone string) *PersonBuilder {
	if b == nil {
		return b
	}
	b.person.Phone = builder.SanitizeString(phone)
	return b
}

// WithBirthdate sets the Birthdate
func (b *PersonBuilder) WithBirthdate(birthdate string) *PersonBuilder {
	if b == nil {
		return b
	}
	b.person.Birthdate = builder.SanitizeString(birthdate)
	return b
}

// WithGender sets the Gender
func (b *PersonBuilder) WithGender(gender string) *PersonBuilder {
	if b == nil {
		return b
	}
	b.person.Gender = builder.SanitizeString(gender)
	return b
}

// WithNationality sets the Nationality
func (b *PersonBuilder) WithNationality(nationality string) *PersonBuilder {
	if b == nil {
		return b
	}
	b.person.Nationality = builder.SanitizeString(nationality)
	return b
}

// WithMaritalStatus sets the MaritalStatus
func (b *PersonBuilder) WithMaritalStatus(maritalStatus string) *PersonBuilder {
	if b == nil {
		return b
	}
	b.person.MaritalStatus = builder.SanitizeString(maritalStatus)
	return b
}

// WithAddress sets the Address
func (b *PersonBuilder) WithAddress(address *AddressBuilder) *PersonBuilder {
	if b == nil {
		return b
	}
	// Handle nested pointer
	b.person.Address = address.BuildPtr()
	return b
}

// WithEducation sets the Education
func (b *PersonBuilder) WithEducation(education *EducationBuilder) *PersonBuilder {
	if b == nil {
		return b
	}
	// Handle nested pointer
	b.person.Education = education.BuildPtr()
	return b
}

// WithProfile sets the Profile
func (b *PersonBuilder) WithProfile(profile interface{}) *PersonBuilder {
	if b == nil {
		return b
	}
	b.person.Profile = profile
	return b
}

// WithBank sets the Bank
func (b *PersonBuilder) WithBank(bank *BankBuilder) *PersonBuilder {
	if b == nil {
		return b
	}
	// Handle nested pointer
	b.person.Bank = bank.BuildPtr()
	return b
}

// WithEmployment sets the Employment
func (b *PersonBuilder) WithEmployment(employment *EmploymentBuilder) *PersonBuilder {
	if b == nil {
		return b
	}
	// Handle nested pointer
	b.person.Employment = employment.BuildPtr()
	return b
}

// WithFriends sets the Friends
func (b *PersonBuilder) WithFriends(friends []*PersonBuilder) *PersonBuilder {
	if b == nil {
		return b
	}
	// Ensure the slice is initialized
	if b.person.Friends == nil {
		b.person.Friends = []*models.Person{}
	}
	// Handle nested slice elements
	// Initialize the slice
	b.person.Friends = make([]*models.Person, 0, len(friends))
	// Convert each builder to its model
	for _, builder := range friends {
		builtValue := builder.Build().(*models.Person)
		b.person.Friends = append(b.person.Friends, builtValue)
	}
	return b
}

// WithFamily sets the Family
func (b *PersonBuilder) WithFamily(family []*FamilyMemberBuilder) *PersonBuilder {
	if b == nil {
		return b
	}
	// Ensure the slice is initialized
	if b.person.Family == nil {
		b.person.Family = []*models.FamilyMember{}
	}
	// Handle nested slice elements
	// Initialize the slice
	b.person.Family = make([]*models.FamilyMember, 0, len(family))
	// Convert each builder to its model
	for _, builder := range family {
		builtValue := builder.Build().(*models.FamilyMember)
		b.person.Family = append(b.person.Family, builtValue)
	}
	return b
}

// WithHealth sets the Health
func (b *PersonBuilder) WithHealth(health interface{}) *PersonBuilder {
	if b == nil {
		return b
	}
	b.person.Health = health
	return b
}

// WithDigital sets the Digital
func (b *PersonBuilder) WithDigital(digital interface{}) *PersonBuilder {
	if b == nil {
		return b
	}
	b.person.Digital = digital
	return b
}

// WithTravelHistory sets the TravelHistory
func (b *PersonBuilder) WithTravelHistory(travelHistory []*TravelBuilder) *PersonBuilder {
	if b == nil {
		return b
	}
	// Ensure the slice is initialized
	if b.person.TravelHistory == nil {
		b.person.TravelHistory = []*models.Travel{}
	}
	// Handle nested slice elements
	// Initialize the slice
	b.person.TravelHistory = make([]*models.Travel, 0, len(travelHistory))
	// Convert each builder to its model
	for _, builder := range travelHistory {
		builtValue := builder.Build().(*models.Travel)
		b.person.TravelHistory = append(b.person.TravelHistory, builtValue)
	}
	return b
}

// WithPreferences sets the Preferences
func (b *PersonBuilder) WithPreferences(preferences *PersonalPreferencesBuilder) *PersonBuilder {
	if b == nil {
		return b
	}
	// Handle nested pointer
	b.person.Preferences = preferences.BuildPtr()
	return b
}

// WithCreatedAt sets the CreatedAt
func (b *PersonBuilder) WithCreatedAt(createdAt string) *PersonBuilder {
	if b == nil {
		return b
	}
	b.person.CreatedAt = builder.SanitizeString(createdAt)
	return b
}

// WithUpdatedAt sets the UpdatedAt
func (b *PersonBuilder) WithUpdatedAt(updatedAt string) *PersonBuilder {
	if b == nil {
		return b
	}
	b.person.UpdatedAt = builder.SanitizeString(updatedAt)
	return b
}

// AddFriend adds a single item to the Friends slice
func (b *PersonBuilder) AddFriend(friend *PersonBuilder) *PersonBuilder {
	if b == nil {
		return b
	}
	// Ensure the slice is initialized with capacity
	if b.person.Friends == nil {
		b.person.Friends = make([]*models.Person, 0, 4) // Pre-allocate capacity
	}
	// Handle nested slice element
	if friend != nil {
		builtValue := friend.Build().(*models.Person)
		b.person.Friends = append(b.person.Friends, builtValue)
	}
	return b
}
// AddFamily adds a single item to the Family slice
func (b *PersonBuilder) AddFamily(family *FamilyMemberBuilder) *PersonBuilder {
	if b == nil {
		return b
	}
	// Ensure the slice is initialized with capacity
	if b.person.Family == nil {
		b.person.Family = make([]*models.FamilyMember, 0, 4) // Pre-allocate capacity
	}
	// Handle nested slice element
	if family != nil {
		builtValue := family.Build().(*models.FamilyMember)
		b.person.Family = append(b.person.Family, builtValue)
	}
	return b
}
// AddTravelHistory adds a single item to the TravelHistory slice
func (b *PersonBuilder) AddTravelHistory(travelHistory *TravelBuilder) *PersonBuilder {
	if b == nil {
		return b
	}
	// Ensure the slice is initialized with capacity
	if b.person.TravelHistory == nil {
		b.person.TravelHistory = make([]*models.Travel, 0, 4) // Pre-allocate capacity
	}
	// Handle nested slice element
	if travelHistory != nil {
		builtValue := travelHistory.Build().(*models.Travel)
		b.person.TravelHistory = append(b.person.TravelHistory, builtValue)
	}
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
	if b == nil || b.person == nil {
		return nil, fmt.Errorf("builder is not properly initialized")
	}

	person := b.person

	// Run custom validation functions
	for i, validationFunc := range b.validationFuncs {
		if validationFunc == nil {
			continue // Skip nil validators
		}
		if err := validationFunc(person); err != nil {
			return nil, fmt.Errorf("custom validation failed at index %d: %w", i, err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := any(person).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return person, fmt.Errorf("model validation failed: %w", err)
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
	if b == nil || b.person == nil {
		return NewPersonBuilder()
	}

	// Deep copy the struct
	clonedPerson := *b.person

	// Create new builder with cloned data
	clonedBuilder := &PersonBuilder{
		person: &clonedPerson,
		validationFuncs: make([]func(*models.Person) error, 0, len(b.validationFuncs)),
	}

	// Copy validation functions safely
	if len(b.validationFuncs) > 0 {
		clonedBuilder.validationFuncs = append(clonedBuilder.validationFuncs, b.validationFuncs...)
	}

	return clonedBuilder
}
