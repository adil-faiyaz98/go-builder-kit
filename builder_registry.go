package builders

import (
	"fmt"
)

// BuilderFunc is a function that creates a new builder
type BuilderFunc func() Builder

// builderRegistry is a registry of builder functions
var builderRegistry = make(map[string]BuilderFunc)

// RegisterBuilder registers a builder function
func RegisterBuilder(name string, fn BuilderFunc) {
	builderRegistry[name] = fn
}

// GetBuilder returns a builder function by name
func GetBuilder(name string) (BuilderFunc, bool) {
	fn, ok := builderRegistry[name]
	return fn, ok
}

// GetAllBuilders returns all registered builder functions
func GetAllBuilders() map[string]BuilderFunc {
	return builderRegistry
}

// CreateBuilder creates a new builder by name
func CreateBuilder(name string) (Builder, error) {
	fn, ok := GetBuilder(name)
	if !ok {
		return nil, fmt.Errorf("builder not found: %s", name)
	}
	return fn(), nil
}
