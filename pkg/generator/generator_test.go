package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerator(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "generator-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a test Go file with a struct
	testFile := filepath.Join(tempDir, "test.go")
	testCode := `package models

// Person represents a person
type Person struct {
	ID        string
	Name      string
	Age       int
	Email     string
	Addresses []Address
	Metadata  map[string]string
}

// Address represents a physical location
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
	Country    string
}
`
	if err := os.WriteFile(testFile, []byte(testCode), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// Create a generator
	gen := NewGenerator(Options{
		PackageName:   "builders",
		ModelsPackage: "github.com/example/models",
		Verbose:       true,
	})

	// Parse the file
	structs, err := gen.ParseFile(testFile)
	if err != nil {
		t.Fatalf("Failed to parse file: %v", err)
	}

	// Check that we found the expected structs
	if len(structs) != 2 {
		t.Fatalf("Expected 2 structs, got %d", len(structs))
	}

	// Check the Person struct
	person := structs[0]
	if person.Name != "Person" {
		t.Errorf("Expected struct name 'Person', got '%s'", person.Name)
	}
	if len(person.Fields) != 6 {
		t.Errorf("Expected 6 fields in Person, got %d", len(person.Fields))
	}

	// Check the Address struct
	address := structs[1]
	if address.Name != "Address" {
		t.Errorf("Expected struct name 'Address', got '%s'", address.Name)
	}
	if len(address.Fields) != 5 {
		t.Errorf("Expected 5 fields in Address, got %d", len(address.Fields))
	}

	// Generate builder code for Person
	personCode, err := gen.GenerateBuilder(person)
	if err != nil {
		t.Fatalf("Failed to generate builder for Person: %v", err)
	}

	// Check that the generated code contains expected elements
	expectedElements := []string{
		"package builders",
		"import",
		"github.com/example/models",
		"PersonBuilder",
		"NewPersonBuilder",
		"WithID",
		"WithName",
		"WithAge",
		"WithEmail",
		"WithAddresses",
		"WithMetadata",
		"Build",
		"BuildWithValidation",
	}

	for _, element := range expectedElements {
		if !contains(personCode, element) {
			t.Errorf("Expected generated code to contain '%s', but it doesn't", element)
		}
	}

	// Generate builder code for Address
	addressCode, err := gen.GenerateBuilder(address)
	if err != nil {
		t.Fatalf("Failed to generate builder for Address: %v", err)
	}

	// Check that the generated code contains expected elements
	expectedElements = []string{
		"package builders",
		"import",
		"github.com/example/models",
		"AddressBuilder",
		"NewAddressBuilder",
		"WithStreet",
		"WithCity",
		"WithState",
		"WithPostalCode",
		"WithCountry",
		"Build",
		"BuildWithValidation",
	}

	for _, element := range expectedElements {
		if !contains(addressCode, element) {
			t.Errorf("Expected generated code to contain '%s', but it doesn't", element)
		}
	}
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}
