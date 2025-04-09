package performance_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

// BenchmarkSimplePersonBuilder benchmarks creating a simple person
func BenchmarkSimplePersonBuilder(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		builders.NewPersonBuilder().
			WithID("123").
			WithName("John Doe").
			WithAge(30).
			WithEmail("john.doe@example.com").
			Build()
	}
}

// BenchmarkComplexPersonBuilder benchmarks creating a complex person with nested objects
func BenchmarkComplexPersonBuilder(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		builders.NewPersonBuilder().
			WithID("123").
			WithName("John Doe").
			WithAge(30).
			WithEmail("john.doe@example.com").
			WithAddress(
				builders.NewAddressBuilder().
					WithStreet("123 Main St").
					WithCity("New York").
					WithState("NY").
					WithPostalCode("10001").
					WithCountry("USA"),
			).
			WithEducation(
				builders.NewEducationBuilder().
					WithDegree("Bachelor of Science").
					WithInstitution("Harvard University").
					WithStartDate("2018-09-01").
					WithEndDate("2022-05-15").
					WithGPA(3.8),
			).
			WithEmployment(
				builders.NewEmploymentBuilder().
					WithPosition("Software Engineer").
					WithDepartment("Engineering").
					WithStartDate("2022-06-01").
					WithSalary(100000.0),
			).
			Build()
	}
}

// BenchmarkPersonBuilderWithValidation benchmarks creating a person with validation
func BenchmarkPersonBuilderWithValidation(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = builders.NewPersonBuilder().
			WithID("123").
			WithName("John Doe").
			WithAge(30).
			WithEmail("john.doe@example.com").
			BuildWithValidation()
	}
}

// BenchmarkBulkPersonCreation benchmarks creating many persons
func BenchmarkBulkPersonCreation(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		createManyPersons(100)
	}
}

// createManyPersons creates many persons
func createManyPersons(count int) {
	persons := make([]interface{}, count)
	for i := 0; i < count; i++ {
		persons[i] = builders.NewPersonBuilder().
			WithID(fmt.Sprintf("ID-%d", i)).
			WithName(fmt.Sprintf("Person %d", i)).
			WithAge(20 + (i % 50)).
			WithEmail(fmt.Sprintf("person%d@example.com", i)).
			Build()
	}
}

// BenchmarkBulkPersonCreationWithReuse benchmarks creating many persons with builder reuse
func BenchmarkBulkPersonCreationWithReuse(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		createManyPersonsWithReuse(100)
	}
}

// createManyPersonsWithReuse creates many persons with builder reuse
func createManyPersonsWithReuse(count int) {
	// Create a base builder
	baseBuilder := builders.NewPersonBuilder().
		WithAge(30).
		WithEmail("example@example.com")

	// Create persons by reusing the base builder
	persons := make([]*models.Person, count)
	for i := 0; i < count; i++ {
		// Clone the base builder and customize
		persons[i] = baseBuilder.Clone().
			WithID(fmt.Sprintf("ID-%d", i)).
			WithName(fmt.Sprintf("Person %d", i)).
			Build().(*models.Person)
	}
}

// BenchmarkDeepNestedStructure benchmarks creating a deep nested structure
func BenchmarkDeepNestedStructure(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		createDeepNestedStructure()
	}
}

