package norm_test

import (
	"os"

	. "github.com/wwwdata/norm"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Integration", func() {
	var neo4jURL string

	BeforeEach(func() {
		neo4jURL = os.Getenv("NORM_NEO4J_URL")
		Expect(neo4jURL).ToNot(BeEmpty(), "NORM_NEO4J_URL env variable must be set")
	})

	Describe("Test Connection", func() {
		It("Connects to a running neo4j instance", func() {
			_, err := Connect(neo4jURL)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
