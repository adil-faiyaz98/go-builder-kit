package models_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

var _ = Describe("Education", func() {
	var (
		educationBuilder *builders.EducationBuilder
		addressBuilder   *builders.AddressBuilder
		courseBuilder    *builders.CourseBuilder
	)

	BeforeEach(func() {
		addressBuilder = builders.NewAddressBuilder().
			WithStreet("123 University Ave").
			WithCity("Cambridge").
			WithState("MA").
			WithPostalCode("02138").
			WithCountry("USA")

		courseBuilder = builders.NewCourseBuilder().
			WithCode("CS101").
			WithName("Introduction to Computer Science").
			WithCredits(3.0).
			WithYear(2020).
			WithSemester("Fall").
			WithGrade("A")

		educationBuilder = builders.NewEducationBuilder().
			WithDegree("Bachelor of Science").
			WithInstitution("Harvard University").
			WithLocation(addressBuilder).
			WithStartDate("2018-09-01").
			WithEndDate("2022-05-15").
			WithGPA(3.8).
			WithMajor("Computer Science").
			WithMinor("Mathematics").
			WithCourse(courseBuilder)
	})

	Describe("Validation", func() {
		Context("when all fields are valid", func() {
			It("should pass validation", func() {
				education := educationBuilder.Build().(*models.Education)
				err := education.Validate()
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when Degree is empty", func() {
			It("should fail validation", func() {
				education := educationBuilder.WithDegree("").Build().(*models.Education)
				err := education.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Degree cannot be empty"))
			})
		})

		Context("when Institution is empty", func() {
			It("should fail validation", func() {
				education := educationBuilder.WithInstitution("").Build().(*models.Education)
				err := education.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Institution cannot be empty"))
			})
		})

		Context("when StartDate is in the future", func() {
			It("should fail validation", func() {
				futureDate := time.Now().AddDate(1, 0, 0).Format("2006-01-02")
				education := educationBuilder.WithStartDate(futureDate).Build().(*models.Education)
				err := education.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("StartDate cannot be in the future"))
			})
		})

		Context("when StartDate is too far in the past", func() {
			It("should fail validation", func() {
				veryOldDate := time.Now().AddDate(-150, 0, 0).Format("2006-01-02")
				education := educationBuilder.WithStartDate(veryOldDate).Build().(*models.Education)
				err := education.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("StartDate cannot be more than 100 years ago"))
			})
		})

		Context("when EndDate is before StartDate", func() {
			It("should fail validation", func() {
				startDate := "2020-01-01"
				endDate := "2019-01-01" // Before start date
				education := educationBuilder.
					WithStartDate(startDate).
					WithEndDate(endDate).
					Build().(models.Education)
				err := education.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("EndDate cannot be before StartDate"))
			})
		})

		Context("when GPA is negative", func() {
			It("should fail validation", func() {
				education := educationBuilder.WithGPA(-1.0).Build().(*models.Education)
				err := education.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("GPA cannot be negative"))
			})
		})

		Context("when GPA is too high", func() {
			It("should fail validation", func() {
				education := educationBuilder.WithGPA(4.5).Build().(*models.Education)
				err := education.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("GPA cannot be greater than 4.0"))
			})
		})

		Context("when Course is invalid", func() {
			It("should fail validation with empty code", func() {
				invalidCourseBuilder := builders.NewCourseBuilder().
					WithCode("").
					WithName("Invalid Course")

				education := educationBuilder.
					WithCourse(invalidCourseBuilder).
					Build().(models.Education)

				err := education.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Code cannot be empty"))
			})

			It("should fail validation with future year", func() {
				futureYear := time.Now().Year() + 5
				invalidCourseBuilder := builders.NewCourseBuilder().
					WithCode("CS999").
					WithName("Future Course").
					WithYear(futureYear)

				education := educationBuilder.
					WithCourse(invalidCourseBuilder).
					Build().(models.Education)

				err := education.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Year cannot be after"))
			})
		})

		Context("when using BuildWithValidation", func() {
			It("should return error for invalid education", func() {
				_, err := educationBuilder.WithGPA(-1.0).BuildWithValidation()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("GPA cannot be negative"))
			})

			It("should return valid education when all fields are valid", func() {
				education, err := educationBuilder.BuildWithValidation()
				Expect(err).NotTo(HaveOccurred())
				Expect(education).NotTo(BeNil())
				Expect(education.Institution).To(Equal("Harvard University"))
			})
		})
	})
})

