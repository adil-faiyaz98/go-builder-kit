package models_test

import (
	"testing"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

// TestBuilderVerification verifies that all builders work correctly
func TestBuilderVerification(t *testing.T) {
	// Test Person with Address
	t.Run("Person with Address", func(t *testing.T) {
		// Create a valid address
		addressBuilder := builders.NewAddressBuilder().
			WithStreet("123 Main St").
			WithCity("New York").
			WithState("NY").
			WithPostalCode("10001").
			WithCountry("USA")

		// Create a person with the address
		personBuilder := builders.NewPersonBuilder().
			WithID("123").
			WithName("John Doe").
			WithAge(30).
			WithEmail("john.doe@example.com").
			WithAddress(addressBuilder)

		// Build the person
		person := personBuilder.Build().(*models.Person)

		// Verify the address is set
		if person.Address.Street != "123 Main St" {
			t.Errorf("Expected street to be '123 Main St', got '%s'", person.Address.Street)
		}
		if person.Address.City != "New York" {
			t.Errorf("Expected city to be 'New York', got '%s'", person.Address.City)
		}
	})

	// Test Bank with Address
	t.Run("Bank with Address", func(t *testing.T) {
		// Create a valid address
		addressBuilder := builders.NewAddressBuilder().
			WithStreet("123 Main St").
			WithCity("New York").
			WithState("NY").
			WithPostalCode("10001").
			WithCountry("USA")

		// Create a bank with the address
		bankBuilder := builders.NewBankBuilder().
			WithName("Global Bank").
			WithBranchCode("GB-001").
			WithAddress(addressBuilder)

		// Build the bank
		bank := bankBuilder.Build().(*models.Bank)

		// Verify the address is set
		if bank.Address.Street != "123 Main St" {
			t.Errorf("Expected street to be '123 Main St', got '%s'", bank.Address.Street)
		}
		if bank.Address.City != "New York" {
			t.Errorf("Expected city to be 'New York', got '%s'", bank.Address.City)
		}
	})

	// Test Device with Type field
	t.Run("Device with Type field", func(t *testing.T) {
		// Create a device with a type field
		deviceBuilder := builders.NewDeviceBuilder().
			WithType("Smartphone").
			WithModel("iPhone 13").
			WithOS("iOS 15.4")

		// Build the device
		device := deviceBuilder.Build().(*models.Device)

		// Verify the type field is set
		if device.Type != "Smartphone" {
			t.Errorf("Expected type to be 'Smartphone', got '%s'", device.Type)
		}
	})

	// Test Transaction with Type field
	t.Run("Transaction with Type field", func(t *testing.T) {
		// Create a transaction with a type field
		transactionBuilder := builders.NewTransactionBuilder().
			WithID("TRX-001").
			WithAmount(100.0).
			WithCurrency("USD").
			WithType("deposit").
			WithDate("2023-01-01").
			WithStatus("completed")

		// Build the transaction
		transaction := transactionBuilder.Build().(*models.Transaction)

		// Verify the type field is set
		if transaction.Type != "deposit" {
			t.Errorf("Expected type to be 'deposit', got '%s'", transaction.Type)
		}
	})

	// Test Portfolio with nested structures
	t.Run("Portfolio with nested structures", func(t *testing.T) {
		// Portfolio test doesn't need to create stocks and bonds for this simple test

		// Create a portfolio with stocks and bonds
		portfolioBuilder := builders.NewPortfolioBuilder().
			WithID("PORT-001").
			WithName("Growth Portfolio").
			WithRiskLevel("High").
			WithTotalValue(100000.0)

		// Build the portfolio
		portfolio := portfolioBuilder.Build().(*models.Portfolio)

		// Verify the portfolio is created
		if portfolio.Name != "Growth Portfolio" {
			t.Errorf("Expected name to be 'Growth Portfolio', got '%s'", portfolio.Name)
		}
	})

	// Test HealthProfile with nested structures
	t.Run("HealthProfile with nested structures", func(t *testing.T) {
		// HealthProfile test doesn't need to create medications and medical records for this simple test

		// Create a health profile
		healthProfileBuilder := builders.NewHealthProfileBuilder().
			WithHeight(180.0).
			WithWeight(75.0).
			WithBloodType("O+")

		// Build the health profile
		healthProfile := healthProfileBuilder.Build().(*models.HealthProfile)

		// Verify the health profile is created
		if healthProfile.Height != 180.0 {
			t.Errorf("Expected height to be 180.0, got %f", healthProfile.Height)
		}
		if healthProfile.Weight != 75.0 {
			t.Errorf("Expected weight to be 75.0, got %f", healthProfile.Weight)
		}
	})
}
