// +build ignore

// +build ignore

import (
	"github.com/adil-faiyaz98/go-builder-kit/builders"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Company", func() {
	Context("Company Builder", func() {
		It("should build a valid Company with all fields", func() {
			// Create an Address builder for the company
			addressBuilder := builders.NewAddressBuilder().
				WithStreet("123 Tech St").
				WithCity("San Francisco").
				WithState("CA").
				WithPostalCode("94105").
				WithCountry("USA").
				WithType("Work")

			// Create a Company builder
			companyBuilder := builders.NewCompanyBuilder().
				WithName("Acme Inc").
				WithIndustry("Technology").
				WithDescription("Software company").
				WithFounded("2000-01-01").
				WithWebsite("https://acme.example.com").
				WithAddress(addressBuilder).
				WithSize("Medium").
				WithRevenue("$10M-$50M").
				WithPublic(false)

			// Build the Company
			company, err := companyBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(company).NotTo(BeNil())
			Expect(company.Name).To(Equal("Acme Inc"))
			Expect(company.Industry).To(Equal("Technology"))
			Expect(company.Description).To(Equal("Software company"))
			Expect(company.Founded).To(Equal("2000-01-01"))
			Expect(company.Website).To(Equal("https://acme.example.com"))
			Expect(company.Size).To(Equal("Medium"))
			Expect(company.Revenue).To(Equal("$10M-$50M"))
			Expect(company.Public).To(BeFalse())

			// Verify Address
			Expect(company.Address).NotTo(BeNil())
			Expect(company.Address.Street).To(Equal("123 Tech St"))
			Expect(company.Address.City).To(Equal("San Francisco"))
		})

		It("should validate required fields", func() {
			// Create a Company builder with missing required fields
			companyBuilder := builders.NewCompanyBuilder()
			// Missing Name and Industry

			// Build and validate should fail
			_, err := companyBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Name cannot be empty"))
			Expect(err.Error()).To(ContainSubstring("Industry cannot be empty"))
		})

		It("should validate date format", func() {
			// Create a Company builder with invalid date
			companyBuilder := builders.NewCompanyBuilder().
				WithName("Acme Inc").
				WithIndustry("Technology").
				WithFounded("01/01/2000") // Wrong format, should be YYYY-MM-DD

			// Build and validate should fail
			_, err := companyBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Founded must be in the format YYYY-MM-DD"))
		})

		It("should validate website format", func() {
			// Create a Company builder with invalid website
			companyBuilder := builders.NewCompanyBuilder().
				WithName("Acme Inc").
				WithIndustry("Technology").
				WithWebsite("acme.example.com") // Missing http:// or https://

			// Build and validate should fail
			_, err := companyBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Website must start with http:// or https://"))
		})

		It("should validate nested structures", func() {
			// Create an Address builder with invalid data
			addressBuilder := builders.NewAddressBuilder().
				WithStreet(""). // Missing required field
				WithCity("San Francisco").
				WithCountry("USA").
				WithType("Work")

			// Create a Company builder with the invalid address
			companyBuilder := builders.NewCompanyBuilder().
				WithName("Acme Inc").
				WithIndustry("Technology").
				WithAddress(addressBuilder)

			// Build and validate should fail
			_, err := companyBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Street cannot be empty"))
		})

		It("should clone the builder correctly", func() {
			// Create a base company builder
			baseBuilder := builders.NewCompanyBuilder().
				WithName("Acme Inc").
				WithIndustry("Technology").
				WithFounded("2000-01-01").
				WithWebsite("https://acme.example.com")

			// Clone the builder and modify it
			clonedBuilder := baseBuilder.Clone().
				WithName("Beta Corp").
				WithIndustry("Finance")

			// Build both companies
			baseCompany := baseBuilder.BuildPtr()
			clonedCompany := clonedBuilder.BuildPtr()

			// Verify the base company
			Expect(baseCompany.Name).To(Equal("Acme Inc"))
			Expect(baseCompany.Industry).To(Equal("Technology"))

			// Verify the cloned company
			Expect(clonedCompany.Name).To(Equal("Beta Corp"))
			Expect(clonedCompany.Industry).To(Equal("Finance"))

			// Verify that the founded date and website are the same
			Expect(clonedCompany.Founded).To(Equal(baseCompany.Founded))
			Expect(clonedCompany.Website).To(Equal(baseCompany.Website))
		})
	})
})
