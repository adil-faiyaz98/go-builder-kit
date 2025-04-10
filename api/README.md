# Go Builder Kit API Documentation

This document provides comprehensive API documentation for the Go Builder Kit library.

## Table of Contents

1. [Installation](#installation)
2. [Builder Generator](#builder-generator)
3. [Generated Builder Methods](#generated-builder-methods)
4. [Working with Nested Structures](#working-with-nested-structures)
5. [Validation](#validation)
6. [Cloning](#cloning)
7. [Builder Registry](#builder-registry)
8. [Testing with Builders](#testing-with-builders)
9. [Best Practices](#best-practices)

## Installation

### Installing the Library

```bash
go get github.com/adil-faiyaz98/go-builder-kit
```

### Installing the Builder Generator

```bash
go install github.com/adil-faiyaz98/go-builder-kit/cmd/builder-gen@latest
```

## Builder Generator

The builder generator is a command-line tool that automatically generates builder implementations for your Go structs.

### Basic Usage

```bash
builder-gen -input path/to/models -output path/to/builders -models-package github.com/yourusername/yourproject/models
```

### Command Line Options

| Option | Description | Required | Default |
|--------|-------------|----------|--------|
| `-input` | Path to the input Go file or directory containing structs | Yes | - |
| `-output` | Output directory for generated builder files | Yes | - |
| `-models-package` | Import path for the models package | Yes | - |
| `-package-name` | Name of the generated package | No | "builders" |
| `-recursive` | Process directories recursively | No | false |
| `-verbose` | Enable verbose output | No | false |

### Example

Given a model like this:

```go
package models

type Person struct {
    ID        string
    Name      string
    Age       int
    Email     string
    Address   *Address
    Friends   []*Person
    Skills    []string
}

type Address struct {
    Street     string
    City       string
    State      string
    PostalCode string
    Country    string
}
```

Generate builders with:

```bash
builder-gen -input models -output builders -models-package github.com/yourusername/yourproject/models
```

## Generated Builder Methods

For each struct, the generator creates a builder class with the following methods:

### Constructor Methods

```go
// Creates a new builder with default values
NewPersonBuilder() *PersonBuilder

// Creates a new builder with sensible defaults (if implemented)
NewPersonBuilderWithDefaults() *PersonBuilder
```

### Setter Methods

For each field in the struct, a setter method is generated:

```go
// For simple fields
WithName(name string) *PersonBuilder
WithAge(age int) *PersonBuilder

// For pointer fields
WithAddress(address *AddressBuilder) *PersonBuilder

// For slice fields
WithFriends(friends []*PersonBuilder) *PersonBuilder
WithSkills(skills []string) *PersonBuilder

// For slice fields, an additional method to add a single item
AddFriend(friend *PersonBuilder) *PersonBuilder
AddSkill(skill string) *PersonBuilder
```

### Build Methods

```go
// Returns the built struct
Build() interface{}

// Returns a pointer to the built struct
BuildPtr() *models.Person

// Builds and validates the struct
BuildAndValidate() (*models.Person, error)

// Builds the struct and panics if validation fails
MustBuild() *models.Person
```

### Other Methods

```go
// Adds a custom validation function
WithValidation(validationFunc func(*models.Person) error) *PersonBuilder

// Creates a deep copy of the builder
Clone() *PersonBuilder
```

## Working with Nested Structures

One of the key strengths of the builder pattern is handling complex nested structures. Go Builder Kit makes this easy with its fluent API.

### Setting Nested Structs

```go
// Create an address builder
addressBuilder := builders.NewAddressBuilder().
    WithStreet("123 Main St").
    WithCity("Anytown").
    WithState("CA").
    WithPostalCode("12345").
    WithCountry("USA")

// Set the address in the person builder
personBuilder := builders.NewPersonBuilder().
    WithName("John Doe").
    WithAddress(addressBuilder)
```

### Working with Slices of Nested Structs

```go
// Create a department with employees
departmentBuilder := builders.NewDepartmentBuilder().
    WithName("Engineering")

// Add multiple employees
employee1 := builders.NewEmployeeBuilder().
    WithName("Alice").
    WithTitle("Senior Engineer")

employee2 := builders.NewEmployeeBuilder().
    WithName("Bob").
    WithTitle("Junior Engineer")

// Method 1: Add employees one at a time
departmentBuilder.AddEmployee(employee1)
departmentBuilder.AddEmployee(employee2)

// Method 2: Set all employees at once
departmentBuilder.WithEmployees([]*builders.EmployeeBuilder{employee1, employee2})
```

### Handling Circular References

The builder pattern can handle circular references elegantly:

```go
// Create employee and department with circular references
employeeBuilder := builders.NewEmployeeBuilder().
    WithName("John Doe").
    WithTitle("Manager")

departmentBuilder := builders.NewDepartmentBuilder().
    WithName("Engineering")

// Set circular references
employeeBuilder.WithDepartment(departmentBuilder)
departmentBuilder.WithManager(employeeBuilder)
```

## Validation

Go Builder Kit provides built-in support for validation.

### Custom Validation Functions

You can add custom validation functions to your builders:

```go
personBuilder.WithValidation(func(p *models.Person) error {
    if p.Age < 0 {
        return fmt.Errorf("age cannot be negative")
    }
    if p.Name == "" {
        return fmt.Errorf("name cannot be empty")
    }
    return nil
})
```

### Building with Validation

Use the `BuildAndValidate` method to build and validate in one step:

```go
person, err := personBuilder.BuildAndValidate()
if err != nil {
    // Handle validation error
    log.Fatalf("Validation failed: %v", err)
}
```

### Panic on Validation Failure

If you prefer to panic on validation failure (e.g., in tests), use `MustBuild`:

```go
// Will panic if validation fails
person := personBuilder.MustBuild()
```

## Cloning

Go Builder Kit supports deep copying of builders, which is useful for creating variations of objects.

```go
// Create a base person
basePersonBuilder := builders.NewPersonBuilder().
    WithName("John Doe").
    WithAge(30).
    WithEmail("john@example.com")

// Clone and modify for a different person
janeDoeBuilder := basePersonBuilder.Clone().
    WithName("Jane Doe").
    WithEmail("jane@example.com")

// Original builder is unaffected
johnDoe := basePersonBuilder.BuildPtr()  // Name is still "John Doe"
janeDoe := janeDoeBuilder.BuildPtr()     // Name is "Jane Doe"
```

## Builder Registry

The builder registry allows you to register and retrieve builders by name.

```go
// Register a builder
builders.RegisterBuilder("person", func() builders.Builder {
    return builders.NewPersonBuilder()
})

// Get a builder by name
builder, ok := builders.GetBuilder("person")
if ok {
    personBuilder := builder.(*builders.PersonBuilder)
    // Use the builder
}
```

## Testing with Builders

Builders are particularly useful in testing scenarios, making it easy to create test fixtures.

### Creating Test Fixtures

```go
func TestPersonService(t *testing.T) {
    // Create a test person with all required fields
    person := builders.NewPersonBuilder().
        WithID("123").
        WithName("Test Person").
        WithAge(30).
        BuildPtr()
    
    // Use the person in your test
    result := service.ProcessPerson(person)
    
    // Assert on the result
    assert.Equal(t, expected, result)
}
```

### Table-Driven Tests

```go
func TestPersonValidation(t *testing.T) {
    tests := []struct {
        name    string
        builder func() *builders.PersonBuilder
        wantErr bool
    }{
        {
            name: "valid person",
            builder: func() *builders.PersonBuilder {
                return builders.NewPersonBuilder().
                    WithName("John Doe").
                    WithAge(30)
            },
            wantErr: false,
        },
        {
            name: "invalid age",
            builder: func() *builders.PersonBuilder {
                return builders.NewPersonBuilder().
                    WithName("John Doe").
                    WithAge(-1)
            },
            wantErr: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            _, err := tt.builder().BuildAndValidate()
            if (err != nil) != tt.wantErr {
                t.Errorf("BuildAndValidate() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

## Best Practices

### When to Use Builders

- **Complex Objects**: Use builders for structs with many fields or nested structures
- **Optional Fields**: Builders are great when many fields are optional
- **Validation**: When you need to validate objects during construction
- **Testing**: For creating test fixtures with different variations

### Tips for Using Builders

1. **Keep Builders in a Separate Package**: This helps maintain separation of concerns
2. **Use Sensible Defaults**: Implement `NewXBuilderWithDefaults()` methods for common use cases
3. **Add Validation**: Use the validation support to ensure objects are valid
4. **Document Builder Methods**: Add comments to explain what each method does
5. **Use Clone for Variations**: When creating similar objects, clone a base builder

### Performance Considerations

- Builders add a small overhead compared to direct struct initialization
- For performance-critical code paths, consider direct initialization
- Use the `BuildPtr()` method to avoid unnecessary copying
