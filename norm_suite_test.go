package norm_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestNorm(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Norm Suite")
}
