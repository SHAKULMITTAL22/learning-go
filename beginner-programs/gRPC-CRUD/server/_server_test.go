package server

import (
	context "context"
	fmt "fmt"
	debug "runtime/debug"
	testing "testing"
	time "time"
	pb "github.com/tannergabriel/learning-go/beginner-programs/gRPC-CRUD/pb"
	bson "go.mongodb.org/mongo-driver/bson"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	mongo "go.mongodb.org/mongo-driver/mongo"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	assert "github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
	options "go.mongodb.org/mongo-driver/mongo/options"
	errors "errors"
)



var Collection *MockCollection

type MockCollection struct {
	mock.Mock
}


/*
ROOST_METHOD_HASH=server_DeleteBlog_a6edcf3dd6
ROOST_METHOD_SIG_HASH=server_DeleteBlog_4e19addbc8

FUNCTION_DEF=func (*server) DeleteBlog(ctx context.Context, request *pb.DeleteBlogRequest) (*pb.DeleteBlogResponse, error) 

*/
func (m *mockMongoCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*mongo.options.DeleteOptions) (*mongo.DeleteResult, error) {
	m.DeleteOneCalled = true
	m.DeleteOneCtx = ctx
	m.DeleteOneFilter = filter
	if m.DeleteOneFunc != nil {

		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		return m.DeleteOneFunc(ctx, filter, opts...)
	}

	return nil, fmt.Errorf("DeleteOneFunc not implemented in mock")
}


/*
ROOST_METHOD_HASH=server_CreateBlog_3046ec9527
ROOST_METHOD_SIG_HASH=server_CreateBlog_48672e73a7

FUNCTION_DEF=func (*server) CreateBlog(ctx context.Context, request *pb.CreateBlogRequest) (*pb.CreateBlogResponse, error) 

*/
func (m *MockCollection) InsertOne(ctx context.Context, document interface{}, opts ...*mongo.InsertOneOptions) (*mongo.InsertOneResult, error) {

	if err := ctx.Err(); err != nil {

		return nil, err
	}
	args := m.Called(ctx, document)
	res := args.Get(0)
	err := args.Error(1)

	if res == nil {
		return nil, err
	}
	return res.(*mongo.InsertOneResult), err
}

