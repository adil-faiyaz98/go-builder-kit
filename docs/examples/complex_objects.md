# Working with Complex Nested Objects

This document demonstrates how to use Go Builder Kit with complex nested objects.

## Creating Nested Objects

Go Builder Kit makes it easy to create complex nested objects:

```go
package main

import (
    "fmt"

    "github.com/adil-faiyaz98/go-builder-kit/builders"
    "github.com/adil-faiyaz98/go-builder-kit/models"
)

func main() {
    // Create a person with nested objects
    person := builders.NewPersonBuilder().
        WithID("123").
        WithName("John Doe").
        WithAge(30).
        WithEmail("john.doe@example.com").
        // Add an address
        WithAddress(
            builders.NewAddressBuilder().
                WithStreet("123 Main St").
                WithCity("New York").
                WithState("NY").
                WithPostalCode("10001").
                WithCountry("USA").
                WithType("Home").
                WithIsPrimary(true),
        ).
        // Add education
        WithEducation(
            builders.NewEducationBuilder().
                WithDegree("Bachelor of Science").
                WithInstitution("Harvard University").
                WithStartDate("2018-09-01").
                WithEndDate("2022-05-15").
                WithGPA(3.8).
                WithMajor("Computer Science").
                WithMinor("Mathematics"),
        ).
        // Add employment
        WithEmployment(
            builders.NewEmploymentBuilder().
                WithPosition("Software Engineer").
                WithDepartment("Engineering").
                WithStartDate("2022-06-01").
                WithSalary(100000.0).
                WithIsCurrent(true).
                // Add company
                WithCompany(
                    builders.NewCompanyBuilder().
                        WithID("COMP-001").
                        WithName("Tech Innovations Inc.").
                        WithIndustry("Technology").
                        // Add company location
                        WithLocation(
                            builders.NewAddressBuilder().
                                WithStreet("456 Tech Blvd").
                                WithCity("San Francisco").
                                WithState("CA").
                                WithPostalCode("94107").
                                WithCountry("USA"),
                        ),
                ),
        ).
        Build().(models.Person)

    fmt.Printf("Created person: %s\n", person.Name)
    fmt.Printf("Address: %s, %s, %s\n", 
        person.Address.Street, person.Address.City, person.Address.Country)
    fmt.Printf("Education: %s in %s from %s\n", 
        person.Education.Degree, person.Education.Major, person.Education.Institution)
    fmt.Printf("Employment: %s at %s\n", 
        person.Employment.Position, person.Employment.Company.Name)
}
```

## Working with Collections of Nested Objects

Go Builder Kit also supports collections of nested objects:

```go
package main

import (
    "fmt"

    "github.com/adil-faiyaz98/go-builder-kit/builders"
    "github.com/adil-faiyaz98/go-builder-kit/models"
)

func main() {
    // Create a bank with multiple accounts
    bank := builders.NewBankBuilder().
        WithName("Global Bank").
        WithBranchCode("GB-001").
        // Add first account
        WithAccount(
            builders.NewAccountBuilder().
                WithNumber("ACCT-12345").
                WithBalance(5000.0).
                WithType("Checking").
                WithCurrency("USD").
                WithStatus("Active").
                // Add transactions to the account
                WithTransaction(
                    builders.NewTransactionBuilder().
                        WithID("TX-001").
                        WithDate("2023-01-15").
                        WithAmount(1000.0).
                        WithType("Deposit").
                        WithDescription("Salary deposit"),
                ).
                WithTransaction(
                    builders.NewTransactionBuilder().
                        WithID("TX-002").
                        WithDate("2023-01-20").
                        WithAmount(200.0).
                        WithType("Withdrawal").
                        WithDescription("ATM withdrawal"),
                ),
        ).
        // Add second account
        WithAccount(
            builders.NewAccountBuilder().
                WithNumber("ACCT-67890").
                WithBalance(10000.0).
                WithType("Savings").
                WithCurrency("USD").
                WithInterestRate(1.5).
                WithStatus("Active"),
        ).
        Build().(models.Bank)

    fmt.Printf("Bank: %s\n", bank.Name)
    fmt.Printf("Number of accounts: %d\n", len(bank.Accounts))
    
    // Access the first account
    checkingAccount := bank.Accounts[0]
    fmt.Printf("Checking account balance: $%.2f\n", checkingAccount.Balance)
    fmt.Printf("Number of transactions: %d\n", len(checkingAccount.Transactions))
    
    // Access the second account
    savingsAccount := bank.Accounts[1]
    fmt.Printf("Savings account balance: $%.2f\n", savingsAccount.Balance)
    fmt.Printf("Interest rate: %.2f%%\n", savingsAccount.InterestRate)
}
```

