package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/models"
	"github.com/adil-faiyaz98/go-builder-kit/pkg/builder"
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
			Name:        "",
			Code:        "",
			Description: "",
			Manager:     interface{}(0),
			Employees:   []interface{}{},
			Budget:      0.0,
			HeadCount:   0,
			Projects:    []interface{}{},
			Location:    nil,
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
	if b == nil {
		return b
	}
	b.department.Name = builder.SanitizeString(name)
	return b
}

// WithCode sets the Code
func (b *DepartmentBuilder) WithCode(code string) *DepartmentBuilder {
	if b == nil {
		return b
	}
	b.department.Code = builder.SanitizeString(code)
	return b
}

// WithDescription sets the Description
func (b *DepartmentBuilder) WithDescription(description string) *DepartmentBuilder {
	if b == nil {
		return b
	}
	b.department.Description = builder.SanitizeString(description)
	return b
}

// WithManager sets the Manager
func (b *DepartmentBuilder) WithManager(manager interface{}) *DepartmentBuilder {
	if b == nil {
		return b
	}
	b.department.Manager = manager
	return b
}

// WithEmployees sets the Employees
func (b *DepartmentBuilder) WithEmployees(employees []interface{}) *DepartmentBuilder {
	if b == nil {
		return b
	}
	b.department.Employees = append(b.department.Employees, employees...)
	return b
}

// WithBudget sets the Budget
func (b *DepartmentBuilder) WithBudget(budget float64) *DepartmentBuilder {
	if b == nil {
		return b
	}
	b.department.Budget = budget
	return b
}

// WithHeadCount sets the HeadCount
func (b *DepartmentBuilder) WithHeadCount(headCount int) *DepartmentBuilder {
	if b == nil {
		return b
	}
	b.department.HeadCount = headCount
	return b
}

// WithProjects sets the Projects
func (b *DepartmentBuilder) WithProjects(projects []interface{}) *DepartmentBuilder {
	if b == nil {
		return b
	}
	b.department.Projects = append(b.department.Projects, projects...)
	return b
}

// WithLocation sets the Location
func (b *DepartmentBuilder) WithLocation(location *AddressBuilder) *DepartmentBuilder {
	if b == nil {
		return b
	}
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
	if b == nil || b.department == nil {
		return nil, fmt.Errorf("builder is not properly initialized")
	}

	department := b.department

	// Run custom validation functions
	for i, validationFunc := range b.validationFuncs {
		if validationFunc == nil {
			continue // Skip nil validators
		}
		if err := validationFunc(department); err != nil {
			return nil, fmt.Errorf("custom validation failed at index %d: %w", i, err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := any(department).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return department, fmt.Errorf("model validation failed: %w", err)
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
	if b == nil || b.department == nil {
		return NewDepartmentBuilder()
	}

	// Deep copy the struct
	clonedDepartment := *b.department

	// Create new builder with cloned data
	clonedBuilder := &DepartmentBuilder{
		department:      &clonedDepartment,
		validationFuncs: make([]func(*models.Department) error, 0, len(b.validationFuncs)),
	}

	// Copy validation functions safely
	if len(b.validationFuncs) > 0 {
		clonedBuilder.validationFuncs = append(clonedBuilder.validationFuncs, b.validationFuncs...)
	}

	return clonedBuilder
}
