package person_test

import (
	"github.com/adil-faiyaz98/go-builder-kit/builders"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Comprehensive Person Scenario", func() {
	// Shared builders that will be used across tests
	var (
		// Person and related builders
		personBuilder    *builders.PersonBuilder
		addressBuilder   *builders.AddressBuilder
		educationBuilder *builders.EducationBuilder

		// Employment related builders
		employmentBuilder *builders.EmploymentBuilder

		// Financial related builders
		bankBuilder *builders.BankBuilder
	)

	BeforeEach(func() {
		// Initialize all builders before each test

		// Create address
		addressBuilder = builders.NewAddressBuilder().
			WithStreet("123 Main St").
			WithCity("San Francisco").
			WithState("CA").
			WithPostalCode("94105").
			WithCountry("USA").
			WithType("Home")

		// Create geo location for the address
		geoLocationBuilder := builders.NewGeoLocationBuilder().
			WithLatitude(37.7749).
			WithLongitude(-122.4194)
		addressBuilder.WithCoordinates(geoLocationBuilder)

		// Create education
		educationBuilder = builders.NewEducationBuilder().
			WithDegree("Master of Science").
			WithMajor("Computer Science").
			WithInstitution("Stanford University").
			WithStartDate("2018-09-01").
			WithEndDate("2020-06-01")

		// Create company for employment
		companyBuilder := builders.NewCompanyBuilder().
			WithName("Acme Inc").
			WithIndustry("Technology")

		// Create employment
		employmentBuilder = builders.NewEmploymentBuilder().
			WithCompany(companyBuilder).
			WithPosition("Senior Software Engineer").
			WithDepartment("Engineering").
			WithStartDate("2020-07-01").
			WithSalary(150000.0).
			WithIsCurrent(true)

		// Create bank
		bankBuilder = builders.NewBankBuilder().
			WithName("First National Bank").
			WithBranchCode("SF001")

		// Create account
		accountBuilder := builders.NewAccountBuilder().
			WithType("Checking").
			WithNumber("1234567890").
			WithBalance(25000.0).
			WithCurrency("USD").
			WithStatus("Active")

		// Add account to bank
		bankBuilder.AddAccount(accountBuilder)

		// Create person
		personBuilder = builders.NewPersonBuilder().
			WithName("John Doe").
			WithEmail("john.doe@example.com").
			WithAge(32).
			WithPhone("555-1234").
			WithGender("Male").
			WithNationality("USA").
			WithMaritalStatus("Married").
			WithAddress(addressBuilder).
			WithEducation(educationBuilder).
			WithEmployment(employmentBuilder).
			WithBank(bankBuilder)
	})

	Context("Complete Person Profile", func() {
		It("should create a person with all related information", func() {
			// Build the person
			person := personBuilder.BuildPtr()

			// Verify basic person information
			Expect(person.Name).To(Equal("John Doe"))
			Expect(person.Email).To(Equal("john.doe@example.com"))
			Expect(person.Age).To(Equal(32))
			Expect(person.Phone).To(Equal("555-1234"))
			Expect(person.Gender).To(Equal("Male"))
			Expect(person.Nationality).To(Equal("USA"))
			Expect(person.MaritalStatus).To(Equal("Married"))

			// Verify address
			Expect(person.Address).NotTo(BeNil())
			Expect(person.Address.Street).To(Equal("123 Main St"))
			Expect(person.Address.City).To(Equal("San Francisco"))
			Expect(person.Address.State).To(Equal("CA"))
			Expect(person.Address.PostalCode).To(Equal("94105"))
			Expect(person.Address.Country).To(Equal("USA"))

			// Verify geo location
			Expect(person.Address.Coordinates).NotTo(BeNil())
			Expect(person.Address.Coordinates.Latitude).To(Equal(37.7749))
			Expect(person.Address.Coordinates.Longitude).To(Equal(-122.4194))

			// Verify education
			Expect(person.Education).NotTo(BeNil())
			Expect(person.Education.Degree).To(Equal("Master of Science"))
			Expect(person.Education.Major).To(Equal("Computer Science"))
			Expect(person.Education.Institution).To(Equal("Stanford University"))
			Expect(person.Education.StartDate).To(Equal("2018-09-01"))
			Expect(person.Education.EndDate).To(Equal("2020-06-01"))

			// Verify employment
			Expect(person.Employment).NotTo(BeNil())
			Expect(person.Employment.Company).NotTo(BeNil())
			Expect(person.Employment.Company.Name).To(Equal("Acme Inc"))
			Expect(person.Employment.Position).To(Equal("Senior Software Engineer"))
			Expect(person.Employment.StartDate).To(Equal("2020-07-01"))
			Expect(person.Employment.Salary).To(Equal(150000.0))

			// Verify department
			Expect(person.Employment.Department).To(Equal("Engineering"))

			// Verify bank
			Expect(person.Bank).NotTo(BeNil())
			Expect(person.Bank.Name).To(Equal("First National Bank"))
			Expect(person.Bank.BranchCode).To(Equal("SF001"))

			// Verify accounts
			Expect(person.Bank.Accounts).To(HaveLen(1))
			account := person.Bank.Accounts[0]
			Expect(account.Type).To(Equal("Checking"))
			Expect(account.Number).To(Equal("1234567890"))
			Expect(account.Balance).To(Equal(25000.0))
			Expect(account.Currency).To(Equal("USD"))
			Expect(account.Status).To(Equal("Active"))
		})
	})

})
