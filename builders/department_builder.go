package builders

import (
	"fmt"

	"github.com/adil-faiyaz98/go-builder-kit/models"
)

// DepartmentBuilder builds a Department model
type DepartmentBuilder struct {
	department      *models.Department
	validationFuncs []func(*models.Department) error
}

// NewDepartmentBuilder creates a new DepartmentBuilder
func NewDepartmentBuilder() *DepartmentBuilder {
	return &DepartmentBuilder{
		department: &models.Department{
			Name:        "",
			Code:        "",
			Description: "",
			Manager:     nil,
			Employees:   []any{},
			Budget:      0.0,
			HeadCount:   0,
			Location:    nil,
			Projects:    []any{},
		},
		validationFuncs: []func(*models.Department) error{},
	}
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

// WithEmployee adds an employee to the Employees slice
func (b *DepartmentBuilder) WithEmployee(employee any) *DepartmentBuilder {
	b.department.Employees = append(b.department.Employees, employee)
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

// WithLocation sets the Location
func (b *DepartmentBuilder) WithLocation(location *AddressBuilder) *DepartmentBuilder {
	builtValue := location.Build().(*models.Address)
	b.department.Location = builtValue
	return b
}

// WithProject adds a project to the Projects slice
func (b *DepartmentBuilder) WithProject(project any) *DepartmentBuilder {
	b.department.Projects = append(b.department.Projects, project)
	return b
}

// WithValidation adds a custom validation function
func (b *DepartmentBuilder) WithValidation(validationFunc func(*models.Department) error) *DepartmentBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Department
func (b *DepartmentBuilder) Build() any {
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
			return department, err
		}
	}

	// Run model's Validate method
	if err := department.Validate(); err != nil {
		return department, err
	}

	return department, nil
}

// MustBuild builds the Department and panics if validation fails
func (b *DepartmentBuilder) MustBuild() *models.Department {
	department, err := b.BuildAndValidate()
	if err != nil {
		panic(fmt.Sprintf("Department validation failed: %s", err.Error()))
	}
	return department
}

// Clone creates a deep copy of the DepartmentBuilder
func (b *DepartmentBuilder) Clone() *DepartmentBuilder {
	clonedDepartment := *b.department
	return &DepartmentBuilder{
		department:      &clonedDepartment,
		validationFuncs: append([]func(*models.Department) error{}, b.validationFuncs...),
	}
}
