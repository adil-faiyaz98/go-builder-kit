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
	if b == nil {
		var zero T
		return zero
	}

	b.cacheLock.RLock()
	if !b.dirty && b.cache != nil {
		result := *b.cache
		b.cacheLock.RUnlock()
		return result
	}
	b.cacheLock.RUnlock()

	b.cacheLock.Lock()
	defer b.cacheLock.Unlock()

	// Double-check pattern to avoid race conditions
	if !b.dirty && b.cache != nil {
		return *b.cache
	}

	if b.GenericBuilder == nil {
		var zero T
		return zero
	}

	obj := b.GenericBuilder.Build()
	b.cache = &obj
	b.dirty = false
	return obj
}

// BuildPtr builds the object and returns a pointer to it
func (b *CachedBuilder[T]) BuildPtr() *T {
	if b == nil {
		return nil
	}

	b.cacheLock.RLock()
	if !b.dirty && b.cache != nil {
		result := b.cache
		b.cacheLock.RUnlock()
		return result
	}
	b.cacheLock.RUnlock()

	b.cacheLock.Lock()
	defer b.cacheLock.Unlock()

	// Double-check pattern to avoid race conditions
	if !b.dirty && b.cache != nil {
		return b.cache
	}

	if b.GenericBuilder == nil {
		return nil
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
