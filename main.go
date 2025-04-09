package main

import (
	"fmt"
	"log"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

func main() {
	fmt.Println("Go Builder Kit - Demo")
	fmt.Println("=====================")

	// Simple example
	simpleExample()

	// Nested structures example
	nestedStructuresExample()

	// Validation example
	validationExample()

	// Custom validation example
	customValidationExample()

	// Builder cloning example
	builderCloningExample()
}

// simpleExample demonstrates basic builder usage
func simpleExample() {
	fmt.Println("\n1. Simple Example")
	fmt.Println("----------------")

	// Create a person with basic fields
	personBuilder := builders.NewPersonBuilder().
		WithID("P12345").
		WithName("John Doe").
		WithAge(30).
		WithEmail("john.doe@example.com").
		WithPhone("555-123-4567").
		WithBirthdate("1993-05-15").
		WithGender("male").
		WithNationality("American").
		WithMaritalStatus("single")

	// Build the person
	person, err := personBuilder.BuildAndValidate()
	if err != nil {
		log.Fatalf("Failed to build person: %v", err)
	}

	fmt.Printf("Created person: %s (ID: %s, Age: %d)\n", person.Name, person.ID, person.Age)
	fmt.Printf("Contact: %s, %s\n", person.Email, person.Phone)
	fmt.Printf("Details: %s, %s, %s\n", person.Gender, person.Nationality, person.MaritalStatus)
}

// nestedStructuresExample demonstrates nested structure support
func nestedStructuresExample() {
	fmt.Println("\n2. Nested Structures Example")
	fmt.Println("---------------------------")

	// Create a GeoLocation builder
	geoLocationBuilder := builders.NewGeoLocationBuilder().
		WithLatitude(37.7749).
		WithLongitude(-122.4194).
		WithAccuracy(10.0)

	// Create an Address builder with nested GeoLocation
	addressBuilder := builders.NewAddressBuilder().
		WithStreet("123 Main St").
		WithCity("San Francisco").
		WithState("CA").
		WithPostalCode("94105").
		WithCountry("USA").
		WithCoordinates(geoLocationBuilder).
		WithType("Home")

	// Create a Company builder with nested Address
	companyBuilder := builders.NewCompanyBuilder().
		WithName("Acme Inc").
		WithIndustry("Technology").
		WithAddress(builders.NewAddressBuilder().
			WithStreet("456 Market St").
			WithCity("San Francisco").
			WithCountry("USA"))

	// Create an Employment builder with nested Company
	employmentBuilder := builders.NewEmploymentBuilder().
		WithPosition("Senior Software Engineer").
		WithCompany(companyBuilder).
		WithDepartment("Engineering").
		WithStartDate("2020-01-15")

	// Create a Person builder with all nested structures
	personBuilder := builders.NewPersonBuilder().
		WithID("P12345").
		WithName("John Doe").
		WithAge(28).
		WithEmail("john.doe@example.com").
		WithAddress(addressBuilder).
		WithEmployment(employmentBuilder)

	// Build the Person with all nested structures
	person, err := personBuilder.BuildAndValidate()
	if err != nil {
		log.Fatalf("Failed to build person with nested structures: %v", err)
	}

	fmt.Printf("Created person: %s (ID: %s)\n", person.Name, person.ID)
	fmt.Printf("Address: %s, %s, %s\n",
		person.Address.Street,
		person.Address.City,
		person.Address.State)
	fmt.Printf("Coordinates: %.4f, %.4f\n",
		person.Address.Coordinates.Latitude,
		person.Address.Coordinates.Longitude)
	fmt.Printf("Employment: %s at %s, %s department\n",
		person.Employment.Position,
		person.Employment.Company.Name,
		person.Employment.Department)
	fmt.Printf("Company Address: %s, %s\n",
		person.Employment.Company.Address.Street,
		person.Employment.Company.Address.City)
}

// validationExample demonstrates validation support
func validationExample() {
	fmt.Println("\n3. Validation Example")
	fmt.Println("-------------------")

	// Create a person with invalid data
	personBuilder := builders.NewPersonBuilder().
		WithID("P12345").
		WithName("John Doe").
		WithAge(-5).                // Invalid: negative age
		WithEmail("invalid").       // Invalid: not a valid email format
		WithBirthdate("01/01/2000") // Invalid: wrong date format

	// Try to build and validate
	_, err := personBuilder.BuildAndValidate()
	if err != nil {
		fmt.Printf("Validation failed as expected: %v\n", err)
	} else {
		log.Fatal("Validation should have failed but didn't")
	}

	// Fix the validation errors
	personBuilder = builders.NewPersonBuilder().
		WithID("P12345").
		WithName("John Doe").
		WithAge(30).
		WithEmail("john.doe@example.com").
		WithBirthdate("2000-01-01")

	// Build and validate again
	person, err := personBuilder.BuildAndValidate()
	if err != nil {
		log.Fatalf("Validation failed unexpectedly: %v", err)
	}

	fmt.Printf("Created valid person: %s (Age: %d, Email: %s)\n",
		person.Name, person.Age, person.Email)
}

// customValidationExample demonstrates custom validation
func customValidationExample() {
	fmt.Println("\n4. Custom Validation Example")
	fmt.Println("--------------------------")

	// Create a person with custom validation
	personBuilder := builders.NewPersonBuilder().
		WithID("P12345").
		WithName("John Doe").
		WithAge(30).
		WithEmail("john.doe@example.com").
		WithValidation(func(p *models.Person) error {
			if p.Age < 21 {
				return fmt.Errorf("person must be at least 21 years old")
			}
			return nil
		})

	// Build and validate
	person, err := personBuilder.BuildAndValidate()
	if err != nil {
		log.Fatalf("Custom validation failed unexpectedly: %v", err)
	}

	fmt.Printf("Created person with custom validation: %s (Age: %d)\n",
		person.Name, person.Age)

	// Create another person that fails custom validation
	personBuilder = builders.NewPersonBuilder().
		WithID("P67890").
		WithName("Jane Smith").
		WithAge(18).
		WithEmail("jane.smith@example.com").
		WithValidation(func(p *models.Person) error {
			if p.Age < 21 {
				return fmt.Errorf("person must be at least 21 years old")
			}
			return nil
		})

	// Try to build and validate
	_, err = personBuilder.BuildAndValidate()
	if err != nil {
		fmt.Printf("Custom validation failed as expected: %v\n", err)
	} else {
		log.Fatal("Custom validation should have failed but didn't")
	}
}

// builderCloningExample demonstrates builder cloning
func builderCloningExample() {
	fmt.Println("\n5. Builder Cloning Example")
	fmt.Println("------------------------")

	// Create a base address builder
	baseAddressBuilder := builders.NewAddressBuilder().
		WithStreet("123 Main St").
		WithCity("San Francisco").
		WithState("CA").
		WithCountry("USA").
		WithType("Home")

	// Create a base person builder
	basePersonBuilder := builders.NewPersonBuilder().
		WithName("John Doe").
		WithAge(30).
		WithEmail("john.doe@example.com").
		WithAddress(baseAddressBuilder)

	// Clone the base person builder and modify it
	clonedPersonBuilder := basePersonBuilder.Clone().
		WithName("Jane Smith").
		WithEmail("jane.smith@example.com")

	// Build both persons
	basePerson := basePersonBuilder.BuildPtr()
	clonedPerson := clonedPersonBuilder.BuildPtr()

	fmt.Printf("Base person: %s (%s)\n", basePerson.Name, basePerson.Email)
	fmt.Printf("Cloned person: %s (%s)\n", clonedPerson.Name, clonedPerson.Email)
	fmt.Printf("Both have the same address: %s, %s\n",
		basePerson.Address.Street, clonedPerson.Address.Street)
}
