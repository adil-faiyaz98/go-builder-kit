package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// MutualFundBuilder builds a MutualFund model
type MutualFundBuilder struct {
	mutualFund *models.MutualFund
	// Custom validation functions
	validationFuncs []func(*models.MutualFund) error
}

// NewMutualFundBuilder creates a new MutualFundBuilder
func NewMutualFundBuilder() *MutualFundBuilder {
	return &MutualFundBuilder{
		mutualFund: &models.MutualFund{
			ID: "",
			Name: "",
			Symbol: "",
			TickerSymbol: "",
			Category: "",
			FundManager: "",
			FundFamily: "",
			ExpenseRatio: 0.0,
			NAV: 0.0,
			AUM: 0.0,
			YTDReturn: 0.0,
			OneYearReturn: 0.0,
			ThreeYearReturn: 0.0,
			FiveYearReturn: 0.0,
			Risk: "",
			Strategy: "",
			Shares: 0.0,
			PurchaseDate: "",
			Holdings: []*models.FundHolding{},
		},
		validationFuncs: []func(*models.MutualFund) error{},
	}
}

// NewMutualFundBuilderWithDefaults creates a new MutualFundBuilder with sensible defaults
func NewMutualFundBuilderWithDefaults() *MutualFundBuilder {
	builder := NewMutualFundBuilder()
	// Add default values here if needed
	return builder
}
// WithID sets the ID
func (b *MutualFundBuilder) WithID(iD string) *MutualFundBuilder {
	b.mutualFund.ID = iD
	return b
}

// WithName sets the Name
func (b *MutualFundBuilder) WithName(name string) *MutualFundBuilder {
	b.mutualFund.Name = name
	return b
}

// WithSymbol sets the Symbol
func (b *MutualFundBuilder) WithSymbol(symbol string) *MutualFundBuilder {
	b.mutualFund.Symbol = symbol
	return b
}

// WithTickerSymbol sets the TickerSymbol
func (b *MutualFundBuilder) WithTickerSymbol(tickerSymbol string) *MutualFundBuilder {
	b.mutualFund.TickerSymbol = tickerSymbol
	return b
}

// WithCategory sets the Category
func (b *MutualFundBuilder) WithCategory(category string) *MutualFundBuilder {
	b.mutualFund.Category = category
	return b
}

// WithFundManager sets the FundManager
func (b *MutualFundBuilder) WithFundManager(fundManager string) *MutualFundBuilder {
	b.mutualFund.FundManager = fundManager
	return b
}

// WithFundFamily sets the FundFamily
func (b *MutualFundBuilder) WithFundFamily(fundFamily string) *MutualFundBuilder {
	b.mutualFund.FundFamily = fundFamily
	return b
}

// WithExpenseRatio sets the ExpenseRatio
func (b *MutualFundBuilder) WithExpenseRatio(expenseRatio float64) *MutualFundBuilder {
	b.mutualFund.ExpenseRatio = expenseRatio
	return b
}

// WithNAV sets the NAV
func (b *MutualFundBuilder) WithNAV(nAV float64) *MutualFundBuilder {
	b.mutualFund.NAV = nAV
	return b
}

// WithAUM sets the AUM
func (b *MutualFundBuilder) WithAUM(aUM float64) *MutualFundBuilder {
	b.mutualFund.AUM = aUM
	return b
}

// WithYTDReturn sets the YTDReturn
func (b *MutualFundBuilder) WithYTDReturn(yTDReturn float64) *MutualFundBuilder {
	b.mutualFund.YTDReturn = yTDReturn
	return b
}

// WithOneYearReturn sets the OneYearReturn
func (b *MutualFundBuilder) WithOneYearReturn(oneYearReturn float64) *MutualFundBuilder {
	b.mutualFund.OneYearReturn = oneYearReturn
	return b
}

// WithThreeYearReturn sets the ThreeYearReturn
func (b *MutualFundBuilder) WithThreeYearReturn(threeYearReturn float64) *MutualFundBuilder {
	b.mutualFund.ThreeYearReturn = threeYearReturn
	return b
}

// WithFiveYearReturn sets the FiveYearReturn
func (b *MutualFundBuilder) WithFiveYearReturn(fiveYearReturn float64) *MutualFundBuilder {
	b.mutualFund.FiveYearReturn = fiveYearReturn
	return b
}

// WithRisk sets the Risk
func (b *MutualFundBuilder) WithRisk(risk string) *MutualFundBuilder {
	b.mutualFund.Risk = risk
	return b
}

// WithStrategy sets the Strategy
func (b *MutualFundBuilder) WithStrategy(strategy string) *MutualFundBuilder {
	b.mutualFund.Strategy = strategy
	return b
}

// WithShares sets the Shares
func (b *MutualFundBuilder) WithShares(shares float64) *MutualFundBuilder {
	b.mutualFund.Shares = shares
	return b
}

// WithPurchaseDate sets the PurchaseDate
func (b *MutualFundBuilder) WithPurchaseDate(purchaseDate string) *MutualFundBuilder {
	b.mutualFund.PurchaseDate = purchaseDate
	return b
}

// WithHoldings sets the Holdings
func (b *MutualFundBuilder) WithHoldings(holdings *FundHoldingBuilder) *MutualFundBuilder {
	// Ensure the slice is initialized
	if b.mutualFund.Holdings == nil {
		b.mutualFund.Holdings = []*models.FundHolding{}
	}
	// Handle nested slice element
	builtValue := holdings.Build().(*models.FundHolding)
	b.mutualFund.Holdings = append(b.mutualFund.Holdings, builtValue)
	return b
}


// WithValidation adds a custom validation function
func (b *MutualFundBuilder) WithValidation(validationFunc func(*models.MutualFund) error) *MutualFundBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the MutualFund
func (b *MutualFundBuilder) Build() interface{} {
	return b.mutualFund
}

// BuildPtr builds the MutualFund and returns a pointer
func (b *MutualFundBuilder) BuildPtr() *models.MutualFund {
	return b.mutualFund
}

// BuildAndValidate builds the MutualFund and validates it
func (b *MutualFundBuilder) BuildAndValidate() (*models.MutualFund, error) {
	mutualFund := b.mutualFund

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(mutualFund); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(mutualFund).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return mutualFund, err
		}
	}

	return mutualFund, nil
}

// MustBuild builds the MutualFund and panics if validation fails
func (b *MutualFundBuilder) MustBuild() *models.MutualFund {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *MutualFundBuilder) Clone() *MutualFundBuilder {
	clonedMutualFund := *b.mutualFund
	return &MutualFundBuilder{
		mutualFund: &clonedMutualFund,
		validationFuncs: append([]func(*models.MutualFund) error{}, b.validationFuncs...),
	}
}
