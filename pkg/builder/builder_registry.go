package builder

import (
	"fmt"
	"sync"
)

// Func is a function that creates a new builder
type Func func() interface{}

// Registry is a registry of builder functions
type Registry struct {
	builders map[string]Func
	mu       sync.RWMutex
}

// NewBuilderRegistry creates a new Registry
func NewBuilderRegistry() *Registry {
	return &Registry{
		builders: make(map[string]Func),
	}
}

// Register registers a builder function
func (r *Registry) Register(name string, fn Func) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.builders[name] = fn
}

// Get returns a builder function by name
func (r *Registry) Get(name string) (Func, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	fn, ok := r.builders[name]
	return fn, ok
}

// GetAll returns all registered builder functions
func (r *Registry) GetAll() map[string]Func {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Create a copy to avoid race conditions
	result := make(map[string]Func, len(r.builders))
	for k, v := range r.builders {
		result[k] = v
	}

	return result
}

// Create creates a new builder by name
func (r *Registry) Create(name string) (interface{}, error) {
	fn, ok := r.Get(name)
	if !ok {
		return nil, fmt.Errorf("builder not found: %s", name)
	}
	return fn(), nil
}

// DefaultRegistry is the default builder registry
var DefaultRegistry = NewBuilderRegistry()

// Register registers a builder function in the default registry
func Register(name string, fn Func) {
	DefaultRegistry.Register(name, fn)
}

// Get returns a builder function by name from the default registry
func Get(name string) (Func, bool) {
	return DefaultRegistry.Get(name)
}

// GetAll returns all registered builder functions from the default registry
func GetAll() map[string]Func {
	return DefaultRegistry.GetAll()
}

// Create creates a new builder by name from the default registry
func Create(name string) (interface{}, error) {
	return DefaultRegistry.Create(name)
}
