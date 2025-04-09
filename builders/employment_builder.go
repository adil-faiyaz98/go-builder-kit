package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// EmploymentBuilder builds a Employment model
type EmploymentBuilder struct {
	employment *models.Employment
	// Custom validation functions
	validationFuncs []func(*models.Employment) error
}

// NewEmploymentBuilder creates a new EmploymentBuilder
func NewEmploymentBuilder() *EmploymentBuilder {
	return &EmploymentBuilder{
		employment: &models.Employment{
			Company: nil,
			Position: "",
			Department: "",
			StartDate: "",
			EndDate: "",
			Salary: 0.0,
			IsCurrent: false,
			Supervisor: any(0),
			Subordinates: []any{},
		},
		validationFuncs: []func(*models.Employment) error{},
	}
}

// NewEmploymentBuilderWithDefaults creates a new EmploymentBuilder with sensible defaults
func NewEmploymentBuilderWithDefaults() *EmploymentBuilder {
	builder := NewEmploymentBuilder()
	// Add default values here if needed
	return builder
}
// WithCompany sets the Company
func (b *EmploymentBuilder) WithCompany(company *CompanyBuilder) *EmploymentBuilder {
	// Handle nested pointer
	b.employment.Company = company.BuildPtr()
	return b
}

// WithPosition sets the Position
func (b *EmploymentBuilder) WithPosition(position string) *EmploymentBuilder {
	b.employment.Position = position
	return b
}

// WithDepartment sets the Department
func (b *EmploymentBuilder) WithDepartment(department string) *EmploymentBuilder {
	b.employment.Department = department
	return b
}

// WithStartDate sets the StartDate
func (b *EmploymentBuilder) WithStartDate(startDate string) *EmploymentBuilder {
	b.employment.StartDate = startDate
	return b
}

// WithEndDate sets the EndDate
func (b *EmploymentBuilder) WithEndDate(endDate string) *EmploymentBuilder {
	b.employment.EndDate = endDate
	return b
}

// WithSalary sets the Salary
func (b *EmploymentBuilder) WithSalary(salary float64) *EmploymentBuilder {
	b.employment.Salary = salary
	return b
}

// WithIsCurrent sets the IsCurrent
func (b *EmploymentBuilder) WithIsCurrent(isCurrent bool) *EmploymentBuilder {
	b.employment.IsCurrent = isCurrent
	return b
}

// WithSupervisor sets the Supervisor
func (b *EmploymentBuilder) WithSupervisor(supervisor any) *EmploymentBuilder {
	b.employment.Supervisor = supervisor
	return b
}

// WithSubordinates sets the Subordinates
func (b *EmploymentBuilder) WithSubordinates(subordinates any) *EmploymentBuilder {
	b.employment.Subordinates = append(b.employment.Subordinates, subordinates)
	return b
}


// WithValidation adds a custom validation function
func (b *EmploymentBuilder) WithValidation(validationFunc func(*models.Employment) error) *EmploymentBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Employment
func (b *EmploymentBuilder) Build() interface{} {
	return b.employment
}

// BuildPtr builds the Employment and returns a pointer
func (b *EmploymentBuilder) BuildPtr() *models.Employment {
	return b.employment
}

// BuildAndValidate builds the Employment and validates it
func (b *EmploymentBuilder) BuildAndValidate() (*models.Employment, error) {
	employment := b.employment

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(employment); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(employment).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return employment, err
		}
	}

	return employment, nil
}

// MustBuild builds the Employment and panics if validation fails
func (b *EmploymentBuilder) MustBuild() *models.Employment {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *EmploymentBuilder) Clone() *EmploymentBuilder {
	clonedEmployment := *b.employment
	return &EmploymentBuilder{
		employment: &clonedEmployment,
		validationFuncs: append([]func(*models.Employment) error{}, b.validationFuncs...),
	}
}