## Creating Complex Investment Portfolios

Here's an example of creating a complex investment portfolio:

```go
package main

import (
    "fmt"

    "github.com/adil-faiyaz98/go-builder-kit/builders"
    "github.com/adil-faiyaz98/go-builder-kit/models"
)

func main() {
    // Create a portfolio
    portfolio := builders.NewPortfolioBuilder().
        // Add stocks
        WithStock(
            builders.NewStockBuilder().
                WithTicker("AAPL").
                WithCompanyName("Apple Inc.").
                WithPurchasePrice(150.0).
                WithCurrentPrice(170.0).
                WithUnits(10).
                WithSector("Technology"),
        ).
        WithStock(
            builders.NewStockBuilder().
                WithTicker("MSFT").
                WithCompanyName("Microsoft Corporation").
                WithPurchasePrice(250.0).
                WithCurrentPrice(280.0).
                WithUnits(5).
                WithSector("Technology"),
        ).
        // Add bonds
        WithBond(
            builders.NewBondBuilder().
                WithIssuer("US Treasury").
                WithFaceValue(1000.0).
                WithCouponRate(2.5).
                WithMaturityDate("2030-01-15").
                WithType("Government").
                WithQuantity(5),
        ).
        // Add ETFs
        WithETF(
            builders.NewETFBuilder().
                WithName("S&P 500 ETF").
                WithTickerSymbol("SPY").
                WithExpenseRatio(0.09).
                WithShares(10.0),
        ).
        // Set allocations
        WithAllocation("Stocks", 60.0).
        WithAllocation("Bonds", 30.0).
        WithAllocation("ETFs", 10.0).
        Build().(models.Portfolio)

    // Create an investment with the portfolio
    investment := builders.NewInvestmentBuilder().
        WithID("INV-001").
        WithName("Retirement Portfolio").
        WithType("401k").
        WithValue(50000.0).
        WithStartDate("2020-01-15").
        WithRisk("Medium").
        WithPortfolio(
            builders.NewPortfolioBuilder().
                WithStocks(portfolio.Stocks).
                WithBonds(portfolio.Bonds).
                WithETFs(portfolio.ETFs).
                WithAllocation("Stocks", 60.0).
                WithAllocation("Bonds", 30.0).
                WithAllocation("ETFs", 10.0),
        ).
        // Add performance records
        WithPerformanceRecord(
            builders.NewPerformanceRecordBuilder().
                WithDate("2022-01-01").
                WithValue(45000.0).
                WithROI(5.0).
                WithNotes("Q1 performance"),
        ).
        WithPerformanceRecord(
            builders.NewPerformanceRecordBuilder().
                WithDate("2022-04-01").
                WithValue(47500.0).
                WithROI(5.5).
                WithNotes("Q2 performance"),
        ).
        Build().(models.Investment)

    fmt.Printf("Investment: %s\n", investment.Name)
    fmt.Printf("Value: $%.2f\n", investment.Value)
    fmt.Printf("Number of stocks: %d\n", len(investment.Portfolio.Stocks))
    fmt.Printf("Number of bonds: %d\n", len(investment.Portfolio.Bonds))
    fmt.Printf("Number of ETFs: %d\n", len(investment.Portfolio.ETFs))
    fmt.Printf("Number of performance records: %d\n", len(investment.Performance))
}
```

## Validating Complex Nested Objects

Validation works with complex nested objects too:

```go
package main

import (
    "fmt"

    "github.com/adil-faiyaz98/go-builder-kit/builders"
)

func main() {
    // Create a bank with validation
    bank, err := builders.NewBankBuilder().
        WithName("Global Bank").
        WithBranchCode("GB-001").
        WithAccount(
            builders.NewAccountBuilder().
                WithNumber("ACCT-12345").
                WithBalance(5000.0).
                WithType("Checking").
                WithCurrency("USD").
                WithStatus("Active"),
        ).
        BuildWithValidation()

    if err != nil {
        fmt.Printf("Validation error: %v\n", err)
        return
    }

    fmt.Printf("Created valid bank: %s\n", bank.Name)

    // Example with validation error in nested object
    _, err = builders.NewBankBuilder().
        WithName("Global Bank").
        WithBranchCode("GB-001").
        WithAccount(
            builders.NewAccountBuilder().
                WithNumber("").  // Invalid: empty account number
                WithBalance(5000.0).
                WithType("Unknown"),  // Invalid: unknown account type
        ).
        BuildWithValidation()

    if err != nil {
        fmt.Printf("Validation error: %v\n", err)
        // Output: Validation error: validation failed: Account[0] validation failed: validation failed: Number cannot be empty; Type must be one of: Checking, Savings, Credit, Loan, Investment, Other
    }
}
```
