package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// CryptoTransactionBuilder builds a CryptoTransaction model
type CryptoTransactionBuilder struct {
	cryptoTransaction *models.CryptoTransaction
	// Custom validation functions
	validationFuncs []func(*models.CryptoTransaction) error
}

// NewCryptoTransactionBuilder creates a new CryptoTransactionBuilder
func NewCryptoTransactionBuilder() *CryptoTransactionBuilder {
	return &CryptoTransactionBuilder{
		cryptoTransaction: &models.CryptoTransaction{
			Date: "",
			Type: "",
			Quantity: 0.0,
			Price: 0.0,
			Fees: 0.0,
			Total: 0.0,
			Exchange: "",
			Network: "",
			TxHash: "",
			WalletFrom: "",
			WalletTo: "",
		},
		validationFuncs: []func(*models.CryptoTransaction) error{},
	}
}

// NewCryptoTransactionBuilderWithDefaults creates a new CryptoTransactionBuilder with sensible defaults
func NewCryptoTransactionBuilderWithDefaults() *CryptoTransactionBuilder {
	builder := NewCryptoTransactionBuilder()
	// Add default values here if needed
	return builder
}
// WithDate sets the Date
func (b *CryptoTransactionBuilder) WithDate(date string) *CryptoTransactionBuilder {
	b.cryptoTransaction.Date = date
	return b
}

// WithType sets the Type
func (b *CryptoTransactionBuilder) WithType(value string) *CryptoTransactionBuilder {
	b.cryptoTransaction.Type = value
	return b
}

// WithQuantity sets the Quantity
func (b *CryptoTransactionBuilder) WithQuantity(quantity float64) *CryptoTransactionBuilder {
	b.cryptoTransaction.Quantity = quantity
	return b
}

// WithPrice sets the Price
func (b *CryptoTransactionBuilder) WithPrice(price float64) *CryptoTransactionBuilder {
	b.cryptoTransaction.Price = price
	return b
}

// WithFees sets the Fees
func (b *CryptoTransactionBuilder) WithFees(fees float64) *CryptoTransactionBuilder {
	b.cryptoTransaction.Fees = fees
	return b
}

// WithTotal sets the Total
func (b *CryptoTransactionBuilder) WithTotal(total float64) *CryptoTransactionBuilder {
	b.cryptoTransaction.Total = total
	return b
}

// WithExchange sets the Exchange
func (b *CryptoTransactionBuilder) WithExchange(exchange string) *CryptoTransactionBuilder {
	b.cryptoTransaction.Exchange = exchange
	return b
}

// WithNetwork sets the Network
func (b *CryptoTransactionBuilder) WithNetwork(network string) *CryptoTransactionBuilder {
	b.cryptoTransaction.Network = network
	return b
}

// WithTxHash sets the TxHash
func (b *CryptoTransactionBuilder) WithTxHash(txHash string) *CryptoTransactionBuilder {
	b.cryptoTransaction.TxHash = txHash
	return b
}

// WithWalletFrom sets the WalletFrom
func (b *CryptoTransactionBuilder) WithWalletFrom(walletFrom string) *CryptoTransactionBuilder {
	b.cryptoTransaction.WalletFrom = walletFrom
	return b
}

// WithWalletTo sets the WalletTo
func (b *CryptoTransactionBuilder) WithWalletTo(walletTo string) *CryptoTransactionBuilder {
	b.cryptoTransaction.WalletTo = walletTo
	return b
}


// WithValidation adds a custom validation function
func (b *CryptoTransactionBuilder) WithValidation(validationFunc func(*models.CryptoTransaction) error) *CryptoTransactionBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the CryptoTransaction
func (b *CryptoTransactionBuilder) Build() interface{} {
	return b.cryptoTransaction
}

// BuildPtr builds the CryptoTransaction and returns a pointer
func (b *CryptoTransactionBuilder) BuildPtr() *models.CryptoTransaction {
	return b.cryptoTransaction
}

// BuildAndValidate builds the CryptoTransaction and validates it
func (b *CryptoTransactionBuilder) BuildAndValidate() (*models.CryptoTransaction, error) {
	cryptoTransaction := b.cryptoTransaction

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(cryptoTransaction); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(cryptoTransaction).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return cryptoTransaction, err
		}
	}

	return cryptoTransaction, nil
}

// MustBuild builds the CryptoTransaction and panics if validation fails
func (b *CryptoTransactionBuilder) MustBuild() *models.CryptoTransaction {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *CryptoTransactionBuilder) Clone() *CryptoTransactionBuilder {
	clonedCryptoTransaction := *b.cryptoTransaction
	return &CryptoTransactionBuilder{
		cryptoTransaction: &clonedCryptoTransaction,
		validationFuncs: append([]func(*models.CryptoTransaction) error{}, b.validationFuncs...),
	}
}