func TestServerCreateBlog(t *testing.T) {

	s := &server{}

	testCases := []struct {
		name                 string
		ctx                  context.Context
		req                  *pb.CreateBlogRequest
		setupMock            func(mockColl *MockCollection, req *pb.CreateBlogRequest)
		expectedResp         *pb.CreateBlogResponse
		expectedErrCode      codes.Code
		expectedErrMsgSubstr string
		validation           string
		importance           string
	}{
		{
			name: "Scenario 1: Successful Blog Creation",
			ctx:  context.Background(),
			req: &pb.CreateBlogRequest{
				Blog: &pb.Blog{
					AuthorId: "author123",
					Title:    "My First Blog",
					Content:  "This is the content.",
				},
			},
			setupMock: func(mockColl *MockCollection, req *pb.CreateBlogRequest) {
				expectedObjectID := primitive.NewObjectID()
				expectedBlogItem := BlogItem{
					AuthorID: req.GetBlog().GetAuthorId(),
					Title:    req.GetBlog().GetTitle(),
					Content:  req.GetBlog().GetContent(),
				}
				mockColl.On("InsertOne", mock.AnythingOfType("*context.emptyCtx"), expectedBlogItem).
					Return(&mongo.InsertOneResult{InsertedID: expectedObjectID}, nil).
					Once()
			},
			expectedResp: &pb.CreateBlogResponse{
				Blog: &pb.Blog{
					Id: primitive.NewObjectID().Hex(),
				},
			},
			expectedErrCode: codes.OK,
			validation:      "Assert nil error and correct response structure with matching ID. Confirms happy path and data integrity.",
			importance:      "Fundamental test for core functionality under normal conditions.",
		},
		{
			name: "Scenario 2: Database Insertion Failure",
			ctx:  context.Background(),
			req: &pb.CreateBlogRequest{
				Blog: &pb.Blog{
					AuthorId: "author456",
					Title:    "DB Error Test",
					Content:  "Content here.",
				},
			},
			setupMock: func(mockColl *MockCollection, req *pb.CreateBlogRequest) {
				dbError := fmt.Errorf("database connection refused")
				expectedBlogItem := BlogItem{
					AuthorID: req.GetBlog().GetAuthorId(),
					Title:    req.GetBlog().GetTitle(),
					Content:  req.GetBlog().GetContent(),
				}
				mockColl.On("InsertOne", mock.AnythingOfType("*context.emptyCtx"), expectedBlogItem).
					Return(nil, dbError).
					Once()
			},
			expectedResp:         nil,
			expectedErrCode:      codes.Internal,
			expectedErrMsgSubstr: "internal error: database connection refused",
			validation:           "Expect nil response, non-nil error, gRPC Internal code, and specific error message propagation.",
			importance:           "Ensures database failures are handled gracefully and reported correctly.",
		},
		{
			name: "Scenario 3: Invalid Inserted ID Type from Database",
			ctx:  context.Background(),
			req: &pb.CreateBlogRequest{
				Blog: &pb.Blog{
					AuthorId: "author789",
					Title:    "Invalid OID Test",
					Content:  "More content.",
				},
			},
			setupMock: func(mockColl *MockCollection, req *pb.CreateBlogRequest) {
				invalidID := "not-an-object-id"
				expectedBlogItem := BlogItem{
					AuthorID: req.GetBlog().GetAuthorId(),
					Title:    req.GetBlog().GetTitle(),
					Content:  req.GetBlog().GetContent(),
				}
				mockColl.On("InsertOne", mock.AnythingOfType("*context.emptyCtx"), expectedBlogItem).
					Return(&mongo.InsertOneResult{InsertedID: invalidID}, nil).
					Once()
			},
			expectedResp:         nil,
			expectedErrCode:      codes.Internal,
			expectedErrMsgSubstr: "cannot convert OID",
			validation:           "Expect nil response, non-nil error, gRPC Internal code, and the specific 'cannot convert OID' message.",
			importance:           "Ensures resilience against unexpected database return types or mock errors.",
		},
		{
			name: "Scenario 4: Context Cancellation During Database Insertion",
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
			}(),
			req: &pb.CreateBlogRequest{
				Blog: &pb.Blog{
					AuthorId: "authorCancel",
					Title:    "Context Cancel Test",
					Content:  "Cancelled content.",
				},
			},
			setupMock: func(mockColl *MockCollection, req *pb.CreateBlogRequest) {

				expectedBlogItem := BlogItem{
					AuthorID: req.GetBlog().GetAuthorId(),
					Title:    req.GetBlog().GetTitle(),
					Content:  req.GetBlog().GetContent(),
				}

				mockColl.On("InsertOne", mock.AnythingOfType("*context.cancelCtx"), expectedBlogItem).
					Return(nil, context.Canceled).
					Once()

			},
			expectedResp:         nil,
			expectedErrCode:      codes.Internal,
			expectedErrMsgSubstr: "internal error: context canceled",
			validation:           "Expect nil response, non-nil error, gRPC Internal code, and error message containing 'context canceled'.",
			importance:           "Verifies correct handling of request lifecycle cancellation/timeouts.",
		},
		{
			name: "Scenario 4b: Context Timeout During Database Insertion",
			ctx: func() context.Context {

				ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)

				_ = cancel
				time.Sleep(5 * time.Millisecond)
				return ctx
			}(),
			req: &pb.CreateBlogRequest{
				Blog: &pb.Blog{
					AuthorId: "authorTimeout",
					Title:    "Context Timeout Test",
					Content:  "Timeout content.",
				},
			},
			setupMock: func(mockColl *MockCollection, req *pb.CreateBlogRequest) {
				expectedBlogItem := BlogItem{
					AuthorID: req.GetBlog().GetAuthorId(),
					Title:    req.GetBlog().GetTitle(),
					Content:  req.GetBlog().GetContent(),
				}

				mockColl.On("InsertOne", mock.AnythingOfType("*context.timerCtx"), expectedBlogItem).
					Return(nil, context.DeadlineExceeded).
					Once()
			},
			expectedResp:         nil,
			expectedErrCode:      codes.Internal,
			expectedErrMsgSubstr: "internal error: context deadline exceeded",
			validation:           "Expect nil response, non-nil error, gRPC Internal code, and error message containing 'context deadline exceeded'.",
			importance:           "Verifies correct handling of request lifecycle cancellation/timeouts.",
		},
		{
			name: "Scenario 5: Handling Request with Empty Fields",
			ctx:  context.Background(),
			req: &pb.CreateBlogRequest{
				Blog: &pb.Blog{
					AuthorId: "",
					Title:    "Empty Fields Test",
					Content:  "",
				},
			},
			setupMock: func(mockColl *MockCollection, req *pb.CreateBlogRequest) {
				expectedObjectID := primitive.NewObjectID()

				expectedBlogItem := BlogItem{
					AuthorID: "",
					Title:    "Empty Fields Test",
					Content:  "",
				}
				mockColl.On("InsertOne", mock.AnythingOfType("*context.emptyCtx"), expectedBlogItem).
					Return(&mongo.InsertOneResult{InsertedID: expectedObjectID}, nil).
					Once()
			},
			expectedResp: &pb.CreateBlogResponse{
				Blog: &pb.Blog{
					Id: primitive.NewObjectID().Hex(),
				},
			},
			expectedErrCode: codes.OK,
			validation:      "Assert nil error and successful response. Verifies function passes empty fields to DB without validation.",
			importance:      "Clarifies that input validation is not performed at this layer; relies on DB constraints or client validation.",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Test panicked: %v\n%s", r, string(debug.Stack()))
				}
			}()

			t.Logf("Running Test Case: %s", tc.name)
			t.Logf("Scenario Description: %s", tc.validation)
			t.Logf("Scenario Importance: %s", tc.importance)

			mockColl := new(MockCollection)

			Collection = mockColl

			var expectedIDHex string
			if tc.setupMock != nil {
				tc.setupMock(mockColl, tc.req)

				if tc.expectedErrCode == codes.OK {
					for _, call := range mockColl.ExpectedCalls {
						if call.Method == "InsertOne" {

							if len(call.ReturnArguments) > 0 && call.ReturnArguments.Get(0) != nil {
								if result, ok := call.ReturnArguments.Get(0).(*mongo.InsertOneResult); ok {
									if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
										expectedIDHex = oid.Hex()

										tc.expectedResp.Blog.Id = expectedIDHex
									}
								}
							}
							break
						}
					}
					if expectedIDHex == "" {
						t.Log("Warning: Could not extract expected ObjectID from mock setup for success case. ID assertion might fail.")

						if tc.expectedResp != nil && tc.expectedResp.Blog != nil {
							expectedIDHex = tc.expectedResp.Blog.Id
						}
					}
				}
			}

			resp, err := s.CreateBlog(tc.ctx, tc.req)

			if tc.expectedErrCode == codes.OK {

				if assert.NoError(t, err, "Expected no error, but got one") {
					assert.NotNil(t, resp, "Response should not be nil on success")
					if resp != nil {
						assert.NotNil(t, resp.Blog, "Response.Blog should not be nil on success")
						if resp.Blog != nil {
							assert.Equal(t, expectedIDHex, resp.Blog.Id, "Response Blog ID does not match expected ID")
							t.Logf("Success: Received expected response with ID: %s", resp.Blog.Id)
						}
					}
				} else {
					st, _ := status.FromError(err)
					t.Logf("Failure: Expected success but got error: Code=%s, Message=%s", st.Code(), st.Message())
				}
			} else {

				if assert.Error(t, err, "Expected an error, but got nil") {
					st, ok := status.FromError(err)
					if assert.True(t, ok, "Error should be a gRPC status error") {
						assert.Equal(t, tc.expectedErrCode, st.Code(), "gRPC error code does not match expected")
						if tc.expectedErrMsgSubstr != "" {
							assert.Contains(t, st.Message(), tc.expectedErrMsgSubstr, "Error message does not contain expected substring")
							t.Logf("Success: Received expected error: Code=%s, Message contains '%s'", st.Code(), tc.expectedErrMsgSubstr)
						} else {
							t.Logf("Success: Received expected error: Code=%s, Message=%s", st.Code(), st.Message())
						}
					}
				}
				assert.Nil(t, resp, "Response should be nil on error")
			}

			mockColl.AssertExpectations(t)
			t.Logf("Finished Test Case: %s", tc.name)
			t.Log("------")
		})
	}
}


