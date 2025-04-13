package builder

import (
	"sync"
)

// CachedBuilder is a builder that caches the built object
type CachedBuilder[T any] struct {
	GenericBuilder[T]
	
	// cache is the cached object
	cache     *T
	cacheLock sync.RWMutex
	dirty     bool
}

// NewCachedBuilder creates a new CachedBuilder
func NewCachedBuilder[T any](builder GenericBuilder[T]) *CachedBuilder[T] {
	return &CachedBuilder[T]{
		GenericBuilder: builder,
		dirty:          true,
	}
}

// Build builds the object and returns it
func (b *CachedBuilder[T]) Build() T {
	b.cacheLock.RLock()
	if !b.dirty && b.cache != nil {
		defer b.cacheLock.RUnlock()
		return *b.cache
	}
	b.cacheLock.RUnlock()
	
	b.cacheLock.Lock()
	defer b.cacheLock.Unlock()
	
	// Check again in case another goroutine updated the cache
	if !b.dirty && b.cache != nil {
		return *b.cache
	}
	
	obj := b.GenericBuilder.Build()
	b.cache = &obj
	b.dirty = false
	return obj
}

// BuildPtr builds the object and returns a pointer to it
func (b *CachedBuilder[T]) BuildPtr() *T {
	b.cacheLock.RLock()
	if !b.dirty && b.cache != nil {
		defer b.cacheLock.RUnlock()
		return b.cache
	}
	b.cacheLock.RUnlock()
	
	b.cacheLock.Lock()
	defer b.cacheLock.Unlock()
	
	// Check again in case another goroutine updated the cache
	if !b.dirty && b.cache != nil {
		return b.cache
	}
	
	obj := b.GenericBuilder.Build()
	b.cache = &obj
	b.dirty = false
	return b.cache
}

// Invalidate invalidates the cache
func (b *CachedBuilder[T]) Invalidate() {
	b.cacheLock.Lock()
	defer b.cacheLock.Unlock()
	
	b.dirty = true
	b.cache = nil
}

// WithBuilder sets the underlying builder
func (b *CachedBuilder[T]) WithBuilder(builder GenericBuilder[T]) *CachedBuilder[T] {
	b.cacheLock.Lock()
	defer b.cacheLock.Unlock()
	
	b.GenericBuilder = builder
	b.dirty = true
	b.cache = nil
	return b
}
