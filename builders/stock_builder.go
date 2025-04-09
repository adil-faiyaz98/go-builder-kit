package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// StockBuilder builds a Stock model
type StockBuilder struct {
	stock *models.Stock
	// Custom validation functions
	validationFuncs []func(*models.Stock) error
}

// NewStockBuilder creates a new StockBuilder
func NewStockBuilder() *StockBuilder {
	return &StockBuilder{
		stock: &models.Stock{
			Ticker: "",
			CompanyName: "",
			Units: 0,
			PurchasePrice: 0.0,
			CurrentPrice: 0.0,
			PurchaseDate: "",
			Sector: "",
			DividendYield: 0.0,
			Trades: []*models.Trade{},
		},
		validationFuncs: []func(*models.Stock) error{},
	}
}

// NewStockBuilderWithDefaults creates a new StockBuilder with sensible defaults
func NewStockBuilderWithDefaults() *StockBuilder {
	builder := NewStockBuilder()
	// Add default values here if needed
	return builder
}
// WithTicker sets the Ticker
func (b *StockBuilder) WithTicker(ticker string) *StockBuilder {
	b.stock.Ticker = ticker
	return b
}

// WithCompanyName sets the CompanyName
func (b *StockBuilder) WithCompanyName(companyName string) *StockBuilder {
	b.stock.CompanyName = companyName
	return b
}

// WithUnits sets the Units
func (b *StockBuilder) WithUnits(units int) *StockBuilder {
	b.stock.Units = units
	return b
}

// WithPurchasePrice sets the PurchasePrice
func (b *StockBuilder) WithPurchasePrice(purchasePrice float64) *StockBuilder {
	b.stock.PurchasePrice = purchasePrice
	return b
}

// WithCurrentPrice sets the CurrentPrice
func (b *StockBuilder) WithCurrentPrice(currentPrice float64) *StockBuilder {
	b.stock.CurrentPrice = currentPrice
	return b
}

// WithPurchaseDate sets the PurchaseDate
func (b *StockBuilder) WithPurchaseDate(purchaseDate string) *StockBuilder {
	b.stock.PurchaseDate = purchaseDate
	return b
}

// WithSector sets the Sector
func (b *StockBuilder) WithSector(sector string) *StockBuilder {
	b.stock.Sector = sector
	return b
}

// WithDividendYield sets the DividendYield
func (b *StockBuilder) WithDividendYield(dividendYield float64) *StockBuilder {
	b.stock.DividendYield = dividendYield
	return b
}

// WithTrades sets the Trades
func (b *StockBuilder) WithTrades(trades *TradeBuilder) *StockBuilder {
	// Ensure the slice is initialized
	if b.stock.Trades == nil {
		b.stock.Trades = []*models.Trade{}
	}
	// Handle nested slice element
	builtValue := trades.Build().(*models.Trade)
	b.stock.Trades = append(b.stock.Trades, builtValue)
	return b
}


// WithValidation adds a custom validation function
func (b *StockBuilder) WithValidation(validationFunc func(*models.Stock) error) *StockBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Stock
func (b *StockBuilder) Build() interface{} {
	return b.stock
}

// BuildPtr builds the Stock and returns a pointer
func (b *StockBuilder) BuildPtr() *models.Stock {
	return b.stock
}

// BuildAndValidate builds the Stock and validates it
func (b *StockBuilder) BuildAndValidate() (*models.Stock, error) {
	stock := b.stock

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(stock); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(stock).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return stock, err
		}
	}

	return stock, nil
}

// MustBuild builds the Stock and panics if validation fails
func (b *StockBuilder) MustBuild() *models.Stock {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *StockBuilder) Clone() *StockBuilder {
	clonedStock := *b.stock
	return &StockBuilder{
		stock: &clonedStock,
		validationFuncs: append([]func(*models.Stock) error{}, b.validationFuncs...),
	}
}
