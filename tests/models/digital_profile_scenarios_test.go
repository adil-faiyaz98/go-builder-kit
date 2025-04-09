package models_test

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

var _ = Describe("DigitalProfile Scenarios", func() {
	// Helper function to create a valid device builder
	createValidDeviceBuilder := func() *builders.DeviceBuilder {
		return builders.NewDeviceBuilder().
			WithType("Smartphone").
			WithModel("iPhone 13").
			WithManufacturer("Apple").
			WithSerialNumber("ABCD1234XYZ").
			WithPurchaseDate("2022-01-15").
			WithOperatingSystem("iOS").
			WithOSVersion("15.4")
	}

	// Helper function to create a valid application builder
	createValidApplicationBuilder := func() *builders.ApplicationBuilder {
		return builders.NewApplicationBuilder().
			WithName("Social Media App").
			WithVersion("2.3.1").
			WithDeveloper("Tech Corp").
			WithCategory("Social").
			WithInstallDate("2022-02-01").
			WithLastUpdated("2022-03-15").
			WithSize(45.6)
	}

	// Helper function to create a valid online account builder
	createValidOnlineAccountBuilder := func() *builders.OnlineAccountBuilder {
		return builders.NewOnlineAccountBuilder().
			WithPlatform("Instagram").
			WithUsername("johndoe").
			WithEmail("john.doe@example.com").
			WithDateCreated("2020-05-10").
			WithLastLogin("2022-04-01").
			WithStatus("active").
			WithSubscription(
				builders.NewSubscriptionBuilder().
					WithPlan("Premium").
					WithStartDate("2022-01-01").
					WithEndDate("2023-01-01").
					WithPrice(9.99).
					WithBillingCycle("monthly").
					WithAutoRenew(true).
					WithStatus("active").
					WithFeatures("HD streaming").
					WithFeatures("Ad-free experience"),
			)
	}

	// Helper function to create a valid geo location builder
	createValidGeoLocationBuilder := func() *builders.GeoLocationBuilder {
		return builders.NewGeoLocationBuilder().
			WithLatitude(37.7749).
			WithLongitude(-122.4194).
			WithAccuracy(5.0)
	}

	// Helper function to create a valid user activity builder
	createValidUserActivityBuilder := func() *builders.UserActivityBuilder {
		return builders.NewUserActivityBuilder().
			WithType("Login").
			WithPlatform("Mobile").
			WithDate("2022-04-01").
			WithTimestamp("2022-04-01T14:30:00Z").
			WithDuration(15).
			WithDescription("User logged in from mobile app").
			WithDevice("iPhone 13").
			WithLocation(createValidGeoLocationBuilder()).
			WithDetails("ip_address", "192.168.1.1").
			WithDetails("browser", "Safari Mobile")
	}

	// Helper function to create a valid user preferences builder
	createValidUserPreferencesBuilder := func() *builders.UserPreferencesBuilder {
		return builders.NewUserPreferencesBuilder().
			WithTheme("dark").
			WithLanguage("en-US").
			WithNotificationsEnabled(true).
			WithPrivacySettings("public_profile", true).
			WithPrivacySettings("show_activity", false)
	}

	Describe("Valid Scenarios", func() {
		Context("when creating a complete digital profile with all valid data", func() {
			It("should build and validate successfully", func() {
				// Create a digital profile with all valid data
				digitalProfileBuilder := builders.NewDigitalProfileBuilder().
					WithDevice(createValidDeviceBuilder()).
					WithAccount(createValidOnlineAccountBuilder()).
					WithPreferences(createValidUserPreferencesBuilder()).
					WithActivity(createValidUserActivityBuilder())

				// Build the digital profile
				digitalProfile := digitalProfileBuilder.Build().(*models.DigitalProfile)

				// Validate the digital profile
				err := digitalProfile.Validate()
				Expect(err).NotTo(HaveOccurred(), "Expected validation to pass for a valid digital profile")

				// Verify the structure is complete
				Expect(digitalProfile.Devices).To(HaveLen(1), "Expected 1 device")
				Expect(digitalProfile.Accounts).To(HaveLen(1), "Expected 1 account")
				Expect(digitalProfile.Activity).To(HaveLen(1), "Expected 1 activity")
				Expect(digitalProfile.Preferences).NotTo(BeNil(), "Expected preferences to be set")

				// Verify nested structures
				device := digitalProfile.Devices[0]
				Expect(device.Type).To(Equal("Smartphone"), "Expected device type to be 'Smartphone'")
				Expect(device.Model).To(Equal("iPhone 13"), "Expected device model to be 'iPhone 13'")

				account := digitalProfile.Accounts[0]
				Expect(account.Platform).To(Equal("Instagram"), "Expected platform to be 'Instagram'")
				Expect(account.Status).To(Equal("active"), "Expected status to be 'active'")
				Expect(account.Subscription).NotTo(BeNil(), "Expected subscription to be set")
				Expect(account.Subscription.Plan).To(Equal("Premium"), "Expected subscription plan to be 'Premium'")
				Expect(account.Subscription.Features).To(ContainElements("HD streaming", "Ad-free experience"), 
					"Expected subscription features to include 'HD streaming' and 'Ad-free experience'")

				activity := digitalProfile.Activity[0]
				Expect(activity.Type).To(Equal("Login"), "Expected activity type to be 'Login'")
				Expect(activity.Platform).To(Equal("Mobile"), "Expected activity platform to be 'Mobile'")
				Expect(activity.Location).NotTo(BeNil(), "Expected location to be set")
				Expect(activity.Location.Latitude).To(Equal(37.7749), "Expected latitude to be 37.7749")
				Expect(activity.Location.Longitude).To(Equal(-122.4194), "Expected longitude to be -122.4194")
				Expect(activity.Details).To(HaveKeyWithValue("ip_address", "192.168.1.1"), 
					"Expected details to include IP address")

				preferences := digitalProfile.Preferences
				Expect(preferences.Theme).To(Equal("dark"), "Expected theme to be 'dark'")
				Expect(preferences.NotificationsEnabled).To(BeTrue(), "Expected notifications to be enabled")
				Expect(preferences.PrivacySettings).To(HaveKeyWithValue("public_profile", true), 
					"Expected privacy settings to include public profile setting")
			})
		})

		Context("when creating a digital profile with minimal valid data", func() {
			It("should build and validate successfully", func() {
				// Create a digital profile with minimal data
				digitalProfileBuilder := builders.NewDigitalProfileBuilder().
					WithDevice(createValidDeviceBuilder())

				// Build the digital profile
				digitalProfile := digitalProfileBuilder.Build().(*models.DigitalProfile)

				// Validate the digital profile
				err := digitalProfile.Validate()
				Expect(err).NotTo(HaveOccurred(), "Expected validation to pass for a minimal digital profile")

				// Verify the structure
				Expect(digitalProfile.Devices).To(HaveLen(1), "Expected 1 device")
				Expect(digitalProfile.Accounts).To(BeEmpty(), "Expected no accounts")
				Expect(digitalProfile.Activity).To(BeEmpty(), "Expected no activity")
				Expect(digitalProfile.Preferences).To(BeNil(), "Expected no preferences")
			})
		})
	})

	Describe("Invalid Scenarios", func() {
		Context("when creating a digital profile with invalid device data", func() {
			It("should fail validation with appropriate error messages", func() {
				// Create a device with invalid data
				invalidDeviceBuilder := builders.NewDeviceBuilder().
					WithType("InvalidType").  // Invalid type
					WithModel("").            // Empty model
					WithManufacturer("").     // Empty manufacturer
					WithPurchaseDate("2025-01-01") // Future date

				// Create a digital profile with the invalid device
				digitalProfileBuilder := builders.NewDigitalProfileBuilder().
					WithDevice(invalidDeviceBuilder)

				// Build the digital profile
				digitalProfile := digitalProfileBuilder.Build().(*models.DigitalProfile)

				// Validate the digital profile
				err := digitalProfile.Validate()
				Expect(err).To(HaveOccurred(), "Expected validation to fail for invalid device data")

				// Check error messages
				errString := err.Error()
				Expect(errString).To(ContainSubstring("Type must be one of"), 
					"Expected error about invalid device type")
				Expect(errString).To(ContainSubstring("Model is required"), 
					"Expected error about empty model")
				Expect(errString).To(ContainSubstring("Manufacturer is required"), 
					"Expected error about empty manufacturer")
				Expect(errString).To(ContainSubstring("PurchaseDate cannot be in the future"), 
					"Expected error about future purchase date")
			})
		})

		Context("when creating a digital profile with invalid account data", func() {
			It("should fail validation with appropriate error messages", func() {
				// Create an account with invalid data
				invalidAccountBuilder := builders.NewOnlineAccountBuilder().
					WithPlatform("").         // Empty platform
					WithUsername("").         // Empty username
					WithEmail("invalid-email") // Invalid email format

				// Create a digital profile with the invalid account
				digitalProfileBuilder := builders.NewDigitalProfileBuilder().
					WithDevice(createValidDeviceBuilder()).
					WithAccount(invalidAccountBuilder)

				// Build the digital profile
				digitalProfile := digitalProfileBuilder.Build().(*models.DigitalProfile)

				// Validate the digital profile
				err := digitalProfile.Validate()
				Expect(err).To(HaveOccurred(), "Expected validation to fail for invalid account data")

				// Check error messages
				errString := err.Error()
				Expect(errString).To(ContainSubstring("Platform is required"), 
					"Expected error about empty platform")
				Expect(errString).To(ContainSubstring("Username is required"), 
					"Expected error about empty username")
				Expect(errString).To(ContainSubstring("Email format is invalid"), 
					"Expected error about invalid email format")
			})
		})

		Context("when creating a digital profile with invalid location data", func() {
			It("should fail validation with appropriate error messages", func() {
				// Create a user activity with invalid location data
				invalidLocationBuilder := builders.NewGeoLocationBuilder().
					WithLatitude(100.0).    // Invalid latitude (out of range)
					WithLongitude(200.0).   // Invalid longitude (out of range)
					WithAccuracy(-5.0)      // Negative accuracy

				invalidActivityBuilder := builders.NewUserActivityBuilder().
					WithType("Login").
					WithPlatform("Mobile").
					WithDate("2022-04-01").
					WithLocation(invalidLocationBuilder)

				// Create a digital profile with the invalid activity
				digitalProfileBuilder := builders.NewDigitalProfileBuilder().
					WithDevice(createValidDeviceBuilder()).
					WithActivity(invalidActivityBuilder)

				// Build the digital profile
				digitalProfile := digitalProfileBuilder.Build().(*models.DigitalProfile)

				// Validate the digital profile
				err := digitalProfile.Validate()
				Expect(err).To(HaveOccurred(), "Expected validation to fail for invalid location data")

				// Check error messages
				errString := err.Error()
				Expect(errString).To(ContainSubstring("Latitude must be between -90 and 90"), 
					"Expected error about invalid latitude")
				Expect(errString).To(ContainSubstring("Longitude must be between -180 and 180"), 
					"Expected error about invalid longitude")
				Expect(errString).To(ContainSubstring("Accuracy cannot be negative"), 
					"Expected error about negative accuracy")
			})
		})

		Context("when creating a digital profile with invalid subscription data", func() {
			It("should fail validation with appropriate error messages", func() {
				// Create an account with invalid subscription data
				invalidSubscriptionBuilder := builders.NewSubscriptionBuilder().
					WithPlan("").                // Empty plan
					WithStartDate("2023-01-01"). // Start date after end date
					WithEndDate("2022-01-01").   // End date before start date
					WithPrice(-9.99).            // Negative price
					WithBillingCycle("invalid"). // Invalid billing cycle
					WithStatus("invalid")        // Invalid status

				invalidAccountBuilder := builders.NewOnlineAccountBuilder().
					WithPlatform("Instagram").
					WithUsername("johndoe").
					WithEmail("john.doe@example.com").
					WithSubscription(invalidSubscriptionBuilder)

				// Create a digital profile with the invalid account
				digitalProfileBuilder := builders.NewDigitalProfileBuilder().
					WithDevice(createValidDeviceBuilder()).
					WithAccount(invalidAccountBuilder)

				// Build the digital profile
				digitalProfile := digitalProfileBuilder.Build().(*models.DigitalProfile)

				// Validate the digital profile
				err := digitalProfile.Validate()
				Expect(err).To(HaveOccurred(), "Expected validation to fail for invalid subscription data")

				// Check error messages
				errString := err.Error()
				Expect(errString).To(ContainSubstring("Plan is required"), 
					"Expected error about empty plan")
				Expect(errString).To(ContainSubstring("EndDate cannot be before StartDate"), 
					"Expected error about end date before start date")
				Expect(errString).To(ContainSubstring("Price cannot be negative"), 
					"Expected error about negative price")
				Expect(errString).To(ContainSubstring("BillingCycle must be one of"), 
					"Expected error about invalid billing cycle")
				Expect(errString).To(ContainSubstring("Status must be one of"), 
					"Expected error about invalid status")
			})
		})
	})

	Describe("Edge Cases", func() {
		Context("when creating a digital profile with multiple devices", func() {
			It("should handle multiple devices correctly", func() {
				// Create a digital profile with multiple devices
				digitalProfileBuilder := builders.NewDigitalProfileBuilder().
					WithDevice(createValidDeviceBuilder().WithModel("iPhone 13")).
					WithDevice(createValidDeviceBuilder().WithModel("MacBook Pro")).
					WithDevice(createValidDeviceBuilder().WithModel("iPad Pro"))

				// Build the digital profile
				digitalProfile := digitalProfileBuilder.Build().(*models.DigitalProfile)

				// Validate the digital profile
				err := digitalProfile.Validate()
				Expect(err).NotTo(HaveOccurred(), "Expected validation to pass for multiple devices")

				// Verify the devices
				Expect(digitalProfile.Devices).To(HaveLen(3), "Expected 3 devices")
				
				deviceModels := []string{}
				for _, device := range digitalProfile.Devices {
					deviceModels = append(deviceModels, device.Model)
				}
				
				Expect(deviceModels).To(ContainElements("iPhone 13", "MacBook Pro", "iPad Pro"), 
					"Expected devices to include all three models")
			})
		})

		Context("when creating a digital profile with a device that has multiple applications", func() {
			It("should handle multiple applications correctly", func() {
				// Create a device with multiple applications
				deviceBuilder := createValidDeviceBuilder()
				
				// Add applications to the device
				deviceBuilder.
					WithApplication(createValidApplicationBuilder().WithName("Social Media App")).
					WithApplication(createValidApplicationBuilder().WithName("Messaging App")).
					WithApplication(createValidApplicationBuilder().WithName("Gaming App")).
					WithApplication(createValidApplicationBuilder().WithName("Productivity App")).
					WithApplication(createValidApplicationBuilder().WithName("Photo Editing App"))

				// Create a digital profile with the device
				digitalProfileBuilder := builders.NewDigitalProfileBuilder().
					WithDevice(deviceBuilder)

				// Build the digital profile
				digitalProfile := digitalProfileBuilder.Build().(*models.DigitalProfile)

				// Validate the digital profile
				err := digitalProfile.Validate()
				Expect(err).NotTo(HaveOccurred(), "Expected validation to pass for device with multiple applications")

				// Verify the applications
				device := digitalProfile.Devices[0]
				Expect(device.Applications).To(HaveLen(5), "Expected 5 applications")
				
				appNames := []string{}
				for _, app := range device.Applications {
					appNames = append(appNames, app.Name)
				}
				
				Expect(appNames).To(ContainElements(
					"Social Media App", 
					"Messaging App", 
					"Gaming App", 
					"Productivity App", 
					"Photo Editing App"), 
					"Expected applications to include all five names")
			})
		})

		Context("when creating a digital profile with an expired subscription", func() {
			It("should handle expired subscriptions correctly", func() {
				// Get current date and format it
				now := time.Now()
				startDate := now.AddDate(-1, 0, 0).Format("2006-01-02") // 1 year ago
				endDate := now.AddDate(0, -1, 0).Format("2006-01-02")   // 1 month ago
				
				// Create an account with an expired subscription
				accountBuilder := createValidOnlineAccountBuilder()
				accountBuilder.WithSubscription(
					builders.NewSubscriptionBuilder().
						WithPlan("Premium").
						WithStartDate(startDate).
						WithEndDate(endDate).
						WithPrice(9.99).
						WithBillingCycle("monthly").
						WithAutoRenew(false).
						WithStatus("expired")
				)

				// Create a digital profile with the account
				digitalProfileBuilder := builders.NewDigitalProfileBuilder().
					WithDevice(createValidDeviceBuilder()).
					WithAccount(accountBuilder)

				// Build the digital profile
				digitalProfile := digitalProfileBuilder.Build().(*models.DigitalProfile)

				// Validate the digital profile
				err := digitalProfile.Validate()
				Expect(err).NotTo(HaveOccurred(), "Expected validation to pass for expired subscription")

				// Verify the subscription
				account := digitalProfile.Accounts[0]
				subscription := account.Subscription
				Expect(subscription).NotTo(BeNil(), "Expected subscription to be set")
				Expect(subscription.Status).To(Equal("expired"), "Expected subscription status to be 'expired'")
				Expect(subscription.AutoRenew).To(BeFalse(), "Expected auto-renew to be false")
				
				// Parse the end date to verify it's in the past
				parsedEndDate, err := time.Parse("2006-01-02", subscription.EndDate)
				Expect(err).NotTo(HaveOccurred(), "Expected end date to be parseable")
				Expect(parsedEndDate.Before(now)).To(BeTrue(), "Expected end date to be in the past")
			})
		})

		Context("when creating a digital profile with extreme values", func() {
			It("should handle extreme values correctly", func() {
				// Create a user activity with a very long duration
				activityBuilder := createValidUserActivityBuilder().
					WithDuration(1000000) // Very long duration (in minutes)

				// Create a device with extreme values
				deviceBuilder := createValidDeviceBuilder().
					WithOSVersion("999.999.999") // Extreme version number

				// Create an account with extreme values
				accountBuilder := createValidOnlineAccountBuilder().
					WithSubscription(
						builders.NewSubscriptionBuilder().
							WithPlan("Ultra Premium Platinum Diamond Elite").
							WithPrice(9999.99) // Very high price
					)

				// Create a digital profile with extreme values
				digitalProfileBuilder := builders.NewDigitalProfileBuilder().
					WithDevice(deviceBuilder).
					WithAccount(accountBuilder).
					WithActivity(activityBuilder)

				// Build the digital profile
				digitalProfile := digitalProfileBuilder.Build().(*models.DigitalProfile)

				// Validate the digital profile
				err := digitalProfile.Validate()
				Expect(err).NotTo(HaveOccurred(), "Expected validation to pass for extreme values")

				// Verify the extreme values
				device := digitalProfile.Devices[0]
				Expect(device.OSVersion).To(Equal("999.999.999"), "Expected extreme OS version")

				activity := digitalProfile.Activity[0]
				Expect(activity.Duration).To(Equal(1000000), "Expected extreme duration")

				account := digitalProfile.Accounts[0]
				subscription := account.Subscription
				Expect(subscription.Plan).To(Equal("Ultra Premium Platinum Diamond Elite"), "Expected extreme plan name")
				Expect(subscription.Price).To(Equal(9999.99), "Expected extreme price")
			})
		})
	})

	Describe("Performance", func() {
		Context("when creating a digital profile with many nested objects", func() {
			It("should handle creating a complex digital profile efficiently", func() {
				// Create a digital profile with many nested objects
				digitalProfileBuilder := builders.NewDigitalProfileBuilder()

				// Add multiple devices
				for i := 0; i < 10; i++ {
					deviceBuilder := createValidDeviceBuilder().
						WithModel(fmt.Sprintf("Device %d", i))

					// Add multiple applications to each device
					for j := 0; j < 5; j++ {
						deviceBuilder.WithApplication(
							createValidApplicationBuilder().
								WithName(fmt.Sprintf("App %d-%d", i, j))
						)
					}

					digitalProfileBuilder.WithDevice(deviceBuilder)
				}

				// Add multiple accounts
				for i := 0; i < 10; i++ {
					digitalProfileBuilder.WithAccount(
						createValidOnlineAccountBuilder().
							WithPlatform(fmt.Sprintf("Platform %d", i)).
							WithUsername(fmt.Sprintf("user%d", i))
					)
				}

				// Add multiple activities
				for i := 0; i < 10; i++ {
					digitalProfileBuilder.WithActivity(
						createValidUserActivityBuilder().
							WithType(fmt.Sprintf("Activity %d", i))
					)
				}

				// Build the digital profile
				startTime := time.Now()
				digitalProfile := digitalProfileBuilder.Build().(*models.DigitalProfile)
				buildTime := time.Since(startTime)

				// Validate the digital profile
				startTime = time.Now()
				err := digitalProfile.Validate()
				validateTime := time.Since(startTime)

				// Verify performance
				Expect(err).NotTo(HaveOccurred(), "Expected validation to pass for complex digital profile")
				Expect(buildTime).To(BeNumerically("<", 100*time.Millisecond), 
					"Expected build time to be less than 100ms")
				Expect(validateTime).To(BeNumerically("<", 100*time.Millisecond), 
					"Expected validation time to be less than 100ms")

				// Verify the structure
				Expect(digitalProfile.Devices).To(HaveLen(10), "Expected 10 devices")
				Expect(digitalProfile.Accounts).To(HaveLen(10), "Expected 10 accounts")
				Expect(digitalProfile.Activity).To(HaveLen(10), "Expected 10 activities")

				// Verify nested structures
				for i, device := range digitalProfile.Devices {
					Expect(device.Model).To(Equal(fmt.Sprintf("Device %d", i)), 
						"Expected device model to match")
					Expect(device.Applications).To(HaveLen(5), 
						"Expected 5 applications per device")
				}
			})
		})
	})
})
