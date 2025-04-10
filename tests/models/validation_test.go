// +build ignore

// +build ignore

import (
	"github.com/adil-faiyaz98/go-builder-kit/builders"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validation", func() {
	Context("Address Validation", func() {
		It("should fail validation with empty required fields", func() {
			// Create an Address builder with missing required fields
			addressBuilder := builders.NewAddressBuilder()
			// Missing Street, City, and Country

			// Build and validate should fail
			_, err := addressBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Street cannot be empty"))
			Expect(err.Error()).To(ContainSubstring("City cannot be empty"))
			Expect(err.Error()).To(ContainSubstring("Country cannot be empty"))
		})

		It("should fail validation with invalid postal code", func() {
			// Create an Address builder with invalid postal code
			addressBuilder := builders.NewAddressBuilder()
			addressBuilder.
				WithStreet("123 Main St").
				WithCity("San Francisco").
				WithState("CA").
				WithPostalCode("94"). // Too short
				WithCountry("USA")

			// Build and validate should fail
			_, err := addressBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("PostalCode format is invalid"))
		})

		It("should fail validation with invalid address type", func() {
			// Create an Address builder with invalid type
			addressBuilder := builders.NewAddressBuilder()
			addressBuilder.
				WithStreet("123 Main St").
				WithCity("San Francisco").
				WithState("CA").
				WithPostalCode("94105").
				WithCountry("USA").
				WithType("Invalid") // Not a valid type

			// Build and validate should fail
			_, err := addressBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Type must be one of"))
		})
	})

	Context("GeoLocation Validation", func() {
		It("should fail validation with invalid latitude", func() {
			// Create a GeoLocation builder with invalid latitude
			geoLocationBuilder := builders.NewGeoLocationBuilder()
			geoLocationBuilder.
				WithLatitude(100.0). // Greater than 90
				WithLongitude(-122.4194).
				WithAccuracy(10.0)

			// Build and validate should fail
			_, err := geoLocationBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Latitude must be between -90 and 90"))
		})

		It("should fail validation with invalid longitude", func() {
			// Create a GeoLocation builder with invalid longitude
			geoLocationBuilder := builders.NewGeoLocationBuilder()
			geoLocationBuilder.
				WithLatitude(37.7749).
				WithLongitude(200.0). // Greater than 180
				WithAccuracy(10.0)

			// Build and validate should fail
			_, err := geoLocationBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Longitude must be between -180 and 180"))
		})

		It("should fail validation with negative accuracy", func() {
			// Create a GeoLocation builder with negative accuracy
			geoLocationBuilder := builders.NewGeoLocationBuilder()
			geoLocationBuilder.
				WithLatitude(37.7749).
				WithLongitude(-122.4194).
				WithAccuracy(-10.0) // Negative

			// Build and validate should fail
			_, err := geoLocationBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Accuracy cannot be negative"))
		})
	})

	Context("Person Validation", func() {
		It("should fail validation with empty required fields", func() {
			// Create a Person builder with missing required fields
			personBuilder := builders.NewPersonBuilder()
			// Missing ID and Name

			// Build and validate should fail
			_, err := personBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("ID cannot be empty"))
			Expect(err.Error()).To(ContainSubstring("Name cannot be empty"))
		})

		It("should fail validation with invalid email", func() {
			// Create a Person builder with invalid email
			personBuilder := builders.NewPersonBuilder()
			personBuilder.
				WithID("P12345").
				WithName("John Doe").
				WithEmail("invalid-email") // Not a valid email

			// Build and validate should fail
			_, err := personBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Email is not valid"))
		})

		It("should fail validation with invalid phone", func() {
			// Create a Person builder with invalid phone
			personBuilder := builders.NewPersonBuilder()
			personBuilder.
				WithID("P12345").
				WithName("John Doe").
				WithPhone("123") // Too short

			// Build and validate should fail
			_, err := personBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Phone number is not valid"))
		})

		It("should fail validation with invalid birthdate", func() {
			// Create a Person builder with invalid birthdate
			personBuilder := builders.NewPersonBuilder()
			personBuilder.
				WithID("P12345").
				WithName("John Doe").
				WithBirthdate("2023/01/01") // Wrong format

			// Build and validate should fail
			_, err := personBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Birthdate must be in the format YYYY-MM-DD"))
		})

		It("should fail validation with future birthdate", func() {
			// Create a Person builder with future birthdate
			personBuilder := builders.NewPersonBuilder()
			personBuilder.
				WithID("P12345").
				WithName("John Doe").
				WithBirthdate("2100-01-01") // Future date

			// Build and validate should fail
			_, err := personBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Birthdate cannot be in the future"))
		})

		It("should fail validation with invalid gender", func() {
			// Create a Person builder with invalid gender
			personBuilder := builders.NewPersonBuilder()
			personBuilder.
				WithID("P12345").
				WithName("John Doe").
				WithGender("invalid") // Not a valid gender

			// Build and validate should fail
			_, err := personBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Gender must be one of"))
		})
	})

	Context("Bond Validation", func() {
		It("should fail validation with empty required fields", func() {
			// Create a Bond builder with missing required fields
			bondBuilder := builders.NewBondBuilder()
			// Missing ISIN, Name, Issuer, etc.

			// Build and validate should fail
			_, err := bondBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("ISIN is required"))
			Expect(err.Error()).To(ContainSubstring("Name is required"))
			Expect(err.Error()).To(ContainSubstring("Issuer is required"))
		})

		It("should fail validation with invalid ISIN", func() {
			// Create a Bond builder with invalid ISIN
			bondBuilder := builders.NewBondBuilder()
			bondBuilder.
				WithISIN("INVALID"). // Not a valid ISIN
				WithName("US Treasury Bond").
				WithIssuer("US Treasury").
				WithType("government").
				WithFaceValue(1000.0).
				WithCouponRate(2.5).
				WithMaturityDate("2030-12-31").
				WithPurchaseDate("2023-01-15").
				WithPurchasePrice(980.0).
				WithCurrentPrice(985.0).
				WithQuantity(10).
				WithCurrency("USD").
				WithPaymentFrequency("semi-annual")

			// Build and validate should fail
			_, err := bondBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("ISIN format is invalid"))
		})

		It("should fail validation with invalid bond type", func() {
			// Create a Bond builder with invalid type
			bondBuilder := builders.NewBondBuilder()
			bondBuilder.
				WithISIN("US123456AB12").
				WithName("US Treasury Bond").
				WithIssuer("US Treasury").
				WithType("invalid"). // Not a valid type
				WithFaceValue(1000.0).
				WithCouponRate(2.5).
				WithMaturityDate("2030-12-31").
				WithPurchaseDate("2023-01-15").
				WithPurchasePrice(980.0).
				WithCurrentPrice(985.0).
				WithQuantity(10).
				WithCurrency("USD").
				WithPaymentFrequency("semi-annual")

			// Build and validate should fail
			_, err := bondBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Type must be one of"))
		})

		It("should fail validation with negative face value", func() {
			// Create a Bond builder with negative face value
			bondBuilder := builders.NewBondBuilder()
			bondBuilder.
				WithISIN("US123456AB12").
				WithName("US Treasury Bond").
				WithIssuer("US Treasury").
				WithType("government").
				WithFaceValue(-1000.0). // Negative
				WithCouponRate(2.5).
				WithMaturityDate("2030-12-31").
				WithPurchaseDate("2023-01-15").
				WithPurchasePrice(980.0).
				WithCurrentPrice(985.0).
				WithQuantity(10).
				WithCurrency("USD").
				WithPaymentFrequency("semi-annual")

			// Build and validate should fail
			_, err := bondBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("FaceValue must be greater than 0"))
		})

		It("should fail validation with maturity date before purchase date", func() {
			// Create a Bond builder with maturity date before purchase date
			bondBuilder := builders.NewBondBuilder()
			bondBuilder.
				WithISIN("US123456AB12").
				WithName("US Treasury Bond").
				WithIssuer("US Treasury").
				WithType("government").
				WithFaceValue(1000.0).
				WithCouponRate(2.5).
				WithMaturityDate("2020-12-31"). // Before purchase date
				WithPurchaseDate("2023-01-15").
				WithPurchasePrice(980.0).
				WithCurrentPrice(985.0).
				WithQuantity(10).
				WithCurrency("USD").
				WithPaymentFrequency("semi-annual")

			// Build and validate should fail
			_, err := bondBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("MaturityDate cannot be before PurchaseDate"))
		})
	})
})
