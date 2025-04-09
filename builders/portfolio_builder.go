package builders

import (
	"fmt"

	"github.com/adil-faiyaz98/go-builder-kit/models"
)

// PortfolioBuilder builds a Portfolio model
type PortfolioBuilder struct {
	portfolio       *models.Portfolio
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

// WithAllocation adds an allocation to the Allocation map
func (b *PortfolioBuilder) WithAllocation(assetClass string, percentage float64) *PortfolioBuilder {
	if b.portfolio.Allocation == nil {
		b.portfolio.Allocation = map[string]float64{}
	}
	b.portfolio.Allocation[assetClass] = percentage
	return b
}

// WithStock adds a stock to the Stocks slice
func (b *PortfolioBuilder) WithStock(stock any) *PortfolioBuilder {
	b.portfolio.Stocks = append(b.portfolio.Stocks, stock)
	return b
}

// WithBond adds a bond to the Bonds slice
func (b *PortfolioBuilder) WithBond(bond *BondBuilder) *PortfolioBuilder {
	builtValue := bond.Build().(*models.Bond)
	b.portfolio.Bonds = append(b.portfolio.Bonds, builtValue)
	return b
}

// WithETF adds an ETF to the ETFs slice
func (b *PortfolioBuilder) WithETF(etf any) *PortfolioBuilder {
	b.portfolio.ETFs = append(b.portfolio.ETFs, etf)
	return b
}

// WithMutualFund adds a mutual fund to the MutualFunds slice
func (b *PortfolioBuilder) WithMutualFund(mutualFund any) *PortfolioBuilder {
	b.portfolio.MutualFunds = append(b.portfolio.MutualFunds, mutualFund)
	return b
}

// WithCryptocurrency adds a cryptocurrency to the Cryptocurrencies slice
func (b *PortfolioBuilder) WithCryptocurrency(cryptocurrency any) *PortfolioBuilder {
	b.portfolio.Cryptocurrencies = append(b.portfolio.Cryptocurrencies, cryptocurrency)
	return b
}

// WithValidation adds a custom validation function
func (b *PortfolioBuilder) WithValidation(validationFunc func(*models.Portfolio) error) *PortfolioBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Portfolio
func (b *PortfolioBuilder) Build() any {
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
			return portfolio, err
		}
	}

	// Run model's Validate method
	if err := portfolio.Validate(); err != nil {
		return portfolio, err
	}

	return portfolio, nil
}

// MustBuild builds the Portfolio and panics if validation fails
func (b *PortfolioBuilder) MustBuild() *models.Portfolio {
	portfolio, err := b.BuildAndValidate()
	if err != nil {
		panic(fmt.Sprintf("Portfolio validation failed: %s", err.Error()))
	}
	return portfolio
}

// Clone creates a deep copy of the PortfolioBuilder
func (b *PortfolioBuilder) Clone() *PortfolioBuilder {
	clonedPortfolio := *b.portfolio
	return &PortfolioBuilder{
		portfolio:       &clonedPortfolio,
		validationFuncs: append([]func(*models.Portfolio) error{}, b.validationFuncs...),
	}
}
