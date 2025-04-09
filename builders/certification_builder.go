package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// CertificationBuilder builds a Certification model
type CertificationBuilder struct {
	certification *models.Certification
	// Custom validation functions
	validationFuncs []func(*models.Certification) error
}

// NewCertificationBuilder creates a new CertificationBuilder
func NewCertificationBuilder() *CertificationBuilder {
	return &CertificationBuilder{
		certification: &models.Certification{
			Name: "",
			Issuer: "",
			IssueDate: "",
			ExpiryDate: "",
			CredentialID: "",
			URL: "",
		},
		validationFuncs: []func(*models.Certification) error{},
	}
}

// NewCertificationBuilderWithDefaults creates a new CertificationBuilder with sensible defaults
func NewCertificationBuilderWithDefaults() *CertificationBuilder {
	builder := NewCertificationBuilder()
	// Add default values here if needed
	return builder
}
// WithName sets the Name
func (b *CertificationBuilder) WithName(name string) *CertificationBuilder {
	b.certification.Name = name
	return b
}

// WithIssuer sets the Issuer
func (b *CertificationBuilder) WithIssuer(issuer string) *CertificationBuilder {
	b.certification.Issuer = issuer
	return b
}

// WithIssueDate sets the IssueDate
func (b *CertificationBuilder) WithIssueDate(issueDate string) *CertificationBuilder {
	b.certification.IssueDate = issueDate
	return b
}

// WithExpiryDate sets the ExpiryDate
func (b *CertificationBuilder) WithExpiryDate(expiryDate string) *CertificationBuilder {
	b.certification.ExpiryDate = expiryDate
	return b
}

// WithCredentialID sets the CredentialID
func (b *CertificationBuilder) WithCredentialID(credentialID string) *CertificationBuilder {
	b.certification.CredentialID = credentialID
	return b
}

// WithURL sets the URL
func (b *CertificationBuilder) WithURL(uRL string) *CertificationBuilder {
	b.certification.URL = uRL
	return b
}


// WithValidation adds a custom validation function
func (b *CertificationBuilder) WithValidation(validationFunc func(*models.Certification) error) *CertificationBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Certification
func (b *CertificationBuilder) Build() interface{} {
	return b.certification
}

// BuildPtr builds the Certification and returns a pointer
func (b *CertificationBuilder) BuildPtr() *models.Certification {
	return b.certification
}

// BuildAndValidate builds the Certification and validates it
func (b *CertificationBuilder) BuildAndValidate() (*models.Certification, error) {
	certification := b.certification

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(certification); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(certification).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return certification, err
		}
	}

	return certification, nil
}

// MustBuild builds the Certification and panics if validation fails
func (b *CertificationBuilder) MustBuild() *models.Certification {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *CertificationBuilder) Clone() *CertificationBuilder {
	clonedCertification := *b.certification
	return &CertificationBuilder{
		certification: &clonedCertification,
		validationFuncs: append([]func(*models.Certification) error{}, b.validationFuncs...),
	}
}
