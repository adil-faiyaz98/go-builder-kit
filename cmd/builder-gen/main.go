package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/adil-faiyaz98/go-builder-kit/pkg/generator"
)

func main() {
	// Parse command line flags
	inputPath := flag.String("input", "", "Path to the input Go file or directory containing structs to generate builders for")
	outputPath := flag.String("output", "", "Output directory for generated builder files")
	modelsPackage := flag.String("models-package", "", "Import path for the models package")
	packageName := flag.String("package-name", "builders", "Name of the generated package")
	verbose := flag.Bool("verbose", false, "Enable verbose output")
	help := flag.Bool("help", false, "Show help")

	flag.Parse()

	// Show help if requested or if required flags are missing
	if *help || *inputPath == "" || *outputPath == "" || *modelsPackage == "" {
		printUsage()
		os.Exit(0)
	}

	// Validate and sanitize inputs
	if err := validateInputs(*inputPath, *outputPath, *modelsPackage, *packageName); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Sanitize paths to prevent directory traversal attacks
	*inputPath = filepath.Clean(*inputPath)
	*outputPath = filepath.Clean(*outputPath)

	// Create generator options
	opts := generator.Options{
		PackageName:   *packageName,
		ModelsPackage: *modelsPackage,
		Verbose:       *verbose,
	}

	// Create generator
	gen := generator.NewGenerator(opts)

	// Check if input path exists
	if _, err := os.Stat(*inputPath); os.IsNotExist(err) {
		fmt.Printf("Error: Input path %s does not exist\n", *inputPath)
		os.Exit(1)
	}

	// Create output directory if it doesn't exist
	if _, err := os.Stat(*outputPath); os.IsNotExist(err) {
		if err := os.MkdirAll(*outputPath, 0755); err != nil {
			fmt.Printf("Error creating output directory: %v\n", err)
			os.Exit(1)
		}
	}

	// Process input path
	if err := processPath(*inputPath, *outputPath, gen); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Builder generation completed successfully!")
}

// processPath processes the input path and generates builders
func processPath(inputPath, outputPath string, gen *generator.Generator) error {
	// Additional security check - ensure paths are absolute or relative without traversal
	if !filepath.IsAbs(inputPath) && strings.Contains(inputPath, "..") {
		return fmt.Errorf("relative input path contains directory traversal")
	}
	if !filepath.IsAbs(outputPath) && strings.Contains(outputPath, "..") {
		return fmt.Errorf("relative output path contains directory traversal")
	}

	// Check if input path is a directory
	fileInfo, err := os.Stat(inputPath)
	if err != nil {
		return fmt.Errorf("failed to stat input path %q: %w", inputPath, err)
	}

	var processedFiles int

	if fileInfo.IsDir() {
		// Process all Go files in the directory
		files, err := os.ReadDir(inputPath)
		if err != nil {
			return fmt.Errorf("failed to read directory %q: %w", inputPath, err)
		}

		for _, file := range files {
			if filepath.Ext(file.Name()) == ".go" && !file.IsDir() {
				// Skip test files and files starting with underscore
				if strings.HasSuffix(file.Name(), "_test.go") || strings.HasPrefix(file.Name(), "_") {
					continue
				}

				filePath := filepath.Join(inputPath, file.Name())
				if err := gen.GenerateBuilders(filePath, outputPath); err != nil {
					return fmt.Errorf("failed to generate builders for %q: %w", filePath, err)
				}
				processedFiles++
			}
		}

		if processedFiles == 0 {
			return fmt.Errorf("no valid Go files found in directory %q", inputPath)
		}
	} else {
		// Process single file
		if filepath.Ext(inputPath) != ".go" {
			return fmt.Errorf("input file %q must be a Go file (.go extension)", inputPath)
		}

		// Skip test files
		if strings.HasSuffix(inputPath, "_test.go") {
			return fmt.Errorf("test files are not supported for builder generation")
		}

		if err := gen.GenerateBuilders(inputPath, outputPath); err != nil {
			return fmt.Errorf("failed to generate builders for %q: %w", inputPath, err)
		}
		processedFiles = 1
	}

	fmt.Printf("Successfully processed %d file(s)\n", processedFiles)
	return nil
}

// validateInputs validates all input parameters for security and correctness
func validateInputs(inputPath, outputPath, modelsPackage, packageName string) error {
	// Validate input path
	if strings.Contains(inputPath, "..") {
		return fmt.Errorf("input path contains directory traversal sequences")
	}

	// Validate output path
	if strings.Contains(outputPath, "..") {
		return fmt.Errorf("output path contains directory traversal sequences")
	}

	// Validate package name
	if err := validatePackageName(packageName); err != nil {
		return fmt.Errorf("invalid package name: %w", err)
	}

	// Validate models package path
	if err := validateModelsPackage(modelsPackage); err != nil {
		return fmt.Errorf("invalid models package: %w", err)
	}

	return nil
}

// validatePackageName validates that the package name is a valid Go package name
func validatePackageName(name string) error {
	if name == "" {
		return fmt.Errorf("package name cannot be empty")
	}

	// Go package names must be valid identifiers
	validPackageName := regexp.MustCompile(`^[a-z][a-z0-9_]*$`)
	if !validPackageName.MatchString(name) {
		return fmt.Errorf("package name must be a valid Go identifier (lowercase letters, numbers, underscores)")
	}

	// Check for reserved keywords
	reservedKeywords := []string{
		"break", "case", "chan", "const", "continue", "default", "defer", "else",
		"fallthrough", "for", "func", "go", "goto", "if", "import", "interface",
		"map", "package", "range", "return", "select", "struct", "switch", "type", "var",
	}

	for _, keyword := range reservedKeywords {
		if name == keyword {
			return fmt.Errorf("package name cannot be a Go reserved keyword: %s", keyword)
		}
	}

	return nil
}

// validateModelsPackage validates the models package import path
func validateModelsPackage(pkg string) error {
	if pkg == "" {
		return fmt.Errorf("models package cannot be empty")
	}

	// Basic validation for import path format
	if strings.Contains(pkg, "..") {
		return fmt.Errorf("models package path contains directory traversal sequences")
	}

	// Check for suspicious characters
	if strings.ContainsAny(pkg, "<>|\"*?") {
		return fmt.Errorf("models package path contains invalid characters")
	}

	return nil
}

// printUsage prints the usage information
func printUsage() {
	fmt.Println("Usage: builder-gen [options]")
	fmt.Println("\nOptions:")
	flag.PrintDefaults()
	fmt.Println("\nExample:")
	fmt.Println("  builder-gen -input path/to/models -output path/to/builders -models-package github.com/yourusername/yourproject/models")
	fmt.Println("\nSecurity Notes:")
	fmt.Println("  - All paths are sanitized to prevent directory traversal attacks")
	fmt.Println("  - Package names are validated against Go naming conventions")
	fmt.Println("  - Input validation is performed on all parameters")
}
