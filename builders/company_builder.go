package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/v2/models"
	"github.com/adil-faiyaz98/go-builder-kit/v2/pkg/builder"
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
func (b *CompanyBuilder) WithID(id string) *CompanyBuilder {
	if b == nil {
		return b
	}
	b.company.ID = builder.SanitizeString(id)
	return b
}

// WithName sets the Name
func (b *CompanyBuilder) WithName(name string) *CompanyBuilder {
	if b == nil {
		return b
	}
	b.company.Name = builder.SanitizeString(name)
	return b
}

// WithIndustry sets the Industry
func (b *CompanyBuilder) WithIndustry(industry string) *CompanyBuilder {
	if b == nil {
		return b
	}
	b.company.Industry = builder.SanitizeString(industry)
	return b
}

// WithDescription sets the Description
func (b *CompanyBuilder) WithDescription(description string) *CompanyBuilder {
	if b == nil {
		return b
	}
	b.company.Description = builder.SanitizeString(description)
	return b
}

// WithFounded sets the Founded
func (b *CompanyBuilder) WithFounded(founded string) *CompanyBuilder {
	if b == nil {
		return b
	}
	b.company.Founded = builder.SanitizeString(founded)
	return b
}

// WithWebsite sets the Website
func (b *CompanyBuilder) WithWebsite(website string) *CompanyBuilder {
	if b == nil {
		return b
	}
	b.company.Website = builder.SanitizeString(website)
	return b
}

// WithAddress sets the Address
func (b *CompanyBuilder) WithAddress(address *AddressBuilder) *CompanyBuilder {
	if b == nil {
		return b
	}
	// Handle nested pointer
	b.company.Address = address.BuildPtr()
	return b
}

// WithLocation sets the Location
func (b *CompanyBuilder) WithLocation(location *AddressBuilder) *CompanyBuilder {
	if b == nil {
		return b
	}
	// Handle nested pointer
	b.company.Location = location.BuildPtr()
	return b
}

// WithSize sets the Size
func (b *CompanyBuilder) WithSize(size string) *CompanyBuilder {
	if b == nil {
		return b
	}
	b.company.Size = builder.SanitizeString(size)
	return b
}

// WithRevenue sets the Revenue
func (b *CompanyBuilder) WithRevenue(revenue string) *CompanyBuilder {
	if b == nil {
		return b
	}
	b.company.Revenue = builder.SanitizeString(revenue)
	return b
}

// WithPublic sets the Public
func (b *CompanyBuilder) WithPublic(public bool) *CompanyBuilder {
	if b == nil {
		return b
	}
	b.company.Public = public
	return b
}

// WithStockSymbol sets the StockSymbol
func (b *CompanyBuilder) WithStockSymbol(stockSymbol string) *CompanyBuilder {
	if b == nil {
		return b
	}
	b.company.StockSymbol = builder.SanitizeString(stockSymbol)
	return b
}

// WithDepartments sets the Departments
func (b *CompanyBuilder) WithDepartments(departments []*DepartmentBuilder) *CompanyBuilder {
	if b == nil {
		return b
	}
	// Ensure the slice is initialized
	if b.company.Departments == nil {
		b.company.Departments = []*models.Department{}
	}
	// Handle nested slice elements
	// Initialize the slice
	b.company.Departments = make([]*models.Department, 0, len(departments))
	// Convert each builder to its model
	for _, builder := range departments {
		builtValue := builder.Build().(*models.Department)
		b.company.Departments = append(b.company.Departments, builtValue)
	}
	return b
}

// AddDepartment adds a single item to the Departments slice
func (b *CompanyBuilder) AddDepartment(department *DepartmentBuilder) *CompanyBuilder {
	if b == nil {
		return b
	}
	// Ensure the slice is initialized with capacity
	if b.company.Departments == nil {
		b.company.Departments = make([]*models.Department, 0, 4) // Pre-allocate capacity
	}
	// Handle nested slice element
	if department != nil {
		builtValue := department.Build().(*models.Department)
		b.company.Departments = append(b.company.Departments, builtValue)
	}
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
	if b == nil || b.company == nil {
		return nil, fmt.Errorf("builder is not properly initialized")
	}

	company := b.company

	// Run custom validation functions
	for i, validationFunc := range b.validationFuncs {
		if validationFunc == nil {
			continue // Skip nil validators
		}
		if err := validationFunc(company); err != nil {
			return nil, fmt.Errorf("custom validation failed at index %d: %w", i, err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := any(company).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return company, fmt.Errorf("model validation failed: %w", err)
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
	if b == nil || b.company == nil {
		return NewCompanyBuilder()
	}

	// Deep copy the struct
	clonedCompany := *b.company

	// Create new builder with cloned data
	clonedBuilder := &CompanyBuilder{
		company: &clonedCompany,
		validationFuncs: make([]func(*models.Company) error, 0, len(b.validationFuncs)),
	}

	// Copy validation functions safely
	if len(b.validationFuncs) > 0 {
		clonedBuilder.validationFuncs = append(clonedBuilder.validationFuncs, b.validationFuncs...)
	}

	return clonedBuilder
}
