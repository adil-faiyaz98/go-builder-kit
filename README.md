# Go Builder Kit

![Release Status](https://github.com/adil-faiyaz98/go-builder-kit/actions/workflows/release.yml/badge.svg)

A powerful toolkit for implementing the Builder pattern in Go, designed to simplify the creation of complex and nested objects with validation and code generation.

[![Go Builder Kit CI/CD](https://github.com/adil-faiyaz98/go-builder-kit/actions/workflows/go.yml/badge.svg)](https://github.com/adil-faiyaz98/go-builder-kit/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/adil-faiyaz98/go-builder-kit)](https://goreportcard.com/report/github.com/adil-faiyaz98/go-builder-kit)
[![GoDoc](https://godoc.org/github.com/adil-faiyaz98/go-builder-kit?status.svg)](https://godoc.org/github.com/adil-faiyaz98/go-builder-kit)

## Key Features

- **Builder Pattern**: Simplifies object creation with a fluent, type-safe API.
- **Nested Structures**: Handles deeply nested structs with ease.
- **Validation**: Built-in and custom validation support for models.
- **Code Generation**: Automatically generate builders for any struct.
- **Performance Optimized**: Efficient implementation with minimal overhead.
- **Testing Support**: Streamlines test data creation for unit and integration tests.

## Installation

Install the library using `go get`:

```bash
go get github.com/adil-faiyaz98/go-builder-kit
```

## Quick Start

### Generate Builders

Use the builder generator to create builders for your structs:

```bash
go install github.com/adil-faiyaz98/go-builder-kit/cmd/builder-gen@latest

builder-gen -input models -output builders -recursive -models-package github.com/adil-faiyaz98/go-builder-kit/models
```

### Create Objects with Builders

```go
package main

import (
    "fmt"
    "github.com/adil-faiyaz98/go-builder-kit/builders"
)

func main() {
    person, err := builders.NewPersonBuilder().
        WithID("123").
        WithName("John Doe").
        WithAge(30).
        WithEmail("john.doe@example.com").
        BuildWithValidation()

    if err != nil {
        fmt.Printf("Validation error: %v\n", err)
        return
    }

    fmt.Printf("Created person: %s (%d)\n", person.Name, person.Age)
}
```

### Nested Structures

```go
address := builders.NewAddressBuilder().
    WithStreet("123 Main St").
    WithCity("New York").
    WithCountry("USA").
    Build().(*models.Address)

person := builders.NewPersonBuilder().
    WithName("Jane Doe").
    WithAddress(address).
    Build().(*models.Person)

fmt.Printf("Person: %s, Address: %s, %s\n", person.Name, person.Address.Street, person.Address.City)
```

## Testing with Builders

Builders simplify test setup and validation:

```go
func TestPersonBuilder(t *testing.T) {
    person := builders.NewPersonBuilder().
        WithName("John Doe").
        WithEmail("john.doe@example.com").
        Build().(*models.Person)

    if person.Name != "John Doe" {
        t.Errorf("expected name to be 'John Doe', got '%s'", person.Name)
    }
}
```

## Why Use Go Builder Kit?

- **Simplifies Complex Object Creation**: Ideal for structs with many optional fields or deeply nested structures.
- **Improves Code Readability**: Fluent API makes object creation intuitive and maintainable.
- **Reduces Boilerplate**: Automatically generates builders, saving time and effort.
- **Enhances Testability**: Easily create reusable test fixtures with minimal code.

## Documentation

For detailed API documentation and examples, see the [API Documentation](docs/api/README.md).

## Contributing

Contributions are welcome! Please submit a Pull Request or open an issue.

### Development

```bash
# Clone the repository
git clone https://github.com/adil-faiyaz98/go-builder-kit.git
cd go-builder-kit

# Install dependencies
go mod download

# Run tests
go test ./...

# Generate builders
go run cmd/builder-gen/main.go -input models -output builders -recursive
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.
