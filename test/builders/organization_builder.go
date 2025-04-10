package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/test/models"
)

// OrganizationBuilder builds a Organization model
type OrganizationBuilder struct {
	organization *models.Organization
	// Custom validation functions
	validationFuncs []func(*models.Organization) error
}

// NewOrganizationBuilder creates a new OrganizationBuilder
func NewOrganizationBuilder() *OrganizationBuilder {
	return &OrganizationBuilder{
		organization: &models.Organization{
			Name: "",
			Type: "",
			Description: "",
			Headquarters: nil,
			Companies: []*models.Company{},
			Departments: []*models.Department{},
			Employees: []*models.Employee{},
			Projects: []*models.Project{},
			Contacts: []*models.Contact{},
		},
		validationFuncs: []func(*models.Organization) error{},
	}
}

// NewOrganizationBuilderWithDefaults creates a new OrganizationBuilder with sensible defaults
func NewOrganizationBuilderWithDefaults() *OrganizationBuilder {
	builder := NewOrganizationBuilder()
	// Add default values here if needed
	return builder
}
// WithName sets the Name
func (b *OrganizationBuilder) WithName(name string) *OrganizationBuilder {
	b.organization.Name = name
	return b
}

// WithType sets the Type
func (b *OrganizationBuilder) WithType(value string) *OrganizationBuilder {
	b.organization.Type = value
	return b
}

// WithDescription sets the Description
func (b *OrganizationBuilder) WithDescription(description string) *OrganizationBuilder {
	b.organization.Description = description
	return b
}

// WithHeadquarters sets the Headquarters
func (b *OrganizationBuilder) WithHeadquarters(headquarters *AddressBuilder) *OrganizationBuilder {
	// Handle nested pointer
	b.organization.Headquarters = headquarters.BuildPtr()
	return b
}

// WithCompanies sets the Companies
func (b *OrganizationBuilder) WithCompanies(companies []*CompanyBuilder) *OrganizationBuilder {
	// Ensure the slice is initialized
	if b.organization.Companies == nil {
		b.organization.Companies = []*models.Company{}
	}
	// Handle nested slice elements
	// Initialize the slice
	b.organization.Companies = make([]*models.Company, 0, len(companies))
	// Convert each builder to its model
	for _, builder := range companies {
		builtValue := builder.Build().(*models.Company)
		b.organization.Companies = append(b.organization.Companies, builtValue)
	}
	return b
}

// WithDepartments sets the Departments
func (b *OrganizationBuilder) WithDepartments(departments []*DepartmentBuilder) *OrganizationBuilder {
	// Ensure the slice is initialized
	if b.organization.Departments == nil {
		b.organization.Departments = []*models.Department{}
	}
	// Handle nested slice elements
	// Initialize the slice
	b.organization.Departments = make([]*models.Department, 0, len(departments))
	// Convert each builder to its model
	for _, builder := range departments {
		builtValue := builder.Build().(*models.Department)
		b.organization.Departments = append(b.organization.Departments, builtValue)
	}
	return b
}

// WithEmployees sets the Employees
func (b *OrganizationBuilder) WithEmployees(employees []*EmployeeBuilder) *OrganizationBuilder {
	// Ensure the slice is initialized
	if b.organization.Employees == nil {
		b.organization.Employees = []*models.Employee{}
	}
	// Handle nested slice elements
	// Initialize the slice
	b.organization.Employees = make([]*models.Employee, 0, len(employees))
	// Convert each builder to its model
	for _, builder := range employees {
		builtValue := builder.Build().(*models.Employee)
		b.organization.Employees = append(b.organization.Employees, builtValue)
	}
	return b
}

// WithProjects sets the Projects
func (b *OrganizationBuilder) WithProjects(projects []*ProjectBuilder) *OrganizationBuilder {
	// Ensure the slice is initialized
	if b.organization.Projects == nil {
		b.organization.Projects = []*models.Project{}
	}
	// Handle nested slice elements
	// Initialize the slice
	b.organization.Projects = make([]*models.Project, 0, len(projects))
	// Convert each builder to its model
	for _, builder := range projects {
		builtValue := builder.Build().(*models.Project)
		b.organization.Projects = append(b.organization.Projects, builtValue)
	}
	return b
}

