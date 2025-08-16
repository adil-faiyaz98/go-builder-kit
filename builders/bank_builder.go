package builders

import (
	"fmt"

	"github.com/adil-faiyaz98/go-builder-kit/models"
	"github.com/adil-faiyaz98/go-builder-kit/pkg/builder"
)

// BankBuilder builds a Bank model
type BankBuilder struct {
	bank *models.Bank
	// Custom validation functions
	validationFuncs []func(*models.Bank) error
}

// NewBankBuilder creates a new BankBuilder
func NewBankBuilder() *BankBuilder {
	return &BankBuilder{
		bank: &models.Bank{
			Name:                  "",
			BranchCode:            "",
			Address:               nil,
			Accounts:              []*models.Account{},
			Stocks:                []any{},
			Loans:                 []any{},
			Investments:           []any{},
			Advisor:               nil,
			RelationshipStartDate: "",
		},
		validationFuncs: []func(*models.Bank) error{},
	}
}

// NewBankBuilderWithDefaults creates a new BankBuilder with sensible defaults
func NewBankBuilderWithDefaults() *BankBuilder {
	builder := NewBankBuilder()
	// Add default values here if needed
	return builder
}

// WithName sets the Name
func (b *BankBuilder) WithName(name string) *BankBuilder {
	if b == nil {
		return b
	}
	b.bank.Name = builder.SanitizeString(name)
	return b
}

// WithBranchCode sets the BranchCode
func (b *BankBuilder) WithBranchCode(branchCode string) *BankBuilder {
	if b == nil {
		return b
	}
	b.bank.BranchCode = builder.SanitizeString(branchCode)
	return b
}

// WithAddress sets the Address
func (b *BankBuilder) WithAddress(address *AddressBuilder) *BankBuilder {
	if b == nil {
		return b
	}
	// Handle nested pointer
	b.bank.Address = address.BuildPtr()
	return b
}

// WithAccounts sets the Accounts
func (b *BankBuilder) WithAccounts(accounts []*AccountBuilder) *BankBuilder {
	if b == nil {
		return b
	}
	// Ensure the slice is initialized
	if b.bank.Accounts == nil {
		b.bank.Accounts = []*models.Account{}
	}
	// Handle nested slice elements
	// Initialize the slice
	b.bank.Accounts = make([]*models.Account, 0, len(accounts))
	// Convert each builder to its model
	for _, builder := range accounts {
		builtValue := builder.Build().(*models.Account)
		b.bank.Accounts = append(b.bank.Accounts, builtValue)
	}
	return b
}

// WithStocks sets the Stocks
func (b *BankBuilder) WithStocks(stocks []any) *BankBuilder {
	if b == nil {
		return b
	}
	b.bank.Stocks = append(b.bank.Stocks, stocks...)
	return b
}

// WithLoans sets the Loans
func (b *BankBuilder) WithLoans(loans []any) *BankBuilder {
	if b == nil {
		return b
	}
	b.bank.Loans = append(b.bank.Loans, loans...)
	return b
}

// WithInvestments sets the Investments
func (b *BankBuilder) WithInvestments(investments []any) *BankBuilder {
	if b == nil {
		return b
	}
	b.bank.Investments = append(b.bank.Investments, investments...)
	return b
}

// WithAdvisor sets the Advisor
func (b *BankBuilder) WithAdvisor(advisor any) *BankBuilder {
	if b == nil {
		return b
	}
	b.bank.Advisor = advisor
	return b
}

// WithRelationshipStartDate sets the RelationshipStartDate
func (b *BankBuilder) WithRelationshipStartDate(relationshipStartDate string) *BankBuilder {
	if b == nil {
		return b
	}
	b.bank.RelationshipStartDate = builder.SanitizeString(relationshipStartDate)
	return b
}

// AddAccount adds a single item to the Accounts slice
func (b *BankBuilder) AddAccount(account *AccountBuilder) *BankBuilder {
	if b == nil {
		return b
	}
	// Ensure the slice is initialized with capacity
	if b.bank.Accounts == nil {
		b.bank.Accounts = make([]*models.Account, 0, 4) // Pre-allocate capacity
	}
	// Handle nested slice element
	if account != nil {
		builtValue := account.Build().(*models.Account)
		b.bank.Accounts = append(b.bank.Accounts, builtValue)
	}
	return b
}

// WithValidation adds a custom validation function
func (b *BankBuilder) WithValidation(validationFunc func(*models.Bank) error) *BankBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Bank
func (b *BankBuilder) Build() interface{} {
	return b.bank
}

// BuildPtr builds the Bank and returns a pointer
func (b *BankBuilder) BuildPtr() *models.Bank {
	return b.bank
}

// BuildAndValidate builds the Bank and validates it
func (b *BankBuilder) BuildAndValidate() (*models.Bank, error) {
	if b == nil || b.bank == nil {
		return nil, fmt.Errorf("builder is not properly initialized")
	}

	bank := b.bank

	// Run custom validation functions
	for i, validationFunc := range b.validationFuncs {
		if validationFunc == nil {
			continue // Skip nil validators
		}
		if err := validationFunc(bank); err != nil {
			return nil, fmt.Errorf("custom validation failed at index %d: %w", i, err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := any(bank).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return bank, fmt.Errorf("model validation failed: %w", err)
		}
	}

	return bank, nil
}

// MustBuild builds the Bank and panics if validation fails
func (b *BankBuilder) MustBuild() *models.Bank {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *BankBuilder) Clone() *BankBuilder {
	if b == nil || b.bank == nil {
		return NewBankBuilder()
	}

	// Deep copy the struct
	clonedBank := *b.bank

	// Create new builder with cloned data
	clonedBuilder := &BankBuilder{
		bank:            &clonedBank,
		validationFuncs: make([]func(*models.Bank) error, 0, len(b.validationFuncs)),
	}

	// Copy validation functions safely
	if len(b.validationFuncs) > 0 {
		clonedBuilder.validationFuncs = append(clonedBuilder.validationFuncs, b.validationFuncs...)
	}

	return clonedBuilder
}
