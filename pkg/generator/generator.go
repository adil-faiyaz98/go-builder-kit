package generator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
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

// StructField represents a field in a struct
type StructField struct {
	Name         string
	Type         string
	IsPointer    bool
	IsSlice      bool
	IsMap        bool
	IsNested     bool
	IsBuiltin    bool
	ElementType  string // For slices
	KeyType      string // For maps
	ValType      string // For maps
	ImportNeeded string // Import path needed for this field
	BuilderName  string // Name of the builder for nested structs
}

// StructInfo represents information about a struct
type StructInfo struct {
	Name   string
	Fields []StructField
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

		// Extract struct information
		structInfo, err := g.extractStructInfo(structType)
		if err != nil {
			return fmt.Errorf("failed to extract struct info for %s: %v", structType.Name.Name, err)
		}

		// Generate builder code
		builderCode, err := g.generateBuilderCode(structInfo)
		if err != nil {
			return fmt.Errorf("failed to generate builder code for %s: %v", structType.Name.Name, err)
		}

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

// extractStructInfo extracts information about a struct
func (g *Generator) extractStructInfo(typeSpec *ast.TypeSpec) (StructInfo, error) {
	structType, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		return StructInfo{}, fmt.Errorf("not a struct type: %s", typeSpec.Name.Name)
	}

	structInfo := StructInfo{
		Name:   typeSpec.Name.Name,
		Fields: []StructField{},
	}

	// Extract fields
	for _, field := range structType.Fields.List {
		if len(field.Names) == 0 {
			// Skip embedded fields
			continue
		}

		fieldName := field.Names[0].Name
		fieldType, isPointer, isSlice, isMap, isNested, isBuiltin, elementType, keyType, valType, importNeeded, builderName := g.extractFieldType(field.Type)

		structInfo.Fields = append(structInfo.Fields, StructField{
			Name:         fieldName,
			Type:         fieldType,
			IsPointer:    isPointer,
			IsSlice:      isSlice,
			IsMap:        isMap,
			IsNested:     isNested,
			IsBuiltin:    isBuiltin,
			ElementType:  elementType,
			KeyType:      keyType,
			ValType:      valType,
			ImportNeeded: importNeeded,
			BuilderName:  builderName,
		})
	}

	return structInfo, nil
}

// extractFieldType extracts type information from a field
func (g *Generator) extractFieldType(expr ast.Expr) (string, bool, bool, bool, bool, bool, string, string, string, string, string) {
	// fieldType, isPointer, isSlice, isMap, isNested, isBuiltin, elementType, keyType, valType, importNeeded, builderName
	switch t := expr.(type) {
	case *ast.Ident:
		// Basic type (e.g., string, int)
		isBuiltin := isBuiltinType(t.Name)
		isNested := !isBuiltin && t.Name != "any" && t.Name != "interface{}"
		builderName := ""
		if isNested {
			builderName = t.Name + "Builder"
		}
		return t.Name, false, false, false, isNested, isBuiltin, "", "", "", "", builderName

	case *ast.StarExpr:
		// Pointer type (e.g., *Person)
		fieldType, _, _, _, isNested, isBuiltin, _, _, _, importNeeded, builderName := g.extractFieldType(t.X)
		return fieldType, true, false, false, isNested, isBuiltin, "", "", "", importNeeded, builderName

	case *ast.ArrayType:
		// Slice type (e.g., []string, []*Person)
		elementType, isPointer, _, _, isNested, isBuiltin, _, _, _, importNeeded, builderName := g.extractFieldType(t.Elt)
		sliceType := ""
		if isPointer {
			sliceType = fmt.Sprintf("[]*%s", elementType)
		} else {
			sliceType = fmt.Sprintf("[]%s", elementType)
		}
		return sliceType, isPointer, true, false, isNested, isBuiltin, elementType, "", "", importNeeded, builderName

	case *ast.MapType:
		// Map type (e.g., map[string]int)
		keyType, _, _, _, _, _, _, _, _, _, _ := g.extractFieldType(t.Key)
		valType, _, _, _, isNested, _, _, _, _, importNeeded, builderName := g.extractFieldType(t.Value)
		return fmt.Sprintf("map[%s]%s", keyType, valType), false, false, true, isNested, false, "", keyType, valType, importNeeded, builderName

	case *ast.SelectorExpr:
		// Qualified type (e.g., time.Time)
		if ident, ok := t.X.(*ast.Ident); ok {
			builderName := ""
			// Only consider it a builder if it's from our models package
			if ident.Name == "models" || ident.Name == g.Options.ModelsPackage || strings.HasSuffix(ident.Name, "/models") {
				builderName = t.Sel.Name + "Builder"
			}
			return t.Sel.Name, false, false, false, true, false, "", "", "", ident.Name, builderName
		}
		return "interface{}", false, false, false, false, false, "", "", "", "", ""

	case *ast.InterfaceType:
		// Interface type
		return "interface{}", false, false, false, false, true, "", "", "", "", ""

	default:
		// Unknown type
		return "interface{}", false, false, false, false, true, "", "", "", "", ""
	}
}

