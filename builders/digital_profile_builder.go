package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// DigitalProfileBuilder builds a DigitalProfile model
type DigitalProfileBuilder struct {
	digitalProfile *models.DigitalProfile
	// Custom validation functions
	validationFuncs []func(*models.DigitalProfile) error
}

// NewDigitalProfileBuilder creates a new DigitalProfileBuilder
func NewDigitalProfileBuilder() *DigitalProfileBuilder {
	return &DigitalProfileBuilder{
		digitalProfile: &models.DigitalProfile{
			Devices: []*models.Device{},
			Accounts: []*models.OnlineAccount{},
			Preferences: nil,
			Activity: []*models.UserActivity{},
		},
		validationFuncs: []func(*models.DigitalProfile) error{},
	}
}

// NewDigitalProfileBuilderWithDefaults creates a new DigitalProfileBuilder with sensible defaults
func NewDigitalProfileBuilderWithDefaults() *DigitalProfileBuilder {
	builder := NewDigitalProfileBuilder()
	// Add default values here if needed
	return builder
}
// WithDevices sets the Devices
func (b *DigitalProfileBuilder) WithDevices(devices *DeviceBuilder) *DigitalProfileBuilder {
	// Ensure the slice is initialized
	if b.digitalProfile.Devices == nil {
		b.digitalProfile.Devices = []*models.Device{}
	}
	// Handle nested slice element
	builtValue := devices.Build().(*models.Device)
	b.digitalProfile.Devices = append(b.digitalProfile.Devices, builtValue)
	return b
}

// WithAccounts sets the Accounts
func (b *DigitalProfileBuilder) WithAccounts(accounts *OnlineAccountBuilder) *DigitalProfileBuilder {
	// Ensure the slice is initialized
	if b.digitalProfile.Accounts == nil {
		b.digitalProfile.Accounts = []*models.OnlineAccount{}
	}
	// Handle nested slice element
	builtValue := accounts.Build().(*models.OnlineAccount)
	b.digitalProfile.Accounts = append(b.digitalProfile.Accounts, builtValue)
	return b
}

// WithPreferences sets the Preferences
func (b *DigitalProfileBuilder) WithPreferences(preferences *UserPreferencesBuilder) *DigitalProfileBuilder {
	// Handle nested pointer
	b.digitalProfile.Preferences = preferences.BuildPtr()
	return b
}

// WithActivity sets the Activity
func (b *DigitalProfileBuilder) WithActivity(activity *UserActivityBuilder) *DigitalProfileBuilder {
	// Ensure the slice is initialized
	if b.digitalProfile.Activity == nil {
		b.digitalProfile.Activity = []*models.UserActivity{}
	}
	// Handle nested slice element
	builtValue := activity.Build().(*models.UserActivity)
	b.digitalProfile.Activity = append(b.digitalProfile.Activity, builtValue)
	return b
}


// WithValidation adds a custom validation function
func (b *DigitalProfileBuilder) WithValidation(validationFunc func(*models.DigitalProfile) error) *DigitalProfileBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the DigitalProfile
func (b *DigitalProfileBuilder) Build() interface{} {
	return b.digitalProfile
}

// BuildPtr builds the DigitalProfile and returns a pointer
func (b *DigitalProfileBuilder) BuildPtr() *models.DigitalProfile {
	return b.digitalProfile
}

// BuildAndValidate builds the DigitalProfile and validates it
func (b *DigitalProfileBuilder) BuildAndValidate() (*models.DigitalProfile, error) {
	digitalProfile := b.digitalProfile

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(digitalProfile); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(digitalProfile).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return digitalProfile, err
		}
	}

	return digitalProfile, nil
}

// MustBuild builds the DigitalProfile and panics if validation fails
func (b *DigitalProfileBuilder) MustBuild() *models.DigitalProfile {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *DigitalProfileBuilder) Clone() *DigitalProfileBuilder {
	clonedDigitalProfile := *b.digitalProfile
	return &DigitalProfileBuilder{
		digitalProfile: &clonedDigitalProfile,
		validationFuncs: append([]func(*models.DigitalProfile) error{}, b.validationFuncs...),
	}
}
