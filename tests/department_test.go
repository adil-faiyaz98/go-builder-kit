// +build ignore

// +build ignore

import (
	"github.com/adil-faiyaz98/go-builder-kit/builders"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Department", func() {
	Context("Department Builder", func() {
		It("should build a valid Department with all fields", func() {
			// Create an Address builder for the department
			addressBuilder := builders.NewAddressBuilder().
				WithStreet("123 Office St").
				WithCity("San Francisco").
				WithState("CA").
				WithPostalCode("94105").
				WithCountry("USA").
				WithType("Work")

			// Create a Department builder
			departmentBuilder := builders.NewDepartmentBuilder().
				WithName("Engineering").
				WithCode("ENG").
				WithDescription("Software Engineering Department").
				WithBudget(1000000.0).
				WithHeadCount(50).
				WithLocation(addressBuilder)

			// Build the Department
			department, err := departmentBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(department).NotTo(BeNil())
			Expect(department.Name).To(Equal("Engineering"))
			Expect(department.Code).To(Equal("ENG"))
			Expect(department.Description).To(Equal("Software Engineering Department"))
			Expect(department.Budget).To(Equal(1000000.0))
			Expect(department.HeadCount).To(Equal(50))

			// Verify Location
			Expect(department.Location).NotTo(BeNil())
			Expect(department.Location.Street).To(Equal("123 Office St"))
			Expect(department.Location.City).To(Equal("San Francisco"))
		})

		It("should validate required fields", func() {
			// Create a Department builder with missing required fields
			departmentBuilder := builders.NewDepartmentBuilder()
			// Missing Name and Code

			// Build and validate should fail
			_, err := departmentBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Name cannot be empty"))
			Expect(err.Error()).To(ContainSubstring("Code cannot be empty"))
		})

		It("should validate budget is positive", func() {
			// Create a Department builder with negative budget
			departmentBuilder := builders.NewDepartmentBuilder().
				WithName("Engineering").
				WithCode("ENG").
				WithBudget(-1000.0) // Negative budget

			// Build and validate should fail
			_, err := departmentBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Budget cannot be negative"))
		})

		It("should validate nested structures", func() {
			// Create an Address builder with invalid data
			addressBuilder := builders.NewAddressBuilder().
				WithStreet(""). // Missing required field
				WithCity("San Francisco").
				WithCountry("USA").
				WithType("Work")

			// Create a Department builder with the invalid address
			departmentBuilder := builders.NewDepartmentBuilder().
				WithName("Engineering").
				WithCode("ENG").
				WithLocation(addressBuilder)

			// Build and validate should fail
			_, err := departmentBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Street cannot be empty"))
		})

		It("should clone the builder correctly", func() {
			// Create a base department builder
			baseBuilder := builders.NewDepartmentBuilder().
				WithName("Engineering").
				WithCode("ENG").
				WithBudget(1000000.0).
				WithHeadCount(50)

			// Clone the builder and modify it
			clonedBuilder := baseBuilder.Clone().
				WithName("Marketing").
				WithCode("MKT").
				WithBudget(500000.0)

			// Build both departments
			baseDepartment := baseBuilder.BuildPtr()
			clonedDepartment := clonedBuilder.BuildPtr()

			// Verify the base department
			Expect(baseDepartment.Name).To(Equal("Engineering"))
			Expect(baseDepartment.Code).To(Equal("ENG"))
			Expect(baseDepartment.Budget).To(Equal(1000000.0))

			// Verify the cloned department
			Expect(clonedDepartment.Name).To(Equal("Marketing"))
			Expect(clonedDepartment.Code).To(Equal("MKT"))
			Expect(clonedDepartment.Budget).To(Equal(500000.0))

			// Verify that the head count is the same
			Expect(clonedDepartment.HeadCount).To(Equal(baseDepartment.HeadCount))
		})
	})
})
