# Performance Optimization

This document demonstrates performance optimization techniques for Go Builder Kit.

## Benchmarking Builder Performance

First, let's create a benchmark to measure the performance of our builders:

```go
package builders_test

import (
    "testing"

    "github.com/adil-faiyaz98/go-builder-kit/builders"
)

func BenchmarkSimplePersonBuilder(b *testing.B) {
    for i := 0; i < b.N; i++ {
        builders.NewPersonBuilder().
            WithID("123").
            WithName("John Doe").
            WithAge(30).
            WithEmail("john.doe@example.com").
            Build()
    }
}

func BenchmarkComplexPersonBuilder(b *testing.B) {
    for i := 0; i < b.N; i++ {
        builders.NewPersonBuilder().
            WithID("123").
            WithName("John Doe").
            WithAge(30).
            WithEmail("john.doe@example.com").
            WithAddress(
                builders.NewAddressBuilder().
                    WithStreet("123 Main St").
                    WithCity("New York").
                    WithState("NY").
                    WithPostalCode("10001").
                    WithCountry("USA"),
            ).
            WithEducation(
                builders.NewEducationBuilder().
                    WithDegree("Bachelor of Science").
                    WithInstitution("Harvard University").
                    WithStartDate("2018-09-01").
                    WithEndDate("2022-05-15").
                    WithGPA(3.8),
            ).
            WithEmployment(
                builders.NewEmploymentBuilder().
                    WithPosition("Software Engineer").
                    WithDepartment("Engineering").
                    WithStartDate("2022-06-01").
                    WithSalary(100000.0),
            ).
            Build()
    }
}

func BenchmarkPersonBuilderWithValidation(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _, _ = builders.NewPersonBuilder().
            WithID("123").
            WithName("John Doe").
            WithAge(30).
            WithEmail("john.doe@example.com").
            BuildWithValidation()
    }
}

func BenchmarkBulkPersonCreation(b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        createManyPersons(100)
    }
}

func createManyPersons(count int) {
    persons := make([]interface{}, count)
    for i := 0; i < count; i++ {
        persons[i] = builders.NewPersonBuilder().
            WithID(fmt.Sprintf("ID-%d", i)).
            WithName(fmt.Sprintf("Person %d", i)).
            WithAge(20 + (i % 50)).
            WithEmail(fmt.Sprintf("person%d@example.com", i)).
            Build()
    }
}
```

## Optimizing Builder Creation

Here are some techniques to optimize builder creation:

### 1. Reusing Builders

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
        WithAge(30).
        WithEmail("example@example.com")

    // Create 1000 persons by reusing the base builder
    persons := make([]models.Person, 1000)
    for i := 0; i < 1000; i++ {
        // Clone the base builder and customize
        persons[i] = baseBuilder.Clone().
            WithID(fmt.Sprintf("ID-%d", i)).
            WithName(fmt.Sprintf("Person %d", i)).
            Build().(models.Person)
    }

    fmt.Printf("Created %d persons\n", len(persons))
}
```

### 2. Using Builder Pools

```go
package main

import (
    "fmt"
    "sync"

    "github.com/adil-faiyaz98/go-builder-kit/builders"
    "github.com/adil-faiyaz98/go-builder-kit/models"
)

// BuilderPool is a pool of reusable builders
type BuilderPool struct {
    pool sync.Pool
}

// NewBuilderPool creates a new builder pool
func NewBuilderPool() *BuilderPool {
    return &BuilderPool{
        pool: sync.Pool{
            New: func() interface{} {
                return builders.NewPersonBuilder()
            },
        },
    }
}

// Get gets a builder from the pool
func (p *BuilderPool) Get() *builders.PersonBuilder {
    return p.pool.Get().(*builders.PersonBuilder)
}

// Put returns a builder to the pool
func (p *BuilderPool) Put(builder *builders.PersonBuilder) {
    // Reset the builder to its initial state
    *builder = *builders.NewPersonBuilder()
    p.pool.Put(builder)
}

func main() {
    // Create a builder pool
    pool := NewBuilderPool()

    // Create 1000 persons using the pool
    persons := make([]models.Person, 1000)
    for i := 0; i < 1000; i++ {
        // Get a builder from the pool
        builder := pool.Get()

        // Use the builder
        persons[i] = builder.
            WithID(fmt.Sprintf("ID-%d", i)).
            WithName(fmt.Sprintf("Person %d", i)).
            WithAge(30).
            WithEmail(fmt.Sprintf("person%d@example.com", i)).
            Build().(models.Person)

        // Return the builder to the pool
        pool.Put(builder)
    }

    fmt.Printf("Created %d persons\n", len(persons))
}
```

### 3. Batch Processing

```go
package main

import (
    "fmt"

    "github.com/adil-faiyaz98/go-builder-kit/builders"
    "github.com/adil-faiyaz98/go-builder-kit/models"
)

// BatchPersonBuilder creates multiple persons efficiently
type BatchPersonBuilder struct {
    baseBuilder *builders.PersonBuilder
    batchSize   int
    persons     []models.Person
}

