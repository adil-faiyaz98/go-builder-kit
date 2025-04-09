package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// OnlineAccountBuilder builds a OnlineAccount model
type OnlineAccountBuilder struct {
	onlineAccount *models.OnlineAccount
	// Custom validation functions
	validationFuncs []func(*models.OnlineAccount) error
}

// NewOnlineAccountBuilder creates a new OnlineAccountBuilder
func NewOnlineAccountBuilder() *OnlineAccountBuilder {
	return &OnlineAccountBuilder{
		onlineAccount: &models.OnlineAccount{
			Platform: "",
			Username: "",
			Email: "",
			CreationDate: "",
			LastLogin: "",
			Status: "",
			Settings: map[string]string{},
			Subscription: nil,
		},
		validationFuncs: []func(*models.OnlineAccount) error{},
	}
}

// NewOnlineAccountBuilderWithDefaults creates a new OnlineAccountBuilder with sensible defaults
func NewOnlineAccountBuilderWithDefaults() *OnlineAccountBuilder {
	builder := NewOnlineAccountBuilder()
	// Add default values here if needed
	return builder
}
// WithPlatform sets the Platform
func (b *OnlineAccountBuilder) WithPlatform(platform string) *OnlineAccountBuilder {
	b.onlineAccount.Platform = platform
	return b
}

// WithUsername sets the Username
func (b *OnlineAccountBuilder) WithUsername(username string) *OnlineAccountBuilder {
	b.onlineAccount.Username = username
	return b
}

// WithEmail sets the Email
func (b *OnlineAccountBuilder) WithEmail(email string) *OnlineAccountBuilder {
	b.onlineAccount.Email = email
	return b
}

// WithCreationDate sets the CreationDate
func (b *OnlineAccountBuilder) WithCreationDate(creationDate string) *OnlineAccountBuilder {
	b.onlineAccount.CreationDate = creationDate
	return b
}

// WithLastLogin sets the LastLogin
func (b *OnlineAccountBuilder) WithLastLogin(lastLogin string) *OnlineAccountBuilder {
	b.onlineAccount.LastLogin = lastLogin
	return b
}

// WithStatus sets the Status
func (b *OnlineAccountBuilder) WithStatus(status string) *OnlineAccountBuilder {
	b.onlineAccount.Status = status
	return b
}

// WithSettings sets the Settings
func (b *OnlineAccountBuilder) WithSettings(key string, val string) *OnlineAccountBuilder {
	if b.onlineAccount.Settings == nil {
		b.onlineAccount.Settings = make(map[string]string)
	}
	b.onlineAccount.Settings[key] = val
	return b
}

// WithSubscription sets the Subscription
func (b *OnlineAccountBuilder) WithSubscription(subscription *SubscriptionBuilder) *OnlineAccountBuilder {
	// Handle nested pointer
	b.onlineAccount.Subscription = subscription.BuildPtr()
	return b
}


// WithValidation adds a custom validation function
func (b *OnlineAccountBuilder) WithValidation(validationFunc func(*models.OnlineAccount) error) *OnlineAccountBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the OnlineAccount
func (b *OnlineAccountBuilder) Build() interface{} {
	return b.onlineAccount
}

// BuildPtr builds the OnlineAccount and returns a pointer
func (b *OnlineAccountBuilder) BuildPtr() *models.OnlineAccount {
	return b.onlineAccount
}

// BuildAndValidate builds the OnlineAccount and validates it
func (b *OnlineAccountBuilder) BuildAndValidate() (*models.OnlineAccount, error) {
	onlineAccount := b.onlineAccount

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(onlineAccount); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(onlineAccount).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return onlineAccount, err
		}
	}

	return onlineAccount, nil
}

// MustBuild builds the OnlineAccount and panics if validation fails
func (b *OnlineAccountBuilder) MustBuild() *models.OnlineAccount {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *OnlineAccountBuilder) Clone() *OnlineAccountBuilder {
	clonedOnlineAccount := *b.onlineAccount
	return &OnlineAccountBuilder{
		onlineAccount: &clonedOnlineAccount,
		validationFuncs: append([]func(*models.OnlineAccount) error{}, b.validationFuncs...),
	}
}
