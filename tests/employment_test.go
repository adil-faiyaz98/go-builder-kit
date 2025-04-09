package tests

import (
	"github.com/adil-faiyaz98/go-builder-kit/builders"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Employment", func() {
	Context("Employment Builder", func() {
		It("should build a valid Employment with all fields", func() {
			// Create a Company builder
			companyBuilder := builders.NewCompanyBuilder().
				WithName("Acme Inc").
				WithIndustry("Technology")

			// Create an Employment builder
			employmentBuilder := builders.NewEmploymentBuilder().
				WithPosition("Senior Software Engineer").
				WithCompany(companyBuilder).
				WithDepartment("Engineering").
				WithStartDate("2020-01-15").
				WithEndDate("2023-06-30").
				WithIsCurrent(false).
				WithSalary(120000.0).
				WithSupervisor(nil).
				WithSubordinate(nil)

			// Build the Employment
			employment, err := employmentBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(employment).NotTo(BeNil())
			Expect(employment.Position).To(Equal("Senior Software Engineer"))
			Expect(employment.Department).To(Equal("Engineering"))
			Expect(employment.StartDate).To(Equal("2020-01-15"))
			Expect(employment.EndDate).To(Equal("2023-06-30"))
			Expect(employment.IsCurrent).To(BeFalse())
			Expect(employment.Salary).To(Equal(120000.0))

			// Verify Company
			Expect(employment.Company).NotTo(BeNil())
			Expect(employment.Company.Name).To(Equal("Acme Inc"))
			Expect(employment.Company.Industry).To(Equal("Technology"))
		})

		It("should validate required fields", func() {
			// Create an Employment builder with missing required fields
			employmentBuilder := builders.NewEmploymentBuilder()
			// Missing Position and Company

			// Build and validate should fail
			_, err := employmentBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Position cannot be empty"))
			Expect(err.Error()).To(ContainSubstring("Department cannot be empty"))
		})

		It("should validate date format", func() {
			// Create a Company builder
			companyBuilder := builders.NewCompanyBuilder().
				WithName("Acme Inc").
				WithIndustry("Technology")

			// Create an Employment builder with invalid date
			employmentBuilder := builders.NewEmploymentBuilder().
				WithPosition("Senior Software Engineer").
				WithCompany(companyBuilder).
				WithStartDate("01/15/2020") // Wrong format, should be YYYY-MM-DD

			// Build and validate should fail
			_, err := employmentBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("StartDate must be in the format YYYY-MM-DD"))
		})

		It("should validate end date is after start date", func() {
			// Create a Company builder
			companyBuilder := builders.NewCompanyBuilder().
				WithName("Acme Inc").
				WithIndustry("Technology")

			// Create an Employment builder with end date before start date
			employmentBuilder := builders.NewEmploymentBuilder().
				WithPosition("Senior Software Engineer").
				WithCompany(companyBuilder).
				WithStartDate("2023-01-15").
				WithEndDate("2020-06-30") // Before start date

			// Build and validate should fail
			_, err := employmentBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("EndDate cannot be before StartDate"))
		})

		It("should validate current job has no end date", func() {
			// Create a Company builder
			companyBuilder := builders.NewCompanyBuilder().
				WithName("Acme Inc").
				WithIndustry("Technology")

			// Create an Employment builder with current job and end date
			employmentBuilder := builders.NewEmploymentBuilder().
				WithPosition("Senior Software Engineer").
				WithCompany(companyBuilder).
				WithDepartment("Engineering").
				WithStartDate("2020-01-15").
				WithEndDate("2023-06-30").
				WithIsCurrent(true) // Current job should not have end date

			// Build and validate should fail
			_, err := employmentBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("EndDate cannot be set for current job"))
		})

		It("should validate salary is positive", func() {
			// Create a Company builder
			companyBuilder := builders.NewCompanyBuilder().
				WithName("Acme Inc").
				WithIndustry("Technology")

			// Create an Employment builder with negative salary
			employmentBuilder := builders.NewEmploymentBuilder().
				WithPosition("Senior Software Engineer").
				WithCompany(companyBuilder).
				WithSalary(-1000.0) // Negative salary

			// Build and validate should fail
			_, err := employmentBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Salary cannot be negative"))
		})

		It("should validate nested structures", func() {
			// Create a Company builder with invalid data
			companyBuilder := builders.NewCompanyBuilder().
				WithName(""). // Missing required field
				WithIndustry("Technology")

			// Create an Employment builder with the invalid company
			employmentBuilder := builders.NewEmploymentBuilder().
				WithPosition("Senior Software Engineer").
				WithCompany(companyBuilder)

			// Build and validate should fail
			_, err := employmentBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Name cannot be empty"))
		})

		It("should clone the builder correctly", func() {
			// Create a Company builder
			companyBuilder := builders.NewCompanyBuilder().
				WithName("Acme Inc").
				WithIndustry("Technology")

			// Create a base employment builder
			baseBuilder := builders.NewEmploymentBuilder().
				WithPosition("Senior Software Engineer").
				WithCompany(companyBuilder).
				WithStartDate("2020-01-15").
				WithIsCurrent(true).
				WithSalary(120000.0)

			// Clone the builder and modify it
			clonedBuilder := baseBuilder.Clone().
				WithPosition("Principal Engineer").
				WithSalary(150000.0)

			// Build both employments
			baseEmployment := baseBuilder.BuildPtr()
			clonedEmployment := clonedBuilder.BuildPtr()

			// Verify the base employment
			Expect(baseEmployment.Position).To(Equal("Senior Software Engineer"))
			Expect(baseEmployment.Salary).To(Equal(120000.0))

			// Verify the cloned employment
			Expect(clonedEmployment.Position).To(Equal("Principal Engineer"))
			Expect(clonedEmployment.Salary).To(Equal(150000.0))

			// Verify that the start date and current status are the same
			Expect(clonedEmployment.StartDate).To(Equal(baseEmployment.StartDate))
			Expect(clonedEmployment.IsCurrent).To(Equal(baseEmployment.IsCurrent))
		})
	})
})
