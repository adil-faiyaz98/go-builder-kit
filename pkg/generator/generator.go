package generator

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"text/template"
)

// Options contains configuration options for the generator
type Options struct {
	PackageName   string
	ModelsPackage string
	Verbose       bool
}

// Generator generates builder code for structs
type Generator struct {
	Options Options
}

// StructInfo contains information about a struct
type StructInfo struct {
	Name    string
	Comment string
	Fields  []FieldInfo
}

// FieldInfo contains information about a struct field
type FieldInfo struct {
	Name        string
	Type        string
	IsSlice     bool
	IsMap       bool
	IsNested    bool
	IsPointer   bool
	IsBuiltin   bool
	KeyType     string
	ValType     string
	ElementType string // For slices, the element type
	Comment     string
	Tags        map[string]string // Field tags
}

// NewGenerator creates a new Generator
func NewGenerator(opts Options) *Generator {
	return &Generator{
		Options: opts,
	}
}

// ParseFile parses a Go file and extracts struct information
func (g *Generator) ParseFile(filename string) ([]StructInfo, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("error parsing file: %w", err)
	}

	var structs []StructInfo

	ast.Inspect(node, func(n ast.Node) bool {
		// Check if this is a type declaration
		typeSpec, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		// Check if this is a struct type
		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			return true
		}

		// Get struct name
		structName := typeSpec.Name.Name

		// Get struct comment
		var comment string
		if typeSpec.Doc != nil {
			comment = typeSpec.Doc.Text()
		}

		// Create struct info
		structInfo := StructInfo{
			Name:    structName,
			Comment: comment,
			Fields:  []FieldInfo{},
		}

		// Process fields
		for _, field := range structType.Fields.List {
			// Skip anonymous fields
			if len(field.Names) == 0 {
				continue
			}

			// Get field name
			fieldName := field.Names[0].Name

			// Get field type
			fieldType := ""
			isSlice := false
			isMap := false
			isNested := false
			isPointer := false
			keyType := ""
			valType := ""
			elementType := ""

			// Process field type
			switch t := field.Type.(type) {
			case *ast.Ident:
				fieldType = t.Name
				isBuiltin := isBuiltinType(fieldType)
				if ast.IsExported(fieldType) && !isBuiltin {
					isNested = true
				}
			case *ast.SelectorExpr:
				if x, ok := t.X.(*ast.Ident); ok {
					fieldType = x.Name + "." + t.Sel.Name
					// Assume imported types are nested
					isNested = true
				}
			case *ast.StarExpr:
				isPointer = true
				switch pt := t.X.(type) {
				case *ast.Ident:
					fieldType = pt.Name
					isBuiltin := isBuiltinType(fieldType)
					if ast.IsExported(pt.Name) && !isBuiltin {
						isNested = true
					}
				case *ast.SelectorExpr:
					if x, ok := pt.X.(*ast.Ident); ok {
						fieldType = x.Name + "." + pt.Sel.Name
						// Assume imported types are nested
						isNested = true
					}
				}
			case *ast.ArrayType:
				isSlice = true
				switch et := t.Elt.(type) {
				case *ast.Ident:
					fieldType = et.Name
					elementType = et.Name
					isBuiltin := isBuiltinType(fieldType)
					if ast.IsExported(et.Name) && !isBuiltin {
						isNested = true
					}
				case *ast.StarExpr:
					switch pt := et.X.(type) {
					case *ast.Ident:
						fieldType = "*" + pt.Name
						elementType = pt.Name
						isBuiltin := isBuiltinType(elementType)
						if ast.IsExported(pt.Name) && !isBuiltin {
							isNested = true
						}
					case *ast.SelectorExpr:
						if x, ok := pt.X.(*ast.Ident); ok {
							fieldType = "*" + x.Name + "." + pt.Sel.Name
							elementType = pt.Sel.Name
							// Assume imported types are nested
							isNested = true
						}
					}
				case *ast.SelectorExpr:
					if x, ok := et.X.(*ast.Ident); ok {
						fieldType = x.Name + "." + et.Sel.Name
						elementType = et.Sel.Name
						// Assume imported types are nested
						isNested = true
					}
				}
			case *ast.MapType:
				// Handle maps
				isMap = true
				switch kt := t.Key.(type) {
				case *ast.Ident:
					keyType = kt.Name
				case *ast.SelectorExpr:
					if x, ok := kt.X.(*ast.Ident); ok {
						keyType = x.Name + "." + kt.Sel.Name
					}
				}
				switch vt := t.Value.(type) {
				case *ast.Ident:
					valType = vt.Name
					isBuiltin := isBuiltinType(valType)
					if ast.IsExported(valType) && !isBuiltin {
						isNested = true
					}
				case *ast.SelectorExpr:
					if x, ok := vt.X.(*ast.Ident); ok {
						valType = x.Name + "." + vt.Sel.Name
						// Assume imported types are nested
						isNested = true
					}
				}
			}

			// Parse field tags
			tags := make(map[string]string)
			if field.Tag != nil {
				tag := field.Tag.Value
				// Remove the backticks
				tag = tag[1 : len(tag)-1]
				// Parse the tag
				for len(tag) > 0 {
					// Skip leading space
					i := 0
					for i < len(tag) && tag[i] == ' ' {
						i++
					}
					if i == len(tag) {
						break
					}
					tag = tag[i:]
					// Scan to colon
					i = 0
					for i < len(tag) && tag[i] != ':' && tag[i] != ' ' {
						i++
					}
					if i == 0 || i+1 >= len(tag) || tag[i] != ':' {
						break
					}
					name := tag[:i]
					tag = tag[i+1:]
					// Scan quoted string
					i = 1
					for i < len(tag) && tag[i] != '"' {
						if tag[i] == '\\' {
							i++
						}
						i++
					}
					if i >= len(tag) {
						break
					}
					value := tag[1:i]
					tag = tag[i+1:]
					tags[name] = value
				}
			}

			// Extract field comment
			comment := ""
			if field.Comment != nil {
				comment = field.Comment.Text()
			}

			// Add field to struct
			structInfo.Fields = append(structInfo.Fields, FieldInfo{
				Name:        fieldName,
				Type:        fieldType,
				IsSlice:     isSlice,
				IsMap:       isMap,
				IsNested:    isNested,
				IsPointer:   isPointer,
				IsBuiltin:   isBuiltinType(fieldType),
				KeyType:     keyType,
				ValType:     valType,
				ElementType: elementType,
				Comment:     comment,
				Tags:        tags,
			})
		}

		// Add struct to list
		structs = append(structs, structInfo)
		return true
	})

	return structs, nil
}

