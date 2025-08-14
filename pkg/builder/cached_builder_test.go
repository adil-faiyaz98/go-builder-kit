package builder

import (
	"sync"
	"testing"
	"time"
)

func TestCachedBuilder(t *testing.T) {
	t.Run("Build caches result", func(t *testing.T) {
		callCount := 0
		baseBuilder := NewBaseGenericBuilder(func() TestPerson {
			callCount++
			return TestPerson{Name: "John", Age: 30}
		})

		cachedBuilder := NewCachedBuilder(baseBuilder)

		// First call should execute build function
		result1 := cachedBuilder.Build()
		if callCount != 1 {
			t.Errorf("Expected 1 call, got %d", callCount)
		}

		// Second call should use cache
		result2 := cachedBuilder.Build()
		if callCount != 1 {
			t.Errorf("Expected 1 call (cached), got %d", callCount)
		}

		if result1.Name != result2.Name || result1.Age != result2.Age {
			t.Error("Cached results should be identical")
		}
	})

	t.Run("BuildPtr caches result", func(t *testing.T) {
		callCount := 0
		baseBuilder := NewBaseGenericBuilder(func() TestPerson {
			callCount++
			return TestPerson{Name: "Jane", Age: 25}
		})

		cachedBuilder := NewCachedBuilder(baseBuilder)

		// First call should execute build function
		result1 := cachedBuilder.BuildPtr()
		if callCount != 1 {
			t.Errorf("Expected 1 call, got %d", callCount)
		}

		// Second call should use cache
		result2 := cachedBuilder.BuildPtr()
		if callCount != 1 {
			t.Errorf("Expected 1 call (cached), got %d", callCount)
		}

		// Should return same pointer
		if result1 != result2 {
			t.Error("BuildPtr should return same cached pointer")
		}
	})

	t.Run("Invalidate clears cache", func(t *testing.T) {
		callCount := 0
		baseBuilder := NewBaseGenericBuilder(func() TestPerson {
			callCount++
			return TestPerson{Name: "Bob", Age: 40}
		})

		cachedBuilder := NewCachedBuilder(baseBuilder)

		// First call
		cachedBuilder.Build()
		if callCount != 1 {
			t.Errorf("Expected 1 call, got %d", callCount)
		}

		// Invalidate cache
		cachedBuilder.Invalidate()

		// Next call should rebuild
		cachedBuilder.Build()
		if callCount != 2 {
			t.Errorf("Expected 2 calls after invalidation, got %d", callCount)
		}
	})

	t.Run("Thread safety", func(t *testing.T) {
		callCount := 0
		var mu sync.Mutex
		baseBuilder := NewBaseGenericBuilder(func() TestPerson {
			mu.Lock()
			callCount++
			mu.Unlock()
			// Add small delay to increase chance of race condition
			time.Sleep(time.Millisecond)
			return TestPerson{Name: "Concurrent", Age: 50}
		})

		cachedBuilder := NewCachedBuilder(baseBuilder)

		// Run multiple goroutines concurrently
		const numGoroutines = 10
		var wg sync.WaitGroup
		results := make([]TestPerson, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(index int) {
				defer wg.Done()
				results[index] = cachedBuilder.Build()
			}(i)
		}

		wg.Wait()

		// Build function should only be called once due to caching
		mu.Lock()
		finalCallCount := callCount
		mu.Unlock()

		if finalCallCount != 1 {
			t.Errorf("Expected 1 call with caching, got %d", finalCallCount)
		}

		// All results should be identical
		for i := 1; i < numGoroutines; i++ {
			if results[i].Name != results[0].Name || results[i].Age != results[0].Age {
				t.Error("All cached results should be identical")
				break
			}
		}
	})

	t.Run("Nil builder handling", func(t *testing.T) {
		var cachedBuilder *CachedBuilder[TestPerson]

		// Should not panic
		result := cachedBuilder.Build()
		if result.Name != "" || result.Age != 0 {
			t.Error("Expected zero value for nil builder")
		}

		ptr := cachedBuilder.BuildPtr()
		if ptr != nil {
			t.Error("Expected nil pointer for nil builder")
		}
	})

	t.Run("Nil base builder handling", func(t *testing.T) {
		cachedBuilder := NewCachedBuilder[TestPerson](nil)

		// Should not panic
		result := cachedBuilder.Build()
		if result.Name != "" || result.Age != 0 {
			t.Error("Expected zero value for nil base builder")
		}

		ptr := cachedBuilder.BuildPtr()
		if ptr != nil {
			t.Error("Expected nil pointer for nil base builder")
		}
	})

	t.Run("Cache invalidation during concurrent access", func(t *testing.T) {
		callCount := 0
		var mu sync.Mutex
		baseBuilder := NewBaseGenericBuilder(func() TestPerson {
			mu.Lock()
			callCount++
			count := callCount
			mu.Unlock()
			return TestPerson{Name: "Test", Age: count}
		})

		cachedBuilder := NewCachedBuilder(baseBuilder)

		// Build once to populate cache
		first := cachedBuilder.Build()
		if first.Age != 1 {
			t.Errorf("Expected age 1, got %d", first.Age)
		}

		// Start goroutine that will invalidate cache
		go func() {
			time.Sleep(10 * time.Millisecond)
			cachedBuilder.Invalidate()
		}()

		// Keep building in a loop
		var lastAge int
		for i := 0; i < 100; i++ {
			result := cachedBuilder.Build()
			if result.Age > lastAge {
				lastAge = result.Age
			}
			time.Sleep(time.Millisecond)
		}

		// Should have been invalidated at least once
		if lastAge <= 1 {
			t.Error("Cache should have been invalidated and rebuilt")
		}
	})
}

func TestNewCachedBuilder(t *testing.T) {
	baseBuilder := NewBaseGenericBuilder(func() TestPerson {
		return TestPerson{Name: "Test", Age: 1}
	})

	cachedBuilder := NewCachedBuilder(baseBuilder)

	if cachedBuilder == nil {
		t.Fatal("NewCachedBuilder returned nil")
	}

	if cachedBuilder.GenericBuilder != baseBuilder {
		t.Error("CachedBuilder should wrap the provided builder")
	}

	// Should start as dirty
	if !cachedBuilder.dirty {
		t.Error("New cached builder should start as dirty")
	}

	if cachedBuilder.cache != nil {
		t.Error("New cached builder should start with nil cache")
	}
}
