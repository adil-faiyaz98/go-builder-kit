# Go Builder Kit

A comprehensive toolkit for implementing the Builder pattern in Go, with support for complex types, nested objects, validation, and code generation.

[![Go Builder Kit CI/CD](https://github.com/adil-faiyaz98/go-builder-kit/actions/workflows/go.yml/badge.svg)](https://github.com/adil-faiyaz98/go-builder-kit/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/adil-faiyaz98/go-builder-kit)](https://goreportcard.com/report/github.com/adil-faiyaz98/go-builder-kit)
[![GoDoc](https://godoc.org/github.com/adil-faiyaz98/go-builder-kit?status.svg)](https://godoc.org/github.com/adil-faiyaz98/go-builder-kit)

## Features

- **Builder Pattern Implementation**: Ready-to-use builder pattern for Go structs
- **Comprehensive Nested Structure Support**: Fully handles deeply nested structures at any level of nesting
- **Automatic Builder Generation**: Generate builder code for any struct
- **Validation Support**: Built-in and custom validation for models with proper error propagation through nested structures
- **Fluent Interface**: Clean, readable API with method chaining
- **Type Safety**: Maintain type safety with Go's type system
- **Performance Optimized**: Efficient implementation without compromising readability
- **Builder Composition**: Support for builder composition and inheritance
- **Comprehensive Documentation**: Extensive examples and API documentation
- **Comprehensive Testing**: Extensive test suite with positive, negative, and performance tests
- **Reserved Keyword Handling**: Properly handles Go reserved keywords like 'type' in struct fields

## Installation

```bash
go get github.com/adil-faiyaz98/go-builder-kit
```

## Releasing

To release a new version of the library:

1. Ensure all tests pass: `go test ./...`
2. Tag the release: `git tag v1.0.0`
3. Push the tag: `git push origin v1.0.0`
4. The library will be available as a Go module: `go get github.com/adil-faiyaz98/go-builder-kit@v1.0.0`

## Quick Start

### Using the Builder Generator

```bash
# Install the builder generator
go install github.com/adil-faiyaz98/go-builder-kit/cmd/builder-gen@latest

# Generate builders for a single file
builder-gen -input models/person.go -output builders -models-package github.com/adil-faiyaz98/go-builder-kit/models

# Generate builders for all files in a directory
builder-gen -input models -output builders -recursive -models-package github.com/adil-faiyaz98/go-builder-kit/models
```

### Using the Generated Builders

```go
package main

import (
    "fmt"

    "github.com/adil-faiyaz98/go-builder-kit/builders"
    "github.com/adil-faiyaz98/go-builder-kit/models"
)

func main() {
    // Create a person with validation
    person, err := builders.NewPersonBuilder().
        WithID("123").
        WithName("John Doe").
        WithAge(30).
        WithEmail("john.doe@example.com").
        WithPhone("+1-555-123-4567").
        WithBirthdate("1993-05-15").
        WithMaritalStatus("Single").
        BuildWithValidation()

    if err != nil {
        fmt.Printf("Validation error: %v\n", err)
        return
    }

    fmt.Printf("Created person: %s (%d)\n", person.Name, person.Age)

    // Create a complex nested structure
    addressBuilder := builders.NewAddressBuilder().
        WithStreet("123 Main St").
        WithCity("New York").
        WithState("NY").
        WithPostalCode("10001").
        WithCountry("USA").
        WithType("Home")

    educationBuilder := builders.NewEducationBuilder().
        WithDegree("Bachelor of Science").
        WithInstitution("Harvard University").
        WithStartDate("2018-09-01").
        WithEndDate("2022-05-15").
        WithGPA(3.8).
        WithMajor("Computer Science")

    complexPerson := builders.NewPersonBuilder().
        WithID("456").
        WithName("Jane Smith").
        WithAge(28).
        WithEmail("jane.smith@example.com").
        WithAddress(addressBuilder).
        WithEducation(educationBuilder).
        Build().(models.Person)

    fmt.Printf("Created complex person: %s from %s\n",
        complexPerson.Name,
        complexPerson.Education.Institution)
}
```

## Documentation

### Examples

- [Basic Usage](docs/examples/basic_usage.md)
- [Complex Objects](docs/examples/complex_objects.md)
- [Builder Composition](docs/examples/builder_composition.md)
- [Performance Optimization](docs/examples/performance_optimization.md)

### API Documentation

- [API Reference](docs/api/README.md)

## Testing with Ginkgo and Gomega

Go Builder Kit works seamlessly with Ginkgo and Gomega for BDD-style testing, which is especially valuable for testing complex nested structures.

### Installation

```bash
go get github.com/onsi/ginkgo/v2
go get github.com/onsi/gomega
```

### Example Test

```go
var _ = Describe("Person Builder", func() {
    Context("with nested structures", func() {
        It("should build a valid Person with nested Address", func() {
            // Create a person with nested address
            person := builders.NewPersonBuilder().
                WithID("123").
                WithName("John Doe").
                WithEmail("john.doe@example.com").
                WithAddress(builders.NewAddressBuilder().
                    WithStreet("123 Main St").
                    WithCity("San Francisco").
                    WithCountry("USA")).
                Build()

            // Verify the person
            Expect(person.ID).To(Equal("123"))
            Expect(person.Name).To(Equal("John Doe"))
            Expect(person.Email).To(Equal("john.doe@example.com"))

            // Verify nested address
            Expect(person.Address).NotTo(BeNil())
            Expect(person.Address.Street).To(Equal("123 Main St"))
            Expect(person.Address.City).To(Equal("San Francisco"))
            Expect(person.Address.Country).To(Equal("USA"))
        })

        It("should fail validation with invalid data", func() {
            // Create a person with invalid data
            _, err := builders.NewPersonBuilder().
                WithID("").  // Invalid: empty ID
                WithName("John Doe").
                WithEmail("invalid-email"). // Invalid: not a valid email
                BuildAndValidate()

            // Verify validation error
            Expect(err).To(HaveOccurred())
            Expect(err.Error()).To(ContainSubstring("ID cannot be empty"))
            Expect(err.Error()).To(ContainSubstring("Email is not valid"))
        })
    })
})
```

### Benefits for Testing Complex Nested Structures

1. **Readable Test Setup**: Fluent builders make test setup code clear and maintainable
2. **Focused Testing**: Test specific aspects of complex objects without building the entire structure
3. **Reusable Test Fixtures**: Create base builders that can be cloned and modified for different test cases
4. **Validation Testing**: Easily test both positive and negative validation scenarios
5. **Expressive Assertions**: Ginkgo and Gomega provide rich assertions for verifying complex object graphs

## Builder Pattern

The builder pattern is a creational design pattern that lets you construct complex objects step by step. It's particularly useful when:

- You need to create objects with many optional parameters
- You want to enforce construction rules
- You need to create immutable objects
- You want to separate construction logic from business logic

### Benefits of the Builder Pattern

1. **Readability**: The fluent interface makes code more readable and self-documenting
2. **Maintainability**: Each builder encapsulates the construction logic for its model
3. **Flexibility**: Builders allow for optional parameters and different construction sequences
4. **Testability**: Makes it easy to create test fixtures with specific configurations
5. **Immutability**: Can support immutable objects by building them completely before returning

## Handling Reserved Keywords

The Go Builder Kit properly handles Go reserved keywords like 'type' in struct fields. When a struct has a field named 'type', the builder will generate a method named 'WithType' that takes a parameter named 'value' instead of 'type'.

```go
// For a struct with a field named 'type'
type Device struct {
    Type string
    Model string
}

// The builder will generate a method like this
func (b *DeviceBuilder) WithType(value string) *DeviceBuilder {
    b.device.Type = value
    return b
}

// Usage
device := builders.NewDeviceBuilder().
    WithType("Smartphone").
    WithModel("iPhone 13").
    Build().(*models.Device)
```

## Validation

Go Builder Kit supports validation for models:

```go
// Create a person with validation
person, err := builders.NewPersonBuilder().
    WithID("123").
    WithName("John Doe").
    WithAge(-5).  // Invalid: negative age
    BuildAndValidate()

if err != nil {
    fmt.Printf("Validation error: %v\n", err)
    // Output: Validation error: validation failed: Age cannot be negative
}
```

### Custom Validation

You can add custom validation rules to your builders:

```go
// Create a person with custom validation
person, err := builders.NewPersonBuilder().
    WithID("123").
    WithName("John Doe").
    WithAge(30).
    WithEmail("john.doe@example.com").
    // Add custom validation rule
    WithValidation(func(p *models.Person) error {
        if !strings.HasPrefix(p.Email, "john.") {
            return fmt.Errorf("email must start with 'john.'")
        }
        return nil
    }).
    BuildWithValidation()

if err != nil {
    fmt.Printf("Validation error: %v\n", err)
    return
}
```

## Nested Structures

Go Builder Kit excels at handling deeply nested structures, which is one of its core strengths. The library provides comprehensive support for building and validating complex object graphs with multiple levels of nesting.

### Example of Deep Nesting

```go
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
    WithAddress(addressBuilder)

// Create an Employment builder with nested Company
employmentBuilder := builders.NewEmploymentBuilder().
    WithPosition("Software Engineer").
    WithDepartment("Engineering").
    WithCompany(companyBuilder)

// Create a Person builder with nested Address and Employment
personBuilder := builders.NewPersonBuilder().
    WithName("John Doe").
    WithAddress(addressBuilder).
    WithEmployment(employmentBuilder)

// Build and validate the entire object graph
person, err := personBuilder.BuildAndValidate()
```

### Key Features for Nested Structures

1. **Automatic Validation Propagation**: Validation errors from nested objects are properly propagated up the object graph
2. **Builder Composition**: Builders can be composed to create complex object graphs
3. **Deep Cloning**: Builders can be cloned to create variations of complex objects
4. **Type Safety**: Maintain type safety throughout the object graph
5. **Fluent Interface**: Clean, readable API even with deeply nested structures

## Performance

Go Builder Kit is designed to be efficient without compromising readability. Here are some performance optimization techniques:

1. **Reusing Builders**: Clone builders to create variations of the same object
2. **Builder Pools**: Use sync.Pool to reuse builders
3. **Batch Processing**: Process multiple objects in a batch
4. **Lazy Validation**: Defer validation until needed
5. **Incremental Validation**: Validate fields incrementally
6. **Minimizing Allocations**: Preallocate capacity for slices and maps

See [Performance Optimization](docs/examples/performance_optimization.md) for more details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### Development

```bash
# Clone the repository
git clone https://github.com/adil-faiyaz98/go-builder-kit.git
cd go-builder-kit

# Install dependencies
go mod download

# Run tests
go test ./...

# Run benchmarks
go test -bench=. ./...

# Generate builders
go run cmd/builder-gen/main.go -input models -output generated-builders -recursive
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
