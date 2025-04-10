// +build ignore

// +build ignore

import (
	"github.com/adil-faiyaz98/go-builder-kit/builders"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Education", func() {
	Context("Education Builder", func() {
		It("should build a valid Education with all fields", func() {
			// Create a Course builder
			courseBuilder := builders.NewCourseBuilder().
				WithCode("CS101").
				WithName("Introduction to Computer Science").
				WithCredits(3.0).
				WithGrade("A").
				WithYear(2015)

			// Create an Education builder
			educationBuilder := builders.NewEducationBuilder().
				WithDegree("Bachelor of Science").
				WithInstitution("Stanford University").
				WithStartDate("2016-09-01").
				WithEndDate("2020-06-15").
				WithGPA(3.8).
				WithHonor("Magna Cum Laude").
				WithMajor("Computer Science").
				WithMinor("Mathematics").
				WithCourse(courseBuilder).
				WithActivity("Chess Club").
				WithActivity("Programming Team")

			// Build the Education
			education, err := educationBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(education).NotTo(BeNil())
			Expect(education.Degree).To(Equal("Bachelor of Science"))
			Expect(education.Institution).To(Equal("Stanford University"))
			Expect(education.StartDate).To(Equal("2016-09-01"))
			Expect(education.EndDate).To(Equal("2020-06-15"))
			Expect(education.GPA).To(Equal(3.8))
			Expect(education.Honors).To(ContainElement("Magna Cum Laude"))
			Expect(education.Major).To(Equal("Computer Science"))
			Expect(education.Minor).To(Equal("Mathematics"))
			Expect(education.Activities).To(HaveLen(2))
			Expect(education.Activities).To(ContainElement("Chess Club"))
			Expect(education.Activities).To(ContainElement("Programming Team"))

			// Verify Courses
			Expect(education.Courses).To(HaveLen(1))
			Expect(education.Courses[0].Code).To(Equal("CS101"))
			Expect(education.Courses[0].Name).To(Equal("Introduction to Computer Science"))
		})

		It("should validate required fields", func() {
			// Create an Education builder with missing required fields
			educationBuilder := builders.NewEducationBuilder()
			// Missing Degree and Institution

			// Build and validate should fail
			_, err := educationBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Degree cannot be empty"))
			Expect(err.Error()).To(ContainSubstring("Institution cannot be empty"))
		})

		It("should validate date format", func() {
			// Create an Education builder with invalid date
			educationBuilder := builders.NewEducationBuilder().
				WithDegree("Bachelor of Science").
				WithInstitution("Stanford University").
				WithStartDate("09/01/2016") // Wrong format, should be YYYY-MM-DD

			// Build and validate should fail
			_, err := educationBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("StartDate must be in the format YYYY-MM-DD"))
		})

		It("should validate end date is after start date", func() {
			// Create an Education builder with end date before start date
			educationBuilder := builders.NewEducationBuilder().
				WithDegree("Bachelor of Science").
				WithInstitution("Stanford University").
				WithStartDate("2020-09-01").
				WithEndDate("2016-06-15") // Before start date

			// Build and validate should fail
			_, err := educationBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("EndDate cannot be before StartDate"))
		})

		It("should validate GPA range", func() {
			// Create an Education builder with invalid GPA
			educationBuilder := builders.NewEducationBuilder().
				WithDegree("Bachelor of Science").
				WithInstitution("Stanford University").
				WithGPA(5.0) // Above maximum

			// Build and validate should fail
			_, err := educationBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("GPA cannot be greater than 4.0"))
		})

		It("should validate nested structures", func() {
			// Create a Course builder with invalid data
			courseBuilder := builders.NewCourseBuilder().
				WithCode(""). // Missing required field
				WithName("Introduction to Computer Science")

			// Create an Education builder with the invalid course
			educationBuilder := builders.NewEducationBuilder().
				WithDegree("Bachelor of Science").
				WithInstitution("Stanford University").
				WithCourse(courseBuilder)

			// Build and validate should fail
			_, err := educationBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Code cannot be empty"))
		})

		It("should clone the builder correctly", func() {
			// Create a base education builder
			baseBuilder := builders.NewEducationBuilder().
				WithDegree("Bachelor of Science").
				WithInstitution("Stanford University").
				WithStartDate("2016-09-01").
				WithEndDate("2020-06-15").
				WithGPA(3.8)

			// Clone the builder and modify it
			clonedBuilder := baseBuilder.Clone().
				WithDegree("Master of Science").
				WithStartDate("2020-09-01").
				WithEndDate("2022-06-15")

			// Build both educations
			baseEducation := baseBuilder.BuildPtr()
			clonedEducation := clonedBuilder.BuildPtr()

			// Verify the base education
			Expect(baseEducation.Degree).To(Equal("Bachelor of Science"))
			Expect(baseEducation.StartDate).To(Equal("2016-09-01"))
			Expect(baseEducation.EndDate).To(Equal("2020-06-15"))

			// Verify the cloned education
			Expect(clonedEducation.Degree).To(Equal("Master of Science"))
			Expect(clonedEducation.StartDate).To(Equal("2020-09-01"))
			Expect(clonedEducation.EndDate).To(Equal("2022-06-15"))

			// Verify that the institution and GPA are the same
			Expect(clonedEducation.Institution).To(Equal(baseEducation.Institution))
			Expect(clonedEducation.GPA).To(Equal(baseEducation.GPA))
		})
	})
})
