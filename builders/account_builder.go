package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/models"
	"github.com/adil-faiyaz98/go-builder-kit/pkg/builder"
)

// AccountBuilder builds a Account model
type AccountBuilder struct {
	account *models.Account
	// Custom validation functions
	validationFuncs []func(*models.Account) error
}

// NewAccountBuilder creates a new AccountBuilder
func NewAccountBuilder() *AccountBuilder {
	return &AccountBuilder{
		account: &models.Account{
			ID:             "",
			Type:           "",
			Number:         "",
			Balance:        0.0,
			Currency:       "",
			OpenDate:       "",
			Status:         "",
			Transactions:   []any{},
			InterestRate:   0.0,
			IsJoint:        false,
			CoOwners:       []any{},
			OverdraftLimit: 0.0,
		},
		validationFuncs: []func(*models.Account) error{},
	}
}

// NewAccountBuilderWithDefaults creates a new AccountBuilder with sensible defaults
func NewAccountBuilderWithDefaults() *AccountBuilder {
	builder := NewAccountBuilder()
	// Add default values here if needed
	return builder
}

// WithID sets the ID
func (b *AccountBuilder) WithID(id string) *AccountBuilder {
	if b == nil {
		return b
	}
	b.account.ID = builder.SanitizeString(id)
	return b
}

// WithType sets the Type
func (b *AccountBuilder) WithType(value string) *AccountBuilder {
	if b == nil {
		return b
	}
	b.account.Type = builder.SanitizeString(value)
	return b
}

// WithNumber sets the Number
func (b *AccountBuilder) WithNumber(number string) *AccountBuilder {
	if b == nil {
		return b
	}
	b.account.Number = builder.SanitizeString(number)
	return b
}

// WithBalance sets the Balance
func (b *AccountBuilder) WithBalance(balance float64) *AccountBuilder {
	if b == nil {
		return b
	}
	b.account.Balance = balance
	return b
}

// WithCurrency sets the Currency
func (b *AccountBuilder) WithCurrency(currency string) *AccountBuilder {
	if b == nil {
		return b
	}
	b.account.Currency = builder.SanitizeString(currency)
	return b
}

// WithOpenDate sets the OpenDate
func (b *AccountBuilder) WithOpenDate(openDate string) *AccountBuilder {
	if b == nil {
		return b
	}
	b.account.OpenDate = builder.SanitizeString(openDate)
	return b
}

// WithStatus sets the Status
func (b *AccountBuilder) WithStatus(status string) *AccountBuilder {
	if b == nil {
		return b
	}
	b.account.Status = builder.SanitizeString(status)
	return b
}

// WithTransactions sets the Transactions
func (b *AccountBuilder) WithTransactions(transactions []any) *AccountBuilder {
	if b == nil {
		return b
	}
	b.account.Transactions = append(b.account.Transactions, transactions...)
	return b
}

// WithInterestRate sets the InterestRate
func (b *AccountBuilder) WithInterestRate(interestRate float64) *AccountBuilder {
	if b == nil {
		return b
	}
	b.account.InterestRate = interestRate
	return b
}

// WithIsJoint sets the IsJoint
func (b *AccountBuilder) WithIsJoint(isJoint bool) *AccountBuilder {
	if b == nil {
		return b
	}
	b.account.IsJoint = isJoint
	return b
}

// WithCoOwners sets the CoOwners
func (b *AccountBuilder) WithCoOwners(coOwners []any) *AccountBuilder {
	if b == nil {
		return b
	}
	b.account.CoOwners = append(b.account.CoOwners, coOwners...)
	return b
}

// WithOverdraftLimit sets the OverdraftLimit
func (b *AccountBuilder) WithOverdraftLimit(overdraftLimit float64) *AccountBuilder {
	if b == nil {
		return b
	}
	b.account.OverdraftLimit = overdraftLimit
	return b
}

// WithValidation adds a custom validation function
func (b *AccountBuilder) WithValidation(validationFunc func(*models.Account) error) *AccountBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Account
func (b *AccountBuilder) Build() interface{} {
	return b.account
}

// BuildPtr builds the Account and returns a pointer
func (b *AccountBuilder) BuildPtr() *models.Account {
	return b.account
}

// BuildAndValidate builds the Account and validates it
func (b *AccountBuilder) BuildAndValidate() (*models.Account, error) {
	if b == nil || b.account == nil {
		return nil, fmt.Errorf("builder is not properly initialized")
	}

	account := b.account

	// Run custom validation functions
	for i, validationFunc := range b.validationFuncs {
		if validationFunc == nil {
			continue // Skip nil validators
		}
		if err := validationFunc(account); err != nil {
			return nil, fmt.Errorf("custom validation failed at index %d: %w", i, err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := any(account).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return account, fmt.Errorf("model validation failed: %w", err)
		}
	}

	return account, nil
}

// MustBuild builds the Account and panics if validation fails
func (b *AccountBuilder) MustBuild() *models.Account {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *AccountBuilder) Clone() *AccountBuilder {
	if b == nil || b.account == nil {
		return NewAccountBuilder()
	}

	// Deep copy the struct
	clonedAccount := *b.account

	// Create new builder with cloned data
	clonedBuilder := &AccountBuilder{
		account:         &clonedAccount,
		validationFuncs: make([]func(*models.Account) error, 0, len(b.validationFuncs)),
	}

	// Copy validation functions safely
	if len(b.validationFuncs) > 0 {
		clonedBuilder.validationFuncs = append(clonedBuilder.validationFuncs, b.validationFuncs...)
	}

	return clonedBuilder
}
