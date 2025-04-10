package builders

import (
	"testing"
)

func TestNestedBuilders(t *testing.T) {
	// Test GeoLocation builder
	geoLocationBuilder := NewGeoLocationBuilder().
		WithLatitude(37.7749).
		WithLongitude(-122.4194).
		WithAltitude(0).
		WithAccuracy(10)

	// Test Address builder with nested GeoLocation
	addressBuilder := NewAddressBuilder().
		WithStreet("123 Main St").
		WithCity("San Francisco").
		WithState("CA").
		WithPostalCode("94105").
		WithCountry("USA").
		WithLocation(geoLocationBuilder)

	// Test Contact builder with nested Address
	contactBuilder := NewContactBuilder().
		WithEmail("john@example.com").
		WithPhone("555-1234").
		WithAddress(addressBuilder)

	// Create a self-referential Contact
	alternativeContactBuilder := NewContactBuilder().
		WithEmail("jane@example.com").
		WithPhone("555-5678")
	
	contactBuilder.WithAlternative(alternativeContactBuilder)

	// Build the Contact
	contact := contactBuilder.BuildPtr()

	// Verify Contact fields
	if contact.Email != "john@example.com" {
		t.Errorf("Expected Email to be 'john@example.com', got '%s'", contact.Email)
	}

	if contact.Phone != "555-1234" {
		t.Errorf("Expected Phone to be '555-1234', got '%s'", contact.Phone)
	}

	// Verify nested Address fields
	if contact.Address == nil {
		t.Fatalf("Expected Address to be non-nil")
	}

	if contact.Address.Street != "123 Main St" {
		t.Errorf("Expected Street to be '123 Main St', got '%s'", contact.Address.Street)
	}

	if contact.Address.City != "San Francisco" {
		t.Errorf("Expected City to be 'San Francisco', got '%s'", contact.Address.City)
	}

	// Verify nested GeoLocation fields
	if contact.Address.Location == nil {
		t.Fatalf("Expected Location to be non-nil")
	}

	if contact.Address.Location.Latitude != 37.7749 {
		t.Errorf("Expected Latitude to be 37.7749, got %f", contact.Address.Location.Latitude)
	}

	if contact.Address.Location.Longitude != -122.4194 {
		t.Errorf("Expected Longitude to be -122.4194, got %f", contact.Address.Location.Longitude)
	}

	// Verify self-referential Contact
	if contact.Alternative == nil {
		t.Fatalf("Expected Alternative to be non-nil")
	}

	if contact.Alternative.Email != "jane@example.com" {
		t.Errorf("Expected Alternative Email to be 'jane@example.com', got '%s'", contact.Alternative.Email)
	}

	// Test complex circular references
	// Create Employee and Department with circular references
	employeeBuilder := NewEmployeeBuilder().
		WithID("E001").
		WithName("John Doe").
		WithTitle("Manager").
		WithContact(contactBuilder)

	departmentBuilder := NewDepartmentBuilder().
		WithName("Engineering").
		WithDescription("Software Engineering").
		WithLocation(addressBuilder)

	// Set circular references
	employeeBuilder.WithDepartment(departmentBuilder)
	departmentBuilder.WithManager(employeeBuilder)

	// Add employee to department
	departmentBuilder.AddEmployee(employeeBuilder)

	// Build the Department
	department := departmentBuilder.BuildPtr()

	// Verify Department fields
	if department.Name != "Engineering" {
		t.Errorf("Expected Name to be 'Engineering', got '%s'", department.Name)
	}

	// Verify circular references
	if department.Manager == nil {
		t.Fatalf("Expected Manager to be non-nil")
	}

	if department.Manager.Name != "John Doe" {
		t.Errorf("Expected Manager Name to be 'John Doe', got '%s'", department.Manager.Name)
	}

	if department.Manager.Department == nil {
		t.Fatalf("Expected Manager's Department to be non-nil")
	}

	if department.Manager.Department.Name != "Engineering" {
		t.Errorf("Expected Manager's Department Name to be 'Engineering', got '%s'", department.Manager.Department.Name)
	}

	// Verify employees slice
	if len(department.Employees) != 1 {
		t.Fatalf("Expected 1 employee, got %d", len(department.Employees))
	}

	if department.Employees[0].Name != "John Doe" {
		t.Errorf("Expected Employee Name to be 'John Doe', got '%s'", department.Employees[0].Name)
	}

	// Test deeply nested structures
	// Create a complex organization structure
	organizationBuilder := NewOrganizationBuilder().
		WithName("Acme Corporation").
		WithType("Corporation").
		WithDescription("A global technology company").
		WithHeadquarters(addressBuilder)

	companyBuilder := NewCompanyBuilder().
		WithName("Acme Tech").
		WithDescription("Technology division").
		WithIndustry("Technology").
		WithHeadquarters(addressBuilder)

	// Add departments and employees to company
	companyBuilder.AddDepartment(departmentBuilder)
	companyBuilder.AddEmployee(employeeBuilder)

	// Add company to organization
	organizationBuilder.AddCompany(companyBuilder)

	// Build the Organization
	organization := organizationBuilder.BuildPtr()

	// Verify Organization fields
	if organization.Name != "Acme Corporation" {
		t.Errorf("Expected Name to be 'Acme Corporation', got '%s'", organization.Name)
	}

	// Verify nested Company
	if len(organization.Companies) != 1 {
		t.Fatalf("Expected 1 company, got %d", len(organization.Companies))
	}

	if organization.Companies[0].Name != "Acme Tech" {
		t.Errorf("Expected Company Name to be 'Acme Tech', got '%s'", organization.Companies[0].Name)
	}

	// Verify deeply nested Department
	if len(organization.Companies[0].Departments) != 1 {
		t.Fatalf("Expected 1 department, got %d", len(organization.Companies[0].Departments))
	}

	if organization.Companies[0].Departments[0].Name != "Engineering" {
		t.Errorf("Expected Department Name to be 'Engineering', got '%s'", organization.Companies[0].Departments[0].Name)
	}

	// Verify deeply nested Employee
	if len(organization.Companies[0].Employees) != 1 {
		t.Fatalf("Expected 1 employee, got %d", len(organization.Companies[0].Employees))
	}

	if organization.Companies[0].Employees[0].Name != "John Doe" {
		t.Errorf("Expected Employee Name to be 'John Doe', got '%s'", organization.Companies[0].Employees[0].Name)
	}

	// Test cloning
	clonedOrganizationBuilder := organizationBuilder.Clone()
	clonedOrganization := clonedOrganizationBuilder.BuildPtr()

	// Modify the original
	organizationBuilder.WithName("Modified Acme")
	modifiedOrganization := organizationBuilder.BuildPtr()

	// Verify the clone wasn't affected
	if clonedOrganization.Name != "Acme Corporation" {
		t.Errorf("Expected cloned Name to be 'Acme Corporation', got '%s'", clonedOrganization.Name)
	}

	// Verify the modified organization
	if modifiedOrganization.Name != "Modified Acme" {
		t.Errorf("Expected modified Name to be 'Modified Acme', got '%s'", modifiedOrganization.Name)
	}
}
