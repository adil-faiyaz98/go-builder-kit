package project_test

import (
	"github.com/adil-faiyaz98/go-builder-kit/builders"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Project Builder", func() {

	Context("Basic Project Creation", func() {
		It("should create a project with basic fields", func() {
			// Create a project with basic fields
			projectBuilder := builders.NewProjectBuilder().
				WithName("Test Project").
				WithDescription("A test project").
				WithStartDate("2023-01-01").
				WithEndDate("2023-12-31").
				WithStatus("Active").
				WithBudget(100000.0)

			// Build the project
			project := projectBuilder.BuildPtr()

			// Verify the project fields
			Expect(project.Name).To(Equal("Test Project"))
			Expect(project.Description).To(Equal("A test project"))
			Expect(project.StartDate).To(Equal("2023-01-01"))
			Expect(project.EndDate).To(Equal("2023-12-31"))
			Expect(project.Status).To(Equal("Active"))
			Expect(project.Budget).To(Equal(100000.0))
		})
	})

	Context("Project with Tasks", func() {
		It("should create a project with tasks", func() {
			// Create tasks
			task1Builder := builders.NewTaskBuilder().
				WithName("Task 1").
				WithDescription("First task").
				WithStatus("In Progress").
				WithPriority("High")

			task2Builder := builders.NewTaskBuilder().
				WithName("Task 2").
				WithDescription("Second task").
				WithStatus("Not Started").
				WithPriority("Medium")

			// Create a project with tasks
			projectBuilder := builders.NewProjectBuilder().
				WithName("Test Project").
				WithDescription("A test project")

			// Add tasks to the project
			projectBuilder.AddTask(task1Builder)
			projectBuilder.AddTask(task2Builder)

			// Build the project
			project := projectBuilder.BuildPtr()

			// Verify the project fields
			Expect(project.Name).To(Equal("Test Project"))
			Expect(project.Description).To(Equal("A test project"))

			// Verify the tasks
			Expect(project.Tasks).To(HaveLen(2))

			// Verify first task
			Expect(project.Tasks[0].Name).To(Equal("Task 1"))
			Expect(project.Tasks[0].Description).To(Equal("First task"))
			Expect(project.Tasks[0].Status).To(Equal("In Progress"))
			Expect(project.Tasks[0].Priority).To(Equal("High"))

			// Verify second task
			Expect(project.Tasks[1].Name).To(Equal("Task 2"))
			Expect(project.Tasks[1].Description).To(Equal("Second task"))
			Expect(project.Tasks[1].Status).To(Equal("Not Started"))
			Expect(project.Tasks[1].Priority).To(Equal("Medium"))
		})
	})

	Context("Project Status", func() {
		It("should create projects with different statuses", func() {
			// Create projects with different statuses
			activeProjectBuilder := builders.NewProjectBuilder().
				WithName("Active Project").
				WithStatus("Active")

			completedProjectBuilder := builders.NewProjectBuilder().
				WithName("Completed Project").
				WithStatus("Completed")

			onHoldProjectBuilder := builders.NewProjectBuilder().
				WithName("On Hold Project").
				WithStatus("On Hold")

			// Build the projects
			activeProject := activeProjectBuilder.BuildPtr()
			completedProject := completedProjectBuilder.BuildPtr()
			onHoldProject := onHoldProjectBuilder.BuildPtr()

			// Verify the statuses
			Expect(activeProject.Status).To(Equal("Active"))
			Expect(completedProject.Status).To(Equal("Completed"))
			Expect(onHoldProject.Status).To(Equal("On Hold"))
		})
	})

	Context("Builder Cloning", func() {
		It("should clone a project builder correctly", func() {
			// Create a base project builder
			baseBuilder := builders.NewProjectBuilder().
				WithName("Base Project").
				WithDescription("Base description").
				WithStatus("Planning").
				WithBudget(50000.0)

			// Clone the builder and modify it
			clonedBuilder := baseBuilder.Clone().
				WithName("Cloned Project").
				WithStatus("Active").
				WithBudget(75000.0)

			// Build both projects
			baseProject := baseBuilder.BuildPtr()
			clonedProject := clonedBuilder.BuildPtr()

			// Verify the base project wasn't affected by changes to the clone
			Expect(baseProject.Name).To(Equal("Base Project"))
			Expect(baseProject.Description).To(Equal("Base description"))
			Expect(baseProject.Status).To(Equal("Planning"))
			Expect(baseProject.Budget).To(Equal(50000.0))

			// Verify the cloned project has the new values
			Expect(clonedProject.Name).To(Equal("Cloned Project"))
			Expect(clonedProject.Description).To(Equal("Base description")) // Unchanged
			Expect(clonedProject.Status).To(Equal("Active"))
			Expect(clonedProject.Budget).To(Equal(75000.0))
		})
	})
})
