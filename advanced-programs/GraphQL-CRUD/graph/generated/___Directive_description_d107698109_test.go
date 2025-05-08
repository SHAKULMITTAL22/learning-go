package generated

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
)

type executionContext struct {}

func (ec *executionContext) ___Directive_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	if obj == nil {
		return graphql.Null
	}
	return graphql.MarshalString(obj.Description)
}

func Test___Directive_description_d107698109(t *testing.T) {
	testCases := []struct {
		name    string
		field   graphql.CollectedField
		obj     *introspection.Directive
		wantErr bool
	}{
		{
			name: "Test Case 1: Successful execution",
			field: graphql.CollectedField{},
			obj: &introspection.Directive{
				Name:        "testDirective",
				Description: "This is a test directive",
				Locations:   []string{"LOCATION1", "LOCATION2"},
				Args:        nil,
				OnOperation: false,
				OnFragment:  false,
				OnField:     false,
			},
			wantErr: false,
		},
		{
			name:    "Test Case 2: Error execution",
			field:   graphql.CollectedField{},
			obj:     nil,
			wantErr: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ec := &executionContext{}
			got := ec.___Directive_description(context.Background(), tt.field, tt.obj)

			if tt.wantErr {
				if got != graphql.Null {
					t.Errorf("___Directive_description() did not returned expected error, got = %v", got)
					t.Logf("Failed! Case: %s", tt.name)
				}else{
				   t.Logf("Success! Case: %s", tt.name)
				}
			} else {
				if got == graphql.Null {
					t.Errorf("___Directive_description() returned unexpected error, got = %v", got)
					t.Logf("Failed! Case: %s", tt.name)
				}else{
				   t.Logf("Success! Case: %s", tt.name)
				}
			}
		})
	}
}
