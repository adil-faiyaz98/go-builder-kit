package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

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
	// Check if input path is a directory
	fileInfo, err := os.Stat(inputPath)
	if err != nil {
		return fmt.Errorf("failed to stat input path: %w", err)
	}

	if fileInfo.IsDir() {
		// Process all Go files in the directory
		files, err := os.ReadDir(inputPath)
		if err != nil {
			return fmt.Errorf("failed to read directory: %w", err)
		}

		for _, file := range files {
			if filepath.Ext(file.Name()) == ".go" && !file.IsDir() {
				filePath := filepath.Join(inputPath, file.Name())
				if err := gen.GenerateBuilders(filePath, outputPath); err != nil {
					return fmt.Errorf("failed to generate builders for %s: %w", filePath, err)
				}
			}
		}
	} else {
		// Process single file
		if filepath.Ext(inputPath) != ".go" {
			return fmt.Errorf("input file must be a Go file")
		}

		if err := gen.GenerateBuilders(inputPath, outputPath); err != nil {
			return fmt.Errorf("failed to generate builders: %w", err)
		}
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
}
