package generator

// BuilderTemplate is the template for generating builder code
const BuilderTemplate = `package {{ .PackageName }}

import (
	"fmt"
	{{ if .ModelsPackage }}
	"{{ .ModelsPackage }}"
	{{ else }}
	"github.com/adil-faiyaz98/go-builder-kit/models"
	{{ end }}
)

// {{ .Struct.Name }}Builder builds a {{ .Struct.Name }} model
type {{ .Struct.Name }}Builder struct {
	{{ ToLowerFirst .Struct.Name }} *models.{{ .Struct.Name }}
	// Custom validation functions
	validationFuncs []func(*models.{{ .Struct.Name }}) error
}

// New{{ .Struct.Name }}Builder creates a new {{ .Struct.Name }}Builder
func New{{ .Struct.Name }}Builder() *{{ .Struct.Name }}Builder {
	return &{{ .Struct.Name }}Builder{
		{{ ToLowerFirst .Struct.Name }}: &models.{{ .Struct.Name }}{
			{{- range .Struct.Fields }}
			{{ .Name }}: {{ if .IsSlice }}[]{{ if .IsNested }}*models.{{ .ElementType }}{{ else }}{{ .Type }}{{ end }}{}{{ else if .IsMap }}map[{{ .KeyType }}]{{ .ValType }}{}{{ else if .IsPointer }}{{ if .IsNested }}nil{{ else }}nil{{ end }}{{ else }}{{ if eq .Type "string" }}""{{ else if eq .Type "int" }}0{{ else if eq .Type "float64" }}0.0{{ else if eq .Type "bool" }}false{{ else if .IsBuiltin }}{{ .Type }}(0){{ else }}models.{{ .Type }}{}{{ end }}{{ end }},
			{{- end }}
		},
		validationFuncs: []func(*models.{{ .Struct.Name }}) error{},
	}
}

// New{{ .Struct.Name }}BuilderWithDefaults creates a new {{ .Struct.Name }}Builder with sensible defaults
func New{{ .Struct.Name }}BuilderWithDefaults() *{{ .Struct.Name }}Builder {
	builder := New{{ .Struct.Name }}Builder()
	// Add default values here if needed
	return builder
}

{{- range .Struct.Fields }}
// With{{ .Name }} sets the {{ .Name }}
func (b *{{ $.Struct.Name }}Builder) With{{ .Name }}({{ if eq (ToLowerFirst .Name) "type" }}value {{ .Type }}{{ else }}{{ if .IsMap }}key {{ .KeyType }}, val {{ .ValType }}{{ else }}{{ ToLowerFirst .Name }} {{ if .IsSlice }}{{ if .IsNested }}*{{ .ElementType }}Builder{{ else }}{{ .Type }}{{ end }}{{ else if .IsPointer }}{{ if .IsNested }}*{{ .Type }}Builder{{ else }}*{{ .Type }}{{ end }}{{ else if .IsNested }}*{{ .Type }}Builder{{ else }}{{ .Type }}{{ end }}{{ end }}{{ end }}) *{{ $.Struct.Name }}Builder {
	{{- if and .IsSlice .IsNested }}
	// Ensure the slice is initialized
	if b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} == nil {
		b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} = []*models.{{ .ElementType }}{}
	}
	{{- end }}
	{{- if .IsSlice }}
	{{- if .IsNested }}
	// Handle nested slice element
	builtValue := {{ if eq (ToLowerFirst .Name) "type" }}value{{ else }}{{ ToLowerFirst .Name }}{{ end }}.Build().(*models.{{ .ElementType }})
	b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} = append(b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }}, builtValue)
	{{- else }}
	b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} = append(b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }}, {{ if eq (ToLowerFirst .Name) "type" }}value{{ else }}{{ ToLowerFirst .Name }}{{ end }})
	{{- end }}
	{{- else if .IsMap }}
	if b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} == nil {
		b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} = make(map[{{ .KeyType }}]{{ .ValType }})
	}
	{{- if .IsNested }}
	b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }}[key] = val.Build().(models.{{ .ValType }})
	{{- else }}
	b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }}[key] = val
	{{- end }}
	{{- else if .IsPointer }}
	{{- if .IsNested }}
	// Handle nested pointer
	b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} = {{ if eq (ToLowerFirst .Name) "type" }}value{{ else }}{{ ToLowerFirst .Name }}{{ end }}.BuildPtr()
	{{- else }}
	b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} = {{ if eq (ToLowerFirst .Name) "type" }}value{{ else }}{{ ToLowerFirst .Name }}{{ end }}
	{{- end }}
	{{- else if .IsNested }}
	builtValue := {{ if eq (ToLowerFirst .Name) "type" }}value{{ else }}{{ ToLowerFirst .Name }}{{ end }}.Build().(*models.{{ .Type }})
	b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} = *builtValue
	{{- else }}
	b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} = {{ if eq (ToLowerFirst .Name) "type" }}value{{ else }}{{ ToLowerFirst .Name }}{{ end }}
	{{- end }}
	return b
}
{{ end }}

// WithValidation adds a custom validation function
func (b *{{ .Struct.Name }}Builder) WithValidation(validationFunc func(*models.{{ .Struct.Name }}) error) *{{ .Struct.Name }}Builder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the {{ .Struct.Name }}
func (b *{{ .Struct.Name }}Builder) Build() interface{} {
	return b.{{ ToLowerFirst .Struct.Name }}
}

// BuildPtr builds the {{ .Struct.Name }} and returns a pointer
func (b *{{ .Struct.Name }}Builder) BuildPtr() *models.{{ .Struct.Name }} {
	return b.{{ ToLowerFirst .Struct.Name }}
}

// BuildAndValidate builds the {{ .Struct.Name }} and validates it
func (b *{{ .Struct.Name }}Builder) BuildAndValidate() (*models.{{ .Struct.Name }}, error) {
	{{ ToLowerFirst .Struct.Name }} := b.{{ ToLowerFirst .Struct.Name }}

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc({{ ToLowerFirst .Struct.Name }}); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}({{ ToLowerFirst .Struct.Name }}).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return {{ ToLowerFirst .Struct.Name }}, err
		}
	}

	return {{ ToLowerFirst .Struct.Name }}, nil
}

// MustBuild builds the {{ .Struct.Name }} and panics if validation fails
func (b *{{ .Struct.Name }}Builder) MustBuild() *models.{{ .Struct.Name }} {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *{{ .Struct.Name }}Builder) Clone() *{{ .Struct.Name }}Builder {
	cloned{{ .Struct.Name }} := *b.{{ ToLowerFirst .Struct.Name }}
	return &{{ .Struct.Name }}Builder{
		{{ ToLowerFirst .Struct.Name }}: &cloned{{ .Struct.Name }},
		validationFuncs: append([]func(*models.{{ .Struct.Name }}) error{}, b.validationFuncs...),
	}
}
`
