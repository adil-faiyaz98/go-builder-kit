package builder

import (
	"sync"
	"testing"
)

func TestBuilderRegistry(t *testing.T) {
	t.Run("Register and Get builder", func(t *testing.T) {
		registry := NewBuilderRegistry()

		// Register a builder
		registry.Register("test", func() any {
			return &TestPerson{Name: "Test", Age: 1}
		})

		// Get the builder
		builderFunc, ok := registry.Get("test")
		if !ok {
			t.Error("Expected to find registered builder")
		}

		if builderFunc == nil {
			t.Error("Expected non-nil builder function")
		}

		// Use the builder
		result := builderFunc()
		person, ok := result.(*TestPerson)
		if !ok {
			t.Error("Expected TestPerson from builder")
		}

		if person.Name != "Test" || person.Age != 1 {
			t.Errorf("Expected {Test 1}, got %+v", person)
		}
	})

	t.Run("Get non-existent builder", func(t *testing.T) {
		registry := NewBuilderRegistry()

		builderFunc, ok := registry.Get("non-existent")
		if ok {
			t.Error("Expected not to find non-existent builder")
		}

		if builderFunc != nil {
			t.Error("Expected nil builder function for non-existent builder")
		}
	})

	t.Run("GetAll returns all builders", func(t *testing.T) {
		registry := NewBuilderRegistry()

		// Register multiple builders
		registry.Register("test1", func() any { return 1 })
		registry.Register("test2", func() any { return 2 })

		all := registry.GetAll()
		if len(all) != 2 {
			t.Errorf("Expected 2 builders, got %d", len(all))
		}

		if _, ok := all["test1"]; !ok {
			t.Error("Expected to find test1 in GetAll")
		}

		if _, ok := all["test2"]; !ok {
			t.Error("Expected to find test2 in GetAll")
		}
	})

	t.Run("Create builder by name", func(t *testing.T) {
		registry := NewBuilderRegistry()

		registry.Register("person", func() any {
			return &TestPerson{Name: "Created", Age: 42}
		})

		result, err := registry.Create("person")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		person, ok := result.(*TestPerson)
		if !ok {
			t.Error("Expected TestPerson from Create")
		}

		if person.Name != "Created" || person.Age != 42 {
			t.Errorf("Expected {Created 42}, got %+v", person)
		}
	})

	t.Run("Create non-existent builder", func(t *testing.T) {
		registry := NewBuilderRegistry()

		result, err := registry.Create("non-existent")
		if err == nil {
			t.Error("Expected error for non-existent builder")
		}

		if result != nil {
			t.Error("Expected nil result for non-existent builder")
		}

		expectedError := "builder not found: non-existent"
		if err.Error() != expectedError {
			t.Errorf("Expected error %q, got %q", expectedError, err.Error())
		}
	})

	t.Run("Override existing builder", func(t *testing.T) {
		registry := NewBuilderRegistry()

		// Register initial builder
		registry.Register("test", func() any {
			return "original"
		})

		// Override with new builder
		registry.Register("test", func() any {
			return "override"
		})

		builderFunc, ok := registry.Get("test")
		if !ok {
			t.Error("Expected to find builder after override")
		}

		result := builderFunc()
		if result != "override" {
			t.Errorf("Expected 'override', got %v", result)
		}
	})

	t.Run("Thread safety", func(t *testing.T) {
		registry := NewBuilderRegistry()
		const numGoroutines = 100

		var wg sync.WaitGroup

		// Concurrent registration
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(index int) {
				defer wg.Done()
				registry.Register("test", func() any {
					return index
				})
			}(i)
		}

		// Concurrent reading
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				registry.Get("test")
				registry.GetAll()
			}()
		}

		wg.Wait()

		// Should not panic and should have a builder registered
		_, ok := registry.Get("test")
		if !ok {
			t.Error("Expected to find builder after concurrent operations")
		}
	})
}

func TestDefaultRegistry(t *testing.T) {
	// Note: These tests might affect global state, so they should be isolated
	// or run in a separate test process in a real application

	t.Run("Register and Get from default registry", func(t *testing.T) {
		Register("default-test", func() any {
			return "default-value"
		})

		builderFunc, ok := Get("default-test")
		if !ok {
			t.Error("Expected to find builder in default registry")
		}

		result := builderFunc()
		if result != "default-value" {
			t.Errorf("Expected 'default-value', got %v", result)
		}
	})

	t.Run("GetAll from default registry", func(t *testing.T) {
		// Register a test builder
		Register("getall-test", func() any { return "test" })

		all := GetAll()
		if len(all) == 0 {
			t.Error("Expected at least one builder in default registry")
		}

		if _, ok := all["getall-test"]; !ok {
			t.Error("Expected to find getall-test in default registry")
		}
	})

	t.Run("Create from default registry", func(t *testing.T) {
		Register("create-test", func() any {
			return &TestPerson{Name: "Default", Age: 100}
		})

		result, err := Create("create-test")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		person, ok := result.(*TestPerson)
		if !ok {
			t.Error("Expected TestPerson from default registry Create")
		}

		if person.Name != "Default" || person.Age != 100 {
			t.Errorf("Expected {Default 100}, got %+v", person)
		}
	})

	t.Run("Create non-existent from default registry", func(t *testing.T) {
		result, err := Create("definitely-non-existent")
		if err == nil {
			t.Error("Expected error for non-existent builder in default registry")
		}

		if result != nil {
			t.Error("Expected nil result for non-existent builder in default registry")
		}
	})
}

func TestNewBuilderRegistry(t *testing.T) {
	registry := NewBuilderRegistry()

	if registry == nil {
		t.Fatal("NewBuilderRegistry returned nil")
	}

	if registry.builders == nil {
		t.Error("Registry builders map should be initialized")
	}

	// Should start empty
	all := registry.GetAll()
	if len(all) != 0 {
		t.Error("New registry should start empty")
	}
}
