package main

import (
	"fmt"
)

func main() {
	// Simple example
	fmt.Println("=== Simple Example ===")
	runSimpleExample()
	fmt.Println()

	// Complex example
	fmt.Println("=== Complex Example ===")
}

func runSimpleExample() {
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

	// Validation example
	_, err := simple.NewPersonBuilder().
		WithID(""). // Invalid: empty ID
		WithName("John Doe").
		WithAge(-5).                // Invalid: negative age
		WithEmail("invalid-email"). // Invalid: not a valid email
		BuildAndValidate()

	if err != nil {
		fmt.Printf("Validation error: %v\n", err)
	}
}
