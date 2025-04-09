package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// LanguageBuilder builds a Language model
type LanguageBuilder struct {
	language *models.Language
	// Custom validation functions
	validationFuncs []func(*models.Language) error
}

// NewLanguageBuilder creates a new LanguageBuilder
func NewLanguageBuilder() *LanguageBuilder {
	return &LanguageBuilder{
		language: &models.Language{
			Name: "",
			Proficiency: "",
			Certification: nil,
		},
		validationFuncs: []func(*models.Language) error{},
	}
}

// NewLanguageBuilderWithDefaults creates a new LanguageBuilder with sensible defaults
func NewLanguageBuilderWithDefaults() *LanguageBuilder {
	builder := NewLanguageBuilder()
	// Add default values here if needed
	return builder
}
// WithName sets the Name
func (b *LanguageBuilder) WithName(name string) *LanguageBuilder {
	b.language.Name = name
	return b
}

// WithProficiency sets the Proficiency
func (b *LanguageBuilder) WithProficiency(proficiency string) *LanguageBuilder {
	b.language.Proficiency = proficiency
	return b
}

// WithCertification sets the Certification
func (b *LanguageBuilder) WithCertification(certification *CertificationBuilder) *LanguageBuilder {
	// Handle nested pointer
	b.language.Certification = certification.BuildPtr()
	return b
}


// WithValidation adds a custom validation function
func (b *LanguageBuilder) WithValidation(validationFunc func(*models.Language) error) *LanguageBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Language
func (b *LanguageBuilder) Build() interface{} {
	return b.language
}

// BuildPtr builds the Language and returns a pointer
func (b *LanguageBuilder) BuildPtr() *models.Language {
	return b.language
}

// BuildAndValidate builds the Language and validates it
func (b *LanguageBuilder) BuildAndValidate() (*models.Language, error) {
	language := b.language

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(language); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(language).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return language, err
		}
	}

	return language, nil
}

// MustBuild builds the Language and panics if validation fails
func (b *LanguageBuilder) MustBuild() *models.Language {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *LanguageBuilder) Clone() *LanguageBuilder {
	clonedLanguage := *b.language
	return &LanguageBuilder{
		language: &clonedLanguage,
		validationFuncs: append([]func(*models.Language) error{}, b.validationFuncs...),
	}
}
