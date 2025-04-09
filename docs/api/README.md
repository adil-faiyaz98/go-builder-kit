# Go Builder Kit API Documentation

This document provides comprehensive API documentation for the Go Builder Kit library.

## Table of Contents

1. [Builder Interface](#builder-interface)
2. [Builder Registry](#builder-registry)
3. [Model Builders](#model-builders)
4. [Validation](#validation)
5. [Builder Generator](#builder-generator)

## Builder Interface

The `Builder` interface is the core of the Go Builder Kit library. It defines the basic functionality that all builders must implement.

```go
// Builder is the interface that all builders must implement
type Builder interface {
    // Build builds the object and returns it
    Build() interface{}
}
```

## Builder Registry

The builder registry allows you to register and retrieve builders by name.

```go
// Register registers a builder function
func Register(name string, fn BuilderFunc)

// Get returns a builder function by name
func Get(name string) (BuilderFunc, bool)

// GetAll returns all registered builder functions
func GetAll() map[string]BuilderFunc

// BuilderFunc is a function that creates a new builder
type BuilderFunc func() Builder
```

### Example

```go
// Register a builder
Register("Person", func() Builder { return NewPersonBuilder() })

// Get a builder by name
builderFunc, ok := Get("Person")
if ok {
    builder := builderFunc()
    person := builder.Build()
}
```

## Model Builders

The library provides builders for all models in the system. Each builder follows the same pattern:

### PersonBuilder

```go
// PersonBuilder builds a Person model
type PersonBuilder struct {
    person *models.Person
    validationFuncs []func(*models.Person) error
}

// NewPersonBuilder creates a new PersonBuilder
func NewPersonBuilder() *PersonBuilder

// WithID sets the ID
func (b *PersonBuilder) WithID(id string) *PersonBuilder

// WithName sets the Name
func (b *PersonBuilder) WithName(name string) *PersonBuilder

// WithAge sets the Age
func (b *PersonBuilder) WithAge(age int) *PersonBuilder

// WithEmail sets the Email
func (b *PersonBuilder) WithEmail(email string) *PersonBuilder

// WithPhone sets the Phone
func (b *PersonBuilder) WithPhone(phone string) *PersonBuilder

// WithBirthdate sets the Birthdate
func (b *PersonBuilder) WithBirthdate(birthdate string) *PersonBuilder

// WithGender sets the Gender
func (b *PersonBuilder) WithGender(gender string) *PersonBuilder

// WithNationality sets the Nationality
func (b *PersonBuilder) WithNationality(nationality string) *PersonBuilder

// WithMaritalStatus sets the MaritalStatus
func (b *PersonBuilder) WithMaritalStatus(maritalStatus string) *PersonBuilder

// WithAddress sets the Address
func (b *PersonBuilder) WithAddress(addressBuilder *AddressBuilder) *PersonBuilder

// WithEducation sets the Education
func (b *PersonBuilder) WithEducation(educationBuilder *EducationBuilder) *PersonBuilder

// WithEmployment sets the Employment
func (b *PersonBuilder) WithEmployment(employmentBuilder *EmploymentBuilder) *PersonBuilder

// WithHealthProfile sets the HealthProfile
func (b *PersonBuilder) WithHealthProfile(healthProfileBuilder *HealthProfileBuilder) *PersonBuilder

// WithFinancialProfile sets the FinancialProfile
func (b *PersonBuilder) WithFinancialProfile(financialProfileBuilder *FinancialProfileBuilder) *PersonBuilder

// WithDigitalProfile sets the DigitalProfile
func (b *PersonBuilder) WithDigitalProfile(digitalProfileBuilder *DigitalProfileBuilder) *PersonBuilder

// WithMetadata adds a metadata entry
func (b *PersonBuilder) WithMetadata(key, value string) *PersonBuilder

// WithValidation adds a custom validation function
func (b *PersonBuilder) WithValidation(validationFunc func(*models.Person) error) *PersonBuilder

// Build builds the Person
func (b *PersonBuilder) Build() interface{}

// BuildWithValidation builds the Person and validates it
func (b *PersonBuilder) BuildWithValidation() (*models.Person, error)

// Clone creates a deep copy of the builder
func (b *PersonBuilder) Clone() *PersonBuilder

// FromModel creates a builder from an existing model
func PersonBuilderFromModel(model models.Person) *PersonBuilder
```

### AddressBuilder

```go
// AddressBuilder builds an Address model
type AddressBuilder struct {
    address *models.Address
    validationFuncs []func(*models.Address) error
}

// NewAddressBuilder creates a new AddressBuilder
func NewAddressBuilder() *AddressBuilder

// WithStreet sets the Street
func (b *AddressBuilder) WithStreet(street string) *AddressBuilder

// WithCity sets the City
func (b *AddressBuilder) WithCity(city string) *AddressBuilder

// WithState sets the State
func (b *AddressBuilder) WithState(state string) *AddressBuilder

// WithPostalCode sets the PostalCode
func (b *AddressBuilder) WithPostalCode(postalCode string) *AddressBuilder

// WithCountry sets the Country
func (b *AddressBuilder) WithCountry(country string) *AddressBuilder

// WithCoordinates sets the Coordinates
func (b *AddressBuilder) WithCoordinates(coordinatesBuilder *GeoLocationBuilder) *AddressBuilder

// WithType sets the Type
func (b *AddressBuilder) WithType(addressType string) *AddressBuilder

// WithIsPrimary sets the IsPrimary flag
func (b *AddressBuilder) WithIsPrimary(isPrimary bool) *AddressBuilder

// WithValidation adds a custom validation function
func (b *AddressBuilder) WithValidation(validationFunc func(*models.Address) error) *AddressBuilder

// Build builds the Address
func (b *AddressBuilder) Build() interface{}

// BuildWithValidation builds the Address and validates it
func (b *AddressBuilder) BuildWithValidation() (*models.Address, error)

// Clone creates a deep copy of the builder
func (b *AddressBuilder) Clone() *AddressBuilder

// FromModel creates a builder from an existing model
func AddressBuilderFromModel(model models.Address) *AddressBuilder
```

## Validation

The library provides validation support for all models. Validation can be performed in two ways:

1. Using the `BuildWithValidation` method on builders
2. Using the `Validate` method on models

### Builder Validation

```go
// Create a person with validation
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
```

### Model Validation

```go
// Create a person without validation
person := builders.NewPersonBuilder().
    WithID("123").
    WithName("John Doe").
    WithAge(30).
    WithEmail("john.doe@example.com").
    Build().(models.Person)

// Validate the person
if err := person.Validate(); err != nil {
    fmt.Printf("Validation error: %v\n", err)
    return
}
```

### Custom Validation

You can add custom validation rules to builders:

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

## Builder Generator

The library includes a code generator that can automatically generate builders for any struct.

### Command-Line Tool

```bash
# Generate builders for a single file
builder-gen -input models/person.go -output builders -models-package github.com/adil-faiyaz98/go-builder-kit/models

# Generate builders for all files in a directory
builder-gen -input models -output builders -recursive -models-package github.com/adil-faiyaz98/go-builder-kit/models
```

### Generator API

```go
// Generator generates builder code for structs
type Generator struct {
    Options Options
}

// Options contains configuration options for the generator
type Options struct {
    PackageName   string
    ModelsPackage string
    Verbose       bool
}

// NewGenerator creates a new Generator
func NewGenerator(opts Options) *Generator

// ParseFile parses a Go file and extracts struct information
func (g *Generator) ParseFile(filename string) ([]StructInfo, error)

// GenerateBuilder generates builder code for a struct
func (g *Generator) GenerateBuilder(structInfo StructInfo) (string, error)
```

### Example

```go
// Create a generator
gen := generator.NewGenerator(generator.Options{
    PackageName:   "builders",
    ModelsPackage: "github.com/adil-faiyaz98/go-builder-kit/models",
    Verbose:       true,
})

// Parse a file
structs, err := gen.ParseFile("models/person.go")
if err != nil {
    log.Fatalf("Error parsing file: %v", err)
}

// Generate builder for each struct
for _, s := range structs {
    code, err := gen.GenerateBuilder(s)
    if err != nil {
        log.Fatalf("Error generating builder: %v", err)
    }

    // Write builder to file
    outputFile := fmt.Sprintf("builders/%s_builder.go", generator.ToSnakeCase(s.Name))
    if err := os.WriteFile(outputFile, []byte(code), 0644); err != nil {
        log.Fatalf("Error writing file: %v", err)
    }
}
```
