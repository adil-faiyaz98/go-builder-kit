package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// SocialMediaBuilder builds a SocialMedia model
type SocialMediaBuilder struct {
	socialMedia *models.SocialMedia
	// Custom validation functions
	validationFuncs []func(*models.SocialMedia) error
}

// NewSocialMediaBuilder creates a new SocialMediaBuilder
func NewSocialMediaBuilder() *SocialMediaBuilder {
	return &SocialMediaBuilder{
		socialMedia: &models.SocialMedia{
			LinkedIn: "",
			Twitter: "",
			Facebook: "",
			Instagram: "",
			GitHub: "",
			Website: "",
			OtherProfiles: map[string]string{},
		},
		validationFuncs: []func(*models.SocialMedia) error{},
	}
}

// NewSocialMediaBuilderWithDefaults creates a new SocialMediaBuilder with sensible defaults
func NewSocialMediaBuilderWithDefaults() *SocialMediaBuilder {
	builder := NewSocialMediaBuilder()
	// Add default values here if needed
	return builder
}
// WithLinkedIn sets the LinkedIn
func (b *SocialMediaBuilder) WithLinkedIn(linkedIn string) *SocialMediaBuilder {
	b.socialMedia.LinkedIn = linkedIn
	return b
}

// WithTwitter sets the Twitter
func (b *SocialMediaBuilder) WithTwitter(twitter string) *SocialMediaBuilder {
	b.socialMedia.Twitter = twitter
	return b
}

// WithFacebook sets the Facebook
func (b *SocialMediaBuilder) WithFacebook(facebook string) *SocialMediaBuilder {
	b.socialMedia.Facebook = facebook
	return b
}

// WithInstagram sets the Instagram
func (b *SocialMediaBuilder) WithInstagram(instagram string) *SocialMediaBuilder {
	b.socialMedia.Instagram = instagram
	return b
}

// WithGitHub sets the GitHub
func (b *SocialMediaBuilder) WithGitHub(gitHub string) *SocialMediaBuilder {
	b.socialMedia.GitHub = gitHub
	return b
}

// WithWebsite sets the Website
func (b *SocialMediaBuilder) WithWebsite(website string) *SocialMediaBuilder {
	b.socialMedia.Website = website
	return b
}

// WithOtherProfiles sets the OtherProfiles
func (b *SocialMediaBuilder) WithOtherProfiles(key string, val string) *SocialMediaBuilder {
	if b.socialMedia.OtherProfiles == nil {
		b.socialMedia.OtherProfiles = make(map[string]string)
	}
	b.socialMedia.OtherProfiles[key] = val
	return b
}


// WithValidation adds a custom validation function
func (b *SocialMediaBuilder) WithValidation(validationFunc func(*models.SocialMedia) error) *SocialMediaBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the SocialMedia
func (b *SocialMediaBuilder) Build() interface{} {
	return b.socialMedia
}

// BuildPtr builds the SocialMedia and returns a pointer
func (b *SocialMediaBuilder) BuildPtr() *models.SocialMedia {
	return b.socialMedia
}

// BuildAndValidate builds the SocialMedia and validates it
func (b *SocialMediaBuilder) BuildAndValidate() (*models.SocialMedia, error) {
	socialMedia := b.socialMedia

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(socialMedia); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(socialMedia).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return socialMedia, err
		}
	}

	return socialMedia, nil
}

// MustBuild builds the SocialMedia and panics if validation fails
func (b *SocialMediaBuilder) MustBuild() *models.SocialMedia {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *SocialMediaBuilder) Clone() *SocialMediaBuilder {
	clonedSocialMedia := *b.socialMedia
	return &SocialMediaBuilder{
		socialMedia: &clonedSocialMedia,
		validationFuncs: append([]func(*models.SocialMedia) error{}, b.validationFuncs...),
	}
}
