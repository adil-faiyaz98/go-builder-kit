package person_test

import (
	"github.com/adil-faiyaz98/go-builder-kit/builders"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Person Builder", func() {
	var personBuilder *builders.PersonBuilder

	BeforeEach(func() {
		// Initialize a new builder before each test
		personBuilder = builders.NewPersonBuilder()
	})

	Context("Basic Person Creation", func() {
		It("should create a person with basic fields", func() {
			// Configure the builder
			personBuilder.
				WithName("John Doe").
				WithEmail("john.doe@example.com").
				WithAge(30).
				WithPhone("555-1234")

			// Build the person
			person := personBuilder.BuildPtr()

			// Verify the person fields
			Expect(person.Name).To(Equal("John Doe"))
			Expect(person.Email).To(Equal("john.doe@example.com"))
			Expect(person.Age).To(Equal(30))
			Expect(person.Phone).To(Equal("555-1234"))
		})
	})

	Context("Person with Address", func() {
		var addressBuilder *builders.AddressBuilder

		BeforeEach(func() {
			// Initialize an address builder before each test
			addressBuilder = builders.NewAddressBuilder().
				WithStreet("123 Main St").
				WithCity("San Francisco").
				WithState("CA").
				WithPostalCode("94105").
				WithCountry("USA").
				WithType("Home")

			// Configure the person builder with basic info
			personBuilder.
				WithName("John Doe").
				WithEmail("john.doe@example.com")
		})

		It("should create a person with an address", func() {
			// Add the address to the person
			personBuilder.WithAddress(addressBuilder)

			// Build the person
			person := personBuilder.BuildPtr()

			// Verify the person fields
			Expect(person.Name).To(Equal("John Doe"))
			Expect(person.Email).To(Equal("john.doe@example.com"))

			// Verify the address fields
			Expect(person.Address).NotTo(BeNil())
			Expect(person.Address.Street).To(Equal("123 Main St"))
			Expect(person.Address.City).To(Equal("San Francisco"))
			Expect(person.Address.State).To(Equal("CA"))
			Expect(person.Address.PostalCode).To(Equal("94105"))
			Expect(person.Address.Country).To(Equal("USA"))
			Expect(person.Address.Type).To(Equal("Home"))
		})
	})

	Describe("Person with Education", func() {
		It("should create a person with education", func() {
			// Create an education
			educationBuilder := builders.NewEducationBuilder().
				WithDegree("Bachelor of Science").
				WithMajor("Computer Science").
				WithInstitution("Stanford University").
				WithStartDate("2016-09-01").
				WithEndDate("2020-06-01")

			// Create a person with education
			personBuilder := builders.NewPersonBuilder().
				WithName("John Doe").
				WithEmail("john.doe@example.com").
				WithEducation(educationBuilder)

			// Build the person
			person := personBuilder.BuildPtr()

			// Verify the person fields
			Expect(person.Name).To(Equal("John Doe"))
			Expect(person.Email).To(Equal("john.doe@example.com"))

			// Verify the education fields
			Expect(person.Education).NotTo(BeNil())
			Expect(person.Education.Degree).To(Equal("Bachelor of Science"))
			Expect(person.Education.Major).To(Equal("Computer Science"))
			Expect(person.Education.Institution).To(Equal("Stanford University"))
			Expect(person.Education.StartDate).To(Equal("2016-09-01"))
			Expect(person.Education.EndDate).To(Equal("2020-06-01"))
		})
	})

	Describe("Builder Cloning", func() {
		It("should clone a person builder correctly", func() {
			// Create a base person builder
			baseBuilder := builders.NewPersonBuilder().
				WithName("John Doe").
				WithEmail("john.doe@example.com").
				WithAge(30)

			// Clone the builder and modify it
			clonedBuilder := baseBuilder.Clone().
				WithName("Jane Doe").
				WithAge(28)

			// Build both persons
			basePerson := baseBuilder.BuildPtr()
			clonedPerson := clonedBuilder.BuildPtr()

			// Verify the base person wasn't affected by changes to the clone
			Expect(basePerson.Name).To(Equal("John Doe"))
			Expect(basePerson.Email).To(Equal("john.doe@example.com"))
			Expect(basePerson.Age).To(Equal(30))

			// Verify the cloned person has the new values
			Expect(clonedPerson.Name).To(Equal("Jane Doe"))
			Expect(clonedPerson.Email).To(Equal("john.doe@example.com")) // Unchanged
			Expect(clonedPerson.Age).To(Equal(28))
		})
	})
})
