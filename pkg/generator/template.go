package generator

// BuilderTemplate is the template for generating builder code
const BuilderTemplate = `package {{ .PackageName }}

import (
{{ .ImportLines }}	"github.com/adil-faiyaz98/go-builder-kit/v2/pkg/builder"
)

// {{ .Struct.Name }}Builder builds a {{ .Struct.Name }} model
type {{ .Struct.Name }}Builder struct {
	{{ ToLowerFirst .Struct.Name }} *{{ .ModelsPackage | base }}.{{ .Struct.Name }}
	// Custom validation functions
	validationFuncs []func(*{{ .ModelsPackage | base }}.{{ .Struct.Name }}) error
}

// New{{ .Struct.Name }}Builder creates a new {{ .Struct.Name }}Builder
func New{{ .Struct.Name }}Builder() *{{ .Struct.Name }}Builder {
	return &{{ .Struct.Name }}Builder{
		{{ ToLowerFirst .Struct.Name }}: &{{ .ModelsPackage | base }}.{{ .Struct.Name }}{
			{{- range .Struct.Fields }}
			{{ .Name }}: {{ if .IsSlice }}{{ if .IsNested }}[]*{{ $.ModelsPackage | base }}.{{ .ElementType }}{}{{ else }}{{ .Type }}{}{{ end }}{{ else if .IsMap }}map[{{ .KeyType }}]{{ .ValType }}{}{{ else if .IsPointer }}{{ if .IsNested }}nil{{ else }}nil{{ end }}{{ else }}{{ if eq .Type "string" }}""{{ else if eq .Type "int" }}0{{ else if eq .Type "float64" }}0.0{{ else if eq .Type "bool" }}false{{ else if eq .Type "Time" }}time.Time{}{{ else if .IsBuiltin }}{{ .Type }}(0){{ else }}{{ $.ModelsPackage | base }}.{{ .Type }}{}{{ end }}{{ end }},
			{{- end }}
		},
		validationFuncs: []func(*{{ .ModelsPackage | base }}.{{ .Struct.Name }}) error{},
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
func (b *{{ $.Struct.Name }}Builder) With{{ .Name }}({{ if eq (ToLowerFirst .Name) "type" }}value {{ .Type }}{{ else }}{{ if .IsMap }}key {{ .KeyType }}, val {{ .ValType }}{{ else }}{{ ToParamName .Name }} {{ if .IsSlice }}{{ if .IsNested }}[]*{{ .BuilderName }}{{ else }}{{ .Type }}{{ end }}{{ else if .IsPointer }}{{ if .IsNested }}*{{ .BuilderName }}{{ else }}*{{ .Type }}{{ end }}{{ else if eq .Type "Time" }}time.Time{{ else if .IsNested }}*{{ .BuilderName }}{{ else }}{{ .Type }}{{ end }}{{ end }}{{ end }}) *{{ $.Struct.Name }}Builder {
	if b == nil {
		return b
	}
	{{- if and .IsSlice .IsNested }}
	// Ensure the slice is initialized
	if b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} == nil {
		b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} = []*{{ $.ModelsPackage | base }}.{{ .ElementType }}{}
	}
	{{- end }}
	{{- if .IsSlice }}
	{{- if .IsNested }}
	// Handle nested slice elements
	// Initialize the slice
	b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} = make([]*{{ $.ModelsPackage | base }}.{{ .ElementType }}, 0, len({{ if eq (ToLowerFirst .Name) "type" }}value{{ else }}{{ ToParamName .Name }}{{ end }}))
	// Convert each builder to its model
	for _, builder := range {{ if eq (ToLowerFirst .Name) "type" }}value{{ else }}{{ ToParamName .Name }}{{ end }} {
		builtValue := builder.Build().(*{{ $.ModelsPackage | base }}.{{ .ElementType }})
		b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} = append(b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }}, builtValue)
	}
	{{- else }}
	{{- if eq .Type "[]string" }}
	b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} = append(b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }}, {{ if eq (ToLowerFirst .Name) "type" }}value{{ else }}{{ ToParamName .Name }}{{ end }}...)
	{{- else }}
	b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} = append(b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }}, {{ if eq (ToLowerFirst .Name) "type" }}value{{ else }}{{ ToParamName .Name }}{{ end }}...)
	{{- end }}
	{{- end }}
	{{- else if .IsMap }}
	if b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} == nil {
		b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} = make(map[{{ .KeyType }}]{{ .ValType }})
	}
	{{- if .IsNested }}
	b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }}[key] = val.Build().({{ $.ModelsPackage | base }}.{{ .ValType }})
	{{- else }}
	b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }}[key] = val
	{{- end }}
	{{- else if .IsPointer }}
	{{- if .IsNested }}
	// Handle nested pointer
	b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} = {{ if eq (ToLowerFirst .Name) "type" }}value{{ else }}{{ ToParamName .Name }}{{ end }}.BuildPtr()
	{{- else }}
	b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} = {{ if eq (ToLowerFirst .Name) "type" }}value{{ else }}{{ ToParamName .Name }}{{ end }}
	{{- end }}
	{{- else if .IsNested }}
	builtValue := {{ if eq (ToLowerFirst .Name) "type" }}value{{ else }}{{ ToParamName .Name }}{{ end }}.Build().(*{{ $.ModelsPackage | base }}.{{ .Type }})
	b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} = *builtValue
	{{- else }}
	{{- if eq .Type "string" }}
	b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} = builder.SanitizeString({{ if eq (ToLowerFirst .Name) "type" }}value{{ else }}{{ ToParamName .Name }}{{ end }})
	{{- else }}
	b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} = {{ if eq (ToLowerFirst .Name) "type" }}value{{ else }}{{ ToParamName .Name }}{{ end }}
	{{- end }}
	{{- end }}
	return b
}
{{ end }}

{{- range .Struct.Fields }}
{{- if and .IsSlice .IsNested }}
// Add{{ .Name | Singular }} adds a single item to the {{ .Name }} slice
func (b *{{ $.Struct.Name }}Builder) Add{{ .Name | Singular }}({{ ToParamName (.Name | Singular) }} *{{ .BuilderName }}) *{{ $.Struct.Name }}Builder {
	if b == nil {
		return b
	}
	// Ensure the slice is initialized with capacity
	if b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} == nil {
		b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} = make([]*{{ $.ModelsPackage | base }}.{{ .ElementType }}, 0, 4) // Pre-allocate capacity
	}
	// Handle nested slice element
	if {{ ToParamName (.Name | Singular) }} != nil {
		builtValue := {{ ToParamName (.Name | Singular) }}.Build().(*{{ $.ModelsPackage | base }}.{{ .ElementType }})
		b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }} = append(b.{{ ToLowerFirst $.Struct.Name }}.{{ .Name }}, builtValue)
	}
	return b
}
{{- end }}
{{- end }}

// WithValidation adds a custom validation function
func (b *{{ .Struct.Name }}Builder) WithValidation(validationFunc func(*{{ .ModelsPackage | base }}.{{ .Struct.Name }}) error) *{{ .Struct.Name }}Builder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the {{ .Struct.Name }}
func (b *{{ .Struct.Name }}Builder) Build() interface{} {
	return b.{{ ToLowerFirst .Struct.Name }}
}

// BuildPtr builds the {{ .Struct.Name }} and returns a pointer
func (b *{{ .Struct.Name }}Builder) BuildPtr() *{{ .ModelsPackage | base }}.{{ .Struct.Name }} {
	return b.{{ ToLowerFirst .Struct.Name }}
}

// BuildAndValidate builds the {{ .Struct.Name }} and validates it
func (b *{{ .Struct.Name }}Builder) BuildAndValidate() (*{{ .ModelsPackage | base }}.{{ .Struct.Name }}, error) {
	if b == nil || b.{{ ToLowerFirst .Struct.Name }} == nil {
		return nil, fmt.Errorf("builder is not properly initialized")
	}

	{{ ToLowerFirst .Struct.Name }} := b.{{ ToLowerFirst .Struct.Name }}

	// Run custom validation functions
	for i, validationFunc := range b.validationFuncs {
		if validationFunc == nil {
			continue // Skip nil validators
		}
		if err := validationFunc({{ ToLowerFirst .Struct.Name }}); err != nil {
			return nil, fmt.Errorf("custom validation failed at index %d: %w", i, err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := any({{ ToLowerFirst .Struct.Name }}).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return {{ ToLowerFirst .Struct.Name }}, fmt.Errorf("model validation failed: %w", err)
		}
	}

	return {{ ToLowerFirst .Struct.Name }}, nil
}

// MustBuild builds the {{ .Struct.Name }} and panics if validation fails
func (b *{{ .Struct.Name }}Builder) MustBuild() *{{ .ModelsPackage | base }}.{{ .Struct.Name }} {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *{{ .Struct.Name }}Builder) Clone() *{{ .Struct.Name }}Builder {
	if b == nil || b.{{ ToLowerFirst .Struct.Name }} == nil {
		return New{{ .Struct.Name }}Builder()
	}

	// Deep copy the struct
	cloned{{ .Struct.Name }} := *b.{{ ToLowerFirst .Struct.Name }}

	// Create new builder with cloned data
	clonedBuilder := &{{ .Struct.Name }}Builder{
		{{ ToLowerFirst .Struct.Name }}: &cloned{{ .Struct.Name }},
		validationFuncs: make([]func(*{{ .ModelsPackage | base }}.{{ .Struct.Name }}) error, 0, len(b.validationFuncs)),
	}

	// Copy validation functions safely
	if len(b.validationFuncs) > 0 {
		clonedBuilder.validationFuncs = append(clonedBuilder.validationFuncs, b.validationFuncs...)
	}

	return clonedBuilder
}
`

