package models_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

var _ = Describe("DigitalProfile", func() {
	var (
		digitalProfileBuilder *builders.DigitalProfileBuilder
		deviceBuilder         *builders.DeviceBuilder
		appBuilder            *builders.ApplicationBuilder
		accountBuilder        *builders.OnlineAccountBuilder
		subscriptionBuilder   *builders.SubscriptionBuilder
		preferencesBuilder    *builders.UserPreferencesBuilder
		activityBuilder       *builders.UserActivityBuilder
		geoLocationBuilder    *builders.GeoLocationBuilder
	)

	BeforeEach(func() {
		appBuilder = builders.NewApplicationBuilder().
			WithName("Social Media App").
			WithVersion("2.1.0").
			WithInstallDate("2022-01-15").
			WithLastUsed("2023-05-01").
			WithPermission("Camera").
			WithPermission("Microphone")

		deviceBuilder = builders.NewDeviceBuilder().
			WithType("Smartphone").
			WithModel("iPhone 13").
			WithSerialNumber("ABCD1234").
			WithPurchaseDate("2022-01-01").
			WithOS("iOS 16.5").
			WithLastUsed("2023-05-01").
			WithApp(appBuilder).
			WithSetting("Dark Mode", "Enabled")

		subscriptionBuilder = builders.NewSubscriptionBuilder().
			WithPlan("Premium").
			WithStartDate("2022-01-01").
			WithEndDate("2023-01-01").
			WithPrice(9.99).
			WithBillingCycle("Monthly").
			WithAutoRenew(true).
			WithFeature("Ad-free").
			WithFeature("Offline access")

		accountBuilder = builders.NewOnlineAccountBuilder().
			WithPlatform("Social Media").
			WithUsername("user123").
			WithEmail("user@example.com").
			WithCreationDate("2020-01-01").
			WithLastLogin("2023-05-01").
			WithStatus("Active").
			WithSetting("Privacy", "Friends only").
			WithSubscription(subscriptionBuilder)

		preferencesBuilder = builders.NewUserPreferencesBuilder().
			WithTheme("Dark").
			WithLanguage("English").
			WithNotification("Email", true).
			WithNotification("Push", true).
			WithNotification("SMS", false).
			WithPrivacySetting("Profile visibility", "Friends only").
			WithAccessibilitySetting("High contrast", false)

		geoLocationBuilder = builders.NewGeoLocationBuilder().
			WithLatitude(37.7749).
			WithLongitude(-122.4194).
			WithAccuracy(10.0)

		activityBuilder = builders.NewUserActivityBuilder().
			WithTimestamp("2023-05-01T12:34:56Z").
			WithType("Login").
			WithPlatform("Mobile").
			WithDevice("iPhone 13").
			WithLocation(geoLocationBuilder).
			WithDetail("IP", "192.168.1.1")

		digitalProfileBuilder = builders.NewDigitalProfileBuilder().
			WithDevice(deviceBuilder).
			WithAccount(accountBuilder).
			WithPreferences(preferencesBuilder).
			WithActivity(activityBuilder)
	})

	Describe("Positive Tests", func() {
		Context("when creating a digital profile with valid data", func() {
			It("should create a digital profile with basic information", func() {
				digitalProfile := digitalProfileBuilder.Build().(*models.DigitalProfile)

				Expect(len(digitalProfile.Devices)).To(Equal(1))
				Expect(len(digitalProfile.Accounts)).To(Equal(1))
				Expect(len(digitalProfile.Activity)).To(Equal(1))
			})

			It("should create a digital profile with device details", func() {
				digitalProfile := digitalProfileBuilder.Build().(*models.DigitalProfile)

				device := digitalProfile.Devices[0]
				Expect(device.Type).To(Equal("Smartphone"))
				Expect(device.Model).To(Equal("iPhone 13"))
				Expect(device.OS).To(Equal("iOS 16.5"))

				Expect(len(device.Apps)).To(Equal(1))
				Expect(device.Apps[0].Name).To(Equal("Social Media App"))
				Expect(len(device.Apps[0].Permissions)).To(Equal(2))
			})

			It("should create a digital profile with account details", func() {
				digitalProfile := digitalProfileBuilder.Build().(*models.DigitalProfile)

				account := digitalProfile.Accounts[0]
				Expect(account.Platform).To(Equal("Social Media"))
				Expect(account.Username).To(Equal("user123"))
				Expect(account.Status).To(Equal("Active"))

				Expect(account.Subscription).NotTo(BeNil())
				Expect(account.Subscription.Plan).To(Equal("Premium"))
				Expect(account.Subscription.Price).To(Equal(9.99))
				Expect(len(account.Subscription.Features)).To(Equal(2))
			})

			It("should create a digital profile with preferences", func() {
				digitalProfile := digitalProfileBuilder.Build().(*models.DigitalProfile)

				prefs := digitalProfile.Preferences
				Expect(prefs.Theme).To(Equal("Dark"))
				Expect(prefs.Language).To(Equal("English"))
				Expect(prefs.Notifications["Email"]).To(BeTrue())
				Expect(prefs.Notifications["Push"]).To(BeTrue())
				Expect(prefs.Notifications["SMS"]).To(BeFalse())
			})

			It("should create a digital profile with activity", func() {
				digitalProfile := digitalProfileBuilder.Build().(*models.DigitalProfile)

				activity := digitalProfile.Activity[0]
				Expect(activity.Type).To(Equal("Login"))
				Expect(activity.Platform).To(Equal("Mobile"))
				Expect(activity.Device).To(Equal("iPhone 13"))
				Expect(activity.Location.Latitude).To(Equal(37.7749))
				Expect(activity.Location.Longitude).To(Equal(-122.4194))
			})
		})
	})

	Describe("Negative Tests", func() {
		Context("when creating a digital profile with invalid data", func() {
			It("should handle invalid device data", func() {
				invalidDeviceBuilder := builders.NewDeviceBuilder().
					WithType("").  // Invalid: empty type
					WithModel(""). // Invalid: empty model
					WithOS("")     // Invalid: empty OS

				digitalProfile := digitalProfileBuilder.
					WithDevice(invalidDeviceBuilder).
					Build().(models.DigitalProfile)

				// If we had validation, this would fail
				// For now, just check the device data
				Expect(digitalProfile.Devices[1].Type).To(Equal(""))
				Expect(digitalProfile.Devices[1].Model).To(Equal(""))
				Expect(digitalProfile.Devices[1].OS).To(Equal(""))
			})

			It("should handle invalid account data", func() {
				invalidAccountBuilder := builders.NewOnlineAccountBuilder().
					WithPlatform("").         // Invalid: empty platform
					WithUsername("").         // Invalid: empty username
					WithEmail("not-an-email") // Invalid: bad email format

				digitalProfile := digitalProfileBuilder.
					WithAccount(invalidAccountBuilder).
					Build().(models.DigitalProfile)

				// If we had validation, this would fail
				// For now, just check the account data
				Expect(digitalProfile.Accounts[1].Platform).To(Equal(""))
				Expect(digitalProfile.Accounts[1].Username).To(Equal(""))
				Expect(digitalProfile.Accounts[1].Email).To(Equal("not-an-email"))
			})

			It("should handle invalid location data", func() {
				invalidLocationBuilder := builders.NewGeoLocationBuilder().
					WithLatitude(100.0).  // Invalid: > 90
					WithLongitude(200.0). // Invalid: > 180
					WithAccuracy(-5.0)    // Invalid: negative

				invalidActivityBuilder := builders.NewUserActivityBuilder().
					WithTimestamp("2023-05-01").
					WithType("Login").
					WithLocation(invalidLocationBuilder)

				digitalProfile := digitalProfileBuilder.
					WithActivity(invalidActivityBuilder).
					Build().(models.DigitalProfile)

				// If we had validation, this would fail
				// For now, just check the location data
				Expect(digitalProfile.Activity[1].Location.Latitude).To(Equal(100.0))
				Expect(digitalProfile.Activity[1].Location.Longitude).To(Equal(200.0))
				Expect(digitalProfile.Activity[1].Location.Accuracy).To(Equal(-5.0))
			})
		})
	})

	Describe("Edge Cases", func() {
		It("should handle a digital profile with no devices", func() {
			digitalProfile := builders.NewDigitalProfileBuilder().
				WithAccount(accountBuilder).
				Build().(models.DigitalProfile)

			Expect(len(digitalProfile.Devices)).To(Equal(0))
			Expect(len(digitalProfile.Accounts)).To(Equal(1))
		})

		It("should handle a digital profile with many devices", func() {
			manyDevicesBuilder := builders.NewDigitalProfileBuilder()

			// Add 5 devices
			for i := 1; i <= 5; i++ {
				deviceBuilder := builders.NewDeviceBuilder().
					WithType("Smartphone").
					WithModel(fmt.Sprintf("Phone %d", i)).
					WithOS("Android")

				manyDevicesBuilder.WithDevice(deviceBuilder)
			}

			digitalProfile := manyDevicesBuilder.Build().(*models.DigitalProfile)

			Expect(len(digitalProfile.Devices)).To(Equal(5))
			Expect(digitalProfile.Devices[0].Model).To(Equal("Phone 1"))
			Expect(digitalProfile.Devices[4].Model).To(Equal("Phone 5"))
		})

		It("should handle a device with many apps", func() {
			manyAppsDeviceBuilder := builders.NewDeviceBuilder().
				WithType("Smartphone").
				WithModel("App Phone")

			// Add 10 apps
			for i := 1; i <= 10; i++ {
				appBuilder := builders.NewApplicationBuilder().
					WithName(fmt.Sprintf("App %d", i)).
					WithVersion("1.0")

				manyAppsDeviceBuilder.WithApp(appBuilder)
			}

			digitalProfile := builders.NewDigitalProfileBuilder().
				WithDevice(manyAppsDeviceBuilder).
				Build().(models.DigitalProfile)

			Expect(len(digitalProfile.Devices[0].Apps)).To(Equal(10))
			Expect(digitalProfile.Devices[0].Apps[0].Name).To(Equal("App 1"))
			Expect(digitalProfile.Devices[0].Apps[9].Name).To(Equal("App 10"))
		})

		It("should handle an account with expired subscription", func() {
			expiredSubscriptionBuilder := builders.NewSubscriptionBuilder().
				WithPlan("Premium").
				WithStartDate("2020-01-01").
				WithEndDate("2020-12-31"). // Past date
				WithAutoRenew(false)

			accountWithExpiredSub := builders.NewOnlineAccountBuilder().
				WithPlatform("Streaming").
				WithUsername("expired_user").
				WithSubscription(expiredSubscriptionBuilder)

			digitalProfile := builders.NewDigitalProfileBuilder().
				WithAccount(accountWithExpiredSub).
				Build().(models.DigitalProfile)

			Expect(digitalProfile.Accounts[0].Subscription.EndDate).To(Equal("2020-12-31"))
			Expect(digitalProfile.Accounts[0].Subscription.AutoRenew).To(BeFalse())
		})
	})
})

