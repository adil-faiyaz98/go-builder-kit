package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/adil-faiyaz98/go-builder-kit/pkg/generator"
)

func main() {
	// Create builders directory if it doesn't exist
	if err := os.MkdirAll("builders", os.ModePerm); err != nil {
		log.Fatalf("Failed to create builders directory: %v", err)
	}

	// Parse the models directory
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, "models", nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("Failed to parse models directory: %v", err)
	}

	// Create a map to store all struct definitions
	structs := make(map[string]*ast.StructType)
	imports := make(map[string]bool)

	// Extract all struct definitions from the AST
	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			// Extract imports
			for _, imp := range file.Imports {
				if imp.Path != nil {
					importPath := strings.Trim(imp.Path.Value, "\"")
					imports[importPath] = true
				}
			}

			// Extract struct definitions
			ast.Inspect(file, func(n ast.Node) bool {
				if typeSpec, ok := n.(*ast.TypeSpec); ok {
					if structType, ok := typeSpec.Type.(*ast.StructType); ok {
						structs[typeSpec.Name.Name] = structType
					}
				}
				return true
			})
		}
	}

	// Generate builder files for each struct
	for structName, structType := range structs {
		generateBuilder(structName, structType, structs)
	}

	// Generate builder_util.go file
	generateBuilderUtil(imports)

	fmt.Println("Builder generation completed successfully!")
}

func generateBuilder(structName string, structType *ast.StructType, allStructs map[string]*ast.StructType) {
	// Create a data structure to hold template data
	data := struct {
		PackageName   string
		ModelsPackage string
		Struct        struct {
			Name   string
			Fields []struct {
				Name        string
				Type        string
				KeyType     string
				ValType     string
				ElementType string
				IsSlice     bool
				IsMap       bool
				IsPointer   bool
				IsNested    bool
				IsBuiltin   bool
			}
		}
	}{
		PackageName:   "builders",
		ModelsPackage: "",
	}

	data.Struct.Name = structName
	data.Struct.Fields = make([]struct {
		Name        string
		Type        string
		KeyType     string
		ValType     string
		ElementType string
		IsSlice     bool
		IsMap       bool
		IsPointer   bool
		IsNested    bool
		IsBuiltin   bool
	}, 0)

	// Process struct fields
	for _, field := range structType.Fields.List {
		if len(field.Names) == 0 {
			// Skip embedded fields
			continue
		}

		fieldName := field.Names[0].Name

		// Skip unexported fields
		if !ast.IsExported(fieldName) {
			continue
		}

		fieldType := ""
		keyType := ""
		valType := ""
		elementType := ""
		isSlice := false
		isMap := false
		isPointer := false
		isNested := false
		isBuiltin := false

		// Determine field type
		switch t := field.Type.(type) {
		case *ast.Ident:
			fieldType = t.Name
			// Check if it's a built-in type
			switch fieldType {
			case "string", "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "float32", "float64", "bool", "byte", "rune":
				isBuiltin = true
			default:
				// Check if it's a nested struct
				if _, ok := allStructs[fieldType]; ok {
					isNested = true
				}
			}
		case *ast.ArrayType:
			isSlice = true
			switch et := t.Elt.(type) {
			case *ast.Ident:
				fieldType = "[]" + et.Name
				elementType = et.Name
				// Check if element is a nested struct
				if _, ok := allStructs[elementType]; ok {
					isNested = true
				}
			case *ast.StarExpr:
				if ident, ok := et.X.(*ast.Ident); ok {
					fieldType = "[]*" + ident.Name
					elementType = ident.Name
					// Check if element is a nested struct
					if _, ok := allStructs[elementType]; ok {
						isNested = true
					}
				}
			}
		case *ast.MapType:
			isMap = true
			keyIdent, keyOk := t.Key.(*ast.Ident)
			valIdent, valOk := t.Value.(*ast.Ident)

			if keyOk && valOk {
				keyType = keyIdent.Name
				valType = valIdent.Name
				fieldType = fmt.Sprintf("map[%s]%s", keyType, valType)

				// Check if value is a nested struct
				if _, ok := allStructs[valType]; ok {
					isNested = true
				}
			} else if keyOk && !valOk {
				keyType = keyIdent.Name

				// Handle pointer value types
				if starExpr, ok := t.Value.(*ast.StarExpr); ok {
					if ident, ok := starExpr.X.(*ast.Ident); ok {
						valType = "*" + ident.Name
						fieldType = fmt.Sprintf("map[%s]%s", keyType, valType)

						// Check if value is a nested struct
						if _, ok := allStructs[ident.Name]; ok {
							isNested = true
						}
					}
				}
			}
		case *ast.StarExpr:
			isPointer = true
			if ident, ok := t.X.(*ast.Ident); ok {
				fieldType = "*" + ident.Name
				elementType = ident.Name

				// Check if it's a nested struct
				if _, ok := allStructs[elementType]; ok {
					isNested = true
				}
			}
		case *ast.SelectorExpr:
			if ident, ok := t.X.(*ast.Ident); ok {
				fieldType = ident.Name + "." + t.Sel.Name
			}
		case *ast.InterfaceType:
			fieldType = "interface{}"
		}

		// Add field to struct data
		data.Struct.Fields = append(data.Struct.Fields, struct {
			Name        string
			Type        string
			KeyType     string
			ValType     string
			ElementType string
			IsSlice     bool
			IsMap       bool
			IsPointer   bool
			IsNested    bool
			IsBuiltin   bool
		}{
			Name:        fieldName,
			Type:        fieldType,
			KeyType:     keyType,
			ValType:     valType,
			ElementType: elementType,
			IsSlice:     isSlice,
			IsMap:       isMap,
			IsPointer:   isPointer,
			IsNested:    isNested,
			IsBuiltin:   isBuiltin,
		})
	}

	// Create template functions
	funcMap := template.FuncMap{
		"ToLowerFirst": func(s string) string {
			if s == "" {
				return ""
			}
			return strings.ToLower(s[:1]) + s[1:]
		},
	}

	// Parse the template
	tmpl, err := template.New("builder").Funcs(funcMap).Parse(generator.BuilderTemplate)
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	// Create the builder file
	builderFileName := fmt.Sprintf("builders/%s_builder.go", toSnakeCase(structName))
	file, err := os.Create(builderFileName)
	if err != nil {
		log.Fatalf("Failed to create builder file: %v", err)
	}
	defer file.Close()

	// Execute the template
	if err := tmpl.Execute(file, data); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	fmt.Printf("Generated builder for %s\n", structName)
}

