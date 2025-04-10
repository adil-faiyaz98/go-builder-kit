package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

// CourseBuilder builds a Course model
type CourseBuilder struct {
	course         *models.Course
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

// WithCode sets the Code
func (b *CourseBuilder) WithCode(code string) *CourseBuilder {
	b.course.Code = code
	return b
}

// WithName sets the Name
func (b *CourseBuilder) WithName(name string) *CourseBuilder {
	b.course.Name = name
	return b
}

// WithDescription sets the Description
func (b *CourseBuilder) WithDescription(description string) *CourseBuilder {
	b.course.Description = description
	return b
}

// WithCredits sets the Credits
func (b *CourseBuilder) WithCredits(credits float64) *CourseBuilder {
	b.course.Credits = credits
	return b
}

// WithGrade sets the Grade
func (b *CourseBuilder) WithGrade(grade string) *CourseBuilder {
	b.course.Grade = grade
	return b
}

// WithSemester sets the Semester
func (b *CourseBuilder) WithSemester(semester string) *CourseBuilder {
	b.course.Semester = semester
	return b
}

// WithYear sets the Year
func (b *CourseBuilder) WithYear(year int) *CourseBuilder {
	b.course.Year = year
	return b
}

// WithInstructor sets the Instructor
func (b *CourseBuilder) WithInstructor(instructor string) *CourseBuilder {
	b.course.Instructor = instructor
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
	course := b.course

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(course); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method
	if err := course.Validate(); err != nil {
		return course, err
	}

	return course, nil
}

// MustBuild builds the Course and panics if validation fails
func (b *CourseBuilder) MustBuild() *models.Course {
	course, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return course
}

// Clone creates a deep copy of the builder
func (b *CourseBuilder) Clone() *CourseBuilder {
	clonedCourse := *b.course
	return &CourseBuilder{
		course:         &clonedCourse,
		validationFuncs: append([]func(*models.Course) error{}, b.validationFuncs...),
	}
}