// GenerateBuilder generates builder code for a struct
func (g *Generator) GenerateBuilder(structInfo StructInfo) (string, error) {
	// Prepare template data
	data := struct {
		PackageName   string
		ModelsPackage string
		Struct        StructInfo
	}{
		PackageName:   g.Options.PackageName,
		ModelsPackage: g.Options.ModelsPackage,
		Struct:        structInfo,
	}

	// Parse template
	tmpl, err := template.New("builder").Funcs(template.FuncMap{
		"ToLower":      strings.ToLower,
		"ToLowerFirst": ToLowerFirst,
		"ToSnakeCase":  ToSnakeCase,
	}).Parse(BuilderTemplate)
	if err != nil {
		return "", fmt.Errorf("error parsing template: %w", err)
	}

	// Execute template
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("error executing template: %w", err)
	}

	return buf.String(), nil
}

// ToLowerFirst returns the string with the first character lowercased
func ToLowerFirst(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = []rune(strings.ToLower(string(r[0])))[0]
	return string(r)
}

// ToSnakeCase converts a string to snake_case
func ToSnakeCase(s string) string {
	var result string
	for i, r := range s {
		if i > 0 && 'A' <= r && r <= 'Z' {
			result += "_"
		}
		result += strings.ToLower(string(r))
	}
	return result
}

// isBuiltinType checks if a type is a built-in Go type
func isBuiltinType(typeName string) bool {
	builtinTypes := map[string]bool{
		"bool":       true,
		"byte":       true,
		"complex64":  true,
		"complex128": true,
		"error":      true,
		"float32":    true,
		"float64":    true,
		"int":        true,
		"int8":       true,
		"int16":      true,
		"int32":      true,
		"int64":      true,
		"rune":       true,
		"string":     true,
		"uint":       true,
		"uint8":      true,
		"uint16":     true,
		"uint32":     true,
		"uint64":     true,
		"uintptr":    true,
		"any":        true,
		"interface{}": true,
	}
	return builtinTypes[typeName]
}
