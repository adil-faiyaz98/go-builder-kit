package education_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestEducation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Education Suite")
}
