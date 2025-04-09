package tests

import (
	"fmt"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
	"github.com/adil-faiyaz98/go-builder-kit/models"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GeoLocation", func() {
	Context("GeoLocation Builder", func() {
		It("should build a valid GeoLocation with all fields", func() {
			// Create a GeoLocation builder
			geoLocationBuilder := builders.NewGeoLocationBuilder().
				WithLatitude(37.7749).
				WithLongitude(-122.4194).
				WithAccuracy(10.0)

			// Build the GeoLocation
			geoLocation, err := geoLocationBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(geoLocation).NotTo(BeNil())
			Expect(geoLocation.Latitude).To(Equal(37.7749))
			Expect(geoLocation.Longitude).To(Equal(-122.4194))
			Expect(geoLocation.Accuracy).To(Equal(10.0))
		})

		It("should validate latitude range", func() {
			// Create a GeoLocation builder with invalid latitude
			geoLocationBuilder := builders.NewGeoLocationBuilder().
				WithLatitude(100.0). // Greater than 90
				WithLongitude(-122.4194).
				WithAccuracy(10.0)

			// Build and validate should fail
			_, err := geoLocationBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Latitude must be between -90 and 90"))
		})

		It("should validate longitude range", func() {
			// Create a GeoLocation builder with invalid longitude
			geoLocationBuilder := builders.NewGeoLocationBuilder().
				WithLatitude(37.7749).
				WithLongitude(200.0). // Greater than 180
				WithAccuracy(10.0)

			// Build and validate should fail
			_, err := geoLocationBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Longitude must be between -180 and 180"))
		})

		It("should validate accuracy is positive", func() {
			// Create a GeoLocation builder with negative accuracy
			geoLocationBuilder := builders.NewGeoLocationBuilder().
				WithLatitude(37.7749).
				WithLongitude(-122.4194).
				WithAccuracy(-1.0) // Negative accuracy

			// Build and validate should fail
			_, err := geoLocationBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Accuracy cannot be negative"))
		})

		It("should support custom validation", func() {
			// Create a GeoLocation builder with custom validation
			geoLocationBuilder := builders.NewGeoLocationBuilder().
				WithLatitude(37.7749).
				WithLongitude(-122.4194).
				WithAccuracy(10.0).
				WithValidation(func(g *models.GeoLocation) error {
					if g.Latitude > 0 && g.Longitude < 0 {
						return fmt.Errorf("location must be in the eastern hemisphere")
					}
					return nil
				})

			// Build and validate should fail
			_, err := geoLocationBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("location must be in the eastern hemisphere"))
		})

		It("should clone the builder correctly", func() {
			// Create a base geolocation builder
			baseBuilder := builders.NewGeoLocationBuilder().
				WithLatitude(37.7749).
				WithLongitude(-122.4194).
				WithAccuracy(10.0)

			// Clone the builder and modify it
			clonedBuilder := baseBuilder.Clone().
				WithLatitude(40.7128).
				WithLongitude(-74.0060)

			// Build both geolocations
			baseGeoLocation := baseBuilder.BuildPtr()
			clonedGeoLocation := clonedBuilder.BuildPtr()

			// Verify the base geolocation
			Expect(baseGeoLocation.Latitude).To(Equal(37.7749))
			Expect(baseGeoLocation.Longitude).To(Equal(-122.4194))

			// Verify the cloned geolocation
			Expect(clonedGeoLocation.Latitude).To(Equal(40.7128))
			Expect(clonedGeoLocation.Longitude).To(Equal(-74.0060))

			// Verify that the accuracy is the same
			Expect(clonedGeoLocation.Accuracy).To(Equal(baseGeoLocation.Accuracy))
		})
	})
})
