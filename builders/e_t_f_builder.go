package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// ETFBuilder builds a ETF model
type ETFBuilder struct {
	eTF *models.ETF
	// Custom validation functions
	validationFuncs []func(*models.ETF) error
}

// NewETFBuilder creates a new ETFBuilder
func NewETFBuilder() *ETFBuilder {
	return &ETFBuilder{
		eTF: &models.ETF{
			ID: "",
			Symbol: "",
			TickerSymbol: "",
			Name: "",
			Provider: "",
			ExpenseRatio: 0.0,
			Category: "",
			Index: "",
			Holdings: map[string]float64{},
			Price: 0.0,
			PurchasePrice: 0.0,
			CurrentPrice: 0.0,
			PurchaseDate: "",
			Quantity: 0,
			Shares: 0.0,
			TotalValue: 0.0,
			YieldPercent: 0.0,
		},
		validationFuncs: []func(*models.ETF) error{},
	}
}

// NewETFBuilderWithDefaults creates a new ETFBuilder with sensible defaults
func NewETFBuilderWithDefaults() *ETFBuilder {
	builder := NewETFBuilder()
	// Add default values here if needed
	return builder
}
// WithID sets the ID
func (b *ETFBuilder) WithID(iD string) *ETFBuilder {
	b.eTF.ID = iD
	return b
}

// WithSymbol sets the Symbol
func (b *ETFBuilder) WithSymbol(symbol string) *ETFBuilder {
	b.eTF.Symbol = symbol
	return b
}

// WithTickerSymbol sets the TickerSymbol
func (b *ETFBuilder) WithTickerSymbol(tickerSymbol string) *ETFBuilder {
	b.eTF.TickerSymbol = tickerSymbol
	return b
}

// WithName sets the Name
func (b *ETFBuilder) WithName(name string) *ETFBuilder {
	b.eTF.Name = name
	return b
}

// WithProvider sets the Provider
func (b *ETFBuilder) WithProvider(provider string) *ETFBuilder {
	b.eTF.Provider = provider
	return b
}

// WithExpenseRatio sets the ExpenseRatio
func (b *ETFBuilder) WithExpenseRatio(expenseRatio float64) *ETFBuilder {
	b.eTF.ExpenseRatio = expenseRatio
	return b
}

// WithCategory sets the Category
func (b *ETFBuilder) WithCategory(category string) *ETFBuilder {
	b.eTF.Category = category
	return b
}

// WithIndex sets the Index
func (b *ETFBuilder) WithIndex(index string) *ETFBuilder {
	b.eTF.Index = index
	return b
}

// WithHoldings sets the Holdings
func (b *ETFBuilder) WithHoldings(key string, val float64) *ETFBuilder {
	if b.eTF.Holdings == nil {
		b.eTF.Holdings = make(map[string]float64)
	}
	b.eTF.Holdings[key] = val
	return b
}

// WithPrice sets the Price
func (b *ETFBuilder) WithPrice(price float64) *ETFBuilder {
	b.eTF.Price = price
	return b
}

// WithPurchasePrice sets the PurchasePrice
func (b *ETFBuilder) WithPurchasePrice(purchasePrice float64) *ETFBuilder {
	b.eTF.PurchasePrice = purchasePrice
	return b
}

// WithCurrentPrice sets the CurrentPrice
func (b *ETFBuilder) WithCurrentPrice(currentPrice float64) *ETFBuilder {
	b.eTF.CurrentPrice = currentPrice
	return b
}

// WithPurchaseDate sets the PurchaseDate
func (b *ETFBuilder) WithPurchaseDate(purchaseDate string) *ETFBuilder {
	b.eTF.PurchaseDate = purchaseDate
	return b
}

// WithQuantity sets the Quantity
func (b *ETFBuilder) WithQuantity(quantity int) *ETFBuilder {
	b.eTF.Quantity = quantity
	return b
}

// WithShares sets the Shares
func (b *ETFBuilder) WithShares(shares float64) *ETFBuilder {
	b.eTF.Shares = shares
	return b
}

// WithTotalValue sets the TotalValue
func (b *ETFBuilder) WithTotalValue(totalValue float64) *ETFBuilder {
	b.eTF.TotalValue = totalValue
	return b
}

// WithYieldPercent sets the YieldPercent
func (b *ETFBuilder) WithYieldPercent(yieldPercent float64) *ETFBuilder {
	b.eTF.YieldPercent = yieldPercent
	return b
}


// WithValidation adds a custom validation function
func (b *ETFBuilder) WithValidation(validationFunc func(*models.ETF) error) *ETFBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the ETF
func (b *ETFBuilder) Build() interface{} {
	return b.eTF
}

// BuildPtr builds the ETF and returns a pointer
func (b *ETFBuilder) BuildPtr() *models.ETF {
	return b.eTF
}

// BuildAndValidate builds the ETF and validates it
func (b *ETFBuilder) BuildAndValidate() (*models.ETF, error) {
	eTF := b.eTF

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(eTF); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(eTF).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return eTF, err
		}
	}

	return eTF, nil
}

// MustBuild builds the ETF and panics if validation fails
func (b *ETFBuilder) MustBuild() *models.ETF {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *ETFBuilder) Clone() *ETFBuilder {
	clonedETF := *b.eTF
	return &ETFBuilder{
		eTF: &clonedETF,
		validationFuncs: append([]func(*models.ETF) error{}, b.validationFuncs...),
	}
}
