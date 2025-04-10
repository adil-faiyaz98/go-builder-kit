package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

// PersonalPreferencesBuilder builds a PersonalPreferences model
type PersonalPreferencesBuilder struct {
	personalPreferences *models.PersonalPreferences
	// Custom validation functions
	validationFuncs []func(*models.PersonalPreferences) error
}

// NewPersonalPreferencesBuilder creates a new PersonalPreferencesBuilder
func NewPersonalPreferencesBuilder() *PersonalPreferencesBuilder {
	return &PersonalPreferencesBuilder{
		personalPreferences: &models.PersonalPreferences{
			FavoriteColor: "",
			FavoriteColors: []string{},
			FavoriteFood: "",
			FavoriteFoods: []string{},
			FavoriteMusic: "",
			MusicTastes: []string{},
			FavoriteMovie: "",
			MovieGenres: []string{},
			FavoriteBook: "",
			BookGenres: []string{},
			FavoriteSport: "",
			FavoriteAnimal: "",
			Hobbies: []string{},
			Interests: []string{},
			Languages: []string{},
			TravelPreferences: map[string]string{},
			ShoppingPreferences: map[string]bool{},
		},
		validationFuncs: []func(*models.PersonalPreferences) error{},
	}
}

// NewPersonalPreferencesBuilderWithDefaults creates a new PersonalPreferencesBuilder with sensible defaults
func NewPersonalPreferencesBuilderWithDefaults() *PersonalPreferencesBuilder {
	builder := NewPersonalPreferencesBuilder()
	// Add default values here if needed
	return builder
}
// WithFavoriteColor sets the FavoriteColor
func (b *PersonalPreferencesBuilder) WithFavoriteColor(favoriteColor string) *PersonalPreferencesBuilder {
	b.personalPreferences.FavoriteColor = favoriteColor
	return b
}

// WithFavoriteColors sets the FavoriteColors
func (b *PersonalPreferencesBuilder) WithFavoriteColors(favoriteColors []string) *PersonalPreferencesBuilder {
	b.personalPreferences.FavoriteColors = append(b.personalPreferences.FavoriteColors, favoriteColors...)
	return b
}

// WithFavoriteFood sets the FavoriteFood
func (b *PersonalPreferencesBuilder) WithFavoriteFood(favoriteFood string) *PersonalPreferencesBuilder {
	b.personalPreferences.FavoriteFood = favoriteFood
	return b
}

// WithFavoriteFoods sets the FavoriteFoods
func (b *PersonalPreferencesBuilder) WithFavoriteFoods(favoriteFoods []string) *PersonalPreferencesBuilder {
	b.personalPreferences.FavoriteFoods = append(b.personalPreferences.FavoriteFoods, favoriteFoods...)
	return b
}

// WithFavoriteMusic sets the FavoriteMusic
func (b *PersonalPreferencesBuilder) WithFavoriteMusic(favoriteMusic string) *PersonalPreferencesBuilder {
	b.personalPreferences.FavoriteMusic = favoriteMusic
	return b
}

// WithMusicTastes sets the MusicTastes
func (b *PersonalPreferencesBuilder) WithMusicTastes(musicTastes []string) *PersonalPreferencesBuilder {
	b.personalPreferences.MusicTastes = append(b.personalPreferences.MusicTastes, musicTastes...)
	return b
}

// WithFavoriteMovie sets the FavoriteMovie
func (b *PersonalPreferencesBuilder) WithFavoriteMovie(favoriteMovie string) *PersonalPreferencesBuilder {
	b.personalPreferences.FavoriteMovie = favoriteMovie
	return b
}

// WithMovieGenres sets the MovieGenres
func (b *PersonalPreferencesBuilder) WithMovieGenres(movieGenres []string) *PersonalPreferencesBuilder {
	b.personalPreferences.MovieGenres = append(b.personalPreferences.MovieGenres, movieGenres...)
	return b
}

