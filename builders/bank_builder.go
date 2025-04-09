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
	builtValue := address.Build().(*models.Address)
	b.bank.Address = builtValue
	return b
}

// WithAccount adds an account to the Accounts slice
func (b *BankBuilder) WithAccount(account *AccountBuilder) *BankBuilder {
	builtValue := account.Build().(*models.Account)
	b.bank.Accounts = append(b.bank.Accounts, builtValue)
	return b
}

// WithStock adds a stock to the Stocks slice
func (b *BankBuilder) WithStock(stock any) *BankBuilder {
	b.bank.Stocks = append(b.bank.Stocks, stock)
	return b
}

// WithLoan adds a loan to the Loans slice
func (b *BankBuilder) WithLoan(loan any) *BankBuilder {
	b.bank.Loans = append(b.bank.Loans, loan)
	return b
}

// WithInvestment adds an investment to the Investments slice
func (b *BankBuilder) WithInvestment(investment any) *BankBuilder {
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
func (b *BankBuilder) Build() any {
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
			return bank, err
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
		panic(fmt.Sprintf("Bank validation failed: %s", err.Error()))
	}
	return bank
}

// Clone creates a deep copy of the BankBuilder
func (b *BankBuilder) Clone() *BankBuilder {
	clonedBank := *b.bank
	return &BankBuilder{
		bank:            &clonedBank,
		validationFuncs: append([]func(*models.Bank) error{}, b.validationFuncs...),
	}
}
