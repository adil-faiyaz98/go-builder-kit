// +build ignore

// +build ignore

import (
	"testing"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
)

func TestEducationBuilder(t *testing.T) {
	// Create a GeoLocation builder
	geoLocationBuilder := builders.NewGeoLocationBuilder().
		WithLatitude(37.7749).
		WithLongitude(-122.4194).
		WithAccuracy(10.0)

	// Create an Address builder
	addressBuilder := builders.NewAddressBuilder().
		WithStreet("123 Main St").
		WithCity("San Francisco").
		WithState("CA").
		WithPostalCode("94105").
		WithCountry("USA").
		WithCoordinates(geoLocationBuilder).
		WithType("Home").
		WithIsPrimary(true)

	// Create a Course builder
	courseBuilder := builders.NewCourseBuilder().
		WithCode("CS101").
		WithName("Introduction to Computer Science").
		WithDescription("Basic concepts of computer science").
		WithCredits(3.0).
		WithGrade("A").
		WithSemester("Fall").
		WithYear(2020).
		WithInstructor("Dr. Smith")

	// Create an Education builder
	educationBuilder := builders.NewEducationBuilder().
		WithDegree("Bachelor of Science").
		WithInstitution("University of Example").
		WithLocation(addressBuilder).
		WithStartDate("2018-09-01").
		WithEndDate("2022-05-31").
		WithGPA(3.8).
		WithMajor("Computer Science").
		WithMinor("Mathematics")

	// Add honors
	educationBuilder.AddHonor("Dean's List")
	educationBuilder.AddHonor("Summa Cum Laude")

	// Add courses
	educationBuilder.AddCourse(courseBuilder)

	// Add activities
	educationBuilder.AddActivity("Chess Club")
	educationBuilder.AddActivity("Coding Competition")

	// Build the Education
	education, err := educationBuilder.BuildAndValidate()
	if err != nil {
		t.Fatalf("Failed to build education: %v", err)
	}

	// Verify Education fields
	if education.Degree != "Bachelor of Science" {
		t.Errorf("Expected degree to be 'Bachelor of Science', got '%s'", education.Degree)
	}

	if education.Institution != "University of Example" {
		t.Errorf("Expected institution to be 'University of Example', got '%s'", education.Institution)
	}

	if education.GPA != 3.8 {
		t.Errorf("Expected GPA to be 3.8, got %f", education.GPA)
	}

	// Verify Address fields
	if education.Location == nil {
		t.Fatalf("Expected location to be non-nil")
	}

	if education.Location.City != "San Francisco" {
		t.Errorf("Expected city to be 'San Francisco', got '%s'", education.Location.City)
	}

	// Verify GeoLocation fields
	if education.Location.Coordinates == nil {
		t.Fatalf("Expected coordinates to be non-nil")
	}

	if education.Location.Coordinates.Latitude != 37.7749 {
		t.Errorf("Expected latitude to be 37.7749, got %f", education.Location.Coordinates.Latitude)
	}

	// Verify Course fields
	if len(education.Courses) != 1 {
		t.Fatalf("Expected 1 course, got %d", len(education.Courses))
	}

	if education.Courses[0].Code != "CS101" {
		t.Errorf("Expected course code to be 'CS101', got '%s'", education.Courses[0].Code)
	}

	// Verify Honors fields
	if len(education.Honors) != 2 {
		t.Fatalf("Expected 2 honors, got %d", len(education.Honors))
	}

	if education.Honors[0] != "Dean's List" {
		t.Errorf("Expected first honor to be 'Dean's List', got '%s'", education.Honors[0])
	}

	// Verify Activities fields
	if len(education.Activities) != 2 {
		t.Fatalf("Expected 2 activities, got %d", len(education.Activities))
	}

	if education.Activities[0] != "Chess Club" {
		t.Errorf("Expected first activity to be 'Chess Club', got '%s'", education.Activities[0])
	}
}
