package norm

import (
	"reflect"

	"github.com/jmcvetta/neoism"
)

// A Database connects to neo4j and initiates all further queries
type Database struct {
	db *neoism.Database
}

// Connect to a neo4j database
func Connect(url string) (database Database, err error) {
	database = Database{}
	database.db, err = neoism.Connect(url)
	return
}

// CreateNode creates a new node with all fields from a struct that implements the `Label` interface
func (d Database) CreateNode(obj Label) *Query {
	query := Query{
		action:  typeWrite,
		db:      d.db,
		newNode: obj,
	}

	return &query
}

// FindNodes finds all nodes with a given label and scans the results into destination
func (d Database) FindNodes(label string, destination interface{}) *Query {
	query := Query{
		action:      typeRead,
		db:          d.db,
		label:       label,
		resultNodes: reflect.ValueOf(destination),
	}

	return &query
}

// ExecuteCypher executes a raw cypher query
func (d Database) ExecuteCypher(cypher string) error {
	return d.db.Cypher(&neoism.CypherQuery{
		Statement: cypher,
	})
}
