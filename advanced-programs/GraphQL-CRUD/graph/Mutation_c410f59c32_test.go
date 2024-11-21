// Expected package line (add the relevant package name here)
package your_package_name_here

// Importing required packages for the test cases
import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/database"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/generated"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
)

type mutationResolver struct{ *Resolver }

// Testing function for method 'Mutation'
func TestMutation(t *testing.T) {
	resolver := &Resolver{Db: database.SetupDatabase()}

	testCases := []struct {
		desc   string
		ctx    context.Context
		output generated.MutationResolver
		err    bool
	}{
		{
			desc:   "Test Mutation Success",                         // Successful mutation test case
			ctx:    context.Background(),
			output: &mutationResolver{resolver},
			err:    false,
		},
		// TODO: Add more test cases here for different scenarios
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output := resolver.Mutation()

			assert.Equal(t, tC.output, output, "Expected output not matched")

			// If the test case expects an error
			if tC.err {
				// If the error returned is nil, fail the test case
				if output == nil {
					t.Fail()
					t.Log("Failed test case: ", tC.desc, " Expected error but got nil")
				} else {
					t.Log("Passed test case: ", tC.desc)
				}
			} else {
				// If the error returned is not nil, fail the test case
				if output != tC.output {
					t.Fail()
					t.Log("Failed test case: ", tC.desc, " Expected no error but got error")
				} else {
					t.Log("Passed test case: ", tC.desc)
				}
			}
		})
	}
}
