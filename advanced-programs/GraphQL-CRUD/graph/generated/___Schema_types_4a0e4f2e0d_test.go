package generated

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/stretchr/testify/mock"
)

type mockResolverInterface struct {
	mock.Mock
}

func (m *mockResolverInterface) Types(ctx context.Context) ([]*introspection.Type, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*introspection.Type), args.Error(1)
}

type executionContext struct {
	Resolver *mockResolverInterface
}

func (e *executionContext) ___Schema_types(ctx context.Context, obj *introspection.Schema) ([]*introspection.Type, error) {
	return e.Resolver.Types(ctx)
}

func Test___Schema_types_4a0e4f2e0d(t *testing.T) {

	// A mock object for the resolver interface
	mockResolver := new(mockResolverInterface)
	// Assign our function to the function in the interface
	mockResolver.On("Types", mock.Anything).Return([]*introspection.Type{}, nil)

	// Initialize execution context
	ec := &executionContext{
		Resolver: mockResolver,
	}

	

	t.Run("Test with correct data", func(t *testing.T) {
		obj := &introspection.Schema{}
		_, err := ec.___Schema_types(context.Background(), obj)
		if err != nil {
			t.Error("Expected successful execution but received error")
		} else {
			t.Log("Test passed with correct data")
		}
	})

	t.Run("Test with nil obj", func(t *testing.T) {
		_, err := ec.___Schema_types(context.Background(), nil)
		if err != nil {
			t.Log("Test passed with nil obj")
		} else {
			t.Error("Expected error but received successful execution")
		}
	})

	mockResolver.AssertExpectations(t)
}
