package builders

import (
	"fmt"

	"github.com/adil-faiyaz98/go-builder-kit/models"
)

// PersonalPreferencesBuilder builds a PersonalPreferences model
type PersonalPreferencesBuilder struct {
	preferences     *models.PersonalPreferences
	validationFuncs []func(*models.PersonalPreferences) error
}

// NewPersonalPreferencesBuilder creates a new PersonalPreferencesBuilder
func NewPersonalPreferencesBuilder() *PersonalPreferencesBuilder {
	return &PersonalPreferencesBuilder{
		preferences: &models.PersonalPreferences{
			FavoriteColor:       "",
			FavoriteColors:      []string{},
			FavoriteFood:        "",
			FavoriteFoods:       []string{},
			FavoriteMusic:       "",
			MusicTastes:         []string{},
			FavoriteMovie:       "",
			MovieGenres:         []string{},
			FavoriteBook:        "",
			BookGenres:          []string{},
			FavoriteSport:       "",
			FavoriteAnimal:      "",
			Hobbies:             []string{},
			Interests:           []string{},
			Languages:           []string{},
			TravelPreferences:   map[string]string{},
			ShoppingPreferences: map[string]bool{},
		},
		validationFuncs: []func(*models.PersonalPreferences) error{},
	}
}

// WithFavoriteColor sets the FavoriteColor
func (b *PersonalPreferencesBuilder) WithFavoriteColor(favoriteColor string) *PersonalPreferencesBuilder {
	b.preferences.FavoriteColor = favoriteColor
	return b
}

// WithFavoriteColors sets the FavoriteColors
func (b *PersonalPreferencesBuilder) WithFavoriteColors(favoriteColors []string) *PersonalPreferencesBuilder {
	b.preferences.FavoriteColors = favoriteColors
	return b
}

// WithFavoriteFood sets the FavoriteFood
func (b *PersonalPreferencesBuilder) WithFavoriteFood(favoriteFood string) *PersonalPreferencesBuilder {
	b.preferences.FavoriteFood = favoriteFood
	return b
}

// WithFavoriteFoods sets the FavoriteFoods
func (b *PersonalPreferencesBuilder) WithFavoriteFoods(favoriteFoods []string) *PersonalPreferencesBuilder {
	b.preferences.FavoriteFoods = favoriteFoods
	return b
}

// WithFavoriteMusic sets the FavoriteMusic
func (b *PersonalPreferencesBuilder) WithFavoriteMusic(favoriteMusic string) *PersonalPreferencesBuilder {
	b.preferences.FavoriteMusic = favoriteMusic
	return b
}

// WithMusicTastes sets the MusicTastes
func (b *PersonalPreferencesBuilder) WithMusicTastes(musicTastes []string) *PersonalPreferencesBuilder {
	b.preferences.MusicTastes = musicTastes
	return b
}

// WithFavoriteMovie sets the FavoriteMovie
func (b *PersonalPreferencesBuilder) WithFavoriteMovie(favoriteMovie string) *PersonalPreferencesBuilder {
	b.preferences.FavoriteMovie = favoriteMovie
	return b
}

// WithMovieGenres sets the MovieGenres
func (b *PersonalPreferencesBuilder) WithMovieGenres(movieGenres []string) *PersonalPreferencesBuilder {
	b.preferences.MovieGenres = movieGenres
	return b
}

// WithFavoriteBook sets the FavoriteBook
func (b *PersonalPreferencesBuilder) WithFavoriteBook(favoriteBook string) *PersonalPreferencesBuilder {
	b.preferences.FavoriteBook = favoriteBook
	return b
}

// WithBookGenres sets the BookGenres
func (b *PersonalPreferencesBuilder) WithBookGenres(bookGenres []string) *PersonalPreferencesBuilder {
	b.preferences.BookGenres = bookGenres
	return b
}

// WithFavoriteSport sets the FavoriteSport
func (b *PersonalPreferencesBuilder) WithFavoriteSport(favoriteSport string) *PersonalPreferencesBuilder {
	b.preferences.FavoriteSport = favoriteSport
	return b
}

// WithFavoriteAnimal sets the FavoriteAnimal
func (b *PersonalPreferencesBuilder) WithFavoriteAnimal(favoriteAnimal string) *PersonalPreferencesBuilder {
	b.preferences.FavoriteAnimal = favoriteAnimal
	return b
}

// WithHobbies sets the Hobbies
func (b *PersonalPreferencesBuilder) WithHobbies(hobbies []string) *PersonalPreferencesBuilder {
	b.preferences.Hobbies = hobbies
	return b
}

// WithInterests sets the Interests
func (b *PersonalPreferencesBuilder) WithInterests(interests []string) *PersonalPreferencesBuilder {
	b.preferences.Interests = interests
	return b
}

// WithLanguages sets the Languages
func (b *PersonalPreferencesBuilder) WithLanguages(languages []string) *PersonalPreferencesBuilder {
	b.preferences.Languages = languages
	return b
}

// WithTravelPreference adds a travel preference
func (b *PersonalPreferencesBuilder) WithTravelPreference(key, value string) *PersonalPreferencesBuilder {
	if b.preferences.TravelPreferences == nil {
		b.preferences.TravelPreferences = make(map[string]string)
	}
	b.preferences.TravelPreferences[key] = value
	return b
}

// WithShoppingPreference adds a shopping preference
func (b *PersonalPreferencesBuilder) WithShoppingPreference(key string, value bool) *PersonalPreferencesBuilder {
	if b.preferences.ShoppingPreferences == nil {
		b.preferences.ShoppingPreferences = make(map[string]bool)
	}
	b.preferences.ShoppingPreferences[key] = value
	return b
}

// WithValidation adds a custom validation function
func (b *PersonalPreferencesBuilder) WithValidation(validationFunc func(*models.PersonalPreferences) error) *PersonalPreferencesBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the PersonalPreferences
func (b *PersonalPreferencesBuilder) Build() any {
	return b.preferences
}

// BuildPtr builds the PersonalPreferences and returns a pointer
func (b *PersonalPreferencesBuilder) BuildPtr() *models.PersonalPreferences {
	return b.preferences
}

// BuildAndValidate builds the PersonalPreferences and validates it
func (b *PersonalPreferencesBuilder) BuildAndValidate() (*models.PersonalPreferences, error) {
	preferences := b.preferences

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(preferences); err != nil {
			return preferences, err
		}
	}

	// Run model's Validate method
	if err := preferences.Validate(); err != nil {
		return preferences, err
	}

	return preferences, nil
}

// MustBuild builds the PersonalPreferences and panics if validation fails
func (b *PersonalPreferencesBuilder) MustBuild() *models.PersonalPreferences {
	preferences, err := b.BuildAndValidate()
	if err != nil {
		panic(fmt.Sprintf("PersonalPreferences validation failed: %s", err.Error()))
	}
	return preferences
}

// Clone creates a deep copy of the PersonalPreferencesBuilder
func (b *PersonalPreferencesBuilder) Clone() *PersonalPreferencesBuilder {
	clonedPreferences := *b.preferences
	return &PersonalPreferencesBuilder{
		preferences:     &clonedPreferences,
		validationFuncs: append([]func(*models.PersonalPreferences) error{}, b.validationFuncs...),
	}
}
