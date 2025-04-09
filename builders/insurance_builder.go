package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// InsuranceBuilder builds a Insurance model
type InsuranceBuilder struct {
	insurance *models.Insurance
	// Custom validation functions
	validationFuncs []func(*models.Insurance) error
}

// NewInsuranceBuilder creates a new InsuranceBuilder
func NewInsuranceBuilder() *InsuranceBuilder {
	return &InsuranceBuilder{
		insurance: &models.Insurance{
			Provider: "",
			PolicyNumber: "",
			Type: "",
			StartDate: "",
			EndDate: "",
			Coverage: map[string]float64{},
			Deductible: 0.0,
			CoPay: 0.0,
			Copay: 0.0,
			OutOfPocketMax: 0.0,
			Beneficiaries: []any{},
		},
		validationFuncs: []func(*models.Insurance) error{},
	}
}

// NewInsuranceBuilderWithDefaults creates a new InsuranceBuilder with sensible defaults
func NewInsuranceBuilderWithDefaults() *InsuranceBuilder {
	builder := NewInsuranceBuilder()
	// Add default values here if needed
	return builder
}
// WithProvider sets the Provider
func (b *InsuranceBuilder) WithProvider(provider string) *InsuranceBuilder {
	b.insurance.Provider = provider
	return b
}

// WithPolicyNumber sets the PolicyNumber
func (b *InsuranceBuilder) WithPolicyNumber(policyNumber string) *InsuranceBuilder {
	b.insurance.PolicyNumber = policyNumber
	return b
}

// WithType sets the Type
func (b *InsuranceBuilder) WithType(value string) *InsuranceBuilder {
	b.insurance.Type = value
	return b
}

// WithStartDate sets the StartDate
func (b *InsuranceBuilder) WithStartDate(startDate string) *InsuranceBuilder {
	b.insurance.StartDate = startDate
	return b
}

// WithEndDate sets the EndDate
func (b *InsuranceBuilder) WithEndDate(endDate string) *InsuranceBuilder {
	b.insurance.EndDate = endDate
	return b
}

// WithCoverage sets the Coverage
func (b *InsuranceBuilder) WithCoverage(key string, val float64) *InsuranceBuilder {
	if b.insurance.Coverage == nil {
		b.insurance.Coverage = make(map[string]float64)
	}
	b.insurance.Coverage[key] = val
	return b
}

// WithDeductible sets the Deductible
func (b *InsuranceBuilder) WithDeductible(deductible float64) *InsuranceBuilder {
	b.insurance.Deductible = deductible
	return b
}

// WithCoPay sets the CoPay
func (b *InsuranceBuilder) WithCoPay(coPay float64) *InsuranceBuilder {
	b.insurance.CoPay = coPay
	return b
}

// WithCopay sets the Copay
func (b *InsuranceBuilder) WithCopay(copay float64) *InsuranceBuilder {
	b.insurance.Copay = copay
	return b
}

// WithOutOfPocketMax sets the OutOfPocketMax
func (b *InsuranceBuilder) WithOutOfPocketMax(outOfPocketMax float64) *InsuranceBuilder {
	b.insurance.OutOfPocketMax = outOfPocketMax
	return b
}

// WithBeneficiaries sets the Beneficiaries
func (b *InsuranceBuilder) WithBeneficiaries(beneficiaries any) *InsuranceBuilder {
	b.insurance.Beneficiaries = append(b.insurance.Beneficiaries, beneficiaries)
	return b
}


// WithValidation adds a custom validation function
func (b *InsuranceBuilder) WithValidation(validationFunc func(*models.Insurance) error) *InsuranceBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Insurance
func (b *InsuranceBuilder) Build() interface{} {
	return b.insurance
}

// BuildPtr builds the Insurance and returns a pointer
func (b *InsuranceBuilder) BuildPtr() *models.Insurance {
	return b.insurance
}

// BuildAndValidate builds the Insurance and validates it
func (b *InsuranceBuilder) BuildAndValidate() (*models.Insurance, error) {
	insurance := b.insurance

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(insurance); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(insurance).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return insurance, err
		}
	}

	return insurance, nil
}

// MustBuild builds the Insurance and panics if validation fails
func (b *InsuranceBuilder) MustBuild() *models.Insurance {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *InsuranceBuilder) Clone() *InsuranceBuilder {
	clonedInsurance := *b.insurance
	return &InsuranceBuilder{
		insurance: &clonedInsurance,
		validationFuncs: append([]func(*models.Insurance) error{}, b.validationFuncs...),
	}
}
