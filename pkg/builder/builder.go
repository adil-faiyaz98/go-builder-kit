package builder

// Builder is the interface that all builders should implement
type Builder[T any] interface {
	// Build builds the object and returns it
	Build() T
	
	// BuildPtr builds the object and returns a pointer to it
	BuildPtr() *T
	
	// BuildAndValidate builds the object, validates it, and returns it along with any validation errors
	BuildAndValidate() (*T, error)
	
	// MustBuild builds the object and panics if validation fails
	MustBuild() *T
	
	// Clone creates a deep copy of the builder
	Clone() Builder[T]
}

// Validator is the interface that objects with validation should implement
type Validator interface {
	// Validate validates the object and returns an error if validation fails
	Validate() error
}
