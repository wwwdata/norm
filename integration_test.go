package norm_test

import (
	"os"

	. "github.com/wwwdata/norm"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type TestDummy struct {
	Name string
	Age  int
}

func (t TestDummy) GetLabel() string {
	return "TestDummy"
}

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

	Describe("Creates and retrieves nodes", func() {
		var db Database

		BeforeEach(func() {
			var err error
			db, err = Connect(neo4jURL)
			Expect(err).ToNot(HaveOccurred())
			err = db.ExecuteCypher(`
				MATCH (n)
				OPTIONAL MATCH (n)-[r]-()
				DELETE n, r
			`)
			Expect(err).ToNot(HaveOccurred())
		})

		It("Creates a single node and retrieves it", func() {
			node := TestDummy{
				Name: "Timmy",
				Age:  7,
			}

			err := db.CreateNode(node).Commit()
			Expect(err).ToNot(HaveOccurred())

			nodesFromDb := []TestDummy{}
			dbObject := db.FindNodes("TestDummy", &nodesFromDb)
			err = dbObject.Commit()
			Expect(err).ToNot(HaveOccurred())
			result := dbObject.GetNodes()
			Expect(*result).To(HaveLen(1))
		})
	})
})
