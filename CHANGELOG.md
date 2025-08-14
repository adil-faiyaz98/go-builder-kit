# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Comprehensive Makefile with development tasks
- Enhanced CI/CD pipeline with cross-platform builds, security scanning, and coverage reporting
- Better Go best practices compliance with improved naming conventions
- Enhanced .gitignore with more comprehensive coverage
- Automated dependency security scanning with govulncheck

### Changed
- Improved naming conventions to reduce stuttering (BuilderFunc -> Func, BuilderRegistry -> Registry)
- Enhanced CI/CD workflows with multi-platform testing and security scanning
- Better error handling and code documentation
- Updated dependencies to latest secure versions

### Fixed
- Fixed golint warnings about stuttering type names
- Improved test coverage and error handling

### Security
- Added vulnerability scanning to CI/CD pipeline
- Updated all dependencies to latest secure versions
- Enhanced security documentation

## [v2.0.0] - Previous Release

### Added
- Security-first design with comprehensive input validation and sanitization
- High-performance optimizations with smart caching and memory efficiency
- Full Go 1.23+ generics support with type-safe builders
- Thread-safe operations with proper synchronization
- Comprehensive validation framework with custom validators
- Automatic builder code generation from struct definitions
- Deep cloning functionality for complex structures
- Builder registry for centralized builder management

### Security
- Input validation and sanitization for all string inputs
- Path traversal protection for file operations
- Memory safety with nil-safe operations and bounds checking
- Thread safety with concurrent-safe operations
- Protection against injection attacks in code generation

### Performance
- 3x faster builder creation compared to previous versions
- 50% reduction in memory allocations
- 4x faster validation speed
- 2x faster clone operations
- Pre-allocated slices and optimized string operations

### Documentation
- Comprehensive README with security features, performance metrics, and usage examples
- Complete API reference documentation
- Migration guide from v1.x
- Best practices guidelines for security and performance
