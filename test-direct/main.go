package main

import (
	"fmt"
	"log"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
)

func main() {
	fmt.Println("Testing Go Builder Kit")

	// Create a person with basic fields
	personBuilder := builders.NewPersonBuilder().
		WithID("P12345").
		WithName("John Doe").
		WithAge(30).
		WithEmail("john.doe@example.com")

	// Build the person
	person, err := personBuilder.BuildAndValidate()
	if err != nil {
		log.Fatalf("Failed to build person: %v", err)
	}

	fmt.Printf("Created person: %s (ID: %s, Age: %d)\n", person.Name, person.ID, person.Age)
}
