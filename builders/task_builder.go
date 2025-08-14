package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/v2/models"
	"github.com/adil-faiyaz98/go-builder-kit/v2/pkg/builder"
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
			Assignee:    interface{}(0),
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
	if b == nil {
		return b
	}
	b.task.Name = builder.SanitizeString(name)
	return b
}

// WithDescription sets the Description
func (b *TaskBuilder) WithDescription(description string) *TaskBuilder {
	if b == nil {
		return b
	}
	b.task.Description = builder.SanitizeString(description)
	return b
}

// WithStartDate sets the StartDate
func (b *TaskBuilder) WithStartDate(startDate string) *TaskBuilder {
	if b == nil {
		return b
	}
	b.task.StartDate = builder.SanitizeString(startDate)
	return b
}

// WithEndDate sets the EndDate
func (b *TaskBuilder) WithEndDate(endDate string) *TaskBuilder {
	if b == nil {
		return b
	}
	b.task.EndDate = builder.SanitizeString(endDate)
	return b
}

// WithStatus sets the Status
func (b *TaskBuilder) WithStatus(status string) *TaskBuilder {
	if b == nil {
		return b
	}
	b.task.Status = builder.SanitizeString(status)
	return b
}

// WithPriority sets the Priority
func (b *TaskBuilder) WithPriority(priority string) *TaskBuilder {
	if b == nil {
		return b
	}
	b.task.Priority = builder.SanitizeString(priority)
	return b
}

// WithAssignee sets the Assignee
func (b *TaskBuilder) WithAssignee(assignee interface{}) *TaskBuilder {
	if b == nil {
		return b
	}
	b.task.Assignee = assignee
	return b
}

// WithSubtasks sets the Subtasks
func (b *TaskBuilder) WithSubtasks(subtasks []*TaskBuilder) *TaskBuilder {
	if b == nil {
		return b
	}
	// Ensure the slice is initialized
	if b.task.Subtasks == nil {
		b.task.Subtasks = []*models.Task{}
	}
	// Handle nested slice elements
	// Initialize the slice
	b.task.Subtasks = make([]*models.Task, 0, len(subtasks))
	// Convert each builder to its model
	for _, builder := range subtasks {
		builtValue := builder.Build().(*models.Task)
		b.task.Subtasks = append(b.task.Subtasks, builtValue)
	}
	return b
}

// AddSubtask adds a single item to the Subtasks slice
func (b *TaskBuilder) AddSubtask(subtask *TaskBuilder) *TaskBuilder {
	if b == nil {
		return b
	}
	// Ensure the slice is initialized with capacity
	if b.task.Subtasks == nil {
		b.task.Subtasks = make([]*models.Task, 0, 4) // Pre-allocate capacity
	}
	// Handle nested slice element
	if subtask != nil {
		builtValue := subtask.Build().(*models.Task)
		b.task.Subtasks = append(b.task.Subtasks, builtValue)
	}
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
	if b == nil || b.task == nil {
		return nil, fmt.Errorf("builder is not properly initialized")
	}

	task := b.task

	// Run custom validation functions
	for i, validationFunc := range b.validationFuncs {
		if validationFunc == nil {
			continue // Skip nil validators
		}
		if err := validationFunc(task); err != nil {
			return nil, fmt.Errorf("custom validation failed at index %d: %w", i, err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := any(task).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return task, fmt.Errorf("model validation failed: %w", err)
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
	if b == nil || b.task == nil {
		return NewTaskBuilder()
	}

	// Deep copy the struct
	clonedTask := *b.task

	// Create new builder with cloned data
	clonedBuilder := &TaskBuilder{
		task:            &clonedTask,
		validationFuncs: make([]func(*models.Task) error, 0, len(b.validationFuncs)),
	}

	// Copy validation functions safely
	if len(b.validationFuncs) > 0 {
		clonedBuilder.validationFuncs = append(clonedBuilder.validationFuncs, b.validationFuncs...)
	}

	return clonedBuilder
}
