package builders

import (
	"fmt"

	"github.com/adil-faiyaz98/go-builder-kit/models"
	"github.com/adil-faiyaz98/go-builder-kit/pkg/builder"
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
			Company:      nil,
			Position:     "",
			Department:   "",
			StartDate:    "",
			EndDate:      "",
			Salary:       0.0,
			IsCurrent:    false,
			Supervisor:   nil,
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
	if b == nil {
		return b
	}
	// Handle nested pointer
	b.employment.Company = company.BuildPtr()
	return b
}

// WithPosition sets the Position
func (b *EmploymentBuilder) WithPosition(position string) *EmploymentBuilder {
	if b == nil {
		return b
	}
	b.employment.Position = builder.SanitizeString(position)
	return b
}

// WithDepartment sets the Department
func (b *EmploymentBuilder) WithDepartment(department string) *EmploymentBuilder {
	if b == nil {
		return b
	}
	b.employment.Department = builder.SanitizeString(department)
	return b
}

// WithStartDate sets the StartDate
func (b *EmploymentBuilder) WithStartDate(startDate string) *EmploymentBuilder {
	if b == nil {
		return b
	}
	b.employment.StartDate = builder.SanitizeString(startDate)
	return b
}

// WithEndDate sets the EndDate
func (b *EmploymentBuilder) WithEndDate(endDate string) *EmploymentBuilder {
	if b == nil {
		return b
	}
	b.employment.EndDate = builder.SanitizeString(endDate)
	return b
}

// WithSalary sets the Salary
func (b *EmploymentBuilder) WithSalary(salary float64) *EmploymentBuilder {
	if b == nil {
		return b
	}
	b.employment.Salary = salary
	return b
}

// WithIsCurrent sets the IsCurrent
func (b *EmploymentBuilder) WithIsCurrent(isCurrent bool) *EmploymentBuilder {
	if b == nil {
		return b
	}
	b.employment.IsCurrent = isCurrent
	return b
}

// WithSupervisor sets the Supervisor
func (b *EmploymentBuilder) WithSupervisor(supervisor any) *EmploymentBuilder {
	if b == nil {
		return b
	}
	b.employment.Supervisor = supervisor
	return b
}

// WithSubordinates sets the Subordinates
func (b *EmploymentBuilder) WithSubordinates(subordinates []any) *EmploymentBuilder {
	if b == nil {
		return b
	}
	b.employment.Subordinates = append(b.employment.Subordinates, subordinates...)
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
	if b == nil || b.employment == nil {
		return nil, fmt.Errorf("builder is not properly initialized")
	}

	employment := b.employment

	// Run custom validation functions
	for i, validationFunc := range b.validationFuncs {
		if validationFunc == nil {
			continue // Skip nil validators
		}
		if err := validationFunc(employment); err != nil {
			return nil, fmt.Errorf("custom validation failed at index %d: %w", i, err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := any(employment).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return employment, fmt.Errorf("model validation failed: %w", err)
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
	if b == nil || b.employment == nil {
		return NewEmploymentBuilder()
	}

	// Deep copy the struct
	clonedEmployment := *b.employment

	// Create new builder with cloned data
	clonedBuilder := &EmploymentBuilder{
		employment:      &clonedEmployment,
		validationFuncs: make([]func(*models.Employment) error, 0, len(b.validationFuncs)),
	}

	// Copy validation functions safely
	if len(b.validationFuncs) > 0 {
		clonedBuilder.validationFuncs = append(clonedBuilder.validationFuncs, b.validationFuncs...)
	}

	return clonedBuilder
}
