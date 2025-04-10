package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

// PersonalPreferencesBuilder builds a PersonalPreferences model
type PersonalPreferencesBuilder struct {
	preferences    *models.PersonalPreferences
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

// AddFavoriteColor adds a favorite color to the FavoriteColors slice
func (b *PersonalPreferencesBuilder) AddFavoriteColor(favoriteColor string) *PersonalPreferencesBuilder {
	b.preferences.FavoriteColors = append(b.preferences.FavoriteColors, favoriteColor)
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

// AddFavoriteFood adds a favorite food to the FavoriteFoods slice
func (b *PersonalPreferencesBuilder) AddFavoriteFood(favoriteFood string) *PersonalPreferencesBuilder {
	b.preferences.FavoriteFoods = append(b.preferences.FavoriteFoods, favoriteFood)
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

// AddMusicTaste adds a music taste to the MusicTastes slice
func (b *PersonalPreferencesBuilder) AddMusicTaste(musicTaste string) *PersonalPreferencesBuilder {
	b.preferences.MusicTastes = append(b.preferences.MusicTastes, musicTaste)
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

// AddMovieGenre adds a movie genre to the MovieGenres slice
func (b *PersonalPreferencesBuilder) AddMovieGenre(movieGenre string) *PersonalPreferencesBuilder {
	b.preferences.MovieGenres = append(b.preferences.MovieGenres, movieGenre)
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

// AddBookGenre adds a book genre to the BookGenres slice
func (b *PersonalPreferencesBuilder) AddBookGenre(bookGenre string) *PersonalPreferencesBuilder {
	b.preferences.BookGenres = append(b.preferences.BookGenres, bookGenre)
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

// AddHobby adds a hobby to the Hobbies slice
func (b *PersonalPreferencesBuilder) AddHobby(hobby string) *PersonalPreferencesBuilder {
	b.preferences.Hobbies = append(b.preferences.Hobbies, hobby)
	return b
}

// WithInterests sets the Interests
func (b *PersonalPreferencesBuilder) WithInterests(interests []string) *PersonalPreferencesBuilder {
	b.preferences.Interests = interests
	return b
}

// AddInterest adds an interest to the Interests slice
func (b *PersonalPreferencesBuilder) AddInterest(interest string) *PersonalPreferencesBuilder {
	b.preferences.Interests = append(b.preferences.Interests, interest)
	return b
}

// WithLanguages sets the Languages
func (b *PersonalPreferencesBuilder) WithLanguages(languages []string) *PersonalPreferencesBuilder {
	b.preferences.Languages = languages
	return b
}

// AddLanguage adds a language to the Languages slice
func (b *PersonalPreferencesBuilder) AddLanguage(language string) *PersonalPreferencesBuilder {
	b.preferences.Languages = append(b.preferences.Languages, language)
	return b
}

// WithTravelPreferences sets the TravelPreferences
func (b *PersonalPreferencesBuilder) WithTravelPreferences(travelPreferences map[string]string) *PersonalPreferencesBuilder {
	b.preferences.TravelPreferences = travelPreferences
	return b
}

// AddTravelPreference adds a travel preference to the TravelPreferences map
func (b *PersonalPreferencesBuilder) AddTravelPreference(key, value string) *PersonalPreferencesBuilder {
	if b.preferences.TravelPreferences == nil {
		b.preferences.TravelPreferences = make(map[string]string)
	}
	b.preferences.TravelPreferences[key] = value
	return b
}

// WithShoppingPreferences sets the ShoppingPreferences
func (b *PersonalPreferencesBuilder) WithShoppingPreferences(shoppingPreferences map[string]bool) *PersonalPreferencesBuilder {
	b.preferences.ShoppingPreferences = shoppingPreferences
	return b
}

// AddShoppingPreference adds a shopping preference to the ShoppingPreferences map
func (b *PersonalPreferencesBuilder) AddShoppingPreference(key string, value bool) *PersonalPreferencesBuilder {
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
func (b *PersonalPreferencesBuilder) Build() interface{} {
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
			return nil, fmt.Errorf("custom validation failed: %w", err)
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
		panic(err)
	}
	return preferences
}

// Clone creates a deep copy of the builder
func (b *PersonalPreferencesBuilder) Clone() *PersonalPreferencesBuilder {
	clonedPreferences := *b.preferences
	return &PersonalPreferencesBuilder{
		preferences:    &clonedPreferences,
		validationFuncs: append([]func(*models.PersonalPreferences) error{}, b.validationFuncs...),
	}
}