/*
ROOST_METHOD_HASH=server_ReadBlog_4959f7714c
ROOST_METHOD_SIG_HASH=server_ReadBlog_49fcec6ebb

FUNCTION_DEF=func (*server) ReadBlog(ctx context.Context, request *pb.ReadBlogRequest) (*pb.ReadBlogResponse, error) 

*/
func (m *mockMongoCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	if m.findOneMock != nil {
		return m.findOneMock(ctx, filter, opts...)
	}

	panic("mockMongoCollection.FindOne called without findOneMock configured")
}


/*
ROOST_METHOD_HASH=server_UpdateBlog_449f48f642
ROOST_METHOD_SIG_HASH=server_UpdateBlog_3ebfee2c2e

FUNCTION_DEF=func (*server) UpdateBlog(ctx context.Context, request *pb.UpdateBlogRequest) (*pb.UpdateBlogResponse, error) 

*/
func (m *mockCollection) FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
	if m.FindOneAndUpdateFunc != nil {
		return m.FindOneAndUpdateFunc(ctx, filter, update, opts...)
	}

	return mongo.NewSingleResultFromDocument(nil, errors.New("FindOneAndUpdateFunc not implemented in mock"), nil)
}

func (m *mockSingleResult) Decode(v interface{}) error {

	return m.err
}

func (m *mockSingleResult) DecodeBytes() (bson.Raw, error) {
	return nil, m.err
}

func (m *mockSingleResult) Err() error {
	return m.err
}

