package builders

import (
	"fmt"

	"github.com/adil-faiyaz98/go-builder-kit/models"
)

// AccountBuilder builds an Account model
type AccountBuilder struct {
	account         *models.Account
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

// WithID sets the ID
func (b *AccountBuilder) WithID(id string) *AccountBuilder {
	b.account.ID = id
	return b
}

// WithType sets the Type
func (b *AccountBuilder) WithType(accountType string) *AccountBuilder {
	b.account.Type = accountType
	return b
}

// WithNumber sets the Number
func (b *AccountBuilder) WithNumber(number string) *AccountBuilder {
	b.account.Number = number
	return b
}

// WithBalance sets the Balance
func (b *AccountBuilder) WithBalance(balance float64) *AccountBuilder {
	b.account.Balance = balance
	return b
}

// WithCurrency sets the Currency
func (b *AccountBuilder) WithCurrency(currency string) *AccountBuilder {
	b.account.Currency = currency
	return b
}

// WithOpenDate sets the OpenDate
func (b *AccountBuilder) WithOpenDate(openDate string) *AccountBuilder {
	b.account.OpenDate = openDate
	return b
}

// WithStatus sets the Status
func (b *AccountBuilder) WithStatus(status string) *AccountBuilder {
	b.account.Status = status
	return b
}

// WithTransactions sets the Transactions
func (b *AccountBuilder) WithTransactions(transactions []any) *AccountBuilder {
	b.account.Transactions = transactions
	return b
}

// AddTransaction adds a transaction to the Transactions slice
func (b *AccountBuilder) AddTransaction(transaction any) *AccountBuilder {
	b.account.Transactions = append(b.account.Transactions, transaction)
	return b
}

// WithInterestRate sets the InterestRate
func (b *AccountBuilder) WithInterestRate(interestRate float64) *AccountBuilder {
	b.account.InterestRate = interestRate
	return b
}

// WithIsJoint sets the IsJoint flag
func (b *AccountBuilder) WithIsJoint(isJoint bool) *AccountBuilder {
	b.account.IsJoint = isJoint
	return b
}

// WithCoOwners sets the CoOwners
func (b *AccountBuilder) WithCoOwners(coOwners []any) *AccountBuilder {
	b.account.CoOwners = coOwners
	return b
}

// AddCoOwner adds a co-owner to the CoOwners slice
func (b *AccountBuilder) AddCoOwner(coOwner any) *AccountBuilder {
	b.account.CoOwners = append(b.account.CoOwners, coOwner)
	return b
}

// WithOverdraftLimit sets the OverdraftLimit
func (b *AccountBuilder) WithOverdraftLimit(overdraftLimit float64) *AccountBuilder {
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
	account := b.account

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(account); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method
	if err := account.Validate(); err != nil {
		return account, err
	}

	return account, nil
}

// MustBuild builds the Account and panics if validation fails
func (b *AccountBuilder) MustBuild() *models.Account {
	account, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return account
}

// Clone creates a deep copy of the builder
func (b *AccountBuilder) Clone() *AccountBuilder {
	clonedAccount := *b.account
	return &AccountBuilder{
		account:         &clonedAccount,
		validationFuncs: append([]func(*models.Account) error{}, b.validationFuncs...),
	}
}