// RegistryTemplate contains the builder registry code
const RegistryTemplate = `import (
	"fmt"
)

// BuilderRegistry is a registry of builder functions
type BuilderRegistry struct {
	builders map[string]func() interface{}
}

// NewBuilderRegistry creates a new BuilderRegistry
func NewBuilderRegistry() *BuilderRegistry {
	return &BuilderRegistry{
		builders: make(map[string]func() interface{}),
	}
}

// Register registers a builder function
func (r *BuilderRegistry) Register(typeName string, builderFunc func() interface{}) {
	r.builders[typeName] = builderFunc
}

// Get gets a builder function by type name
func (r *BuilderRegistry) Get(typeName string) (func() interface{}, bool) {
	builderFunc, ok := r.builders[typeName]
	return builderFunc, ok
}

// Create creates a new builder for the given type
func (r *BuilderRegistry) Create(typeName string) (interface{}, error) {
	builderFunc, ok := r.Get(typeName)
	if !ok {
		return nil, fmt.Errorf("no builder registered for type %s", typeName)
	}
	return builderFunc(), nil
}

// DefaultRegistry is the default builder registry
var DefaultRegistry = NewBuilderRegistry()

// Register registers a builder function with the default registry
func Register(typeName string, builderFunc func() interface{}) {
	DefaultRegistry.Register(typeName, builderFunc)
}

// Get gets a builder function by type name from the default registry
func Get(typeName string) (func() interface{}, bool) {
	return DefaultRegistry.Get(typeName)
}

// Create creates a new builder for the given type using the default registry
func Create(typeName string) (interface{}, error) {
	return DefaultRegistry.Create(typeName)
}
`
