package simple_test

import (
	"fmt"
	"testing"

	"github.com/adil-faiyaz98/go-builder-kit/examples/simple"
)

func TestPersonBuilder(t *testing.T) {
	// Create a simple person
	person := simple.NewPersonBuilder().
		WithID("123").
		WithName("John Doe").
		WithAge(30).
		WithEmail("john.doe@example.com").
		WithTag("employee").
		WithTag("developer").
		WithMetadata("department", "engineering").
		Build()

	// Verify the person
	if person.ID != "123" {
		t.Errorf("Expected ID to be '123', got '%s'", person.ID)
	}
	if person.Name != "John Doe" {
		t.Errorf("Expected Name to be 'John Doe', got '%s'", person.Name)
	}
	if person.Age != 30 {
		t.Errorf("Expected Age to be 30, got %d", person.Age)
	}
	if person.Email != "john.doe@example.com" {
		t.Errorf("Expected Email to be 'john.doe@example.com', got '%s'", person.Email)
	}
	if len(person.Tags) != 2 {
		t.Errorf("Expected 2 tags, got %d", len(person.Tags))
	}
	if person.Metadata["department"] != "engineering" {
		t.Errorf("Expected department to be 'engineering', got '%s'", person.Metadata["department"])
	}
}

func TestPersonBuilderWithAddress(t *testing.T) {
	// Create an address
	addressBuilder := simple.NewAddressBuilder().
		WithStreet("123 Main St").
		WithCity("New York").
		WithState("NY").
		WithPostalCode("10001").
		WithCountry("USA")

	// Create a person with an address
	person := simple.NewPersonBuilder().
		WithID("456").
		WithName("Jane Smith").
		WithAge(28).
		WithEmail("jane.smith@example.com").
		WithAddress(addressBuilder).
		Build()

	// Verify the person
	if person.ID != "456" {
		t.Errorf("Expected ID to be '456', got '%s'", person.ID)
	}
	if person.Name != "Jane Smith" {
		t.Errorf("Expected Name to be 'Jane Smith', got '%s'", person.Name)
	}

	// Verify the address
	if person.Address == nil {
		t.Errorf("Expected Address to be non-nil")
	} else {
		if person.Address.Street != "123 Main St" {
			t.Errorf("Expected Street to be '123 Main St', got '%s'", person.Address.Street)
		}
		if person.Address.City != "New York" {
			t.Errorf("Expected City to be 'New York', got '%s'", person.Address.City)
		}
		if person.Address.Country != "USA" {
			t.Errorf("Expected Country to be 'USA', got '%s'", person.Address.Country)
		}
	}
}

func TestPersonBuilderValidation(t *testing.T) {
	// Create a person with invalid data
	_, err := simple.NewPersonBuilder().
		WithID("").  // Invalid: empty ID
		WithName("John Doe").
		WithAge(-5). // Invalid: negative age
		WithEmail("invalid-email"). // Invalid: not a valid email
		BuildAndValidate()

	// Verify validation error
	if err == nil {
		t.Errorf("Expected validation error, got nil")
	} else {
		fmt.Printf("Validation error: %v\n", err)
	}

	// Create a person with valid data
	person, err := simple.NewPersonBuilder().
		WithID("123").
		WithName("John Doe").
		WithAge(30).
		WithEmail("john.doe@example.com").
		BuildAndValidate()

	// Verify no validation error
	if err != nil {
		t.Errorf("Expected no validation error, got %v", err)
	}
	if person == nil {
		t.Errorf("Expected person to be non-nil")
	}
}

func TestPersonBuilderClone(t *testing.T) {
	// Create a base person builder
	baseBuilder := simple.NewPersonBuilder().
		WithID("123").
		WithName("John Doe").
		WithAge(30).
		WithEmail("john.doe@example.com")

	// Clone the builder and modify it
	clonedBuilder := baseBuilder.Clone().
		WithName("Jane Smith").
		WithEmail("jane.smith@example.com")

	// Build both persons
	basePerson := baseBuilder.Build()
	clonedPerson := clonedBuilder.Build()

	// Verify the base person
	if basePerson.Name != "John Doe" {
		t.Errorf("Expected base Name to be 'John Doe', got '%s'", basePerson.Name)
	}
	if basePerson.Email != "john.doe@example.com" {
		t.Errorf("Expected base Email to be 'john.doe@example.com', got '%s'", basePerson.Email)
	}

	// Verify the cloned person
	if clonedPerson.Name != "Jane Smith" {
		t.Errorf("Expected cloned Name to be 'Jane Smith', got '%s'", clonedPerson.Name)
	}
	if clonedPerson.Email != "jane.smith@example.com" {
		t.Errorf("Expected cloned Email to be 'jane.smith@example.com', got '%s'", clonedPerson.Email)
	}

	// Verify that the ID and Age are the same
	if clonedPerson.ID != basePerson.ID {
		t.Errorf("Expected cloned ID to be '%s', got '%s'", basePerson.ID, clonedPerson.ID)
	}
	if clonedPerson.Age != basePerson.Age {
		t.Errorf("Expected cloned Age to be %d, got %d", basePerson.Age, clonedPerson.Age)
	}
}

func ExamplePersonBuilder() {
	// Create an address
	addressBuilder := simple.NewAddressBuilder().
		WithStreet("123 Main St").
		WithCity("New York").
		WithState("NY").
		WithPostalCode("10001").
		WithCountry("USA")

	// Create a person with an address
	person := simple.NewPersonBuilder().
		WithID("456").
		WithName("Jane Smith").
		WithAge(28).
		WithEmail("jane.smith@example.com").
		WithAddress(addressBuilder).
		WithTag("employee").
		WithMetadata("department", "engineering").
		Build()

	fmt.Printf("Created person: %s (%d)\n", person.Name, person.Age)
	fmt.Printf("Email: %s\n", person.Email)
	fmt.Printf("Address: %s, %s, %s\n", person.Address.Street, person.Address.City, person.Address.Country)
	fmt.Printf("Tags: %v\n", person.Tags)
	fmt.Printf("Department: %s\n", person.Metadata["department"])

	// Output:
	// Created person: Jane Smith (28)
	// Email: jane.smith@example.com
	// Address: 123 Main St, New York, USA
	// Tags: [employee]
	// Department: engineering
}
