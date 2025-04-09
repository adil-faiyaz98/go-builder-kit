# Basic Usage of Go Builder Kit

This document demonstrates the basic usage of the Go Builder Kit library.

## Creating a Simple Builder

Let's start with a simple example of creating a builder for a `Person` struct:

```go
package main

import (
    "fmt"

    "github.com/adil-faiyaz98/go-builder-kit/builders"
    "github.com/adil-faiyaz98/go-builder-kit/models"
)

func main() {
    // Create a person using the builder
    person := builders.NewPersonBuilder().
        WithID("123").
        WithName("John Doe").
        WithAge(30).
        WithEmail("john.doe@example.com").
        Build().(models.Person)

    fmt.Printf("Created person: %s (ID: %s, Age: %d)\n", person.Name, person.ID, person.Age)
}
```

## Using Validation

The builder pattern can be combined with validation to ensure that the created objects are valid:

```go
package main

import (
    "fmt"

    "github.com/adil-faiyaz98/go-builder-kit/builders"
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
        BuildWithValidation()

    if err != nil {
        fmt.Printf("Validation error: %v\n", err)
        return
    }

    fmt.Printf("Created valid person: %s\n", person.Name)

    // Example with validation error
    invalidPerson, err := builders.NewPersonBuilder().
        WithID("456").
        WithName("").  // Invalid: empty name
        WithAge(-5).   // Invalid: negative age
        WithEmail("not-an-email").  // Invalid: bad email format
        BuildWithValidation()

    if err != nil {
        fmt.Printf("Validation error: %v\n", err)
        // Output: Validation error: validation failed: Name cannot be empty; Age cannot be negative; Email is not valid
    } else {
        fmt.Printf("Created person: %s\n", invalidPerson.Name)
    }
}
```

## Custom Validation Rules

You can add custom validation rules to your builders:

```go
package main

import (
    "fmt"
    "strings"

    "github.com/adil-faiyaz98/go-builder-kit/builders"
    "github.com/adil-faiyaz98/go-builder-kit/models"
)

func main() {
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

    fmt.Printf("Created valid person: %s\n", person.Name)
}
```

## Cloning Builders

You can clone builders to create variations of the same object:

```go
package main

import (
    "fmt"

    "github.com/adil-faiyaz98/go-builder-kit/builders"
    "github.com/adil-faiyaz98/go-builder-kit/models"
)

func main() {
    // Create a base builder
    baseBuilder := builders.NewPersonBuilder().
        WithName("John Doe").
        WithAge(30).
        WithEmail("john.doe@example.com")

    // Clone the builder and customize
    employee := baseBuilder.Clone().
        WithID("EMP-123").
        WithEmployment(
            builders.NewEmploymentBuilder().
                WithPosition("Software Engineer").
                WithSalary(100000.0),
        ).
        Build().(models.Person)

    // Clone again for a different variation
    customer := baseBuilder.Clone().
        WithID("CUST-456").
        WithAddress(
            builders.NewAddressBuilder().
                WithStreet("123 Main St").
                WithCity("New York").
                WithCountry("USA"),
        ).
        Build().(models.Person)

    fmt.Printf("Employee: %s (ID: %s, Position: %s)\n", 
        employee.Name, employee.ID, employee.Employment.Position)
    
    fmt.Printf("Customer: %s (ID: %s, City: %s)\n", 
        customer.Name, customer.ID, customer.Address.City)
}
```

## Creating Builders from Existing Models

You can create a builder from an existing model:

```go
package main

import (
    "fmt"

    "github.com/adil-faiyaz98/go-builder-kit/builders"
    "github.com/adil-faiyaz98/go-builder-kit/models"
)

func main() {
    // Create a person directly
    existingPerson := models.Person{
        ID:    "123",
        Name:  "John Doe",
        Age:   30,
        Email: "john.doe@example.com",
    }

    // Create a builder from the existing person
    builder := builders.PersonBuilderFromModel(existingPerson)

    // Modify the person using the builder
    modifiedPerson := builder.
        WithAge(31).
        WithPhone("+1-555-123-4567").
        Build().(models.Person)

    fmt.Printf("Modified person: %s (Age: %d, Phone: %s)\n", 
        modifiedPerson.Name, modifiedPerson.Age, modifiedPerson.Phone)
}
```
