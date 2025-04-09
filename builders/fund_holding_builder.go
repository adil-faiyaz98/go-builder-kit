package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// FundHoldingBuilder builds a FundHolding model
type FundHoldingBuilder struct {
	fundHolding *models.FundHolding
	// Custom validation functions
	validationFuncs []func(*models.FundHolding) error
}

// NewFundHoldingBuilder creates a new FundHoldingBuilder
func NewFundHoldingBuilder() *FundHoldingBuilder {
	return &FundHoldingBuilder{
		fundHolding: &models.FundHolding{
			Name: "",
			Type: "",
			Percentage: 0.0,
			Value: 0.0,
			PurchaseDate: "",
			Performance: 0.0,
		},
		validationFuncs: []func(*models.FundHolding) error{},
	}
}

// NewFundHoldingBuilderWithDefaults creates a new FundHoldingBuilder with sensible defaults
func NewFundHoldingBuilderWithDefaults() *FundHoldingBuilder {
	builder := NewFundHoldingBuilder()
	// Add default values here if needed
	return builder
}
// WithName sets the Name
func (b *FundHoldingBuilder) WithName(name string) *FundHoldingBuilder {
	b.fundHolding.Name = name
	return b
}

// WithType sets the Type
func (b *FundHoldingBuilder) WithType(value string) *FundHoldingBuilder {
	b.fundHolding.Type = value
	return b
}

// WithPercentage sets the Percentage
func (b *FundHoldingBuilder) WithPercentage(percentage float64) *FundHoldingBuilder {
	b.fundHolding.Percentage = percentage
	return b
}

// WithValue sets the Value
func (b *FundHoldingBuilder) WithValue(value float64) *FundHoldingBuilder {
	b.fundHolding.Value = value
	return b
}

// WithPurchaseDate sets the PurchaseDate
func (b *FundHoldingBuilder) WithPurchaseDate(purchaseDate string) *FundHoldingBuilder {
	b.fundHolding.PurchaseDate = purchaseDate
	return b
}

// WithPerformance sets the Performance
func (b *FundHoldingBuilder) WithPerformance(performance float64) *FundHoldingBuilder {
	b.fundHolding.Performance = performance
	return b
}


// WithValidation adds a custom validation function
func (b *FundHoldingBuilder) WithValidation(validationFunc func(*models.FundHolding) error) *FundHoldingBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the FundHolding
func (b *FundHoldingBuilder) Build() interface{} {
	return b.fundHolding
}

// BuildPtr builds the FundHolding and returns a pointer
func (b *FundHoldingBuilder) BuildPtr() *models.FundHolding {
	return b.fundHolding
}

// BuildAndValidate builds the FundHolding and validates it
func (b *FundHoldingBuilder) BuildAndValidate() (*models.FundHolding, error) {
	fundHolding := b.fundHolding

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(fundHolding); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(fundHolding).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return fundHolding, err
		}
	}

	return fundHolding, nil
}

// MustBuild builds the FundHolding and panics if validation fails
func (b *FundHoldingBuilder) MustBuild() *models.FundHolding {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *FundHoldingBuilder) Clone() *FundHoldingBuilder {
	clonedFundHolding := *b.fundHolding
	return &FundHoldingBuilder{
		fundHolding: &clonedFundHolding,
		validationFuncs: append([]func(*models.FundHolding) error{}, b.validationFuncs...),
	}
}
