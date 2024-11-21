// Test generated by RoostGPT for test roost-test using AI Type Azure Open AI and AI Model roost-gpt4-32k

package middleware

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// mockMongoCollection is a helper struct that we can utilise to mock the behaviour of mgo.collection struct
type mockMongoCollection struct {
	mock.Mock
}

func (m *mockMongoCollection) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(*mongo.DeleteResult), args.Error(1)
}

func Test_deleteAllTask(t *testing.T) {
	cases := []struct {
		name                 string
		result               mongo.DeleteResult
		error                error
		expectedDeletedCount int64
		expectedError        error
	}{
		{
			name:                 "Successful Deletion of All Tasks",
			result:               mongo.DeleteResult{DeletedCount: 10},
			error:                nil,
			expectedDeletedCount: 10,
			expectedError:        nil,
		},
		{
			name:                 "Deletion of Tasks from Empty Collection",
			result:               mongo.DeleteResult{DeletedCount: 0},
			error:                nil,
			expectedDeletedCount: 0,
			expectedError:        nil,
		},
		{
			name:                 "Fails to delete",
			result:               mongo.DeleteResult{DeletedCount: 0},
			error:                mongo.ErrInvalidCursor,
			expectedDeletedCount: 0,
			expectedError:        mongo.ErrInvalidCursor,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockCollection := new(mockMongoCollection)
			mockCollection.On("DeleteMany", ctx, bson.D{{}}).Return(&tc.result, tc.error)
			collection = mockCollection
			osStdOut := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			deletedCount := deleteAllTask()
			fmt.Fprintln(os.Stdout)
			w.Close()
			os.Stdout = osStdOut
			var buf bytes.Buffer
			io.Copy(&buf, r)
			assert.Contains(t, buf.String(), fmt.Sprint(tc.expectedDeletedCount))
			assert.Equal(t, tc.expectedDeletedCount, deletedCount)
		})
	}
}
