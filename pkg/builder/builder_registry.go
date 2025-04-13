package builder

import (
	"fmt"
	"sync"
)

// BuilderFunc is a function that creates a new builder
type BuilderFunc func() interface{}

// BuilderRegistry is a registry of builder functions
type BuilderRegistry struct {
	builders map[string]BuilderFunc
	mu       sync.RWMutex
}

// NewBuilderRegistry creates a new BuilderRegistry
func NewBuilderRegistry() *BuilderRegistry {
	return &BuilderRegistry{
		builders: make(map[string]BuilderFunc),
	}
}

// Register registers a builder function
func (r *BuilderRegistry) Register(name string, fn BuilderFunc) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.builders[name] = fn
}

// Get returns a builder function by name
func (r *BuilderRegistry) Get(name string) (BuilderFunc, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	fn, ok := r.builders[name]
	return fn, ok
}

// GetAll returns all registered builder functions
func (r *BuilderRegistry) GetAll() map[string]BuilderFunc {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	// Create a copy to avoid race conditions
	result := make(map[string]BuilderFunc, len(r.builders))
	for k, v := range r.builders {
		result[k] = v
	}
	
	return result
}

// Create creates a new builder by name
func (r *BuilderRegistry) Create(name string) (interface{}, error) {
	fn, ok := r.Get(name)
	if !ok {
		return nil, fmt.Errorf("builder not found: %s", name)
	}
	return fn(), nil
}

// DefaultRegistry is the default builder registry
var DefaultRegistry = NewBuilderRegistry()

// Register registers a builder function in the default registry
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

// Create creates a new builder by name from the default registry
func Create(name string) (interface{}, error) {
	return DefaultRegistry.Create(name)
}
