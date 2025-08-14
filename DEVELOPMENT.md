# Development Guide

This guide covers development practices and workflows for the Go Builder Kit project.

## Prerequisites

- Go 1.22+ (recommended: Go 1.23+)
- Git
- Make (optional, but recommended)

## Development Setup

1. **Clone the repository:**
   ```bash
   git clone https://github.com/adil-faiyaz98/go-builder-kit.git
   cd go-builder-kit
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Install development tools:**
   ```bash
   make install-tools
   # OR manually:
   go install honnef.co/go/tools/cmd/staticcheck@latest
   go install golang.org/x/lint/golint@latest
   go install golang.org/x/vuln/cmd/govulncheck@latest
   ```

## Development Workflow

### Using Make (Recommended)

The project includes a comprehensive Makefile with common development tasks:

```bash
# Run all checks and build
make all

# Run tests
make test

# Run tests with coverage
make coverage

# Format code
make fmt

# Run linters
make lint

# Build binary
make build

# Run security scan
make security

# Generate builders for examples
make generate

# Run full CI pipeline locally
make ci

# Clean build artifacts
make clean
```

### Manual Commands

If you prefer not to use Make:

```bash
# Format code
go fmt ./...

# Run static analysis
go vet ./...
staticcheck ./...

# Run tests
go test -v -race ./...

# Run tests with coverage
go test -v -race -coverprofile=coverage.out ./...
go tool cover -func=coverage.out

# Build
go build -ldflags "-s -w" -o bin/builder-gen ./cmd/builder-gen

# Security scan
govulncheck ./...
```

## Code Quality Standards

### Go Best Practices

1. **Formatting**: All code must be formatted with `go fmt`
2. **Linting**: Code must pass `go vet`, `staticcheck`, and `golint`
3. **Testing**: Minimum 70% test coverage (current: ~72%)
4. **Documentation**: All public APIs must have comprehensive documentation
5. **Error Handling**: Proper error handling with meaningful error messages
6. **Security**: All inputs must be validated and sanitized

### Naming Conventions

- Use clear, descriptive names
- Avoid stuttering (e.g., `builder.Registry` not `builder.BuilderRegistry`)
- Follow Go naming conventions (exported vs unexported)
- Use consistent naming across the codebase

### Testing Standards

1. **Coverage**: Maintain minimum 70% test coverage
2. **Race Conditions**: All tests must pass with `-race` flag
3. **Table-Driven Tests**: Use table-driven tests where appropriate
4. **Subtests**: Use subtests for logical grouping
5. **Cleanup**: Properly clean up test resources

### Security Standards

1. **Input Validation**: Validate all external inputs
2. **String Sanitization**: Sanitize strings to prevent injection
3. **Path Validation**: Validate file paths to prevent traversal
4. **Dependency Updates**: Keep dependencies up to date
5. **Vulnerability Scanning**: Run `govulncheck` before releases

## Project Structure

```
├── cmd/builder-gen/     # Main command-line tool
├── pkg/                 # Core library packages
│   ├── builder/         # Builder framework
│   └── generator/       # Code generation
├── models/              # Example model structs
├── builders/            # Generated builder implementations
├── tests/               # Integration tests
├── examples/            # Usage examples
├── .github/workflows/   # CI/CD configuration
└── docs/                # Documentation
```

### Package Guidelines

- **cmd/**: Command-line tools and applications
- **pkg/**: Reusable library code
- **internal/**: Private packages (if needed)
- **tests/**: Integration and end-to-end tests
- **examples/**: Usage examples and demos

## Contributing Workflow

### Branch Strategy

- `main`: Stable release branch
- `develop`: Development integration branch (if used)
- `feature/*`: Feature development branches
- `fix/*`: Bug fix branches
- `release/*`: Release preparation branches

### Pull Request Process

1. **Create Feature Branch:**
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Develop and Test:**
   ```bash
   # Make your changes
   make ci  # Run full CI pipeline locally
   ```

3. **Commit Changes:**
   ```bash
   git add .
   git commit -m "feat: add your feature description"
   ```

4. **Push and Create PR:**
   ```bash
   git push origin feature/your-feature-name
   # Create pull request via GitHub
   ```

### Commit Message Format

Follow [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` New features
- `fix:` Bug fixes
- `docs:` Documentation changes
- `style:` Code style changes (formatting, etc.)
- `refactor:` Code refactoring
- `perf:` Performance improvements
- `test:` Adding or updating tests
- `ci:` CI/CD changes
- `chore:` Maintenance tasks

## Release Process

### Preparing a Release

1. **Update Version Numbers:**
   - Update `go.mod` if needed
   - Update `CHANGELOG.md`
   - Update version references in documentation

2. **Run Release Checks:**
   ```bash
   make release-check
   ```

3. **Create Release Tag:**
   ```bash
   git tag -a v2.1.0 -m "Release v2.1.0"
   git push origin v2.1.0
   ```

### Automated Release

The release process is automated via GitHub Actions:

1. **Tag Creation** triggers the release workflow
2. **Tests and Security Scans** run automatically
3. **Cross-Platform Binaries** are built
4. **GitHub Release** is created with artifacts
5. **Go Package Registry** is updated

## Debugging and Troubleshooting

### Common Issues

1. **Test Failures:**
   ```bash
   # Run specific test
   go test -v ./pkg/builder -run TestSpecific
   
   # Run with race detection
   go test -race ./...
   ```

2. **Build Issues:**
   ```bash
   # Clean module cache
   go clean -modcache
   go mod download
   ```

3. **Linting Issues:**
   ```bash
   # Fix formatting
   go fmt ./...
   
   # Check specific issues
   staticcheck ./path/to/package
   ```

### Performance Profiling

```bash
# CPU profiling
go test -cpuprofile=cpu.prof -bench=.

# Memory profiling
go test -memprofile=mem.prof -bench=.

# View profiles
go tool pprof cpu.prof
go tool pprof mem.prof
```

## IDE Configuration

### VS Code

Recommended extensions:
- Go extension (official)
- Go Test Explorer
- Code Spell Checker

Settings (`.vscode/settings.json`):
```json
{
  "go.lintTool": "staticcheck",
  "go.testFlags": ["-v", "-race"],
  "go.buildFlags": ["-v"],
  "editor.formatOnSave": true
}
```

### GoLand/IntelliJ

- Enable Go modules support
- Configure code style to use gofmt
- Enable static analysis tools
- Set up test configurations with race detection

## Resources

- [Go Documentation](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Go Best Practices](https://golang.org/doc/effective_go.html)
- [Security Guidelines](https://golang.org/doc/security.html)
