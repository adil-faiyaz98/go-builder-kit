package tests

import (
	"github.com/adil-faiyaz98/go-builder-kit/builders"
	"github.com/adil-faiyaz98/go-builder-kit/models"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Account", func() {
	Context("Account Builder", func() {
		It("should build a valid Account with all fields", func() {
			// Create an Account builder
			accountBuilder := builders.NewAccountBuilder().
				WithID("A12345").
				WithType("Checking").
				WithNumber("123456789").
				WithBalance(1000.0).
				WithCurrency("USD").
				WithOpenDate("2023-01-01").
				WithStatus("active").
				WithInterestRate(0.01).
				WithIsJoint(true).
				WithOverdraftLimit(500.0)

			// Build the Account
			account, err := accountBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(account).NotTo(BeNil())
			Expect(account.ID).To(Equal("A12345"))
			Expect(account.Type).To(Equal("Checking"))
			Expect(account.Number).To(Equal("123456789"))
			Expect(account.Balance).To(Equal(1000.0))
			Expect(account.Currency).To(Equal("USD"))
			Expect(account.OpenDate).To(Equal("2023-01-01"))
			Expect(account.Status).To(Equal("active"))
			Expect(account.InterestRate).To(Equal(0.01))
			Expect(account.IsJoint).To(BeTrue())
			Expect(account.OverdraftLimit).To(Equal(500.0))
		})

		It("should validate required fields", func() {
			// Create an Account builder with missing required fields
			accountBuilder := builders.NewAccountBuilder()
			// Missing ID, Type, and Number

			// Build and validate should fail
			_, err := accountBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("ID is required"))
			Expect(err.Error()).To(ContainSubstring("Type is required"))
			Expect(err.Error()).To(ContainSubstring("Number is required"))
		})

		It("should validate date format", func() {
			// Create an Account builder with invalid date
			accountBuilder := builders.NewAccountBuilder().
				WithID("A12345").
				WithType("Checking").
				WithNumber("123456789").
				WithOpenDate("01/01/2023") // Wrong format, should be YYYY-MM-DD

			// Build and validate should fail
			_, err := accountBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("OpenDate must be in the format YYYY-MM-DD"))
		})

		It("should support custom validation", func() {
			// Create an Account builder with custom validation
			accountBuilder := builders.NewAccountBuilder().
				WithID("A12345").
				WithType("Checking").
				WithNumber("123456789").
				WithBalance(-100.0). // Negative balance
				WithCurrency("USD").
				WithStatus("active").
				WithValidation(func(a *models.Account) error {
					if a.Balance < 0 {
						return nil // Allow negative balance in custom validation
					}
					return nil
				})

			// Build and validate should succeed despite negative balance
			account, err := accountBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(account).NotTo(BeNil())
			Expect(account.Balance).To(Equal(-100.0))
		})

		It("should clone the builder correctly", func() {
			// Create a base account builder
			baseBuilder := builders.NewAccountBuilder().
				WithID("A12345").
				WithType("Checking").
				WithNumber("123456789").
				WithBalance(1000.0).
				WithCurrency("USD")

			// Clone the builder and modify it
			clonedBuilder := baseBuilder.Clone().
				WithID("A67890").
				WithBalance(2000.0)

			// Build both accounts
			baseAccount := baseBuilder.BuildPtr()
			clonedAccount := clonedBuilder.BuildPtr()

			// Verify the base account
			Expect(baseAccount.ID).To(Equal("A12345"))
			Expect(baseAccount.Balance).To(Equal(1000.0))

			// Verify the cloned account
			Expect(clonedAccount.ID).To(Equal("A67890"))
			Expect(clonedAccount.Balance).To(Equal(2000.0))

			// Verify that the currency is the same
			Expect(clonedAccount.Currency).To(Equal(baseAccount.Currency))
		})
	})
})