// WithContacts sets the Contacts
func (b *OrganizationBuilder) WithContacts(contacts []*ContactBuilder) *OrganizationBuilder {
	// Ensure the slice is initialized
	if b.organization.Contacts == nil {
		b.organization.Contacts = []*models.Contact{}
	}
	// Handle nested slice elements
	// Initialize the slice
	b.organization.Contacts = make([]*models.Contact, 0, len(contacts))
	// Convert each builder to its model
	for _, builder := range contacts {
		builtValue := builder.Build().(*models.Contact)
		b.organization.Contacts = append(b.organization.Contacts, builtValue)
	}
	return b
}

// AddCompany adds a single item to the Companies slice
func (b *OrganizationBuilder) AddCompany(company *CompanyBuilder) *OrganizationBuilder {
	// Ensure the slice is initialized
	if b.organization.Companies == nil {
		b.organization.Companies = []*models.Company{}
	}
	// Handle nested slice element
	builtValue := company.Build().(*models.Company)
	b.organization.Companies = append(b.organization.Companies, builtValue)
	return b
}
// AddDepartment adds a single item to the Departments slice
func (b *OrganizationBuilder) AddDepartment(department *DepartmentBuilder) *OrganizationBuilder {
	// Ensure the slice is initialized
	if b.organization.Departments == nil {
		b.organization.Departments = []*models.Department{}
	}
	// Handle nested slice element
	builtValue := department.Build().(*models.Department)
	b.organization.Departments = append(b.organization.Departments, builtValue)
	return b
}
// AddEmployee adds a single item to the Employees slice
func (b *OrganizationBuilder) AddEmployee(employee *EmployeeBuilder) *OrganizationBuilder {
	// Ensure the slice is initialized
	if b.organization.Employees == nil {
		b.organization.Employees = []*models.Employee{}
	}
	// Handle nested slice element
	builtValue := employee.Build().(*models.Employee)
	b.organization.Employees = append(b.organization.Employees, builtValue)
	return b
}
// AddProject adds a single item to the Projects slice
func (b *OrganizationBuilder) AddProject(project *ProjectBuilder) *OrganizationBuilder {
	// Ensure the slice is initialized
	if b.organization.Projects == nil {
		b.organization.Projects = []*models.Project{}
	}
	// Handle nested slice element
	builtValue := project.Build().(*models.Project)
	b.organization.Projects = append(b.organization.Projects, builtValue)
	return b
}
// AddContact adds a single item to the Contacts slice
func (b *OrganizationBuilder) AddContact(contact *ContactBuilder) *OrganizationBuilder {
	// Ensure the slice is initialized
	if b.organization.Contacts == nil {
		b.organization.Contacts = []*models.Contact{}
	}
	// Handle nested slice element
	builtValue := contact.Build().(*models.Contact)
	b.organization.Contacts = append(b.organization.Contacts, builtValue)
	return b
}

// WithValidation adds a custom validation function
func (b *OrganizationBuilder) WithValidation(validationFunc func(*models.Organization) error) *OrganizationBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Organization
func (b *OrganizationBuilder) Build() interface{} {
	return b.organization
}

// BuildPtr builds the Organization and returns a pointer
func (b *OrganizationBuilder) BuildPtr() *models.Organization {
	return b.organization
}

// BuildAndValidate builds the Organization and validates it
func (b *OrganizationBuilder) BuildAndValidate() (*models.Organization, error) {
	organization := b.organization

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(organization); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(organization).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return organization, err
		}
	}

	return organization, nil
}

// MustBuild builds the Organization and panics if validation fails
func (b *OrganizationBuilder) MustBuild() *models.Organization {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *OrganizationBuilder) Clone() *OrganizationBuilder {
	clonedOrganization := *b.organization
	return &OrganizationBuilder{
		organization: &clonedOrganization,
		validationFuncs: append([]func(*models.Organization) error{}, b.validationFuncs...),
	}
}
