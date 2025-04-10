package builders

import (
	"fmt"

	"github.com/adil-faiyaz98/go-builder-kit/models"
)

// BankBuilder builds a Bank model
type BankBuilder struct {
	bank            *models.Bank
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
	b.bank.Address = address.BuildPtr()
	return b
}

// WithAccounts sets the Accounts
func (b *BankBuilder) WithAccounts(accounts []*models.Account) *BankBuilder {
	b.bank.Accounts = accounts
	return b
}

// AddAccount adds an account to the Accounts slice
func (b *BankBuilder) AddAccount(account *AccountBuilder) *BankBuilder {
	b.bank.Accounts = append(b.bank.Accounts, account.BuildPtr())
	return b
}

// WithStocks sets the Stocks
func (b *BankBuilder) WithStocks(stocks []any) *BankBuilder {
	b.bank.Stocks = stocks
	return b
}

// AddStock adds a stock to the Stocks slice
func (b *BankBuilder) AddStock(stock any) *BankBuilder {
	b.bank.Stocks = append(b.bank.Stocks, stock)
	return b
}

// WithLoans sets the Loans
func (b *BankBuilder) WithLoans(loans []any) *BankBuilder {
	b.bank.Loans = loans
	return b
}

// AddLoan adds a loan to the Loans slice
func (b *BankBuilder) AddLoan(loan any) *BankBuilder {
	b.bank.Loans = append(b.bank.Loans, loan)
	return b
}

// WithInvestments sets the Investments
func (b *BankBuilder) WithInvestments(investments []any) *BankBuilder {
	b.bank.Investments = investments
	return b
}

// AddInvestment adds an investment to the Investments slice
func (b *BankBuilder) AddInvestment(investment any) *BankBuilder {
	b.bank.Investments = append(b.bank.Investments, investment)
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

	// Run model's Validate method
	if err := bank.Validate(); err != nil {
		return bank, err
	}

	return bank, nil
}

// MustBuild builds the Bank and panics if validation fails
func (b *BankBuilder) MustBuild() *models.Bank {
	bank, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return bank
}

// Clone creates a deep copy of the builder
func (b *BankBuilder) Clone() *BankBuilder {
	clonedBank := *b.bank
	return &BankBuilder{
		bank:            &clonedBank,
		validationFuncs: append([]func(*models.Bank) error{}, b.validationFuncs...),
	}
}
