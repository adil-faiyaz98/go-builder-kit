package models

// Address represents a physical address
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
	Country    string
	Location   *GeoLocation
}

// GeoLocation represents geographic coordinates
type GeoLocation struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
	Accuracy  float64
}

// Contact represents contact information
type Contact struct {
	Email       string
	Phone       string
	Address     *Address
	Alternative *Contact
}

// Department represents a department in an organization
type Department struct {
	Name        string
	Description string
	Manager     *Employee
	Employees   []*Employee
	SubDepts    []*Department
	Location    *Address
}

// Project represents a project
type Project struct {
	Name        string
	Description string
	StartDate   string
	EndDate     string
	Status      string
	Budget      float64
	Manager     *Employee
	Team        []*Employee
	Tasks       []*Task
	SubProjects []*Project
}

// Task represents a task in a project
type Task struct {
	Name        string
	Description string
	Status      string
	Priority    string
	DueDate     string
	Assignee    *Employee
	SubTasks    []*Task
	Project     *Project
}

// Employee represents an employee
type Employee struct {
	ID          string
	Name        string
	Title       string
	Department  *Department
	Manager     *Employee
	Subordinates []*Employee
	Contact     *Contact
	Projects    []*Project
	Skills      []string
	StartDate   string
}

// Company represents a company
type Company struct {
	Name        string
	Description string
	Founded     string
	Industry    string
	Website     string
	Headquarters *Address
	Departments []*Department
	Employees   []*Employee
	Projects    []*Project
	Subsidiaries []*Company
	ParentCompany *Company
}

// Organization represents a complex organization structure
type Organization struct {
	Name         string
	Type         string
	Description  string
	Headquarters *Address
	Companies    []*Company
	Departments  []*Department
	Employees    []*Employee
	Projects     []*Project
	Contacts     []*Contact
}
