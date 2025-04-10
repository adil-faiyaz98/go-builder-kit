// +build ignore

// +build ignore
// +build ignore

package models_test

import (
	"github.com/adil-faiyaz98/go-builder-kit/builders"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Nested Structs", func() {
	Context("Deep Nesting", func() {
		It("should handle deeply nested structs correctly", func() {
			// Create a GeoLocation builder
			geoLocationBuilder := builders.NewGeoLocationBuilder()
			geoLocationBuilder.
				WithLatitude(37.7749).
				WithLongitude(-122.4194).
				WithAccuracy(10.0)

			// Create an Address builder for the person
			personAddressBuilder := builders.NewAddressBuilder()
			personAddressBuilder.
				WithStreet("123 Main St").
				WithCity("San Francisco").
				WithState("CA").
				WithPostalCode("94105").
				WithCountry("USA").
				WithCoordinates(geoLocationBuilder).
				WithType("Home").
				WithIsPrimary(true)

			// Create an Address builder for the company
			companyAddressBuilder := builders.NewAddressBuilder()
			companyAddressBuilder.
				WithStreet("456 Market St").
				WithCity("San Francisco").
				WithState("CA").
				WithPostalCode("94103").
				WithCountry("USA").
				WithType("Work")

			// Create a Company builder
			companyBuilder := builders.NewCompanyBuilder()
			companyBuilder.
				WithName("Acme Inc").
				WithIndustry("Technology").
				WithDescription("Software company").
				WithFounded("2000-01-01").
				WithWebsite("https://acme.example.com").
				WithAddress(companyAddressBuilder).
				WithSize("Medium").
				WithRevenue("$10M-$50M").
				WithPublic(false)

			// Create a Task builder
			taskBuilder := builders.NewTaskBuilder()
			taskBuilder.
				WithName("Implement Feature X").
				WithDescription("Add new functionality").
				WithStartDate("2023-01-15").
				WithEndDate("2023-02-15").
				WithStatus("in-progress").
				WithPriority("high")

			// Create a Project builder
			projectBuilder := builders.NewProjectBuilder()
			projectBuilder.
				WithName("Product Redesign").
				WithDescription("Redesign the product UI/UX").
				WithStartDate("2023-01-01").
				WithEndDate("2023-06-30").
				WithStatus("in-progress").
				WithBudget(200000.0).
				WithTask(taskBuilder)

			// Create an Employment builder
			employmentBuilder := builders.NewEmploymentBuilder()
			employmentBuilder.
				WithPosition("Senior Software Engineer").
				WithCompany(companyBuilder).
				WithStartDate("2020-01-15").
				WithIsCurrent(true).
				WithSalary(120000.0).
				WithDepartment("Engineering")

			// Create a Course builder
			courseBuilder := builders.NewCourseBuilder()
			courseBuilder.
				WithCode("CS101").
				WithName("Introduction to Computer Science").
				WithDescription("Basic programming concepts").
				WithCredits(3.0).
				WithGrade("A").
				WithSemester("Fall").
				WithYear(2015).
				WithInstructor("Dr. Smith")

			// Create an Education builder
			educationBuilder := builders.NewEducationBuilder()
			educationBuilder.
				WithDegree("Bachelor of Science").
				WithInstitution("Stanford University").
				WithStartDate("2012-09-01").
				WithEndDate("2016-06-15").
				WithGPA(3.8).
				WithHonor("Magna Cum Laude").
				WithMajor("Computer Science").
				WithMinor("Mathematics").
				WithCourse(courseBuilder).
				WithActivity("Chess Club")

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
				WithAddress(personAddressBuilder).
				WithEducation(educationBuilder).
				WithEmployment(employmentBuilder).
				WithCreatedAt("2023-01-01T12:00:00Z").
				WithUpdatedAt("2023-01-01T12:00:00Z")

			// Build the Person
			person, err := personBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(person).NotTo(BeNil())

			// Verify Person fields
			Expect(person.Name).To(Equal("John Doe"))
			Expect(person.Email).To(Equal("john.doe@example.com"))

			// Verify Address fields
			Expect(person.Address).NotTo(BeNil())
			Expect(person.Address.Street).To(Equal("123 Main St"))
			Expect(person.Address.City).To(Equal("San Francisco"))

			// Verify GeoLocation fields
			Expect(person.Address.Coordinates).NotTo(BeNil())
			Expect(person.Address.Coordinates.Latitude).To(Equal(37.7749))
			Expect(person.Address.Coordinates.Longitude).To(Equal(-122.4194))

			// Verify Education fields
			Expect(person.Education).NotTo(BeNil())
			Expect(person.Education.Degree).To(Equal("Bachelor of Science"))
			Expect(person.Education.Institution).To(Equal("Stanford University"))

			// Verify Course fields
			Expect(person.Education.Courses).To(HaveLen(1))
			Expect(person.Education.Courses[0].Name).To(Equal("Introduction to Computer Science"))
			Expect(person.Education.Courses[0].Grade).To(Equal("A"))

			// Verify Employment fields
			Expect(person.Employment).NotTo(BeNil())
			Expect(person.Employment.Position).To(Equal("Senior Software Engineer"))
			Expect(person.Employment.Salary).To(Equal(120000.0))

			// Verify Company fields
			Expect(person.Employment.Company).NotTo(BeNil())
			Expect(person.Employment.Company.Name).To(Equal("Acme Inc"))
			Expect(person.Employment.Company.Industry).To(Equal("Technology"))

			// Verify Company Address fields
			Expect(person.Employment.Company.Address).NotTo(BeNil())
			Expect(person.Employment.Company.Address.Street).To(Equal("456 Market St"))
			Expect(person.Employment.Company.Address.City).To(Equal("San Francisco"))
		})
	})
})
