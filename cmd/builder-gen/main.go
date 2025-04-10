package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/adil-faiyaz98/go-builder-kit/pkg/generator"
)

func main() {
	// Define command-line flags
	inputFile := flag.String("input", "", "Input Go file or directory containing structs to generate builders for")
	outputDir := flag.String("output", "", "Output directory for generated builder files")
	packageName := flag.String("package", "builders", "Package name for generated builder files")
	modelsPackage := flag.String("models-package", "", "Package path for the models (e.g., github.com/user/repo/models)")
	recursive := flag.Bool("recursive", false, "Recursively process all Go files in the input directory")
	verbose := flag.Bool("verbose", false, "Enable verbose output")

	// Parse command-line flags
	flag.Parse()

	// Validate input
	if *inputFile == "" {
		fmt.Println("Error: input file or directory is required")
		flag.Usage()
		os.Exit(1)
	}

	if *outputDir == "" {
		fmt.Println("Error: output directory is required")
		flag.Usage()
		os.Exit(1)
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Create generator options
	opts := generator.Options{
		PackageName:   *packageName,
		ModelsPackage: *modelsPackage,
		Verbose:       *verbose,
	}

	// Process input
	fileInfo, err := os.Stat(*inputFile)
	if err != nil {
		fmt.Printf("Error accessing input: %v\n", err)
		os.Exit(1)
	}

	if fileInfo.IsDir() {
		// Process directory
		if *recursive {
			err = processDirectory(*inputFile, *outputDir, opts)
		} else {
			err = processFiles([]string{*inputFile}, *outputDir, opts)
		}
	} else {
		// Process single file
		err = processFiles([]string{*inputFile}, *outputDir, opts)
	}

	if err != nil {
		fmt.Printf("Error generating builders: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Builder generation completed successfully!")
}

func processDirectory(inputDir, outputDir string, opts generator.Options) error {
	var files []string

	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".go" {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("error walking directory: %w", err)
	}

	return processFiles(files, outputDir, opts)
}

func processFiles(files []string, outputDir string, opts generator.Options) error {
	gen := generator.NewGenerator(opts)

	// Write utility file first
	utilFile := filepath.Join(outputDir, "builder_util.go")
	utilCode := fmt.Sprintf("package %s\n\n%s", opts.PackageName, generator.UtilTemplate)
	if err := os.WriteFile(utilFile, []byte(utilCode), 0644); err != nil {
		return fmt.Errorf("error writing utility file %s: %w", utilFile, err)
	}

	if opts.Verbose {
		fmt.Printf("Generated utility file: %s\n", utilFile)
	}

	// Write builder registry file
	registryFile := filepath.Join(outputDir, "builder_registry.go")
	registryCode := fmt.Sprintf("package %s\n\n%s", opts.PackageName, generator.RegistryTemplate)
	if err := os.WriteFile(registryFile, []byte(registryCode), 0644); err != nil {
		return fmt.Errorf("error writing registry file %s: %w", registryFile, err)
	}

	if opts.Verbose {
		fmt.Printf("Generated registry file: %s\n", registryFile)
	}

	for _, file := range files {
		if opts.Verbose {
			fmt.Printf("Processing file: %s\n", file)
		}

		// Process the file
		err := gen.ProcessFile(file, outputDir)
		if err != nil {
			return fmt.Errorf("error processing file %s: %w", file, err)
		}
	}

	return nil
}
