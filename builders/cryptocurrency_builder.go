package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// CryptocurrencyBuilder builds a Cryptocurrency model
type CryptocurrencyBuilder struct {
	cryptocurrency *models.Cryptocurrency
	// Custom validation functions
	validationFuncs []func(*models.Cryptocurrency) error
}

// NewCryptocurrencyBuilder creates a new CryptocurrencyBuilder
func NewCryptocurrencyBuilder() *CryptocurrencyBuilder {
	return &CryptocurrencyBuilder{
		cryptocurrency: &models.Cryptocurrency{
			ID: "",
			Symbol: "",
			Name: "",
			Quantity: 0.0,
			PurchasePrice: 0.0,
			CurrentPrice: 0.0,
			Value: 0.0,
			Wallet: "",
			PurchaseDate: "",
			Transactions: []*models.CryptoTransaction{},
		},
		validationFuncs: []func(*models.Cryptocurrency) error{},
	}
}

// NewCryptocurrencyBuilderWithDefaults creates a new CryptocurrencyBuilder with sensible defaults
func NewCryptocurrencyBuilderWithDefaults() *CryptocurrencyBuilder {
	builder := NewCryptocurrencyBuilder()
	// Add default values here if needed
	return builder
}
// WithID sets the ID
func (b *CryptocurrencyBuilder) WithID(iD string) *CryptocurrencyBuilder {
	b.cryptocurrency.ID = iD
	return b
}

// WithSymbol sets the Symbol
func (b *CryptocurrencyBuilder) WithSymbol(symbol string) *CryptocurrencyBuilder {
	b.cryptocurrency.Symbol = symbol
	return b
}

// WithName sets the Name
func (b *CryptocurrencyBuilder) WithName(name string) *CryptocurrencyBuilder {
	b.cryptocurrency.Name = name
	return b
}

// WithQuantity sets the Quantity
func (b *CryptocurrencyBuilder) WithQuantity(quantity float64) *CryptocurrencyBuilder {
	b.cryptocurrency.Quantity = quantity
	return b
}

// WithPurchasePrice sets the PurchasePrice
func (b *CryptocurrencyBuilder) WithPurchasePrice(purchasePrice float64) *CryptocurrencyBuilder {
	b.cryptocurrency.PurchasePrice = purchasePrice
	return b
}

// WithCurrentPrice sets the CurrentPrice
func (b *CryptocurrencyBuilder) WithCurrentPrice(currentPrice float64) *CryptocurrencyBuilder {
	b.cryptocurrency.CurrentPrice = currentPrice
	return b
}

// WithValue sets the Value
func (b *CryptocurrencyBuilder) WithValue(value float64) *CryptocurrencyBuilder {
	b.cryptocurrency.Value = value
	return b
}

// WithWallet sets the Wallet
func (b *CryptocurrencyBuilder) WithWallet(wallet string) *CryptocurrencyBuilder {
	b.cryptocurrency.Wallet = wallet
	return b
}

// WithPurchaseDate sets the PurchaseDate
func (b *CryptocurrencyBuilder) WithPurchaseDate(purchaseDate string) *CryptocurrencyBuilder {
	b.cryptocurrency.PurchaseDate = purchaseDate
	return b
}

// WithTransactions sets the Transactions
func (b *CryptocurrencyBuilder) WithTransactions(transactions *CryptoTransactionBuilder) *CryptocurrencyBuilder {
	// Ensure the slice is initialized
	if b.cryptocurrency.Transactions == nil {
		b.cryptocurrency.Transactions = []*models.CryptoTransaction{}
	}
	// Handle nested slice element
	builtValue := transactions.Build().(*models.CryptoTransaction)
	b.cryptocurrency.Transactions = append(b.cryptocurrency.Transactions, builtValue)
	return b
}


// WithValidation adds a custom validation function
func (b *CryptocurrencyBuilder) WithValidation(validationFunc func(*models.Cryptocurrency) error) *CryptocurrencyBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Cryptocurrency
func (b *CryptocurrencyBuilder) Build() interface{} {
	return b.cryptocurrency
}

// BuildPtr builds the Cryptocurrency and returns a pointer
func (b *CryptocurrencyBuilder) BuildPtr() *models.Cryptocurrency {
	return b.cryptocurrency
}

// BuildAndValidate builds the Cryptocurrency and validates it
func (b *CryptocurrencyBuilder) BuildAndValidate() (*models.Cryptocurrency, error) {
	cryptocurrency := b.cryptocurrency

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(cryptocurrency); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(cryptocurrency).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return cryptocurrency, err
		}
	}

	return cryptocurrency, nil
}

// MustBuild builds the Cryptocurrency and panics if validation fails
func (b *CryptocurrencyBuilder) MustBuild() *models.Cryptocurrency {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *CryptocurrencyBuilder) Clone() *CryptocurrencyBuilder {
	clonedCryptocurrency := *b.cryptocurrency
	return &CryptocurrencyBuilder{
		cryptocurrency: &clonedCryptocurrency,
		validationFuncs: append([]func(*models.Cryptocurrency) error{}, b.validationFuncs...),
	}
}
