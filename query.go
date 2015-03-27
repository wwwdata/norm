package norm

import (
	"errors"
	"reflect"

	"github.com/jmcvetta/neoism"
)

const (
	typeRead  = "read"
	typeWrite = "write"
)

// The Query object which constructs all the magic. Do not instantiate this directly, there is a chainable api for everything!
type Query struct {
	db          *neoism.Database
	action      string // can be either typeRead or typeWrite
	label       string
	resultNodes interface{} // must be a pointer to a slice of something that implements the interface `Label`
	newNode     Label
	lastID      int
}

func (q *Query) createNode() error {
	props := neoism.Props{}
	reflectValue := reflect.ValueOf(q.newNode)
	if reflectValue.Kind() != reflect.Struct {
		return errors.New("obj must be of type struct")
	}

	for i := 0; i < reflectValue.NumField(); i++ {
		field := reflectValue.Field(i)
		name := reflectValue.Type().Field(i).Name
		props[name] = field.Interface()
	}

	node, err := q.db.CreateNode(props)
	if err != nil {
		return err
	}

	err = node.AddLabel(q.newNode.GetLabel())
	if err != nil {
		return err
	}

	q.lastID = node.Id()

	return nil
}

func (q *Query) findNodes() error {
	refValue := reflect.ValueOf(q.resultNodes)
	if refValue.Type().Kind() != reflect.Ptr || refValue.Elem().Kind() != reflect.Slice {
		return errors.New("desination must be a pointer to a slice")
	}

	destinationType := reflect.TypeOf(q.resultNodes).Elem().Elem()
	result := []interface{}{}

	nodes, err := q.db.NodesByLabel(q.label)
	if err != nil {
		return err
	}

	for _, node := range nodes {
		props, err := node.Properties()
		if err != nil {
			return err
		}

		destination := reflect.New(destinationType).Elem()

		for key, prop := range props {
			field := destination.FieldByName(key)
			field.Set(reflect.ValueOf(prop).Convert(field.Type()))
		}

		result = append(result, destination.Interface())
	}

	// TODO: this does not work, the reference is gone
	q.resultNodes = &result

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

// GetNodes shit, I intended to directly write in the variable passed by reference from the db object... figure out how to do that
func (q *Query) GetNodes() *[]interface{} {
	return q.resultNodes.(*[]interface{})
}

// GetLastID returns the internal neo4j ID of the last generated node. This should not be interesting, because the IDs
// in neo4j are always different if you restart it. If a node could not be created, you will notice it by an error
func (q *Query) GetLastID() int {
	return q.lastID
}
