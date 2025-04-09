package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
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
			Name: "",
			BranchCode: "",
			Address: nil,
			Accounts: []*models.Account{},
			Stocks: []*models.Stock{},
			Loans: []*models.Loan{},
			Investments: []*models.Investment{},
			Advisor: any(0),
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
	b.bank.Name = name
	return b
}

// WithBranchCode sets the BranchCode
func (b *BankBuilder) WithBranchCode(branchCode string) *BankBuilder {
	b.bank.BranchCode = branchCode
	return b
}

// WithAddress sets the Address
func (b *BankBuilder) WithAddress(address *AddressBuilder) *BankBuilder {
	// Handle nested pointer
	b.bank.Address = address.BuildPtr()
	return b
}

// WithAccounts sets the Accounts
func (b *BankBuilder) WithAccounts(accounts *AccountBuilder) *BankBuilder {
	// Ensure the slice is initialized
	if b.bank.Accounts == nil {
		b.bank.Accounts = []*models.Account{}
	}
	// Handle nested slice element
	builtValue := accounts.Build().(*models.Account)
	b.bank.Accounts = append(b.bank.Accounts, builtValue)
	return b
}

// WithStocks sets the Stocks
func (b *BankBuilder) WithStocks(stocks *StockBuilder) *BankBuilder {
	// Ensure the slice is initialized
	if b.bank.Stocks == nil {
		b.bank.Stocks = []*models.Stock{}
	}
	// Handle nested slice element
	builtValue := stocks.Build().(*models.Stock)
	b.bank.Stocks = append(b.bank.Stocks, builtValue)
	return b
}

// WithLoans sets the Loans
func (b *BankBuilder) WithLoans(loans *LoanBuilder) *BankBuilder {
	// Ensure the slice is initialized
	if b.bank.Loans == nil {
		b.bank.Loans = []*models.Loan{}
	}
	// Handle nested slice element
	builtValue := loans.Build().(*models.Loan)
	b.bank.Loans = append(b.bank.Loans, builtValue)
	return b
}

// WithInvestments sets the Investments
func (b *BankBuilder) WithInvestments(investments *InvestmentBuilder) *BankBuilder {
	// Ensure the slice is initialized
	if b.bank.Investments == nil {
		b.bank.Investments = []*models.Investment{}
	}
	// Handle nested slice element
	builtValue := investments.Build().(*models.Investment)
	b.bank.Investments = append(b.bank.Investments, builtValue)
	return b
}

// WithAdvisor sets the Advisor
func (b *BankBuilder) WithAdvisor(advisor any) *BankBuilder {
	b.bank.Advisor = advisor
	return b
}

// WithRelationshipStartDate sets the RelationshipStartDate
func (b *BankBuilder) WithRelationshipStartDate(relationshipStartDate string) *BankBuilder {
	b.bank.RelationshipStartDate = relationshipStartDate
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
	bank := b.bank

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(bank); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(bank).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return bank, err
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
	clonedBank := *b.bank
	return &BankBuilder{
		bank: &clonedBank,
		validationFuncs: append([]func(*models.Bank) error{}, b.validationFuncs...),
	}
}
