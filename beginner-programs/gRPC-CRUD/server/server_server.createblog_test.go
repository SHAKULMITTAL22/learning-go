```go
package server

import (
	"context"
	"fmt"
	"runtime/debug"
	"testing"
	// "time" // Import time if needed for delays in mocks, not strictly required for these scenarios

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tannergabriel/learning-go/beginner-programs/gRPC-CRUD/pb" // Assumes pb types are needed and accessible
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	// Removed unused imports: log, net, os, os/signal, bson, grpc, reflection
)

// --- Mocking Setup ---

// Define an interface matching the methods used from mongo.Collection by CreateBlog.
// This allows mocking the dependency without relying on the concrete mongo.Collection type directly.
type mongoCollectionInterface interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
}

// Mock implementation for the mongoCollectionInterface.
type mockMongoCollection struct {
	mock.Mock
}

// Implement the InsertOne method for the mock.
func (m *mockMongoCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	// Pass opts as a single slice argument to mock.Called for matching flexibility.
	// This allows using mock.Anything or specific matchers for the options slice.
	args := m.Called(ctx, document, opts)

	// Handle potential nil return for the result pointer.
	res, _ := args.Get(0).(*mongo.InsertOneResult)
	return res, args.Error(1)
}

// --- Struct Definitions (Locally defined if unexported in original package) ---

// Define necessary structs locally if they are unexported in the original 'server' package.
// This assumes the structure matches the one used internally by the CreateBlog function.
// If these structs are exported in the original package, these local definitions are not needed
// and could cause conflicts if this test file is in the same package.
// TODO: Verify if 'server' and 'BlogItem' are exported or need local definition.
// Commenting out local definitions assuming they exist in the package. If not, uncomment.
/*
type server struct {
	// Add fields here if the original server struct has any.
	// pb.UnimplementedBlogServiceServer // Likely needed if implementing a gRPC service
}

type BlogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}
*/

// --- Global Variable Handling for Mocking ---

// Global variable dependency: Use the interface type for the package-level 'Collection'.
// This assumes the original 'Collection' variable in the 'server' package
// is compatible with this interface for testing purposes (e.g., it's also an interface
// or can be replaced). Using global variables for dependencies makes testing harder;
// dependency injection into the 'server' struct is generally preferred.
// TODO: Confirm how the actual 'Collection' is managed and if this replacement strategy is valid in the target project.
var Collection mongoCollectionInterface

// Store the original value if Collection might be pre-initialized or used by other tests.
var originalCollection mongoCollectionInterface // Needs correct type matching the actual global variable

// --- Test Function ---

func TestServerCreateBlog(t *testing.T) {
	// Backup and restore the global Collection variable around the test suite.
	// This is crucial if 'Collection' is truly a global variable shared across tests or package state.
	originalCollection = Collection // Backup before tests run
	defer func() {
		Collection = originalCollection // Restore after all tests in this function complete
	}()

	// Define test cases using table-driven approach
	testCases := []struct {
		name                 string
		setupMock            func(mockColl *mockMongoCollection, testID primitive.ObjectID) // Pass testID for consistent ID checking
		request              *pb.CreateBlogRequest
		expectErr            bool
		expectedErrCode      codes.Code
		expectedErrMsgContains string // Substring to check in error message
		expectedID           primitive.ObjectID // Store the ID used in mock setup for assertion
	}{
		{
			name: "Scenario 1: Successful Blog Creation",
			request: &pb.CreateBlogRequest{
				Blog: &pb.Blog{
					AuthorId: "author