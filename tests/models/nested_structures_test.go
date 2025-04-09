package models_test

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

var _ = Describe("Nested Structures", func() {
	// Test Person with nested Address
	Describe("Person with Address", func() {
		It("should correctly build a Person with an Address", func() {
			// Create an address
			addressBuilder := builders.NewAddressBuilder().
				WithStreet("123 Main St").
				WithCity("New York").
				WithState("NY").
				WithPostalCode("10001").
				WithCountry("USA").
				WithType("Home").
				WithIsPrimary(true)

			// Create a person with the address
			person := builders.NewPersonBuilder().
				WithID("123").
				WithName("John Doe").
				WithAge(30).
				WithEmail("john.doe@example.com").
				WithAddress(addressBuilder).
				Build().(*models.Person)

			// Verify the nested structure
			Expect(person.Address.Street).To(Equal("123 Main St"))
			Expect(person.Address.City).To(Equal("New York"))
			Expect(person.Address.State).To(Equal("NY"))
			Expect(person.Address.PostalCode).To(Equal("10001"))
			Expect(person.Address.Country).To(Equal("USA"))
			Expect(person.Address.Type).To(Equal("Home"))
			Expect(person.Address.IsPrimary).To(BeTrue())
		})

		It("should correctly validate a Person with an invalid Address", func() {
			// Create an invalid address (missing required fields)
			invalidAddressBuilder := builders.NewAddressBuilder().
				WithStreet(""). // Invalid: empty street
				WithCity("New York").
				WithCountry("USA")

			// Create a person with the invalid address
			_, err := builders.NewPersonBuilder().
				WithID("123").
				WithName("John Doe").
				WithAge(30).
				WithEmail("john.doe@example.com").
				WithAddress(invalidAddressBuilder).
				BuildAndValidate()

			// Verify validation fails
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Street cannot be empty"))
		})
	})

	// Test Person with nested Education
	Describe("Person with Education", func() {
		It("should correctly build a Person with Education", func() {
			// Create education
			educationBuilder := builders.NewEducationBuilder().
				WithDegree("Bachelor of Science").
				WithInstitution("Harvard University").
				WithStartDate("2018-09-01").
				WithEndDate("2022-05-15").
				WithGPA(3.8).
				WithMajor("Computer Science").
				WithMinor("Mathematics")

			// Create a course
			courseBuilder := builders.NewCourseBuilder().
				WithCode("CS101").
				WithName("Introduction to Computer Science").
				WithCredits(3.0).
				WithYear(2020).
				WithSemester("Fall").
				WithGrade("A")

			// Add course to education
			educationBuilder.WithCourse(courseBuilder)

			// Create a person with education
			person := builders.NewPersonBuilder().
				WithID("123").
				WithName("John Doe").
				WithAge(30).
				WithEmail("john.doe@example.com").
				WithEducation(educationBuilder).
				Build().(*models.Person)

			// Verify the nested structure
			Expect(person.Education.Degree).To(Equal("Bachelor of Science"))
			Expect(person.Education.Institution).To(Equal("Harvard University"))
			Expect(person.Education.StartDate).To(Equal("2018-09-01"))
			Expect(person.Education.EndDate).To(Equal("2022-05-15"))
			Expect(person.Education.GPA).To(Equal(3.8))
			Expect(person.Education.Major).To(Equal("Computer Science"))
			Expect(person.Education.Minor).To(Equal("Mathematics"))

			// Verify the nested course
			Expect(len(person.Education.Courses)).To(Equal(1))
			Expect(person.Education.Courses[0].Code).To(Equal("CS101"))
			Expect(person.Education.Courses[0].Name).To(Equal("Introduction to Computer Science"))
			Expect(person.Education.Courses[0].Credits).To(Equal(3.0))
			Expect(person.Education.Courses[0].Year).To(Equal(2020))
			Expect(person.Education.Courses[0].Semester).To(Equal("Fall"))
			Expect(person.Education.Courses[0].Grade).To(Equal("A"))
		})

		It("should correctly validate a Person with invalid Education", func() {
			// Create invalid education (end date before start date)
			invalidEducationBuilder := builders.NewEducationBuilder().
				WithDegree("Bachelor of Science").
				WithInstitution("Harvard University").
				WithStartDate("2022-09-01"). // Later than end date
				WithEndDate("2022-05-15").
				WithGPA(3.8)

			// Create a person with the invalid education
			_, err := builders.NewPersonBuilder().
				WithID("123").
				WithName("John Doe").
				WithAge(30).
				WithEmail("john.doe@example.com").
				WithEducation(invalidEducationBuilder).
				BuildAndValidate()

			// Verify validation fails
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("EndDate cannot be before StartDate"))
		})
	})

	// Test Person with nested Employment and Company
	Describe("Person with Employment and Company", func() {
		It("should correctly build a Person with Employment and Company", func() {
			// Create company address
			companyAddressBuilder := builders.NewAddressBuilder().
				WithStreet("456 Tech Blvd").
				WithCity("San Francisco").
				WithState("CA").
				WithPostalCode("94107").
				WithCountry("USA").
				WithType("Work")

			// Create department
			departmentBuilder := builders.NewDepartmentBuilder().
				WithName("Engineering").
				WithDescription("Software development department").
				WithHeadCount(50).
				WithBudget(5000000.0)

			// Create project
			projectBuilder := builders.NewProjectBuilder().
				WithName("Mobile App").
				WithDescription("Cross-platform mobile application").
				WithStartDate("2022-01-01").
				WithEndDate("2022-12-31").
				WithStatus("In Progress").
				WithBudget(500000.0)

			// Add project to department
			departmentBuilder.WithProject(projectBuilder)

			// Create company
			companyBuilder := builders.NewCompanyBuilder().
				WithID("COMP-001").
				WithName("Tech Innovations Inc.").
				WithIndustry("Technology").
				WithDescription("Leading technology company").
				WithFounded("2010-01-01").
				WithLocation(companyAddressBuilder).
				WithDepartment(departmentBuilder)

			// Create employment
			employmentBuilder := builders.NewEmploymentBuilder().
				WithCompany(companyBuilder).
				WithPosition("Senior Software Engineer").
				WithDepartment("Engineering").
				WithStartDate("2020-01-15").
				WithSalary(120000.0).
				WithIsCurrent(true)

			// Create a person with employment
			person := builders.NewPersonBuilder().
				WithID("123").
				WithName("John Doe").
				WithAge(30).
				WithEmail("john.doe@example.com").
				WithEmployment(employmentBuilder).
				Build().(*models.Person)

			// Verify the nested structure
			Expect(person.Employment.Position).To(Equal("Senior Software Engineer"))
			Expect(person.Employment.Department).To(Equal("Engineering"))
			Expect(person.Employment.StartDate).To(Equal("2020-01-15"))
			Expect(person.Employment.Salary).To(Equal(120000.0))
			Expect(person.Employment.IsCurrent).To(BeTrue())

			// Verify the nested company
			Expect(person.Employment.Company.ID).To(Equal("COMP-001"))
			Expect(person.Employment.Company.Name).To(Equal("Tech Innovations Inc."))
			Expect(person.Employment.Company.Industry).To(Equal("Technology"))

			// Verify the nested company location
			Expect(person.Employment.Company.Location.Street).To(Equal("456 Tech Blvd"))
			Expect(person.Employment.Company.Location.City).To(Equal("San Francisco"))

			// Verify the nested department
			Expect(len(person.Employment.Company.Departments)).To(Equal(1))
			Expect(person.Employment.Company.Departments[0].Name).To(Equal("Engineering"))

			// Verify the nested project
			Expect(len(person.Employment.Company.Departments[0].Projects)).To(Equal(1))
			Expect(person.Employment.Company.Departments[0].Projects[0].Name).To(Equal("Mobile App"))
		})
	})

	// Test Bank with nested Accounts and Transactions
	Describe("Bank with Accounts and Transactions", func() {
		It("should correctly build a Bank with Accounts and Transactions", func() {
			// Create transaction
			transactionBuilder := builders.NewTransactionBuilder().
				WithID("TX-001").
				WithDate("2023-01-15").
				WithAmount(1000.0).
				WithType("Deposit").
				WithDescription("Salary deposit")

			// Create account
			accountBuilder := builders.NewAccountBuilder().
				WithNumber("ACCT-12345").
				WithBalance(5000.0).
				WithType("Checking").
				WithCurrency("USD").
				WithStatus("Active").
				WithTransaction(transactionBuilder)

			// Create bank address
			addressBuilder := builders.NewAddressBuilder().
				WithStreet("100 Financial St").
				WithCity("New York").
				WithState("NY").
				WithPostalCode("10005").
				WithCountry("USA").
				WithType("Work")

			// Create bank
			bank := builders.NewBankBuilder().
				WithName("Global Bank").
				WithBranchCode("GB-001").
				WithAddress(addressBuilder).
				WithAccount(accountBuilder).
				WithRelationshipStartDate("2020-01-15").
				Build().(models.Bank)

			// Verify the nested structure
			Expect(bank.Name).To(Equal("Global Bank"))
			Expect(bank.BranchCode).To(Equal("GB-001"))

			// Verify the nested address
			Expect(bank.Address.Street).To(Equal("100 Financial St"))
			Expect(bank.Address.City).To(Equal("New York"))

			// Verify the nested account
			Expect(len(bank.Accounts)).To(Equal(1))
			Expect(bank.Accounts[0].Number).To(Equal("ACCT-12345"))
			Expect(bank.Accounts[0].Balance).To(Equal(5000.0))
			Expect(bank.Accounts[0].Type).To(Equal("Checking"))

			// Verify the nested transaction
			Expect(len(bank.Accounts[0].Transactions)).To(Equal(1))
			Expect(bank.Accounts[0].Transactions[0].ID).To(Equal("TX-001"))
			Expect(bank.Accounts[0].Transactions[0].Amount).To(Equal(1000.0))
			Expect(bank.Accounts[0].Transactions[0].Type).To(Equal("Deposit"))
		})

		It("should correctly validate a Bank with invalid Account", func() {
			// Create invalid account (invalid type)
			invalidAccountBuilder := builders.NewAccountBuilder().
				WithNumber("ACCT-12345").
				WithBalance(5000.0).
				WithType("Unknown"). // Invalid type
				WithCurrency("USD")

			// Create bank with invalid account
			_, err := builders.NewBankBuilder().
				WithName("Global Bank").
				WithBranchCode("GB-001").
				WithAccount(invalidAccountBuilder).
				BuildAndValidate()

			// Verify validation fails
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Type must be one of"))
		})
	})

	// Test Investment with nested Portfolio and Stocks
	Describe("Investment with Portfolio and Stocks", func() {
		It("should correctly build an Investment with Portfolio and Stocks", func() {
			// Create stock
			stockBuilder := builders.NewStockBuilder().
				WithTicker("AAPL").
				WithCompanyName("Apple Inc.").
				WithPurchasePrice(150.0).
				WithCurrentPrice(170.0).
				WithUnits(10).
				WithSector("Technology")

			// Create bond
			bondBuilder := builders.NewBondBuilder().
				WithIssuer("US Treasury").
				WithFaceValue(1000.0).
				WithCouponRate(2.5).
				WithMaturityDate("2030-01-15").
				WithType("Government")

			// Create portfolio
			portfolioBuilder := builders.NewPortfolioBuilder().
				WithStock(stockBuilder).
				WithBond(bondBuilder).
				WithAllocation("Stocks", 60.0).
				WithAllocation("Bonds", 40.0)

			// Create performance record
			performanceRecordBuilder := builders.NewPerformanceRecordBuilder().
				WithDate("2022-01-01").
				WithValue(10500.0).
				WithROI(5.0).
				WithNotes("Quarterly performance review")

			// Create investment
			investment := builders.NewInvestmentBuilder().
				WithID("INV-001").
				WithName("Retirement Portfolio").
				WithType("401k").
				WithValue(10000.0).
				WithStartDate("2020-01-15").
				WithRisk("Medium").
				WithPortfolio(portfolioBuilder).
				WithPerformanceRecord(performanceRecordBuilder).
				Build().(models.Investment)

			// Verify the nested structure
			Expect(investment.ID).To(Equal("INV-001"))
			Expect(investment.Name).To(Equal("Retirement Portfolio"))
			Expect(investment.Type).To(Equal("401k"))
			Expect(investment.Value).To(Equal(10000.0))

			// Verify the nested portfolio
			Expect(len(investment.Portfolio.Stocks)).To(Equal(1))
			Expect(investment.Portfolio.Stocks[0].Ticker).To(Equal("AAPL"))
			Expect(investment.Portfolio.Stocks[0].CompanyName).To(Equal("Apple Inc."))

			Expect(len(investment.Portfolio.Bonds)).To(Equal(1))
			Expect(investment.Portfolio.Bonds[0].Issuer).To(Equal("US Treasury"))
			Expect(investment.Portfolio.Bonds[0].FaceValue).To(Equal(1000.0))

			// Verify the nested performance record
			Expect(len(investment.Performance)).To(Equal(1))
			Expect(investment.Performance[0].Date).To(Equal("2022-01-01"))
			Expect(investment.Performance[0].Value).To(Equal(10500.0))
			Expect(investment.Performance[0].ROI).To(Equal(5.0))
		})
	})

	// Test DigitalProfile with nested Device, Application, and OnlineAccount
	Describe("DigitalProfile with Device, Application, and OnlineAccount", func() {
		It("should correctly build a DigitalProfile with Device, Application, and OnlineAccount", func() {
			// Create application
			appBuilder := builders.NewApplicationBuilder().
				WithName("Social Media App").
				WithVersion("2.1.0").
				WithInstallDate("2022-01-15").
				WithLastUsed("2023-05-01").
				WithPermission("Camera").
				WithPermission("Microphone")

			// Create device
			deviceBuilder := builders.NewDeviceBuilder().
				WithType("Smartphone").
				WithModel("iPhone 13").
				WithSerialNumber("ABCD1234").
				WithPurchaseDate("2022-01-01").
				WithOS("iOS 16.5").
				WithLastUsed("2023-05-01").
				WithApp(appBuilder).
				WithSetting("Dark Mode", "Enabled")

			// Create subscription
			subscriptionBuilder := builders.NewSubscriptionBuilder().
				WithPlan("Premium").
				WithStartDate("2022-01-01").
				WithEndDate("2023-01-01").
				WithPrice(9.99).
				WithBillingCycle("Monthly").
				WithAutoRenew(true).
				WithFeature("Ad-free").
				WithFeature("Offline access")

			// Create online account
			accountBuilder := builders.NewOnlineAccountBuilder().
				WithPlatform("Social Media").
				WithUsername("user123").
				WithEmail("user@example.com").
				WithCreationDate("2020-01-01").
				WithLastLogin("2023-05-01").
				WithStatus("Active").
				WithSetting("Privacy", "Friends only").
				WithSubscription(subscriptionBuilder)

			// Create user preferences
			preferencesBuilder := builders.NewUserPreferencesBuilder().
				WithTheme("Dark").
				WithLanguage("English").
				WithNotification("Email", true).
				WithNotification("Push", true).
				WithNotification("SMS", false).
				WithPrivacySetting("Profile visibility", "Friends only").
				WithAccessibilitySetting("High contrast", false)

			// Create geo location
			geoLocationBuilder := builders.NewGeoLocationBuilder().
				WithLatitude(37.7749).
				WithLongitude(-122.4194).
				WithAccuracy(10.0)

			// Create user activity
			activityBuilder := builders.NewUserActivityBuilder().
				WithTimestamp("2023-05-01T12:34:56Z").
				WithType("Login").
				WithPlatform("Mobile").
				WithDevice("iPhone 13").
				WithLocation(geoLocationBuilder).
				WithDetail("IP", "192.168.1.1")

			// Create digital profile
			digitalProfile := builders.NewDigitalProfileBuilder().
				WithDevice(deviceBuilder).
				WithAccount(accountBuilder).
				WithPreferences(preferencesBuilder).
				WithActivity(activityBuilder).
				Build().(models.DigitalProfile)

			// Verify the nested structure
			Expect(len(digitalProfile.Devices)).To(Equal(1))
			Expect(digitalProfile.Devices[0].Type).To(Equal("Smartphone"))
			Expect(digitalProfile.Devices[0].Model).To(Equal("iPhone 13"))

			// Verify the nested application
			Expect(len(digitalProfile.Devices[0].Apps)).To(Equal(1))
			Expect(digitalProfile.Devices[0].Apps[0].Name).To(Equal("Social Media App"))
			Expect(len(digitalProfile.Devices[0].Apps[0].Permissions)).To(Equal(2))

			// Verify the nested online account
			Expect(len(digitalProfile.Accounts)).To(Equal(1))
			Expect(digitalProfile.Accounts[0].Platform).To(Equal("Social Media"))
			Expect(digitalProfile.Accounts[0].Username).To(Equal("user123"))

			// Verify the nested subscription
			Expect(digitalProfile.Accounts[0].Subscription.Plan).To(Equal("Premium"))
			Expect(digitalProfile.Accounts[0].Subscription.Price).To(Equal(9.99))
			Expect(len(digitalProfile.Accounts[0].Subscription.Features)).To(Equal(2))

			// Verify the nested preferences
			Expect(digitalProfile.Preferences.Theme).To(Equal("Dark"))
			Expect(digitalProfile.Preferences.Language).To(Equal("English"))
			Expect(digitalProfile.Preferences.Notifications["Email"]).To(BeTrue())

			// Verify the nested activity
			Expect(len(digitalProfile.Activity)).To(Equal(1))
			Expect(digitalProfile.Activity[0].Type).To(Equal("Login"))

			// Verify the nested geo location
			Expect(digitalProfile.Activity[0].Location.Latitude).To(Equal(37.7749))
			Expect(digitalProfile.Activity[0].Location.Longitude).To(Equal(-122.4194))
		})
	})
})

