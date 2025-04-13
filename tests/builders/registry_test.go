package builders_test

import (
	"github.com/adil-faiyaz98/go-builder-kit/builders"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Builder Registry", func() {
	var registry *builders.BuilderRegistry

	BeforeEach(func() {
		// Create a new registry for each test
		registry = builders.NewBuilderRegistry()
	})

	Context("Registering and retrieving builders", func() {
		It("should register and retrieve a builder", func() {
			// Register a builder
			registry.Register("person", func() interface{} {
				return builders.NewPersonBuilder()
			})

			// Get the builder by name
			builderFunc, ok := registry.Get("person")
			builder := builderFunc()

			// Verify the builder was retrieved
			Expect(ok).To(BeTrue())
			Expect(builder).NotTo(BeNil())

			// Verify the builder is of the correct type
			personBuilder, ok := builder.(*builders.PersonBuilder)
			Expect(ok).To(BeTrue())
			Expect(personBuilder).NotTo(BeNil())

			// Use the builder
			person := personBuilder.
				WithName("John Doe").
				WithAge(30).
				BuildPtr()

			// Verify the person was built correctly
			Expect(person.Name).To(Equal("John Doe"))
			Expect(person.Age).To(Equal(30))
		})

		It("should return false when getting a non-existent builder", func() {
			// Get a builder that doesn't exist
			builderFunc, ok := registry.Get("non-existent")

			// Verify the builder was not found
			Expect(ok).To(BeFalse())
			Expect(builderFunc).To(BeNil())
		})

		It("should register multiple builders", func() {
			// Register multiple builders
			registry.Register("person", func() interface{} {
				return builders.NewPersonBuilder()
			})

			registry.Register("address", func() interface{} {
				return builders.NewAddressBuilder()
			})

			// Get the builders by name
			personBuilderFunc, personOk := registry.Get("person")
			addressBuilderFunc, addressOk := registry.Get("address")

			// Verify the builders were retrieved
			Expect(personOk).To(BeTrue())
			Expect(personBuilderFunc).NotTo(BeNil())

			Expect(addressOk).To(BeTrue())
			Expect(addressBuilderFunc).NotTo(BeNil())
		})

		It("should override a builder when registering with the same name", func() {
			// Register a builder
			registry.Register("person", func() interface{} {
				return builders.NewPersonBuilder().WithName("Original")
			})

			// Register another builder with the same name
			registry.Register("person", func() interface{} {
				return builders.NewPersonBuilder().WithName("Override")
			})

			// Get the builder by name
			builderFunc, ok := registry.Get("person")

			// Verify the builder was retrieved
			Expect(ok).To(BeTrue())
			Expect(builderFunc).NotTo(BeNil())

			// Verify the builder is the overridden one
			builder := builderFunc()
			personBuilder, ok := builder.(*builders.PersonBuilder)
			Expect(ok).To(BeTrue())

			person := personBuilder.BuildPtr()
			Expect(person.Name).To(Equal("Override"))
		})
	})
})
