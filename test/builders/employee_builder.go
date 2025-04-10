package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/test/models"
)

// EmployeeBuilder builds a Employee model
type EmployeeBuilder struct {
	employee *models.Employee
	// Custom validation functions
	validationFuncs []func(*models.Employee) error
}

// NewEmployeeBuilder creates a new EmployeeBuilder
func NewEmployeeBuilder() *EmployeeBuilder {
	return &EmployeeBuilder{
		employee: &models.Employee{
			ID: "",
			Name: "",
			Title: "",
			Department: nil,
			Manager: nil,
			Subordinates: []*models.Employee{},
			Contact: nil,
			Projects: []*models.Project{},
			Skills: []string{},
			StartDate: "",
		},
		validationFuncs: []func(*models.Employee) error{},
	}
}

// NewEmployeeBuilderWithDefaults creates a new EmployeeBuilder with sensible defaults
func NewEmployeeBuilderWithDefaults() *EmployeeBuilder {
	builder := NewEmployeeBuilder()
	// Add default values here if needed
	return builder
}
// WithID sets the ID
func (b *EmployeeBuilder) WithID(id string) *EmployeeBuilder {
	b.employee.ID = id
	return b
}

// WithName sets the Name
func (b *EmployeeBuilder) WithName(name string) *EmployeeBuilder {
	b.employee.Name = name
	return b
}

// WithTitle sets the Title
func (b *EmployeeBuilder) WithTitle(title string) *EmployeeBuilder {
	b.employee.Title = title
	return b
}

// WithDepartment sets the Department
func (b *EmployeeBuilder) WithDepartment(department *DepartmentBuilder) *EmployeeBuilder {
	// Handle nested pointer
	b.employee.Department = department.BuildPtr()
	return b
}

// WithManager sets the Manager
func (b *EmployeeBuilder) WithManager(manager *EmployeeBuilder) *EmployeeBuilder {
	// Handle nested pointer
	b.employee.Manager = manager.BuildPtr()
	return b
}

// WithSubordinates sets the Subordinates
func (b *EmployeeBuilder) WithSubordinates(subordinates []*EmployeeBuilder) *EmployeeBuilder {
	// Ensure the slice is initialized
	if b.employee.Subordinates == nil {
		b.employee.Subordinates = []*models.Employee{}
	}
	// Handle nested slice elements
	// Initialize the slice
	b.employee.Subordinates = make([]*models.Employee, 0, len(subordinates))
	// Convert each builder to its model
	for _, builder := range subordinates {
		builtValue := builder.Build().(*models.Employee)
		b.employee.Subordinates = append(b.employee.Subordinates, builtValue)
	}
	return b
}

// WithContact sets the Contact
func (b *EmployeeBuilder) WithContact(contact *ContactBuilder) *EmployeeBuilder {
	// Handle nested pointer
	b.employee.Contact = contact.BuildPtr()
	return b
}

// WithProjects sets the Projects
func (b *EmployeeBuilder) WithProjects(projects []*ProjectBuilder) *EmployeeBuilder {
	// Ensure the slice is initialized
	if b.employee.Projects == nil {
		b.employee.Projects = []*models.Project{}
	}
	// Handle nested slice elements
	// Initialize the slice
	b.employee.Projects = make([]*models.Project, 0, len(projects))
	// Convert each builder to its model
	for _, builder := range projects {
		builtValue := builder.Build().(*models.Project)
		b.employee.Projects = append(b.employee.Projects, builtValue)
	}
	return b
}

// WithSkills sets the Skills
func (b *EmployeeBuilder) WithSkills(skills []string) *EmployeeBuilder {
	b.employee.Skills = append(b.employee.Skills, skills...)
	return b
}

// WithStartDate sets the StartDate
func (b *EmployeeBuilder) WithStartDate(startDate string) *EmployeeBuilder {
	b.employee.StartDate = startDate
	return b
}

// AddSubordinate adds a single item to the Subordinates slice
func (b *EmployeeBuilder) AddSubordinate(subordinate *EmployeeBuilder) *EmployeeBuilder {
	// Ensure the slice is initialized
	if b.employee.Subordinates == nil {
		b.employee.Subordinates = []*models.Employee{}
	}
	// Handle nested slice element
	builtValue := subordinate.Build().(*models.Employee)
	b.employee.Subordinates = append(b.employee.Subordinates, builtValue)
	return b
}
// AddProject adds a single item to the Projects slice
func (b *EmployeeBuilder) AddProject(project *ProjectBuilder) *EmployeeBuilder {
	// Ensure the slice is initialized
	if b.employee.Projects == nil {
		b.employee.Projects = []*models.Project{}
	}
	// Handle nested slice element
	builtValue := project.Build().(*models.Project)
	b.employee.Projects = append(b.employee.Projects, builtValue)
	return b
}

// WithValidation adds a custom validation function
func (b *EmployeeBuilder) WithValidation(validationFunc func(*models.Employee) error) *EmployeeBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Employee
func (b *EmployeeBuilder) Build() interface{} {
	return b.employee
}

// BuildPtr builds the Employee and returns a pointer
func (b *EmployeeBuilder) BuildPtr() *models.Employee {
	return b.employee
}

// BuildAndValidate builds the Employee and validates it
func (b *EmployeeBuilder) BuildAndValidate() (*models.Employee, error) {
	employee := b.employee

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(employee); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(employee).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return employee, err
		}
	}

	return employee, nil
}

// MustBuild builds the Employee and panics if validation fails
func (b *EmployeeBuilder) MustBuild() *models.Employee {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *EmployeeBuilder) Clone() *EmployeeBuilder {
	clonedEmployee := *b.employee
	return &EmployeeBuilder{
		employee: &clonedEmployee,
		validationFuncs: append([]func(*models.Employee) error{}, b.validationFuncs...),
	}
}
