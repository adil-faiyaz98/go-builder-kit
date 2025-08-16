package address_test

import (
	"github.com/adil-faiyaz98/go-builder-kit/builders"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Address Builder", func() {

	Context("Basic Address Creation", func() {
		It("should create an address with basic fields", func() {
			// Create an address with basic fields
			addressBuilder := builders.NewAddressBuilder().
				WithStreet("123 Main St").
				WithCity("San Francisco").
				WithState("CA").
				WithPostalCode("94105").
				WithCountry("USA").
				WithType("Home")

			// Build the address
			address := addressBuilder.BuildPtr()

			// Verify the address fields
			Expect(address.Street).To(Equal("123 Main St"))
			Expect(address.City).To(Equal("San Francisco"))
			Expect(address.State).To(Equal("CA"))
			Expect(address.PostalCode).To(Equal("94105"))
			Expect(address.Country).To(Equal("USA"))
			Expect(address.Type).To(Equal("Home"))
		})
	})

	Context("Address with GeoLocation", func() {
		It("should create an address with geo coordinates", func() {
			// Create a GeoLocation
			geoLocationBuilder := builders.NewGeoLocationBuilder().
				WithLatitude(37.7749).
				WithLongitude(-122.4194)

			// Create an address with geo coordinates
			addressBuilder := builders.NewAddressBuilder().
				WithStreet("123 Main St").
				WithCity("San Francisco").
				WithState("CA").
				WithCoordinates(geoLocationBuilder)

			// Build the address
			address := addressBuilder.BuildPtr()

			// Verify the address fields
			Expect(address.Street).To(Equal("123 Main St"))
			Expect(address.City).To(Equal("San Francisco"))
			Expect(address.State).To(Equal("CA"))

			// Verify the geo coordinates
			Expect(address.Coordinates).NotTo(BeNil())
			Expect(address.Coordinates.Latitude).To(Equal(37.7749))
			Expect(address.Coordinates.Longitude).To(Equal(-122.4194))
		})
	})

	Context("Address Types", func() {
		It("should create addresses with different types", func() {
			// Create a home address
			homeAddressBuilder := builders.NewAddressBuilder().
				WithStreet("123 Main St").
				WithCity("San Francisco").
				WithState("CA").
				WithType("Home")

			// Create a work address
			workAddressBuilder := builders.NewAddressBuilder().
				WithStreet("456 Market St").
				WithCity("San Francisco").
				WithState("CA").
				WithType("Work")

			// Build the addresses
			homeAddress := homeAddressBuilder.BuildPtr()
			workAddress := workAddressBuilder.BuildPtr()

			// Verify the home address
			Expect(homeAddress.Street).To(Equal("123 Main St"))
			Expect(homeAddress.Type).To(Equal("Home"))

			// Verify the work address
			Expect(workAddress.Street).To(Equal("456 Market St"))
			Expect(workAddress.Type).To(Equal("Work"))
		})
	})

	Context("Builder Cloning", func() {
		It("should clone an address builder correctly", func() {
			// Create a base address builder
			baseBuilder := builders.NewAddressBuilder().
				WithStreet("123 Main St").
				WithCity("San Francisco").
				WithState("CA").
				WithCountry("USA")

			// Clone the builder and modify it
			clonedBuilder := baseBuilder.Clone().
				WithStreet("456 Market St").
				WithPostalCode("94105")

			// Build both addresses
			baseAddress := baseBuilder.BuildPtr()
			clonedAddress := clonedBuilder.BuildPtr()

			// Verify the base address wasn't affected by changes to the clone
			Expect(baseAddress.Street).To(Equal("123 Main St"))
			Expect(baseAddress.PostalCode).To(Equal(""))

			// Verify the cloned address has the new values
			Expect(clonedAddress.Street).To(Equal("456 Market St"))
			Expect(clonedAddress.City).To(Equal("San Francisco")) // Unchanged
			Expect(clonedAddress.State).To(Equal("CA"))           // Unchanged
			Expect(clonedAddress.PostalCode).To(Equal("94105"))
		})
	})
})