// NewBatchPersonBuilder creates a new BatchPersonBuilder
func NewBatchPersonBuilder(batchSize int) *BatchPersonBuilder {
    return &BatchPersonBuilder{
        baseBuilder: builders.NewPersonBuilder(),
        batchSize:   batchSize,
        persons:     make([]models.Person, 0, batchSize),
    }
}

// WithBaseBuilder sets the base builder
func (b *BatchPersonBuilder) WithBaseBuilder(baseBuilder *builders.PersonBuilder) *BatchPersonBuilder {
    b.baseBuilder = baseBuilder
    return b
}

// AddPerson adds a person to the batch
func (b *BatchPersonBuilder) AddPerson(customize func(*builders.PersonBuilder) *builders.PersonBuilder) *BatchPersonBuilder {
    // Clone the base builder
    builder := b.baseBuilder.Clone()

    // Apply customizations
    builder = customize(builder)

    // Build the person and add to the batch
    b.persons = append(b.persons, builder.Build().(models.Person))

    return b
}

// Build builds all persons in the batch
func (b *BatchPersonBuilder) Build() []models.Person {
    return b.persons
}

func main() {
    // Create a base builder with common attributes
    baseBuilder := builders.NewPersonBuilder().
        WithAge(30).
        WithEmail("example@example.com")

    // Create a batch builder
    batchBuilder := NewBatchPersonBuilder(1000).
        WithBaseBuilder(baseBuilder)

    // Add 1000 persons to the batch
    for i := 0; i < 1000; i++ {
        batchBuilder.AddPerson(func(b *builders.PersonBuilder) *builders.PersonBuilder {
            return b.
                WithID(fmt.Sprintf("ID-%d", i)).
                WithName(fmt.Sprintf("Person %d", i))
        })
    }

    // Build all persons
    persons := batchBuilder.Build()

    fmt.Printf("Created %d persons\n", len(persons))
}
```

## Optimizing Validation

Validation can be expensive, especially for complex objects. Here are some techniques to optimize validation:

### 1. Lazy Validation

```go
package main

import (
    "fmt"

    "github.com/adil-faiyaz98/go-builder-kit/builders"
    "github.com/adil-faiyaz98/go-builder-kit/models"
)

// LazyValidator defers validation until needed
type LazyValidator struct {
    object     interface{}
    validated  bool
    validationErr error
}

// NewLazyValidator creates a new LazyValidator
func NewLazyValidator(object interface{}) *LazyValidator {
    return &LazyValidator{
        object:    object,
        validated: false,
    }
}

// Validate validates the object if not already validated
func (v *LazyValidator) Validate() error {
    if !v.validated {
        // Check if the object implements Validate() error
        if validator, ok := v.object.(interface{ Validate() error }); ok {
            v.validationErr = validator.Validate()
        }
        v.validated = true
    }
    return v.validationErr
}

// Get returns the object
func (v *LazyValidator) Get() interface{} {
    return v.object
}

// GetValidated returns the object after validation
func (v *LazyValidator) GetValidated() (interface{}, error) {
    err := v.Validate()
    return v.object, err
}

func main() {
    // Create a person builder
    personBuilder := builders.NewPersonBuilder().
        WithID("123").
        WithName("John Doe").
        WithAge(30).
        WithEmail("john.doe@example.com")

    // Build the person without validation
    person := personBuilder.Build().(models.Person)

    // Create a lazy validator
    validator := NewLazyValidator(person)

    // Use the person without validation
    fmt.Printf("Person: %s\n", person.Name)

    // Validate only when needed
    if err := validator.Validate(); err != nil {
        fmt.Printf("Validation error: %v\n", err)
    } else {
        fmt.Println("Person is valid")
    }
}
```

### 2. Incremental Validation

```go
package main

import (
    "fmt"

    "github.com/adil-faiyaz98/go-builder-kit/builders"
    "github.com/adil-faiyaz98/go-builder-kit/models"
)

// IncrementalValidator validates fields incrementally
type IncrementalValidator struct {
    person *models.Person
    errors []string
}

// NewIncrementalValidator creates a new IncrementalValidator
func NewIncrementalValidator(person *models.Person) *IncrementalValidator {
    return &IncrementalValidator{
        person: person,
        errors: make([]string, 0),
    }
}

// ValidateID validates the ID field
func (v *IncrementalValidator) ValidateID() *IncrementalValidator {
    if v.person.ID == "" {
        v.errors = append(v.errors, "ID cannot be empty")
    }
    return v
}

// ValidateName validates the Name field
func (v *IncrementalValidator) ValidateName() *IncrementalValidator {
    if v.person.Name == "" {
        v.errors = append(v.errors, "Name cannot be empty")
    } else if len(v.person.Name) < 2 {
        v.errors = append(v.errors, "Name must be at least 2 characters long")
    }
    return v
}

// ValidateAge validates the Age field
func (v *IncrementalValidator) ValidateAge() *IncrementalValidator {
    if v.person.Age < 0 {
        v.errors = append(v.errors, "Age cannot be negative")
    } else if v.person.Age > 150 {
        v.errors = append(v.errors, "Age cannot be greater than 150")
    }
    return v
}

