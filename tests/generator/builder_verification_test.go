package generator_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/adil-faiyaz98/go-builder-kit/models"
	"github.com/adil-faiyaz98/go-builder-kit/pkg/generator"
)

// TestGeneratedBuilders tests that the generated builders correctly handle all nested types
func TestGeneratedBuilders(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "builder-verification-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Get all model files
	modelFiles, err := filepath.Glob("../../models/*.go")
	if err != nil {
		t.Fatalf("Failed to find model files: %v", err)
	}

	// Create a generator
	gen := generator.NewGenerator(generator.Options{
		PackageName:   "builders",
		ModelsPackage: "github.com/adil-faiyaz98/go-builder-kit/models",
		Verbose:       true,
	})

	// Process each model file
	for _, modelFile := range modelFiles {
		t.Logf("Processing model file: %s", modelFile)

		// Parse the file
		structs, err := gen.ParseFile(modelFile)
		if err != nil {
			t.Fatalf("Failed to parse file %s: %v", modelFile, err)
		}

		// Process each struct
		for _, structInfo := range structs {
			t.Logf("  Verifying struct: %s", structInfo.Name)

			// Generate builder code
			code, err := gen.GenerateBuilder(structInfo)
			if err != nil {
				t.Fatalf("Failed to generate builder for %s: %v", structInfo.Name, err)
			}

			// Write builder to file
			outputFile := filepath.Join(tempDir, fmt.Sprintf("%s_builder.go", generator.ToSnakeCase(structInfo.Name)))
			if err := os.WriteFile(outputFile, []byte(code), 0644); err != nil {
				t.Fatalf("Failed to write builder file %s: %v", outputFile, err)
			}

			// Verify that the builder has methods for all fields
			verifyBuilderMethods(t, structInfo, code)

			// Verify that the builder handles nested types
			verifyNestedTypes(t, structInfo, code)
		}
	}
}

// verifyBuilderMethods verifies that the builder has methods for all fields
func verifyBuilderMethods(t *testing.T, structInfo generator.StructInfo, code string) {
	for _, field := range structInfo.Fields {
		methodName := fmt.Sprintf("With%s", field.Name)
		if !strings.Contains(code, methodName) {
			t.Errorf("Builder for %s is missing method %s", structInfo.Name, methodName)
		}
	}
}

// verifyNestedTypes verifies that the builder handles nested types
func verifyNestedTypes(t *testing.T, structInfo generator.StructInfo, code string) {
	for _, field := range structInfo.Fields {
		if field.IsNested {
			// Check if the field is a slice of nested types
			if field.IsSlice {
				// Verify that the builder handles slices of nested types
				expectedCode := fmt.Sprintf("b.%s.%s = append(b.%s.%s, %s.Build().(models.%s))",
					generator.ToLowerFirst(structInfo.Name), field.Name,
					generator.ToLowerFirst(structInfo.Name), field.Name,
					generator.ToLowerFirst(field.Name), field.ElementType)
				if !strings.Contains(code, expectedCode) {
					t.Errorf("Builder for %s does not correctly handle slice of nested type %s", structInfo.Name, field.Name)
				}
			} else {
				// Verify that the builder handles nested types
				expectedCode1 := fmt.Sprintf("b.%s.%s = %s.Build().(models.%s)",
					generator.ToLowerFirst(structInfo.Name), field.Name,
					generator.ToLowerFirst(field.Name), field.Type)
				expectedCode2 := fmt.Sprintf("value := %s.Build().(models.%s)",
					generator.ToLowerFirst(field.Name), field.Type)
				if !strings.Contains(code, expectedCode1) && !strings.Contains(code, expectedCode2) {
					t.Errorf("Builder for %s does not correctly handle nested type %s", structInfo.Name, field.Name)
				}
			}
		}
	}
}

// TestModelReflection uses reflection to verify that all models have corresponding builders
func TestModelReflection(t *testing.T) {
	// Get all types in the models package
	modelTypes := getModelTypes()

	// Verify that each model has a corresponding builder
	for typeName, modelType := range modelTypes {
		t.Logf("Checking model: %s", typeName)

		// Skip interfaces and non-struct types
		if modelType.Kind() != reflect.Struct {
			continue
		}

		// Verify that the model has a corresponding builder
		builderName := fmt.Sprintf("%sBuilder", typeName)
		builderPath := fmt.Sprintf("../../builders/%s_builder.go", generator.ToSnakeCase(typeName))
		if _, err := os.Stat(builderPath); os.IsNotExist(err) {
			t.Errorf("Model %s does not have a corresponding builder %s", typeName, builderName)
		}

		// Verify that all fields in the model are handled by the builder
		verifyAllFieldsHandled(t, typeName, modelType)
	}
}

