package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// PaymentBuilder builds a Payment model
type PaymentBuilder struct {
	payment *models.Payment
	// Custom validation functions
	validationFuncs []func(*models.Payment) error
}

// NewPaymentBuilder creates a new PaymentBuilder
func NewPaymentBuilder() *PaymentBuilder {
	return &PaymentBuilder{
		payment: &models.Payment{
			Date: "",
			Amount: 0.0,
			Principal: 0.0,
			Interest: 0.0,
			Status: "",
			Transaction: nil,
		},
		validationFuncs: []func(*models.Payment) error{},
	}
}

// NewPaymentBuilderWithDefaults creates a new PaymentBuilder with sensible defaults
func NewPaymentBuilderWithDefaults() *PaymentBuilder {
	builder := NewPaymentBuilder()
	// Add default values here if needed
	return builder
}
// WithDate sets the Date
func (b *PaymentBuilder) WithDate(date string) *PaymentBuilder {
	b.payment.Date = date
	return b
}

// WithAmount sets the Amount
func (b *PaymentBuilder) WithAmount(amount float64) *PaymentBuilder {
	b.payment.Amount = amount
	return b
}

// WithPrincipal sets the Principal
func (b *PaymentBuilder) WithPrincipal(principal float64) *PaymentBuilder {
	b.payment.Principal = principal
	return b
}

// WithInterest sets the Interest
func (b *PaymentBuilder) WithInterest(interest float64) *PaymentBuilder {
	b.payment.Interest = interest
	return b
}

// WithStatus sets the Status
func (b *PaymentBuilder) WithStatus(status string) *PaymentBuilder {
	b.payment.Status = status
	return b
}

// WithTransaction sets the Transaction
func (b *PaymentBuilder) WithTransaction(transaction *TransactionBuilder) *PaymentBuilder {
	// Handle nested pointer
	b.payment.Transaction = transaction.BuildPtr()
	return b
}


// WithValidation adds a custom validation function
func (b *PaymentBuilder) WithValidation(validationFunc func(*models.Payment) error) *PaymentBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Payment
func (b *PaymentBuilder) Build() interface{} {
	return b.payment
}

// BuildPtr builds the Payment and returns a pointer
func (b *PaymentBuilder) BuildPtr() *models.Payment {
	return b.payment
}

// BuildAndValidate builds the Payment and validates it
func (b *PaymentBuilder) BuildAndValidate() (*models.Payment, error) {
	payment := b.payment

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(payment); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(payment).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return payment, err
		}
	}

	return payment, nil
}

// MustBuild builds the Payment and panics if validation fails
func (b *PaymentBuilder) MustBuild() *models.Payment {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *PaymentBuilder) Clone() *PaymentBuilder {
	clonedPayment := *b.payment
	return &PaymentBuilder{
		payment: &clonedPayment,
		validationFuncs: append([]func(*models.Payment) error{}, b.validationFuncs...),
	}
}
