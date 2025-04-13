package education_test

import (
	"github.com/adil-faiyaz98/go-builder-kit/builders"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Education Builder", func() {

	Context("Basic Education Creation", func() {
		It("should create an education with basic fields", func() {
			// Create an education with basic fields
			educationBuilder := builders.NewEducationBuilder().
				WithDegree("Bachelor of Science").
				WithMajor("Computer Science").
				WithInstitution("Stanford University").
				WithStartDate("2016-09-01").
				WithEndDate("2020-06-01")

			// Build the education
			education := educationBuilder.BuildPtr()

			// Verify the education fields
			Expect(education.Degree).To(Equal("Bachelor of Science"))
			Expect(education.Major).To(Equal("Computer Science"))
			Expect(education.Institution).To(Equal("Stanford University"))
			Expect(education.StartDate).To(Equal("2016-09-01"))
			Expect(education.EndDate).To(Equal("2020-06-01"))
		})
	})

	Context("Different Degree Types", func() {
		It("should create education entries with different degree types", func() {
			// Create a Bachelor's degree
			bsBuilder := builders.NewEducationBuilder().
				WithDegree("Bachelor of Science").
				WithMajor("Computer Science").
				WithInstitution("Stanford University").
				WithStartDate("2014-09-01").
				WithEndDate("2018-06-01")

			// Create a Master's degree
			msBuilder := builders.NewEducationBuilder().
				WithDegree("Master of Science").
				WithMajor("Artificial Intelligence").
				WithInstitution("MIT").
				WithStartDate("2018-09-01").
				WithEndDate("2020-06-01")

			// Create a PhD
			phdBuilder := builders.NewEducationBuilder().
				WithDegree("Doctor of Philosophy").
				WithMajor("Computer Science").
				WithInstitution("UC Berkeley").
				WithStartDate("2020-09-01").
				WithEndDate("2023-06-01")

			// Build the education entries
			bs := bsBuilder.BuildPtr()
			ms := msBuilder.BuildPtr()
			phd := phdBuilder.BuildPtr()

			// Verify the Bachelor's degree
			Expect(bs.Degree).To(Equal("Bachelor of Science"))
			Expect(bs.Major).To(Equal("Computer Science"))
			Expect(bs.Institution).To(Equal("Stanford University"))
			Expect(bs.StartDate).To(Equal("2014-09-01"))
			Expect(bs.EndDate).To(Equal("2018-06-01"))

			// Verify the Master's degree
			Expect(ms.Degree).To(Equal("Master of Science"))
			Expect(ms.Major).To(Equal("Artificial Intelligence"))
			Expect(ms.Institution).To(Equal("MIT"))
			Expect(ms.StartDate).To(Equal("2018-09-01"))
			Expect(ms.EndDate).To(Equal("2020-06-01"))

			// Verify the PhD
			Expect(phd.Degree).To(Equal("Doctor of Philosophy"))
			Expect(phd.Major).To(Equal("Computer Science"))
			Expect(phd.Institution).To(Equal("UC Berkeley"))
			Expect(phd.StartDate).To(Equal("2020-09-01"))
			Expect(phd.EndDate).To(Equal("2023-06-01"))
		})
	})

	Context("Builder Cloning", func() {
		It("should clone an education builder correctly", func() {
			// Create a base education builder
			baseBuilder := builders.NewEducationBuilder().
				WithDegree("Bachelor of Science").
				WithMajor("Computer Science").
				WithInstitution("Stanford University").
				WithStartDate("2016-09-01").
				WithEndDate("2020-06-01")

			// Clone the builder and modify it
			clonedBuilder := baseBuilder.Clone().
				WithDegree("Master of Science").
				WithStartDate("2020-09-01").
				WithEndDate("2022-06-01")

			// Build both education entries
			baseEducation := baseBuilder.BuildPtr()
			clonedEducation := clonedBuilder.BuildPtr()

			// Verify the base education wasn't affected by changes to the clone
			Expect(baseEducation.Degree).To(Equal("Bachelor of Science"))
			Expect(baseEducation.StartDate).To(Equal("2016-09-01"))
			Expect(baseEducation.EndDate).To(Equal("2020-06-01"))

			// Verify the cloned education has the new values
			Expect(clonedEducation.Degree).To(Equal("Master of Science"))
			Expect(clonedEducation.Major).To(Equal("Computer Science"))          // Unchanged
			Expect(clonedEducation.Institution).To(Equal("Stanford University")) // Unchanged
			Expect(clonedEducation.StartDate).To(Equal("2020-09-01"))
			Expect(clonedEducation.EndDate).To(Equal("2022-06-01"))
		})
	})
})
