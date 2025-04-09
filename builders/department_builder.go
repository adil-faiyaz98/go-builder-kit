package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// DepartmentBuilder builds a Department model
type DepartmentBuilder struct {
	department *models.Department
	// Custom validation functions
	validationFuncs []func(*models.Department) error
}

// NewDepartmentBuilder creates a new DepartmentBuilder
func NewDepartmentBuilder() *DepartmentBuilder {
	return &DepartmentBuilder{
		department: &models.Department{
			Name: "",
			Code: "",
			Description: "",
			Manager: any(0),
			Employees: []any{},
			Budget: 0.0,
			HeadCount: 0,
			Projects: []*models.Project{},
			Location: nil,
		},
		validationFuncs: []func(*models.Department) error{},
	}
}

// NewDepartmentBuilderWithDefaults creates a new DepartmentBuilder with sensible defaults
func NewDepartmentBuilderWithDefaults() *DepartmentBuilder {
	builder := NewDepartmentBuilder()
	// Add default values here if needed
	return builder
}
// WithName sets the Name
func (b *DepartmentBuilder) WithName(name string) *DepartmentBuilder {
	b.department.Name = name
	return b
}

// WithCode sets the Code
func (b *DepartmentBuilder) WithCode(code string) *DepartmentBuilder {
	b.department.Code = code
	return b
}

// WithDescription sets the Description
func (b *DepartmentBuilder) WithDescription(description string) *DepartmentBuilder {
	b.department.Description = description
	return b
}

// WithManager sets the Manager
func (b *DepartmentBuilder) WithManager(manager any) *DepartmentBuilder {
	b.department.Manager = manager
	return b
}

// WithEmployees sets the Employees
func (b *DepartmentBuilder) WithEmployees(employees any) *DepartmentBuilder {
	b.department.Employees = append(b.department.Employees, employees)
	return b
}

// WithBudget sets the Budget
func (b *DepartmentBuilder) WithBudget(budget float64) *DepartmentBuilder {
	b.department.Budget = budget
	return b
}

// WithHeadCount sets the HeadCount
func (b *DepartmentBuilder) WithHeadCount(headCount int) *DepartmentBuilder {
	b.department.HeadCount = headCount
	return b
}

// WithProjects sets the Projects
func (b *DepartmentBuilder) WithProjects(projects *ProjectBuilder) *DepartmentBuilder {
	// Ensure the slice is initialized
	if b.department.Projects == nil {
		b.department.Projects = []*models.Project{}
	}
	// Handle nested slice element
	builtValue := projects.Build().(*models.Project)
	b.department.Projects = append(b.department.Projects, builtValue)
	return b
}

// WithLocation sets the Location
func (b *DepartmentBuilder) WithLocation(location *AddressBuilder) *DepartmentBuilder {
	// Handle nested pointer
	b.department.Location = location.BuildPtr()
	return b
}


// WithValidation adds a custom validation function
func (b *DepartmentBuilder) WithValidation(validationFunc func(*models.Department) error) *DepartmentBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Department
func (b *DepartmentBuilder) Build() interface{} {
	return b.department
}

// BuildPtr builds the Department and returns a pointer
func (b *DepartmentBuilder) BuildPtr() *models.Department {
	return b.department
}

// BuildAndValidate builds the Department and validates it
func (b *DepartmentBuilder) BuildAndValidate() (*models.Department, error) {
	department := b.department

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(department); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(department).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return department, err
		}
	}

	return department, nil
}

// MustBuild builds the Department and panics if validation fails
func (b *DepartmentBuilder) MustBuild() *models.Department {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *DepartmentBuilder) Clone() *DepartmentBuilder {
	clonedDepartment := *b.department
	return &DepartmentBuilder{
		department: &clonedDepartment,
		validationFuncs: append([]func(*models.Department) error{}, b.validationFuncs...),
	}
}
