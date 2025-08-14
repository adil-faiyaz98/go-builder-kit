package department_test

import (
	"github.com/adil-faiyaz98/go-builder-kit/v2/builders"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Department Builder", func() {

	Context("Basic Department Creation", func() {
		It("should create a department with basic fields", func() {
			// Create a department with basic fields
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
	})

	Context("Department with Location", func() {
		It("should create a department with a location", func() {
			// Create an address
			addressBuilder := builders.NewAddressBuilder().
				WithStreet("123 Main St").
				WithCity("San Francisco").
				WithState("CA").
				WithPostalCode("94105").
				WithCountry("USA")

			// Create a department with a location
			departmentBuilder := builders.NewDepartmentBuilder().
				WithName("Engineering").
				WithCode("ENG").
				WithLocation(addressBuilder)

			// Build the department
			department := departmentBuilder.BuildPtr()

			// Verify the department fields
			Expect(department.Name).To(Equal("Engineering"))
			Expect(department.Code).To(Equal("ENG"))

			// Verify the location
			Expect(department.Location).NotTo(BeNil())
			Expect(department.Location.Street).To(Equal("123 Main St"))
			Expect(department.Location.City).To(Equal("San Francisco"))
			Expect(department.Location.State).To(Equal("CA"))
			Expect(department.Location.PostalCode).To(Equal("94105"))
			Expect(department.Location.Country).To(Equal("USA"))
		})
	})

	Context("Department Types", func() {
		It("should create departments with different types", func() {
			// Create departments with different types
			engineeringBuilder := builders.NewDepartmentBuilder().
				WithName("Engineering").
				WithCode("ENG").
				WithBudget(1000000.0)

			marketingBuilder := builders.NewDepartmentBuilder().
				WithName("Marketing").
				WithCode("MKT").
				WithBudget(500000.0)

			financeBuilder := builders.NewDepartmentBuilder().
				WithName("Finance").
				WithCode("FIN").
				WithBudget(750000.0)

			// Build the departments
			engineering := engineeringBuilder.BuildPtr()
			marketing := marketingBuilder.BuildPtr()
			finance := financeBuilder.BuildPtr()

			// Verify the departments
			Expect(engineering.Name).To(Equal("Engineering"))
			Expect(engineering.Code).To(Equal("ENG"))
			Expect(engineering.Budget).To(Equal(1000000.0))

			Expect(marketing.Name).To(Equal("Marketing"))
			Expect(marketing.Code).To(Equal("MKT"))
			Expect(marketing.Budget).To(Equal(500000.0))

			Expect(finance.Name).To(Equal("Finance"))
			Expect(finance.Code).To(Equal("FIN"))
			Expect(finance.Budget).To(Equal(750000.0))
		})
	})

	Context("Builder Cloning", func() {
		It("should clone a department builder correctly", func() {
			// Create a base department builder
			baseBuilder := builders.NewDepartmentBuilder().
				WithName("Engineering").
				WithCode("ENG").
				WithDescription("Engineering department").
				WithBudget(1000000.0).
				WithHeadCount(50)

			// Clone the builder and modify it
			clonedBuilder := baseBuilder.Clone().
				WithName("Research & Development").
				WithCode("R&D").
				WithBudget(1500000.0)

			// Build both departments
			baseDepartment := baseBuilder.BuildPtr()
			clonedDepartment := clonedBuilder.BuildPtr()

			// Verify the base department wasn't affected by changes to the clone
			Expect(baseDepartment.Name).To(Equal("Engineering"))
			Expect(baseDepartment.Code).To(Equal("ENG"))
			Expect(baseDepartment.Budget).To(Equal(1000000.0))

			// Verify the cloned department has the new values
			Expect(clonedDepartment.Name).To(Equal("Research & Development"))
			Expect(clonedDepartment.Code).To(Equal("R&D"))
			Expect(clonedDepartment.Description).To(Equal("Engineering department")) // Unchanged
			Expect(clonedDepartment.Budget).To(Equal(1500000.0))
			Expect(clonedDepartment.HeadCount).To(Equal(50)) // Unchanged
		})
	})
})
