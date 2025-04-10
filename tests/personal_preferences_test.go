// +build ignore

// +build ignore

import (
	"fmt"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
	"github.com/adil-faiyaz98/go-builder-kit/models"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("PersonalPreferences", func() {
	Context("PersonalPreferences Builder", func() {
		It("should build a valid PersonalPreferences with all fields", func() {
			// Create a PersonalPreferences builder
			preferencesBuilder := builders.NewPersonalPreferencesBuilder().
				WithFavoriteColor("Blue").
				WithFavoriteColors([]string{"Blue", "Green", "Purple"}).
				WithFavoriteFood("Pizza").
				WithFavoriteFoods([]string{"Pizza", "Sushi", "Tacos"}).
				WithFavoriteMusic("Rock").
				WithMusicTastes([]string{"Rock", "Jazz", "Classical"}).
				WithFavoriteMovie("The Matrix").
				WithMovieGenres([]string{"Sci-Fi", "Action", "Drama"}).
				WithFavoriteBook("1984").
				WithBookGenres([]string{"Sci-Fi", "Dystopian", "Classic"}).
				WithFavoriteSport("Basketball").
				WithFavoriteAnimal("Dog").
				WithHobbies([]string{"Reading", "Hiking", "Coding"}).
				WithInterests([]string{"Technology", "Science", "Art"}).
				WithLanguages([]string{"English", "Spanish", "French"}).
				WithTravelPreference("Accommodation", "Hotel").
				WithTravelPreference("Transportation", "Plane").
				WithShoppingPreference("Online", true).
				WithShoppingPreference("Eco-friendly", true)

			// Build the PersonalPreferences
			preferences, err := preferencesBuilder.BuildAndValidate()
			Expect(err).To(BeNil())
			Expect(preferences).NotTo(BeNil())
			Expect(preferences.FavoriteColor).To(Equal("Blue"))
			Expect(preferences.FavoriteColors).To(HaveLen(3))
			Expect(preferences.FavoriteColors).To(ContainElement("Green"))
			Expect(preferences.FavoriteFood).To(Equal("Pizza"))
			Expect(preferences.FavoriteFoods).To(HaveLen(3))
			Expect(preferences.FavoriteFoods).To(ContainElement("Sushi"))
			Expect(preferences.FavoriteMusic).To(Equal("Rock"))
			Expect(preferences.MusicTastes).To(HaveLen(3))
			Expect(preferences.MusicTastes).To(ContainElement("Jazz"))
			Expect(preferences.FavoriteMovie).To(Equal("The Matrix"))
			Expect(preferences.MovieGenres).To(HaveLen(3))
			Expect(preferences.MovieGenres).To(ContainElement("Action"))
			Expect(preferences.FavoriteBook).To(Equal("1984"))
			Expect(preferences.BookGenres).To(HaveLen(3))
			Expect(preferences.BookGenres).To(ContainElement("Dystopian"))
			Expect(preferences.FavoriteSport).To(Equal("Basketball"))
			Expect(preferences.FavoriteAnimal).To(Equal("Dog"))
			Expect(preferences.Hobbies).To(HaveLen(3))
			Expect(preferences.Hobbies).To(ContainElement("Hiking"))
			Expect(preferences.Interests).To(HaveLen(3))
			Expect(preferences.Interests).To(ContainElement("Science"))
			Expect(preferences.Languages).To(HaveLen(3))
			Expect(preferences.Languages).To(ContainElement("Spanish"))
			Expect(preferences.TravelPreferences).To(HaveLen(2))
			Expect(preferences.TravelPreferences["Accommodation"]).To(Equal("Hotel"))
			Expect(preferences.TravelPreferences["Transportation"]).To(Equal("Plane"))
			Expect(preferences.ShoppingPreferences).To(HaveLen(2))
			Expect(preferences.ShoppingPreferences["Online"]).To(BeTrue())
			Expect(preferences.ShoppingPreferences["Eco-friendly"]).To(BeTrue())
		})

		It("should support custom validation", func() {
			// Create a PersonalPreferences builder with custom validation
			preferencesBuilder := builders.NewPersonalPreferencesBuilder().
				WithFavoriteColor("Ultraviolet"). // Not a common color
				WithValidation(func(p *models.PersonalPreferences) error {
					if p.FavoriteColor == "Ultraviolet" {
						return fmt.Errorf("Ultraviolet is not a visible color")
					}
					return nil
				})

			// Build and validate should fail
			_, err := preferencesBuilder.BuildAndValidate()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("Ultraviolet is not a visible color"))
		})

		It("should clone the builder correctly", func() {
			// Create a base preferences builder
			baseBuilder := builders.NewPersonalPreferencesBuilder().
				WithFavoriteColor("Blue").
				WithFavoriteFood("Pizza").
				WithFavoriteMusic("Rock").
				WithHobbies([]string{"Reading", "Hiking", "Coding"})

			// Clone the builder and modify it
			clonedBuilder := baseBuilder.Clone().
				WithFavoriteColor("Green").
				WithFavoriteFood("Sushi")

			// Build both preferences
			basePreferences := baseBuilder.BuildPtr()
			clonedPreferences := clonedBuilder.BuildPtr()

			// Verify the base preferences
			Expect(basePreferences.FavoriteColor).To(Equal("Blue"))
			Expect(basePreferences.FavoriteFood).To(Equal("Pizza"))

			// Verify the cloned preferences
			Expect(clonedPreferences.FavoriteColor).To(Equal("Green"))
			Expect(clonedPreferences.FavoriteFood).To(Equal("Sushi"))

			// Verify that the music and hobbies are the same
			Expect(clonedPreferences.FavoriteMusic).To(Equal(basePreferences.FavoriteMusic))
			Expect(clonedPreferences.Hobbies).To(Equal(basePreferences.Hobbies))
		})
	})
})
