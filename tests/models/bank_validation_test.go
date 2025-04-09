package models_test

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

var _ = Describe("Bank Validation", func() {
	var (
		bankBuilder    *builders.BankBuilder
		addressBuilder *builders.AddressBuilder
		accountBuilder *builders.AccountBuilder
	)

	BeforeEach(func() {
		addressBuilder = builders.NewAddressBuilder().
			WithStreet("100 Financial St").
			WithCity("New York").
			WithState("NY").
			WithPostalCode("10005").
			WithCountry("USA").
			WithType("Work")

		accountBuilder = builders.NewAccountBuilder().
			WithNumber("ACCT-12345").
			WithBalance(5000.0).
			WithType("Checking").
			WithCurrency("USD").
			WithStatus("Active")

		bankBuilder = builders.NewBankBuilder().
			WithName("Global Bank").
			WithBranchCode("GB-001").
			WithAddress(addressBuilder).
			WithAccount(accountBuilder).
			WithRelationshipStartDate("2020-01-15")
	})

	Describe("Validation", func() {
		Context("when all fields are valid", func() {
			It("should pass validation", func() {
				bank, err := bankBuilder.BuildWithValidation()
				Expect(err).NotTo(HaveOccurred())
				Expect(bank).NotTo(BeNil())
				Expect(bank.Name).To(Equal("Global Bank"))
			})
		})

		Context("when required fields are missing", func() {
			It("should fail validation when Name is empty", func() {
				_, err := bankBuilder.
					WithName("").
					BuildWithValidation()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Name cannot be empty"))
			})

			It("should fail validation when BranchCode is empty", func() {
				_, err := bankBuilder.
					WithBranchCode("").
					BuildWithValidation()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("BranchCode cannot be empty"))
			})
		})

		Context("when field values are invalid", func() {
			It("should fail validation when BranchCode format is invalid", func() {
				_, err := bankBuilder.
					WithBranchCode("123").
					BuildWithValidation()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("BranchCode must be in the format XX-000"))
			})

			It("should fail validation when RelationshipStartDate format is invalid", func() {
				_, err := bankBuilder.
					WithRelationshipStartDate("01/15/2020").
					BuildWithValidation()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("RelationshipStartDate must be in the format YYYY-MM-DD"))
			})
		})

		Context("when nested objects are invalid", func() {
			It("should fail validation when Address is invalid", func() {
				invalidAddressBuilder := builders.NewAddressBuilder().
					WithStreet(""). // Invalid: empty street
					WithCity("New York").
					WithCountry("USA")

				_, err := bankBuilder.
					WithAddress(invalidAddressBuilder).
					BuildWithValidation()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Street cannot be empty"))
			})

			It("should fail validation when Account is invalid", func() {
				invalidAccountBuilder := builders.NewAccountBuilder().
					WithNumber(""). // Invalid: empty number
					WithBalance(5000.0).
					WithType("Checking")

				_, err := bankBuilder.
					WithAccount(invalidAccountBuilder).
					BuildWithValidation()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Number cannot be empty"))
			})

			It("should fail validation when Account type is invalid", func() {
				invalidAccountBuilder := builders.NewAccountBuilder().
					WithNumber("ACCT-12345").
					WithBalance(5000.0).
					WithType("Unknown") // Invalid type

				_, err := bankBuilder.
					WithAccount(invalidAccountBuilder).
					BuildWithValidation()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Type must be one of"))
			})

			It("should fail validation when Account currency is invalid", func() {
				invalidAccountBuilder := builders.NewAccountBuilder().
					WithNumber("ACCT-12345").
					WithBalance(5000.0).
					WithType("Checking").
					WithCurrency("US") // Invalid: should be 3 letters

				_, err := bankBuilder.
					WithAccount(invalidAccountBuilder).
					BuildWithValidation()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Currency must be a 3-letter code"))
			})
		})
	})

	// Performance tests
	Describe("Performance", func() {
		Context("when creating many banks", func() {
			It("should handle creating 1000 banks efficiently", func() {
				startTime := time.Now()

				// Create 1000 banks
				for i := 1; i <= 1000; i++ {
					bank := builders.NewBankBuilder().
						WithName(fmt.Sprintf("Bank %d", i)).
						WithBranchCode(fmt.Sprintf("BB-%03d", i%1000)).
						Build().(models.Bank)

					// Just to ensure the compiler doesn't optimize away the creation
					Expect(bank.Name).To(Equal(fmt.Sprintf("Bank %d", i)))
				}

				duration := time.Since(startTime)

				// This is a performance test, so we're just checking that it completes in a reasonable time
				// Adjust the threshold based on your machine's performance
				Expect(duration.Seconds()).To(BeNumerically("<", 1.0), "Creating 1000 banks should take less than 1 second")
			})

			It("should handle creating and validating 100 banks efficiently", func() {
				startTime := time.Now()

				// Create and validate 100 banks
				for i := 1; i <= 100; i++ {
					_, err := builders.NewBankBuilder().
						WithName(fmt.Sprintf("Bank %d", i)).
						WithBranchCode(fmt.Sprintf("BB-%03d", i%1000)).
						WithRelationshipStartDate("2020-01-15").
						BuildWithValidation()

					Expect(err).NotTo(HaveOccurred())
				}

				duration := time.Since(startTime)

				// This is a performance test, so we're just checking that it completes in a reasonable time
				// Adjust the threshold based on your machine's performance
				Expect(duration.Seconds()).To(BeNumerically("<", 0.5), "Creating and validating 100 banks should take less than 0.5 seconds")
			})

			It("should handle creating a bank with many accounts efficiently", func() {
				startTime := time.Now()

				manyAccountsBuilder := builders.NewBankBuilder().
					WithName("Many Accounts Bank").
					WithBranchCode("MA-001").
					WithAddress(addressBuilder)

				// Add 100 accounts
				for i := 1; i <= 100; i++ {
					accountBuilder := builders.NewAccountBuilder().
						WithNumber(fmt.Sprintf("ACCT-%d", i)).
						WithBalance(float64(i * 1000)).
						WithType("Checking").
						WithCurrency("USD").
						WithStatus("Active")

					// Add 10 transactions to each account
					for j := 1; j <= 10; j++ {
						transactionBuilder := builders.NewTransactionBuilder().
							WithID(fmt.Sprintf("TX-%d-%d", i, j)).
							WithDate("2023-01-15").
							WithAmount(float64(j * 100)).
							WithType("Deposit")

						accountBuilder.WithTransaction(transactionBuilder)
					}

					manyAccountsBuilder.WithAccount(accountBuilder)
				}

				bank, err := manyAccountsBuilder.BuildWithValidation()

				duration := time.Since(startTime)

				Expect(err).NotTo(HaveOccurred())
				Expect(bank).NotTo(BeNil())
				Expect(len(bank.Accounts)).To(Equal(100))
				Expect(len(bank.Accounts[0].Transactions)).To(Equal(10))

				// This is a performance test, so we're just checking that it completes in a reasonable time
				// Adjust the threshold based on your machine's performance
				Expect(duration.Seconds()).To(BeNumerically("<", 0.5), "Creating a bank with many accounts should take less than 0.5 seconds")
			})
		})
	})
})

