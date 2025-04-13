package department_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDepartment(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Department Suite")
}
