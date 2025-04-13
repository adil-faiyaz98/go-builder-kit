package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

// PortfolioBuilder builds a Portfolio model
type PortfolioBuilder struct {
	portfolio *models.Portfolio
	// Custom validation functions
	validationFuncs []func(*models.Portfolio) error
}

// NewPortfolioBuilder creates a new PortfolioBuilder
func NewPortfolioBuilder() *PortfolioBuilder {
	return &PortfolioBuilder{
		portfolio: &models.Portfolio{
			ID: "",
			Name: "",
			Description: "",
			RiskLevel: "",
			TotalValue: 0.0,
			Allocation: map[string]float64{},
			Stocks: []any{},
			Bonds: []*models.Bond{},
			ETFs: []any{},
			MutualFunds: []any{},
			Cryptocurrencies: []any{},
		},
		validationFuncs: []func(*models.Portfolio) error{},
	}
}

// NewPortfolioBuilderWithDefaults creates a new PortfolioBuilder with sensible defaults
func NewPortfolioBuilderWithDefaults() *PortfolioBuilder {
	builder := NewPortfolioBuilder()
	// Add default values here if needed
	return builder
}
// WithID sets the ID
func (b *PortfolioBuilder) WithID(id string) *PortfolioBuilder {
	b.portfolio.ID = id
	return b
}

// WithName sets the Name
func (b *PortfolioBuilder) WithName(name string) *PortfolioBuilder {
	b.portfolio.Name = name
	return b
}

// WithDescription sets the Description
func (b *PortfolioBuilder) WithDescription(description string) *PortfolioBuilder {
	b.portfolio.Description = description
	return b
}

// WithRiskLevel sets the RiskLevel
func (b *PortfolioBuilder) WithRiskLevel(riskLevel string) *PortfolioBuilder {
	b.portfolio.RiskLevel = riskLevel
	return b
}

// WithTotalValue sets the TotalValue
func (b *PortfolioBuilder) WithTotalValue(totalValue float64) *PortfolioBuilder {
	b.portfolio.TotalValue = totalValue
	return b
}

// WithAllocation sets the Allocation
func (b *PortfolioBuilder) WithAllocation(key string, val float64) *PortfolioBuilder {
	if b.portfolio.Allocation == nil {
		b.portfolio.Allocation = make(map[string]float64)
	}
	b.portfolio.Allocation[key] = val
	return b
}

// WithStocks sets the Stocks
func (b *PortfolioBuilder) WithStocks(stocks []any) *PortfolioBuilder {
	b.portfolio.Stocks = append(b.portfolio.Stocks, stocks...)
	return b
}

// WithBonds sets the Bonds
func (b *PortfolioBuilder) WithBonds(bonds []*BondBuilder) *PortfolioBuilder {
	// Ensure the slice is initialized
	if b.portfolio.Bonds == nil {
		b.portfolio.Bonds = []*models.Bond{}
	}
	// Handle nested slice elements
	// Initialize the slice
	b.portfolio.Bonds = make([]*models.Bond, 0, len(bonds))
	// Convert each builder to its model
	for _, builder := range bonds {
		builtValue := builder.Build().(*models.Bond)
		b.portfolio.Bonds = append(b.portfolio.Bonds, builtValue)
	}
	return b
}

// WithETFs sets the ETFs
func (b *PortfolioBuilder) WithETFs(eTFs []any) *PortfolioBuilder {
	b.portfolio.ETFs = append(b.portfolio.ETFs, eTFs...)
	return b
}

// WithMutualFunds sets the MutualFunds
func (b *PortfolioBuilder) WithMutualFunds(mutualFunds []any) *PortfolioBuilder {
	b.portfolio.MutualFunds = append(b.portfolio.MutualFunds, mutualFunds...)
	return b
}

// WithCryptocurrencies sets the Cryptocurrencies
func (b *PortfolioBuilder) WithCryptocurrencies(cryptocurrencies []any) *PortfolioBuilder {
	b.portfolio.Cryptocurrencies = append(b.portfolio.Cryptocurrencies, cryptocurrencies...)
	return b
}

// AddBond adds a single item to the Bonds slice
func (b *PortfolioBuilder) AddBond(bond *BondBuilder) *PortfolioBuilder {
	// Ensure the slice is initialized
	if b.portfolio.Bonds == nil {
		b.portfolio.Bonds = []*models.Bond{}
	}
	// Handle nested slice element
	builtValue := bond.Build().(*models.Bond)
	b.portfolio.Bonds = append(b.portfolio.Bonds, builtValue)
	return b
}

// WithValidation adds a custom validation function
func (b *PortfolioBuilder) WithValidation(validationFunc func(*models.Portfolio) error) *PortfolioBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Portfolio
func (b *PortfolioBuilder) Build() interface{} {
	return b.portfolio
}

// BuildPtr builds the Portfolio and returns a pointer
func (b *PortfolioBuilder) BuildPtr() *models.Portfolio {
	return b.portfolio
}

// BuildAndValidate builds the Portfolio and validates it
func (b *PortfolioBuilder) BuildAndValidate() (*models.Portfolio, error) {
	portfolio := b.portfolio

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(portfolio); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(portfolio).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return portfolio, err
		}
	}

	return portfolio, nil
}

// MustBuild builds the Portfolio and panics if validation fails
func (b *PortfolioBuilder) MustBuild() *models.Portfolio {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *PortfolioBuilder) Clone() *PortfolioBuilder {
	clonedPortfolio := *b.portfolio
	return &PortfolioBuilder{
		portfolio: &clonedPortfolio,
		validationFuncs: append([]func(*models.Portfolio) error{}, b.validationFuncs...),
	}
}
