package builders_test

import (
	"testing"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBuilders(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Builders Suite")
}

var _ = Describe("Builders", func() {
	Describe("Project Builder", func() {
		It("should handle string date fields correctly", func() {
			// Create date strings
			startDate := "2023-01-01"
			endDate := "2023-12-31"

			// Create a project with date fields
			projectBuilder := builders.NewProjectBuilder().
				WithName("Test Project").
				WithDescription("A test project").
				WithStartDate(startDate).
				WithEndDate(endDate)

			// Build the project
			project := projectBuilder.BuildPtr()

			// Verify the date fields
			Expect(project.StartDate).To(Equal(startDate))
			Expect(project.EndDate).To(Equal(endDate))
		})

		It("should handle string enum types correctly", func() {
			// Create a project with string enum fields
			projectBuilder := builders.NewProjectBuilder().
				WithName("Test Project").
				WithDescription("A test project").
				WithStatus("Active") // Status is a string enum

			// Build the project
			project := projectBuilder.BuildPtr()

			// Verify the string enum field
			Expect(project.Status).To(Equal("Active"))
		})
	})
})
