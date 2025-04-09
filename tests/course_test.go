package tests

import (
	"github.com/adil-faiyaz98/go-builder-kit/builders"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Course", func() {
	Context("Course Builder", func() {
		It("should build a valid Course with all fields", func() {
			// Create a Course builder
			courseBuilder := builders.NewCourseBuilder().
				WithCode("CS101").
				WithName("Introduction to Computer Science").
				WithDescription("Basic programming concepts").
				WithCredits(3.0).
				WithGrade("A").
				WithSemester("Fall").
				WithYear(2020).
				WithInstructor("Dr. Smith")

			// Build the Course
			course, err := courseBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(course).NotTo(BeNil())
			Expect(course.Code).To(Equal("CS101"))
			Expect(course.Name).To(Equal("Introduction to Computer Science"))
			Expect(course.Description).To(Equal("Basic programming concepts"))
			Expect(course.Credits).To(Equal(3.0))
			Expect(course.Grade).To(Equal("A"))
			Expect(course.Semester).To(Equal("Fall"))
			Expect(course.Year).To(Equal(2020))
			Expect(course.Instructor).To(Equal("Dr. Smith"))
		})

		It("should validate required fields", func() {
			// Create a Course builder with missing required fields
			courseBuilder := builders.NewCourseBuilder()
			// Missing Code and Name

			// Build and validate should fail
			_, err := courseBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Code cannot be empty"))
			Expect(err.Error()).To(ContainSubstring("Name cannot be empty"))
		})

		It("should validate credits are positive", func() {
			// Create a Course builder with negative credits
			courseBuilder := builders.NewCourseBuilder().
				WithCode("CS101").
				WithName("Introduction to Computer Science").
				WithCredits(-1.0) // Negative credits

			// Build and validate should fail
			_, err := courseBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Credits cannot be negative"))
		})

		It("should validate year is reasonable", func() {
			// Create a Course builder with unreasonable year
			courseBuilder := builders.NewCourseBuilder().
				WithCode("CS101").
				WithName("Introduction to Computer Science").
				WithYear(1800). // Too old
				WithGrade("Z")  // Invalid grade

			// Build and validate should fail
			_, err := courseBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Year cannot be before 1900"))
		})

		It("should validate year is reasonable", func() {
			// Create a Course builder with unreasonable year
			courseBuilder := builders.NewCourseBuilder().
				WithCode("CS101").
				WithName("Introduction to Computer Science").
				WithYear(1800) // Too old

			// Build and validate should fail
			_, err := courseBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Year cannot be before 1900"))
		})

		It("should clone the builder correctly", func() {
			// Create a base course builder
			baseBuilder := builders.NewCourseBuilder().
				WithCode("CS101").
				WithName("Introduction to Computer Science").
				WithCredits(3.0).
				WithYear(2020)

			// Clone the builder and modify it
			clonedBuilder := baseBuilder.Clone().
				WithCode("CS102").
				WithName("Data Structures")

			// Build both courses
			baseCourse := baseBuilder.BuildPtr()
			clonedCourse := clonedBuilder.BuildPtr()

			// Verify the base course
			Expect(baseCourse.Code).To(Equal("CS101"))
			Expect(baseCourse.Name).To(Equal("Introduction to Computer Science"))

			// Verify the cloned course
			Expect(clonedCourse.Code).To(Equal("CS102"))
			Expect(clonedCourse.Name).To(Equal("Data Structures"))

			// Verify that the credits and year are the same
			Expect(clonedCourse.Credits).To(Equal(baseCourse.Credits))
			Expect(clonedCourse.Year).To(Equal(baseCourse.Year))
		})
	})
})
