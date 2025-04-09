package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// BondBuilder builds a Bond model
type BondBuilder struct {
	bond *models.Bond
	// Custom validation functions
	validationFuncs []func(*models.Bond) error
}

// NewBondBuilder creates a new BondBuilder
func NewBondBuilder() *BondBuilder {
	return &BondBuilder{
		bond: &models.Bond{
			ID: "",
			ISIN: "",
			Name: "",
			Issuer: "",
			Type: "",
			FaceValue: 0.0,
			CouponRate: 0.0,
			MaturityDate: "",
			PurchaseDate: "",
			PurchasePrice: 0.0,
			CurrentPrice: 0.0,
			Quantity: 0,
			Currency: "",
			PaymentFrequency: "",
			Rating: "",
			Yield: 0.0,
		},
		validationFuncs: []func(*models.Bond) error{},
	}
}

// NewBondBuilderWithDefaults creates a new BondBuilder with sensible defaults
func NewBondBuilderWithDefaults() *BondBuilder {
	builder := NewBondBuilder()
	// Add default values here if needed
	return builder
}
// WithID sets the ID
func (b *BondBuilder) WithID(iD string) *BondBuilder {
	b.bond.ID = iD
	return b
}

// WithISIN sets the ISIN
func (b *BondBuilder) WithISIN(iSIN string) *BondBuilder {
	b.bond.ISIN = iSIN
	return b
}

// WithName sets the Name
func (b *BondBuilder) WithName(name string) *BondBuilder {
	b.bond.Name = name
	return b
}

// WithIssuer sets the Issuer
func (b *BondBuilder) WithIssuer(issuer string) *BondBuilder {
	b.bond.Issuer = issuer
	return b
}

// WithType sets the Type
func (b *BondBuilder) WithType(value string) *BondBuilder {
	b.bond.Type = value
	return b
}

// WithFaceValue sets the FaceValue
func (b *BondBuilder) WithFaceValue(faceValue float64) *BondBuilder {
	b.bond.FaceValue = faceValue
	return b
}

// WithCouponRate sets the CouponRate
func (b *BondBuilder) WithCouponRate(couponRate float64) *BondBuilder {
	b.bond.CouponRate = couponRate
	return b
}

// WithMaturityDate sets the MaturityDate
func (b *BondBuilder) WithMaturityDate(maturityDate string) *BondBuilder {
	b.bond.MaturityDate = maturityDate
	return b
}

// WithPurchaseDate sets the PurchaseDate
func (b *BondBuilder) WithPurchaseDate(purchaseDate string) *BondBuilder {
	b.bond.PurchaseDate = purchaseDate
	return b
}

// WithPurchasePrice sets the PurchasePrice
func (b *BondBuilder) WithPurchasePrice(purchasePrice float64) *BondBuilder {
	b.bond.PurchasePrice = purchasePrice
	return b
}

// WithCurrentPrice sets the CurrentPrice
func (b *BondBuilder) WithCurrentPrice(currentPrice float64) *BondBuilder {
	b.bond.CurrentPrice = currentPrice
	return b
}

// WithQuantity sets the Quantity
func (b *BondBuilder) WithQuantity(quantity int) *BondBuilder {
	b.bond.Quantity = quantity
	return b
}

// WithCurrency sets the Currency
func (b *BondBuilder) WithCurrency(currency string) *BondBuilder {
	b.bond.Currency = currency
	return b
}

// WithPaymentFrequency sets the PaymentFrequency
func (b *BondBuilder) WithPaymentFrequency(paymentFrequency string) *BondBuilder {
	b.bond.PaymentFrequency = paymentFrequency
	return b
}

// WithRating sets the Rating
func (b *BondBuilder) WithRating(rating string) *BondBuilder {
	b.bond.Rating = rating
	return b
}

// WithYield sets the Yield
func (b *BondBuilder) WithYield(yield float64) *BondBuilder {
	b.bond.Yield = yield
	return b
}


// WithValidation adds a custom validation function
func (b *BondBuilder) WithValidation(validationFunc func(*models.Bond) error) *BondBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Bond
func (b *BondBuilder) Build() interface{} {
	return b.bond
}

// BuildPtr builds the Bond and returns a pointer
func (b *BondBuilder) BuildPtr() *models.Bond {
	return b.bond
}

// BuildAndValidate builds the Bond and validates it
func (b *BondBuilder) BuildAndValidate() (*models.Bond, error) {
	bond := b.bond

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(bond); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(bond).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return bond, err
		}
	}

	return bond, nil
}

// MustBuild builds the Bond and panics if validation fails
func (b *BondBuilder) MustBuild() *models.Bond {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *BondBuilder) Clone() *BondBuilder {
	clonedBond := *b.bond
	return &BondBuilder{
		bond: &clonedBond,
		validationFuncs: append([]func(*models.Bond) error{}, b.validationFuncs...),
	}
}
