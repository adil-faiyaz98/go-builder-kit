package builders

import (
	"fmt"

	"github.com/adil-faiyaz98/go-builder-kit/models"
)

// EducationBuilder builds an Education model
type EducationBuilder struct {
	education       *models.Education
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

// WithDegree sets the Degree
func (b *EducationBuilder) WithDegree(degree string) *EducationBuilder {
	b.education.Degree = degree
	return b
}

// WithInstitution sets the Institution
func (b *EducationBuilder) WithInstitution(institution string) *EducationBuilder {
	b.education.Institution = institution
	return b
}

// WithLocation sets the Location
func (b *EducationBuilder) WithLocation(location *AddressBuilder) *EducationBuilder {
	b.education.Location = location.BuildPtr()
	return b
}

// WithStartDate sets the StartDate
func (b *EducationBuilder) WithStartDate(startDate string) *EducationBuilder {
	b.education.StartDate = startDate
	return b
}

// WithEndDate sets the EndDate
func (b *EducationBuilder) WithEndDate(endDate string) *EducationBuilder {
	b.education.EndDate = endDate
	return b
}

// WithGPA sets the GPA
func (b *EducationBuilder) WithGPA(gpa float64) *EducationBuilder {
	b.education.GPA = gpa
	return b
}

// WithHonors sets the Honors
func (b *EducationBuilder) WithHonors(honors []string) *EducationBuilder {
	b.education.Honors = honors
	return b
}

// AddHonor adds an honor to the Honors slice
func (b *EducationBuilder) AddHonor(honor string) *EducationBuilder {
	b.education.Honors = append(b.education.Honors, honor)
	return b
}

// WithMajor sets the Major
func (b *EducationBuilder) WithMajor(major string) *EducationBuilder {
	b.education.Major = major
	return b
}

// WithMinor sets the Minor
func (b *EducationBuilder) WithMinor(minor string) *EducationBuilder {
	b.education.Minor = minor
	return b
}

// WithCourses sets the Courses
func (b *EducationBuilder) WithCourses(courses []*models.Course) *EducationBuilder {
	b.education.Courses = courses
	return b
}

// AddCourse adds a course to the Courses slice
func (b *EducationBuilder) AddCourse(course *CourseBuilder) *EducationBuilder {
	b.education.Courses = append(b.education.Courses, course.BuildPtr())
	return b
}

// WithActivities sets the Activities
func (b *EducationBuilder) WithActivities(activities []string) *EducationBuilder {
	b.education.Activities = activities
	return b
}

// AddActivity adds an activity to the Activities slice
func (b *EducationBuilder) AddActivity(activity string) *EducationBuilder {
	b.education.Activities = append(b.education.Activities, activity)
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
	education := b.education

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(education); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method
	if err := education.Validate(); err != nil {
		return education, err
	}

	return education, nil
}

// MustBuild builds the Education and panics if validation fails
func (b *EducationBuilder) MustBuild() *models.Education {
	education, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return education
}

// Clone creates a deep copy of the builder
func (b *EducationBuilder) Clone() *EducationBuilder {
	clonedEducation := *b.education
	return &EducationBuilder{
		education:       &clonedEducation,
		validationFuncs: append([]func(*models.Education) error{}, b.validationFuncs...),
	}
}
