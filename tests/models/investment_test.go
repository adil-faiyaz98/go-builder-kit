package models_test

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

var _ = Describe("Investment", func() {
	var (
		investmentBuilder        *builders.InvestmentBuilder
		portfolioBuilder         *builders.PortfolioBuilder
		performanceRecordBuilder *builders.PerformanceRecordBuilder
		stockBuilder             *builders.StockBuilder
		bondBuilder              *builders.BondBuilder
		etfBuilder               *builders.ETFBuilder
	)

	BeforeEach(func() {
		stockBuilder = builders.NewStockBuilder().
			WithTicker("AAPL").
			WithCompanyName("Apple Inc.").
			WithPurchasePrice(150.0).
			WithCurrentPrice(155.0).
			WithUnits(10).
			WithSector("Technology")

		bondBuilder = builders.NewBondBuilder().
			WithIssuer("US Treasury").
			WithFaceValue(1000.0).
			WithCouponRate(2.5).
			WithMaturityDate("2030-01-15").
			WithType("Government")

		etfBuilder = builders.NewETFBuilder().
			WithName("S&P 500 ETF").
			WithTickerSymbol("SPY").
			WithExpenseRatio(0.09).
			WithShares(5.0)

		portfolioBuilder = builders.NewPortfolioBuilder().
			WithStock(stockBuilder).
			WithBond(bondBuilder).
			WithETF(etfBuilder).
			WithAllocation("Stocks", 60.0).
			WithAllocation("Bonds", 30.0).
			WithAllocation("ETFs", 10.0)

		performanceRecordBuilder = builders.NewPerformanceRecordBuilder().
			WithDate(time.Now().AddDate(0, -1, 0).Format("2006-01-02")). // One month ago
			WithValue(10500.0).
			WithROI(5.0).
			WithNotes("Quarterly performance review")

		investmentBuilder = builders.NewInvestmentBuilder().
			WithID("INV-001").
			WithName("Retirement Portfolio").
			WithType("401k").
			WithValue(10000.0).
			WithStartDate("2020-01-15").
			WithRisk("Medium").
			WithPortfolio(portfolioBuilder).
			WithPerformanceRecord(performanceRecordBuilder)
	})

	Describe("Validation", func() {
		Context("when all fields are valid", func() {
			It("should pass validation", func() {
				investment, err := investmentBuilder.BuildWithValidation()
				Expect(err).NotTo(HaveOccurred())
				Expect(investment).NotTo(BeNil())
				Expect(investment.Name).To(Equal("Retirement Portfolio"))
			})
		})

		Context("when required fields are missing", func() {
			It("should fail validation when ID is empty", func() {
				_, err := investmentBuilder.
					WithID("").
					BuildWithValidation()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("ID cannot be empty"))
			})

			It("should fail validation when Name is empty", func() {
				_, err := investmentBuilder.
					WithName("").
					BuildWithValidation()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Name cannot be empty"))
			})
		})

		Context("when field values are invalid", func() {
			It("should fail validation when Type is invalid", func() {
				_, err := investmentBuilder.
					WithType("Unknown").
					BuildWithValidation()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Type must be one of"))
			})

			It("should fail validation when Value is negative", func() {
				_, err := investmentBuilder.
					WithValue(-1000.0).
					BuildWithValidation()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Value cannot be negative"))
			})

			It("should fail validation when StartDate is in the future", func() {
				futureDate := time.Now().AddDate(1, 0, 0).Format("2006-01-02")
				_, err := investmentBuilder.
					WithStartDate(futureDate).
					BuildWithValidation()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("StartDate cannot be in the future"))
			})

			It("should fail validation when StartDate format is invalid", func() {
				_, err := investmentBuilder.
					WithStartDate("01/15/2020").
					BuildWithValidation()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("StartDate must be in the format YYYY-MM-DD"))
			})

			It("should fail validation when Risk is invalid", func() {
				_, err := investmentBuilder.
					WithRisk("Extreme").
					BuildWithValidation()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Risk must be one of: Low, Medium, High"))
			})
		})

		Context("when nested objects are invalid", func() {
			It("should fail validation when Portfolio allocation is invalid", func() {
				invalidPortfolioBuilder := builders.NewPortfolioBuilder().
					WithStock(stockBuilder).
					WithBond(bondBuilder).
					WithAllocation("Stocks", 60.0).
					WithAllocation("Bonds", 20.0) // Total allocation is 80%, not 100%

				_, err := investmentBuilder.
					WithPortfolio(invalidPortfolioBuilder).
					BuildWithValidation()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Total allocation percentage should be 100%"))
			})

			It("should fail validation when PerformanceRecord date is in the future", func() {
				futureDate := time.Now().AddDate(1, 0, 0).Format("2006-01-02")
				invalidRecordBuilder := builders.NewPerformanceRecordBuilder().
					WithDate(futureDate).
					WithValue(11000.0).
					WithROI(10.0)

				_, err := investmentBuilder.
					WithPerformanceRecord(invalidRecordBuilder).
					BuildWithValidation()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Date cannot be in the future"))
			})

			It("should fail validation when PerformanceRecord value is negative", func() {
				invalidRecordBuilder := builders.NewPerformanceRecordBuilder().
					WithDate("2022-01-15").
					WithValue(-1000.0).
					WithROI(-10.0)

				_, err := investmentBuilder.
					WithPerformanceRecord(invalidRecordBuilder).
					BuildWithValidation()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Value cannot be negative"))
			})
		})
	})

	// Performance tests
	Describe("Performance", func() {
		Context("when creating many investments", func() {
			It("should handle creating 1000 investments efficiently", func() {
				startTime := time.Now()

				// Create 1000 investments
				for i := 1; i <= 1000; i++ {
					investment := builders.NewInvestmentBuilder().
						WithID(fmt.Sprintf("INV-%04d", i)).
						WithName(fmt.Sprintf("Investment %d", i)).
						WithType("401k").
						WithValue(float64(i * 1000)).
						WithStartDate("2020-01-15").
						WithRisk("Medium").
						Build().(models.Investment)

					// Just to ensure the compiler doesn't optimize away the creation
					Expect(investment.ID).To(Equal(fmt.Sprintf("INV-%04d", i)))
				}

				duration := time.Since(startTime)

				// This is a performance test, so we're just checking that it completes in a reasonable time
				// Adjust the threshold based on your machine's performance
				Expect(duration.Seconds()).To(BeNumerically("<", 1.0), "Creating 1000 investments should take less than 1 second")
			})

			It("should handle creating and validating 100 investments efficiently", func() {
				startTime := time.Now()

				// Create and validate 100 investments
				for i := 1; i <= 100; i++ {
					_, err := builders.NewInvestmentBuilder().
						WithID(fmt.Sprintf("INV-%04d", i)).
						WithName(fmt.Sprintf("Investment %d", i)).
						WithType("401k").
						WithValue(float64(i * 1000)).
						WithStartDate("2020-01-15").
						WithRisk("Medium").
						BuildWithValidation()

					Expect(err).NotTo(HaveOccurred())
				}

				duration := time.Since(startTime)

				// This is a performance test, so we're just checking that it completes in a reasonable time
				// Adjust the threshold based on your machine's performance
				Expect(duration.Seconds()).To(BeNumerically("<", 0.5), "Creating and validating 100 investments should take less than 0.5 seconds")
			})

			It("should handle creating a complex investment with many nested objects efficiently", func() {
				startTime := time.Now()

				complexPortfolioBuilder := builders.NewPortfolioBuilder()

				// Add 100 stocks to the portfolio
				for i := 1; i <= 100; i++ {
					stockBuilder := builders.NewStockBuilder().
						WithTicker(fmt.Sprintf("STOCK%d", i)).
						WithCompanyName(fmt.Sprintf("Company %d", i)).
						WithPurchasePrice(float64(i * 10)).
						WithCurrentPrice(float64(i * 10)).
						WithUnits(i)

					complexPortfolioBuilder.WithStock(stockBuilder)
				}

				// Add 50 performance records
				complexInvestmentBuilder := builders.NewInvestmentBuilder().
					WithID("COMPLEX-001").
					WithName("Complex Investment").
					WithType("401k").
					WithValue(1000000.0).
					WithStartDate("2020-01-15").
					WithRisk("Medium").
					WithPortfolio(complexPortfolioBuilder)

				for i := 1; i <= 50; i++ {
					date := time.Now().AddDate(0, -i, 0).Format("2006-01-02")
					recordBuilder := builders.NewPerformanceRecordBuilder().
						WithDate(date).
						WithValue(1000000.0 + float64(i*1000)).
						WithROI(float64(i) / 10.0)

					complexInvestmentBuilder.WithPerformanceRecord(recordBuilder)
				}

				investment, err := complexInvestmentBuilder.BuildWithValidation()

				duration := time.Since(startTime)

				Expect(err).NotTo(HaveOccurred())
				Expect(investment).NotTo(BeNil())
				Expect(len(investment.Portfolio.Stocks)).To(Equal(100))
				Expect(len(investment.Performance)).To(Equal(50))

				// This is a performance test, so we're just checking that it completes in a reasonable time
				// Adjust the threshold based on your machine's performance
				Expect(duration.Seconds()).To(BeNumerically("<", 0.1), "Creating a complex investment should take less than 0.1 seconds")
			})
		})
	})
})