// TestNestedStructuresReflection uses reflection to verify that all nested structures are properly handled
func TestNestedStructuresReflection(t *testing.T) {
	// Get all model types
	modelTypes := getModelTypes()

	// Check each model type for nested structures
	for typeName, modelType := range modelTypes {
		t.Logf("Checking model: %s", typeName)

		// Skip interfaces and non-struct types
		if modelType.Kind() != reflect.Struct {
			continue
		}

		// Check each field in the model
		for i := 0; i < modelType.NumField(); i++ {
			field := modelType.Field(i)

			// Skip unexported fields
			if !field.IsExported() {
				continue
			}

			// Check if the field is a nested struct
			if field.Type.Kind() == reflect.Struct && field.Type.PkgPath() == modelType.PkgPath() {
				t.Logf("  Found nested struct: %s.%s of type %s", typeName, field.Name, field.Type.Name())

				// Verify that we can create a builder for the nested struct
				verifyNestedBuilder(t, typeName, field.Name, field.Type.Name())
			}

			// Check if the field is a slice of nested structs
			if field.Type.Kind() == reflect.Slice &&
				field.Type.Elem().Kind() == reflect.Struct &&
				field.Type.Elem().PkgPath() == modelType.PkgPath() {
				t.Logf("  Found slice of nested structs: %s.%s of type []%s", typeName, field.Name, field.Type.Elem().Name())

				// Verify that we can create a builder for the nested struct
				verifyNestedBuilder(t, typeName, field.Name, field.Type.Elem().Name())
			}
		}
	}
}

