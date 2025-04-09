package tests

import (
	"github.com/adil-faiyaz98/go-builder-kit/builders"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bank", func() {
	Context("Bank Builder", func() {
		It("should build a valid Bank with all fields", func() {
			// Create an Address builder for the bank
			addressBuilder := builders.NewAddressBuilder().
				WithStreet("123 Wall St").
				WithCity("New York").
				WithState("NY").
				WithPostalCode("10005").
				WithCountry("USA").
				WithType("Work")

			// Create an Account builder
			accountBuilder := builders.NewAccountBuilder().
				WithID("A12345").
				WithType("Checking").
				WithNumber("123456789").
				WithBalance(1000.0).
				WithCurrency("USD").
				WithStatus("active")

			// Create a Bank builder
			bankBuilder := builders.NewBankBuilder().
				WithName("Global Bank").
				WithBranchCode("NYC001").
				WithAddress(addressBuilder).
				WithAccount(accountBuilder).
				WithRelationshipStartDate("2020-01-01")

			// Build the Bank
			bank, err := bankBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(bank).NotTo(BeNil())
			Expect(bank.Name).To(Equal("Global Bank"))
			Expect(bank.BranchCode).To(Equal("NYC001"))
			Expect(bank.RelationshipStartDate).To(Equal("2020-01-01"))

			// Verify Address
			Expect(bank.Address).NotTo(BeNil())
			Expect(bank.Address.Street).To(Equal("123 Wall St"))
			Expect(bank.Address.City).To(Equal("New York"))

			// Verify Accounts
			Expect(bank.Accounts).To(HaveLen(1))
			Expect(bank.Accounts[0].ID).To(Equal("A12345"))
			Expect(bank.Accounts[0].Type).To(Equal("Checking"))
		})

		It("should validate required fields", func() {
			// Create a Bank builder with missing required fields
			bankBuilder := builders.NewBankBuilder()
			// Missing Name and BranchCode

			// Build and validate should fail
			_, err := bankBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Name is required"))
			Expect(err.Error()).To(ContainSubstring("BranchCode is required"))
		})

		It("should validate date format", func() {
			// Create a Bank builder with invalid date
			bankBuilder := builders.NewBankBuilder().
				WithName("Global Bank").
				WithBranchCode("NYC001").
				WithRelationshipStartDate("01/01/2020") // Wrong format, should be YYYY-MM-DD

			// Build and validate should fail
			_, err := bankBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("RelationshipStartDate must be in the format YYYY-MM-DD"))
		})

		It("should validate nested structures", func() {
			// Create an Address builder with invalid data
			addressBuilder := builders.NewAddressBuilder().
				WithStreet(""). // Missing required field
				WithCity("New York").
				WithCountry("USA").
				WithType("Work")

			// Create a Bank builder with the invalid address
			bankBuilder := builders.NewBankBuilder().
				WithName("Global Bank").
				WithBranchCode("NYC001").
				WithAddress(addressBuilder)

			// Build and validate should fail
			_, err := bankBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Street cannot be empty"))
		})

		It("should clone the builder correctly", func() {
			// Create a base bank builder
			baseBuilder := builders.NewBankBuilder().
				WithName("Global Bank").
				WithBranchCode("NYC001").
				WithRelationshipStartDate("2020-01-01")

			// Clone the builder and modify it
			clonedBuilder := baseBuilder.Clone().
				WithName("Local Bank").
				WithBranchCode("LA001")

			// Build both banks
			baseBank := baseBuilder.BuildPtr()
			clonedBank := clonedBuilder.BuildPtr()

			// Verify the base bank
			Expect(baseBank.Name).To(Equal("Global Bank"))
			Expect(baseBank.BranchCode).To(Equal("NYC001"))

			// Verify the cloned bank
			Expect(clonedBank.Name).To(Equal("Local Bank"))
			Expect(clonedBank.BranchCode).To(Equal("LA001"))

			// Verify that the relationship start date is the same
			Expect(clonedBank.RelationshipStartDate).To(Equal(baseBank.RelationshipStartDate))
		})
	})
})
