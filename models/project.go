package models

import (
	"fmt"
	"strings"
	"time"
)

// Project represents a project
type Project struct {
	Name        string
	Description string
	StartDate   string
	EndDate     string
	Status      string
	Budget      float64
	Manager     interface{}
	Team        []interface{}
	Members     []interface{}
	Tasks       []*Task
}

// Validate validates the Project model
func (p *Project) Validate() error {
	var errors []string

	// Validate Name
	if p.Name == "" {
		errors = append(errors, "Name cannot be empty")
	}

	// Validate StartDate if provided
	if p.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", p.StartDate)
		if err != nil {
			errors = append(errors, "StartDate must be in the format YYYY-MM-DD")
		} else {
			// Validate EndDate if provided
			if p.EndDate != "" {
				endDate, err := time.Parse("2006-01-02", p.EndDate)
				if err != nil {
					errors = append(errors, "EndDate must be in the format YYYY-MM-DD")
				} else if endDate.Before(startDate) {
					errors = append(errors, "EndDate cannot be before StartDate")
				}
			}
		}
	}

	// Validate Budget
	if p.Budget < 0 {
		errors = append(errors, "Budget cannot be negative")
	}

	// Validate Status if provided
	if p.Status != "" {
		validStatuses := []string{"planning", "in-progress", "on-hold", "completed", "cancelled"}
		isValidStatus := false
		for _, status := range validStatuses {
			if strings.ToLower(p.Status) == status {
				isValidStatus = true
				break
			}
		}
		if !isValidStatus {
			errors = append(errors, "Status must be one of: planning, in-progress, on-hold, completed, cancelled")
		}
	}

	// Validate Tasks if provided
	for i, task := range p.Tasks {
		if task != nil {
			if err := task.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("Task[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// Task represents a task in a project
type Task struct {
	Name        string
	Description string
	StartDate   string
	EndDate     string
	Status      string
	Priority    string
	Assignee    interface{}
	Subtasks    []*Task
}

// Validate validates the Task model
func (t *Task) Validate() error {
	var errors []string

	// Validate Name
	if t.Name == "" {
		errors = append(errors, "Name cannot be empty")
	}

	// Validate StartDate if provided
	if t.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", t.StartDate)
		if err != nil {
			errors = append(errors, "StartDate must be in the format YYYY-MM-DD")
		} else {
			// Validate EndDate if provided
			if t.EndDate != "" {
				endDate, err := time.Parse("2006-01-02", t.EndDate)
				if err != nil {
					errors = append(errors, "EndDate must be in the format YYYY-MM-DD")
				} else if endDate.Before(startDate) {
					errors = append(errors, "EndDate cannot be before StartDate")
				}
			}
		}
	}

	// Validate Status if provided
	if t.Status != "" {
		validStatuses := []string{"not-started", "in-progress", "completed", "blocked", "deferred"}
		isValidStatus := false
		for _, status := range validStatuses {
			if strings.ToLower(t.Status) == status {
				isValidStatus = true
				break
			}
		}
		if !isValidStatus {
			errors = append(errors, "Status must be one of: not-started, in-progress, completed, blocked, deferred")
		}
	}

	// Validate Priority if provided
	if t.Priority != "" {
		validPriorities := []string{"low", "medium", "high", "critical"}
		isValidPriority := false
		for _, priority := range validPriorities {
			if strings.ToLower(t.Priority) == priority {
				isValidPriority = true
				break
			}
		}
		if !isValidPriority {
			errors = append(errors, "Priority must be one of: low, medium, high, critical")
		}
	}

	// Validate Subtasks if provided
	for i, subtask := range t.Subtasks {
		if subtask != nil {
			if err := subtask.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("Subtask[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}