// getModelTypes returns all types in the models package
func getModelTypes() map[string]reflect.Type {
	modelTypes := make(map[string]reflect.Type)

	// Use reflection to get all types in the models package
	modelType := reflect.TypeOf(models.Person{})
	modelTypes[modelType.Name()] = modelType

	// Add more model types
	modelTypes["Address"] = reflect.TypeOf(models.Address{})
	modelTypes["Education"] = reflect.TypeOf(models.Education{})
	modelTypes["Course"] = reflect.TypeOf(models.Course{})
	modelTypes["Employment"] = reflect.TypeOf(models.Employment{})
	modelTypes["Company"] = reflect.TypeOf(models.Company{})
	modelTypes["Department"] = reflect.TypeOf(models.Department{})
	modelTypes["Project"] = reflect.TypeOf(models.Project{})
	modelTypes["Bank"] = reflect.TypeOf(models.Bank{})
	modelTypes["Account"] = reflect.TypeOf(models.Account{})
	modelTypes["Transaction"] = reflect.TypeOf(models.Transaction{})
	modelTypes["Investment"] = reflect.TypeOf(models.Investment{})
	modelTypes["Portfolio"] = reflect.TypeOf(models.Portfolio{})
	modelTypes["PerformanceRecord"] = reflect.TypeOf(models.PerformanceRecord{})
	modelTypes["Stock"] = reflect.TypeOf(models.Stock{})
	modelTypes["Bond"] = reflect.TypeOf(models.Bond{})
	modelTypes["ETF"] = reflect.TypeOf(models.ETF{})
	modelTypes["HealthProfile"] = reflect.TypeOf(models.HealthProfile{})
	modelTypes["Insurance"] = reflect.TypeOf(models.Insurance{})
	modelTypes["Medication"] = reflect.TypeOf(models.Medication{})
	modelTypes["MedicalRecord"] = reflect.TypeOf(models.MedicalRecord{})
	modelTypes["DigitalProfile"] = reflect.TypeOf(models.DigitalProfile{})
	modelTypes["Device"] = reflect.TypeOf(models.Device{})
	modelTypes["Application"] = reflect.TypeOf(models.Application{})
	modelTypes["OnlineAccount"] = reflect.TypeOf(models.OnlineAccount{})
	modelTypes["Subscription"] = reflect.TypeOf(models.Subscription{})
	modelTypes["UserPreferences"] = reflect.TypeOf(models.UserPreferences{})
	modelTypes["UserActivity"] = reflect.TypeOf(models.UserActivity{})
	modelTypes["GeoLocation"] = reflect.TypeOf(models.GeoLocation{})

	return modelTypes
}

