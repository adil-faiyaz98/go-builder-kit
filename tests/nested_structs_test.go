// +build ignore

// +build ignore

import (
	"github.com/adil-faiyaz98/go-builder-kit/builders"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Nested Structs", func() {
	Context("Deep Nesting", func() {
		It("should handle deeply nested structs correctly", func() {
			// Create a GeoLocation builder
			geoLocationBuilder := builders.NewGeoLocationBuilder().
				WithLatitude(37.7749).
				WithLongitude(-122.4194).
				WithAccuracy(10.0)

			// Create an Address builder for the person
			personAddressBuilder := builders.NewAddressBuilder().
				WithStreet("123 Main St").
				WithCity("San Francisco").
				WithState("CA").
				WithPostalCode("94105").
				WithCountry("USA").
				WithCoordinates(geoLocationBuilder).
				WithType("Home").
				WithIsPrimary(true)

			// Create an Address builder for the company
			companyAddressBuilder := builders.NewAddressBuilder().
				WithStreet("456 Market St").
				WithCity("San Francisco").
				WithState("CA").
				WithPostalCode("94103").
				WithCountry("USA").
				WithType("Work")

			// Create a Company builder
			companyBuilder := builders.NewCompanyBuilder().
				WithName("Acme Inc").
				WithIndustry("Technology").
				WithDescription("Software company").
				WithFounded("2000-01-01").
				WithWebsite("https://acme.example.com").
				WithAddress(companyAddressBuilder).
				WithSize("Medium").
				WithRevenue("$10M-$50M").
				WithPublic(false)

			// Create an Employment builder
			employmentBuilder := builders.NewEmploymentBuilder().
				WithPosition("Senior Software Engineer").
				WithCompany(companyBuilder).
				WithStartDate("2020-01-15").
				WithIsCurrent(true).
				WithSalary(120000.0).
				WithDepartment("Engineering")

			// Create a Course builder
			courseBuilder := builders.NewCourseBuilder().
				WithCode("CS101").
				WithName("Introduction to Computer Science").
				WithDescription("Basic programming concepts").
				WithCredits(3.0).
				WithGrade("A").
				WithSemester("Fall").
				WithYear(2015).
				WithInstructor("Dr. Smith")

			// Create an Education builder
			educationBuilder := builders.NewEducationBuilder().
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
			personBuilder := builders.NewPersonBuilder().
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

	Context("Validation", func() {
		It("should validate nested structures", func() {
			// Create an Address builder with invalid coordinates
			addressBuilder := builders.NewAddressBuilder().
				WithStreet("123 Main St").
				WithCity("San Francisco").
				WithState("CA").
				WithPostalCode("94105").
				WithCountry("USA").
				WithCoordinates(
					builders.NewGeoLocationBuilder().
						WithLatitude(100.0). // Invalid: latitude must be between -90 and 90
						WithLongitude(-122.4194).
						WithAccuracy(10.0),
				)

			// Create a Person builder with the invalid address
			personBuilder := builders.NewPersonBuilder().
				WithID("P12345").
				WithName("John Doe").
				WithEmail("john.doe@example.com").
				WithAddress(addressBuilder)

			// Build and validate should fail
			_, err := personBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Latitude must be between -90 and 90"))
		})

		It("should validate all nested levels", func() {
			// Create a Course builder with invalid data
			courseBuilder := builders.NewCourseBuilder().
				WithCode(""). // Invalid: empty code
				WithName("Introduction to Computer Science").
				WithCredits(-1.0) // Invalid: negative credits

			// Create an Education builder with the invalid course
			educationBuilder := builders.NewEducationBuilder().
				WithDegree("Bachelor of Science").
				WithInstitution("Stanford University").
				WithCourse(courseBuilder)

			// Create a Person builder with the invalid education
			personBuilder := builders.NewPersonBuilder().
				WithID("P12345").
				WithName("John Doe").
				WithEmail("john.doe@example.com").
				WithEducation(educationBuilder)

			// Build and validate should fail
			_, err := personBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Code cannot be empty"))
			Expect(err.Error()).To(ContainSubstring("Credits cannot be negative"))
		})
	})

	Context("Builder Cloning", func() {
		It("should clone builders with nested structures", func() {
			// Create a base person builder with address
			baseBuilder := builders.NewPersonBuilder().
				WithID("P12345").
				WithName("John Doe").
				WithEmail("john.doe@example.com").
				WithAddress(
					builders.NewAddressBuilder().
						WithStreet("123 Main St").
						WithCity("San Francisco").
						WithCountry("USA"),
				)

			// Clone the builder and modify it
			clonedBuilder := baseBuilder.Clone().
				WithName("Jane Smith").
				WithEmail("jane.smith@example.com")

			// Build both persons
			basePerson := baseBuilder.BuildPtr()
			clonedPerson := clonedBuilder.BuildPtr()

			// Verify the base person
			Expect(basePerson.Name).To(Equal("John Doe"))
			Expect(basePerson.Email).To(Equal("john.doe@example.com"))

			// Verify the cloned person
			Expect(clonedPerson.Name).To(Equal("Jane Smith"))
			Expect(clonedPerson.Email).To(Equal("jane.smith@example.com"))

			// Verify that the ID is the same
			Expect(clonedPerson.ID).To(Equal(basePerson.ID))

			// Verify that the address is cloned
			Expect(clonedPerson.Address).NotTo(BeNil())
			Expect(clonedPerson.Address.Street).To(Equal("123 Main St"))
			Expect(clonedPerson.Address.City).To(Equal("San Francisco"))
			Expect(clonedPerson.Address.Country).To(Equal("USA"))

			// Modify the cloned address
			clonedBuilder.WithAddress(
				builders.NewAddressBuilder().
					WithStreet("456 Oak St").
					WithCity("New York").
					WithCountry("USA"),
			)

			// Build the cloned person again
			clonedPerson = clonedBuilder.BuildPtr()

			// Verify that the base person's address is unchanged
			Expect(basePerson.Address.Street).To(Equal("123 Main St"))
			Expect(basePerson.Address.City).To(Equal("San Francisco"))

			// Verify that the cloned person's address is changed
			Expect(clonedPerson.Address.Street).To(Equal("456 Oak St"))
			Expect(clonedPerson.Address.City).To(Equal("New York"))
		})
	})
})
