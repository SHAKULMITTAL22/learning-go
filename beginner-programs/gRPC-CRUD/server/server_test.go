package main

import (
	"context"
	"fmt"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/tannergabriel/learning-go/beginner-programs/gRPC-CRUD/pb"
	"errors"
	"time"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/options"
)



var mockCollection *mongo.Collection
var mockSingleResult *mongo.SingleResult

type ExpectedExec struct {
	queryBasedExpectation
	result driver.Result
	delay  time.Duration
}
type T struct {
	common
	isEnvSet bool
	context  *testContext
}


/*
ROOST_METHOD_HASH=DeleteBlog_ec0a92d5fe
ROOST_METHOD_SIG_HASH=DeleteBlog_4e19addbc8


 */
func TestserverDeleteBlog(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer db.Close()

	tests := []struct {
		name       string
		blogId     string
		setupMock  func()
		wantStatus bool
		wantErr    codes.Code
	}{
		{
			name:   "Valid Blog Deletion",
			blogId: "60c72b2f9b1e4d3f8c2f1234",
			setupMock: func() {
				mock.ExpectExec("DELETE FROM blogs WHERE (.+)").
					WithArgs("60c72b2f9b1e4d3f8c2f1234").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantStatus: true,
			wantErr:    codes.OK,
		},
		{
			name:   "Invalid Blog ID Format",
			blogId: "invalid_id",
			setupMock: func() {

			},
			wantStatus: false,
			wantErr:    codes.InvalidArgument,
		},
		{
			name:   "Non-Existent Blog ID",
			blogId: "60c72b2f9b1e4d3f8c2f5678",
			setupMock: func() {
				mock.ExpectExec("DELETE FROM blogs WHERE (.+)").
					WithArgs("60c72b2f9b1e4d3f8c2f5678").
					WillReturnError(mongo.ErrNoDocuments)
			},
			wantStatus: false,
			wantErr:    codes.NotFound,
		},
		{
			name:   "Database Connection Failure",
			blogId: "60c72b2f9b1e4d3f8c2f1234",
			setupMock: func() {
				mock.ExpectExec("DELETE FROM blogs WHERE (.+)").
					WithArgs("60c72b2f9b1e4d3f8c2f1234").
					WillReturnError(fmt.Errorf("connection error"))
			},
			wantStatus: false,
			wantErr:    codes.Internal,
		},
		{
			name:   "Deleting Already Deleted Blog",
			blogId: "60c72b2f9b1e4d3f8c2f1234",
			setupMock: func() {
				mock.ExpectExec("DELETE FROM blogs WHERE (.+)").
					WithArgs("60c72b2f9b1e4d3f8c2f1234").
					WillReturnError(mongo.ErrNoDocuments)
			},
			wantStatus: false,
			wantErr:    codes.NotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			s := &server{}
			req := &pb.DeleteBlogRequest{BlogId: tt.blogId}
			resp, err := s.DeleteBlog(context.Background(), req)

			if tt.wantErr != codes.OK {
				if err == nil {
					t.Fatalf("expected error code %v, got nil", tt.wantErr)
				}
				st, ok := status.FromError(err)
				if !ok || st.Code() != tt.wantErr {
					t.Fatalf("expected error code %v, got %v", tt.wantErr, st.Code())
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if resp.Status != tt.wantStatus {
					t.Fatalf("expected status %v, got %v", tt.wantStatus, resp.Status)
				}
			}
		})
	}
}


/*
ROOST_METHOD_HASH=CreateBlog_f35cab12bb
ROOST_METHOD_SIG_HASH=CreateBlog_48672e73a7


 */
func TestserverCreateBlog(t *testing.T) {
	type testCase struct {
		name         string
		request      *pb.CreateBlogRequest
		mockBehavior func(mock sqlmock.Sqlmock)
		expectedResp *pb.CreateBlogResponse
		expectedErr  error
	}

	validBlog := &pb.Blog{
		AuthorId: "author123",
		Title:    "Test Title",
		Content:  "Test Content",
	}

	tests := []testCase{
		{
			name: "Successful Blog Creation",
			request: &pb.CreateBlogRequest{
				Blog: validBlog,
			},
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO blogs").WithArgs(validBlog.AuthorId, validBlog.Title, validBlog.Content).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedResp: &pb.CreateBlogResponse{
				Blog: &pb.Blog{
					Id: "507f191e810c19729de860ea",
				},
			},
			expectedErr: nil,
		},
		{
			name: "Database Insertion Failure",
			request: &pb.CreateBlogRequest{
				Blog: validBlog,
			},
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO blogs").WithArgs(validBlog.AuthorId, validBlog.Title, validBlog.Content).
					WillReturnError(errors.New("failed to insert"))
			},
			expectedResp: nil,
			expectedErr:  status.Errorf(codes.Internal, "internal error: failed to insert"),
		},
		{
			name: "Invalid ObjectID Conversion",
			request: &pb.CreateBlogRequest{
				Blog: validBlog,
			},
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO blogs").WithArgs(validBlog.AuthorId, validBlog.Title, validBlog.Content).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedResp: nil,
			expectedErr:  status.Errorf(codes.Internal, "cannot convert OID"),
		},
		{
			name: "Missing Blog Data in Request",
			request: &pb.CreateBlogRequest{
				Blog: nil,
			},
			mockBehavior: func(mock sqlmock.Sqlmock) {

			},
			expectedResp: nil,
			expectedErr:  status.Errorf(codes.InvalidArgument, "missing blog data"),
		},
		{
			name: "Empty Blog Fields",
			request: &pb.CreateBlogRequest{
				Blog: &pb.Blog{
					AuthorId: "",
					Title:    "",
					Content:  "",
				},
			},
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO blogs").WithArgs("", "", "").
					WillReturnError(errors.New("failed to insert"))
			},
			expectedResp: nil,
			expectedErr:  status.Errorf(codes.InvalidArgument, "empty blog fields"),
		},
		{
			name: "Context Cancellation",
			request: &pb.CreateBlogRequest{
				Blog: validBlog,
			},
			mockBehavior: func(mock sqlmock.Sqlmock) {

			},
			expectedResp: nil,
			expectedErr:  status.Errorf(codes.Canceled, "context canceled"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockDB, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer mockDB.Close()

			tt.mockBehavior(mock)

			Collection = &mongo.Collection{}

			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			if tt.name == "Context Cancellation" {
				cancel()
			}

			resp, err := new(server).CreateBlog(ctx, tt.request)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResp.Blog.Id, resp.Blog.Id)
			}

			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}


