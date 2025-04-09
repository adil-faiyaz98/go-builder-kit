# Builder Composition and Inheritance

This document demonstrates how to use composition and inheritance patterns with Go Builder Kit.

## Builder Composition

Builder composition allows you to create complex objects by combining multiple builders:

```go
package main

import (
    "fmt"

    "github.com/adil-faiyaz98/go-builder-kit/builders"
    "github.com/adil-faiyaz98/go-builder-kit/models"
)

// CreateHomeAddress is a factory function that creates a standard home address
func CreateHomeAddress(street, city, state, postalCode, country string) *builders.AddressBuilder {
    return builders.NewAddressBuilder().
        WithStreet(street).
        WithCity(city).
        WithState(state).
        WithPostalCode(postalCode).
        WithCountry(country).
        WithType("Home").
        WithIsPrimary(true)
}

// CreateWorkAddress is a factory function that creates a standard work address
func CreateWorkAddress(street, city, state, postalCode, country string) *builders.AddressBuilder {
    return builders.NewAddressBuilder().
        WithStreet(street).
        WithCity(city).
        WithState(state).
        WithPostalCode(postalCode).
        WithCountry(country).
        WithType("Work").
        WithIsPrimary(false)
}

// CreateBasicPerson is a factory function that creates a person with basic information
func CreateBasicPerson(id, name string, age int, email string) *builders.PersonBuilder {
    return builders.NewPersonBuilder().
        WithID(id).
        WithName(name).
        WithAge(age).
        WithEmail(email)
}

func main() {
    // Create a person with home and work addresses using composition
    person := CreateBasicPerson("123", "John Doe", 30, "john.doe@example.com").
        WithAddress(
            CreateHomeAddress("123 Main St", "New York", "NY", "10001", "USA"),
        ).
        WithAddress(
            CreateWorkAddress("456 Business Ave", "New York", "NY", "10002", "USA"),
        ).
        Build().(models.Person)

    fmt.Printf("Created person: %s\n", person.Name)
    fmt.Printf("Home address: %s, %s\n", person.Address.Street, person.Address.City)
    
    // Create another person with the same pattern
    anotherPerson := CreateBasicPerson("456", "Jane Smith", 28, "jane.smith@example.com").
        WithAddress(
            CreateHomeAddress("789 Residential St", "Boston", "MA", "02115", "USA"),
        ).
        WithAddress(
            CreateWorkAddress("101 Office Blvd", "Boston", "MA", "02116", "USA"),
        ).
        Build().(models.Person)

    fmt.Printf("Created another person: %s\n", anotherPerson.Name)
    fmt.Printf("Home address: %s, %s\n", anotherPerson.Address.Street, anotherPerson.Address.City)
}
```

## Builder Inheritance

You can use Go's embedding to create "inherited" builders:

