package person_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPerson(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Person Suite")
}
