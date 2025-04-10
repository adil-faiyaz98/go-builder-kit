package builders

import (
	"testing"
)

func TestBuilders(t *testing.T) {
	// Create a geo location
	geoLocationBuilder := NewGeoLocationBuilder().
		WithLatitude(37.7749).
		WithLongitude(-122.4194).
		WithAccuracy(10)

	// Create an address
	addressBuilder := NewAddressBuilder().
		WithStreet("123 Main St").
		WithCity("San Francisco").
		WithState("CA").
		WithPostalCode("94105").
		WithCountry("USA").
		WithCoordinates(geoLocationBuilder)

	// Create a course
	courseBuilder := NewCourseBuilder().
		WithName("Computer Science 101").
		WithCode("CS101").
		WithDescription("Introduction to Computer Science").
		WithCredits(3)

	// Create an education
	educationBuilder := NewEducationBuilder().
		WithDegree("Bachelor of Science").
		WithInstitution("Stanford University").
		WithLocation(addressBuilder).
		WithStartDate("2018-09-01").
		WithEndDate("2022-06-15").
		WithGPA(3.8).
		WithMajor("Computer Science").
		WithMinor("Mathematics").
		WithHonors([]string{"Magna Cum Laude", "Dean's List"}).
		WithCourses([]*CourseBuilder{courseBuilder})

	// Create a person
	personBuilder := NewPersonBuilder().
		WithName("John Doe").
		WithEmail("john.doe@example.com").
		WithPhone("555-1234").
		WithAddress(addressBuilder).
		WithEducation(educationBuilder)

	// Build the person
	person := personBuilder.BuildPtr()

	// Verify fields
	if person.Name != "John Doe" {
		t.Errorf("Expected Name to be 'John Doe', got '%s'", person.Name)
	}

	if person.Email != "john.doe@example.com" {
		t.Errorf("Expected Email to be 'john.doe@example.com', got '%s'", person.Email)
	}

	if person.Address.City != "San Francisco" {
		t.Errorf("Expected Address.City to be 'San Francisco', got '%s'", person.Address.City)
	}

	if person.Education.Degree != "Bachelor of Science" {
		t.Errorf("Expected Education.Degree to be 'Bachelor of Science', got '%s'", person.Education.Degree)
	}

	if len(person.Education.Honors) != 2 {
		t.Errorf("Expected 2 honors, got %d", len(person.Education.Honors))
	}

	if len(person.Education.Courses) != 1 {
		t.Errorf("Expected 1 course, got %d", len(person.Education.Courses))
	}
}