```go
package main

import (
    "fmt"

    "github.com/adil-faiyaz98/go-builder-kit/builders"
    "github.com/adil-faiyaz98/go-builder-kit/models"
)

// EmployeeBuilder extends PersonBuilder with employee-specific methods
type EmployeeBuilder struct {
    *builders.PersonBuilder
}

// NewEmployeeBuilder creates a new EmployeeBuilder
func NewEmployeeBuilder() *EmployeeBuilder {
    return &EmployeeBuilder{
        PersonBuilder: builders.NewPersonBuilder(),
    }
}

// WithEmployeeID sets the employee ID
func (b *EmployeeBuilder) WithEmployeeID(employeeID string) *EmployeeBuilder {
    b.WithID("EMP-" + employeeID)
    return b
}

// WithDepartment sets the department
func (b *EmployeeBuilder) WithDepartment(department string) *EmployeeBuilder {
    b.WithEmployment(
        builders.NewEmploymentBuilder().
            WithDepartment(department),
    )
    return b
}

// WithPosition sets the position
func (b *EmployeeBuilder) WithPosition(position string) *EmployeeBuilder {
    // Get the existing employment or create a new one
    var employment *models.Employment
    if b.PersonBuilder.Build().(models.Person).Employment != nil {
        employment = b.PersonBuilder.Build().(models.Person).Employment
    } else {
        employment = &models.Employment{}
    }
    
    // Update the position
    employment.Position = position
    
    // Set the updated employment
    b.WithEmployment(
        builders.EmploymentBuilderFromModel(*employment),
    )
    
    return b
}

// WithSalary sets the salary
func (b *EmployeeBuilder) WithSalary(salary float64) *EmployeeBuilder {
    // Get the existing employment or create a new one
    var employment *models.Employment
    if b.PersonBuilder.Build().(models.Person).Employment != nil {
        employment = b.PersonBuilder.Build().(models.Person).Employment
    } else {
        employment = &models.Employment{}
    }
    
    // Update the salary
    employment.Salary = salary
    
    // Set the updated employment
    b.WithEmployment(
        builders.EmploymentBuilderFromModel(*employment),
    )
    
    return b
}

func main() {
    // Create an employee using the extended builder
    employee := NewEmployeeBuilder().
        WithName("John Doe").
        WithAge(30).
        WithEmail("john.doe@example.com").
        WithEmployeeID("12345").
        WithDepartment("Engineering").
        WithPosition("Software Engineer").
        WithSalary(100000.0).
        Build().(models.Person)

    fmt.Printf("Created employee: %s\n", employee.Name)
    fmt.Printf("Employee ID: %s\n", employee.ID)
    fmt.Printf("Department: %s\n", employee.Employment.Department)
    fmt.Printf("Position: %s\n", employee.Employment.Position)
    fmt.Printf("Salary: $%.2f\n", employee.Employment.Salary)
}
```

## Combining Composition and Inheritance

You can combine both patterns for maximum flexibility:

```go
package main

import (
    "fmt"

    "github.com/adil-faiyaz98/go-builder-kit/builders"
    "github.com/adil-faiyaz98/go-builder-kit/models"
)

// EmployeeBuilder extends PersonBuilder with employee-specific methods
type EmployeeBuilder struct {
    *builders.PersonBuilder
}

// NewEmployeeBuilder creates a new EmployeeBuilder
func NewEmployeeBuilder() *EmployeeBuilder {
    return &EmployeeBuilder{
        PersonBuilder: builders.NewPersonBuilder(),
    }
}

// WithEmployeeID sets the employee ID
func (b *EmployeeBuilder) WithEmployeeID(employeeID string) *EmployeeBuilder {
    b.WithID("EMP-" + employeeID)
    return b
}

// WithDepartment sets the department
func (b *EmployeeBuilder) WithDepartment(department string) *EmployeeBuilder {
    b.WithEmployment(
        builders.NewEmploymentBuilder().
            WithDepartment(department),
    )
    return b
}

// WithPosition sets the position
func (b *EmployeeBuilder) WithPosition(position string) *EmployeeBuilder {
    // Get the existing employment or create a new one
    var employment *models.Employment
    if b.PersonBuilder.Build().(models.Person).Employment != nil {
        employment = b.PersonBuilder.Build().(models.Person).Employment
    } else {
        employment = &models.Employment{}
    }
    
    // Update the position
    employment.Position = position
    
    // Set the updated employment
    b.WithEmployment(
        builders.EmploymentBuilderFromModel(*employment),
    )
    
    return b
}

// ManagerBuilder extends EmployeeBuilder with manager-specific methods
type ManagerBuilder struct {
    *EmployeeBuilder
}

// NewManagerBuilder creates a new ManagerBuilder
func NewManagerBuilder() *ManagerBuilder {
    employeeBuilder := NewEmployeeBuilder()
    employeeBuilder.WithPosition("Manager")
    
    return &ManagerBuilder{
        EmployeeBuilder: employeeBuilder,
    }
}

// WithTeamSize sets the team size
func (b *ManagerBuilder) WithTeamSize(teamSize int) *ManagerBuilder {
    // Add team size as metadata
    b.WithMetadata("teamSize", fmt.Sprintf("%d", teamSize))
    return b
}

// WithTeamMember adds a team member
func (b *ManagerBuilder) WithTeamMember(teamMember *builders.PersonBuilder) *ManagerBuilder {
    // In a real implementation, you would add the team member to a team collection
    // For this example, we'll just add them as metadata
    person := teamMember.Build().(models.Person)
    b.WithMetadata("teamMember_"+person.ID, person.Name)
    return b
}

func main() {
    // Create a manager using the extended builder
    manager := NewManagerBuilder().
        WithName("Jane Smith").
        WithAge(35).
        WithEmail("jane.smith@example.com").
        WithEmployeeID("M-789").
        WithDepartment("Engineering").
        WithTeamSize(5).
        WithTeamMember(
            builders.NewPersonBuilder().
                WithID("EMP-123").
                WithName("John Doe"),
        ).
        WithTeamMember(
            builders.NewPersonBuilder().
                WithID("EMP-456").
                WithName("Alice Johnson"),
        ).
        Build().(models.Person)

    fmt.Printf("Created manager: %s\n", manager.Name)
    fmt.Printf("Employee ID: %s\n", manager.ID)
    fmt.Printf("Department: %s\n", manager.Employment.Department)
    fmt.Printf("Position: %s\n", manager.Employment.Position)
    fmt.Printf("Team size: %s\n", manager.Metadata["teamSize"])
    fmt.Printf("Team member 1: %s\n", manager.Metadata["teamMember_EMP-123"])
    fmt.Printf("Team member 2: %s\n", manager.Metadata["teamMember_EMP-456"])
}
```

