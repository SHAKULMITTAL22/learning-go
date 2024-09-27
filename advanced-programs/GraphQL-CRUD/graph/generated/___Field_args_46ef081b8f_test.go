package generated

import (
	"bytes"
	"context"
	"testing"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
)

var testObj = &introspection.Field{}

var resolverMiddleware = func(ctx context.Context) (interface{}, error) {
	return nil, nil
}

type executionContext struct {
	ResolverMiddleware func(ctx context.Context) (interface{}, error)
}

func (ec *executionContext) ___Field_args(ctx context.Context, collectedField graphql.CollectedField, obj *introspection.Field) interface{} {
	res, _ := ec.ResolverMiddleware(ctx)
	return res
}

func Test___Field_args_46ef081b8f(t *testing.T) {
	t.Run("Arguments are returned successfully", func(t *testing.T) {
		ec := &executionContext{
			ResolverMiddleware: resolverMiddleware,
		}
		
		collectedField := graphql.CollectedField{}
		ret := ec.___Field_args(context.Background(), collectedField, testObj)
		
		if assert.NotNil(t, ret) {
			t.Log("Arguments are retrieved successfully")
		} else {
			t.Error("Failed to retrieve arguments", "CollectedField:", collectedField, "Object:", testObj)
		}
	})

	t.Run("Resolver middleware returns an error", func(t *testing.T) {
		errorMsg := "test error message"
		errorResolver := func(ctx context.Context) (interface{}, error) {
			return nil, errors.New(errorMsg)
		}
		
		ec := &executionContext{
			ResolverMiddleware: errorResolver,
		}

		collectedField := graphql.CollectedField{}
		ret := ec.___Field_args(context.Background(), collectedField, testObj)
		
		if assert.Equal(t, nil, ret) {
			t.Log("Correctly handled an error from the resolver middleware")
		} else {
			t.Error("Failed to handle an error from the resolver middleware", "CollectedField:", collectedField, "Object:", testObj)
		}
	})

	t.Run("Return value from resolver middleware is nil", func(t *testing.T) {
		nilResolver := func(ctx context.Context) (interface{}, error) {
			return nil, nil
		}
		
		ec := &executionContext{
			ResolverMiddleware: nilResolver,
		}

		collectedField := graphql.CollectedField{}
		ret := ec.___Field_args(context.Background(), collectedField, testObj)
		
		if assert.Equal(t, nil, ret) {
			t.Log("Correctly handled nil return from the resolver middleware")
		} else {
			t.Error("Failed to handle nil return from the resolver middleware", "CollectedField:", collectedField, "Object:", testObj)
		}
	})
}