// WithFavoriteBook sets the FavoriteBook
func (b *PersonalPreferencesBuilder) WithFavoriteBook(favoriteBook string) *PersonalPreferencesBuilder {
	b.personalPreferences.FavoriteBook = favoriteBook
	return b
}

// WithBookGenres sets the BookGenres
func (b *PersonalPreferencesBuilder) WithBookGenres(bookGenres []string) *PersonalPreferencesBuilder {
	b.personalPreferences.BookGenres = append(b.personalPreferences.BookGenres, bookGenres...)
	return b
}

// WithFavoriteSport sets the FavoriteSport
func (b *PersonalPreferencesBuilder) WithFavoriteSport(favoriteSport string) *PersonalPreferencesBuilder {
	b.personalPreferences.FavoriteSport = favoriteSport
	return b
}

// WithFavoriteAnimal sets the FavoriteAnimal
func (b *PersonalPreferencesBuilder) WithFavoriteAnimal(favoriteAnimal string) *PersonalPreferencesBuilder {
	b.personalPreferences.FavoriteAnimal = favoriteAnimal
	return b
}

// WithHobbies sets the Hobbies
func (b *PersonalPreferencesBuilder) WithHobbies(hobbies []string) *PersonalPreferencesBuilder {
	b.personalPreferences.Hobbies = append(b.personalPreferences.Hobbies, hobbies...)
	return b
}

// WithInterests sets the Interests
func (b *PersonalPreferencesBuilder) WithInterests(interests []string) *PersonalPreferencesBuilder {
	b.personalPreferences.Interests = append(b.personalPreferences.Interests, interests...)
	return b
}

// WithLanguages sets the Languages
func (b *PersonalPreferencesBuilder) WithLanguages(languages []string) *PersonalPreferencesBuilder {
	b.personalPreferences.Languages = append(b.personalPreferences.Languages, languages...)
	return b
}

// WithTravelPreferences sets the TravelPreferences
func (b *PersonalPreferencesBuilder) WithTravelPreferences(key string, val string) *PersonalPreferencesBuilder {
	if b.personalPreferences.TravelPreferences == nil {
		b.personalPreferences.TravelPreferences = make(map[string]string)
	}
	b.personalPreferences.TravelPreferences[key] = val
	return b
}

// WithShoppingPreferences sets the ShoppingPreferences
func (b *PersonalPreferencesBuilder) WithShoppingPreferences(key string, val bool) *PersonalPreferencesBuilder {
	if b.personalPreferences.ShoppingPreferences == nil {
		b.personalPreferences.ShoppingPreferences = make(map[string]bool)
	}
	b.personalPreferences.ShoppingPreferences[key] = val
	return b
}


// WithValidation adds a custom validation function
func (b *PersonalPreferencesBuilder) WithValidation(validationFunc func(*models.PersonalPreferences) error) *PersonalPreferencesBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the PersonalPreferences
func (b *PersonalPreferencesBuilder) Build() interface{} {
	return b.personalPreferences
}

// BuildPtr builds the PersonalPreferences and returns a pointer
func (b *PersonalPreferencesBuilder) BuildPtr() *models.PersonalPreferences {
	return b.personalPreferences
}

// BuildAndValidate builds the PersonalPreferences and validates it
func (b *PersonalPreferencesBuilder) BuildAndValidate() (*models.PersonalPreferences, error) {
	personalPreferences := b.personalPreferences

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(personalPreferences); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(personalPreferences).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return personalPreferences, err
		}
	}

	return personalPreferences, nil
}

// MustBuild builds the PersonalPreferences and panics if validation fails
func (b *PersonalPreferencesBuilder) MustBuild() *models.PersonalPreferences {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *PersonalPreferencesBuilder) Clone() *PersonalPreferencesBuilder {
	clonedPersonalPreferences := *b.personalPreferences
	return &PersonalPreferencesBuilder{
		personalPreferences: &clonedPersonalPreferences,
		validationFuncs: append([]func(*models.PersonalPreferences) error{}, b.validationFuncs...),
	}
}