// createDeepNestedStructure creates a deep nested structure
func createDeepNestedStructure() interface{} {
	// Create a deep nested structure
	return builders.NewPersonBuilder().
		WithID("123").
		WithName("John Doe").
		WithAge(30).
		WithEmail("john.doe@example.com").
		WithAddress(
			builders.NewAddressBuilder().
				WithStreet("123 Main St").
				WithCity("New York").
				WithState("NY").
				WithPostalCode("10001").
				WithCountry("USA"),
		).
		WithEmployment(
			builders.NewEmploymentBuilder().
				WithPosition("Senior Software Engineer").
				WithDepartment("Engineering").
				WithStartDate("2020-01-15").
				WithSalary(120000.0).
				WithIsCurrent(true).
				WithCompany(
					builders.NewCompanyBuilder().
						WithID("COMP-001").
						WithName("Tech Innovations Inc.").
						WithIndustry("Technology").
						WithDescription("Leading technology company").
						WithFounded("2010-01-01").
						WithLocation(
							builders.NewAddressBuilder().
								WithStreet("456 Tech Blvd").
								WithCity("San Francisco").
								WithState("CA").
								WithPostalCode("94107").
								WithCountry("USA"),
						).
						WithDepartment(
							builders.NewDepartmentBuilder().
								WithName("Engineering").
								WithDescription("Software development department").
								WithHeadCount(50).
								WithBudget(5000000.0).
								WithProject(
									builders.NewProjectBuilder().
										WithName("Mobile App").
										WithDescription("Cross-platform mobile application").
										WithStartDate("2022-01-01").
										WithEndDate("2022-12-31").
										WithStatus("In Progress").
										WithBudget(500000.0),
								),
						),
				),
		).
		WithDigitalProfile(
			builders.NewDigitalProfileBuilder().
				WithDevice(
					builders.NewDeviceBuilder().
						WithType("Smartphone").
						WithModel("iPhone 13").
						WithSerialNumber("ABCD1234").
						WithPurchaseDate("2022-01-01").
						WithOS("iOS 16.5").
						WithLastUsed("2023-05-01").
						WithApp(
							builders.NewApplicationBuilder().
								WithName("Social Media App").
								WithVersion("2.1.0").
								WithInstallDate("2022-01-15").
								WithLastUsed("2023-05-01").
								WithPermission("Camera").
								WithPermission("Microphone"),
						),
				).
				WithAccount(
					builders.NewOnlineAccountBuilder().
						WithPlatform("Social Media").
						WithUsername("user123").
						WithEmail("user@example.com").
						WithCreationDate("2020-01-01").
						WithLastLogin("2023-05-01").
						WithStatus("Active").
						WithSubscription(
							builders.NewSubscriptionBuilder().
								WithPlan("Premium").
								WithStartDate("2022-01-01").
								WithEndDate("2023-01-01").
								WithPrice(9.99).
								WithBillingCycle("Monthly").
								WithAutoRenew(true).
								WithFeature("Ad-free").
								WithFeature("Offline access"),
						),
				),
		).
		Build()
}

// TestPerformanceComparison compares the performance of different builder approaches
func TestPerformanceComparison(t *testing.T) {
	// Skip in short mode
	if testing.Short() {
		t.Skip("Skipping performance comparison in short mode")
	}

	// Test creating a simple person
	t.Run("SimplePersonBuilder", func(t *testing.T) {
		start := time.Now()
		iterations := 10000

		for i := 0; i < iterations; i++ {
			builders.NewPersonBuilder().
				WithID("123").
				WithName("John Doe").
				WithAge(30).
				WithEmail("john.doe@example.com").
				Build()
		}

		duration := time.Since(start)
		t.Logf("Created %d simple persons in %v (%.2f ns/op)", iterations, duration, float64(duration.Nanoseconds())/float64(iterations))
	})

	// Test creating a complex person
	t.Run("ComplexPersonBuilder", func(t *testing.T) {
		start := time.Now()
		iterations := 1000

		for i := 0; i < iterations; i++ {
			builders.NewPersonBuilder().
				WithID("123").
				WithName("John Doe").
				WithAge(30).
				WithEmail("john.doe@example.com").
				WithAddress(
					builders.NewAddressBuilder().
						WithStreet("123 Main St").
						WithCity("New York").
						WithState("NY").
						WithPostalCode("10001").
						WithCountry("USA"),
				).
				WithEducation(
					builders.NewEducationBuilder().
						WithDegree("Bachelor of Science").
						WithInstitution("Harvard University").
						WithStartDate("2018-09-01").
						WithEndDate("2022-05-15").
						WithGPA(3.8),
				).
				WithEmployment(
					builders.NewEmploymentBuilder().
						WithPosition("Software Engineer").
						WithDepartment("Engineering").
						WithStartDate("2022-06-01").
						WithSalary(100000.0),
				).
				Build()
		}

		duration := time.Since(start)
		t.Logf("Created %d complex persons in %v (%.2f ns/op)", iterations, duration, float64(duration.Nanoseconds())/float64(iterations))
	})

	// Test creating many persons with and without reuse
	t.Run("BulkCreationComparison", func(t *testing.T) {
		count := 1000

		// Without reuse
		start := time.Now()
		createManyPersons(count)
		durationWithoutReuse := time.Since(start)

		// With reuse
		start = time.Now()
		createManyPersonsWithReuse(count)
		durationWithReuse := time.Since(start)

		t.Logf("Created %d persons without reuse in %v", count, durationWithoutReuse)
		t.Logf("Created %d persons with reuse in %v", count, durationWithReuse)
		t.Logf("Improvement: %.2f%%", 100*(1-float64(durationWithReuse.Nanoseconds())/float64(durationWithoutReuse.Nanoseconds())))
	})

	// Test creating a deep nested structure
	t.Run("DeepNestedStructure", func(t *testing.T) {
		start := time.Now()
		iterations := 100

		for i := 0; i < iterations; i++ {
			createDeepNestedStructure()
		}

		duration := time.Since(start)
		t.Logf("Created %d deep nested structures in %v (%.2f ns/op)", iterations, duration, float64(duration.Nanoseconds())/float64(iterations))
	})
}
