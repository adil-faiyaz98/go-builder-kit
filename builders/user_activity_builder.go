package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// UserActivityBuilder builds a UserActivity model
type UserActivityBuilder struct {
	userActivity *models.UserActivity
	// Custom validation functions
	validationFuncs []func(*models.UserActivity) error
}

// NewUserActivityBuilder creates a new UserActivityBuilder
func NewUserActivityBuilder() *UserActivityBuilder {
	return &UserActivityBuilder{
		userActivity: &models.UserActivity{
			Type: "",
			Platform: "",
			Date: "",
			Timestamp: "",
			Duration: 0,
			Description: "",
			Metadata: map[string]string{},
			Device: "",
			Location: nil,
			Details: map[string]string{},
		},
		validationFuncs: []func(*models.UserActivity) error{},
	}
}

// NewUserActivityBuilderWithDefaults creates a new UserActivityBuilder with sensible defaults
func NewUserActivityBuilderWithDefaults() *UserActivityBuilder {
	builder := NewUserActivityBuilder()
	// Add default values here if needed
	return builder
}
// WithType sets the Type
func (b *UserActivityBuilder) WithType(value string) *UserActivityBuilder {
	b.userActivity.Type = value
	return b
}

// WithPlatform sets the Platform
func (b *UserActivityBuilder) WithPlatform(platform string) *UserActivityBuilder {
	b.userActivity.Platform = platform
	return b
}

// WithDate sets the Date
func (b *UserActivityBuilder) WithDate(date string) *UserActivityBuilder {
	b.userActivity.Date = date
	return b
}

// WithTimestamp sets the Timestamp
func (b *UserActivityBuilder) WithTimestamp(timestamp string) *UserActivityBuilder {
	b.userActivity.Timestamp = timestamp
	return b
}

// WithDuration sets the Duration
func (b *UserActivityBuilder) WithDuration(duration int) *UserActivityBuilder {
	b.userActivity.Duration = duration
	return b
}

// WithDescription sets the Description
func (b *UserActivityBuilder) WithDescription(description string) *UserActivityBuilder {
	b.userActivity.Description = description
	return b
}

// WithMetadata sets the Metadata
func (b *UserActivityBuilder) WithMetadata(key string, val string) *UserActivityBuilder {
	if b.userActivity.Metadata == nil {
		b.userActivity.Metadata = make(map[string]string)
	}
	b.userActivity.Metadata[key] = val
	return b
}

// WithDevice sets the Device
func (b *UserActivityBuilder) WithDevice(device string) *UserActivityBuilder {
	b.userActivity.Device = device
	return b
}

// WithLocation sets the Location
func (b *UserActivityBuilder) WithLocation(location *GeoLocationBuilder) *UserActivityBuilder {
	// Handle nested pointer
	b.userActivity.Location = location.BuildPtr()
	return b
}

// WithDetails sets the Details
func (b *UserActivityBuilder) WithDetails(key string, val string) *UserActivityBuilder {
	if b.userActivity.Details == nil {
		b.userActivity.Details = make(map[string]string)
	}
	b.userActivity.Details[key] = val
	return b
}


// WithValidation adds a custom validation function
func (b *UserActivityBuilder) WithValidation(validationFunc func(*models.UserActivity) error) *UserActivityBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the UserActivity
func (b *UserActivityBuilder) Build() interface{} {
	return b.userActivity
}

// BuildPtr builds the UserActivity and returns a pointer
func (b *UserActivityBuilder) BuildPtr() *models.UserActivity {
	return b.userActivity
}

// BuildAndValidate builds the UserActivity and validates it
func (b *UserActivityBuilder) BuildAndValidate() (*models.UserActivity, error) {
	userActivity := b.userActivity

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(userActivity); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(userActivity).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return userActivity, err
		}
	}

	return userActivity, nil
}

// MustBuild builds the UserActivity and panics if validation fails
func (b *UserActivityBuilder) MustBuild() *models.UserActivity {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *UserActivityBuilder) Clone() *UserActivityBuilder {
	clonedUserActivity := *b.userActivity
	return &UserActivityBuilder{
		userActivity: &clonedUserActivity,
		validationFuncs: append([]func(*models.UserActivity) error{}, b.validationFuncs...),
	}
}
