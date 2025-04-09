package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// LoanBuilder builds a Loan model
type LoanBuilder struct {
	loan *models.Loan
	// Custom validation functions
	validationFuncs []func(*models.Loan) error
}

// NewLoanBuilder creates a new LoanBuilder
func NewLoanBuilder() *LoanBuilder {
	return &LoanBuilder{
		loan: &models.Loan{
			ID: "",
			Amount: 0.0,
			InterestRate: 0.0,
			StartDate: "",
			EndDate: "",
			MonthlyPayment: 0.0,
			RemainingBalance: 0.0,
			Type: "",
			Collateral: "",
			Status: "",
			Payments: []*models.Payment{},
		},
		validationFuncs: []func(*models.Loan) error{},
	}
}

// NewLoanBuilderWithDefaults creates a new LoanBuilder with sensible defaults
func NewLoanBuilderWithDefaults() *LoanBuilder {
	builder := NewLoanBuilder()
	// Add default values here if needed
	return builder
}
// WithID sets the ID
func (b *LoanBuilder) WithID(iD string) *LoanBuilder {
	b.loan.ID = iD
	return b
}

// WithAmount sets the Amount
func (b *LoanBuilder) WithAmount(amount float64) *LoanBuilder {
	b.loan.Amount = amount
	return b
}

// WithInterestRate sets the InterestRate
func (b *LoanBuilder) WithInterestRate(interestRate float64) *LoanBuilder {
	b.loan.InterestRate = interestRate
	return b
}

// WithStartDate sets the StartDate
func (b *LoanBuilder) WithStartDate(startDate string) *LoanBuilder {
	b.loan.StartDate = startDate
	return b
}

// WithEndDate sets the EndDate
func (b *LoanBuilder) WithEndDate(endDate string) *LoanBuilder {
	b.loan.EndDate = endDate
	return b
}

// WithMonthlyPayment sets the MonthlyPayment
func (b *LoanBuilder) WithMonthlyPayment(monthlyPayment float64) *LoanBuilder {
	b.loan.MonthlyPayment = monthlyPayment
	return b
}

// WithRemainingBalance sets the RemainingBalance
func (b *LoanBuilder) WithRemainingBalance(remainingBalance float64) *LoanBuilder {
	b.loan.RemainingBalance = remainingBalance
	return b
}

// WithType sets the Type
func (b *LoanBuilder) WithType(value string) *LoanBuilder {
	b.loan.Type = value
	return b
}

// WithCollateral sets the Collateral
func (b *LoanBuilder) WithCollateral(collateral string) *LoanBuilder {
	b.loan.Collateral = collateral
	return b
}

// WithStatus sets the Status
func (b *LoanBuilder) WithStatus(status string) *LoanBuilder {
	b.loan.Status = status
	return b
}

// WithPayments sets the Payments
func (b *LoanBuilder) WithPayments(payments *PaymentBuilder) *LoanBuilder {
	// Ensure the slice is initialized
	if b.loan.Payments == nil {
		b.loan.Payments = []*models.Payment{}
	}
	// Handle nested slice element
	builtValue := payments.Build().(*models.Payment)
	b.loan.Payments = append(b.loan.Payments, builtValue)
	return b
}


// WithValidation adds a custom validation function
func (b *LoanBuilder) WithValidation(validationFunc func(*models.Loan) error) *LoanBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Loan
func (b *LoanBuilder) Build() interface{} {
	return b.loan
}

// BuildPtr builds the Loan and returns a pointer
func (b *LoanBuilder) BuildPtr() *models.Loan {
	return b.loan
}

// BuildAndValidate builds the Loan and validates it
func (b *LoanBuilder) BuildAndValidate() (*models.Loan, error) {
	loan := b.loan

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(loan); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(loan).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return loan, err
		}
	}

	return loan, nil
}

// MustBuild builds the Loan and panics if validation fails
func (b *LoanBuilder) MustBuild() *models.Loan {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *LoanBuilder) Clone() *LoanBuilder {
	clonedLoan := *b.loan
	return &LoanBuilder{
		loan: &clonedLoan,
		validationFuncs: append([]func(*models.Loan) error{}, b.validationFuncs...),
	}
}
