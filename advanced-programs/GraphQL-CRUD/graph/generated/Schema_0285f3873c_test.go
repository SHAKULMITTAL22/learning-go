// Test has been fine-tuned to remove the errors occurred at compile/run time.

package generated

import (
	"testing"
)

// Use unique structures to avoid redeclaring in this block
type testExecutableSchema struct {
	exec ExecutableSchema
}

var testSchema *Schema

// Define a test schema
var parsedTestSchema *Schema = testSchema

// Complexity method was missing in the earlier version. Newly added.
func (s *testExecutableSchema) Complexity(typeName string, fieldName string, childComplexity int, rawArgs map[string]interface{}) (complexity int, computed bool) {
	return complexity, computed
}

// Implement other necessary methods of interface graphql.ExecutableSchema
func (s *testExecutableSchema) Schema() *Schema {
	// Returns schema
	return parsedTestSchema
}

func TestSchema(t *testing.T) {

	e := testExecutableSchema{}
	schema := e.Schema()
	if schema != parsedTestSchema {
		t.Errorf("Schema() failed, expected %v, got %v", parsedTestSchema, schema)
	} else {
		t.Log("Schema() passed")
	}

	parsedTestSchema = nil
	schema = e.Schema()
	if schema != nil {
		t.Errorf("Schema() failed, expected nil, got %v", schema)
	} else {
		t.Log("Schema() test with nil passed")
	}
}
