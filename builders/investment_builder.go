package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// InvestmentBuilder builds a Investment model
type InvestmentBuilder struct {
	investment *models.Investment
	// Custom validation functions
	validationFuncs []func(*models.Investment) error
}

// NewInvestmentBuilder creates a new InvestmentBuilder
func NewInvestmentBuilder() *InvestmentBuilder {
	return &InvestmentBuilder{
		investment: &models.Investment{
			ID: "",
			Name: "",
			Type: "",
			Value: 0.0,
			StartDate: "",
			EndDate: "",
			Risk: "",
			Portfolio: nil,
			Performance: []*models.PerformanceRecord{},
		},
		validationFuncs: []func(*models.Investment) error{},
	}
}

// NewInvestmentBuilderWithDefaults creates a new InvestmentBuilder with sensible defaults
func NewInvestmentBuilderWithDefaults() *InvestmentBuilder {
	builder := NewInvestmentBuilder()
	// Add default values here if needed
	return builder
}
// WithID sets the ID
func (b *InvestmentBuilder) WithID(iD string) *InvestmentBuilder {
	b.investment.ID = iD
	return b
}

// WithName sets the Name
func (b *InvestmentBuilder) WithName(name string) *InvestmentBuilder {
	b.investment.Name = name
	return b
}

// WithType sets the Type
func (b *InvestmentBuilder) WithType(value string) *InvestmentBuilder {
	b.investment.Type = value
	return b
}

// WithValue sets the Value
func (b *InvestmentBuilder) WithValue(value float64) *InvestmentBuilder {
	b.investment.Value = value
	return b
}

// WithStartDate sets the StartDate
func (b *InvestmentBuilder) WithStartDate(startDate string) *InvestmentBuilder {
	b.investment.StartDate = startDate
	return b
}

// WithEndDate sets the EndDate
func (b *InvestmentBuilder) WithEndDate(endDate string) *InvestmentBuilder {
	b.investment.EndDate = endDate
	return b
}

// WithRisk sets the Risk
func (b *InvestmentBuilder) WithRisk(risk string) *InvestmentBuilder {
	b.investment.Risk = risk
	return b
}

// WithPortfolio sets the Portfolio
func (b *InvestmentBuilder) WithPortfolio(portfolio *PortfolioBuilder) *InvestmentBuilder {
	// Handle nested pointer
	b.investment.Portfolio = portfolio.BuildPtr()
	return b
}

// WithPerformance sets the Performance
func (b *InvestmentBuilder) WithPerformance(performance *PerformanceRecordBuilder) *InvestmentBuilder {
	// Ensure the slice is initialized
	if b.investment.Performance == nil {
		b.investment.Performance = []*models.PerformanceRecord{}
	}
	// Handle nested slice element
	builtValue := performance.Build().(*models.PerformanceRecord)
	b.investment.Performance = append(b.investment.Performance, builtValue)
	return b
}


// WithValidation adds a custom validation function
func (b *InvestmentBuilder) WithValidation(validationFunc func(*models.Investment) error) *InvestmentBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Investment
func (b *InvestmentBuilder) Build() interface{} {
	return b.investment
}

// BuildPtr builds the Investment and returns a pointer
func (b *InvestmentBuilder) BuildPtr() *models.Investment {
	return b.investment
}

// BuildAndValidate builds the Investment and validates it
func (b *InvestmentBuilder) BuildAndValidate() (*models.Investment, error) {
	investment := b.investment

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(investment); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(investment).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return investment, err
		}
	}

	return investment, nil
}

// MustBuild builds the Investment and panics if validation fails
func (b *InvestmentBuilder) MustBuild() *models.Investment {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *InvestmentBuilder) Clone() *InvestmentBuilder {
	clonedInvestment := *b.investment
	return &InvestmentBuilder{
		investment: &clonedInvestment,
		validationFuncs: append([]func(*models.Investment) error{}, b.validationFuncs...),
	}
}
