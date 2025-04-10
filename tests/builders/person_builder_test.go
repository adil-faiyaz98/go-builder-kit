package builders_test

import (
	"testing"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
)

func TestPersonBuilder(t *testing.T) {
	// Create a GeoLocation builder
	geoLocationBuilder := builders.NewGeoLocationBuilder().
		WithLatitude(37.7749).
		WithLongitude(-122.4194).
		WithAccuracy(10.0)

	// Create an Address builder
	addressBuilder := builders.NewAddressBuilder().
		WithStreet("123 Main St").
		WithCity("San Francisco").
		WithState("CA").
		WithPostalCode("94105").
		WithCountry("USA").
		WithCoordinates(geoLocationBuilder).
		WithType("Home").
		WithIsPrimary(true)

	// Create a Course builder
	courseBuilder := builders.NewCourseBuilder().
		WithCode("CS101").
		WithName("Introduction to Computer Science").
		WithDescription("Basic concepts of computer science").
		WithCredits(3.0).
		WithGrade("A").
		WithSemester("Fall").
		WithYear(2020).
		WithInstructor("Dr. Smith")

	// Create an Education builder
	educationBuilder := builders.NewEducationBuilder().
		WithDegree("Bachelor of Science").
		WithInstitution("University of Example").
		WithLocation(addressBuilder.Clone()).
		WithStartDate("2018-09-01").
		WithEndDate("2022-05-31").
		WithGPA(3.8).
		WithMajor("Computer Science").
		WithMinor("Mathematics").
		AddCourse(courseBuilder)

	// Create a Department builder
	departmentBuilder := builders.NewDepartmentBuilder().
		WithName("Engineering").
		WithCode("ENG").
		WithDescription("Engineering department").
		WithManager("John Doe").
		WithBudget(1000000.0).
		WithHeadCount(50)

	// Create a Task builder
	taskBuilder := builders.NewTaskBuilder().
		WithID("TASK-001").
		WithTitle("Implement feature X").
		WithDescription("Implement the new feature X").
		WithStatus("In Progress").
		WithPriority("High").
		WithDueDate("2023-12-31").
		WithAssignedTo("Jane Smith").
		WithCreatedBy("John Doe").
		WithCreatedAt("2023-01-01").
		WithUpdatedAt("2023-01-15").
		AddTag("feature").
		AddTag("backend").
		WithEstimatedHours(8.0).
		WithActualHours(6.5).
		WithProgress(75.0)

	// Create a Project builder (not used in this test but would be used in a real application)
	_ = builders.NewProjectBuilder().
		WithName("Project Alpha").
		WithDescription("A project to build a new system").
		WithStartDate("2023-01-01").
		WithEndDate("2023-12-31").
		WithStatus("In Progress").
		WithBudget(500000.0).
		WithManager("John Doe").
		AddTeamMember("John Doe").
		AddTeamMember("Jane Smith").
		AddTask(taskBuilder)

	// Create a Company builder
	companyBuilder := builders.NewCompanyBuilder().
		WithName("ACME Corp").
		WithIndustry("Technology").
		WithDescription("A leading technology company").
		WithFounded("2000-01-01").
		WithWebsite("https://acmecorp.example.com").
		WithAddress(addressBuilder.Clone()).
		WithSize("large").
		WithRevenue("$1B+").
		WithPublic(true).
		WithStockSymbol("ACME").
		AddDepartment(departmentBuilder)

	// Create an Employment builder
	employmentBuilder := builders.NewEmploymentBuilder().
		WithCompany(companyBuilder).
		WithPosition("Software Engineer").
		WithDepartment("Engineering").
		WithStartDate("2022-06-01").
		WithEndDate("").
		WithSalary(100000.0).
		WithIsCurrent(true)

	// Create a Bond builder (not used in this test but would be used in a real application)
	_ = builders.NewBondBuilder().
		WithID("BOND-1234").
		WithISIN("US123456AB12").
		WithName("US Treasury Bond").
		WithIssuer("US Treasury").
		WithType("government").
		WithMaturityDate("2030-01-01").
		WithPurchaseDate("2020-01-01").
		WithCouponRate(2.5).
		WithFaceValue(1000.0).
		WithPurchasePrice(980.0).
		WithCurrentPrice(950.0).
		WithQuantity(10).
		WithCurrency("USD").
		WithPaymentFrequency("semi-annual").
		WithRating("AAA").
		WithYield(2.8)

	// Create an Account builder
	accountBuilder := builders.NewAccountBuilder().
		WithID("ACC-1234").
		WithNumber("1234567890").
		WithType("Savings").
		WithBalance(5000.0).
		WithCurrency("USD").
		WithOpenDate("2020-01-01").
		WithStatus("Active").
		WithInterestRate(1.5).
		WithIsJoint(false).
		WithOverdraftLimit(1000.0)

	// Create a Bank builder
	bankBuilder := builders.NewBankBuilder().
		WithName("Example Bank").
		WithBranchCode("Downtown").
		WithAddress(addressBuilder.Clone()).
		AddAccount(accountBuilder)

	// Create a Personal Preferences builder
	preferencesBuilder := builders.NewPersonalPreferencesBuilder().
		WithFavoriteColor("Blue").
		AddFavoriteColor("Green").
		WithFavoriteFood("Pizza").
		AddFavoriteFood("Sushi").
		WithFavoriteMusic("Rock").
		AddMusicTaste("Jazz").
		WithFavoriteMovie("The Matrix").
		AddMovieGenre("Sci-Fi").
		WithFavoriteBook("1984").
		AddBookGenre("Dystopian").
		WithFavoriteSport("Basketball").
		WithFavoriteAnimal("Dog").
		AddHobby("Reading").
		AddHobby("Hiking").
		AddInterest("Technology").
		AddInterest("Travel").
		AddLanguage("English").
		AddLanguage("Spanish").
		AddTravelPreference("Accommodation", "Hotel").
		AddShoppingPreference("Online", true)

	// Create a Person builder
	personBuilder := builders.NewPersonBuilder().
		WithID("P12345").
		WithName("John Smith").
		WithAge(30).
		WithEmail("john.smith@example.com").
		WithPhone("+1-555-123-4567").
		WithBirthdate("1993-05-15").
		WithGender("Male").
		WithNationality("American").
		WithMaritalStatus("Single").
		WithAddress(addressBuilder).
		WithEducation(educationBuilder).
		WithBank(bankBuilder).
		WithEmployment(employmentBuilder).
		WithPreferences(preferencesBuilder).
		WithCreatedAt("2023-01-01").
		WithUpdatedAt("2023-01-15")

	// Build the Person
	person, err := personBuilder.BuildAndValidate()
	if err != nil {
		t.Fatalf("Failed to build person: %v", err)
	}

	// Verify Person fields
	if person.ID != "P12345" {
		t.Errorf("Expected ID to be 'P12345', got '%s'", person.ID)
	}

	if person.Name != "John Smith" {
		t.Errorf("Expected name to be 'John Smith', got '%s'", person.Name)
	}

	if person.Age != 30 {
		t.Errorf("Expected age to be 30, got %d", person.Age)
	}

	// Verify Address fields
	if person.Address == nil {
		t.Fatalf("Expected address to be non-nil")
	}

	if person.Address.City != "San Francisco" {
		t.Errorf("Expected city to be 'San Francisco', got '%s'", person.Address.City)
	}

	// Verify Education fields
	if person.Education == nil {
		t.Fatalf("Expected education to be non-nil")
	}

	if person.Education.Degree != "Bachelor of Science" {
		t.Errorf("Expected degree to be 'Bachelor of Science', got '%s'", person.Education.Degree)
	}

	if len(person.Education.Courses) != 1 {
		t.Fatalf("Expected 1 course, got %d", len(person.Education.Courses))
	}

	// Verify Employment fields
	if person.Employment == nil {
		t.Fatalf("Expected employment to be non-nil")
	}

	if person.Employment.Company == nil {
		t.Fatalf("Expected company to be non-nil")
	}

	if person.Employment.Company.Name != "ACME Corp" {
		t.Errorf("Expected company name to be 'ACME Corp', got '%s'", person.Employment.Company.Name)
	}

	if person.Employment.Position != "Software Engineer" {
		t.Errorf("Expected position to be 'Software Engineer', got '%s'", person.Employment.Position)
	}

	// Verify Bank fields
	if person.Bank == nil {
		t.Fatalf("Expected bank to be non-nil")
	}

	if person.Bank.Name != "Example Bank" {
		t.Errorf("Expected bank name to be 'Example Bank', got '%s'", person.Bank.Name)
	}

	if len(person.Bank.Accounts) != 1 {
		t.Fatalf("Expected 1 account, got %d", len(person.Bank.Accounts))
	}

	// Verify Preferences fields
	if person.Preferences == nil {
		t.Fatalf("Expected preferences to be non-nil")
	}

	if person.Preferences.FavoriteColor != "Blue" {
		t.Errorf("Expected favorite color to be 'Blue', got '%s'", person.Preferences.FavoriteColor)
	}

	if len(person.Preferences.Hobbies) != 2 {
		t.Fatalf("Expected 2 hobbies, got %d", len(person.Preferences.Hobbies))
	}
}
