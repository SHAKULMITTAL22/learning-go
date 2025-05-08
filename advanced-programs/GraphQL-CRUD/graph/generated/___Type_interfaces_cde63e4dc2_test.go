package generated

import (
	"bytes"
	"context"
	"errors"
	"reflect"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

func Test___Type_interfaces_cde63e4dc2(t *testing.T) {
	// Test with valid data
	t.Run("should execute ___Type_interfaces successfully", func(t *testing.T) {
		var collectedObjField graphql.CollectedField
		execContext := executionContext{} 
		TypeInterface := introspection.Type{}

		_, err := execContext.___Type_interfaces(context.Background(), collectedObjField, &TypeInterface)
		if err != nil {
			t.Error("Failed to execute ___Type_interfaces, error: ", err.Error())
		} else {
			t.Log("___Type_interfaces executed successfully.")
		}
	})

	// Test with incorrect data
	t.Run("should fail to execute ___Type_interfaces", func(t *testing.T) {
		execContext := executionContext{} 
		TypeInterface := introspection.Type{}

		_, err := execContext.___Type_interfaces(context.Background(), nil, &TypeInterface)
		if err == nil {
			t.Fail()
			t.Log("___Type_interfaces execution supposed to fail, but it didn't")
		} else {
			t.Log("___Type_interfaces execution failed as expected.")
		}
	})
}
