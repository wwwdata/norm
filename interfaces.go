package norm

// The Label interface must be implemented for every struct and defines the neo4j label name
type Label interface {
	GetLabel() string
}
