# Go Builder Kit

[![Go Version](https://img.shields.io/badge/go-1.23+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/adil-faiyaz98/go-builder-kit)](https://goreportcard.com/report/github.com/adil-faiyaz98/go-builder-kit)
[![GoDoc](https://godoc.org/github.com/adil-faiyaz98/go-builder-kit?status.svg)](https://pkg.go.dev/github.com/adil-faiyaz98/go-builder-kit)

A powerful Go library for implementing the Builder pattern with automatic code generation.

## Features

- Automatic builder generation from struct definitions
- Thread-safe operations
- Type-safe builder methods
- Deep cloning support
- Comprehensive validation framework
- Zero runtime dependencies

## Installation

```bash
go get github.com/adil-faiyaz98/go-builder-kit@latest
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

### Using Builders

```go
// Create a person with basic fields
person := builders.NewPersonBuilder().
    WithID("123").
    WithName("John Doe").
    WithEmail("john.doe@example.com").
    WithAge(30).
    BuildPtr()

// Create a person with nested structures
person := builders.NewPersonBuilder().
    WithID("123").
    WithName("John Doe").
    WithEmail("john.doe@example.com").
    WithAddress(
        builders.NewAddressBuilder().
            WithStreet("123 Main St").
            WithCity("San Francisco").
            WithState("CA").
            WithPostalCode("94105").
            WithCountry("USA"),
    ).
    BuildPtr()
```

## Builder Registry

```go
import "github.com/adil-faiyaz98/go-builder-kit/pkg/builder"

// Register builders with the global registry
builder.Register("person", func() any {
    return builders.NewPersonBuilder()
})

// Create instances directly from registry
person, err := builder.Create("person")
if err != nil {
    fmt.Printf("Failed to create person builder: %v\n", err)
    return
}
```

## Validation

```go
import "github.com/adil-faiyaz98/go-builder-kit/pkg/builder"

// Built-in validation utilities
var errors builder.ValidationErrors

// Validate required fields
builder.ValidateRequired(person.Name, "name", &errors)
builder.ValidateRequired(person.Email, "email", &errors)

// Validate string length
builder.ValidateStringLength(person.Name, "name", 1, 100, &errors)

// Check for validation errors
if errors.HasErrors() {
    fmt.Println("Validation failed:", errors.Error())
    return
}
```

## API Reference

### Core Builder Interface

```go
type GenericBuilder[T any] interface {
    Build() T
    BuildPtr() *T
    BuildAndValidate() (*T, error)
    MustBuild() T
    WithValidation(func(T) error) GenericBuilder[T]
    Clone() GenericBuilder[T]
}
```

### Generated Builder Methods

For each struct field, the following methods are generated:

```go
// Basic setters
func (b *PersonBuilder) WithName(name string) *PersonBuilder
func (b *PersonBuilder) WithAge(age int) *PersonBuilder
func (b *PersonBuilder) WithEmail(email string) *PersonBuilder

// Slice operations
func (b *PersonBuilder) WithSkills(skills []string) *PersonBuilder
func (b *PersonBuilder) AddSkill(skill string) *PersonBuilder
func (b *PersonBuilder) AddSkills(skills ...string) *PersonBuilder

// Build methods
func (b *PersonBuilder) Build() any
func (b *PersonBuilder) BuildPtr() *models.Person
func (b *PersonBuilder) BuildAndValidate() (*models.Person, error)
func (b *PersonBuilder) Clone() *PersonBuilder
```

## Testing with Go Builder Kit

```go
func TestPersonValidation(t *testing.T) {
    // Test valid person
    person, err := builders.NewPersonBuilder().
        WithName("John Doe").
        WithAge(30).
        WithEmail("john@example.com").
        BuildAndValidate()

    assert.NoError(t, err)
    assert.Equal(t, "John Doe", person.Name)

    // Test invalid person
    _, err = builders.NewPersonBuilder().
        WithName("Invalid Person").
        WithAge(-5).
        BuildAndValidate()

    assert.Error(t, err)
}
```

## Contributing

We welcome contributions! Please see our [Contributing Guidelines](CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Features

- Automatic builder generation from struct definitions
- Thread-safe operations
- Type-safe builder methods
- Deep cloning support
- Comprehensive validation framework
- Zero runtime dependencies

## Installation

```bash
go get github.com/adil-faiyaz98/go-builder-kit@latest
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

## Builder Registry

```go
import "github.com/adil-faiyaz98/go-builder-kit/pkg/builder"

// Register builders with the global registry
builder.Register("person", func() any {
    return builders.NewPersonBuilder()
})

// Create instances directly from registry
person, err := builder.Create("person")
if err != nil {
    fmt.Printf("Failed to create person builder: %v\n", err)
    return
}
```

## Validation

```go
import "github.com/adil-faiyaz98/go-builder-kit/pkg/builder"

// Built-in validation utilities
var errors builder.ValidationErrors

// Validate required fields
builder.ValidateRequired(person.Name, "name", &errors)
builder.ValidateRequired(person.Email, "email", &errors)

// Validate string length
builder.ValidateStringLength(person.Name, "name", 1, 100, &errors)

// Check for validation errors
if errors.HasErrors() {
    fmt.Println("Validation failed:", errors.Error())
    return
}
```

## API Reference

### Core Builder Interface

```go
type GenericBuilder[T any] interface {
    Build() T
    BuildPtr() *T
    BuildAndValidate() (*T, error)
    MustBuild() T
    WithValidation(func(T) error) GenericBuilder[T]
    Clone() GenericBuilder[T]
}
```

### Generated Builder Methods

For each struct field, the following methods are generated:

```go
// Basic setters (with automatic sanitization for strings)
func (b *PersonBuilder) WithName(name string) *PersonBuilder
func (b *PersonBuilder) WithAge(age int) *PersonBuilder
func (b *PersonBuilder) WithEmail(email string) *PersonBuilder

// Slice operations
func (b *PersonBuilder) WithSkills(skills []string) *PersonBuilder
func (b *PersonBuilder) AddSkill(skill string) *PersonBuilder
func (b *PersonBuilder) AddSkills(skills ...string) *PersonBuilder

// Nested object setters
func (b *PersonBuilder) WithAddress(address *AddressBuilder) *PersonBuilder

// Build methods
func (b *PersonBuilder) Build() any
func (b *PersonBuilder) BuildPtr() *models.Person
func (b *PersonBuilder) BuildAndValidate() (*models.Person, error)
func (b *PersonBuilder) Clone() *PersonBuilder
```

## Testing with Go Builder Kit

```go
func TestPersonValidation(t *testing.T) {
    // Test valid person
    person, err := builders.NewPersonBuilder().
        WithName("John Doe").
        WithAge(30).
        WithEmail("john@example.com").
        BuildAndValidate()

    assert.NoError(t, err)
    assert.Equal(t, "John Doe", person.Name)

    // Test invalid person
    _, err = builders.NewPersonBuilder().
        WithName("Invalid Person").
        WithAge(-5).
        BuildAndValidate()

    assert.Error(t, err)
}
```

## Contributing

We welcome contributions! Please see our [Contributing Guidelines](CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
