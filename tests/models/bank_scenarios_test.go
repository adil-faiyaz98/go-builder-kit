package models_test

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

var _ = Describe("Bank Scenarios", func() {
	// Helper function to create a valid address builder
	createValidAddressBuilder := func() *builders.AddressBuilder {
		return builders.NewAddressBuilder().
			WithStreet("123 Main St").
			WithCity("New York").
			WithState("NY").
			WithZipCode("10001").
			WithCountry("USA")
	}

	// Helper function to create a valid transaction builder
	createValidTransactionBuilder := func() *builders.TransactionBuilder {
		return builders.NewTransactionBuilder().
			WithID("txn-123").
			WithType("deposit").
			WithAmount(1000.00).
			WithCurrency("USD").
			WithDate(time.Now().AddDate(0, 0, -5).Format("2006-01-02")). // 5 days ago
			WithDescription("Salary deposit").
			WithCategory("income").
			WithStatus("completed").
			WithReference("REF123456")
	}

	// Helper function to create a valid account builder
	createValidAccountBuilder := func() *builders.AccountBuilder {
		return builders.NewAccountBuilder().
			WithID("acc-123").
			WithType("checking").
			WithNumber("1234567890").
			WithBalance(5000.00).
			WithCurrency("USD").
			WithOpenDate(time.Now().AddDate(-1, 0, 0).Format("2006-01-02")). // 1 year ago
			WithStatus("active").
			WithInterestRate(0.01).
			WithIsJoint(false).
			WithOverdraftLimit(1000.00)
	}

	// Helper function to create a valid loan builder
	createValidLoanBuilder := func() *builders.LoanBuilder {
		return builders.NewLoanBuilder().
			WithID("loan-123").
			WithType("mortgage").
			WithAmount(250000.00).
			WithInterestRate(3.5).
			WithTerm(30).
			WithStartDate(time.Now().AddDate(-2, 0, 0).Format("2006-01-02")). // 2 years ago
			WithEndDate(time.Now().AddDate(28, 0, 0).Format("2006-01-02")).   // 28 years from now
			WithPaymentAmount(1120.00).
			WithPaymentFrequency("monthly").
			WithStatus("active")
	}

	// Helper function to create a valid stock builder
	createValidStockBuilder := func() *builders.StockBuilder {
		return builders.NewStockBuilder().
			WithSymbol("AAPL").
			WithName("Apple Inc.").
			WithPurchasePrice(150.00).
			WithCurrentPrice(175.00).
			WithQuantity(10).
			WithPurchaseDate(time.Now().AddDate(0, -6, 0).Format("2006-01-02")). // 6 months ago
			WithCurrency("USD").
			WithSector("Technology")
	}

	// Helper function to create a valid investment builder
	createValidInvestmentBuilder := func() *builders.InvestmentBuilder {
		return builders.NewInvestmentBuilder().
			WithID("inv-123").
			WithName("Retirement Fund").
			WithType("mutual fund").
			WithAmount(50000.00).
			WithRisk("medium").
			WithStartDate(time.Now().AddDate(-3, 0, 0).Format("2006-01-02")). // 3 years ago
			WithStatus("active")
	}

	Describe("Valid Scenarios", func() {
		Context("when creating a complete bank with all valid data", func() {
			It("should build and validate successfully", func() {
				// Create a bank with all valid data
				bankBuilder := builders.NewBankBuilder().
					WithName("Global Bank").
					WithBranchCode("GB-NYC-001").
					WithAddress(createValidAddressBuilder())

				// Add accounts with transactions
				account1 := createValidAccountBuilder().
					WithID("acc-123").
					WithType("checking").
					WithNumber("1234567890")

				account2 := createValidAccountBuilder().
					WithID("acc-456").
					WithType("savings").
					WithNumber("0987654321").
					WithInterestRate(1.5)

				// Add transactions to accounts
				for i := 0; i < 3; i++ {
					account1.WithTransaction(
						createValidTransactionBuilder().
							WithID(fmt.Sprintf("txn-%d", i)).
							WithAmount(float64(100 * (i + 1))).
							WithDescription(fmt.Sprintf("Transaction %d", i+1)),
					)
				}

				// Add accounts to bank
				bankBuilder.WithAccount(account1).WithAccount(account2)

				// Add loans
				bankBuilder.WithLoan(createValidLoanBuilder())

				// Add stocks
				bankBuilder.WithStock(createValidStockBuilder())

				// Add investments
				bankBuilder.WithInvestment(createValidInvestmentBuilder())

				// Set relationship start date
				bankBuilder.WithRelationshipStartDate(time.Now().AddDate(-5, 0, 0).Format("2006-01-02")) // 5 years ago

				// Build the bank
				bank := bankBuilder.Build().(*models.Bank)

				// Validate the bank
				err := bank.Validate()
				Expect(err).NotTo(HaveOccurred(), "Expected validation to pass for a valid bank")

				// Verify the structure is complete
				Expect(bank.Name).To(Equal("Global Bank"), "Expected bank name to be 'Global Bank'")
				Expect(bank.BranchCode).To(Equal("GB-NYC-001"), "Expected branch code to be 'GB-NYC-001'")
				Expect(bank.Address).NotTo(BeNil(), "Expected address to be set")
				Expect(bank.Accounts).To(HaveLen(2), "Expected 2 accounts")
				Expect(bank.Loans).To(HaveLen(1), "Expected 1 loan")
				Expect(bank.Stocks).To(HaveLen(1), "Expected 1 stock")
				Expect(bank.Investments).To(HaveLen(1), "Expected 1 investment")

				// Verify nested structures
				address := bank.Address
				Expect(address.Street).To(Equal("123 Main St"), "Expected street to be '123 Main St'")
				Expect(address.City).To(Equal("New York"), "Expected city to be 'New York'")

				// Verify accounts
				checkingAccount := bank.Accounts[0]
				savingsAccount := bank.Accounts[1]

				Expect(checkingAccount.Type).To(Equal("checking"), "Expected account type to be 'checking'")
				Expect(checkingAccount.Number).To(Equal("1234567890"), "Expected account number to be '1234567890'")
				Expect(checkingAccount.Transactions).To(HaveLen(3), "Expected 3 transactions in checking account")

				Expect(savingsAccount.Type).To(Equal("savings"), "Expected account type to be 'savings'")
				Expect(savingsAccount.InterestRate).To(Equal(1.5), "Expected interest rate to be 1.5")

				// Verify transactions in checking account
				for i, transaction := range checkingAccount.Transactions {
					Expect(transaction.ID).To(Equal(fmt.Sprintf("txn-%d", i)),
						"Expected transaction ID to match")
					Expect(transaction.Amount).To(Equal(float64(100*(i+1))),
						"Expected transaction amount to match")
				}

				// Verify loan
				loan := bank.Loans[0]
				Expect(loan.Type).To(Equal("mortgage"), "Expected loan type to be 'mortgage'")
				Expect(loan.Amount).To(Equal(250000.00), "Expected loan amount to be 250000.00")
				Expect(loan.Status).To(Equal("active"), "Expected loan status to be 'active'")

				// Verify stock
				stock := bank.Stocks[0]
				Expect(stock.Symbol).To(Equal("AAPL"), "Expected stock symbol to be 'AAPL'")
				Expect(stock.Name).To(Equal("Apple Inc."), "Expected stock name to be 'Apple Inc.'")

				// Verify investment
				investment := bank.Investments[0]
				Expect(investment.Name).To(Equal("Retirement Fund"), "Expected investment name to be 'Retirement Fund'")
				Expect(investment.Type).To(Equal("mutual fund"), "Expected investment type to be 'mutual fund'")
				Expect(investment.Risk).To(Equal("medium"), "Expected investment risk to be 'medium'")
			})
		})

		Context("when creating a bank with minimal valid data", func() {
			It("should build and validate successfully", func() {
				// Create a bank with minimal data
				bankBuilder := builders.NewBankBuilder().
					WithName("Minimal Bank").
					WithBranchCode("MB-001").
					WithAddress(createValidAddressBuilder())

				// Build the bank
				bank := bankBuilder.Build().(*models.Bank)

				// Validate the bank
				err := bank.Validate()
				Expect(err).NotTo(HaveOccurred(), "Expected validation to pass for a minimal bank")

				// Verify the structure
				Expect(bank.Name).To(Equal("Minimal Bank"), "Expected bank name to be 'Minimal Bank'")
				Expect(bank.BranchCode).To(Equal("MB-001"), "Expected branch code to be 'MB-001'")
				Expect(bank.Address).NotTo(BeNil(), "Expected address to be set")
				Expect(bank.Accounts).To(BeEmpty(), "Expected no accounts")
				Expect(bank.Loans).To(BeEmpty(), "Expected no loans")
				Expect(bank.Stocks).To(BeEmpty(), "Expected no stocks")
				Expect(bank.Investments).To(BeEmpty(), "Expected no investments")
			})
		})
	})

	Describe("Invalid Scenarios", func() {
		Context("when creating a bank with missing required fields", func() {
			It("should fail validation with appropriate error messages for missing name", func() {
				// Create a bank with missing name
				bankBuilder := builders.NewBankBuilder().
					WithBranchCode("GB-NYC-001").
					WithAddress(createValidAddressBuilder())

				// Build the bank
				bank := bankBuilder.Build().(*models.Bank)

				// Validate the bank
				err := bank.Validate()
				Expect(err).To(HaveOccurred(), "Expected validation to fail for missing name")

				// Check error messages
				errString := err.Error()
				Expect(errString).To(ContainSubstring("Name is required"),
					"Expected error about missing name")
			})

			It("should fail validation with appropriate error messages for missing branch code", func() {
				// Create a bank with missing branch code
				bankBuilder := builders.NewBankBuilder().
					WithName("Global Bank").
					WithAddress(createValidAddressBuilder())

				// Build the bank
				bank := bankBuilder.Build().(*models.Bank)

				// Validate the bank
				err := bank.Validate()
				Expect(err).To(HaveOccurred(), "Expected validation to fail for missing branch code")

				// Check error messages
				errString := err.Error()
				Expect(errString).To(ContainSubstring("BranchCode is required"),
					"Expected error about missing branch code")
			})

			It("should fail validation with appropriate error messages for missing address", func() {
				// Create a bank with missing address
				bankBuilder := builders.NewBankBuilder().
					WithName("Global Bank").
					WithBranchCode("GB-NYC-001")

				// Build the bank
				bank := bankBuilder.Build().(*models.Bank)

				// Validate the bank
				err := bank.Validate()
				Expect(err).To(HaveOccurred(), "Expected validation to fail for missing address")

				// Check error messages
				errString := err.Error()
				Expect(errString).To(ContainSubstring("Address is required"),
					"Expected error about missing address")
			})
		})

		Context("when creating a bank with invalid account data", func() {
			It("should fail validation with appropriate error messages", func() {
				// Create an account with invalid data
				invalidAccountBuilder := builders.NewAccountBuilder().
					WithID("").               // Empty ID
					WithType("invalid-type"). // Invalid type
					WithNumber("").           // Empty number
					WithBalance(-1000.00).    // Negative balance
					WithCurrency("").         // Empty currency
					WithStatus("invalid")     // Invalid status

				// Create a bank with the invalid account
				bankBuilder := builders.NewBankBuilder().
					WithName("Global Bank").
					WithBranchCode("GB-NYC-001").
					WithAddress(createValidAddressBuilder()).
					WithAccount(invalidAccountBuilder)

				// Build the bank
				bank := bankBuilder.Build().(*models.Bank)

				// Validate the bank
				err := bank.Validate()
				Expect(err).To(HaveOccurred(), "Expected validation to fail for invalid account data")

				// Check error messages
				errString := err.Error()
				Expect(errString).To(ContainSubstring("ID is required"),
					"Expected error about empty ID")
				Expect(errString).To(ContainSubstring("Type must be one of"),
					"Expected error about invalid account type")
				Expect(errString).To(ContainSubstring("Number is required"),
					"Expected error about empty number")
				Expect(errString).To(ContainSubstring("Balance cannot be negative"),
					"Expected error about negative balance")
				Expect(errString).To(ContainSubstring("Currency is required"),
					"Expected error about empty currency")
				Expect(errString).To(ContainSubstring("Status must be one of"),
					"Expected error about invalid status")
			})
		})

		Context("when creating a bank with invalid transaction data", func() {
			It("should fail validation with appropriate error messages", func() {
				// Create a transaction with invalid data
				invalidTransactionBuilder := builders.NewTransactionBuilder().
					WithID("").               // Empty ID
					WithType("invalid-type"). // Invalid type
					WithAmount(-100.00).      // Negative amount
					WithCurrency("").         // Empty currency
					WithDate("invalid-date"). // Invalid date
					WithStatus("invalid")     // Invalid status

				// Create an account with the invalid transaction
				accountBuilder := createValidAccountBuilder().
					WithTransaction(invalidTransactionBuilder)

				// Create a bank with the account
				bankBuilder := builders.NewBankBuilder().
					WithName("Global Bank").
					WithBranchCode("GB-NYC-001").
					WithAddress(createValidAddressBuilder()).
					WithAccount(accountBuilder)

				// Build the bank
				bank := bankBuilder.Build().(*models.Bank)

				// Validate the bank
				err := bank.Validate()
				Expect(err).To(HaveOccurred(), "Expected validation to fail for invalid transaction data")

				// Check error messages
				errString := err.Error()
				Expect(errString).To(ContainSubstring("ID is required"),
					"Expected error about empty transaction ID")
				Expect(errString).To(ContainSubstring("Type must be one of"),
					"Expected error about invalid transaction type")
				Expect(errString).To(ContainSubstring("Amount cannot be negative"),
					"Expected error about negative amount")
				Expect(errString).To(ContainSubstring("Currency is required"),
					"Expected error about empty currency")
				Expect(errString).To(ContainSubstring("Date format is invalid"),
					"Expected error about invalid date format")
				Expect(errString).To(ContainSubstring("Status must be one of"),
					"Expected error about invalid status")
			})
		})

		Context("when creating a bank with invalid loan data", func() {
			It("should fail validation with appropriate error messages", func() {
				// Create a loan with invalid data
				invalidLoanBuilder := builders.NewLoanBuilder().
					WithID("").                  // Empty ID
					WithType("invalid-type").    // Invalid type
					WithAmount(-250000.00).      // Negative amount
					WithInterestRate(-3.5).      // Negative interest rate
					WithTerm(-30).               // Negative term
					WithStartDate("2023-01-01"). // Start date after end date
					WithEndDate("2022-01-01").   // End date before start date
					WithStatus("invalid")        // Invalid status

				// Create a bank with the invalid loan
				bankBuilder := builders.NewBankBuilder().
					WithName("Global Bank").
					WithBranchCode("GB-NYC-001").
					WithAddress(createValidAddressBuilder()).
					WithLoan(invalidLoanBuilder)

				// Build the bank
				bank := bankBuilder.Build().(*models.Bank)

				// Validate the bank
				err := bank.Validate()
				Expect(err).To(HaveOccurred(), "Expected validation to fail for invalid loan data")

				// Check error messages
				errString := err.Error()
				Expect(errString).To(ContainSubstring("ID is required"),
					"Expected error about empty loan ID")
				Expect(errString).To(ContainSubstring("Type must be one of"),
					"Expected error about invalid loan type")
				Expect(errString).To(ContainSubstring("Amount cannot be negative"),
					"Expected error about negative amount")
				Expect(errString).To(ContainSubstring("InterestRate cannot be negative"),
					"Expected error about negative interest rate")
				Expect(errString).To(ContainSubstring("Term cannot be negative"),
					"Expected error about negative term")
				Expect(errString).To(ContainSubstring("EndDate cannot be before StartDate"),
					"Expected error about end date before start date")
				Expect(errString).To(ContainSubstring("Status must be one of"),
					"Expected error about invalid status")
			})
		})
	})

	Describe("Edge Cases", func() {
		Context("when creating a bank with many accounts", func() {
			It("should handle many accounts correctly", func() {
				// Create a bank with many accounts
				bankBuilder := builders.NewBankBuilder().
					WithName("Multi-Account Bank").
					WithBranchCode("MAB-001").
					WithAddress(createValidAddressBuilder())

				// Add many accounts
				accountCount := 20
				for i := 0; i < accountCount; i++ {
					accountBuilder := createValidAccountBuilder().
						WithID(fmt.Sprintf("acc-%d", i)).
						WithNumber(fmt.Sprintf("1000%d", i)).
						WithBalance(float64(1000 * (i + 1)))

					bankBuilder.WithAccount(accountBuilder)
				}

				// Build the bank
				bank := bankBuilder.Build().(*models.Bank)

				// Validate the bank
				err := bank.Validate()
				Expect(err).NotTo(HaveOccurred(), "Expected validation to pass for bank with many accounts")

				// Verify the accounts
				Expect(bank.Accounts).To(HaveLen(accountCount),
					fmt.Sprintf("Expected %d accounts", accountCount))

				// Verify each account
				for i, account := range bank.Accounts {
					Expect(account.ID).To(Equal(fmt.Sprintf("acc-%d", i)),
						"Expected account ID to match")
					Expect(account.Number).To(Equal(fmt.Sprintf("1000%d", i)),
						"Expected account number to match")
					Expect(account.Balance).To(Equal(float64(1000*(i+1))),
						"Expected account balance to match")
				}
			})
		})

		Context("when creating a bank with an account that has many transactions", func() {
			It("should handle many transactions correctly", func() {
				// Create an account with many transactions
				accountBuilder := createValidAccountBuilder()

				// Add many transactions
				transactionCount := 50
				for i := 0; i < transactionCount; i++ {
					transactionBuilder := createValidTransactionBuilder().
						WithID(fmt.Sprintf("txn-%d", i)).
						WithAmount(float64(10 * (i + 1))).
						WithDescription(fmt.Sprintf("Transaction %d", i+1)).
						WithDate(time.Now().AddDate(0, 0, -i).Format("2006-01-02")) // i days ago

					accountBuilder.WithTransaction(transactionBuilder)
				}

				// Create a bank with the account
				bankBuilder := builders.NewBankBuilder().
					WithName("Transaction-Heavy Bank").
					WithBranchCode("THB-001").
					WithAddress(createValidAddressBuilder()).
					WithAccount(accountBuilder)

				// Build the bank
				bank := bankBuilder.Build().(*models.Bank)

				// Validate the bank
				err := bank.Validate()
				Expect(err).NotTo(HaveOccurred(), "Expected validation to pass for account with many transactions")

				// Verify the transactions
				account := bank.Accounts[0]
				Expect(account.Transactions).To(HaveLen(transactionCount),
					fmt.Sprintf("Expected %d transactions", transactionCount))

				// Verify each transaction
				for i, transaction := range account.Transactions {
					Expect(transaction.ID).To(Equal(fmt.Sprintf("txn-%d", i)),
						"Expected transaction ID to match")
					Expect(transaction.Amount).To(Equal(float64(10*(i+1))),
						"Expected transaction amount to match")
					Expect(transaction.Description).To(Equal(fmt.Sprintf("Transaction %d", i+1)),
						"Expected transaction description to match")
				}
			})
		})

		Context("when creating a bank with extreme values", func() {
			It("should handle extreme values correctly", func() {
				// Create an account with extreme values
				accountBuilder := createValidAccountBuilder().
					WithBalance(9999999999.99).    // Very large balance
					WithInterestRate(99.99).       // Very high interest rate
					WithOverdraftLimit(1000000.00) // Very high overdraft limit

				// Create a loan with extreme values
				loanBuilder := createValidLoanBuilder().
					WithAmount(10000000.00). // Very large loan amount
					WithInterestRate(29.99). // Very high interest rate
					WithTerm(40)             // Very long term

				// Create a bank with extreme values
				bankBuilder := builders.NewBankBuilder().
					WithName("Extreme Values Bank").
					WithBranchCode("EVB-001").
					WithAddress(createValidAddressBuilder()).
					WithAccount(accountBuilder).
					WithLoan(loanBuilder)

				// Build the bank
				bank := bankBuilder.Build().(*models.Bank)

				// Validate the bank
				err := bank.Validate()
				Expect(err).NotTo(HaveOccurred(), "Expected validation to pass for extreme values")

				// Verify the extreme values
				account := bank.Accounts[0]
				Expect(account.Balance).To(Equal(9999999999.99), "Expected extreme balance")
				Expect(account.InterestRate).To(Equal(99.99), "Expected extreme interest rate")
				Expect(account.OverdraftLimit).To(Equal(1000000.00), "Expected extreme overdraft limit")

				loan := bank.Loans[0]
				Expect(loan.Amount).To(Equal(10000000.00), "Expected extreme loan amount")
				Expect(loan.InterestRate).To(Equal(29.99), "Expected extreme interest rate")
				Expect(loan.Term).To(Equal(30), "Expected extreme term")
			})
		})

		Context("when creating a bank with zero values", func() {
			It("should handle zero values correctly", func() {
				// Create an account with zero values
				accountBuilder := createValidAccountBuilder().
					WithBalance(0.00).       // Zero balance
					WithInterestRate(0.00).  // Zero interest rate
					WithOverdraftLimit(0.00) // Zero overdraft limit

				// Create a bank with zero values
				bankBuilder := builders.NewBankBuilder().
					WithName("Zero Values Bank").
					WithBranchCode("ZVB-001").
					WithAddress(createValidAddressBuilder()).
					WithAccount(accountBuilder)

				// Build the bank
				bank := bankBuilder.Build().(*models.Bank)

				// Validate the bank
				err := bank.Validate()
				Expect(err).NotTo(HaveOccurred(), "Expected validation to pass for zero values")

				// Verify the zero values
				account := bank.Accounts[0]
				Expect(account.Balance).To(Equal(0.00), "Expected zero balance")
				Expect(account.InterestRate).To(Equal(0.00), "Expected zero interest rate")
				Expect(account.OverdraftLimit).To(Equal(0.00), "Expected zero overdraft limit")
			})
		})
	})

	Describe("Performance", func() {
		Context("when creating a bank with many nested objects", func() {
			It("should handle creating a complex bank efficiently", func() {
				// Create a bank with many nested objects
				bankBuilder := builders.NewBankBuilder().
					WithName("Complex Bank").
					WithBranchCode("CB-001").
					WithAddress(createValidAddressBuilder())

				// Add multiple accounts
				for i := 0; i < 10; i++ {
					accountBuilder := createValidAccountBuilder().
						WithID(fmt.Sprintf("acc-%d", i)).
						WithNumber(fmt.Sprintf("1000%d", i))

					// Add multiple transactions to each account
					for j := 0; j < 10; j++ {
						accountBuilder.WithTransaction(
							createValidTransactionBuilder().
								WithID(fmt.Sprintf("txn-%d-%d", i, j)).
								WithAmount(float64(100 * (j + 1))).
								WithDescription(fmt.Sprintf("Transaction %d-%d", i, j)),
						)
					}

					bankBuilder.WithAccount(accountBuilder)
				}

				// Add multiple loans
				for i := 0; i < 5; i++ {
					bankBuilder.WithLoan(
						createValidLoanBuilder().
							WithID(fmt.Sprintf("loan-%d", i)).
							WithAmount(100000.00 * float64(i+1)),
					)
				}

				// Add multiple stocks
				for i := 0; i < 5; i++ {
					bankBuilder.WithStock(
						createValidStockBuilder().
							WithSymbol(fmt.Sprintf("STOCK%d", i)).
							WithName(fmt.Sprintf("Stock %d", i)),
					)
				}

				// Add multiple investments
				for i := 0; i < 5; i++ {
					bankBuilder.WithInvestment(
						createValidInvestmentBuilder().
							WithID(fmt.Sprintf("inv-%d", i)).
							WithName(fmt.Sprintf("Investment %d", i)),
					)
				}

				// Build the bank
				startTime := time.Now()
				bank := bankBuilder.Build().(*models.Bank)
				buildTime := time.Since(startTime)

				// Validate the bank
				startTime = time.Now()
				err := bank.Validate()
				validateTime := time.Since(startTime)

				// Verify performance
				Expect(err).NotTo(HaveOccurred(), "Expected validation to pass for complex bank")
				Expect(buildTime).To(BeNumerically("<", 100*time.Millisecond),
					"Expected build time to be less than 100ms")
				Expect(validateTime).To(BeNumerically("<", 100*time.Millisecond),
					"Expected validation time to be less than 100ms")

				// Verify the structure
				Expect(bank.Accounts).To(HaveLen(10), "Expected 10 accounts")
				Expect(bank.Loans).To(HaveLen(5), "Expected 5 loans")
				Expect(bank.Stocks).To(HaveLen(5), "Expected 5 stocks")
				Expect(bank.Investments).To(HaveLen(5), "Expected 5 investments")

				// Verify nested structures
				for i, account := range bank.Accounts {
					Expect(account.ID).To(Equal(fmt.Sprintf("acc-%d", i)),
						"Expected account ID to match")
					Expect(account.Transactions).To(HaveLen(10),
						"Expected 10 transactions per account")
				}
			})
		})
	})
})
