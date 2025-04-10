package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/test/models"
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
			Name: "",
			Description: "",
			Founded: "",
			Industry: "",
			Website: "",
			Headquarters: nil,
			Departments: []*models.Department{},
			Employees: []*models.Employee{},
			Projects: []*models.Project{},
			Subsidiaries: []*models.Company{},
			ParentCompany: nil,
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
// WithName sets the Name
func (b *CompanyBuilder) WithName(name string) *CompanyBuilder {
	b.company.Name = name
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

// WithIndustry sets the Industry
func (b *CompanyBuilder) WithIndustry(industry string) *CompanyBuilder {
	b.company.Industry = industry
	return b
}

// WithWebsite sets the Website
func (b *CompanyBuilder) WithWebsite(website string) *CompanyBuilder {
	b.company.Website = website
	return b
}

// WithHeadquarters sets the Headquarters
func (b *CompanyBuilder) WithHeadquarters(headquarters *AddressBuilder) *CompanyBuilder {
	// Handle nested pointer
	b.company.Headquarters = headquarters.BuildPtr()
	return b
}

// WithDepartments sets the Departments
func (b *CompanyBuilder) WithDepartments(departments []*DepartmentBuilder) *CompanyBuilder {
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

// WithEmployees sets the Employees
func (b *CompanyBuilder) WithEmployees(employees []*EmployeeBuilder) *CompanyBuilder {
	// Ensure the slice is initialized
	if b.company.Employees == nil {
		b.company.Employees = []*models.Employee{}
	}
	// Handle nested slice elements
	// Initialize the slice
	b.company.Employees = make([]*models.Employee, 0, len(employees))
	// Convert each builder to its model
	for _, builder := range employees {
		builtValue := builder.Build().(*models.Employee)
		b.company.Employees = append(b.company.Employees, builtValue)
	}
	return b
}

// WithProjects sets the Projects
func (b *CompanyBuilder) WithProjects(projects []*ProjectBuilder) *CompanyBuilder {
	// Ensure the slice is initialized
	if b.company.Projects == nil {
		b.company.Projects = []*models.Project{}
	}
	// Handle nested slice elements
	// Initialize the slice
	b.company.Projects = make([]*models.Project, 0, len(projects))
	// Convert each builder to its model
	for _, builder := range projects {
		builtValue := builder.Build().(*models.Project)
		b.company.Projects = append(b.company.Projects, builtValue)
	}
	return b
}

// WithSubsidiaries sets the Subsidiaries
func (b *CompanyBuilder) WithSubsidiaries(subsidiaries []*CompanyBuilder) *CompanyBuilder {
	// Ensure the slice is initialized
	if b.company.Subsidiaries == nil {
		b.company.Subsidiaries = []*models.Company{}
	}
	// Handle nested slice elements
	// Initialize the slice
	b.company.Subsidiaries = make([]*models.Company, 0, len(subsidiaries))
	// Convert each builder to its model
	for _, builder := range subsidiaries {
		builtValue := builder.Build().(*models.Company)
		b.company.Subsidiaries = append(b.company.Subsidiaries, builtValue)
	}
	return b
}

// WithParentCompany sets the ParentCompany
func (b *CompanyBuilder) WithParentCompany(parentCompany *CompanyBuilder) *CompanyBuilder {
	// Handle nested pointer
	b.company.ParentCompany = parentCompany.BuildPtr()
	return b
}

// AddDepartment adds a single item to the Departments slice
func (b *CompanyBuilder) AddDepartment(department *DepartmentBuilder) *CompanyBuilder {
	// Ensure the slice is initialized
	if b.company.Departments == nil {
		b.company.Departments = []*models.Department{}
	}
	// Handle nested slice element
	builtValue := department.Build().(*models.Department)
	b.company.Departments = append(b.company.Departments, builtValue)
	return b
}
// AddEmployee adds a single item to the Employees slice
func (b *CompanyBuilder) AddEmployee(employee *EmployeeBuilder) *CompanyBuilder {
	// Ensure the slice is initialized
	if b.company.Employees == nil {
		b.company.Employees = []*models.Employee{}
	}
	// Handle nested slice element
	builtValue := employee.Build().(*models.Employee)
	b.company.Employees = append(b.company.Employees, builtValue)
	return b
}
// AddProject adds a single item to the Projects slice
func (b *CompanyBuilder) AddProject(project *ProjectBuilder) *CompanyBuilder {
	// Ensure the slice is initialized
	if b.company.Projects == nil {
		b.company.Projects = []*models.Project{}
	}
	// Handle nested slice element
	builtValue := project.Build().(*models.Project)
	b.company.Projects = append(b.company.Projects, builtValue)
	return b
}
// AddSubsidiary adds a single item to the Subsidiaries slice
func (b *CompanyBuilder) AddSubsidiary(subsidiary *CompanyBuilder) *CompanyBuilder {
	// Ensure the slice is initialized
	if b.company.Subsidiaries == nil {
		b.company.Subsidiaries = []*models.Company{}
	}
	// Handle nested slice element
	builtValue := subsidiary.Build().(*models.Company)
	b.company.Subsidiaries = append(b.company.Subsidiaries, builtValue)
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
