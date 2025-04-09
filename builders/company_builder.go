package builders

import (
	"fmt"

	"github.com/adil-faiyaz98/go-builder-kit/models"
)

// CompanyBuilder builds a Company model
type CompanyBuilder struct {
	company         *models.Company
	validationFuncs []func(*models.Company) error
}

// NewCompanyBuilder creates a new CompanyBuilder
func NewCompanyBuilder() *CompanyBuilder {
	return &CompanyBuilder{
		company: &models.Company{
			Name:        "",
			Industry:    "",
			Description: "",
			Founded:     "",
			Website:     "",
			Address:     nil,
			Location:    nil,
			Size:        "",
			Revenue:     "",
			Public:      false,
			StockSymbol: "",
		},
		validationFuncs: []func(*models.Company) error{},
	}
}

// WithName sets the Name
func (b *CompanyBuilder) WithName(name string) *CompanyBuilder {
	b.company.Name = name
	return b
}

// WithIndustry sets the Industry
func (b *CompanyBuilder) WithIndustry(industry string) *CompanyBuilder {
	b.company.Industry = industry
	return b
}

// WithDescription sets the Description
func (b *CompanyBuilder) WithDescription(description string) *CompanyBuilder {
	b.company.Description = description
	return b
}

// WithFounded sets the Founded
func (b *CompanyBuilder) WithFounded(founded string) *CompanyBuilder {
	b.company.Founded = founded
	return b
}

// WithWebsite sets the Website
func (b *CompanyBuilder) WithWebsite(website string) *CompanyBuilder {
	b.company.Website = website
	return b
}

// WithAddress sets the Address
func (b *CompanyBuilder) WithAddress(address *AddressBuilder) *CompanyBuilder {
	builtValue := address.Build().(*models.Address)
	b.company.Address = builtValue
	return b
}

// WithLocation sets the Location
func (b *CompanyBuilder) WithLocation(location *AddressBuilder) *CompanyBuilder {
	builtValue := location.Build().(*models.Address)
	b.company.Location = builtValue
	return b
}

// WithSize sets the Size
func (b *CompanyBuilder) WithSize(size string) *CompanyBuilder {
	b.company.Size = size
	return b
}

// WithRevenue sets the Revenue
func (b *CompanyBuilder) WithRevenue(revenue string) *CompanyBuilder {
	b.company.Revenue = revenue
	return b
}

// WithPublic sets the Public
func (b *CompanyBuilder) WithPublic(public bool) *CompanyBuilder {
	b.company.Public = public
	return b
}

// WithStockSymbol sets the StockSymbol
func (b *CompanyBuilder) WithStockSymbol(stockSymbol string) *CompanyBuilder {
	b.company.StockSymbol = stockSymbol
	return b
}

// WithValidation adds a custom validation function
func (b *CompanyBuilder) WithValidation(validationFunc func(*models.Company) error) *CompanyBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Company
func (b *CompanyBuilder) Build() any {
	return b.company
}

// BuildPtr builds the Company and returns a pointer
func (b *CompanyBuilder) BuildPtr() *models.Company {
	return b.company
}

// BuildAndValidate builds the Company and validates it
func (b *CompanyBuilder) BuildAndValidate() (*models.Company, error) {
	company := b.company

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(company); err != nil {
			return company, err
		}
	}

	// Run model's Validate method
	if err := company.Validate(); err != nil {
		return company, err
	}

	return company, nil
}

// MustBuild builds the Company and panics if validation fails
func (b *CompanyBuilder) MustBuild() *models.Company {
	company, err := b.BuildAndValidate()
	if err != nil {
		panic(fmt.Sprintf("Company validation failed: %s", err.Error()))
	}
	return company
}

// Clone creates a deep copy of the CompanyBuilder
func (b *CompanyBuilder) Clone() *CompanyBuilder {
	clonedCompany := *b.company
	return &CompanyBuilder{
		company:         &clonedCompany,
		validationFuncs: append([]func(*models.Company) error{}, b.validationFuncs...),
	}
}
