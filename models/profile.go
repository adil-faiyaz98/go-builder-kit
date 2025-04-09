package models

import (
	"fmt"
	"strings"
	"time"
)

// Profile represents a person's profile information
type Profile struct {
	Address        Address
	Education      []string // Just store education IDs to avoid circular imports
	Skills         []*Skill
	Certifications []*Certification
	SocialMedia    SocialMedia
	Biography      string
	Interests      []string
	Languages      []*Language
}

// Validate validates the Profile model
func (p *Profile) Validate() error {
	var errors []string

	// Validate Address
	if err := p.Address.Validate(); err != nil {
		errors = append(errors, fmt.Sprintf("Address validation failed: %s", err.Error()))
	}

	// Validate Skills if provided
	for i, skill := range p.Skills {
		if skill != nil {
			if err := skill.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("Skill[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Validate Certifications if provided
	for i, cert := range p.Certifications {
		if cert != nil {
			if err := cert.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("Certification[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Validate SocialMedia
	if err := p.SocialMedia.Validate(); err != nil {
		errors = append(errors, fmt.Sprintf("SocialMedia validation failed: %s", err.Error()))
	}

	// Validate Languages if provided
	for i, lang := range p.Languages {
		if lang != nil {
			if err := lang.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("Language[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// Skill represents a person's skill or ability
type Skill struct {
	Name              string
	Category          string
	Level             int // 1-5 scale
	YearsOfExperience int
	Endorsements      int
}

// Validate validates the Skill model
func (s *Skill) Validate() error {
	var errors []string

	// Validate Name
	if s.Name == "" {
		errors = append(errors, "Name cannot be empty")
	}

	// Validate Level
	if s.Level < 1 || s.Level > 5 {
		errors = append(errors, "Level must be between 1 and 5")
	}

	// Validate YearsOfExperience
	if s.YearsOfExperience < 0 {
		errors = append(errors, "YearsOfExperience cannot be negative")
	} else if s.YearsOfExperience > 100 {
		errors = append(errors, "YearsOfExperience cannot be greater than 100")
	}

	// Validate Endorsements
	if s.Endorsements < 0 {
		errors = append(errors, "Endorsements cannot be negative")
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// Certification represents a professional certification
type Certification struct {
	Name         string
	Issuer       string
	IssueDate    string
	ExpiryDate   string
	CredentialID string
	URL          string
}

// Validate validates the Certification model
func (c *Certification) Validate() error {
	var errors []string

	// Validate Name
	if c.Name == "" {
		errors = append(errors, "Name cannot be empty")
	}

	// Validate Issuer
	if c.Issuer == "" {
		errors = append(errors, "Issuer cannot be empty")
	}

	// Validate IssueDate if provided
	if c.IssueDate != "" {
		issueDate, err := time.Parse("2006-01-02", c.IssueDate)
		if err != nil {
			errors = append(errors, "IssueDate must be in the format YYYY-MM-DD")
		} else {
			// Check if issue date is in the future
			if issueDate.After(time.Now()) {
				errors = append(errors, "IssueDate cannot be in the future")
			}

			// Validate ExpiryDate if provided
			if c.ExpiryDate != "" {
				expiryDate, err := time.Parse("2006-01-02", c.ExpiryDate)
				if err != nil {
					errors = append(errors, "ExpiryDate must be in the format YYYY-MM-DD")
				} else if expiryDate.Before(issueDate) {
					errors = append(errors, "ExpiryDate cannot be before IssueDate")
				}
			}
		}
	}

	// Validate URL if provided
	if c.URL != "" {
		// Simple validation for URL format
		if !strings.HasPrefix(c.URL, "http://") && !strings.HasPrefix(c.URL, "https://") {
			errors = append(errors, "URL must start with http:// or https://")
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// SocialMedia represents a person's social media presence
type SocialMedia struct {
	LinkedIn      string
	Twitter       string
	Facebook      string
	Instagram     string
	GitHub        string
	Website       string
	OtherProfiles map[string]string
}

// Validate validates the SocialMedia model
func (s *SocialMedia) Validate() error {
	var errors []string

	// Validate LinkedIn if provided
	if s.LinkedIn != "" && !strings.HasPrefix(s.LinkedIn, "https://www.linkedin.com/") {
		errors = append(errors, "LinkedIn URL must start with https://www.linkedin.com/")
	}

	// Validate Twitter if provided
	if s.Twitter != "" && !strings.HasPrefix(s.Twitter, "https://twitter.com/") {
		errors = append(errors, "Twitter URL must start with https://twitter.com/")
	}

	// Validate Facebook if provided
	if s.Facebook != "" && !strings.HasPrefix(s.Facebook, "https://www.facebook.com/") {
		errors = append(errors, "Facebook URL must start with https://www.facebook.com/")
	}

	// Validate Instagram if provided
	if s.Instagram != "" && !strings.HasPrefix(s.Instagram, "https://www.instagram.com/") {
		errors = append(errors, "Instagram URL must start with https://www.instagram.com/")
	}

	// Validate GitHub if provided
	if s.GitHub != "" && !strings.HasPrefix(s.GitHub, "https://github.com/") {
		errors = append(errors, "GitHub URL must start with https://github.com/")
	}

	// Validate Website if provided
	if s.Website != "" {
		// Simple validation for website format
		if !strings.HasPrefix(s.Website, "http://") && !strings.HasPrefix(s.Website, "https://") {
			errors = append(errors, "Website must start with http:// or https://")
		}
	}

	// Validate OtherProfiles if provided
	if s.OtherProfiles != nil {
		for platform, url := range s.OtherProfiles {
			if url == "" {
				errors = append(errors, fmt.Sprintf("URL for %s cannot be empty", platform))
			} else if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
				errors = append(errors, fmt.Sprintf("URL for %s must start with http:// or https://", platform))
			}
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// Language represents a language a person knows
type Language struct {
	Name          string
	Proficiency   string // Beginner, Intermediate, Advanced, Native
	Certification *Certification
}

// Validate validates the Language model
func (l *Language) Validate() error {
	var errors []string

	// Validate Name
	if l.Name == "" {
		errors = append(errors, "Name cannot be empty")
	}

	// Validate Proficiency
	if l.Proficiency != "" {
		validProficiencies := []string{"beginner", "intermediate", "advanced", "native"}
		isValidProficiency := false
		for _, proficiency := range validProficiencies {
			if strings.ToLower(l.Proficiency) == proficiency {
				isValidProficiency = true
				break
			}
		}
		if !isValidProficiency {
			errors = append(errors, "Proficiency must be one of: beginner, intermediate, advanced, native")
		}
	}

	// Validate Certification if provided
	if l.Certification != nil {
		if err := l.Certification.Validate(); err != nil {
			errors = append(errors, fmt.Sprintf("Certification validation failed: %s", err.Error()))
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}
