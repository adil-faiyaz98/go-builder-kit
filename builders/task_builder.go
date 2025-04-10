package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

// TaskBuilder builds a Task model
type TaskBuilder struct {
	task           *models.Task
	validationFuncs []func(*models.Task) error
}

// NewTaskBuilder creates a new TaskBuilder
func NewTaskBuilder() *TaskBuilder {
	return &TaskBuilder{
		task: &models.Task{
			ID:          "",
			Title:       "",
			Description: "",
			Status:      "",
			Priority:    "",
			DueDate:     "",
			AssignedTo:  "",
			CreatedBy:   "",
			CreatedAt:   "",
			UpdatedAt:   "",
			Tags:        []string{},
			Comments:    []string{},
			Attachments: []string{},
			SubTasks:    []*models.Task{},
			ParentTask:  nil,
			EstimatedHours: 0.0,
			ActualHours: 0.0,
			Progress:    0.0,
		},
		validationFuncs: []func(*models.Task) error{},
	}
}

// WithID sets the ID
func (b *TaskBuilder) WithID(id string) *TaskBuilder {
	b.task.ID = id
	return b
}

// WithTitle sets the Title
func (b *TaskBuilder) WithTitle(title string) *TaskBuilder {
	b.task.Title = title
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

// WithAssignedTo sets the AssignedTo
func (b *TaskBuilder) WithAssignedTo(assignedTo string) *TaskBuilder {
	b.task.AssignedTo = assignedTo
	return b
}

// WithCreatedBy sets the CreatedBy
func (b *TaskBuilder) WithCreatedBy(createdBy string) *TaskBuilder {
	b.task.CreatedBy = createdBy
	return b
}

// WithCreatedAt sets the CreatedAt
func (b *TaskBuilder) WithCreatedAt(createdAt string) *TaskBuilder {
	b.task.CreatedAt = createdAt
	return b
}

// WithUpdatedAt sets the UpdatedAt
func (b *TaskBuilder) WithUpdatedAt(updatedAt string) *TaskBuilder {
	b.task.UpdatedAt = updatedAt
	return b
}

// WithTags sets the Tags
func (b *TaskBuilder) WithTags(tags []string) *TaskBuilder {
	b.task.Tags = tags
	return b
}

// AddTag adds a tag to the Tags slice
func (b *TaskBuilder) AddTag(tag string) *TaskBuilder {
	b.task.Tags = append(b.task.Tags, tag)
	return b
}

// WithComments sets the Comments
func (b *TaskBuilder) WithComments(comments []string) *TaskBuilder {
	b.task.Comments = comments
	return b
}

// AddComment adds a comment to the Comments slice
func (b *TaskBuilder) AddComment(comment string) *TaskBuilder {
	b.task.Comments = append(b.task.Comments, comment)
	return b
}

// WithAttachments sets the Attachments
func (b *TaskBuilder) WithAttachments(attachments []string) *TaskBuilder {
	b.task.Attachments = attachments
	return b
}

// AddAttachment adds an attachment to the Attachments slice
func (b *TaskBuilder) AddAttachment(attachment string) *TaskBuilder {
	b.task.Attachments = append(b.task.Attachments, attachment)
	return b
}

// WithSubTasks sets the SubTasks
func (b *TaskBuilder) WithSubTasks(subTasks []*models.Task) *TaskBuilder {
	b.task.SubTasks = subTasks
	return b
}

// AddSubTask adds a sub-task to the SubTasks slice
func (b *TaskBuilder) AddSubTask(subTask *TaskBuilder) *TaskBuilder {
	b.task.SubTasks = append(b.task.SubTasks, subTask.BuildPtr())
	return b
}

// WithParentTask sets the ParentTask
func (b *TaskBuilder) WithParentTask(parentTask *TaskBuilder) *TaskBuilder {
	b.task.ParentTask = parentTask.BuildPtr()
	return b
}

// WithEstimatedHours sets the EstimatedHours
func (b *TaskBuilder) WithEstimatedHours(estimatedHours float64) *TaskBuilder {
	b.task.EstimatedHours = estimatedHours
	return b
}

// WithActualHours sets the ActualHours
func (b *TaskBuilder) WithActualHours(actualHours float64) *TaskBuilder {
	b.task.ActualHours = actualHours
	return b
}

// WithProgress sets the Progress
func (b *TaskBuilder) WithProgress(progress float64) *TaskBuilder {
	b.task.Progress = progress
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

	// Run model's Validate method
	if err := task.Validate(); err != nil {
		return task, err
	}

	return task, nil
}

// MustBuild builds the Task and panics if validation fails
func (b *TaskBuilder) MustBuild() *models.Task {
	task, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return task
}

// Clone creates a deep copy of the builder
func (b *TaskBuilder) Clone() *TaskBuilder {
	clonedTask := *b.task
	return &TaskBuilder{
		task:           &clonedTask,
		validationFuncs: append([]func(*models.Task) error{}, b.validationFuncs...),
	}
}
