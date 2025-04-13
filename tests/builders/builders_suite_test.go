package builders_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBuildersPackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Builders Package Suite")
}
