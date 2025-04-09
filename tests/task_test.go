package tests

import (
	"github.com/adil-faiyaz98/go-builder-kit/builders"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Task", func() {
	Context("Task Builder", func() {
		It("should build a valid Task with all fields", func() {
			// Create a Subtask builder
			subtaskBuilder := builders.NewTaskBuilder().
				WithName("Research API").
				WithDescription("Research available APIs").
				WithStartDate("2023-01-15").
				WithEndDate("2023-01-20").
				WithStatus("completed").
				WithPriority("medium")

			// Create a Task builder
			taskBuilder := builders.NewTaskBuilder().
				WithName("Implement Feature X").
				WithDescription("Add new functionality").
				WithStartDate("2023-01-15").
				WithEndDate("2023-02-15").
				WithStatus("in-progress").
				WithPriority("high").
				WithSubtask(subtaskBuilder)

			// Build the Task
			task, err := taskBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(task).NotTo(BeNil())
			Expect(task.Name).To(Equal("Implement Feature X"))
			Expect(task.Description).To(Equal("Add new functionality"))
			Expect(task.StartDate).To(Equal("2023-01-15"))
			Expect(task.EndDate).To(Equal("2023-02-15"))
			Expect(task.Status).To(Equal("in-progress"))
			Expect(task.Priority).To(Equal("high"))

			// Verify Subtasks
			Expect(task.Subtasks).To(HaveLen(1))
			Expect(task.Subtasks[0].Name).To(Equal("Research API"))
			Expect(task.Subtasks[0].Status).To(Equal("completed"))
			Expect(task.Subtasks[0].Priority).To(Equal("medium"))
		})

		It("should validate required fields", func() {
			// Create a Task builder with missing required fields
			taskBuilder := builders.NewTaskBuilder()
			// Missing Name

			// Build and validate should fail
			_, err := taskBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Name cannot be empty"))
		})

		It("should validate date format", func() {
			// Create a Task builder with invalid date
			taskBuilder := builders.NewTaskBuilder().
				WithName("Implement Feature X").
				WithStartDate("01/15/2023") // Wrong format, should be YYYY-MM-DD

			// Build and validate should fail
			_, err := taskBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("StartDate must be in the format YYYY-MM-DD"))
		})

		It("should validate end date is after start date", func() {
			// Create a Task builder with end date before start date
			taskBuilder := builders.NewTaskBuilder().
				WithName("Implement Feature X").
				WithStartDate("2023-02-15").
				WithEndDate("2023-01-15") // Before start date

			// Build and validate should fail
			_, err := taskBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("EndDate cannot be before StartDate"))
		})

		It("should validate status values", func() {
			// Create a Task builder with invalid status
			taskBuilder := builders.NewTaskBuilder().
				WithName("Implement Feature X").
				WithStatus("invalid") // Not a valid status

			// Build and validate should fail
			_, err := taskBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Status must be one of"))
		})

		It("should validate priority values", func() {
			// Create a Task builder with invalid priority
			taskBuilder := builders.NewTaskBuilder().
				WithName("Implement Feature X").
				WithPriority("invalid") // Not a valid priority

			// Build and validate should fail
			_, err := taskBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Priority must be one of"))
		})

		It("should validate nested structures", func() {
			// Create a Subtask builder with invalid data
			subtaskBuilder := builders.NewTaskBuilder().
				WithName(""). // Missing required field
				WithDescription("Research available APIs")

			// Create a Task builder with the invalid subtask
			taskBuilder := builders.NewTaskBuilder().
				WithName("Implement Feature X").
				WithSubtask(subtaskBuilder)

			// Build and validate should fail
			_, err := taskBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Name cannot be empty"))
		})

		It("should clone the builder correctly", func() {
			// Create a base task builder
			baseBuilder := builders.NewTaskBuilder().
				WithName("Implement Feature X").
				WithDescription("Add new functionality").
				WithStartDate("2023-01-15").
				WithEndDate("2023-02-15").
				WithStatus("in-progress").
				WithPriority("high")

			// Clone the builder and modify it
			clonedBuilder := baseBuilder.Clone().
				WithName("Implement Feature Y").
				WithEndDate("2023-03-15")

			// Build both tasks
			baseTask := baseBuilder.BuildPtr()
			clonedTask := clonedBuilder.BuildPtr()

			// Verify the base task
			Expect(baseTask.Name).To(Equal("Implement Feature X"))
			Expect(baseTask.EndDate).To(Equal("2023-02-15"))

			// Verify the cloned task
			Expect(clonedTask.Name).To(Equal("Implement Feature Y"))
			Expect(clonedTask.EndDate).To(Equal("2023-03-15"))

			// Verify that the description, start date, status, and priority are the same
			Expect(clonedTask.Description).To(Equal(baseTask.Description))
			Expect(clonedTask.StartDate).To(Equal(baseTask.StartDate))
			Expect(clonedTask.Status).To(Equal(baseTask.Status))
			Expect(clonedTask.Priority).To(Equal(baseTask.Priority))
		})
	})
})
