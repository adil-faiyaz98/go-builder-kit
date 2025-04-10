package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/test/models"
)

// ProjectBuilder builds a Project model
type ProjectBuilder struct {
	project *models.Project
	// Custom validation functions
	validationFuncs []func(*models.Project) error
}

// NewProjectBuilder creates a new ProjectBuilder
func NewProjectBuilder() *ProjectBuilder {
	return &ProjectBuilder{
		project: &models.Project{
			Name: "",
			Description: "",
			StartDate: "",
			EndDate: "",
			Status: "",
			Budget: 0.0,
			Manager: nil,
			Team: []*models.Employee{},
			Tasks: []*models.Task{},
			SubProjects: []*models.Project{},
		},
		validationFuncs: []func(*models.Project) error{},
	}
}

// NewProjectBuilderWithDefaults creates a new ProjectBuilder with sensible defaults
func NewProjectBuilderWithDefaults() *ProjectBuilder {
	builder := NewProjectBuilder()
	// Add default values here if needed
	return builder
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
func (b *ProjectBuilder) WithManager(manager *EmployeeBuilder) *ProjectBuilder {
	// Handle nested pointer
	b.project.Manager = manager.BuildPtr()
	return b
}

// WithTeam sets the Team
func (b *ProjectBuilder) WithTeam(team []*EmployeeBuilder) *ProjectBuilder {
	// Ensure the slice is initialized
	if b.project.Team == nil {
		b.project.Team = []*models.Employee{}
	}
	// Handle nested slice elements
	// Initialize the slice
	b.project.Team = make([]*models.Employee, 0, len(team))
	// Convert each builder to its model
	for _, builder := range team {
		builtValue := builder.Build().(*models.Employee)
		b.project.Team = append(b.project.Team, builtValue)
	}
	return b
}

// WithTasks sets the Tasks
func (b *ProjectBuilder) WithTasks(tasks []*TaskBuilder) *ProjectBuilder {
	// Ensure the slice is initialized
	if b.project.Tasks == nil {
		b.project.Tasks = []*models.Task{}
	}
	// Handle nested slice elements
	// Initialize the slice
	b.project.Tasks = make([]*models.Task, 0, len(tasks))
	// Convert each builder to its model
	for _, builder := range tasks {
		builtValue := builder.Build().(*models.Task)
		b.project.Tasks = append(b.project.Tasks, builtValue)
	}
	return b
}

// WithSubProjects sets the SubProjects
func (b *ProjectBuilder) WithSubProjects(subProjects []*ProjectBuilder) *ProjectBuilder {
	// Ensure the slice is initialized
	if b.project.SubProjects == nil {
		b.project.SubProjects = []*models.Project{}
	}
	// Handle nested slice elements
	// Initialize the slice
	b.project.SubProjects = make([]*models.Project, 0, len(subProjects))
	// Convert each builder to its model
	for _, builder := range subProjects {
		builtValue := builder.Build().(*models.Project)
		b.project.SubProjects = append(b.project.SubProjects, builtValue)
	}
	return b
}

// AddTeam adds a single item to the Team slice
func (b *ProjectBuilder) AddTeam(team *EmployeeBuilder) *ProjectBuilder {
	// Ensure the slice is initialized
	if b.project.Team == nil {
		b.project.Team = []*models.Employee{}
	}
	// Handle nested slice element
	builtValue := team.Build().(*models.Employee)
	b.project.Team = append(b.project.Team, builtValue)
	return b
}
// AddTask adds a single item to the Tasks slice
func (b *ProjectBuilder) AddTask(task *TaskBuilder) *ProjectBuilder {
	// Ensure the slice is initialized
	if b.project.Tasks == nil {
		b.project.Tasks = []*models.Task{}
	}
	// Handle nested slice element
	builtValue := task.Build().(*models.Task)
	b.project.Tasks = append(b.project.Tasks, builtValue)
	return b
}
// AddSubProject adds a single item to the SubProjects slice
func (b *ProjectBuilder) AddSubProject(subProject *ProjectBuilder) *ProjectBuilder {
	// Ensure the slice is initialized
	if b.project.SubProjects == nil {
		b.project.SubProjects = []*models.Project{}
	}
	// Handle nested slice element
	builtValue := subProject.Build().(*models.Project)
	b.project.SubProjects = append(b.project.SubProjects, builtValue)
	return b
}

// WithValidation adds a custom validation function
func (b *ProjectBuilder) WithValidation(validationFunc func(*models.Project) error) *ProjectBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Project
func (b *ProjectBuilder) Build() interface{} {
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
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(project).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return project, err
		}
	}

	return project, nil
}

// MustBuild builds the Project and panics if validation fails
func (b *ProjectBuilder) MustBuild() *models.Project {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *ProjectBuilder) Clone() *ProjectBuilder {
	clonedProject := *b.project
	return &ProjectBuilder{
		project: &clonedProject,
		validationFuncs: append([]func(*models.Project) error{}, b.validationFuncs...),
	}
}