## Factory Pattern with Builders

You can use the factory pattern with builders to create objects with predefined configurations:

```go
package main

import (
    "fmt"

    "github.com/adil-faiyaz98/go-builder-kit/builders"
    "github.com/adil-faiyaz98/go-builder-kit/models"
)

// PersonFactory provides factory methods for creating different types of persons
type PersonFactory struct{}

// CreateEmployee creates an employee
func (f *PersonFactory) CreateEmployee(name string, age int, department string, position string) models.Person {
    return builders.NewPersonBuilder().
        WithName(name).
        WithAge(age).
        WithEmployment(
            builders.NewEmploymentBuilder().
                WithDepartment(department).
                WithPosition(position),
        ).
        Build().(models.Person)
}

// CreateCustomer creates a customer
func (f *PersonFactory) CreateCustomer(name string, age int, customerID string) models.Person {
    return builders.NewPersonBuilder().
        WithName(name).
        WithAge(age).
        WithID("CUST-" + customerID).
        WithMetadata("customerType", "Regular").
        Build().(models.Person)
}

// CreateVIPCustomer creates a VIP customer
func (f *PersonFactory) CreateVIPCustomer(name string, age int, customerID string) models.Person {
    return builders.NewPersonBuilder().
        WithName(name).
        WithAge(age).
        WithID("VIP-" + customerID).
        WithMetadata("customerType", "VIP").
        WithMetadata("vipLevel", "Gold").
        Build().(models.Person)
}

func main() {
    factory := &PersonFactory{}

    // Create an employee
    employee := factory.CreateEmployee("John Doe", 30, "Engineering", "Software Engineer")
    fmt.Printf("Created employee: %s\n", employee.Name)
    fmt.Printf("Department: %s\n", employee.Employment.Department)
    fmt.Printf("Position: %s\n", employee.Employment.Position)

    // Create a regular customer
    customer := factory.CreateCustomer("Alice Johnson", 25, "C-123")
    fmt.Printf("Created customer: %s\n", customer.Name)
    fmt.Printf("Customer ID: %s\n", customer.ID)
    fmt.Printf("Customer type: %s\n", customer.Metadata["customerType"])

    // Create a VIP customer
    vipCustomer := factory.CreateVIPCustomer("Bob Smith", 40, "C-456")
    fmt.Printf("Created VIP customer: %s\n", vipCustomer.Name)
    fmt.Printf("Customer ID: %s\n", vipCustomer.ID)
    fmt.Printf("Customer type: %s\n", vipCustomer.Metadata["customerType"])
    fmt.Printf("VIP level: %s\n", vipCustomer.Metadata["vipLevel"])
}
```
