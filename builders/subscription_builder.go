package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// SubscriptionBuilder builds a Subscription model
type SubscriptionBuilder struct {
	subscription *models.Subscription
	// Custom validation functions
	validationFuncs []func(*models.Subscription) error
}

// NewSubscriptionBuilder creates a new SubscriptionBuilder
func NewSubscriptionBuilder() *SubscriptionBuilder {
	return &SubscriptionBuilder{
		subscription: &models.Subscription{
			Plan: "",
			StartDate: "",
			EndDate: "",
			Price: 0.0,
			BillingCycle: "",
			AutoRenew: false,
			Status: "",
			Features: []string{},
		},
		validationFuncs: []func(*models.Subscription) error{},
	}
}

// NewSubscriptionBuilderWithDefaults creates a new SubscriptionBuilder with sensible defaults
func NewSubscriptionBuilderWithDefaults() *SubscriptionBuilder {
	builder := NewSubscriptionBuilder()
	// Add default values here if needed
	return builder
}
// WithPlan sets the Plan
func (b *SubscriptionBuilder) WithPlan(plan string) *SubscriptionBuilder {
	b.subscription.Plan = plan
	return b
}

// WithStartDate sets the StartDate
func (b *SubscriptionBuilder) WithStartDate(startDate string) *SubscriptionBuilder {
	b.subscription.StartDate = startDate
	return b
}

// WithEndDate sets the EndDate
func (b *SubscriptionBuilder) WithEndDate(endDate string) *SubscriptionBuilder {
	b.subscription.EndDate = endDate
	return b
}

// WithPrice sets the Price
func (b *SubscriptionBuilder) WithPrice(price float64) *SubscriptionBuilder {
	b.subscription.Price = price
	return b
}

// WithBillingCycle sets the BillingCycle
func (b *SubscriptionBuilder) WithBillingCycle(billingCycle string) *SubscriptionBuilder {
	b.subscription.BillingCycle = billingCycle
	return b
}

// WithAutoRenew sets the AutoRenew
func (b *SubscriptionBuilder) WithAutoRenew(autoRenew bool) *SubscriptionBuilder {
	b.subscription.AutoRenew = autoRenew
	return b
}

// WithStatus sets the Status
func (b *SubscriptionBuilder) WithStatus(status string) *SubscriptionBuilder {
	b.subscription.Status = status
	return b
}

// WithFeatures sets the Features
func (b *SubscriptionBuilder) WithFeatures(features string) *SubscriptionBuilder {
	b.subscription.Features = append(b.subscription.Features, features)
	return b
}


// WithValidation adds a custom validation function
func (b *SubscriptionBuilder) WithValidation(validationFunc func(*models.Subscription) error) *SubscriptionBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Subscription
func (b *SubscriptionBuilder) Build() interface{} {
	return b.subscription
}

// BuildPtr builds the Subscription and returns a pointer
func (b *SubscriptionBuilder) BuildPtr() *models.Subscription {
	return b.subscription
}

// BuildAndValidate builds the Subscription and validates it
func (b *SubscriptionBuilder) BuildAndValidate() (*models.Subscription, error) {
	subscription := b.subscription

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(subscription); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(subscription).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return subscription, err
		}
	}

	return subscription, nil
}

// MustBuild builds the Subscription and panics if validation fails
func (b *SubscriptionBuilder) MustBuild() *models.Subscription {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *SubscriptionBuilder) Clone() *SubscriptionBuilder {
	clonedSubscription := *b.subscription
	return &SubscriptionBuilder{
		subscription: &clonedSubscription,
		validationFuncs: append([]func(*models.Subscription) error{}, b.validationFuncs...),
	}
}
