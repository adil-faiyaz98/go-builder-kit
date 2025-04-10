package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/test/models"
)

// TaskBuilder builds a Task model
type TaskBuilder struct {
	task *models.Task
	// Custom validation functions
	validationFuncs []func(*models.Task) error
}

// NewTaskBuilder creates a new TaskBuilder
func NewTaskBuilder() *TaskBuilder {
	return &TaskBuilder{
		task: &models.Task{
			Name: "",
			Description: "",
			Status: "",
			Priority: "",
			DueDate: "",
			Assignee: nil,
			SubTasks: []*models.Task{},
			Project: nil,
		},
		validationFuncs: []func(*models.Task) error{},
	}
}

// NewTaskBuilderWithDefaults creates a new TaskBuilder with sensible defaults
func NewTaskBuilderWithDefaults() *TaskBuilder {
	builder := NewTaskBuilder()
	// Add default values here if needed
	return builder
}
// WithName sets the Name
func (b *TaskBuilder) WithName(name string) *TaskBuilder {
	b.task.Name = name
	return b
}

// WithDescription sets the Description
func (b *TaskBuilder) WithDescription(description string) *TaskBuilder {
	b.task.Description = description
	return b
}

// WithStatus sets the Status
func (b *TaskBuilder) WithStatus(status string) *TaskBuilder {
	b.task.Status = status
	return b
}

// WithPriority sets the Priority
func (b *TaskBuilder) WithPriority(priority string) *TaskBuilder {
	b.task.Priority = priority
	return b
}

// WithDueDate sets the DueDate
func (b *TaskBuilder) WithDueDate(dueDate string) *TaskBuilder {
	b.task.DueDate = dueDate
	return b
}

// WithAssignee sets the Assignee
func (b *TaskBuilder) WithAssignee(assignee *EmployeeBuilder) *TaskBuilder {
	// Handle nested pointer
	b.task.Assignee = assignee.BuildPtr()
	return b
}

// WithSubTasks sets the SubTasks
func (b *TaskBuilder) WithSubTasks(subTasks []*TaskBuilder) *TaskBuilder {
	// Ensure the slice is initialized
	if b.task.SubTasks == nil {
		b.task.SubTasks = []*models.Task{}
	}
	// Handle nested slice elements
	// Initialize the slice
	b.task.SubTasks = make([]*models.Task, 0, len(subTasks))
	// Convert each builder to its model
	for _, builder := range subTasks {
		builtValue := builder.Build().(*models.Task)
		b.task.SubTasks = append(b.task.SubTasks, builtValue)
	}
	return b
}

// WithProject sets the Project
func (b *TaskBuilder) WithProject(project *ProjectBuilder) *TaskBuilder {
	// Handle nested pointer
	b.task.Project = project.BuildPtr()
	return b
}

// AddSubTask adds a single item to the SubTasks slice
func (b *TaskBuilder) AddSubTask(subTask *TaskBuilder) *TaskBuilder {
	// Ensure the slice is initialized
	if b.task.SubTasks == nil {
		b.task.SubTasks = []*models.Task{}
	}
	// Handle nested slice element
	builtValue := subTask.Build().(*models.Task)
	b.task.SubTasks = append(b.task.SubTasks, builtValue)
	return b
}

// WithValidation adds a custom validation function
func (b *TaskBuilder) WithValidation(validationFunc func(*models.Task) error) *TaskBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Task
func (b *TaskBuilder) Build() interface{} {
	return b.task
}

// BuildPtr builds the Task and returns a pointer
func (b *TaskBuilder) BuildPtr() *models.Task {
	return b.task
}

// BuildAndValidate builds the Task and validates it
func (b *TaskBuilder) BuildAndValidate() (*models.Task, error) {
	task := b.task

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(task); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(task).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return task, err
		}
	}

	return task, nil
}

// MustBuild builds the Task and panics if validation fails
func (b *TaskBuilder) MustBuild() *models.Task {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *TaskBuilder) Clone() *TaskBuilder {
	clonedTask := *b.task
	return &TaskBuilder{
		task: &clonedTask,
		validationFuncs: append([]func(*models.Task) error{}, b.validationFuncs...),
	}
}
