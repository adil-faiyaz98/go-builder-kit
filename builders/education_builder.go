package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/models"
	"github.com/adil-faiyaz98/go-builder-kit/pkg/builder"
)

// EducationBuilder builds a Education model
type EducationBuilder struct {
	education *models.Education
	// Custom validation functions
	validationFuncs []func(*models.Education) error
}

// NewEducationBuilder creates a new EducationBuilder
func NewEducationBuilder() *EducationBuilder {
	return &EducationBuilder{
		education: &models.Education{
			Degree:      "",
			Institution: "",
			Location:    nil,
			StartDate:   "",
			EndDate:     "",
			GPA:         0.0,
			Honors:      []string{},
			Major:       "",
			Minor:       "",
			Courses:     []*models.Course{},
			Activities:  []string{},
		},
		validationFuncs: []func(*models.Education) error{},
	}
}

// NewEducationBuilderWithDefaults creates a new EducationBuilder with sensible defaults
func NewEducationBuilderWithDefaults() *EducationBuilder {
	builder := NewEducationBuilder()
	// Add default values here if needed
	return builder
}

// WithDegree sets the Degree
func (b *EducationBuilder) WithDegree(degree string) *EducationBuilder {
	if b == nil {
		return b
	}
	b.education.Degree = builder.SanitizeString(degree)
	return b
}

// WithInstitution sets the Institution
func (b *EducationBuilder) WithInstitution(institution string) *EducationBuilder {
	if b == nil {
		return b
	}
	b.education.Institution = builder.SanitizeString(institution)
	return b
}

// WithLocation sets the Location
func (b *EducationBuilder) WithLocation(location *AddressBuilder) *EducationBuilder {
	if b == nil {
		return b
	}
	// Handle nested pointer
	b.education.Location = location.BuildPtr()
	return b
}

// WithStartDate sets the StartDate
func (b *EducationBuilder) WithStartDate(startDate string) *EducationBuilder {
	if b == nil {
		return b
	}
	b.education.StartDate = builder.SanitizeString(startDate)
	return b
}

// WithEndDate sets the EndDate
func (b *EducationBuilder) WithEndDate(endDate string) *EducationBuilder {
	if b == nil {
		return b
	}
	b.education.EndDate = builder.SanitizeString(endDate)
	return b
}

// WithGPA sets the GPA
func (b *EducationBuilder) WithGPA(gPA float64) *EducationBuilder {
	if b == nil {
		return b
	}
	b.education.GPA = gPA
	return b
}

// WithHonors sets the Honors
func (b *EducationBuilder) WithHonors(honors []string) *EducationBuilder {
	if b == nil {
		return b
	}
	b.education.Honors = append(b.education.Honors, honors...)
	return b
}

// WithMajor sets the Major
func (b *EducationBuilder) WithMajor(major string) *EducationBuilder {
	if b == nil {
		return b
	}
	b.education.Major = builder.SanitizeString(major)
	return b
}

// WithMinor sets the Minor
func (b *EducationBuilder) WithMinor(minor string) *EducationBuilder {
	if b == nil {
		return b
	}
	b.education.Minor = builder.SanitizeString(minor)
	return b
}

// WithCourses sets the Courses
func (b *EducationBuilder) WithCourses(courses []*CourseBuilder) *EducationBuilder {
	if b == nil {
		return b
	}
	// Ensure the slice is initialized
	if b.education.Courses == nil {
		b.education.Courses = []*models.Course{}
	}
	// Handle nested slice elements
	// Initialize the slice
	b.education.Courses = make([]*models.Course, 0, len(courses))
	// Convert each builder to its model
	for _, builder := range courses {
		builtValue := builder.Build().(*models.Course)
		b.education.Courses = append(b.education.Courses, builtValue)
	}
	return b
}

// WithActivities sets the Activities
func (b *EducationBuilder) WithActivities(activities []string) *EducationBuilder {
	if b == nil {
		return b
	}
	b.education.Activities = append(b.education.Activities, activities...)
	return b
}

// AddCours adds a single item to the Courses slice
func (b *EducationBuilder) AddCours(cours *CourseBuilder) *EducationBuilder {
	if b == nil {
		return b
	}
	// Ensure the slice is initialized with capacity
	if b.education.Courses == nil {
		b.education.Courses = make([]*models.Course, 0, 4) // Pre-allocate capacity
	}
	// Handle nested slice element
	if cours != nil {
		builtValue := cours.Build().(*models.Course)
		b.education.Courses = append(b.education.Courses, builtValue)
	}
	return b
}

// WithValidation adds a custom validation function
func (b *EducationBuilder) WithValidation(validationFunc func(*models.Education) error) *EducationBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Education
func (b *EducationBuilder) Build() interface{} {
	return b.education
}

// BuildPtr builds the Education and returns a pointer
func (b *EducationBuilder) BuildPtr() *models.Education {
	return b.education
}

// BuildAndValidate builds the Education and validates it
func (b *EducationBuilder) BuildAndValidate() (*models.Education, error) {
	if b == nil || b.education == nil {
		return nil, fmt.Errorf("builder is not properly initialized")
	}

	education := b.education

	// Run custom validation functions
	for i, validationFunc := range b.validationFuncs {
		if validationFunc == nil {
			continue // Skip nil validators
		}
		if err := validationFunc(education); err != nil {
			return nil, fmt.Errorf("custom validation failed at index %d: %w", i, err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := any(education).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return education, fmt.Errorf("model validation failed: %w", err)
		}
	}

	return education, nil
}

// MustBuild builds the Education and panics if validation fails
func (b *EducationBuilder) MustBuild() *models.Education {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *EducationBuilder) Clone() *EducationBuilder {
	if b == nil || b.education == nil {
		return NewEducationBuilder()
	}

	// Deep copy the struct
	clonedEducation := *b.education

	// Create new builder with cloned data
	clonedBuilder := &EducationBuilder{
		education:       &clonedEducation,
		validationFuncs: make([]func(*models.Education) error, 0, len(b.validationFuncs)),
	}

	// Copy validation functions safely
	if len(b.validationFuncs) > 0 {
		clonedBuilder.validationFuncs = append(clonedBuilder.validationFuncs, b.validationFuncs...)
	}

	return clonedBuilder
}
