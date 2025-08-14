package builders

import (
	"fmt"
	"github.com/adil-faiyaz98/go-builder-kit/v2/models"
	"github.com/adil-faiyaz98/go-builder-kit/v2/pkg/builder"
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
	if b == nil {
		return b
	}
	b.personalPreferences.FavoriteColor = builder.SanitizeString(favoriteColor)
	return b
}

// WithFavoriteColors sets the FavoriteColors
func (b *PersonalPreferencesBuilder) WithFavoriteColors(favoriteColors []string) *PersonalPreferencesBuilder {
	if b == nil {
		return b
	}
	b.personalPreferences.FavoriteColors = append(b.personalPreferences.FavoriteColors, favoriteColors...)
	return b
}

// WithFavoriteFood sets the FavoriteFood
func (b *PersonalPreferencesBuilder) WithFavoriteFood(favoriteFood string) *PersonalPreferencesBuilder {
	if b == nil {
		return b
	}
	b.personalPreferences.FavoriteFood = builder.SanitizeString(favoriteFood)
	return b
}

// WithFavoriteFoods sets the FavoriteFoods
func (b *PersonalPreferencesBuilder) WithFavoriteFoods(favoriteFoods []string) *PersonalPreferencesBuilder {
	if b == nil {
		return b
	}
	b.personalPreferences.FavoriteFoods = append(b.personalPreferences.FavoriteFoods, favoriteFoods...)
	return b
}

// WithFavoriteMusic sets the FavoriteMusic
func (b *PersonalPreferencesBuilder) WithFavoriteMusic(favoriteMusic string) *PersonalPreferencesBuilder {
	if b == nil {
		return b
	}
	b.personalPreferences.FavoriteMusic = builder.SanitizeString(favoriteMusic)
	return b
}

// WithMusicTastes sets the MusicTastes
func (b *PersonalPreferencesBuilder) WithMusicTastes(musicTastes []string) *PersonalPreferencesBuilder {
	if b == nil {
		return b
	}
	b.personalPreferences.MusicTastes = append(b.personalPreferences.MusicTastes, musicTastes...)
	return b
}

// WithFavoriteMovie sets the FavoriteMovie
func (b *PersonalPreferencesBuilder) WithFavoriteMovie(favoriteMovie string) *PersonalPreferencesBuilder {
	if b == nil {
		return b
	}
	b.personalPreferences.FavoriteMovie = builder.SanitizeString(favoriteMovie)
	return b
}

// WithMovieGenres sets the MovieGenres
func (b *PersonalPreferencesBuilder) WithMovieGenres(movieGenres []string) *PersonalPreferencesBuilder {
	if b == nil {
		return b
	}
	b.personalPreferences.MovieGenres = append(b.personalPreferences.MovieGenres, movieGenres...)
	return b
}

// WithFavoriteBook sets the FavoriteBook
func (b *PersonalPreferencesBuilder) WithFavoriteBook(favoriteBook string) *PersonalPreferencesBuilder {
	if b == nil {
		return b
	}
	b.personalPreferences.FavoriteBook = builder.SanitizeString(favoriteBook)
	return b
}

// WithBookGenres sets the BookGenres
func (b *PersonalPreferencesBuilder) WithBookGenres(bookGenres []string) *PersonalPreferencesBuilder {
	if b == nil {
		return b
	}
	b.personalPreferences.BookGenres = append(b.personalPreferences.BookGenres, bookGenres...)
	return b
}

// WithFavoriteSport sets the FavoriteSport
func (b *PersonalPreferencesBuilder) WithFavoriteSport(favoriteSport string) *PersonalPreferencesBuilder {
	if b == nil {
		return b
	}
	b.personalPreferences.FavoriteSport = builder.SanitizeString(favoriteSport)
	return b
}

// WithFavoriteAnimal sets the FavoriteAnimal
func (b *PersonalPreferencesBuilder) WithFavoriteAnimal(favoriteAnimal string) *PersonalPreferencesBuilder {
	if b == nil {
		return b
	}
	b.personalPreferences.FavoriteAnimal = builder.SanitizeString(favoriteAnimal)
	return b
}

// WithHobbies sets the Hobbies
func (b *PersonalPreferencesBuilder) WithHobbies(hobbies []string) *PersonalPreferencesBuilder {
	if b == nil {
		return b
	}
	b.personalPreferences.Hobbies = append(b.personalPreferences.Hobbies, hobbies...)
	return b
}

// WithInterests sets the Interests
func (b *PersonalPreferencesBuilder) WithInterests(interests []string) *PersonalPreferencesBuilder {
	if b == nil {
		return b
	}
	b.personalPreferences.Interests = append(b.personalPreferences.Interests, interests...)
	return b
}

// WithLanguages sets the Languages
func (b *PersonalPreferencesBuilder) WithLanguages(languages []string) *PersonalPreferencesBuilder {
	if b == nil {
		return b
	}
	b.personalPreferences.Languages = append(b.personalPreferences.Languages, languages...)
	return b
}

// WithTravelPreferences sets the TravelPreferences
func (b *PersonalPreferencesBuilder) WithTravelPreferences(key string, val string) *PersonalPreferencesBuilder {
	if b == nil {
		return b
	}
	if b.personalPreferences.TravelPreferences == nil {
		b.personalPreferences.TravelPreferences = make(map[string]string)
	}
	b.personalPreferences.TravelPreferences[key] = val
	return b
}

// WithShoppingPreferences sets the ShoppingPreferences
func (b *PersonalPreferencesBuilder) WithShoppingPreferences(key string, val bool) *PersonalPreferencesBuilder {
	if b == nil {
		return b
	}
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
	if b == nil || b.personalPreferences == nil {
		return nil, fmt.Errorf("builder is not properly initialized")
	}

	personalPreferences := b.personalPreferences

	// Run custom validation functions
	for i, validationFunc := range b.validationFuncs {
		if validationFunc == nil {
			continue // Skip nil validators
		}
		if err := validationFunc(personalPreferences); err != nil {
			return nil, fmt.Errorf("custom validation failed at index %d: %w", i, err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := any(personalPreferences).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return personalPreferences, fmt.Errorf("model validation failed: %w", err)
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
	if b == nil || b.personalPreferences == nil {
		return NewPersonalPreferencesBuilder()
	}

	// Deep copy the struct
	clonedPersonalPreferences := *b.personalPreferences

	// Create new builder with cloned data
	clonedBuilder := &PersonalPreferencesBuilder{
		personalPreferences: &clonedPersonalPreferences,
		validationFuncs: make([]func(*models.PersonalPreferences) error, 0, len(b.validationFuncs)),
	}

	// Copy validation functions safely
	if len(b.validationFuncs) > 0 {
		clonedBuilder.validationFuncs = append(clonedBuilder.validationFuncs, b.validationFuncs...)
	}

	return clonedBuilder
}
