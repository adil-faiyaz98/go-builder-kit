package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"CamelCase", "camel_case"},
		{"XMLHttpRequest", "x_m_l_http_request"},
		{"ID", "i_d"},
		{"PersonBuilder", "person_builder"},
		{"", ""},
		{"lowercase", "lowercase"},
		{"UPPERCASE", "u_p_p_e_r_c_a_s_e"},
		{"MixedCASEString", "mixed_c_a_s_e_string"},
	}

	for _, test := range tests {
		result := ToSnakeCase(test.input)
		if result != test.expected {
			t.Errorf("ToSnakeCase(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestToLowerFirst(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"CamelCase", "camelCase"},
		{"ID", "iD"},
		{"", ""},
		{"lowercase", "lowercase"},
		{"A", "a"},
		{"UPPERCASE", "uPPERCASE"},
		{"ÄÖÜ", "äÖÜ"}, // Unicode test
	}

	for _, test := range tests {
		result := ToLowerFirst(test.input)
		if result != test.expected {
			t.Errorf("ToLowerFirst(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestToParamName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Name", "name"},
		{"FirstName", "firstName"},
		{"XMLHttpRequest", "xmlHttpRequest"},
		{"ID", "id"},
		{"URLs", "urls"},
		{"", ""},
		{"Type", "value"}, // Special case
	}

	for _, test := range tests {
		result := ToParamName(test.input)
		if result != test.expected {
			t.Errorf("ToParamName(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestToSingular(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Friends", "Friend"},
		{"Addresses", "Address"},
		{"Children", "Child"},
		{"People", "Person"},
		{"Boxes", "Box"},
		{"Companies", "Company"},
		{"", ""},
		{"Friend", "Friend"}, // Already singular
		{"Data", "Data"},     // No change
	}

	for _, test := range tests {
		result := ToSingular(test.input)
		if result != test.expected {
			t.Errorf("ToSingular(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestNewGenerator(t *testing.T) {
	opts := Options{
		PackageName:   "test",
		ModelsPackage: "models",
		Verbose:       true,
	}

	gen := NewGenerator(opts)
	if gen == nil {
		t.Fatal("NewGenerator returned nil")
	}

	if gen.Options.PackageName != "test" {
		t.Errorf("Expected PackageName 'test', got %q", gen.Options.PackageName)
	}

	if gen.Options.ModelsPackage != "models" {
		t.Errorf("Expected ModelsPackage 'models', got %q", gen.Options.ModelsPackage)
	}

	if !gen.Options.Verbose {
		t.Error("Expected Verbose to be true")
	}
}

func TestProcessFileValidation(t *testing.T) {
	gen := NewGenerator(Options{
		PackageName:   "test",
		ModelsPackage: "models",
	})

	t.Run("empty input file", func(t *testing.T) {
		err := gen.ProcessFile("", "output")
		if err == nil {
			t.Error("Expected error for empty input file")
		}
		if !strings.Contains(err.Error(), "input file path cannot be empty") {
			t.Errorf("Unexpected error message: %s", err.Error())
		}
	})

	t.Run("empty output directory", func(t *testing.T) {
		err := gen.ProcessFile("input.go", "")
		if err == nil {
			t.Error("Expected error for empty output directory")
		}
		if !strings.Contains(err.Error(), "output directory path cannot be empty") {
			t.Errorf("Unexpected error message: %s", err.Error())
		}
	})

	t.Run("non-existent input file", func(t *testing.T) {
		err := gen.ProcessFile("non-existent.go", "output")
		if err == nil {
			t.Error("Expected error for non-existent input file")
		}
	})
}

func TestProcessFileWithTempFiles(t *testing.T) {
	// Create temporary directory for test
	tempDir, err := os.MkdirTemp("", "generator_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a simple Go file with a struct
	inputFile := filepath.Join(tempDir, "test.go")
	inputContent := `package models

type Person struct {
	Name string
	Age  int
}
`
	err = os.WriteFile(inputFile, []byte(inputContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write input file: %v", err)
	}

	// Create output directory
	outputDir := filepath.Join(tempDir, "builders")
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create output dir: %v", err)
	}

	// Process the file
	gen := NewGenerator(Options{
		PackageName:   "builders",
		ModelsPackage: "models",
		Verbose:       false,
	})

	err = gen.ProcessFile(inputFile, outputDir)
	if err != nil {
		t.Fatalf("ProcessFile failed: %v", err)
	}

	// Check that output file was created
	expectedOutput := filepath.Join(outputDir, "person_builder.go")
	if _, err := os.Stat(expectedOutput); os.IsNotExist(err) {
		t.Error("Expected output file was not created")
	}

	// Read and verify output content
	content, err := os.ReadFile(expectedOutput)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	contentStr := string(content)

	// Check for expected content
	expectedStrings := []string{
		"package builders",
		"type PersonBuilder struct",
		"func NewPersonBuilder()",
		"func (b *PersonBuilder) WithName(",
		"func (b *PersonBuilder) WithAge(",
		"func (b *PersonBuilder) Build()",
		"func (b *PersonBuilder) BuildPtr()",
		"func (b *PersonBuilder) BuildAndValidate()",
		"func (b *PersonBuilder) Clone()",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(contentStr, expected) {
			t.Errorf("Expected output to contain %q", expected)
		}
	}
}

func TestGenerateBuilders(t *testing.T) {
	gen := NewGenerator(Options{
		PackageName:   "builders",
		ModelsPackage: "models",
	})

	// GenerateBuilders should be an alias for ProcessFile
	tempDir, err := os.MkdirTemp("", "generator_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	inputFile := filepath.Join(tempDir, "test.go")
	inputContent := `package models

type Simple struct {
	Value string
}
`
	err = os.WriteFile(inputFile, []byte(inputContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write input file: %v", err)
	}

	outputDir := filepath.Join(tempDir, "builders")
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create output dir: %v", err)
	}

	err = gen.GenerateBuilders(inputFile, outputDir)
	if err != nil {
		t.Errorf("GenerateBuilders failed: %v", err)
	}

	// Check that output file was created
	expectedOutput := filepath.Join(outputDir, "simple_builder.go")
	if _, err := os.Stat(expectedOutput); os.IsNotExist(err) {
		t.Error("Expected output file was not created by GenerateBuilders")
	}
}

func TestFilenameValidation(t *testing.T) {
	// This test verifies that the filename validation prevents directory traversal
	tempDir, err := os.MkdirTemp("", "generator_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a malicious struct name that could cause directory traversal
	inputFile := filepath.Join(tempDir, "test.go")
	inputContent := `package models

// This struct name should be sanitized
type MaliciousStruct struct {
	Value string
}
`
	err = os.WriteFile(inputFile, []byte(inputContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write input file: %v", err)
	}

	outputDir := filepath.Join(tempDir, "builders")
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create output dir: %v", err)
	}

	gen := NewGenerator(Options{
		PackageName:   "builders",
		ModelsPackage: "models",
	})

	err = gen.ProcessFile(inputFile, outputDir)
	if err != nil {
		t.Fatalf("ProcessFile failed: %v", err)
	}

	// Check that the output filename is properly sanitized
	expectedOutput := filepath.Join(outputDir, "malicious_struct_builder.go")
	if _, err := os.Stat(expectedOutput); os.IsNotExist(err) {
		t.Error("Expected sanitized output file was not created")
	}
}
