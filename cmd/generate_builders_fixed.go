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
	"unicode"
)

func main() {
	// Create builders directory if it doesn't exist
	if err := os.MkdirAll("builders", os.ModePerm); err != nil {
		log.Fatalf("Failed to create builders directory: %v", err)
	}

	// Parse models directory
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, "models", nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("Failed to parse models directory: %v", err)
	}

	// Find all structs in the models package
	allStructs := make(map[string]*ast.StructType)
	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			ast.Inspect(file, func(n ast.Node) bool {
				if typeSpec, ok := n.(*ast.TypeSpec); ok {
					if structType, ok := typeSpec.Type.(*ast.StructType); ok {
						allStructs[typeSpec.Name.Name] = structType
					}
				}
				return true
			})
		}
	}

	// Generate builders for each struct
	imports := make(map[string]bool)
	for structName, structType := range allStructs {
		generateBuilder(structName, structType, allStructs)
		// Collect imports
		ast.Inspect(structType, func(n ast.Node) bool {
			if ident, ok := n.(*ast.Ident); ok {
				if _, exists := allStructs[ident.Name]; exists {
					imports["github.com/adil-faiyaz98/go-builder-kit/models"] = true
				}
			}
			return true
		})
	}

	// Generate builder utility file
	generateBuilderUtil(imports)

	fmt.Println("Builder generation completed successfully!")
}