// generateBuilderCode generates builder code for a struct
func (g *Generator) generateBuilderCode(structInfo StructInfo) (string, error) {
	// Create template data
	data := struct {
		PackageName   string
		ModelsPackage string
		Struct        StructInfo
		ImportLines   string
	}{
		PackageName:   g.Options.PackageName,
		ModelsPackage: g.Options.ModelsPackage,
		Struct:        structInfo,
		ImportLines:   "",
	}

	// Add required imports
	imports := map[string]bool{
		"fmt":                   true, // Always needed for error formatting
		g.Options.ModelsPackage: true, // Always needed for model types
	}

	// We don't actually need reflect for most builders
	// Only add it if we're using it for specific functionality
	// imports["reflect"] = true

	// Add imports for field types
	for _, field := range structInfo.Fields {
		if field.ImportNeeded != "" && field.ImportNeeded != "models" {
			imports[field.ImportNeeded] = true
		}
	}

	// We don't need to automatically add the time package
	// Only add it if we're actually using time.Time types
	// imports["time"] = true

	// Convert map to sorted slice for consistent output
	var importList []string
	for imp := range imports {
		importList = append(importList, imp)
	}
	sort.Strings(importList)

	// Generate import lines
	for _, imp := range importList {
		data.ImportLines += fmt.Sprintf("\t\"%s\"\n", imp)
	}

	// Parse template
	tmpl, err := template.New("builder").Funcs(template.FuncMap{
		"ToLowerFirst": ToLowerFirst,
		"ToParamName":  ToParamName,
		"replace":      strings.ReplaceAll,
		"Singular":     ToSingular,
		"base": func(path string) string {
			// Extract the base package name from a full import path
			parts := strings.Split(path, "/")
			return parts[len(parts)-1]
		},
	}).Parse(BuilderTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %v", err)
	}

	// Execute template
	var buf strings.Builder
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %v", err)
	}

	return buf.String(), nil
}

// isBuiltinType checks if a type is a built-in Go type
func isBuiltinType(typeName string) bool {
	builtinTypes := map[string]bool{
		"string":     true,
		"bool":       true,
		"int":        true,
		"int8":       true,
		"int16":      true,
		"int32":      true,
		"int64":      true,
		"uint":       true,
		"uint8":      true,
		"uint16":     true,
		"uint32":     true,
		"uint64":     true,
		"uintptr":    true,
		"byte":       true,
		"rune":       true,
		"float32":    true,
		"float64":    true,
		"complex64":  true,
		"complex128": true,
		"error":      true,
	}
	return builtinTypes[typeName]
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

// ToLowerFirst converts the first character of a string to lowercase
func ToLowerFirst(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

// ToParamName converts a field name to a parameter name
// Examples: ID -> id, UserID -> userID, HTTPRequest -> httpRequest
func ToParamName(s string) string {
	if s == "" {
		return ""
	}

	// Special case for "ID" exactly
	if s == "ID" {
		return "id"
	}

	// Handle common abbreviations that should be all lowercase when they appear as suffixes
	commonSuffixes := map[string]string{
		"ID":    "id",
		"URL":   "url",
		"URI":   "uri",
		"API":   "api",
		"HTTP":  "http",
		"HTTPS": "https",
		"HTML":  "html",
		"XML":   "xml",
		"JSON":  "json",
		"SQL":   "sql",
		"CSS":   "css",
		"JS":    "js",
	}

	// Check for common suffixes
	for suffix, replacement := range commonSuffixes {
		if strings.HasSuffix(s, suffix) && len(s) > len(suffix) {
			// Convert the prefix to camelCase and append the lowercase suffix
			prefix := s[:len(s)-len(suffix)]
			return ToLowerFirst(prefix) + replacement
		}
	}

	// Default case: just convert first character to lowercase
	return ToLowerFirst(s)
}

// ToSingular converts a plural field name to singular
// Examples: Friends -> Friend, Addresses -> Address, Children -> Child
func ToSingular(s string) string {
	if s == "" {
		return ""
	}

	// Common irregular plurals
	irregularPlurals := map[string]string{
		"Children":  "Child",
		"People":    "Person",
		"Men":       "Man",
		"Women":     "Woman",
		"Feet":      "Foot",
		"Teeth":     "Tooth",
		"Geese":     "Goose",
		"Mice":      "Mouse",
		"Criteria":  "Criterion",
		"Phenomena": "Phenomenon",
	}

	// Check for irregular plurals
	if singular, ok := irregularPlurals[s]; ok {
		return singular
	}

	// Handle common plural endings
	// -ies -> -y
	if strings.HasSuffix(s, "ies") && len(s) > 3 {
		return s[:len(s)-3] + "y"
	}

	// -es -> -e or remove -es
	if strings.HasSuffix(s, "es") && len(s) > 2 {
		// Special cases like Addresses -> Address
		if strings.HasSuffix(s[:len(s)-2], "ss") {
			return s[:len(s)-2]
		}
		// Cases like Boxes -> Box
		if len(s) > 3 && (s[len(s)-3] == 'x' || s[len(s)-3] == 's' || s[len(s)-3] == 'z' ||
			(s[len(s)-3] == 'h' && (s[len(s)-4] == 'c' || s[len(s)-4] == 's'))) {
			return s[:len(s)-2]
		}
	}

	// Default: remove trailing 's'
	if strings.HasSuffix(s, "s") && len(s) > 1 {
		return s[:len(s)-1]
	}

	// If no rule applies, return as is
	return s
}
