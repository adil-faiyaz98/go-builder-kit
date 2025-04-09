# Contributing to Go Builder Kit

Thank you for considering contributing to Go Builder Kit! This document provides guidelines and instructions for contributing.

## Code of Conduct

By participating in this project, you agree to abide by our [Code of Conduct](CODE_OF_CONDUCT.md).

## How Can I Contribute?

### Reporting Bugs

- Check if the bug has already been reported in the [Issues](https://github.com/adil-faiyaz98/go-builder-kit/issues)
- If not, create a new issue with a clear title and description
- Include steps to reproduce, expected behavior, and actual behavior
- Include code samples, error messages, and Go version information

### Suggesting Enhancements

- Check if the enhancement has already been suggested in the [Issues](https://github.com/adil-faiyaz98/go-builder-kit/issues)
- If not, create a new issue with a clear title and description
- Explain why this enhancement would be useful to most users
- Include any relevant code examples or documentation

### Pull Requests

1. Fork the repository
2. Create a new branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests (`go test ./...`)
5. Run linters (`golangci-lint run`)
6. Commit your changes (`git commit -m 'Add some amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

## Development Workflow

### Setting Up Development Environment

1. Install Go (version 1.19 or later)
2. Clone the repository
3. Install dependencies: `go mod download`
4. Install development tools:
   ```bash
   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
   go install golang.org/x/tools/cmd/goimports@latest
   ```

### Running Tests

```bash
go test ./...
```

For tests with coverage:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Code Style

- Follow standard Go code style and conventions
- Use `goimports` to format your code
- Run `golangci-lint run` before submitting a PR

## Release Process

1. Update version in relevant files
2. Update CHANGELOG.md
3. Create a new tag: `git tag v1.x.x`
4. Push the tag: `git push origin v1.x.x`
5. GitHub Actions will automatically create a release

## Documentation

- Update documentation for any new features or changes
- Include examples in the documentation
- Keep the README.md up to date

## Questions?

If you have any questions, feel free to open an issue or contact the maintainers directly.
