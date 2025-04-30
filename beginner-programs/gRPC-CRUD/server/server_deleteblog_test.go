package main

import (
	"context"
	"testing"
	"reflect"
	"github.com/tannergabriel/learning-go/beginner-programs/gRPC-CRUD/pb"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/DATA-DOG/go-sqlmock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MockedCollection struct {
	Err error
}

type ExpectedExec struct {
	queryBasedExpectation
	result driver.Result
	delay  time.Duration
}




type DeleteResult struct {
	DeletedCount int64 `bson:"n"` // The number of documents deleted.
}


type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func (mc *MockedCollection) DeleteOne(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	return &mongo.DeleteResult{}, mc.Err
}
func TestserverDeleteBlog(t *testing.T) {

	testCases := []struct {
		name    string
		blogId  string
		want    *pb.DeleteBlogResponse
		wantErr bool
		err     error
	}{
		{
			name:    "Successful Deletion of a Blog",
			blogId:  "5f697986e1822b422c84f8f9",
			want:    &pb.DeleteBlogResponse{Status: true},
			wantErr: false,
			err:     nil,
		},
		{
			name:    "Deletion of a Blog with Invalid ID",
			blogId:  "123",
			want:    nil,
			wantErr: true,
			err:     status.Errorf(codes.InvalidArgument, "Could not convert to ObjectId: invalid ObjectId in hex: 123"),
		},
		{
			name:    "Deletion of a Non-Existing Blog",
			blogId:  "5f697986e1822b422c84f8ff",
			want:    nil,
			wantErr: true,
			err:     status.Errorf(codes.NotFound, "Could not find/delete blog with id 5f697986e1822b422c84f8ff: mongo: no documents in result"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mock.ExpectExec("DELETE FROM blogs WHERE id = ?").
				WithArgs(tc.blogId).
				WillReturnResult(sqlmock.NewResult(1, 1))

			Collection = &MockedCollection{Err: tc.err}

			request := &pb.DeleteBlogRequest{BlogId: tc.blogId}

			s := server{}
			got, err := s.DeleteBlog(context.Background(), request)

			if (err != nil) != tc.wantErr {
				t.Errorf("server.DeleteBlog() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if err != nil && err.Error() != tc.err.Error() {
				t.Errorf("server.DeleteBlog() = %v, want %v", err, tc.err)
				return
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("server.DeleteBlog() = %v, want %v", got, tc.want)
			}
		})
	}
}