// getModelTypes returns all types in the models package
func getModelTypes() map[string]reflect.Type {
	modelTypes := make(map[string]reflect.Type)

	// Use reflection to get all types in the models package
	modelType := reflect.TypeOf(models.Person{})
	modelTypes[modelType.Name()] = modelType

	// Add more model types
	modelTypes["Address"] = reflect.TypeOf(models.Address{})
	modelTypes["Education"] = reflect.TypeOf(models.Education{})
	modelTypes["Course"] = reflect.TypeOf(models.Course{})
	modelTypes["Employment"] = reflect.TypeOf(models.Employment{})
	modelTypes["Company"] = reflect.TypeOf(models.Company{})
	modelTypes["Department"] = reflect.TypeOf(models.Department{})
	modelTypes["Project"] = reflect.TypeOf(models.Project{})
	modelTypes["Bank"] = reflect.TypeOf(models.Bank{})
	modelTypes["Account"] = reflect.TypeOf(models.Account{})
	modelTypes["Transaction"] = reflect.TypeOf(models.Transaction{})
	modelTypes["Investment"] = reflect.TypeOf(models.Investment{})
	modelTypes["Portfolio"] = reflect.TypeOf(models.Portfolio{})
	modelTypes["PerformanceRecord"] = reflect.TypeOf(models.PerformanceRecord{})
	modelTypes["Stock"] = reflect.TypeOf(models.Stock{})
	modelTypes["Bond"] = reflect.TypeOf(models.Bond{})
	modelTypes["ETF"] = reflect.TypeOf(models.ETF{})
	modelTypes["HealthProfile"] = reflect.TypeOf(models.HealthProfile{})
	modelTypes["Insurance"] = reflect.TypeOf(models.Insurance{})
	modelTypes["Medication"] = reflect.TypeOf(models.Medication{})
	modelTypes["MedicalRecord"] = reflect.TypeOf(models.MedicalRecord{})
	modelTypes["DigitalProfile"] = reflect.TypeOf(models.DigitalProfile{})
	modelTypes["Device"] = reflect.TypeOf(models.Device{})
	modelTypes["Application"] = reflect.TypeOf(models.Application{})
	modelTypes["OnlineAccount"] = reflect.TypeOf(models.OnlineAccount{})
	modelTypes["Subscription"] = reflect.TypeOf(models.Subscription{})
	modelTypes["UserPreferences"] = reflect.TypeOf(models.UserPreferences{})
	modelTypes["UserActivity"] = reflect.TypeOf(models.UserActivity{})
	modelTypes["GeoLocation"] = reflect.TypeOf(models.GeoLocation{})

	return modelTypes
}

// verifyAllFieldsHandled verifies that all fields in the model are handled by the builder
func verifyAllFieldsHandled(t *testing.T, typeName string, modelType reflect.Type) {
	// Read the builder file
	builderPath := fmt.Sprintf("../../builders/%s_builder.go", generator.ToSnakeCase(typeName))
	builderCode, err := os.ReadFile(builderPath)
	if err != nil {
		t.Errorf("Failed to read builder file %s: %v", builderPath, err)
		return
	}

	// Check each field in the model
	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)

		// Skip unexported fields
		if !field.IsExported() {
			continue
		}

		// Check if the builder has a method for this field
		methodName := fmt.Sprintf("With%s", field.Name)
		if !strings.Contains(string(builderCode), methodName) {
			t.Errorf("Builder for %s is missing method %s", typeName, methodName)
		}

		// Check if the field is a nested type
		if field.Type.Kind() == reflect.Struct && field.Type.PkgPath() == modelType.PkgPath() {
			// Check if the builder handles nested types
			expectedCode := fmt.Sprintf("%s.Build().(models.%s)",
				generator.ToLowerFirst(field.Name), field.Type.Name())
			if !strings.Contains(string(builderCode), expectedCode) {
				t.Errorf("Builder for %s does not correctly handle nested type %s", typeName, field.Name)
			}
		}

		// Check if the field is a slice of nested types
		if field.Type.Kind() == reflect.Slice &&
			field.Type.Elem().Kind() == reflect.Struct &&
			field.Type.Elem().PkgPath() == modelType.PkgPath() {
			// Check if the builder handles slices of nested types
			expectedCode := fmt.Sprintf("append(b.%s.%s, %s.Build().(models.%s))",
				generator.ToLowerFirst(typeName), field.Name,
				generator.ToLowerFirst(field.Name), field.Type.Elem().Name())
			if !strings.Contains(string(builderCode), expectedCode) {
				t.Errorf("Builder for %s does not correctly handle slice of nested type %s", typeName, field.Name)
			}
		}
	}
}