func generateBuilderUtil(imports map[string]bool) {
	// Create builder_util.go file
	file, err := os.Create("builders/builder_util.go")
	if err != nil {
		log.Fatalf("Failed to create builder_util.go: %v", err)
	}
	defer file.Close()

	// Write package declaration
	if _, err := file.WriteString("package builders\n\n"); err != nil {
		log.Fatalf("Failed to write package declaration: %v", err)
	}

	// Write imports
	if _, err := file.WriteString("import (\n"); err != nil {
		log.Fatalf("Failed to write import opening: %v", err)
	}
	if _, err := file.WriteString("\t\"github.com/adil-faiyaz98/go-builder-kit/models\"\n"); err != nil {
		log.Fatalf("Failed to write models import: %v", err)
	}
	for imp := range imports {
		if imp != "" && !strings.Contains(imp, "github.com/adil-faiyaz98/go-builder-kit") {
			if _, err := file.WriteString(fmt.Sprintf("\t\"%s\"\n", imp)); err != nil {
				log.Fatalf("Failed to write import %s: %v", imp, err)
			}
		}
	}
	if _, err := file.WriteString(")\n\n"); err != nil {
		log.Fatalf("Failed to write import closing: %v", err)
	}

	// Write utility functions
	if _, err := file.WriteString(`// BuilderInterface defines the common interface for all builders`); err != nil {
		log.Fatalf("Failed to write utility functions: %v", err)
	}
type BuilderInterface interface {
	Build() interface{}
	BuildAndValidate() (interface{}, error)
}

// ValidateFunc is a function type for custom validation
type ValidateFunc func(interface{}) error

// BuilderOption is a function type for builder options
type BuilderOption func(BuilderInterface)

// WithValidation adds a validation function to a builder
func WithValidation(validateFunc ValidateFunc) BuilderOption {
	return func(builder BuilderInterface) {
		// Implementation depends on the specific builder
	}
}
`)

	fmt.Println("Generated builder_util.go")
}

// Helper function to convert CamelCase to snake_case
func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && 'A' <= r && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}
