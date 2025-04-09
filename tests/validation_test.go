package tests

import (
	"fmt"
	"strings"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
	"github.com/adil-faiyaz98/go-builder-kit/models"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validation", func() {
	Context("Person Validation", func() {
		It("should validate required fields", func() {
			// Create a Person builder with missing required fields
			personBuilder := builders.NewPersonBuilder()
			// Missing ID and Name

			// Build and validate should fail
			_, err := personBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("ID cannot be empty"))
			Expect(err.Error()).To(ContainSubstring("Name cannot be empty"))
		})

		It("should validate email format", func() {
			// Create a Person builder with invalid email
			personBuilder := builders.NewPersonBuilder().
				WithID("P12345").
				WithName("John Doe").
				WithEmail("invalid-email") // Not a valid email format

			// Build and validate should fail
			_, err := personBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Email is not valid"))
		})

		It("should validate phone format", func() {
			// Create a Person builder with invalid phone
			personBuilder := builders.NewPersonBuilder().
				WithID("P12345").
				WithName("John Doe").
				WithEmail("john.doe@example.com").
				WithPhone("123") // Too short

			// Build and validate should fail
			_, err := personBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Phone number is not valid"))
		})

		It("should validate birthdate format", func() {
			// Create a Person builder with invalid birthdate
			personBuilder := builders.NewPersonBuilder().
				WithID("P12345").
				WithName("John Doe").
				WithEmail("john.doe@example.com").
				WithBirthdate("01/01/2000") // Wrong format, should be YYYY-MM-DD

			// Build and validate should fail
			_, err := personBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Birthdate must be in the format YYYY-MM-DD"))
		})

		It("should validate future birthdate", func() {
			// Create a Person builder with future birthdate
			personBuilder := builders.NewPersonBuilder().
				WithID("P12345").
				WithName("John Doe").
				WithEmail("john.doe@example.com").
				WithBirthdate("2100-01-01") // Future date

			// Build and validate should fail
			_, err := personBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Birthdate cannot be in the future"))
		})

		It("should validate gender values", func() {
			// Create a Person builder with invalid gender
			personBuilder := builders.NewPersonBuilder().
				WithID("P12345").
				WithName("John Doe").
				WithEmail("john.doe@example.com").
				WithGender("invalid") // Not a valid gender option

			// Build and validate should fail
			_, err := personBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Gender must be one of"))
		})

		It("should validate marital status values", func() {
			// Create a Person builder with invalid marital status
			personBuilder := builders.NewPersonBuilder().
				WithID("P12345").
				WithName("John Doe").
				WithEmail("john.doe@example.com").
				WithMaritalStatus("invalid") // Not a valid marital status option

			// Build and validate should fail
			_, err := personBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("MaritalStatus must be one of"))
		})

		It("should pass validation with valid data", func() {
			// Create a Person builder with valid data
			personBuilder := builders.NewPersonBuilder().
				WithID("P12345").
				WithName("John Doe").
				WithEmail("john.doe@example.com").
				WithPhone("555-123-4567").
				WithBirthdate("1995-05-15").
				WithGender("male").
				WithMaritalStatus("single")

			// Build and validate should succeed
			person, err := personBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(person).NotTo(BeNil())
			Expect(person.ID).To(Equal("P12345"))
			Expect(person.Name).To(Equal("John Doe"))
			Expect(person.Email).To(Equal("john.doe@example.com"))
			Expect(person.Phone).To(Equal("555-123-4567"))
			Expect(person.Birthdate).To(Equal("1995-05-15"))
			Expect(person.Gender).To(Equal("male"))
			Expect(person.MaritalStatus).To(Equal("single"))
		})
	})

	Context("Custom Validation", func() {
		It("should support custom validation functions", func() {
			// Create a Person builder with custom validation
			personBuilder := builders.NewPersonBuilder().
				WithID("P12345").
				WithName("John Doe").
				WithEmail("john.doe@example.com").
				WithValidation(func(p *models.Person) error {
					if p.Name != "John Smith" {
						return fmt.Errorf("name must be 'John Smith'")
					}
					return nil
				})

			// Build and validate should fail
			_, err := personBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("name must be 'John Smith'"))

			// Fix the name and try again
			personBuilder.WithName("John Smith")
			person, err := personBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(person).NotTo(BeNil())
			Expect(person.Name).To(Equal("John Smith"))
		})

		It("should support multiple custom validation functions", func() {
			// Create a Person builder with multiple custom validations
			personBuilder := builders.NewPersonBuilder().
				WithID("P12345").
				WithName("John Doe").
				WithEmail("john.doe@example.com").
				WithValidation(func(p *models.Person) error {
					if len(p.ID) < 6 {
						return fmt.Errorf("ID must be at least 6 characters")
					}
					return nil
				}).
				WithValidation(func(p *models.Person) error {
					if !strings.HasPrefix(p.Email, "john.") {
						return fmt.Errorf("email must start with 'john.'")
					}
					return nil
				})

			// Build and validate should succeed
			person, err := personBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(person).NotTo(BeNil())

			// Change email to invalid and try again
			personBuilder.WithEmail("jane.doe@example.com")
			_, err = personBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("email must start with 'john.'"))
		})
	})
})
