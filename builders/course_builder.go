package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/v2/models"
	"github.com/adil-faiyaz98/go-builder-kit/v2/pkg/builder"
)

// CourseBuilder builds a Course model
type CourseBuilder struct {
	course *models.Course
	// Custom validation functions
	validationFuncs []func(*models.Course) error
}

// NewCourseBuilder creates a new CourseBuilder
func NewCourseBuilder() *CourseBuilder {
	return &CourseBuilder{
		course: &models.Course{
			Code:        "",
			Name:        "",
			Description: "",
			Credits:     0.0,
			Grade:       "",
			Semester:    "",
			Year:        0,
			Instructor:  "",
		},
		validationFuncs: []func(*models.Course) error{},
	}
}

// NewCourseBuilderWithDefaults creates a new CourseBuilder with sensible defaults
func NewCourseBuilderWithDefaults() *CourseBuilder {
	builder := NewCourseBuilder()
	// Add default values here if needed
	return builder
}

// WithCode sets the Code
func (b *CourseBuilder) WithCode(code string) *CourseBuilder {
	if b == nil {
		return b
	}
	b.course.Code = builder.SanitizeString(code)
	return b
}

// WithName sets the Name
func (b *CourseBuilder) WithName(name string) *CourseBuilder {
	if b == nil {
		return b
	}
	b.course.Name = builder.SanitizeString(name)
	return b
}

// WithDescription sets the Description
func (b *CourseBuilder) WithDescription(description string) *CourseBuilder {
	if b == nil {
		return b
	}
	b.course.Description = builder.SanitizeString(description)
	return b
}

// WithCredits sets the Credits
func (b *CourseBuilder) WithCredits(credits float64) *CourseBuilder {
	if b == nil {
		return b
	}
	b.course.Credits = credits
	return b
}

// WithGrade sets the Grade
func (b *CourseBuilder) WithGrade(grade string) *CourseBuilder {
	if b == nil {
		return b
	}
	b.course.Grade = builder.SanitizeString(grade)
	return b
}

// WithSemester sets the Semester
func (b *CourseBuilder) WithSemester(semester string) *CourseBuilder {
	if b == nil {
		return b
	}
	b.course.Semester = builder.SanitizeString(semester)
	return b
}

// WithYear sets the Year
func (b *CourseBuilder) WithYear(year int) *CourseBuilder {
	if b == nil {
		return b
	}
	b.course.Year = year
	return b
}

// WithInstructor sets the Instructor
func (b *CourseBuilder) WithInstructor(instructor string) *CourseBuilder {
	if b == nil {
		return b
	}
	b.course.Instructor = builder.SanitizeString(instructor)
	return b
}

// WithValidation adds a custom validation function
func (b *CourseBuilder) WithValidation(validationFunc func(*models.Course) error) *CourseBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Course
func (b *CourseBuilder) Build() interface{} {
	return b.course
}

// BuildPtr builds the Course and returns a pointer
func (b *CourseBuilder) BuildPtr() *models.Course {
	return b.course
}

// BuildAndValidate builds the Course and validates it
func (b *CourseBuilder) BuildAndValidate() (*models.Course, error) {
	if b == nil || b.course == nil {
		return nil, fmt.Errorf("builder is not properly initialized")
	}

	course := b.course

	// Run custom validation functions
	for i, validationFunc := range b.validationFuncs {
		if validationFunc == nil {
			continue // Skip nil validators
		}
		if err := validationFunc(course); err != nil {
			return nil, fmt.Errorf("custom validation failed at index %d: %w", i, err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := any(course).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return course, fmt.Errorf("model validation failed: %w", err)
		}
	}

	return course, nil
}

// MustBuild builds the Course and panics if validation fails
func (b *CourseBuilder) MustBuild() *models.Course {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *CourseBuilder) Clone() *CourseBuilder {
	if b == nil || b.course == nil {
		return NewCourseBuilder()
	}

	// Deep copy the struct
	clonedCourse := *b.course

	// Create new builder with cloned data
	clonedBuilder := &CourseBuilder{
		course:          &clonedCourse,
		validationFuncs: make([]func(*models.Course) error, 0, len(b.validationFuncs)),
	}

	// Copy validation functions safely
	if len(b.validationFuncs) > 0 {
		clonedBuilder.validationFuncs = append(clonedBuilder.validationFuncs, b.validationFuncs...)
	}

	return clonedBuilder
}
