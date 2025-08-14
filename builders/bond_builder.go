package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/v2/models"
	"github.com/adil-faiyaz98/go-builder-kit/v2/pkg/builder"
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
func (b *BondBuilder) WithID(id string) *BondBuilder {
	if b == nil {
		return b
	}
	b.bond.ID = builder.SanitizeString(id)
	return b
}

// WithISIN sets the ISIN
func (b *BondBuilder) WithISIN(iSIN string) *BondBuilder {
	if b == nil {
		return b
	}
	b.bond.ISIN = builder.SanitizeString(iSIN)
	return b
}

// WithName sets the Name
func (b *BondBuilder) WithName(name string) *BondBuilder {
	if b == nil {
		return b
	}
	b.bond.Name = builder.SanitizeString(name)
	return b
}

// WithIssuer sets the Issuer
func (b *BondBuilder) WithIssuer(issuer string) *BondBuilder {
	if b == nil {
		return b
	}
	b.bond.Issuer = builder.SanitizeString(issuer)
	return b
}

// WithType sets the Type
func (b *BondBuilder) WithType(value string) *BondBuilder {
	if b == nil {
		return b
	}
	b.bond.Type = builder.SanitizeString(value)
	return b
}

// WithFaceValue sets the FaceValue
func (b *BondBuilder) WithFaceValue(faceValue float64) *BondBuilder {
	if b == nil {
		return b
	}
	b.bond.FaceValue = faceValue
	return b
}

// WithCouponRate sets the CouponRate
func (b *BondBuilder) WithCouponRate(couponRate float64) *BondBuilder {
	if b == nil {
		return b
	}
	b.bond.CouponRate = couponRate
	return b
}

// WithMaturityDate sets the MaturityDate
func (b *BondBuilder) WithMaturityDate(maturityDate string) *BondBuilder {
	if b == nil {
		return b
	}
	b.bond.MaturityDate = builder.SanitizeString(maturityDate)
	return b
}

// WithPurchaseDate sets the PurchaseDate
func (b *BondBuilder) WithPurchaseDate(purchaseDate string) *BondBuilder {
	if b == nil {
		return b
	}
	b.bond.PurchaseDate = builder.SanitizeString(purchaseDate)
	return b
}

// WithPurchasePrice sets the PurchasePrice
func (b *BondBuilder) WithPurchasePrice(purchasePrice float64) *BondBuilder {
	if b == nil {
		return b
	}
	b.bond.PurchasePrice = purchasePrice
	return b
}

// WithCurrentPrice sets the CurrentPrice
func (b *BondBuilder) WithCurrentPrice(currentPrice float64) *BondBuilder {
	if b == nil {
		return b
	}
	b.bond.CurrentPrice = currentPrice
	return b
}

// WithQuantity sets the Quantity
func (b *BondBuilder) WithQuantity(quantity int) *BondBuilder {
	if b == nil {
		return b
	}
	b.bond.Quantity = quantity
	return b
}

// WithCurrency sets the Currency
func (b *BondBuilder) WithCurrency(currency string) *BondBuilder {
	if b == nil {
		return b
	}
	b.bond.Currency = builder.SanitizeString(currency)
	return b
}

// WithPaymentFrequency sets the PaymentFrequency
func (b *BondBuilder) WithPaymentFrequency(paymentFrequency string) *BondBuilder {
	if b == nil {
		return b
	}
	b.bond.PaymentFrequency = builder.SanitizeString(paymentFrequency)
	return b
}

// WithRating sets the Rating
func (b *BondBuilder) WithRating(rating string) *BondBuilder {
	if b == nil {
		return b
	}
	b.bond.Rating = builder.SanitizeString(rating)
	return b
}

// WithYield sets the Yield
func (b *BondBuilder) WithYield(yield float64) *BondBuilder {
	if b == nil {
		return b
	}
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
	if b == nil || b.bond == nil {
		return nil, fmt.Errorf("builder is not properly initialized")
	}

	bond := b.bond

	// Run custom validation functions
	for i, validationFunc := range b.validationFuncs {
		if validationFunc == nil {
			continue // Skip nil validators
		}
		if err := validationFunc(bond); err != nil {
			return nil, fmt.Errorf("custom validation failed at index %d: %w", i, err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := any(bond).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return bond, fmt.Errorf("model validation failed: %w", err)
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
	if b == nil || b.bond == nil {
		return NewBondBuilder()
	}

	// Deep copy the struct
	clonedBond := *b.bond

	// Create new builder with cloned data
	clonedBuilder := &BondBuilder{
		bond: &clonedBond,
		validationFuncs: make([]func(*models.Bond) error, 0, len(b.validationFuncs)),
	}

	// Copy validation functions safely
	if len(b.validationFuncs) > 0 {
		clonedBuilder.validationFuncs = append(clonedBuilder.validationFuncs, b.validationFuncs...)
	}

	return clonedBuilder
}
