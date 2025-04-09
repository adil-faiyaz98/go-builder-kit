# Go Builder Kit

A comprehensive toolkit for implementing the Builder pattern in Go, with support for complex types, nested objects, validation, and code generation.

[![Go Builder Kit CI/CD](https://github.com/adil-faiyaz98/go-builder-kit/actions/workflows/ci.yml/badge.svg)](https://github.com/adil-faiyaz98/go-builder-kit/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/adil-faiyaz98/go-builder-kit)](https://goreportcard.com/report/github.com/adil-faiyaz98/go-builder-kit)
[![GoDoc](https://godoc.org/github.com/adil-faiyaz98/go-builder-kit?status.svg)](https://godoc.org/github.com/adil-faiyaz98/go-builder-kit)

## Features

- **Builder Pattern Implementation**: Ready-to-use builder pattern for Go structs
- **Complex Type Support**: Handles complex nested structures and collections
- **Automatic Builder Generation**: Generate builder code for any struct
- **Validation Support**: Built-in and custom validation for models
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