/*
ROOST_METHOD_HASH=ReadBlog_e7690a63b9
ROOST_METHOD_SIG_HASH=ReadBlog_49fcec6ebb


 */
func TestserverReadBlog(t *testing.T) {

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	tests := []struct {
		name         string
		blogID       string
		mockSetup    func()
		expectedResp *pb.ReadBlogResponse
		expectedErr  error
	}{
		{
			name:   "Valid Blog ID",
			blogID: "60d5ec49f8d2e3f1b2f4b7a6",
			mockSetup: func() {
				oid, _ := primitive.ObjectIDFromHex("60d5ec49f8d2e3f1b2f4b7a6")
				mockSingleResult = &mongo.SingleResult{

					rdr: bson.Raw{
						0x0a, 0x00, 0x00, 0x00,
						0x02, 'a', 'u', 't', 'h', 'o', 'r', 'I', 'd', 0x00, 0x0a, 0x00, 0x00, 0x00, 'A', 'u', 't', 'h', 'o', 'r', '1', '2', '3', 0x00,
						0x02, 't', 'i', 't', 'l', 'e', 0x00, 0x0d, 0x00, 0x00, 0x00, 'S', 'a', 'm', 'p', 'l', 'e', ' ', 'T', 'i', 't', 'l', 'e', 0x00,
						0x02, 'c', 'o', 'n', 't', 'e', 'n', 't', 0x00, 0x0f, 0x00, 0x00, 0x00, 'S', 'a', 'm', 'p', 'l', 'e', ' ', 'C', 'o', 'n', 't', 'e', 'n', 't', 0x00,
						0x00,
					},
				}

				Collection = mockCollection

				mockCollection.FindOne = func(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
					return mockSingleResult
				}
			},
			expectedResp: &pb.ReadBlogResponse{
				Blog: &pb.Blog{
					Id:       "60d5ec49f8d2e3f1b2f4b7a6",
					AuthorId: "Author123",
					Title:    "Sample Title",
					Content:  "Sample Content",
				},
			},
			expectedErr: nil,
		},
		{
			name:   "Invalid Blog ID Format",
			blogID: "invalidID",
			mockSetup: func() {

			},
			expectedResp: nil,
			expectedErr:  status.Errorf(codes.InvalidArgument, "cannot parse ID"),
		},
		{
			name:   "Nonexistent Blog ID",
			blogID: "60d5ec49f8d2e3f1b2f4b7a7",
			mockSetup: func() {
				oid, _ := primitive.ObjectIDFromHex("60d5ec49f8d2e3f1b2f4b7a7")
				mockSingleResult = &mongo.SingleResult{
					err: mongo.ErrNoDocuments,
				}
				Collection = mockCollection
				mockCollection.FindOne = func(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
					return mockSingleResult
				}
			},
			expectedResp: nil,
			expectedErr:  status.Errorf(codes.NotFound, "document not found"),
		},
		{
			name:   "Database Connection Error",
			blogID: "60d5ec49f8d2e3f1b2f4b7a8",
			mockSetup: func() {
				mockSingleResult = &mongo.SingleResult{
					err: fmt.Errorf("database connection error"),
				}
				Collection = mockCollection
				mockCollection.FindOne = func(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
					return mockSingleResult
				}
			},
			expectedResp: nil,
			expectedErr:  status.Errorf(codes.Internal, "database connection error"),
		},
		{
			name:   "Blog ID as Empty String",
			blogID: "",
			mockSetup: func() {

			},
			expectedResp: nil,
			expectedErr:  status.Errorf(codes.InvalidArgument, "cannot parse ID"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.mockSetup()

			s := &server{}
			req := &pb.ReadBlogRequest{BlogId: tt.blogID}

			resp, err := s.ReadBlog(context.Background(), req)

			if tt.expectedErr != nil {
				assert.Nil(t, resp)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NotNil(t, resp)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResp, resp)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=UpdateBlog_0c13632208
ROOST_METHOD_SIG_HASH=UpdateBlog_3ebfee2c2e


 */
func TestserverUpdateBlog(t *testing.T) {

	mockDB, _, _ := sqlmock.New()
	defer mockDB.Close()

	collection := &mongo.Collection{}
	Collection = collection

	tests := []struct {
		name           string
		request        *pb.UpdateBlogRequest
		mockSetup      func()
		expectedError  codes.Code
		expectedResult *pb.UpdateBlogResponse
	}{
		{
			name: "Successful Update of an Existing Blog",
			request: &pb.UpdateBlogRequest{
				BlogId: "60b8d295f1d2b2a5a4c8f0e1",
				Blog: &pb.Blog{
					AuthorId: "12345",
					Title:    "Updated Title",
					Content:  "Updated Content",
				},
			},
			mockSetup: func() {

			},
			expectedError: codes.OK,
			expectedResult: &pb.UpdateBlogResponse{
				Blog: &pb.Blog{
					Id:       "60b8d295f1d2b2a5a4c8f0e1",
					AuthorId: "12345",
					Title:    "Updated Title",
					Content:  "Updated Content",
				},
			},
		},
		{
			name: "Blog ID Not Found",
			request: &pb.UpdateBlogRequest{
				BlogId: "nonexistentid",
				Blog: &pb.Blog{
					AuthorId: "12345",
					Title:    "Title",
					Content:  "Content",
				},
			},
			mockSetup: func() {

			},
			expectedError:  codes.NotFound,
			expectedResult: nil,
		},
		{
			name: "Invalid Blog ID Format",
			request: &pb.UpdateBlogRequest{
				BlogId: "invalidformat",
				Blog: &pb.Blog{
					AuthorId: "12345",
					Title:    "Title",
					Content:  "Content",
				},
			},
			mockSetup: func() {

			},
			expectedError:  codes.InvalidArgument,
			expectedResult: nil,
		},
		{
			name: "Database Update Failure",
			request: &pb.UpdateBlogRequest{
				BlogId: "60b8d295f1d2b2a5a4c8f0e1",
				Blog: &pb.Blog{
					AuthorId: "12345",
					Title:    "Title",
					Content:  "Content",
				},
			},
			mockSetup: func() {

			},
			expectedError:  codes.Internal,
			expectedResult: nil,
		},
		{
			name: "No Changes in Update Request",
			request: &pb.UpdateBlogRequest{
				BlogId: "60b8d295f1d2b2a5a4c8f0e1",
				Blog: &pb.Blog{
					AuthorId: "12345",
					Title:    "Same Title",
					Content:  "Same Content",
				},
			},
			mockSetup: func() {

			},
			expectedError: codes.OK,
			expectedResult: &pb.UpdateBlogResponse{
				Blog: &pb.Blog{
					Id:       "60b8d295f1d2b2a5a4c8f0e1",
					AuthorId: "12345",
					Title:    "Same Title",
					Content:  "Same Content",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			test.mockSetup()

			s := &server{}
			result, err := s.UpdateBlog(context.Background(), test.request)

			if test.expectedError != codes.OK {
				if err == nil {
					t.Errorf("Expected error code %v, got nil", test.expectedError)
				} else if status.Code(err) != test.expectedError {
					t.Errorf("Expected error code %v, got %v", test.expectedError, status.Code(err))
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
				if result.Blog.Id != test.expectedResult.Blog.Id ||
					result.Blog.AuthorId != test.expectedResult.Blog.AuthorId ||
					result.Blog.Title != test.expectedResult.Blog.Title ||
					result.Blog.Content != test.expectedResult.Blog.Content {
					t.Errorf("Expected result %v, got %v", test.expectedResult, result)
				}
			}
		})
	}
}

