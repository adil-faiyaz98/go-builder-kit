package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// TransactionBuilder builds a Transaction model
type TransactionBuilder struct {
	transaction *models.Transaction
	// Custom validation functions
	validationFuncs []func(*models.Transaction) error
}

// NewTransactionBuilder creates a new TransactionBuilder
func NewTransactionBuilder() *TransactionBuilder {
	return &TransactionBuilder{
		transaction: &models.Transaction{
			ID: "",
			Type: "",
			Amount: 0.0,
			Currency: "",
			Date: "",
			Description: "",
			Category: "",
			Status: "",
			FromAccount: nil,
			ToAccount: nil,
			Reference: "",
			Tags: []string{},
		},
		validationFuncs: []func(*models.Transaction) error{},
	}
}

// NewTransactionBuilderWithDefaults creates a new TransactionBuilder with sensible defaults
func NewTransactionBuilderWithDefaults() *TransactionBuilder {
	builder := NewTransactionBuilder()
	// Add default values here if needed
	return builder
}
// WithID sets the ID
func (b *TransactionBuilder) WithID(iD string) *TransactionBuilder {
	b.transaction.ID = iD
	return b
}

// WithType sets the Type
func (b *TransactionBuilder) WithType(value string) *TransactionBuilder {
	b.transaction.Type = value
	return b
}

// WithAmount sets the Amount
func (b *TransactionBuilder) WithAmount(amount float64) *TransactionBuilder {
	b.transaction.Amount = amount
	return b
}

// WithCurrency sets the Currency
func (b *TransactionBuilder) WithCurrency(currency string) *TransactionBuilder {
	b.transaction.Currency = currency
	return b
}

// WithDate sets the Date
func (b *TransactionBuilder) WithDate(date string) *TransactionBuilder {
	b.transaction.Date = date
	return b
}

// WithDescription sets the Description
func (b *TransactionBuilder) WithDescription(description string) *TransactionBuilder {
	b.transaction.Description = description
	return b
}

// WithCategory sets the Category
func (b *TransactionBuilder) WithCategory(category string) *TransactionBuilder {
	b.transaction.Category = category
	return b
}

// WithStatus sets the Status
func (b *TransactionBuilder) WithStatus(status string) *TransactionBuilder {
	b.transaction.Status = status
	return b
}

// WithFromAccount sets the FromAccount
func (b *TransactionBuilder) WithFromAccount(fromAccount *AccountBuilder) *TransactionBuilder {
	// Handle nested pointer
	b.transaction.FromAccount = fromAccount.BuildPtr()
	return b
}

// WithToAccount sets the ToAccount
func (b *TransactionBuilder) WithToAccount(toAccount *AccountBuilder) *TransactionBuilder {
	// Handle nested pointer
	b.transaction.ToAccount = toAccount.BuildPtr()
	return b
}

// WithReference sets the Reference
func (b *TransactionBuilder) WithReference(reference string) *TransactionBuilder {
	b.transaction.Reference = reference
	return b
}

// WithTags sets the Tags
func (b *TransactionBuilder) WithTags(tags string) *TransactionBuilder {
	b.transaction.Tags = append(b.transaction.Tags, tags)
	return b
}


// WithValidation adds a custom validation function
func (b *TransactionBuilder) WithValidation(validationFunc func(*models.Transaction) error) *TransactionBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Transaction
func (b *TransactionBuilder) Build() interface{} {
	return b.transaction
}

// BuildPtr builds the Transaction and returns a pointer
func (b *TransactionBuilder) BuildPtr() *models.Transaction {
	return b.transaction
}

// BuildAndValidate builds the Transaction and validates it
func (b *TransactionBuilder) BuildAndValidate() (*models.Transaction, error) {
	transaction := b.transaction

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(transaction); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(transaction).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return transaction, err
		}
	}

	return transaction, nil
}

// MustBuild builds the Transaction and panics if validation fails
func (b *TransactionBuilder) MustBuild() *models.Transaction {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *TransactionBuilder) Clone() *TransactionBuilder {
	clonedTransaction := *b.transaction
	return &TransactionBuilder{
		transaction: &clonedTransaction,
		validationFuncs: append([]func(*models.Transaction) error{}, b.validationFuncs...),
	}
}
