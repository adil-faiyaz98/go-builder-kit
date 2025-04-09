package models_test

import (
	"testing"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

// TestNestedStructuresSimple tests that nested structures are properly handled
func TestNestedStructuresSimple(t *testing.T) {
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

	// Test Bank with Account
	t.Run("Bank with Account", func(t *testing.T) {
		// Create a valid address
		addressBuilder := builders.NewAddressBuilder().
			WithStreet("123 Main St").
			WithCity("New York").
			WithState("NY").
			WithPostalCode("10001").
			WithCountry("USA")

		// Create a bank
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

	// Test Portfolio with Stock
	t.Run("Portfolio with Stock", func(t *testing.T) {
		// Create a portfolio
		portfolioBuilder := builders.NewPortfolioBuilder().
			WithID("PORT-001").
			WithName("Growth Portfolio").
			WithRiskLevel("High").
			WithTotalValue(100000.0)

		// Build the portfolio
		portfolio := portfolioBuilder.Build().(*models.Portfolio)

		// Verify the portfolio is set
		if portfolio.Name != "Growth Portfolio" {
			t.Errorf("Expected name to be 'Growth Portfolio', got '%s'", portfolio.Name)
		}
		if portfolio.RiskLevel != "High" {
			t.Errorf("Expected risk level to be 'High', got '%s'", portfolio.RiskLevel)
		}
	})

	// Test DigitalProfile with Device
	t.Run("DigitalProfile with Device", func(t *testing.T) {
		// Create a device
		deviceBuilder := builders.NewDeviceBuilder().
			WithType("Smartphone").
			WithModel("iPhone 13").
			WithManufacturer("Apple").
			WithSerialNumber("ABCD1234").
			WithPurchaseDate("2022-01-01").
			WithOperatingSystem("iOS").
			WithOSVersion("15.4")

		// Create a digital profile
		digitalProfileBuilder := builders.NewDigitalProfileBuilder()

		// Build the digital profile
		digitalProfile := digitalProfileBuilder.Build().(*models.DigitalProfile)

		// Verify the digital profile is created
		if digitalProfile == nil {
			t.Errorf("Expected digital profile to be created")
		}
	})

	// Test UserActivity with GeoLocation
	t.Run("UserActivity with GeoLocation", func(t *testing.T) {
		// Create a geo location
		geoLocationBuilder := builders.NewGeoLocationBuilder().
			WithLatitude(37.7749).
			WithLongitude(-122.4194).
			WithAccuracy(5.0)

		// Create a user activity
		userActivityBuilder := builders.NewUserActivityBuilder().
			WithType("Login").
			WithPlatform("Mobile").
			WithDate("2022-04-01").
			WithTimestamp("2022-04-01T14:30:00Z").
			WithDuration(15).
			WithDescription("User logged in from mobile app").
			WithDevice("iPhone 13").
			WithLocation(geoLocationBuilder)

		// Build the user activity
		userActivity := userActivityBuilder.Build().(*models.UserActivity)

		// Verify the geo location is set
		if userActivity.Location == nil {
			t.Errorf("Expected location to be set")
		} else {
			if userActivity.Location.Latitude != 37.7749 {
				t.Errorf("Expected latitude to be 37.7749, got %f", userActivity.Location.Latitude)
			}
			if userActivity.Location.Longitude != -122.4194 {
				t.Errorf("Expected longitude to be -122.4194, got %f", userActivity.Location.Longitude)
			}
		}
	})
}
