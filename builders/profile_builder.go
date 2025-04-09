package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// ProfileBuilder builds a Profile model
type ProfileBuilder struct {
	profile *models.Profile
	// Custom validation functions
	validationFuncs []func(*models.Profile) error
}

// NewProfileBuilder creates a new ProfileBuilder
func NewProfileBuilder() *ProfileBuilder {
	return &ProfileBuilder{
		profile: &models.Profile{
			Address: models.Address{},
			Education: []string{},
			Skills: []*models.Skill{},
			Certifications: []*models.Certification{},
			SocialMedia: models.SocialMedia{},
			Biography: "",
			Interests: []string{},
			Languages: []*models.Language{},
		},
		validationFuncs: []func(*models.Profile) error{},
	}
}

// NewProfileBuilderWithDefaults creates a new ProfileBuilder with sensible defaults
func NewProfileBuilderWithDefaults() *ProfileBuilder {
	builder := NewProfileBuilder()
	// Add default values here if needed
	return builder
}
// WithAddress sets the Address
func (b *ProfileBuilder) WithAddress(address *AddressBuilder) *ProfileBuilder {
	builtValue := address.Build().(*models.Address)
	b.profile.Address = *builtValue
	return b
}

// WithEducation sets the Education
func (b *ProfileBuilder) WithEducation(education string) *ProfileBuilder {
	b.profile.Education = append(b.profile.Education, education)
	return b
}

// WithSkills sets the Skills
func (b *ProfileBuilder) WithSkills(skills *SkillBuilder) *ProfileBuilder {
	// Ensure the slice is initialized
	if b.profile.Skills == nil {
		b.profile.Skills = []*models.Skill{}
	}
	// Handle nested slice element
	builtValue := skills.Build().(*models.Skill)
	b.profile.Skills = append(b.profile.Skills, builtValue)
	return b
}

// WithCertifications sets the Certifications
func (b *ProfileBuilder) WithCertifications(certifications *CertificationBuilder) *ProfileBuilder {
	// Ensure the slice is initialized
	if b.profile.Certifications == nil {
		b.profile.Certifications = []*models.Certification{}
	}
	// Handle nested slice element
	builtValue := certifications.Build().(*models.Certification)
	b.profile.Certifications = append(b.profile.Certifications, builtValue)
	return b
}

// WithSocialMedia sets the SocialMedia
func (b *ProfileBuilder) WithSocialMedia(socialMedia *SocialMediaBuilder) *ProfileBuilder {
	builtValue := socialMedia.Build().(*models.SocialMedia)
	b.profile.SocialMedia = *builtValue
	return b
}

// WithBiography sets the Biography
func (b *ProfileBuilder) WithBiography(biography string) *ProfileBuilder {
	b.profile.Biography = biography
	return b
}

// WithInterests sets the Interests
func (b *ProfileBuilder) WithInterests(interests string) *ProfileBuilder {
	b.profile.Interests = append(b.profile.Interests, interests)
	return b
}

// WithLanguages sets the Languages
func (b *ProfileBuilder) WithLanguages(languages *LanguageBuilder) *ProfileBuilder {
	// Ensure the slice is initialized
	if b.profile.Languages == nil {
		b.profile.Languages = []*models.Language{}
	}
	// Handle nested slice element
	builtValue := languages.Build().(*models.Language)
	b.profile.Languages = append(b.profile.Languages, builtValue)
	return b
}


// WithValidation adds a custom validation function
func (b *ProfileBuilder) WithValidation(validationFunc func(*models.Profile) error) *ProfileBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Profile
func (b *ProfileBuilder) Build() interface{} {
	return b.profile
}

// BuildPtr builds the Profile and returns a pointer
func (b *ProfileBuilder) BuildPtr() *models.Profile {
	return b.profile
}

// BuildAndValidate builds the Profile and validates it
func (b *ProfileBuilder) BuildAndValidate() (*models.Profile, error) {
	profile := b.profile

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(profile); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(profile).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return profile, err
		}
	}

	return profile, nil
}

// MustBuild builds the Profile and panics if validation fails
func (b *ProfileBuilder) MustBuild() *models.Profile {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *ProfileBuilder) Clone() *ProfileBuilder {
	clonedProfile := *b.profile
	return &ProfileBuilder{
		profile: &clonedProfile,
		validationFuncs: append([]func(*models.Profile) error{}, b.validationFuncs...),
	}
}
