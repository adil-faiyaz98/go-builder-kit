// +build ignore

// +build ignore

import (
	"fmt"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
	"github.com/adil-faiyaz98/go-builder-kit/models"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Address", func() {
	Context("Address Builder", func() {
		It("should build a valid Address with all fields", func() {
			// Create a GeoLocation builder
			geoLocationBuilder := builders.NewGeoLocationBuilder().
				WithLatitude(37.7749).
				WithLongitude(-122.4194).
				WithAccuracy(10.0)

			// Create an Address builder
			addressBuilder := builders.NewAddressBuilder().
				WithStreet("123 Main St").
				WithCity("San Francisco").
				WithState("CA").
				WithPostalCode("94105").
				WithCountry("USA").
				WithCoordinates(geoLocationBuilder).
				WithType("Home").
				WithIsPrimary(true)

			// Build the Address
			address, err := addressBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(address).NotTo(BeNil())
			Expect(address.Street).To(Equal("123 Main St"))
			Expect(address.City).To(Equal("San Francisco"))
			Expect(address.State).To(Equal("CA"))
			Expect(address.PostalCode).To(Equal("94105"))
			Expect(address.Country).To(Equal("USA"))
			Expect(address.Type).To(Equal("Home"))
			Expect(address.IsPrimary).To(BeTrue())

			// Verify nested GeoLocation
			Expect(address.Coordinates).NotTo(BeNil())
			Expect(address.Coordinates.Latitude).To(Equal(37.7749))
			Expect(address.Coordinates.Longitude).To(Equal(-122.4194))
			Expect(address.Coordinates.Accuracy).To(Equal(10.0))
		})

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
			addressBuilder := builders.NewAddressBuilder().
				WithStreet("123 Main St").
				WithCity("San Francisco").
				WithState("CA").
				WithPostalCode("94"). // Too short
				WithCountry("USA").
				WithType("Home")

			// Build and validate should fail
			_, err := addressBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("PostalCode format is invalid"))
		})

		It("should fail validation with invalid address type", func() {
			// Create an Address builder with invalid type
			addressBuilder := builders.NewAddressBuilder().
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

		It("should fail validation with invalid coordinates", func() {
			// Create a GeoLocation builder with invalid latitude
			geoLocationBuilder := builders.NewGeoLocationBuilder().
				WithLatitude(100.0). // Greater than 90
				WithLongitude(-122.4194).
				WithAccuracy(10.0)

			// Create an Address builder with invalid coordinates
			addressBuilder := builders.NewAddressBuilder().
				WithStreet("123 Main St").
				WithCity("San Francisco").
				WithState("CA").
				WithPostalCode("94105").
				WithCountry("USA").
				WithCoordinates(geoLocationBuilder)

			// Build and validate should fail
			_, err := addressBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Latitude must be between -90 and 90"))
		})

		It("should support custom validation", func() {
			// Create an Address builder with custom validation
			addressBuilder := builders.NewAddressBuilder().
				WithStreet("123 Main St").
				WithCity("San Francisco").
				WithState("CA").
				WithPostalCode("94105").
				WithCountry("USA").
				WithValidation(func(a *models.Address) error {
					if a.State != "NY" {
						return fmt.Errorf("State must be NY")
					}
					return nil
				})

			// Build and validate should fail
			_, err := addressBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("State must be NY"))

			// Fix the state and try again
			addressBuilder.WithState("NY")
			address, err := addressBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(address).NotTo(BeNil())
			Expect(address.State).To(Equal("NY"))
		})

		It("should clone the builder correctly", func() {
			// Create a base address builder
			baseBuilder := builders.NewAddressBuilder().
				WithStreet("123 Main St").
				WithCity("San Francisco").
				WithState("CA").
				WithPostalCode("94105").
				WithCountry("USA")

			// Clone the builder and modify it
			clonedBuilder := baseBuilder.Clone().
				WithStreet("456 Oak St").
				WithCity("New York")

			// Build both addresses
			baseAddress := baseBuilder.Build()
			clonedAddress := clonedBuilder.Build()

			// Verify the base address
			Expect(baseAddress.(*models.Address).Street).To(Equal("123 Main St"))
			Expect(baseAddress.(*models.Address).City).To(Equal("San Francisco"))

			// Verify the cloned address
			Expect(clonedAddress.(*models.Address).Street).To(Equal("456 Oak St"))
			Expect(clonedAddress.(*models.Address).City).To(Equal("New York"))

			// Verify that the country is the same
			Expect(clonedAddress.(*models.Address).Country).To(Equal(baseAddress.(*models.Address).Country))
		})
	})
})