func generateBuilder(structName string, structType *ast.StructType, allStructs map[string]*ast.StructType) {
	// Create a data structure to hold template data
	data := struct {
		PackageName   string
		StructName    string
		Fields        []struct{ Name, Type string }
		NestedStructs []struct{ Name, Type string }
	}{
		PackageName: "builders",
		StructName:  structName,
		Fields:      []struct{ Name, Type string }{},
		NestedStructs: []struct {
			Name, Type string
		}{},
	}

	// Extract fields from the struct
	for _, field := range structType.Fields.List {
		if len(field.Names) > 0 {
			fieldName := field.Names[0].Name
			var fieldType string

			// Get the field type as a string
			switch t := field.Type.(type) {
			case *ast.Ident:
				fieldType = t.Name
				// Check if this is a nested struct
				if _, exists := allStructs[fieldType]; exists {
					data.NestedStructs = append(data.NestedStructs, struct{ Name, Type string }{
						Name: fieldName,
						Type: fieldType,
					})
				}
			case *ast.ArrayType:
				if ident, ok := t.Elt.(*ast.Ident); ok {
					fieldType = "[]" + ident.Name
					// Check if this is an array of nested structs
					if _, exists := allStructs[ident.Name]; exists {
						data.NestedStructs = append(data.NestedStructs, struct{ Name, Type string }{
							Name: fieldName,
							Type: ident.Name,
						})
					}
				} else {
					fieldType = "[]interface{}"
				}
			case *ast.MapType:
				if keyIdent, ok := t.Key.(*ast.Ident); ok {
					if valueIdent, ok := t.Value.(*ast.Ident); ok {
						fieldType = "map[" + keyIdent.Name + "]" + valueIdent.Name
					} else {
						fieldType = "map[" + keyIdent.Name + "]interface{}"
					}
				} else {
					fieldType = "map[string]interface{}"
				}
			case *ast.InterfaceType:
				fieldType = "interface{}"
			case *ast.StarExpr:
				if ident, ok := t.X.(*ast.Ident); ok {
					fieldType = "*" + ident.Name
					// Check if this is a pointer to a nested struct
					if _, exists := allStructs[ident.Name]; exists {
						data.NestedStructs = append(data.NestedStructs, struct{ Name, Type string }{
							Name: fieldName,
							Type: ident.Name,
						})
					}
				} else {
					fieldType = "*interface{}"
				}
			default:
				fieldType = "interface{}"
			}

			data.Fields = append(data.Fields, struct{ Name, Type string }{
				Name: fieldName,
				Type: fieldType,
			})
		}
	}

	// Create builder file
	file, err := os.Create(fmt.Sprintf("builders/%s_builder.go", toSnakeCase(structName)))
	if err != nil {
		log.Fatalf("Failed to create builder file for %s: %v", structName, err)
	}
	defer file.Close()

	// Parse and execute the template
	tmpl := template.Must(template.New("builder").Parse(`package {{.PackageName}}

import (
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

// {{.StructName}}Builder is a builder for {{.StructName}}
type {{.StructName}}Builder struct {
	{{.StructName | toLower}} *models.{{.StructName}}
	validationFuncs []func(*models.{{.StructName}}) error
}

// New{{.StructName}}Builder creates a new {{.StructName}}Builder
func New{{.StructName}}Builder() *{{.StructName}}Builder {
	return &{{.StructName}}Builder{
		{{.StructName | toLower}}: &models.{{.StructName}}{},
	}
}

// Build builds the {{.StructName}}
func (b *{{.StructName}}Builder) Build() models.{{.StructName}} {
	return *b.{{.StructName | toLower}}
}

// BuildPtr builds the {{.StructName}} and returns a pointer
func (b *{{.StructName}}Builder) BuildPtr() *models.{{.StructName}} {
	return b.{{.StructName | toLower}}
}

// BuildAndValidate builds the {{.StructName}} and validates it
func (b *{{.StructName}}Builder) BuildAndValidate() (models.{{.StructName}}, error) {
	// Run custom validation functions
	for _, validateFunc := range b.validationFuncs {
		if err := validateFunc(b.{{.StructName | toLower}}); err != nil {
			return models.{{.StructName}}{}, err
		}
	}

	// Run model's Validate method
	if err := b.{{.StructName | toLower}}.Validate(); err != nil {
		return models.{{.StructName}}{}, err
	}

	return *b.{{.StructName | toLower}}, nil
}

// Clone creates a deep copy of the builder
func (b *{{.StructName}}Builder) Clone() *{{.StructName}}Builder {
	clone := New{{.StructName}}Builder()
	*clone.{{.StructName | toLower}} = *b.{{.StructName | toLower}}
	clone.validationFuncs = append([]func(*models.{{.StructName}}) error{}, b.validationFuncs...)
	return clone
}

// WithValidation adds a custom validation function
func (b *{{.StructName}}Builder) WithValidation(validateFunc func(*models.{{.StructName}}) error) *{{.StructName}}Builder {
	b.validationFuncs = append(b.validationFuncs, validateFunc)
	return b
}

{{range .Fields}}
// With{{.Name}} sets the {{.Name}} field
func (b *{{$.StructName}}Builder) With{{.Name}}({{.Name | toLower}} {{.Type}}) *{{$.StructName}}Builder {
	b.{{$.StructName | toLower}}.{{.Name}} = {{.Name | toLower}}
	return b
}
{{end}}

{{range .NestedStructs}}
// With{{.Name}} sets the {{.Name}} field using a {{.Type}}Builder
func (b *{{$.StructName}}Builder) With{{.Name}}(builder *{{.Type}}Builder) *{{$.StructName}}Builder {
	b.{{$.StructName | toLower}}.{{.Name}} = builder.BuildPtr()
	return b
}

// Add{{.Name}} adds a {{.Type}} to the {{.Name}} slice
func (b *{{$.StructName}}Builder) Add{{.Name}}(builder *{{.Type}}Builder) *{{$.StructName}}Builder {
	b.{{$.StructName | toLower}}.{{.Name}} = append(b.{{$.StructName | toLower}}.{{.Name}}, builder.BuildPtr())
	return b
}
{{end}}
`))

	// Add custom template functions
	tmpl.Funcs(template.FuncMap{
		"toLower": func(s string) string {
			if s == "" {
				return ""
			}
			return strings.ToLower(s[:1]) + s[1:]
		},
	})

	if err := tmpl.Execute(file, data); err != nil {
		log.Fatalf("Failed to execute template for %s: %v", structName, err)
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
	utilContent := `// BuilderInterface defines the common interface for all builders
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
}`

	if _, err := file.WriteString(utilContent); err != nil {
		log.Fatalf("Failed to write utility functions: %v", err)
	}

	fmt.Println("Generated builder_util.go")
}

// Helper function to convert CamelCase to snake_case
func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && 'A' <= r && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(unicode.ToLower(r))
	}
	return result.String()
}