// ValidateEmail validates the Email field
func (v *IncrementalValidator) ValidateEmail() *IncrementalValidator {
    if v.person.Email != "" {
        // Simple email validation
        if !strings.Contains(v.person.Email, "@") {
            v.errors = append(v.errors, "Email is not valid")
        }
    }
    return v
}

// GetErrors returns all validation errors
func (v *IncrementalValidator) GetErrors() []string {
    return v.errors
}

// HasErrors returns true if there are validation errors
func (v *IncrementalValidator) HasErrors() bool {
    return len(v.errors) > 0
}

func main() {
    // Create a person builder
    personBuilder := builders.NewPersonBuilder().
        WithID("123").
        WithName("John Doe").
        WithAge(30).
        WithEmail("john.doe@example.com")

    // Build the person without validation
    person := personBuilder.Build().(models.Person)

    // Create an incremental validator
    validator := NewIncrementalValidator(&person)

    // Validate only the fields we care about
    validator.ValidateID().ValidateName()

    // Check for errors
    if validator.HasErrors() {
        fmt.Printf("Validation errors: %v\n", validator.GetErrors())
    } else {
        fmt.Println("ID and Name are valid")
    }

    // Validate more fields later
    validator.ValidateAge().ValidateEmail()

    // Check for errors again
    if validator.HasErrors() {
        fmt.Printf("Validation errors: %v\n", validator.GetErrors())
    } else {
        fmt.Println("All fields are valid")
    }
}
```

## Memory Optimization

Here are some techniques to optimize memory usage:

### 1. Using Value Semantics

```go
package main

import (
    "fmt"

    "github.com/adil-faiyaz98/go-builder-kit/builders"
    "github.com/adil-faiyaz98/go-builder-kit/models"
)

func main() {
    // Create 1000 persons using value semantics
    persons := make([]models.Person, 1000)
    builder := builders.NewPersonBuilder()

    for i := 0; i < 1000; i++ {
        // Reset the builder for each person
        builder = builders.NewPersonBuilder().
            WithID(fmt.Sprintf("ID-%d", i)).
            WithName(fmt.Sprintf("Person %d", i)).
            WithAge(30)

        // Build the person using value semantics
        persons[i] = builder.Build().(models.Person)
    }

    fmt.Printf("Created %d persons\n", len(persons))
}
```

### 2. Minimizing Allocations

```go
package main

import (
    "fmt"

    "github.com/adil-faiyaz98/go-builder-kit/builders"
    "github.com/adil-faiyaz98/go-builder-kit/models"
)

// PreallocatedPersonBuilder minimizes allocations
type PreallocatedPersonBuilder struct {
    person models.Person
}

// NewPreallocatedPersonBuilder creates a new PreallocatedPersonBuilder
func NewPreallocatedPersonBuilder() *PreallocatedPersonBuilder {
    return &PreallocatedPersonBuilder{
        person: models.Person{
            Addresses:  make([]models.Address, 0, 2),  // Preallocate capacity for 2 addresses
            Education:  models.Education{},
            Employment: &models.Employment{},
            Metadata:   make(map[string]string, 5),    // Preallocate capacity for 5 metadata entries
        },
    }
}

// WithID sets the ID
func (b *PreallocatedPersonBuilder) WithID(id string) *PreallocatedPersonBuilder {
    b.person.ID = id
    return b
}

// WithName sets the Name
func (b *PreallocatedPersonBuilder) WithName(name string) *PreallocatedPersonBuilder {
    b.person.Name = name
    return b
}

// WithAge sets the Age
func (b *PreallocatedPersonBuilder) WithAge(age int) *PreallocatedPersonBuilder {
    b.person.Age = age
    return b
}

// WithEmail sets the Email
func (b *PreallocatedPersonBuilder) WithEmail(email string) *PreallocatedPersonBuilder {
    b.person.Email = email
    return b
}

// WithAddress adds an Address
func (b *PreallocatedPersonBuilder) WithAddress(address models.Address) *PreallocatedPersonBuilder {
    b.person.Addresses = append(b.person.Addresses, address)
    return b
}

// WithMetadata adds a metadata entry
func (b *PreallocatedPersonBuilder) WithMetadata(key, value string) *PreallocatedPersonBuilder {
    b.person.Metadata[key] = value
    return b
}

// Build builds the Person
func (b *PreallocatedPersonBuilder) Build() models.Person {
    return b.person
}

func main() {
    // Create 1000 persons using preallocated builders
    persons := make([]models.Person, 1000)
    
    for i := 0; i < 1000; i++ {
        builder := NewPreallocatedPersonBuilder().
            WithID(fmt.Sprintf("ID-%d", i)).
            WithName(fmt.Sprintf("Person %d", i)).
            WithAge(30).
            WithEmail(fmt.Sprintf("person%d@example.com", i))
            
        // Add some metadata
        builder.WithMetadata("created", "2023-05-01")
        builder.WithMetadata("type", "customer")
            
        // Build the person
        persons[i] = builder.Build()
    }

    fmt.Printf("Created %d persons\n", len(persons))
}
```
