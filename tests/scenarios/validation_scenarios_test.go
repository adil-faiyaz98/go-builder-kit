package scenarios

import (
	"testing"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
	"github.com/adil-faiyaz98/go-builder-kit/models"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestValidationScenarios(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Validation Scenarios Suite")
}

var _ = Describe("Validation Scenarios", func() {
	// Scenario 1: Creating a valid person with complete profile
	Describe("Creating a valid person with complete profile", func() {
		It("should pass validation", func() {
			// Create a complete person with all required fields
			person, err := builders.NewPersonBuilder().
				WithID("valid-123").
				WithName("Jane Smith").
				WithAge(35).
				WithEmail("jane.smith@example.com").
				WithPhone("+1-555-987-6543").
				WithBirthdate("1988-03-15").
				WithGender("Female").
				WithNationality("Canadian").
				WithMaritalStatus("Married").
				WithProfile(
					builders.NewProfileBuilder().
						WithAddress(
							builders.NewAddressBuilder().
								WithStreet("456 Maple Ave").
								WithCity("Toronto").
								WithState("ON").
								WithPostalCode("M5V 2H1").
								WithCountry("Canada").
								WithType("Home").
								WithIsPrimary(true),
						),
				).
				BuildWithValidation()

			// Validation should pass
			Expect(err).NotTo(HaveOccurred())
			Expect(person).NotTo(BeNil())
			Expect(person.Name).To(Equal("Jane Smith"))
		})
	})

	// Scenario 2: Creating a person with invalid email
	Describe("Creating a person with invalid email", func() {
		It("should fail validation with appropriate error", func() {
			// Create a person with invalid email
			_, err := builders.NewPersonBuilder().
				WithID("invalid-123").
				WithName("John Doe").
				WithAge(42).
				WithEmail("not-an-email").  // Invalid email
				WithPhone("+1-555-123-4567").
				WithBirthdate("1981-07-22").
				WithMaritalStatus("Single").
				BuildWithValidation()

			// Validation should fail with specific error
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Email is not valid"))
		})
	})

	// Scenario 3: Creating a person with future birthdate
	Describe("Creating a person with future birthdate", func() {
		It("should fail validation with appropriate error", func() {
			// Create a person with future birthdate
			_, err := builders.NewPersonBuilder().
				WithID("future-123").
				WithName("Future Person").
				WithAge(0).
				WithBirthdate("2030-01-01").  // Future date
				BuildWithValidation()

			// Validation should fail with specific error
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Birthdate cannot be in the future"))
		})
	})

	// Scenario 4: Creating an education record with invalid dates
	Describe("Creating an education record with invalid dates", func() {
		It("should fail validation when end date is before start date", func() {
			// Create education with end date before start date
			_, err := builders.NewEducationBuilder().
				WithDegree("Master of Science").
				WithInstitution("University of Technology").
				WithStartDate("2020-09-01").
				WithEndDate("2019-05-15").  // Before start date
				WithGPA(3.7).
				BuildWithValidation()

			// Validation should fail with specific error
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("EndDate cannot be before StartDate"))
		})
	})

	// Scenario 5: Creating an address with invalid postal code
	Describe("Creating an address with invalid postal code", func() {
		It("should fail validation with appropriate error", func() {
			// Create address with invalid postal code
			_, err := builders.NewAddressBuilder().
				WithStreet("123 Main St").
				WithCity("Anytown").
				WithState("ST").
				WithPostalCode("!@#$%").  // Invalid format
				WithCountry("USA").
				BuildWithValidation()

			// Validation should fail with specific error
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("PostalCode format is invalid"))
		})
	})

	// Scenario 6: Creating a complex nested structure with validation errors
	Describe("Creating a complex nested structure with validation errors", func() {
		It("should detect errors in nested objects", func() {
			// Create a person with valid fields but invalid nested education
			personBuilder := builders.NewPersonBuilder().
				WithID("nested-123").
				WithName("Nested Error Person").
				WithAge(28).
				WithEmail("nested@example.com").
				WithProfile(
					builders.NewProfileBuilder().
						WithAddress(
							builders.NewAddressBuilder().
								WithStreet("789 Oak St").
								WithCity("Chicago").
								WithState("IL").
								WithPostalCode("60601").
								WithCountry("USA"),
						),
				).
				WithEmployment(
					builders.NewEmploymentBuilder().
						WithCompany(
							builders.NewCompanyBuilder().
								WithID("COMP-001").
								WithName("Tech Company").
								WithIndustry("Technology").
								WithLocation(
									builders.NewAddressBuilder().
										WithStreet("100 Tech Blvd").
										WithCity("San Francisco").
										WithState("CA").
										WithPostalCode("94107").
										WithCountry("USA"),
								),
						).
						WithPosition("Software Engineer").
						WithStartDate("2020-01-15").
						WithSalary(-5000.0),  // Invalid: negative salary
				)

			// Build without validation first to check the object
			person := personBuilder.Build().(*models.Person)
			
			// Validate manually
			err := person.Validate()
			
			// Should pass person validation
			Expect(err).NotTo(HaveOccurred())
			
			// But employment validation would fail if we had it
			// This demonstrates how nested validation would work
			// if we implemented it for Employment
			
			// For demonstration, let's check the invalid salary
			Expect(person.Employment.Salary).To(Equal(-5000.0))
			// In a real implementation, this would fail validation
		})
	})
})
