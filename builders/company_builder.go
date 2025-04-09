package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// CompanyBuilder builds a Company model
type CompanyBuilder struct {
	company *models.Company
	// Custom validation functions
	validationFuncs []func(*models.Company) error
}

// NewCompanyBuilder creates a new CompanyBuilder
func NewCompanyBuilder() *CompanyBuilder {
	return &CompanyBuilder{
		company: &models.Company{
			ID: "",
			Name: "",
			Industry: "",
			Description: "",
			Founded: "",
			Website: "",
			Address: nil,
			Location: nil,
			Size: "",
			Revenue: "",
			Public: false,
			StockSymbol: "",
			Departments: []*models.Department{},
		},
		validationFuncs: []func(*models.Company) error{},
	}
}

// NewCompanyBuilderWithDefaults creates a new CompanyBuilder with sensible defaults
func NewCompanyBuilderWithDefaults() *CompanyBuilder {
	builder := NewCompanyBuilder()
	// Add default values here if needed
	return builder
}
// WithID sets the ID
func (b *CompanyBuilder) WithID(iD string) *CompanyBuilder {
	b.company.ID = iD
	return b
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
	// Handle nested pointer
	b.company.Address = address.BuildPtr()
	return b
}

// WithLocation sets the Location
func (b *CompanyBuilder) WithLocation(location *AddressBuilder) *CompanyBuilder {
	// Handle nested pointer
	b.company.Location = location.BuildPtr()
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

// WithDepartments sets the Departments
func (b *CompanyBuilder) WithDepartments(departments *DepartmentBuilder) *CompanyBuilder {
	// Ensure the slice is initialized
	if b.company.Departments == nil {
		b.company.Departments = []*models.Department{}
	}
	// Handle nested slice element
	builtValue := departments.Build().(*models.Department)
	b.company.Departments = append(b.company.Departments, builtValue)
	return b
}


// WithValidation adds a custom validation function
func (b *CompanyBuilder) WithValidation(validationFunc func(*models.Company) error) *CompanyBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Company
func (b *CompanyBuilder) Build() interface{} {
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
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(company).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return company, err
		}
	}

	return company, nil
}

// MustBuild builds the Company and panics if validation fails
func (b *CompanyBuilder) MustBuild() *models.Company {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *CompanyBuilder) Clone() *CompanyBuilder {
	clonedCompany := *b.company
	return &CompanyBuilder{
		company: &clonedCompany,
		validationFuncs: append([]func(*models.Company) error{}, b.validationFuncs...),
	}
}
