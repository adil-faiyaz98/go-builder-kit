package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// ApplicationBuilder builds a Application model
type ApplicationBuilder struct {
	application *models.Application
	// Custom validation functions
	validationFuncs []func(*models.Application) error
}

// NewApplicationBuilder creates a new ApplicationBuilder
func NewApplicationBuilder() *ApplicationBuilder {
	return &ApplicationBuilder{
		application: &models.Application{
			Name: "",
			Version: "",
			InstallDate: "",
			LastUsed: "",
			Settings: map[string]string{},
			Permissions: []string{},
		},
		validationFuncs: []func(*models.Application) error{},
	}
}

// NewApplicationBuilderWithDefaults creates a new ApplicationBuilder with sensible defaults
func NewApplicationBuilderWithDefaults() *ApplicationBuilder {
	builder := NewApplicationBuilder()
	// Add default values here if needed
	return builder
}
// WithName sets the Name
func (b *ApplicationBuilder) WithName(name string) *ApplicationBuilder {
	b.application.Name = name
	return b
}

// WithVersion sets the Version
func (b *ApplicationBuilder) WithVersion(version string) *ApplicationBuilder {
	b.application.Version = version
	return b
}

// WithInstallDate sets the InstallDate
func (b *ApplicationBuilder) WithInstallDate(installDate string) *ApplicationBuilder {
	b.application.InstallDate = installDate
	return b
}

// WithLastUsed sets the LastUsed
func (b *ApplicationBuilder) WithLastUsed(lastUsed string) *ApplicationBuilder {
	b.application.LastUsed = lastUsed
	return b
}

// WithSettings sets the Settings
func (b *ApplicationBuilder) WithSettings(key string, val string) *ApplicationBuilder {
	if b.application.Settings == nil {
		b.application.Settings = make(map[string]string)
	}
	b.application.Settings[key] = val
	return b
}

// WithPermissions sets the Permissions
func (b *ApplicationBuilder) WithPermissions(permissions string) *ApplicationBuilder {
	b.application.Permissions = append(b.application.Permissions, permissions)
	return b
}


// WithValidation adds a custom validation function
func (b *ApplicationBuilder) WithValidation(validationFunc func(*models.Application) error) *ApplicationBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Application
func (b *ApplicationBuilder) Build() interface{} {
	return b.application
}

// BuildPtr builds the Application and returns a pointer
func (b *ApplicationBuilder) BuildPtr() *models.Application {
	return b.application
}

// BuildAndValidate builds the Application and validates it
func (b *ApplicationBuilder) BuildAndValidate() (*models.Application, error) {
	application := b.application

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(application); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(application).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return application, err
		}
	}

	return application, nil
}

// MustBuild builds the Application and panics if validation fails
func (b *ApplicationBuilder) MustBuild() *models.Application {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *ApplicationBuilder) Clone() *ApplicationBuilder {
	clonedApplication := *b.application
	return &ApplicationBuilder{
		application: &clonedApplication,
		validationFuncs: append([]func(*models.Application) error{}, b.validationFuncs...),
	}
}
