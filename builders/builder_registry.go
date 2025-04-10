package builders

import (
	"fmt"
)

// BuilderRegistry is a registry of builder functions
type BuilderRegistry struct {
	builders map[string]func() interface{}
}

// NewBuilderRegistry creates a new BuilderRegistry
func NewBuilderRegistry() *BuilderRegistry {
	return &BuilderRegistry{
		builders: make(map[string]func() interface{}),
	}
}

// Register registers a builder function
func (r *BuilderRegistry) Register(typeName string, builderFunc func() interface{}) {
	r.builders[typeName] = builderFunc
}

// Get gets a builder function by type name
func (r *BuilderRegistry) Get(typeName string) (func() interface{}, bool) {
	builderFunc, ok := r.builders[typeName]
	return builderFunc, ok
}

// Create creates a new builder for the given type
func (r *BuilderRegistry) Create(typeName string) (interface{}, error) {
	builderFunc, ok := r.Get(typeName)
	if !ok {
		return nil, fmt.Errorf("no builder registered for type %s", typeName)
	}
	return builderFunc(), nil
}

// DefaultRegistry is the default builder registry
var DefaultRegistry = NewBuilderRegistry()

// Register registers a builder function with the default registry
func Register(typeName string, builderFunc func() interface{}) {
	DefaultRegistry.Register(typeName, builderFunc)
}

// Get gets a builder function by type name from the default registry
func Get(typeName string) (func() interface{}, bool) {
	return DefaultRegistry.Get(typeName)
}

// Create creates a new builder for the given type using the default registry
func Create(typeName string) (interface{}, error) {
	return DefaultRegistry.Create(typeName)
}
