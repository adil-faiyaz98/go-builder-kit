package main

import (
	"fmt"
	"time"

	"github.com/adil-faiyaz98/go-builder-kit/pkg/builder"
)

// ComplexObject is a complex object that takes time to build
type ComplexObject struct {
	ID     string
	Name   string
	Data   []int
	Matrix [][]int
}

// ComplexObjectBuilder is a builder for ComplexObject
type ComplexObjectBuilder struct {
	*builder.BaseGenericBuilder[ComplexObject]
	obj ComplexObject
}

// NewComplexObjectBuilder creates a new ComplexObjectBuilder
func NewComplexObjectBuilder() *ComplexObjectBuilder {
	b := &ComplexObjectBuilder{
		obj: ComplexObject{
			Data:   make([]int, 0),
			Matrix: make([][]int, 0),
		},
	}

	b.BaseGenericBuilder = builder.NewBaseGenericBuilder(func() ComplexObject {
		// Simulate a complex build process
		time.Sleep(100 * time.Millisecond)
		return b.obj
	})

	return b
}

// WithID sets the ID of the object
func (b *ComplexObjectBuilder) WithID(id string) *ComplexObjectBuilder {
	b.obj.ID = id
	return b
}

// WithName sets the name of the object
func (b *ComplexObjectBuilder) WithName(name string) *ComplexObjectBuilder {
	b.obj.Name = name
	return b
}

// WithData sets the data of the object
func (b *ComplexObjectBuilder) WithData(data []int) *ComplexObjectBuilder {
	b.obj.Data = make([]int, len(data))
	copy(b.obj.Data, data)
	return b
}

// WithMatrix sets the matrix of the object
func (b *ComplexObjectBuilder) WithMatrix(matrix [][]int) *ComplexObjectBuilder {
	b.obj.Matrix = make([][]int, len(matrix))
	for i, row := range matrix {
		b.obj.Matrix[i] = make([]int, len(row))
		copy(b.obj.Matrix[i], row)
	}
	return b
}

func main() {
	// Create a complex object builder
	builder := NewComplexObjectBuilder().
		WithID("123").
		WithName("Complex Object").
		WithData([]int{1, 2, 3, 4, 5}).
		WithMatrix([][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		})

	// Create a cached builder wrapper
	cachedBuilder := &struct {
		builder *ComplexObjectBuilder
		cache   *ComplexObject
		dirty   bool
	}{
		builder: builder,
		dirty:   true,
	}

	// Define a Build method that uses caching
	buildWithCache := func() ComplexObject {
		if !cachedBuilder.dirty && cachedBuilder.cache != nil {
			return *cachedBuilder.cache
		}

		obj := builder.Build()
		cachedBuilder.cache = &obj
		cachedBuilder.dirty = false
		return obj
	}

	// Define an Invalidate method
	invalidateCache := func() {
		cachedBuilder.dirty = true
		cachedBuilder.cache = nil
	}

	// Measure time without caching
	start := time.Now()
	for i := 0; i < 10; i++ {
		obj := builder.Build()
		_ = obj
	}
	fmt.Printf("Time without caching: %v\n", time.Since(start))

	// Measure time with caching
	start = time.Now()
	for i := 0; i < 10; i++ {
		obj := buildWithCache()
		_ = obj
	}
	fmt.Printf("Time with caching: %v\n", time.Since(start))

	// Invalidate the cache and modify the object
	invalidateCache()
	builder.WithName("Modified Object")

	// Build again
	obj := buildWithCache()
	fmt.Println("Modified object name:", obj.Name)
}
