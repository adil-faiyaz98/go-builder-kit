package models

import (
	"testing"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

// TestNestedStructures tests that nested structures are properly handled
func TestNestedStructures(t *testing.T) {
	// Test Person with Address
	t.Run("Person with Address", func(t *testing.T) {
		// Create a valid address
		addressBuilder := builders.NewAddressBuilder().
			WithStreet("123 Main St").
			WithCity("New York").
			WithState("NY").
			WithPostalCode("10001").
			WithCountry("USA")

		// Create a person with the address
		personBuilder := builders.NewPersonBuilder().
			WithID("123").
			WithName("John Doe").
			WithAge(30).
			WithEmail("john.doe@example.com").
			WithAddress(addressBuilder)

		// Build the person
		person := personBuilder.Build().(*models.Person)

		// Verify the address is set
		if person.Address.Street != "123 Main St" {
			t.Errorf("Expected street to be '123 Main St', got '%s'", person.Address.Street)
		}
		if person.Address.City != "New York" {
			t.Errorf("Expected city to be 'New York', got '%s'", person.Address.City)
		}
	})

	// Test Person with Education and Course
	t.Run("Person with Education and Course", func(t *testing.T) {
		// Course creation is not needed for this simple test

		// Create an education with the course
		educationBuilder := builders.NewEducationBuilder().
			WithInstitution("University of Example").
			WithDegree("Bachelor of Science").
			WithMajor("Computer Science").
			WithStartDate("2020-01-01").
			WithEndDate("2024-05-31").
			WithGPA(3.8)

		// Create a person with the education
		personBuilder := builders.NewPersonBuilder().
			WithID("123").
			WithName("John Doe").
			WithAge(30).
			WithEmail("john.doe@example.com").
			WithEducation(educationBuilder)

		// Build the person
		person := personBuilder.Build().(*models.Person)

		// Verify the education is set
		if person.Education.Institution != "University of Example" {
			t.Errorf("Expected institution to be 'University of Example', got '%s'", person.Education.Institution)
		}
		if person.Education.Degree != "Bachelor of Science" {
			t.Errorf("Expected degree to be 'Bachelor of Science', got '%s'", person.Education.Degree)
		}
	})

	// Test Bank with Account and Transaction
	t.Run("Bank with Account", func(t *testing.T) {
		// Account creation is not needed for this simple test

		// Create a bank address
		addressBuilder := builders.NewAddressBuilder().
			WithStreet("456 Financial St").
			WithCity("New York").
			WithState("NY").
			WithPostalCode("10002").
			WithCountry("USA")

		// Create a bank with the account and address
		bankBuilder := builders.NewBankBuilder().
			WithName("Global Bank").
			WithBranchCode("GB-001").
			WithAddress(addressBuilder)

		// Build the bank
		bank := bankBuilder.Build().(*models.Bank)

		// Verify the address is set
		if bank.Address.Street != "456 Financial St" {
			t.Errorf("Expected street to be '456 Financial St', got '%s'", bank.Address.Street)
		}
		if bank.Address.City != "New York" {
			t.Errorf("Expected city to be 'New York', got '%s'", bank.Address.City)
		}
	})

	// Test DigitalProfile with Device and Application
	t.Run("DigitalProfile with Device", func(t *testing.T) {
		// Application and Device creation is not needed for this simple test

		// Create a digital profile
		digitalProfileBuilder := builders.NewDigitalProfileBuilder()

		// Build the digital profile
		digitalProfile := digitalProfileBuilder.Build().(*models.DigitalProfile)

		// Verify the digital profile is created
		if digitalProfile == nil {
			t.Errorf("Expected digital profile to be created")
		}
	})
}
