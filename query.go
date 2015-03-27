package norm

import "errors"

const (
	typeRead  = "read"
	typeWrite = "write"
)

// The Query object which constructs all the magic. Do not instantiate this directly, there is a chainable api for everything!
type Query struct {
	db          Database
	action      string // can be either typeRead or typeWrite
	label       string
	resultNodes *[]interface{}
	newNode     Label
}

func (q *Query) createNode() error {
	return nil
}

func (q *Query) findNodes() error {
	return nil
}

// Commit executes the defined steps of the chainable api before
func (q *Query) Commit() error {
	switch q.action {
	case typeRead:
		return q.findNodes()
	case typeWrite:
		return q.createNode()
	default:
		return errors.New("Something really went terribly wrong")
	}
}
