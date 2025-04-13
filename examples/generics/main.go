package main

import (
	"fmt"

	"github.com/adil-faiyaz98/go-builder-kit/pkg/builder"
)

// Person is a simple struct for demonstration
type Person struct {
	ID    string
	Name  string
	Age   int
	Email string
}

// PersonBuilder is a generic builder for Person
type PersonBuilder struct {
	*builder.BaseGenericBuilder[Person]
	person Person
}

// NewPersonBuilder creates a new PersonBuilder
func NewPersonBuilder() *PersonBuilder {
	b := &PersonBuilder{
		person: Person{},
	}
	
	b.BaseGenericBuilder = builder.NewBaseGenericBuilder(func() Person {
		return b.person
	})
	
	return b
}

// WithID sets the ID of the person
func (b *PersonBuilder) WithID(id string) *PersonBuilder {
	b.person.ID = id
	return b
}

// WithName sets the name of the person
func (b *PersonBuilder) WithName(name string) *PersonBuilder {
	b.person.Name = name
	return b
}

// WithAge sets the age of the person
func (b *PersonBuilder) WithAge(age int) *PersonBuilder {
	b.person.Age = age
	return b
}

// WithEmail sets the email of the person
func (b *PersonBuilder) WithEmail(email string) *PersonBuilder {
	b.person.Email = email
	return b
}

// Clone creates a deep copy of the builder
func (b *PersonBuilder) Clone() *PersonBuilder {
	clone := &PersonBuilder{
		person: b.person,
	}
	
	clone.BaseGenericBuilder = builder.NewBaseGenericBuilder(func() Person {
		return clone.person
	})
	
	return clone
}

func main() {
	// Create a person builder
	personBuilder := NewPersonBuilder().
		WithID("123").
		WithName("John Doe").
		WithAge(30).
		WithEmail("john.doe@example.com")
	
	// Add validation
	personBuilder.WithValidation(func(p Person) error {
		if p.ID == "" {
			return fmt.Errorf("ID cannot be empty")
		}
		if p.Name == "" {
			return fmt.Errorf("name cannot be empty")
		}
		if p.Age < 0 {
			return fmt.Errorf("age cannot be negative")
		}
		return nil
	})
	
	// Build and validate
	person, err := personBuilder.BuildAndValidate()
	if err != nil {
		fmt.Println("Validation failed:", err)
		return
	}
	
	// Use the person
	fmt.Println("Person created:", person.Name)
	
	// Clone the builder and modify it
	clonedBuilder := personBuilder.Clone().
		WithName("Jane Doe").
		WithAge(28)
	
	// Build the cloned person
	clonedPerson := clonedBuilder.BuildPtr()
	
	// Use the cloned person
	fmt.Println("Cloned person created:", clonedPerson.Name)
}
