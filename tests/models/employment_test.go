package models_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

var _ = Describe("Employment", func() {
	var (
		employmentBuilder *builders.EmploymentBuilder
		companyBuilder    *builders.CompanyBuilder
		addressBuilder    *builders.AddressBuilder
		departmentBuilder *builders.DepartmentBuilder
		projectBuilder    *builders.ProjectBuilder
	)

	BeforeEach(func() {
		addressBuilder = builders.NewAddressBuilder().
			WithStreet("123 Corporate Dr").
			WithCity("San Francisco").
			WithState("CA").
			WithPostalCode("94105").
			WithCountry("USA").
			WithType("Business")

		projectBuilder = builders.NewProjectBuilder().
			WithName("Mobile App").
			WithDescription("Cross-platform mobile application").
			WithStartDate("2022-01-01").
			WithEndDate("2022-12-31").
			WithStatus("In Progress").
			WithBudget(500000.0)

		departmentBuilder = builders.NewDepartmentBuilder().
			WithName("Engineering").
			WithDescription("Software development department").
			WithHeadCount(50).
			WithBudget(5000000.0).
			WithProject(projectBuilder)

		companyBuilder = builders.NewCompanyBuilder().
			WithID("COMP-001").
			WithName("Tech Innovations Inc.").
			WithIndustry("Technology").
			WithDescription("Leading technology company").
			WithFounded("2010-01-01").
			WithLocation(addressBuilder).
			WithDepartment(departmentBuilder)

		employmentBuilder = builders.NewEmploymentBuilder().
			WithCompany(companyBuilder).
			WithPosition("Senior Software Engineer").
			WithDepartment("Engineering").
			WithStartDate("2020-01-15").
			WithSalary(120000.0).
			WithIsCurrent(true)
	})

	Describe("Positive Tests", func() {
		Context("when creating employment with valid data", func() {
			It("should create employment with basic information", func() {
				employment := employmentBuilder.Build().(*models.Employment)
				
				Expect(employment.Position).To(Equal("Senior Software Engineer"))
				Expect(employment.Department).To(Equal("Engineering"))
				Expect(employment.Salary).To(Equal(120000.0))
				Expect(employment.IsCurrent).To(BeTrue())
			})

			It("should create employment with company details", func() {
				employment := employmentBuilder.Build().(*models.Employment)
				
				Expect(employment.Company.Name).To(Equal("Tech Innovations Inc."))
				Expect(employment.Company.Industry).To(Equal("Technology"))
				Expect(employment.Company.Location.City).To(Equal("San Francisco"))
			})

			It("should create employment with department and project details", func() {
				employment := employmentBuilder.Build().(*models.Employment)
				
				Expect(len(employment.Company.Departments)).To(Equal(1))
				Expect(employment.Company.Departments[0].Name).To(Equal("Engineering"))
				Expect(len(employment.Company.Departments[0].Projects)).To(Equal(1))
				Expect(employment.Company.Departments[0].Projects[0].Name).To(Equal("Mobile App"))
			})

			It("should create employment with end date for past jobs", func() {
				employment := employmentBuilder.
					WithEndDate("2022-12-31").
					WithIsCurrent(false).
					Build().(models.Employment)
				
				Expect(employment.EndDate).To(Equal("2022-12-31"))
				Expect(employment.IsCurrent).To(BeFalse())
			})
		})
	})

	Describe("Negative Tests", func() {
		Context("when creating employment with invalid data", func() {
			It("should handle empty position", func() {
				employment := employmentBuilder.
					WithPosition("").
					Build().(models.Employment)
				
				// If we had validation, this would fail
				// For now, just check that the position is empty
				Expect(employment.Position).To(Equal(""))
			})

			It("should handle negative salary", func() {
				employment := employmentBuilder.
					WithSalary(-50000.0).
					Build().(models.Employment)
				
				// If we had validation, this would fail
				// For now, just check that the salary is negative
				Expect(employment.Salary).To(Equal(-50000.0))
			})

			It("should handle end date before start date", func() {
				employment := employmentBuilder.
					WithStartDate("2020-01-15").
					WithEndDate("2019-12-31").  // Before start date
					Build().(models.Employment)
				
				// If we had validation, this would fail
				// For now, just check the dates
				Expect(employment.StartDate).To(Equal("2020-01-15"))
				Expect(employment.EndDate).To(Equal("2019-12-31"))
			})

			It("should handle invalid company data", func() {
				invalidCompanyBuilder := builders.NewCompanyBuilder().
					WithID("").  // Invalid: empty ID
					WithName("").  // Invalid: empty name
					WithIndustry("")  // Invalid: empty industry

				employment := employmentBuilder.
					WithCompany(invalidCompanyBuilder).
					Build().(models.Employment)
				
				// If we had validation, this would fail
				// For now, just check the company data
				Expect(employment.Company.ID).To(Equal(""))
				Expect(employment.Company.Name).To(Equal(""))
				Expect(employment.Company.Industry).To(Equal(""))
			})
		})
	})

	Describe("Edge Cases", func() {
		It("should handle employment with no company", func() {
			employment := builders.NewEmploymentBuilder().
				WithPosition("Freelancer").
				WithStartDate("2020-01-01").
				WithIsCurrent(true).
				Build().(models.Employment)
			
			// Company will be zero value
			Expect(employment.Company.Name).To(Equal(""))
			Expect(employment.Company.ID).To(Equal(""))
		})

		It("should handle employment with multiple projects", func() {
			project1 := builders.NewProjectBuilder().
				WithName("Project 1").
				WithStatus("In Progress")

			project2 := builders.NewProjectBuilder().
				WithName("Project 2").
				WithStatus("Planning")

			departmentWithProjects := builders.NewDepartmentBuilder().
				WithName("Multi-Project Department").
				WithProject(project1).
				WithProject(project2)

			companyWithProjects := builders.NewCompanyBuilder().
				WithName("Multi-Project Company").
				WithDepartment(departmentWithProjects)

			employment := employmentBuilder.
				WithCompany(companyWithProjects).
				Build().(models.Employment)
			
			Expect(len(employment.Company.Departments[0].Projects)).To(Equal(2))
			Expect(employment.Company.Departments[0].Projects[0].Name).To(Equal("Project 1"))
			Expect(employment.Company.Departments[0].Projects[1].Name).To(Equal("Project 2"))
		})

		It("should handle employment with very high salary", func() {
			employment := employmentBuilder.
				WithSalary(10000000.0).  // $10 million
				Build().(models.Employment)
			
			Expect(employment.Salary).To(Equal(10000000.0))
		})

		It("should handle employment with zero salary", func() {
			employment := employmentBuilder.
				WithSalary(0.0).  // Volunteer or unpaid internship
				Build().(models.Employment)
			
			Expect(employment.Salary).To(Equal(0.0))
		})
	})
})

