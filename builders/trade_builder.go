package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// TradeBuilder builds a Trade model
type TradeBuilder struct {
	trade *models.Trade
	// Custom validation functions
	validationFuncs []func(*models.Trade) error
}

// NewTradeBuilder creates a new TradeBuilder
func NewTradeBuilder() *TradeBuilder {
	return &TradeBuilder{
		trade: &models.Trade{
			Date: "",
			Type: "",
			Units: 0,
			Price: 0.0,
			Fees: 0.0,
			Total: 0.0,
			Exchange: "",
			ExecutedBy: "",
		},
		validationFuncs: []func(*models.Trade) error{},
	}
}

// NewTradeBuilderWithDefaults creates a new TradeBuilder with sensible defaults
func NewTradeBuilderWithDefaults() *TradeBuilder {
	builder := NewTradeBuilder()
	// Add default values here if needed
	return builder
}
// WithDate sets the Date
func (b *TradeBuilder) WithDate(date string) *TradeBuilder {
	b.trade.Date = date
	return b
}

// WithType sets the Type
func (b *TradeBuilder) WithType(value string) *TradeBuilder {
	b.trade.Type = value
	return b
}

// WithUnits sets the Units
func (b *TradeBuilder) WithUnits(units int) *TradeBuilder {
	b.trade.Units = units
	return b
}

// WithPrice sets the Price
func (b *TradeBuilder) WithPrice(price float64) *TradeBuilder {
	b.trade.Price = price
	return b
}

// WithFees sets the Fees
func (b *TradeBuilder) WithFees(fees float64) *TradeBuilder {
	b.trade.Fees = fees
	return b
}

// WithTotal sets the Total
func (b *TradeBuilder) WithTotal(total float64) *TradeBuilder {
	b.trade.Total = total
	return b
}

// WithExchange sets the Exchange
func (b *TradeBuilder) WithExchange(exchange string) *TradeBuilder {
	b.trade.Exchange = exchange
	return b
}

// WithExecutedBy sets the ExecutedBy
func (b *TradeBuilder) WithExecutedBy(executedBy string) *TradeBuilder {
	b.trade.ExecutedBy = executedBy
	return b
}


// WithValidation adds a custom validation function
func (b *TradeBuilder) WithValidation(validationFunc func(*models.Trade) error) *TradeBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Trade
func (b *TradeBuilder) Build() interface{} {
	return b.trade
}

// BuildPtr builds the Trade and returns a pointer
func (b *TradeBuilder) BuildPtr() *models.Trade {
	return b.trade
}

// BuildAndValidate builds the Trade and validates it
func (b *TradeBuilder) BuildAndValidate() (*models.Trade, error) {
	trade := b.trade

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(trade); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(trade).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return trade, err
		}
	}

	return trade, nil
}

// MustBuild builds the Trade and panics if validation fails
func (b *TradeBuilder) MustBuild() *models.Trade {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *TradeBuilder) Clone() *TradeBuilder {
	clonedTrade := *b.trade
	return &TradeBuilder{
		trade: &clonedTrade,
		validationFuncs: append([]func(*models.Trade) error{}, b.validationFuncs...),
	}
}
