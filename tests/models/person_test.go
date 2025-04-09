package models_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

var _ = Describe("Person", func() {
	var (
		personBuilder *builders.PersonBuilder
	)

	BeforeEach(func() {
		personBuilder = builders.NewPersonBuilder().
			WithID("test-123").
			WithName("John Doe").
			WithAge(30).
			WithEmail("john.doe@example.com").
			WithPhone("+1-555-123-4567").
			WithBirthdate("1993-05-15").
			WithGender("Male").
			WithNationality("American").
			WithMaritalStatus("Single").
			WithEducation(
				builders.NewEducationBuilder().
					WithDegree("Bachelor of Science").
					WithInstitution("Harvard University").
					WithStartDate("2018-09-01").
					WithEndDate("2022-05-15").
					WithGPA(3.8).
					WithMajor("Computer Science"),
			)
	})

	Describe("Validation", func() {
		Context("when all fields are valid", func() {
			It("should pass validation", func() {
				person := personBuilder.Build().(*models.Person)
				err := person.Validate()
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when ID is empty", func() {
			It("should fail validation", func() {
				person := personBuilder.WithID("").Build().(*models.Person)
				err := person.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("ID cannot be empty"))
			})
		})

		Context("when Name is empty", func() {
			It("should fail validation", func() {
				person := personBuilder.WithName("").Build().(*models.Person)
				err := person.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Name cannot be empty"))
			})
		})

		Context("when Name is too short", func() {
			It("should fail validation", func() {
				person := personBuilder.WithName("A").Build().(*models.Person)
				err := person.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Name must be at least 2 characters long"))
			})
		})

		Context("when Age is negative", func() {
			It("should fail validation", func() {
				person := personBuilder.WithAge(-1).Build().(*models.Person)
				err := person.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Age cannot be negative"))
			})
		})

		Context("when Age is too high", func() {
			It("should fail validation", func() {
				person := personBuilder.WithAge(151).Build().(*models.Person)
				err := person.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Age cannot be greater than 150"))
			})
		})

		Context("when Email is invalid", func() {
			It("should fail validation", func() {
				person := personBuilder.WithEmail("not-an-email").Build().(*models.Person)
				err := person.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Email is not valid"))
			})
		})

		Context("when Phone is invalid", func() {
			It("should fail validation", func() {
				person := personBuilder.WithPhone("123").Build().(*models.Person)
				err := person.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Phone number is not valid"))
			})
		})

		Context("when Birthdate is in the future", func() {
			It("should fail validation", func() {
				futureDate := time.Now().AddDate(1, 0, 0).Format("2006-01-02")
				person := personBuilder.WithBirthdate(futureDate).Build().(*models.Person)
				err := person.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Birthdate cannot be in the future"))
			})
		})

		Context("when Birthdate is too far in the past", func() {
			It("should fail validation", func() {
				veryOldDate := time.Now().AddDate(-200, 0, 0).Format("2006-01-02")
				person := personBuilder.WithBirthdate(veryOldDate).Build().(*models.Person)
				err := person.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Birthdate cannot be more than 150 years ago"))
			})
		})

		Context("when Birthdate format is invalid", func() {
			It("should fail validation", func() {
				person := personBuilder.WithBirthdate("05/15/1993").Build().(*models.Person)
				err := person.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Birthdate must be in the format YYYY-MM-DD"))
			})
		})

		Context("when MaritalStatus is invalid", func() {
			It("should fail validation", func() {
				person := personBuilder.WithMaritalStatus("Unknown").Build().(*models.Person)
				err := person.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("MaritalStatus must be one of: Single, Married, Divorced, Widowed, Separated"))
			})
		})

		Context("when using BuildWithValidation", func() {
			It("should return error for invalid person", func() {
				_, err := personBuilder.WithAge(-1).BuildWithValidation()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Age cannot be negative"))
			})

			It("should return valid person when all fields are valid", func() {
				person, err := personBuilder.BuildWithValidation()
				Expect(err).NotTo(HaveOccurred())
				Expect(person).NotTo(BeNil())
				Expect(person.Name).To(Equal("John Doe"))
			})
		})
	})
})

