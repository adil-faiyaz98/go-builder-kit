package models

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// DigitalProfile represents a person's digital presence
type DigitalProfile struct {
	Devices     []*Device
	Accounts    []*OnlineAccount
	Preferences *UserPreferences
	Activity    []*UserActivity
}

// Validate validates the DigitalProfile model
func (d *DigitalProfile) Validate() error {
	var errors []string

	// Validate Devices if provided
	for i, device := range d.Devices {
		if device != nil {
			if err := device.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("Device[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Validate Accounts if provided
	for i, account := range d.Accounts {
		if account != nil {
			if err := account.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("Account[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Validate Preferences if provided
	if d.Preferences != nil {
		if err := d.Preferences.Validate(); err != nil {
			errors = append(errors, fmt.Sprintf("Preferences validation failed: %s", err.Error()))
		}
	}

	// Validate Activity if provided
	for i, activity := range d.Activity {
		if activity != nil {
			if err := activity.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("Activity[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// Device represents an electronic device
type Device struct {
	Type         string // Smartphone, Laptop, Tablet, etc.
	Model        string
	SerialNumber string
	PurchaseDate string
	OS           string
	LastUsed     string
	Apps         []*Application
	Settings     map[string]string
}

// Validate validates the Device model
func (d *Device) Validate() error {
	var errors []string

	// Validate Type
	if d.Type == "" {
		errors = append(errors, "Type cannot be empty")
	} else {
		validTypes := []string{"smartphone", "laptop", "tablet", "desktop", "smartwatch", "other"}
		isValidType := false
		for _, deviceType := range validTypes {
			if strings.ToLower(d.Type) == deviceType {
				isValidType = true
				break
			}
		}
		if !isValidType {
			errors = append(errors, "Type must be one of: smartphone, laptop, tablet, desktop, smartwatch, other")
		}
	}

	// Validate Model
	if d.Model == "" {
		errors = append(errors, "Model cannot be empty")
	}

	// Validate PurchaseDate if provided
	if d.PurchaseDate != "" {
		purchaseDate, err := time.Parse("2006-01-02", d.PurchaseDate)
		if err != nil {
			errors = append(errors, "PurchaseDate must be in the format YYYY-MM-DD")
		} else if purchaseDate.After(time.Now()) {
			errors = append(errors, "PurchaseDate cannot be in the future")
		}
	}

	// Validate LastUsed if provided
	if d.LastUsed != "" {
		lastUsed, err := time.Parse("2006-01-02", d.LastUsed)
		if err != nil {
			errors = append(errors, "LastUsed must be in the format YYYY-MM-DD")
		} else if lastUsed.After(time.Now()) {
			errors = append(errors, "LastUsed cannot be in the future")
		}
	}

	// Validate Apps if provided
	for i, app := range d.Apps {
		if app != nil {
			if err := app.Validate(); err != nil {
				errors = append(errors, fmt.Sprintf("App[%d] validation failed: %s", i, err.Error()))
			}
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// Application represents a software application
type Application struct {
	Name        string
	Version     string
	InstallDate string
	LastUsed    string
	Settings    map[string]string
	Permissions []string
}

// Validate validates the Application model
func (a *Application) Validate() error {
	var errors []string

	// Validate Name
	if a.Name == "" {
		errors = append(errors, "Name cannot be empty")
	}

	// Validate InstallDate if provided
	if a.InstallDate != "" {
		installDate, err := time.Parse("2006-01-02", a.InstallDate)
		if err != nil {
			errors = append(errors, "InstallDate must be in the format YYYY-MM-DD")
		} else if installDate.After(time.Now()) {
			errors = append(errors, "InstallDate cannot be in the future")
		}

		// Validate LastUsed if provided
		if a.LastUsed != "" {
			lastUsed, err := time.Parse("2006-01-02", a.LastUsed)
			if err != nil {
				errors = append(errors, "LastUsed must be in the format YYYY-MM-DD")
			} else if lastUsed.After(time.Now()) {
				errors = append(errors, "LastUsed cannot be in the future")
			} else if lastUsed.Before(installDate) {
				errors = append(errors, "LastUsed cannot be before InstallDate")
			}
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// OnlineAccount represents an online service account
type OnlineAccount struct {
	Platform     string
	Username     string
	Email        string
	CreationDate string
	LastLogin    string
	Status       string // Active, Inactive, Suspended
	Settings     map[string]string
	Subscription *Subscription
}

// Validate validates the OnlineAccount model
func (o *OnlineAccount) Validate() error {
	var errors []string

	// Validate Platform
	if o.Platform == "" {
		errors = append(errors, "Platform cannot be empty")
	}

	// Validate Username or Email
	if o.Username == "" && o.Email == "" {
		errors = append(errors, "Either Username or Email must be provided")
	}

	// Validate Email if provided
	if o.Email != "" {
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(o.Email) {
			errors = append(errors, "Email is not valid")
		}
	}

	// Validate CreationDate if provided
	if o.CreationDate != "" {
		creationDate, err := time.Parse("2006-01-02", o.CreationDate)
		if err != nil {
			errors = append(errors, "CreationDate must be in the format YYYY-MM-DD")
		} else if creationDate.After(time.Now()) {
			errors = append(errors, "CreationDate cannot be in the future")
		}

		// Validate LastLogin if provided
		if o.LastLogin != "" {
			lastLogin, err := time.Parse("2006-01-02", o.LastLogin)
			if err != nil {
				errors = append(errors, "LastLogin must be in the format YYYY-MM-DD")
			} else if lastLogin.After(time.Now()) {
				errors = append(errors, "LastLogin cannot be in the future")
			} else if lastLogin.Before(creationDate) {
				errors = append(errors, "LastLogin cannot be before CreationDate")
			}
		}
	}

	// Validate Status if provided
	if o.Status != "" {
		validStatuses := []string{"active", "inactive", "suspended", "closed"}
		isValidStatus := false
		for _, status := range validStatuses {
			if strings.ToLower(o.Status) == status {
				isValidStatus = true
				break
			}
		}
		if !isValidStatus {
			errors = append(errors, "Status must be one of: active, inactive, suspended, closed")
		}
	}

	// Validate Subscription if provided
	if o.Subscription != nil {
		if err := o.Subscription.Validate(); err != nil {
			errors = append(errors, fmt.Sprintf("Subscription validation failed: %s", err.Error()))
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// Subscription represents a service subscription
type Subscription struct {
	Plan         string
	StartDate    string
	EndDate      string
	Price        float64
	BillingCycle string
	AutoRenew    bool
	Status       string
	Features     []string
}

// Validate validates the Subscription model
func (s *Subscription) Validate() error {
	var errors []string

	// Validate Plan
	if s.Plan == "" {
		errors = append(errors, "Plan cannot be empty")
	}

	// Validate StartDate if provided
	if s.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", s.StartDate)
		if err != nil {
			errors = append(errors, "StartDate must be in the format YYYY-MM-DD")
		} else {
			// Validate EndDate if provided
			if s.EndDate != "" {
				endDate, err := time.Parse("2006-01-02", s.EndDate)
				if err != nil {
					errors = append(errors, "EndDate must be in the format YYYY-MM-DD")
				} else if endDate.Before(startDate) {
					errors = append(errors, "EndDate cannot be before StartDate")
				}
			}
		}
	}

	// Validate Price
	if s.Price < 0 {
		errors = append(errors, "Price cannot be negative")
	}

	// Validate BillingCycle if provided
	if s.BillingCycle != "" {
		validCycles := []string{"monthly", "quarterly", "annually", "one-time"}
		isValidCycle := false
		for _, cycle := range validCycles {
			if strings.ToLower(s.BillingCycle) == cycle {
				isValidCycle = true
				break
			}
		}
		if !isValidCycle {
			errors = append(errors, "BillingCycle must be one of: monthly, quarterly, annually, one-time")
		}
	}

	// Validate Status if provided
	if s.Status != "" {
		validStatuses := []string{"active", "expired", "cancelled", "pending"}
		isValidStatus := false
		for _, status := range validStatuses {
			if strings.ToLower(s.Status) == status {
				isValidStatus = true
				break
			}
		}
		if !isValidStatus {
			errors = append(errors, "Status must be one of: active, expired, cancelled, pending")
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// UserPreferences represents user preferences for digital services
type UserPreferences struct {
	Theme                string
	Language             string
	Notifications        map[string]bool
	Privacy              map[string]string
	Accessibility        map[string]bool
	DefaultCommunication string
}

// Validate validates the UserPreferences model
func (u *UserPreferences) Validate() error {
	var errors []string

	// Validate Theme if provided
	if u.Theme != "" {
		validThemes := []string{"light", "dark", "system", "custom"}
		isValidTheme := false
		for _, theme := range validThemes {
			if strings.ToLower(u.Theme) == theme {
				isValidTheme = true
				break
			}
		}
		if !isValidTheme {
			errors = append(errors, "Theme must be one of: light, dark, system, custom")
		}
	}

	// Validate DefaultCommunication if provided
	if u.DefaultCommunication != "" {
		validCommunications := []string{"email", "phone", "sms", "app", "none"}
		isValidCommunication := false
		for _, comm := range validCommunications {
			if strings.ToLower(u.DefaultCommunication) == comm {
				isValidCommunication = true
				break
			}
		}
		if !isValidCommunication {
			errors = append(errors, "DefaultCommunication must be one of: email, phone, sms, app, none")
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

// UserActivity represents a user's digital activity
type UserActivity struct {
	Type        string
	Platform    string
	Date        string
	Timestamp   string
	Duration    int // in minutes
	Description string
	Metadata    map[string]string
	Device      string
	Location    *GeoLocation
	Details     map[string]string
}

// Validate validates the UserActivity model
func (u *UserActivity) Validate() error {
	var errors []string

	// Validate Type
	if u.Type == "" {
		errors = append(errors, "Type cannot be empty")
	}

	// Validate Platform
	if u.Platform == "" {
		errors = append(errors, "Platform cannot be empty")
	}

	// Validate Date
	if u.Date == "" {
		errors = append(errors, "Date cannot be empty")
	} else {
		date, err := time.Parse("2006-01-02T15:04:05Z", u.Date)
		if err != nil {
			errors = append(errors, "Date must be in the format YYYY-MM-DDThh:mm:ssZ")
		} else if date.After(time.Now()) {
			errors = append(errors, "Date cannot be in the future")
		}
	}

	// Validate Duration
	if u.Duration < 0 {
		errors = append(errors, "Duration cannot be negative")
	}

	// Validate Location if provided
	if u.Location != nil {
		if err := u.Location.Validate(); err != nil {
			errors = append(errors, fmt.Sprintf("Location validation failed: %s", err.Error()))
		}
	}

	// Return errors if any
	if len(errors) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return nil
}
