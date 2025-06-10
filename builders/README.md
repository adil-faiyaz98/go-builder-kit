# go-builder-kit/builders

This package provides a set of generic builder utilities to help you construct and configure Go structs in a flexible, readable, and composable way. It follows the "builder pattern," making it easier to create complex objects with customizable options, especially useful for testing, configuration, and fluent API design.

## Features

- **Generic Builders**: Compose and apply multiple configuration options to struct instances.
- **Composable Options**: Easily combine and reuse option functions.
- **Type Safety**: Leverage Go generics for type-safe builder utilities.

## Quick Start

### 1. Import the Package

```go
import "github.com/adil-faiyaz98/go-builder-kit/builders"
```

### 2. Define Your Struct and Options

```go
type User struct {
    Name  string
    Email string
    Age   int
}

// Create option functions using builders.Option
func WithName(name string) builders.Option[User] {
    return func(u *User) { u.Name = name }
}

func WithEmail(email string) builders.Option[User] {
    return func(u *User) { u.Email = email }
}

func WithAge(age int) builders.Option[User] {
    return func(u *User) { u.Age = age }
}
```

### 3. Build Instances Easily

```go
user := builders.Build(
    User{},
    WithName("Alice"),
    WithEmail("alice@example.com"),
    WithAge(30),
)
```

### 4. Compose and Reuse Options

```go
basicUser := builders.Compose(
    WithName("Bob"),
    WithEmail("bob@example.com"),
)

user := builders.Build(User{}, basicUser, WithAge(25))
```

## API Overview

- `Option[T]`: Type alias for a function that modifies a pointer to T.
- `Build[T]`: Constructs a value of type T, applying all given options in order.
- `Compose[T]`: Composes multiple options into a single option for reuse.

## Example

```go
type Config struct {
    Debug bool
    Port  int
}

cfg := builders.Build(
    Config{},
    func(c *Config) { c.Debug = true },
    func(c *Config) { c.Port = 8080 },
)
```

## When to Use

- When you want a flexible, readable way to configure struct instances.
- When testing: easily create test fixtures with variations.
- When exposing fluent configuration APIs.

## Source & Docs

- [Source on GitHub](https://github.com/adil-faiyaz98/go-builder-kit/tree/main/builders)
- [API Documentation](https://pkg.go.dev/github.com/adil-faiyaz98/go-builder-kit/builders)

---

> **Tip:** See the [`builders` package documentation](https://pkg.go.dev/github.com/adil-faiyaz98/go-builder-kit/builders) for more advanced usage and composition examples.