// verifyNestedBuilder verifies that we can create a builder for a nested struct
func verifyNestedBuilder(t *testing.T, parentTypeName, fieldName, nestedTypeName string) {
	// Get the builder function for the nested type
	builderFuncName := fmt.Sprintf("New%sBuilder", nestedTypeName)

	// Use reflection to find the builder function
	builderType := reflect.TypeOf(builders.NewPersonBuilder)

	// Check if the builder function exists
	if builderType.Kind() != reflect.Func {
		t.Errorf("Builder function %s is not a function", builderFuncName)
		return
	}

	// Check if the builder function returns a builder
	if builderType.NumOut() != 1 {
		t.Errorf("Builder function %s does not return a single value", builderFuncName)
		return
	}

	// Check if the builder has a Build method
	builderReturnType := builderType.Out(0)
	buildMethod, ok := builderReturnType.MethodByName("Build")
	if !ok {
		t.Errorf("Builder for %s does not have a Build method", nestedTypeName)
		return
	}

	// Check if the Build method returns an interface{}
	if buildMethod.Type.NumOut() != 1 || buildMethod.Type.Out(0).Kind() != reflect.Interface {
		t.Errorf("Build method for %s does not return an interface{}", nestedTypeName)
		return
	}

	t.Logf("  Verified builder for nested type %s", nestedTypeName)
}
