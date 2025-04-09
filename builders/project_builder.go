package builders

import (
	"fmt"

	"github.com/adil-faiyaz98/go-builder-kit/models"
)

// ProjectBuilder builds a Project model
type ProjectBuilder struct {
	project         *models.Project
	validationFuncs []func(*models.Project) error
}

// NewProjectBuilder creates a new ProjectBuilder
func NewProjectBuilder() *ProjectBuilder {
	return &ProjectBuilder{
		project: &models.Project{
			Name:        "",
			Description: "",
			StartDate:   "",
			EndDate:     "",
			Status:      "",
			Budget:      0.0,
			Manager:     nil,
			Team:        []any{},
			Members:     []any{},
			Tasks:       []*models.Task{},
		},
		validationFuncs: []func(*models.Project) error{},
	}
}

// WithName sets the Name
func (b *ProjectBuilder) WithName(name string) *ProjectBuilder {
	b.project.Name = name
	return b
}

// WithDescription sets the Description
func (b *ProjectBuilder) WithDescription(description string) *ProjectBuilder {
	b.project.Description = description
	return b
}

// WithStartDate sets the StartDate
func (b *ProjectBuilder) WithStartDate(startDate string) *ProjectBuilder {
	b.project.StartDate = startDate
	return b
}

// WithEndDate sets the EndDate
func (b *ProjectBuilder) WithEndDate(endDate string) *ProjectBuilder {
	b.project.EndDate = endDate
	return b
}

// WithStatus sets the Status
func (b *ProjectBuilder) WithStatus(status string) *ProjectBuilder {
	b.project.Status = status
	return b
}

// WithBudget sets the Budget
func (b *ProjectBuilder) WithBudget(budget float64) *ProjectBuilder {
	b.project.Budget = budget
	return b
}

// WithManager sets the Manager
func (b *ProjectBuilder) WithManager(manager any) *ProjectBuilder {
	b.project.Manager = manager
	return b
}

// WithTeamMember adds a team member to the Team slice
func (b *ProjectBuilder) WithTeamMember(teamMember any) *ProjectBuilder {
	b.project.Team = append(b.project.Team, teamMember)
	return b
}

// WithMember adds a member to the Members slice
func (b *ProjectBuilder) WithMember(member any) *ProjectBuilder {
	b.project.Members = append(b.project.Members, member)
	return b
}

// WithTask adds a task to the Tasks slice
func (b *ProjectBuilder) WithTask(task *TaskBuilder) *ProjectBuilder {
	builtValue := task.Build().(*models.Task)
	b.project.Tasks = append(b.project.Tasks, builtValue)
	return b
}

// WithValidation adds a custom validation function
func (b *ProjectBuilder) WithValidation(validationFunc func(*models.Project) error) *ProjectBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Project
func (b *ProjectBuilder) Build() any {
	return b.project
}

// BuildPtr builds the Project and returns a pointer
func (b *ProjectBuilder) BuildPtr() *models.Project {
	return b.project
}

// BuildAndValidate builds the Project and validates it
func (b *ProjectBuilder) BuildAndValidate() (*models.Project, error) {
	project := b.project

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(project); err != nil {
			return project, err
		}
	}

	// Run model's Validate method
	if err := project.Validate(); err != nil {
		return project, err
	}

	return project, nil
}

// MustBuild builds the Project and panics if validation fails
func (b *ProjectBuilder) MustBuild() *models.Project {
	project, err := b.BuildAndValidate()
	if err != nil {
		panic(fmt.Sprintf("Project validation failed: %s", err.Error()))
	}
	return project
}

// Clone creates a deep copy of the ProjectBuilder
func (b *ProjectBuilder) Clone() *ProjectBuilder {
	clonedProject := *b.project
	return &ProjectBuilder{
		project:         &clonedProject,
		validationFuncs: append([]func(*models.Project) error{}, b.validationFuncs...),
	}
}
