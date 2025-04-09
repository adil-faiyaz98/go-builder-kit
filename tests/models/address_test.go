package models_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

var _ = Describe("Address", func() {
	var (
		addressBuilder *builders.AddressBuilder
	)

	BeforeEach(func() {
		addressBuilder = builders.NewAddressBuilder().
			WithStreet("123 Main St").
			WithCity("New York").
			WithState("NY").
			WithPostalCode("10001").
			WithCountry("USA").
			WithType("Home").
			WithIsPrimary(true)
	})

	Describe("Validation", func() {
		Context("when all fields are valid", func() {
			It("should pass validation", func() {
				address := addressBuilder.Build().(*models.Address)
				err := address.Validate()
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when Street is empty", func() {
			It("should fail validation", func() {
				address := addressBuilder.WithStreet("").Build().(*models.Address)
				err := address.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Street cannot be empty"))
			})
		})

		Context("when City is empty", func() {
			It("should fail validation", func() {
				address := addressBuilder.WithCity("").Build().(*models.Address)
				err := address.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("City cannot be empty"))
			})
		})

		Context("when Country is empty", func() {
			It("should fail validation", func() {
				address := addressBuilder.WithCountry("").Build().(*models.Address)
				err := address.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Country cannot be empty"))
			})
		})

		Context("when PostalCode is invalid", func() {
			It("should fail validation", func() {
				address := addressBuilder.WithPostalCode("!@#$%").Build().(*models.Address)
				err := address.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("PostalCode format is invalid"))
			})
		})

		Context("when Type is invalid", func() {
			It("should fail validation", func() {
				address := addressBuilder.WithType("Unknown").Build().(*models.Address)
				err := address.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Type must be one of: Home, Work, Mailing, Billing, Other"))
			})
		})

		Context("when Coordinates are invalid", func() {
			It("should fail validation with invalid latitude", func() {
				geoLocationBuilder := builders.NewGeoLocationBuilder().
					WithLatitude(100). // Invalid: > 90
					WithLongitude(0).
					WithAccuracy(10)

				address := addressBuilder.WithCoordinates(geoLocationBuilder).Build().(*models.Address)
				err := address.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Latitude must be between -90 and 90"))
			})

			It("should fail validation with invalid longitude", func() {
				geoLocationBuilder := builders.NewGeoLocationBuilder().
					WithLatitude(0).
					WithLongitude(200). // Invalid: > 180
					WithAccuracy(10)

				address := addressBuilder.WithCoordinates(geoLocationBuilder).Build().(*models.Address)
				err := address.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Longitude must be between -180 and 180"))
			})

			It("should fail validation with negative accuracy", func() {
				geoLocationBuilder := builders.NewGeoLocationBuilder().
					WithLatitude(0).
					WithLongitude(0).
					WithAccuracy(-5) // Invalid: negative

				address := addressBuilder.WithCoordinates(geoLocationBuilder).Build().(*models.Address)
				err := address.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Accuracy cannot be negative"))
			})
		})
	})
})

