package generator

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"unicode"
)

// Builder is the interface that all builders must implement
type Builder interface {
	Build() interface{}
}

// BuilderFunc is a function that creates a new builder
type BuilderFunc func() Builder

// BuilderRegistry manages builder functions
type BuilderRegistry struct {
	registry map[string]BuilderFunc
	mutex    sync.RWMutex
}

// NewBuilderRegistry creates a new BuilderRegistry
func NewBuilderRegistry() *BuilderRegistry {
	return &BuilderRegistry{
		registry: make(map[string]BuilderFunc),
	}
}

// Register registers a builder function
func (r *BuilderRegistry) Register(name string, fn BuilderFunc) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.registry[name] = fn
}

// Get returns a builder function by name
func (r *BuilderRegistry) Get(name string) (BuilderFunc, bool) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	fn, ok := r.registry[name]
	return fn, ok
}

// GetAll returns all registered builder functions
func (r *BuilderRegistry) GetAll() map[string]BuilderFunc {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	// Create a copy to avoid concurrent map access
	copy := make(map[string]BuilderFunc, len(r.registry))
	for k, v := range r.registry {
		copy[k] = v
	}
	return copy
}

// CreateBuilder creates a new builder for the given type
func (r *BuilderRegistry) CreateBuilder(typeName string) (Builder, error) {
	builderFn, ok := r.Get(typeName)
	if !ok {
		return nil, fmt.Errorf("no builder registered for type %s", typeName)
	}
	return builderFn(), nil
}

// DefaultRegistry is the default builder registry
var DefaultRegistry = NewBuilderRegistry()

// Register registers a builder function with the default registry
func Register(name string, fn BuilderFunc) {
	DefaultRegistry.Register(name, fn)
}

// Get returns a builder function by name from the default registry
func Get(name string) (BuilderFunc, bool) {
	return DefaultRegistry.Get(name)
}

// GetAll returns all registered builder functions from the default registry
func GetAll() map[string]BuilderFunc {
	return DefaultRegistry.GetAll()
}

// CreateBuilder creates a new builder for the given type using the default registry
func CreateBuilder(typeName string) (Builder, error) {
	return DefaultRegistry.CreateBuilder(typeName)
}

// CamelToSnake converts a camelCase string to snake_case
func CamelToSnake(s string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// SnakeToCamel converts a snake_case string to camelCase
func SnakeToCamel(s string) string {
	var result string
	for _, word := range strings.Split(s, "_") {
		if word == "" {
			continue
		}
		if len(word) > 0 {
			first := unicode.ToUpper(rune(word[0]))
			rest := word[1:]
			result += string(first) + rest
		}
	}
	return result
}
