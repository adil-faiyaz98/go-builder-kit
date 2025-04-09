package tests

import (
	"github.com/adil-faiyaz98/go-builder-kit/builders"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Project", func() {
	Context("Project Builder", func() {
		It("should build a valid Project with all fields", func() {
			// Create a Task builder
			taskBuilder := builders.NewTaskBuilder().
				WithName("Implement Feature X").
				WithDescription("Add new functionality").
				WithStartDate("2023-01-15").
				WithEndDate("2023-02-15").
				WithStatus("in-progress").
				WithPriority("high")

			// Create a Project builder
			projectBuilder := builders.NewProjectBuilder().
				WithName("Product Redesign").
				WithDescription("Redesign the product UI/UX").
				WithStartDate("2023-01-01").
				WithEndDate("2023-06-30").
				WithStatus("in-progress").
				WithBudget(200000.0).
				WithTask(taskBuilder)

			// Build the Project
			project, err := projectBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(project).NotTo(BeNil())
			Expect(project.Name).To(Equal("Product Redesign"))
			Expect(project.Description).To(Equal("Redesign the product UI/UX"))
			Expect(project.StartDate).To(Equal("2023-01-01"))
			Expect(project.EndDate).To(Equal("2023-06-30"))
			Expect(project.Status).To(Equal("in-progress"))
			Expect(project.Budget).To(Equal(200000.0))

			// Verify Tasks
			Expect(project.Tasks).To(HaveLen(1))
			Expect(project.Tasks[0].Name).To(Equal("Implement Feature X"))
			Expect(project.Tasks[0].Status).To(Equal("in-progress"))
			Expect(project.Tasks[0].Priority).To(Equal("high"))
		})

		It("should validate required fields", func() {
			// Create a Project builder with missing required fields
			projectBuilder := builders.NewProjectBuilder()
			// Missing Name

			// Build and validate should fail
			_, err := projectBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Name cannot be empty"))
		})

		It("should validate date format", func() {
			// Create a Project builder with invalid date
			projectBuilder := builders.NewProjectBuilder().
				WithName("Product Redesign").
				WithStartDate("01/01/2023") // Wrong format, should be YYYY-MM-DD

			// Build and validate should fail
			_, err := projectBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("StartDate must be in the format YYYY-MM-DD"))
		})

		It("should validate end date is after start date", func() {
			// Create a Project builder with end date before start date
			projectBuilder := builders.NewProjectBuilder().
				WithName("Product Redesign").
				WithStartDate("2023-06-30").
				WithEndDate("2023-01-01") // Before start date

			// Build and validate should fail
			_, err := projectBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("EndDate cannot be before StartDate"))
		})

		It("should validate budget is positive", func() {
			// Create a Project builder with negative budget
			projectBuilder := builders.NewProjectBuilder().
				WithName("Product Redesign").
				WithBudget(-1000.0) // Negative budget

			// Build and validate should fail
			_, err := projectBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Budget cannot be negative"))
		})

		It("should validate status values", func() {
			// Create a Project builder with invalid status
			projectBuilder := builders.NewProjectBuilder().
				WithName("Product Redesign").
				WithStatus("invalid") // Not a valid status

			// Build and validate should fail
			_, err := projectBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Status must be one of"))
		})

		It("should validate nested structures", func() {
			// Create a Task builder with invalid data
			taskBuilder := builders.NewTaskBuilder().
				WithName(""). // Missing required field
				WithDescription("Add new functionality")

			// Create a Project builder with the invalid task
			projectBuilder := builders.NewProjectBuilder().
				WithName("Product Redesign").
				WithTask(taskBuilder)

			// Build and validate should fail
			_, err := projectBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Name cannot be empty"))
		})

		It("should clone the builder correctly", func() {
			// Create a base project builder
			baseBuilder := builders.NewProjectBuilder().
				WithName("Product Redesign").
				WithDescription("Redesign the product UI/UX").
				WithStartDate("2023-01-01").
				WithEndDate("2023-06-30").
				WithBudget(200000.0)

			// Clone the builder and modify it
			clonedBuilder := baseBuilder.Clone().
				WithName("Website Redesign").
				WithBudget(150000.0)

			// Build both projects
			baseProject := baseBuilder.BuildPtr()
			clonedProject := clonedBuilder.BuildPtr()

			// Verify the base project
			Expect(baseProject.Name).To(Equal("Product Redesign"))
			Expect(baseProject.Budget).To(Equal(200000.0))

			// Verify the cloned project
			Expect(clonedProject.Name).To(Equal("Website Redesign"))
			Expect(clonedProject.Budget).To(Equal(150000.0))

			// Verify that the description and dates are the same
			Expect(clonedProject.Description).To(Equal(baseProject.Description))
			Expect(clonedProject.StartDate).To(Equal(baseProject.StartDate))
			Expect(clonedProject.EndDate).To(Equal(baseProject.EndDate))
		})
	})
})
