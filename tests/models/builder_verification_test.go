// +build ignore

// +build ignore
// +build ignore

package models_test

import (
	"github.com/adil-faiyaz98/go-builder-kit/builders"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Builder Verification", func() {
	Context("Address Builder", func() {
		It("should build a valid Address with all fields", func() {
			// Create a GeoLocation builder
			geoLocationBuilder := builders.NewGeoLocationBuilder()
			geoLocationBuilder.
				WithLatitude(37.7749).
				WithLongitude(-122.4194).
				WithAccuracy(10.0)

			// Create an Address builder
			addressBuilder := builders.NewAddressBuilder()
			addressBuilder.
				WithStreet("123 Main St").
				WithCity("San Francisco").
				WithState("CA").
				WithPostalCode("94105").
				WithCountry("USA").
				WithCoordinates(geoLocationBuilder).
				WithType("Home").
				WithIsPrimary(true)

			// Build the Address
			address, err := addressBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(address).NotTo(BeNil())
			Expect(address.Street).To(Equal("123 Main St"))
			Expect(address.City).To(Equal("San Francisco"))
			Expect(address.State).To(Equal("CA"))
			Expect(address.PostalCode).To(Equal("94105"))
			Expect(address.Country).To(Equal("USA"))
			Expect(address.Type).To(Equal("Home"))
			Expect(address.IsPrimary).To(BeTrue())

			// Verify nested GeoLocation
			Expect(address.Coordinates).NotTo(BeNil())
			Expect(address.Coordinates.Latitude).To(Equal(37.7749))
			Expect(address.Coordinates.Longitude).To(Equal(-122.4194))
			Expect(address.Coordinates.Accuracy).To(Equal(10.0))
		})
	})

	Context("Education Builder", func() {
		It("should build a valid Education with all fields including nested structs", func() {
			// Create an Address builder for the location
			locationBuilder := builders.NewAddressBuilder()
			locationBuilder.
				WithStreet("456 University Ave").
				WithCity("Palo Alto").
				WithState("CA").
				WithPostalCode("94301").
				WithCountry("USA").
				WithType("Work")

			// Create a Course builder
			courseBuilder := builders.NewCourseBuilder()
			courseBuilder.
				WithCode("CS101").
				WithName("Introduction to Computer Science").
				WithDescription("Basic programming concepts").
				WithCredits(3.0).
				WithGrade("A").
				WithSemester("Fall").
				WithYear(2020).
				WithInstructor("Dr. Smith")

			// Create an Education builder
			educationBuilder := builders.NewEducationBuilder()
			educationBuilder.
				WithDegree("Bachelor of Science").
				WithInstitution("Stanford University").
				WithLocation(locationBuilder).
				WithStartDate("2018-09-01").
				WithEndDate("2022-06-15").
				WithGPA(3.8).
				WithHonor("Magna Cum Laude").
				WithMajor("Computer Science").
				WithMinor("Mathematics").
				WithCourse(courseBuilder).
				WithActivity("Chess Club")

			// Build the Education
			education, err := educationBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(education).NotTo(BeNil())
			Expect(education.Degree).To(Equal("Bachelor of Science"))
			Expect(education.Institution).To(Equal("Stanford University"))
			Expect(education.StartDate).To(Equal("2018-09-01"))
			Expect(education.EndDate).To(Equal("2022-06-15"))
			Expect(education.GPA).To(Equal(3.8))
			Expect(education.Major).To(Equal("Computer Science"))
			Expect(education.Minor).To(Equal("Mathematics"))
			Expect(education.Honors).To(ContainElement("Magna Cum Laude"))
			Expect(education.Activities).To(ContainElement("Chess Club"))

			// Verify nested Location
			Expect(education.Location).NotTo(BeNil())
			Expect(education.Location.Street).To(Equal("456 University Ave"))
			Expect(education.Location.City).To(Equal("Palo Alto"))
			Expect(education.Location.State).To(Equal("CA"))
			Expect(education.Location.PostalCode).To(Equal("94301"))
			Expect(education.Location.Country).To(Equal("USA"))
			Expect(education.Location.Type).To(Equal("Work"))

			// Verify nested Course
			Expect(education.Courses).To(HaveLen(1))
			course := education.Courses[0]
			Expect(course.Code).To(Equal("CS101"))
			Expect(course.Name).To(Equal("Introduction to Computer Science"))
			Expect(course.Description).To(Equal("Basic programming concepts"))
			Expect(course.Credits).To(Equal(3.0))
			Expect(course.Grade).To(Equal("A"))
			Expect(course.Semester).To(Equal("Fall"))
			Expect(course.Year).To(Equal(2020))
			Expect(course.Instructor).To(Equal("Dr. Smith"))
		})
	})

	Context("Person Builder", func() {
		It("should build a valid Person with all fields including nested structs", func() {
			// Create an Address builder
			addressBuilder := builders.NewAddressBuilder()
			addressBuilder.
				WithStreet("789 Oak St").
				WithCity("San Francisco").
				WithState("CA").
				WithPostalCode("94110").
				WithCountry("USA").
				WithType("Home").
				WithIsPrimary(true)

			// Create an Education builder
			educationBuilder := builders.NewEducationBuilder()
			educationBuilder.
				WithDegree("Master of Science").
				WithInstitution("UC Berkeley").
				WithStartDate("2022-09-01").
				WithEndDate("2024-06-15").
				WithGPA(3.9).
				WithMajor("Data Science")

			// Create a Person builder
			personBuilder := builders.NewPersonBuilder()
			personBuilder.
				WithID("P12345").
				WithName("John Doe").
				WithAge(28).
				WithEmail("john.doe@example.com").
				WithPhone("555-123-4567").
				WithBirthdate("1995-05-15").
				WithGender("male").
				WithNationality("American").
				WithMaritalStatus("single").
				WithAddress(addressBuilder).
				WithEducation(educationBuilder).
				WithCreatedAt("2023-01-01T12:00:00Z").
				WithUpdatedAt("2023-01-01T12:00:00Z")

			// Build the Person
			person, err := personBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(person).NotTo(BeNil())
			Expect(person.ID).To(Equal("P12345"))
			Expect(person.Name).To(Equal("John Doe"))
			Expect(person.Age).To(Equal(28))
			Expect(person.Email).To(Equal("john.doe@example.com"))
			Expect(person.Phone).To(Equal("555-123-4567"))
			Expect(person.Birthdate).To(Equal("1995-05-15"))
			Expect(person.Gender).To(Equal("male"))
			Expect(person.Nationality).To(Equal("American"))
			Expect(person.MaritalStatus).To(Equal("single"))
			Expect(person.CreatedAt).To(Equal("2023-01-01T12:00:00Z"))
			Expect(person.UpdatedAt).To(Equal("2023-01-01T12:00:00Z"))

			// Verify nested Address
			Expect(person.Address).NotTo(BeNil())
			Expect(person.Address.Street).To(Equal("789 Oak St"))
			Expect(person.Address.City).To(Equal("San Francisco"))
			Expect(person.Address.State).To(Equal("CA"))
			Expect(person.Address.PostalCode).To(Equal("94110"))
			Expect(person.Address.Country).To(Equal("USA"))
			Expect(person.Address.Type).To(Equal("Home"))
			Expect(person.Address.IsPrimary).To(BeTrue())

			// Verify nested Education
			Expect(person.Education).NotTo(BeNil())
			Expect(person.Education.Degree).To(Equal("Master of Science"))
			Expect(person.Education.Institution).To(Equal("UC Berkeley"))
			Expect(person.Education.StartDate).To(Equal("2022-09-01"))
			Expect(person.Education.EndDate).To(Equal("2024-06-15"))
			Expect(person.Education.GPA).To(Equal(3.9))
			Expect(person.Education.Major).To(Equal("Data Science"))
		})
	})

	Context("Department Builder", func() {
		It("should build a valid Department with all fields including nested structs", func() {
			// Create an Address builder for the location
			locationBuilder := builders.NewAddressBuilder()
			locationBuilder.
				WithStreet("101 Corporate Dr").
				WithCity("New York").
				WithState("NY").
				WithPostalCode("10001").
				WithCountry("USA").
				WithType("Work")

			// Create a Project builder
			projectBuilder := builders.NewProjectBuilder()
			projectBuilder.
				WithName("Digital Transformation").
				WithDescription("Modernize IT infrastructure").
				WithStartDate("2023-01-15").
				WithEndDate("2023-12-31").
				WithStatus("in-progress").
				WithBudget(500000.0)

			// Create a Department builder
			departmentBuilder := builders.NewDepartmentBuilder()
			departmentBuilder.
				WithName("Information Technology").
				WithCode("IT").
				WithDescription("Manages company IT infrastructure").
				WithBudget(1000000.0).
				WithLocation(locationBuilder).
				WithProject(projectBuilder)

			// Build the Department
			department, err := departmentBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(department).NotTo(BeNil())
			Expect(department.Name).To(Equal("Information Technology"))
			Expect(department.Code).To(Equal("IT"))
			Expect(department.Description).To(Equal("Manages company IT infrastructure"))
			Expect(department.Budget).To(Equal(1000000.0))

			// Verify nested Location
			Expect(department.Location).NotTo(BeNil())
			Expect(department.Location.Street).To(Equal("101 Corporate Dr"))
			Expect(department.Location.City).To(Equal("New York"))
			Expect(department.Location.State).To(Equal("NY"))
			Expect(department.Location.PostalCode).To(Equal("10001"))
			Expect(department.Location.Country).To(Equal("USA"))
			Expect(department.Location.Type).To(Equal("Work"))
		})
	})

	Context("Bond Builder", func() {
		It("should build a valid Bond with all fields", func() {
			// Create a Bond builder
			bondBuilder := builders.NewBondBuilder()
			bondBuilder.
				WithID("B12345").
				WithISIN("US123456AB12").
				WithName("US Treasury Bond").
				WithIssuer("US Treasury").
				WithType("government").
				WithFaceValue(1000.0).
				WithCouponRate(2.5).
				WithMaturityDate("2030-12-31").
				WithPurchaseDate("2023-01-15").
				WithPurchasePrice(980.0).
				WithCurrentPrice(985.0).
				WithQuantity(10).
				WithCurrency("USD").
				WithPaymentFrequency("semi-annual").
				WithRating("AAA").
				WithYield(2.7)

			// Build the Bond
			bond, err := bondBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(bond).NotTo(BeNil())
			Expect(bond.ID).To(Equal("B12345"))
			Expect(bond.ISIN).To(Equal("US123456AB12"))
			Expect(bond.Name).To(Equal("US Treasury Bond"))
			Expect(bond.Issuer).To(Equal("US Treasury"))
			Expect(bond.Type).To(Equal("government"))
			Expect(bond.FaceValue).To(Equal(1000.0))
			Expect(bond.CouponRate).To(Equal(2.5))
			Expect(bond.MaturityDate).To(Equal("2030-12-31"))
			Expect(bond.PurchaseDate).To(Equal("2023-01-15"))
			Expect(bond.PurchasePrice).To(Equal(980.0))
			Expect(bond.CurrentPrice).To(Equal(985.0))
			Expect(bond.Quantity).To(Equal(10))
			Expect(bond.Currency).To(Equal("USD"))
			Expect(bond.PaymentFrequency).To(Equal("semi-annual"))
			Expect(bond.Rating).To(Equal("AAA"))
			Expect(bond.Yield).To(Equal(2.7))
		})
	})

	Context("Portfolio Builder", func() {
		It("should build a valid Portfolio with all fields including nested structs", func() {
			// Create a Bond builder
			bondBuilder := builders.NewBondBuilder()
			bondBuilder.
				WithID("B12345").
				WithISIN("US123456AB12").
				WithName("US Treasury Bond").
				WithIssuer("US Treasury").
				WithType("government").
				WithFaceValue(1000.0).
				WithCouponRate(2.5).
				WithMaturityDate("2030-12-31").
				WithPurchaseDate("2023-01-15").
				WithPurchasePrice(980.0).
				WithCurrentPrice(985.0).
				WithQuantity(10).
				WithCurrency("USD").
				WithPaymentFrequency("semi-annual").
				WithRating("AAA").
				WithYield(2.7)

			// Create a Portfolio builder
			portfolioBuilder := builders.NewPortfolioBuilder()
			portfolioBuilder.
				WithID("P12345").
				WithName("Retirement Portfolio").
				WithDescription("Long-term investment portfolio").
				WithRiskLevel("medium").
				WithTotalValue(100000.0).
				WithAllocation("Bonds", 40.0).
				WithAllocation("Stocks", 50.0).
				WithAllocation("Cash", 10.0).
				WithBond(bondBuilder)

			// Build the Portfolio
			portfolio, err := portfolioBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(portfolio).NotTo(BeNil())
			Expect(portfolio.ID).To(Equal("P12345"))
			Expect(portfolio.Name).To(Equal("Retirement Portfolio"))
			Expect(portfolio.Description).To(Equal("Long-term investment portfolio"))
			Expect(portfolio.RiskLevel).To(Equal("medium"))
			Expect(portfolio.TotalValue).To(Equal(100000.0))

			// Verify Allocation
			Expect(portfolio.Allocation).To(HaveLen(3))
			Expect(portfolio.Allocation["Bonds"]).To(Equal(40.0))
			Expect(portfolio.Allocation["Stocks"]).To(Equal(50.0))
			Expect(portfolio.Allocation["Cash"]).To(Equal(10.0))

			// Verify nested Bond
			Expect(portfolio.Bonds).To(HaveLen(1))
			bond := portfolio.Bonds[0]
			Expect(bond.ID).To(Equal("B12345"))
			Expect(bond.ISIN).To(Equal("US123456AB12"))
			Expect(bond.Name).To(Equal("US Treasury Bond"))
			Expect(bond.Issuer).To(Equal("US Treasury"))
			Expect(bond.Type).To(Equal("government"))
			Expect(bond.FaceValue).To(Equal(1000.0))
			Expect(bond.CouponRate).To(Equal(2.5))
			Expect(bond.MaturityDate).To(Equal("2030-12-31"))
			Expect(bond.PurchaseDate).To(Equal("2023-01-15"))
			Expect(bond.PurchasePrice).To(Equal(980.0))
			Expect(bond.CurrentPrice).To(Equal(985.0))
			Expect(bond.Quantity).To(Equal(10))
			Expect(bond.Currency).To(Equal("USD"))
			Expect(bond.PaymentFrequency).To(Equal("semi-annual"))
			Expect(bond.Rating).To(Equal("AAA"))
			Expect(bond.Yield).To(Equal(2.7))
		})
	})
})
