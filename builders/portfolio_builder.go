package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/models"
	"github.com/adil-faiyaz98/go-builder-kit/pkg/builder"
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
			ID:               "",
			Name:             "",
			Description:      "",
			RiskLevel:        "",
			TotalValue:       0.0,
			Allocation:       map[string]float64{},
			Stocks:           []any{},
			Bonds:            []*models.Bond{},
			ETFs:             []any{},
			MutualFunds:      []any{},
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
	if b == nil {
		return b
	}
	b.portfolio.ID = builder.SanitizeString(id)
	return b
}

// WithName sets the Name
func (b *PortfolioBuilder) WithName(name string) *PortfolioBuilder {
	if b == nil {
		return b
	}
	b.portfolio.Name = builder.SanitizeString(name)
	return b
}

// WithDescription sets the Description
func (b *PortfolioBuilder) WithDescription(description string) *PortfolioBuilder {
	if b == nil {
		return b
	}
	b.portfolio.Description = builder.SanitizeString(description)
	return b
}

// WithRiskLevel sets the RiskLevel
func (b *PortfolioBuilder) WithRiskLevel(riskLevel string) *PortfolioBuilder {
	if b == nil {
		return b
	}
	b.portfolio.RiskLevel = builder.SanitizeString(riskLevel)
	return b
}

// WithTotalValue sets the TotalValue
func (b *PortfolioBuilder) WithTotalValue(totalValue float64) *PortfolioBuilder {
	if b == nil {
		return b
	}
	b.portfolio.TotalValue = totalValue
	return b
}

// WithAllocation sets the Allocation
func (b *PortfolioBuilder) WithAllocation(key string, val float64) *PortfolioBuilder {
	if b == nil {
		return b
	}
	if b.portfolio.Allocation == nil {
		b.portfolio.Allocation = make(map[string]float64)
	}
	b.portfolio.Allocation[key] = val
	return b
}

// WithStocks sets the Stocks
func (b *PortfolioBuilder) WithStocks(stocks []any) *PortfolioBuilder {
	if b == nil {
		return b
	}
	b.portfolio.Stocks = append(b.portfolio.Stocks, stocks...)
	return b
}

// WithBonds sets the Bonds
func (b *PortfolioBuilder) WithBonds(bonds []*BondBuilder) *PortfolioBuilder {
	if b == nil {
		return b
	}
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
	if b == nil {
		return b
	}
	b.portfolio.ETFs = append(b.portfolio.ETFs, eTFs...)
	return b
}

// WithMutualFunds sets the MutualFunds
func (b *PortfolioBuilder) WithMutualFunds(mutualFunds []any) *PortfolioBuilder {
	if b == nil {
		return b
	}
	b.portfolio.MutualFunds = append(b.portfolio.MutualFunds, mutualFunds...)
	return b
}

// WithCryptocurrencies sets the Cryptocurrencies
func (b *PortfolioBuilder) WithCryptocurrencies(cryptocurrencies []any) *PortfolioBuilder {
	if b == nil {
		return b
	}
	b.portfolio.Cryptocurrencies = append(b.portfolio.Cryptocurrencies, cryptocurrencies...)
	return b
}

// AddBond adds a single item to the Bonds slice
func (b *PortfolioBuilder) AddBond(bond *BondBuilder) *PortfolioBuilder {
	if b == nil {
		return b
	}
	// Ensure the slice is initialized with capacity
	if b.portfolio.Bonds == nil {
		b.portfolio.Bonds = make([]*models.Bond, 0, 4) // Pre-allocate capacity
	}
	// Handle nested slice element
	if bond != nil {
		builtValue := bond.Build().(*models.Bond)
		b.portfolio.Bonds = append(b.portfolio.Bonds, builtValue)
	}
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
	if b == nil || b.portfolio == nil {
		return nil, fmt.Errorf("builder is not properly initialized")
	}

	portfolio := b.portfolio

	// Run custom validation functions
	for i, validationFunc := range b.validationFuncs {
		if validationFunc == nil {
			continue // Skip nil validators
		}
		if err := validationFunc(portfolio); err != nil {
			return nil, fmt.Errorf("custom validation failed at index %d: %w", i, err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := any(portfolio).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return portfolio, fmt.Errorf("model validation failed: %w", err)
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
	if b == nil || b.portfolio == nil {
		return NewPortfolioBuilder()
	}

	// Deep copy the struct
	clonedPortfolio := *b.portfolio

	// Create new builder with cloned data
	clonedBuilder := &PortfolioBuilder{
		portfolio:       &clonedPortfolio,
		validationFuncs: make([]func(*models.Portfolio) error, 0, len(b.validationFuncs)),
	}

	// Copy validation functions safely
	if len(b.validationFuncs) > 0 {
		clonedBuilder.validationFuncs = append(clonedBuilder.validationFuncs, b.validationFuncs...)
	}

	return clonedBuilder
}
