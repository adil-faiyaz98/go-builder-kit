package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// UserPreferencesBuilder builds a UserPreferences model
type UserPreferencesBuilder struct {
	userPreferences *models.UserPreferences
	// Custom validation functions
	validationFuncs []func(*models.UserPreferences) error
}

// NewUserPreferencesBuilder creates a new UserPreferencesBuilder
func NewUserPreferencesBuilder() *UserPreferencesBuilder {
	return &UserPreferencesBuilder{
		userPreferences: &models.UserPreferences{
			Theme: "",
			Language: "",
			Notifications: map[string]bool{},
			Privacy: map[string]string{},
			Accessibility: map[string]bool{},
			DefaultCommunication: "",
		},
		validationFuncs: []func(*models.UserPreferences) error{},
	}
}

// NewUserPreferencesBuilderWithDefaults creates a new UserPreferencesBuilder with sensible defaults
func NewUserPreferencesBuilderWithDefaults() *UserPreferencesBuilder {
	builder := NewUserPreferencesBuilder()
	// Add default values here if needed
	return builder
}
// WithTheme sets the Theme
func (b *UserPreferencesBuilder) WithTheme(theme string) *UserPreferencesBuilder {
	b.userPreferences.Theme = theme
	return b
}

// WithLanguage sets the Language
func (b *UserPreferencesBuilder) WithLanguage(language string) *UserPreferencesBuilder {
	b.userPreferences.Language = language
	return b
}

// WithNotifications sets the Notifications
func (b *UserPreferencesBuilder) WithNotifications(key string, val bool) *UserPreferencesBuilder {
	if b.userPreferences.Notifications == nil {
		b.userPreferences.Notifications = make(map[string]bool)
	}
	b.userPreferences.Notifications[key] = val
	return b
}

// WithPrivacy sets the Privacy
func (b *UserPreferencesBuilder) WithPrivacy(key string, val string) *UserPreferencesBuilder {
	if b.userPreferences.Privacy == nil {
		b.userPreferences.Privacy = make(map[string]string)
	}
	b.userPreferences.Privacy[key] = val
	return b
}

// WithAccessibility sets the Accessibility
func (b *UserPreferencesBuilder) WithAccessibility(key string, val bool) *UserPreferencesBuilder {
	if b.userPreferences.Accessibility == nil {
		b.userPreferences.Accessibility = make(map[string]bool)
	}
	b.userPreferences.Accessibility[key] = val
	return b
}

// WithDefaultCommunication sets the DefaultCommunication
func (b *UserPreferencesBuilder) WithDefaultCommunication(defaultCommunication string) *UserPreferencesBuilder {
	b.userPreferences.DefaultCommunication = defaultCommunication
	return b
}


// WithValidation adds a custom validation function
func (b *UserPreferencesBuilder) WithValidation(validationFunc func(*models.UserPreferences) error) *UserPreferencesBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the UserPreferences
func (b *UserPreferencesBuilder) Build() interface{} {
	return b.userPreferences
}

// BuildPtr builds the UserPreferences and returns a pointer
func (b *UserPreferencesBuilder) BuildPtr() *models.UserPreferences {
	return b.userPreferences
}

// BuildAndValidate builds the UserPreferences and validates it
func (b *UserPreferencesBuilder) BuildAndValidate() (*models.UserPreferences, error) {
	userPreferences := b.userPreferences

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(userPreferences); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(userPreferences).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return userPreferences, err
		}
	}

	return userPreferences, nil
}

// MustBuild builds the UserPreferences and panics if validation fails
func (b *UserPreferencesBuilder) MustBuild() *models.UserPreferences {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *UserPreferencesBuilder) Clone() *UserPreferencesBuilder {
	clonedUserPreferences := *b.userPreferences
	return &UserPreferencesBuilder{
		userPreferences: &clonedUserPreferences,
		validationFuncs: append([]func(*models.UserPreferences) error{}, b.validationFuncs...),
	}
}
