package generated

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/99designs/gqlgen/graphql"
)

// TestConfig model for testing
type TestConfig struct {
	Resolvers  interface{}
	Directives interface{}
	Complexity interface{}
}

// TestExecutableSchema model for testing
type TestExecutableSchema struct {
	resolvers  interface{}
	directives interface{}
	complexity interface{}
}

// NewTestExecutableSchema creates a new executable schema
func NewTestExecutableSchema(cfg TestConfig) graphql.ExecutableSchema {
	return &TestExecutableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

// TestNewExecutableSchema_1f81986322 tests the 'NewExecutableSchema' function
func TestNewExecutableSchema_1f81986322(t *testing.T) {
	// setup resolvers, directives and complexity
	resolvers := "resolvers"
	directives := "directives"
	complexity := "complexity"

	// configure the TestConfig object
	config := TestConfig{
		Resolvers:  resolvers,
		Directives: directives,
		Complexity: complexity,
	}

	// create an new TestExecutableSchema
	es := NewTestExecutableSchema(config)

	// Type assertion to *TestExecutableSchema
	schema, ok := es.(*TestExecutableSchema)

	// Check if the type assertion was successful
	if !ok {
		t.Error("NewTestExecutableSchema() returned a wrong type")
	}
	
	// Test if the returned schema matches the config
	assert.Equal(t, config.Resolvers, schema.resolvers, "Resolvers do not match")
	assert.Equal(t, config.Directives, schema.directives,"Directives do not match")
	assert.Equal(t, config.Complexity, schema.complexity, "Complexity does not match")
	
	t.Log("NewTestExecutableSchema function passed the test.")
}
