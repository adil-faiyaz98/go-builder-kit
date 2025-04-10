package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/models"
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
			Name:        "",
			Description: "",
			StartDate:   "",
			EndDate:     "",
			Status:      "",
			Priority:    "",
			Assignee:    nil,
			Subtasks:    []*models.Task{},
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

// WithStartDate sets the StartDate
func (b *TaskBuilder) WithStartDate(startDate string) *TaskBuilder {
	b.task.StartDate = startDate
	return b
}

// WithEndDate sets the EndDate
func (b *TaskBuilder) WithEndDate(endDate string) *TaskBuilder {
	b.task.EndDate = endDate
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

// WithAssignee sets the Assignee
func (b *TaskBuilder) WithAssignee(assignee interface{}) *TaskBuilder {
	b.task.Assignee = assignee
	return b
}

// WithSubtasks sets the Subtasks
func (b *TaskBuilder) WithSubtasks(subtasks []*models.Task) *TaskBuilder {
	b.task.Subtasks = subtasks
	return b
}

// AddSubtask adds a sub-task to the Subtasks slice
func (b *TaskBuilder) AddSubtask(subtask *TaskBuilder) *TaskBuilder {
	b.task.Subtasks = append(b.task.Subtasks, subtask.BuildPtr())
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
		task:           &clonedTask,
		validationFuncs: append([]func(*models.Task) error{}, b.validationFuncs...),
	}
}
