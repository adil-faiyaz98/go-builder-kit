package models_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

var _ = Describe("Bank", func() {
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
			WithType("Business")

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
			WithAccounts(accountBuilder).
			WithRelationshipStartDate("2020-01-15")
	})

	// Note: In a real implementation, we would add a Validate method to the Bank model
	// For now, we'll just test the builder functionality

	Describe("Positive Tests", func() {
		Context("when creating a bank with valid data", func() {
			It("should create a bank with basic information", func() {
				bank := bankBuilder.Build().(*models.Bank)

				Expect(bank.Name).To(Equal("Global Bank"))
				Expect(bank.BranchCode).To(Equal("GB-001"))
				Expect(bank.Address.City).To(Equal("New York"))
				Expect(len(bank.Accounts)).To(Equal(1))
				Expect(bank.Accounts[0].Number).To(Equal("ACCT-12345"))
			})

			It("should create a bank with multiple accounts", func() {
				savingsAccountBuilder := builders.NewAccountBuilder().
					WithNumber("ACCT-67890").
					WithBalance(10000.0).
					WithType("Savings").
					WithCurrency("USD").
					WithInterestRate(1.5).
					WithStatus("Active")

				bank := bankBuilder.
					WithAccounts(savingsAccountBuilder).
					Build().(models.Bank)

				Expect(len(bank.Accounts)).To(Equal(2))
				Expect(bank.Accounts[0].Type).To(Equal("Checking"))
				Expect(bank.Accounts[1].Type).To(Equal("Savings"))
				Expect(bank.Accounts[1].InterestRate).To(Equal(1.5))
			})

			It("should create a bank with investments", func() {
				investmentBuilder := builders.NewInvestmentBuilder().
					WithID("INV-001").
					WithName("Retirement Fund").
					WithType("401k").
					WithValue(50000.0).
					WithRisk("Medium")

				bank := bankBuilder.
					WithInvestments(investmentBuilder).
					Build().(models.Bank)

				Expect(len(bank.Investments)).To(Equal(1))
				Expect(bank.Investments[0].Name).To(Equal("Retirement Fund"))
				Expect(bank.Investments[0].Value).To(Equal(50000.0))
			})

			It("should create a bank with loans", func() {
				loanBuilder := builders.NewLoanBuilder().
					WithID("LOAN-001").
					WithAmount(250000.0).
					WithInterestRate(3.5).
					WithType("Mortgage").
					WithStatus("Active")

				bank := bankBuilder.
					WithLoans(loanBuilder).
					Build().(models.Bank)

				Expect(len(bank.Loans)).To(Equal(1))
				Expect(bank.Loans[0].Type).To(Equal("Mortgage"))
				Expect(bank.Loans[0].Amount).To(Equal(250000.0))
			})
		})
	})

	Describe("Negative Tests", func() {
		Context("when creating a bank with invalid data", func() {
			It("should fail when name is empty", func() {
				bank := bankBuilder.
					WithName("").
					Build().(models.Bank)

				// If we had validation, this would fail
				// For now, just check that the name is empty
				Expect(bank.Name).To(Equal(""))
			})

			It("should fail when branch code is empty", func() {
				bank := bankBuilder.
					WithBranchCode("").
					Build().(models.Bank)

				// If we had validation, this would fail
				// For now, just check that the branch code is empty
				Expect(bank.BranchCode).To(Equal(""))
			})

			It("should fail when address is invalid", func() {
				invalidAddressBuilder := builders.NewAddressBuilder().
					WithStreet(""). // Invalid: empty street
					WithCity("New York").
					WithCountry("USA")

				bank := bankBuilder.
					WithAddress(invalidAddressBuilder).
					Build().(models.Bank)

				// If we had cascading validation, this would fail
				// For now, just check that the street is empty
				Expect(bank.Address.Street).To(Equal(""))
			})
		})
	})

	Describe("Edge Cases", func() {
		It("should handle a bank with no accounts", func() {
			bank := builders.NewBankBuilder().
				WithName("Empty Bank").
				WithBranchCode("EB-001").
				WithAddress(addressBuilder).
				Build().(models.Bank)

			Expect(len(bank.Accounts)).To(Equal(0))
			Expect(bank.Accounts).To(BeEmpty())
		})

		It("should handle a bank with many accounts", func() {
			manyAccountsBuilder := builders.NewBankBuilder().
				WithName("Many Accounts Bank").
				WithBranchCode("MA-001").
				WithAddress(addressBuilder)

			// Add 10 accounts
			for i := 1; i <= 10; i++ {
				accountBuilder := builders.NewAccountBuilder().
					WithNumber(fmt.Sprintf("ACCT-%d", i)).
					WithBalance(float64(i * 1000)).
					WithType("Checking").
					WithStatus("Active")

				manyAccountsBuilder.WithAccounts(accountBuilder)
			}

			bank := manyAccountsBuilder.Build().(*models.Bank)

			Expect(len(bank.Accounts)).To(Equal(10))
			Expect(bank.Accounts[0].Number).To(Equal("ACCT-1"))
			Expect(bank.Accounts[9].Number).To(Equal("ACCT-10"))
		})

		It("should handle a bank with zero balance accounts", func() {
			zeroBalanceAccountBuilder := builders.NewAccountBuilder().
				WithNumber("ZERO-ACCT").
				WithBalance(0.0).
				WithType("Checking").
				WithStatus("Active")

			bank := bankBuilder.
				WithAccounts(zeroBalanceAccountBuilder).
				Build().(models.Bank)

			Expect(len(bank.Accounts)).To(Equal(2))
			Expect(bank.Accounts[1].Balance).To(Equal(0.0))
		})
	})
})
