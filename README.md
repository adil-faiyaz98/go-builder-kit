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
- **Slice and Map Handling**: Special methods for working with slices and maps.
- **Clone Support**: Deep copy functionality for builders.
- **Zero Dependencies**: No external runtime dependencies.

## Installation

Install the library using `go get`:

```bash
go get github.com/adil-faiyaz98/go-builder-kit
```

## Quick Start

### Generate Builders

Use the builder generator to create builders for your structs:

```bash
# Install the builder generator
go install github.com/adil-faiyaz98/go-builder-kit/cmd/builder-gen@latest

# Generate builders for your models
builder-gen -input path/to/models -output path/to/builders -models-package github.com/yourusername/yourproject/models
```

#### Command Line Options

- `-input`: Path to the input Go file or directory containing structs to generate builders for
- `-output`: Output directory for generated builder files
- `-models-package`: Import path for the models package (required)
- `-package-name`: Name of the generated package (default: "builders")
- `-recursive`: Process directories recursively
- `-verbose`: Enable verbose output

### Using Generated Builders

Once you've generated builders, you can use them to create instances of your structs with a fluent API:

```go
// Create an address builder
addressBuilder := builders.NewAddressBuilder().
    WithStreet("123 Main St").
    WithCity("Anytown").
    WithState("CA").
    WithPostalCode("12345").
    WithCountry("USA")

// Create a person builder with a nested address
personBuilder := builders.NewPersonBuilder().
    WithID("123").
    WithName("John Doe").
    WithAge(30).
    WithEmail("john@example.com").
    WithAddress(addressBuilder)

// Add items to slices
personBuilder.AddSkill("Go")
personBuilder.AddSkill("Java")

// Build the person
person := personBuilder.BuildPtr()
```

### Advanced Features

#### Validation

Add custom validation to your builders:

```go
personBuilder.WithValidation(func(p *models.Person) error {
    if p.Age < 0 {
        return fmt.Errorf("age cannot be negative")
    }
    return nil
})

// Build with validation
person, err := personBuilder.BuildAndValidate()
if err != nil {
    // Handle validation error
}
```

#### Cloning

Create deep copies of builders:

```go
clonedBuilder := personBuilder.Clone()
clonedBuilder.WithName("Jane Doe")

// Original builder is unaffected
person1 := personBuilder.BuildPtr()    // Name is still "John Doe"
person2 := clonedBuilder.BuildPtr()    // Name is "Jane Doe"
```

#### Working with Nested Structures

The builder pattern shines when working with complex nested structures:

```go
// Create department with employees
departmentBuilder := builders.NewDepartmentBuilder().
    WithName("Engineering").
    WithDescription("Software Engineering")

// Add multiple employees
for i := 1; i <= 3; i++ {
    employeeBuilder := builders.NewEmployeeBuilder().
        WithID(fmt.Sprintf("E%03d", i)).
        WithName(fmt.Sprintf("Engineer %d", i))

    departmentBuilder.AddEmployee(employeeBuilder)
}

// Build the department with all its employees
department := departmentBuilder.BuildPtr()
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

## Recent Improvements

- Fixed import path issues for better package handling
- Improved parameter naming for better code readability
- Enhanced handling of nested structs at any depth
- Added support for slice operations with proper type handling
- Removed unnecessary dependencies
- Fixed formatting issues in generated code

## Using GoReleaser

When using GoReleaser with this project, you may encounter a "git is in a dirty state" error. This can happen if there are untracked files or changes in your working directory. To bypass this validation, use the `--snapshot` flag:

```bash
goreleaser release --snapshot --clean
```

This will generate an unversioned snapshot release, skipping all validations without publishing any artifacts.

### GitHub Actions Workflow

This project includes a GitHub Actions workflow (`.github/workflows/release.yml`) that automatically runs GoReleaser with the `--snapshot --clean` flags when a new tag is pushed. This ensures that releases can be generated even if there are untracked files in the repository.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
