package builders_test

import (
	"github.com/adil-faiyaz98/go-builder-kit/builders"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// TestBuilders is not needed as we already have TestBuildersPackage
// func TestBuilders(t *testing.T) {
// 	RegisterFailHandler(Fail)
// 	RunSpecs(t, "Builders Suite")
// }

var _ = Describe("Builders", func() {
	Describe("Project Builder", func() {
		It("should handle string date fields correctly", func() {
			// Create date strings
			startDate := "2023-01-01"
			endDate := "2023-12-31"

			// Create a project with date fields
			projectBuilder := builders.NewProjectBuilder().
				WithName("Test Project").
				WithDescription("A test project").
				WithStartDate(startDate).
				WithEndDate(endDate)

			// Build the project
			project := projectBuilder.BuildPtr()

			// Verify the date fields
			Expect(project.StartDate).To(Equal(startDate))
			Expect(project.EndDate).To(Equal(endDate))
		})

		It("should handle string enum types correctly", func() {
			// Create a project with string enum fields
			projectBuilder := builders.NewProjectBuilder().
				WithName("Test Project").
				WithDescription("A test project").
				WithStatus("Active") // Status is a string enum

			// Build the project
			project := projectBuilder.BuildPtr()

			// Verify the string enum field
			Expect(project.Status).To(Equal("Active"))
		})

		It("should handle task relationships correctly", func() {
			// Create a task
			taskBuilder := builders.NewTaskBuilder().
				WithName("Task 1").
				WithDescription("First task").
				WithStatus("In Progress")

			// Create a project with a task
			projectBuilder := builders.NewProjectBuilder().
				WithName("Test Project").
				WithDescription("A test project")

			// Add the task to the project
			projectBuilder.AddTask(taskBuilder)

			// Build the project
			project := projectBuilder.BuildPtr()

			// Verify the task
			Expect(project.Tasks).To(HaveLen(1))
			Expect(project.Tasks[0].Name).To(Equal("Task 1"))
			Expect(project.Tasks[0].Description).To(Equal("First task"))
			Expect(project.Tasks[0].Status).To(Equal("In Progress"))
		})
	})

	Describe("Address Builder", func() {
		It("should build an Address with all fields", func() {
			// Create an address
			addressBuilder := builders.NewAddressBuilder().
				WithStreet("123 Main St").
				WithCity("San Francisco").
				WithState("CA").
				WithPostalCode("94105").
				WithCountry("USA").
				WithType("Home")

			// Build the address
			address := addressBuilder.BuildPtr()

			// Verify the address fields
			Expect(address.Street).To(Equal("123 Main St"))
			Expect(address.City).To(Equal("San Francisco"))
			Expect(address.State).To(Equal("CA"))
			Expect(address.PostalCode).To(Equal("94105"))
			Expect(address.Country).To(Equal("USA"))
			Expect(address.Type).To(Equal("Home"))
		})

		It("should handle nested GeoLocation", func() {
			// Create a GeoLocation
			geoLocationBuilder := builders.NewGeoLocationBuilder().
				WithLatitude(37.7749).
				WithLongitude(-122.4194)

			// Create an address with a GeoLocation
			addressBuilder := builders.NewAddressBuilder().
				WithStreet("123 Main St").
				WithCity("San Francisco").
				WithCoordinates(geoLocationBuilder)

			// Build the address
			address := addressBuilder.BuildPtr()

			// Verify the GeoLocation
			Expect(address.Coordinates).NotTo(BeNil())
			Expect(address.Coordinates.Latitude).To(Equal(37.7749))
			Expect(address.Coordinates.Longitude).To(Equal(-122.4194))
		})
	})

	Describe("Department Builder", func() {
		It("should build a Department with all fields", func() {
			// Create a department
			departmentBuilder := builders.NewDepartmentBuilder().
				WithName("Engineering").
				WithCode("ENG").
				WithDescription("Engineering department").
				WithBudget(1000000.0).
				WithHeadCount(50)

			// Build the department
			department := departmentBuilder.BuildPtr()

			// Verify the department fields
			Expect(department.Name).To(Equal("Engineering"))
			Expect(department.Code).To(Equal("ENG"))
			Expect(department.Description).To(Equal("Engineering department"))
			Expect(department.Budget).To(Equal(1000000.0))
			Expect(department.HeadCount).To(Equal(50))
		})

		It("should handle nested Address", func() {
			// Create an address
			addressBuilder := builders.NewAddressBuilder().
				WithStreet("123 Main St").
				WithCity("San Francisco")

			// Create a department with an address
			departmentBuilder := builders.NewDepartmentBuilder().
				WithName("Engineering").
				WithLocation(addressBuilder)

			// Build the department
			department := departmentBuilder.BuildPtr()

			// Verify the address
			Expect(department.Location).NotTo(BeNil())
			Expect(department.Location.Street).To(Equal("123 Main St"))
			Expect(department.Location.City).To(Equal("San Francisco"))
		})
	})

	Describe("Builder Cloning", func() {
		It("should clone builders correctly", func() {
			// Create a base builder
			baseBuilder := builders.NewProjectBuilder().
				WithName("Base Project").
				WithDescription("Base description").
				WithStatus("Planning")

			// Clone the builder and modify it
			clonedBuilder := baseBuilder.Clone().
				WithName("Cloned Project").
				WithStatus("Active")

			// Build both objects
			baseProject := baseBuilder.BuildPtr()
			clonedProject := clonedBuilder.BuildPtr()

			// Verify the base object wasn't affected by changes to the clone
			Expect(baseProject.Name).To(Equal("Base Project"))
			Expect(baseProject.Description).To(Equal("Base description"))
			Expect(baseProject.Status).To(Equal("Planning"))

			// Verify the cloned object has the new values
			Expect(clonedProject.Name).To(Equal("Cloned Project"))
			Expect(clonedProject.Description).To(Equal("Base description")) // Unchanged
			Expect(clonedProject.Status).To(Equal("Active"))
		})
	})

})
