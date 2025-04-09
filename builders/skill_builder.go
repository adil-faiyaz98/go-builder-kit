package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// SkillBuilder builds a Skill model
type SkillBuilder struct {
	skill *models.Skill
	// Custom validation functions
	validationFuncs []func(*models.Skill) error
}

// NewSkillBuilder creates a new SkillBuilder
func NewSkillBuilder() *SkillBuilder {
	return &SkillBuilder{
		skill: &models.Skill{
			Name: "",
			Category: "",
			Level: 0,
			YearsOfExperience: 0,
			Endorsements: 0,
		},
		validationFuncs: []func(*models.Skill) error{},
	}
}

// NewSkillBuilderWithDefaults creates a new SkillBuilder with sensible defaults
func NewSkillBuilderWithDefaults() *SkillBuilder {
	builder := NewSkillBuilder()
	// Add default values here if needed
	return builder
}
// WithName sets the Name
func (b *SkillBuilder) WithName(name string) *SkillBuilder {
	b.skill.Name = name
	return b
}

// WithCategory sets the Category
func (b *SkillBuilder) WithCategory(category string) *SkillBuilder {
	b.skill.Category = category
	return b
}

// WithLevel sets the Level
func (b *SkillBuilder) WithLevel(level int) *SkillBuilder {
	b.skill.Level = level
	return b
}

// WithYearsOfExperience sets the YearsOfExperience
func (b *SkillBuilder) WithYearsOfExperience(yearsOfExperience int) *SkillBuilder {
	b.skill.YearsOfExperience = yearsOfExperience
	return b
}

// WithEndorsements sets the Endorsements
func (b *SkillBuilder) WithEndorsements(endorsements int) *SkillBuilder {
	b.skill.Endorsements = endorsements
	return b
}


// WithValidation adds a custom validation function
func (b *SkillBuilder) WithValidation(validationFunc func(*models.Skill) error) *SkillBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Skill
func (b *SkillBuilder) Build() interface{} {
	return b.skill
}

// BuildPtr builds the Skill and returns a pointer
func (b *SkillBuilder) BuildPtr() *models.Skill {
	return b.skill
}

// BuildAndValidate builds the Skill and validates it
func (b *SkillBuilder) BuildAndValidate() (*models.Skill, error) {
	skill := b.skill

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(skill); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(skill).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return skill, err
		}
	}

	return skill, nil
}

// MustBuild builds the Skill and panics if validation fails
func (b *SkillBuilder) MustBuild() *models.Skill {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *SkillBuilder) Clone() *SkillBuilder {
	clonedSkill := *b.skill
	return &SkillBuilder{
		skill: &clonedSkill,
		validationFuncs: append([]func(*models.Skill) error{}, b.validationFuncs...),
	}
}
