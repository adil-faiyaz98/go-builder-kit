package generator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// Options contains options for the generator
type Options struct {
	PackageName   string
	ModelsPackage string
	Verbose       bool
}

// Generator generates builder code for structs
type Generator struct {
	Options Options
}

// NewGenerator creates a new Generator
func NewGenerator(opts Options) *Generator {
	return &Generator{
		Options: opts,
	}
}

// ProcessFile processes a single Go file and generates builders for all structs
func (g *Generator) ProcessFile(inputFile, outputDir string) error {
	if g.Options.Verbose {
		fmt.Printf("Processing file: %s\n", inputFile)
	}

	// Parse the Go file
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, inputFile, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse file %s: %v", inputFile, err)
	}

	// Extract package name if not provided
	if g.Options.ModelsPackage == "" {
		g.Options.ModelsPackage = node.Name.Name
	}

	// Find all structs in the file
	var structs []*ast.TypeSpec
	for _, decl := range node.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if _, ok := typeSpec.Type.(*ast.StructType); ok {
						structs = append(structs, typeSpec)
					}
				}
			}
		}
	}

	// Generate builders for each struct
	for _, structType := range structs {
		if g.Options.Verbose {
			fmt.Printf("Generating builder for struct: %s\n", structType.Name.Name)
		}

		// Generate builder code
		builderCode := g.generateBuilderCode(structType)

		// Write builder to file
		outputFile := filepath.Join(outputDir, ToSnakeCase(structType.Name.Name)+"_builder.go")
		if err := os.WriteFile(outputFile, []byte(builderCode), 0644); err != nil {
			return fmt.Errorf("failed to write builder file %s: %v", outputFile, err)
		}

		if g.Options.Verbose {
			fmt.Printf("Generated builder file: %s\n", outputFile)
		}
	}

	return nil
}

// ProcessDirectory processes all Go files in a directory and generates builders for all structs
func (g *Generator) ProcessDirectory(inputDir, outputDir string, recursive bool) error {
	if g.Options.Verbose {
		fmt.Printf("Processing directory: %s\n", inputDir)
	}

	// Get all Go files in the directory
	var goFiles []string
	if recursive {
		err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && strings.HasSuffix(path, ".go") {
				goFiles = append(goFiles, path)
			}
			return nil
		})
		if err != nil {
			return fmt.Errorf("failed to walk directory %s: %v", inputDir, err)
		}
	} else {
		files, err := os.ReadDir(inputDir)
		if err != nil {
			return fmt.Errorf("failed to read directory %s: %v", inputDir, err)
		}
		for _, file := range files {
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".go") {
				goFiles = append(goFiles, filepath.Join(inputDir, file.Name()))
			}
		}
	}

	// Process each Go file
	for _, file := range goFiles {
		if err := g.ProcessFile(file, outputDir); err != nil {
			return err
		}
	}

	return nil
}

// generateBuilderCode generates builder code for a struct
func (g *Generator) generateBuilderCode(structType *ast.TypeSpec) string {
	// This is a placeholder - in a real implementation, this would generate the actual builder code
	return fmt.Sprintf(`package %s

import (
	"%s"
)

// %sBuilder is a builder for %s
type %sBuilder struct {
	%s *%s.%s
}

// New%sBuilder creates a new %sBuilder
func New%sBuilder() *%sBuilder {
	return &%sBuilder{
		%s: &%s.%s{},
	}
}

// Build builds the %s
func (b *%sBuilder) Build() %s.%s {
	return *b.%s
}

// BuildPtr builds the %s and returns a pointer
func (b *%sBuilder) BuildPtr() *%s.%s {
	return b.%s
}
`,
		g.Options.PackageName,
		g.Options.ModelsPackage,
		structType.Name.Name, structType.Name.Name,
		structType.Name.Name,
		strings.ToLower(structType.Name.Name), g.Options.ModelsPackage, structType.Name.Name,
		structType.Name.Name, structType.Name.Name,
		structType.Name.Name, structType.Name.Name,
		structType.Name.Name,
		strings.ToLower(structType.Name.Name), g.Options.ModelsPackage, structType.Name.Name,
		structType.Name.Name,
		structType.Name.Name, g.Options.ModelsPackage, structType.Name.Name,
		strings.ToLower(structType.Name.Name),
		structType.Name.Name,
		structType.Name.Name, g.Options.ModelsPackage, structType.Name.Name,
		strings.ToLower(structType.Name.Name),
	)
}

// ToSnakeCase converts a string from CamelCase to snake_case
func ToSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && 'A' <= r && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(unicode.ToLower(r))
	}
	return result.String()
}
