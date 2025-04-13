package builder

// GenericBuilder is a generic builder interface for creating objects of type T
type GenericBuilder[T any] interface {
	// Build builds the object and returns it
	Build() T
	
	// BuildPtr builds the object and returns a pointer to it
	BuildPtr() *T
	
	// BuildAndValidate builds the object, validates it, and returns it along with any validation errors
	BuildAndValidate() (*T, error)
	
	// MustBuild builds the object and panics if validation fails
	MustBuild() *T
	
	// Clone creates a deep copy of the builder
	Clone() GenericBuilder[T]
}

// GenericValidator is a generic validator interface for validating objects of type T
type GenericValidator[T any] interface {
	// Validate validates the object and returns an error if validation fails
	Validate(obj T) error
}

// GenericBuilderWithValidation is a generic builder interface with validation support
type GenericBuilderWithValidation[T any] interface {
	GenericBuilder[T]
	
	// WithValidation adds a validation function to the builder
	WithValidation(fn func(T) error) GenericBuilderWithValidation[T]
}

// BaseGenericBuilder is a base implementation of GenericBuilderWithValidation
type BaseGenericBuilder[T any] struct {
	// build is a function that builds the object
	build func() T
	
	// validators is a slice of validation functions
	validators []func(T) error
}

// NewBaseGenericBuilder creates a new BaseGenericBuilder
func NewBaseGenericBuilder[T any](buildFn func() T) *BaseGenericBuilder[T] {
	return &BaseGenericBuilder[T]{
		build:      buildFn,
		validators: make([]func(T) error, 0),
	}
}

// Build builds the object and returns it
func (b *BaseGenericBuilder[T]) Build() T {
	return b.build()
}

// BuildPtr builds the object and returns a pointer to it
func (b *BaseGenericBuilder[T]) BuildPtr() *T {
	obj := b.build()
	return &obj
}

// WithValidation adds a validation function to the builder
func (b *BaseGenericBuilder[T]) WithValidation(fn func(T) error) GenericBuilderWithValidation[T] {
	b.validators = append(b.validators, fn)
	return b
}

// BuildAndValidate builds the object, validates it, and returns it along with any validation errors
func (b *BaseGenericBuilder[T]) BuildAndValidate() (*T, error) {
	obj := b.build()
	
	for _, validator := range b.validators {
		if err := validator(obj); err != nil {
			return nil, err
		}
	}
	
	return &obj, nil
}

// MustBuild builds the object and panics if validation fails
func (b *BaseGenericBuilder[T]) MustBuild() *T {
	obj, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return obj
}

// Clone creates a deep copy of the builder
func (b *BaseGenericBuilder[T]) Clone() GenericBuilder[T] {
	clone := &BaseGenericBuilder[T]{
		build:      b.build,
		validators: make([]func(T) error, len(b.validators)),
	}
	
	copy(clone.validators, b.validators)
	return clone
}
