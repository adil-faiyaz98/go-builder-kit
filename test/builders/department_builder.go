package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/test/models"
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
			Description: "",
			Manager: nil,
			Employees: []*models.Employee{},
			SubDepts: []*models.Department{},
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

// WithDescription sets the Description
func (b *DepartmentBuilder) WithDescription(description string) *DepartmentBuilder {
	b.department.Description = description
	return b
}

// WithManager sets the Manager
func (b *DepartmentBuilder) WithManager(manager *EmployeeBuilder) *DepartmentBuilder {
	// Handle nested pointer
	b.department.Manager = manager.BuildPtr()
	return b
}

// WithEmployees sets the Employees
func (b *DepartmentBuilder) WithEmployees(employees []*EmployeeBuilder) *DepartmentBuilder {
	// Ensure the slice is initialized
	if b.department.Employees == nil {
		b.department.Employees = []*models.Employee{}
	}
	// Handle nested slice elements
	// Initialize the slice
	b.department.Employees = make([]*models.Employee, 0, len(employees))
	// Convert each builder to its model
	for _, builder := range employees {
		builtValue := builder.Build().(*models.Employee)
		b.department.Employees = append(b.department.Employees, builtValue)
	}
	return b
}

// WithSubDepts sets the SubDepts
func (b *DepartmentBuilder) WithSubDepts(subDepts []*DepartmentBuilder) *DepartmentBuilder {
	// Ensure the slice is initialized
	if b.department.SubDepts == nil {
		b.department.SubDepts = []*models.Department{}
	}
	// Handle nested slice elements
	// Initialize the slice
	b.department.SubDepts = make([]*models.Department, 0, len(subDepts))
	// Convert each builder to its model
	for _, builder := range subDepts {
		builtValue := builder.Build().(*models.Department)
		b.department.SubDepts = append(b.department.SubDepts, builtValue)
	}
	return b
}

// WithLocation sets the Location
func (b *DepartmentBuilder) WithLocation(location *AddressBuilder) *DepartmentBuilder {
	// Handle nested pointer
	b.department.Location = location.BuildPtr()
	return b
}

// AddEmployee adds a single item to the Employees slice
func (b *DepartmentBuilder) AddEmployee(employee *EmployeeBuilder) *DepartmentBuilder {
	// Ensure the slice is initialized
	if b.department.Employees == nil {
		b.department.Employees = []*models.Employee{}
	}
	// Handle nested slice element
	builtValue := employee.Build().(*models.Employee)
	b.department.Employees = append(b.department.Employees, builtValue)
	return b
}
// AddSubDept adds a single item to the SubDepts slice
func (b *DepartmentBuilder) AddSubDept(subDept *DepartmentBuilder) *DepartmentBuilder {
	// Ensure the slice is initialized
	if b.department.SubDepts == nil {
		b.department.SubDepts = []*models.Department{}
	}
	// Handle nested slice element
	builtValue := subDept.Build().(*models.Department)
	b.department.SubDepts = append(b.department.SubDepts, builtValue)
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
