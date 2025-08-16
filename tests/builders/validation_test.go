package builders_test

import (
	"fmt"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
	"github.com/adil-faiyaz98/go-builder-kit/models"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Builder Validation", func() {
	Context("Custom Validation Functions", func() {
		It("should validate a person with custom validation", func() {
			// Create a person builder with custom validation
			personBuilder := builders.NewPersonBuilder().
				WithID("123").
				WithName("John Doe").
				WithAge(30).
				WithEmail("john.doe@example.com")

			// Add custom validation function
			personBuilder.WithValidation(func(p *models.Person) error {
				if p.Age < 0 {
					return fmt.Errorf("age cannot be negative")
				}
				if p.Name == "" {
					return fmt.Errorf("name cannot be empty")
				}
				return nil
			})

			// Build and validate the person
			person, err := personBuilder.BuildAndValidate()

			// Verify validation passed
			Expect(err).To(BeNil())
			Expect(person).NotTo(BeNil())
			Expect(person.Name).To(Equal("John Doe"))
			Expect(person.Age).To(Equal(30))
		})

		It("should fail validation when validation function returns an error", func() {
			// Create a person builder with invalid age
			personBuilder := builders.NewPersonBuilder().
				WithID("123").
				WithName("John Doe").
				WithAge(-1) // Invalid age

			// Add custom validation function
			personBuilder.WithValidation(func(p *models.Person) error {
				if p.Age < 0 {
					return fmt.Errorf("age cannot be negative")
				}
				return nil
			})

			// Build and validate the person
			person, err := personBuilder.BuildAndValidate()

			// Verify validation failed
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("age cannot be negative"))
			Expect(person).To(BeNil())
		})

		It("should panic when using MustBuild with invalid data", func() {
			// Create a person builder with invalid name
			personBuilder := builders.NewPersonBuilder().
				WithID("123").
				WithName(""). // Invalid name
				WithAge(30)

			// Add custom validation function
			personBuilder.WithValidation(func(p *models.Person) error {
				if p.Name == "" {
					return fmt.Errorf("name cannot be empty")
				}
				return nil
			})

			// Verify MustBuild panics
			Expect(func() {
				personBuilder.MustBuild()
			}).To(Panic())
		})
	})

	Context("Table-Driven Tests", func() {
		DescribeTable("Person validation scenarios",
			func(name string, age int, expectError bool) {
				// Create a person builder
				personBuilder := builders.NewPersonBuilder().
					WithID("123").
					WithName(name).
					WithAge(age)

				// Add custom validation function
				personBuilder.WithValidation(func(p *models.Person) error {
					if p.Age < 0 {
						return fmt.Errorf("age cannot be negative")
					}
					if p.Name == "" {
						return fmt.Errorf("name cannot be empty")
					}
					return nil
				})

				// Build and validate the person
				_, err := personBuilder.BuildAndValidate()

				// Verify validation result
				if expectError {
					Expect(err).NotTo(BeNil())
				} else {
					Expect(err).To(BeNil())
				}
			},
			Entry("Valid person", "John Doe", 30, false),
			Entry("Empty name", "", 30, true),
			Entry("Negative age", "John Doe", -1, true),
			Entry("Both invalid", "", -1, true),
		)
	})
})
