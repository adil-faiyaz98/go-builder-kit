package tests

import (
	"github.com/adil-faiyaz98/go-builder-kit/builders"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bond", func() {
	Context("Bond Builder", func() {
		It("should build a valid Bond with all fields", func() {
			// Create a Bond builder
			bondBuilder := builders.NewBondBuilder().
				WithID("B12345").
				WithISIN("US1234567890").
				WithName("US Treasury Bond").
				WithIssuer("US Treasury").
				WithType("treasury").
				WithFaceValue(1000.0).
				WithCouponRate(0.025).
				WithMaturityDate("2030-01-01").
				WithPurchaseDate("2020-01-01").
				WithPurchasePrice(980.0).
				WithCurrentPrice(1020.0).
				WithQuantity(10).
				WithCurrency("USD").
				WithPaymentFrequency("semi-annual").
				WithRating("AAA").
				WithYield(0.03)

			// Build the Bond
			bond, err := bondBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(bond).NotTo(BeNil())
			Expect(bond.ID).To(Equal("B12345"))
			Expect(bond.ISIN).To(Equal("US1234567890"))
			Expect(bond.Name).To(Equal("US Treasury Bond"))
			Expect(bond.Issuer).To(Equal("US Treasury"))
			Expect(bond.Type).To(Equal("treasury"))
			Expect(bond.FaceValue).To(Equal(1000.0))
			Expect(bond.CouponRate).To(Equal(0.025))
			Expect(bond.MaturityDate).To(Equal("2030-01-01"))
			Expect(bond.PurchaseDate).To(Equal("2020-01-01"))
			Expect(bond.PurchasePrice).To(Equal(980.0))
			Expect(bond.CurrentPrice).To(Equal(1020.0))
			Expect(bond.Quantity).To(Equal(10))
			Expect(bond.Currency).To(Equal("USD"))
			Expect(bond.PaymentFrequency).To(Equal("semi-annual"))
			Expect(bond.Rating).To(Equal("AAA"))
			Expect(bond.Yield).To(Equal(0.03))
		})

		It("should validate required fields", func() {
			// Create a Bond builder with missing required fields
			bondBuilder := builders.NewBondBuilder()
			// Missing ISIN, Name, and Issuer

			// Build and validate should fail
			_, err := bondBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("ISIN is required"))
			Expect(err.Error()).To(ContainSubstring("Name is required"))
			Expect(err.Error()).To(ContainSubstring("Issuer is required"))
		})

		It("should validate date format", func() {
			// Create a Bond builder with invalid date
			bondBuilder := builders.NewBondBuilder().
				WithID("B12345").
				WithISIN("US1234567890").
				WithName("US Treasury Bond").
				WithIssuer("US Treasury").
				WithType("treasury").
				WithFaceValue(1000.0).
				WithPurchasePrice(980.0).
				WithCurrentPrice(1020.0).
				WithQuantity(10).
				WithCurrency("USD").
				WithPaymentFrequency("semi-annual").
				WithMaturityDate("01/01/2030") // Wrong format, should be YYYY-MM-DD

			// Build and validate should fail
			_, err := bondBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("MaturityDate must be in the format YYYY-MM-DD"))
		})

		It("should validate maturity date is after purchase date", func() {
			// Create a Bond builder with maturity date before purchase date
			bondBuilder := builders.NewBondBuilder().
				WithID("B12345").
				WithISIN("US1234567890").
				WithName("US Treasury Bond").
				WithIssuer("US Treasury").
				WithType("treasury").
				WithPurchaseDate("2030-01-01").
				WithMaturityDate("2020-01-01") // Before purchase date

			// Build and validate should fail
			_, err := bondBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("MaturityDate cannot be before PurchaseDate"))
		})

		It("should clone the builder correctly", func() {
			// Create a base bond builder
			baseBuilder := builders.NewBondBuilder().
				WithID("B12345").
				WithISIN("US1234567890").
				WithName("US Treasury Bond").
				WithIssuer("US Treasury").
				WithType("treasury").
				WithFaceValue(1000.0).
				WithCouponRate(0.025).
				WithPurchaseDate("2020-01-01").
				WithMaturityDate("2030-01-01")

			// Clone the builder and modify it
			clonedBuilder := baseBuilder.Clone().
				WithID("B67890").
				WithName("Corporate Bond").
				WithIssuer("Acme Corp")

			// Build both bonds
			baseBond := baseBuilder.BuildPtr()
			clonedBond := clonedBuilder.BuildPtr()

			// Verify the base bond
			Expect(baseBond.ID).To(Equal("B12345"))
			Expect(baseBond.Name).To(Equal("US Treasury Bond"))
			Expect(baseBond.Issuer).To(Equal("US Treasury"))

			// Verify the cloned bond
			Expect(clonedBond.ID).To(Equal("B67890"))
			Expect(clonedBond.Name).To(Equal("Corporate Bond"))
			Expect(clonedBond.Issuer).To(Equal("Acme Corp"))

			// Verify that the dates and rates are the same
			Expect(clonedBond.PurchaseDate).To(Equal(baseBond.PurchaseDate))
			Expect(clonedBond.MaturityDate).To(Equal(baseBond.MaturityDate))
			Expect(clonedBond.CouponRate).To(Equal(baseBond.CouponRate))
		})
	})
})
