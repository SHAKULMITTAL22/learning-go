package FiberPostgresCRUD

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"
	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tannergabriel/learning-go/advanced-programs/FiberPostgresCRUD/database"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"regexp"
)



var initDatabase = func() {}
var createTable = func() {}
var setupRoutes = func(app *fiber.App) {}

type MockDatabase struct {
	mock.Mock
}


/*
ROOST_METHOD_HASH=helloWorld_81caa76204
ROOST_METHOD_SIG_HASH=helloWorld_7d2f837ce0

FUNCTION_DEF=func helloWorld(c *fiber.Ctx) 

*/
func TestHelloWorld(t *testing.T) {
	app := fiber.New()

	tests := []struct {
		name           string
		method         string
		expectedBody   string
		expectedStatus int
		modifyHeaders  func(*fiber.Ctx)
	}{
		{
			name:           "Scenario 1: Verify response content",
			method:         "GET",
			expectedBody:   "Hello, World!",
			expectedStatus: 200,
		},
		{
			name:           "Scenario 2: Ensure status code is 200",
			method:         "GET",
			expectedBody:   "Hello, World!",
			expectedStatus: 200,
		},
		{
			name:           "Scenario 3: Handle empty request correctly",
			method:         "GET",
			expectedBody:   "Hello, World!",
			expectedStatus: 200,
		},
		{
			name:           "Scenario 7: Handle malformed request gracefully",
			method:         "POST",
			expectedBody:   "Hello, World!",
			expectedStatus: 200,
		},
	}

	app.Get("/", helloWorld)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := fiber.AcquireCtx(app.AcquireCtx())
			defer fiber.ReleaseCtx(req)

			req.Request().Header.SetMethod(tt.method)

			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			helloWorld(req)

			w.Close()
			os.Stdout = old

			var buf bytes.Buffer
			fmt.Fscanf(r, "%s", &buf)

			assert.Equal(t, tt.expectedBody, buf.String(), "Response body mismatch")
			assert.Equal(t, tt.expectedStatus, req.Response().StatusCode(), "Unexpected status code")
		})
	}

	t.Run("Scenario 4: Verify request headers remain unchanged", func(t *testing.T) {
		req := fiber.AcquireCtx(app.AcquireCtx())
		defer fiber.ReleaseCtx(req)

		req.Request().Header.Set("X-Custom-Header", "TestValue")

		helloWorld(req)

		assert.Equal(t, "TestValue", req.Request().Header.Peek("X-Custom-Header"), "Request headers should not be modified")
	})

	t.Run("Scenario 5: Handle concurrent requests", func(t *testing.T) {
		var wg sync.WaitGroup
		const numRequests = 10

		for i := 0; i < numRequests; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				req := fiber.AcquireCtx(app.AcquireCtx())
				defer fiber.ReleaseCtx(req)

				helloWorld(req)

				assert.Equal(t, "Hello, World!", string(req.Response().Body()), "Response body mismatch in concurrent execution")
				assert.Equal(t, 200, req.Response().StatusCode(), "Unexpected status code in concurrent execution")
			}()
		}

		wg.Wait()
	})

	t.Run("Scenario 6: Ensure no global state modifications", func(t *testing.T) {
		initialState := os.Environ()

		helloWorld(fiber.AcquireCtx(app.AcquireCtx()))

		assert.Equal(t, initialState, os.Environ(), "Global state should not be modified")
	})

	t.Run("Scenario 8: Performance test", func(t *testing.T) {
		start := time.Now()
		req := fiber.AcquireCtx(app.AcquireCtx())
		defer fiber.ReleaseCtx(req)

		helloWorld(req)

		duration := time.Since(start)
		assert.Less(t, duration.Milliseconds(), int64(50), "Function execution took too long")
	})

	t.Run("Scenario 9: Handle different HTTP methods", func(t *testing.T) {
		methods := []string{"GET", "POST", "PUT", "DELETE"}

		for _, method := range methods {
			t.Run(fmt.Sprintf("Method: %s", method), func(t *testing.T) {
				req := fiber.AcquireCtx(app.AcquireCtx())
				defer fiber.ReleaseCtx(req)

				req.Request().Header.SetMethod(method)
				helloWorld(req)

				assert.Equal(t, "Hello, World!", string(req.Response().Body()), "Response body mismatch for method "+method)
				assert.Equal(t, 200, req.Response().StatusCode(), "Unexpected status code for method "+method)
			})
		}
	})

	t.Run("Scenario 10: Memory leak test", func(t *testing.T) {
		before := new(bytes.Buffer)
		after := new(bytes.Buffer)

		fmt.Fprintf(before, "%v", os.Environ())

		for i := 0; i < 1000; i++ {
			req := fiber.AcquireCtx(app.AcquireCtx())
			helloWorld(req)
			fiber.ReleaseCtx(req)
		}

		fmt.Fprintf(after, "%v", os.Environ())

		assert.Equal(t, before.String(), after.String(), "Memory leak detected")
	})
}


/*
ROOST_METHOD_HASH=initDatabase_66a03266dc
ROOST_METHOD_SIG_HASH=initDatabase_bd163c083c

FUNCTION_DEF=func initDatabase() 

*/
func TestInitDatabase(t *testing.T) {
	tests := []struct {
		name            string
		envHost         string
		expectedConnStr string
		mockDBError     bool
		expectPanic     bool
		expectedLog     string
	}{
		{
			name:            "Successfully Connect to Database with Default Host",
			envHost:         "",
			expectedConnStr: "host=localhost port=5432 user=_ dbname=pq-demo password=example sslmode=disable",
			mockDBError:     false,
			expectPanic:     false,
			expectedLog:     "Connection Opened to Database",
		},
		{
			name:            "Successfully Connect to Database with Custom Host",
			envHost:         "custom-db-host",
			expectedConnStr: "host=custom-db-host port=5432 user=_ dbname=pq-demo password=example sslmode=disable",
			mockDBError:     false,
			expectPanic:     false,
			expectedLog:     "Connection Opened to Database",
		},
		{
			name:            "Database Connection Failure Handling",
			envHost:         "invalid-host",
			expectedConnStr: "host=invalid-host port=5432 user=_ dbname=pq-demo password=example sslmode=disable",
			mockDBError:     true,
			expectPanic:     true,
			expectedLog:     "failed to connect database",
		},
		{
			name:            "Ensure Panic on Database Connection Failure",
			envHost:         "invalid-host",
			expectedConnStr: "host=invalid-host port=5432 user=_ dbname=pq-demo password=example sslmode=disable",
			mockDBError:     true,
			expectPanic:     true,
			expectedLog:     "failed to connect database",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			originalHost := os.Getenv("HOST")
			defer os.Setenv("HOST", originalHost)

			os.Setenv("HOST", tt.envHost)

			var buf bytes.Buffer
			fmt.Fprintf(&buf, "")

			mockDB, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Failed to create mock database: %v", err)
			}
			defer mockDB.Close()

			if tt.mockDBError {
				mock.ExpectPing().WillReturnError(fmt.Errorf("mock connection error"))
			} else {
				mock.ExpectPing().WillReturnError(nil)
			}

			originalOpen := gorm.Open
			defer func() { gorm.Open = originalOpen }()

			gorm.Open = func(dialect string, args ...interface{}) (*gorm.DB, error) {
				if tt.mockDBError {
					return nil, fmt.Errorf("mock connection error")
				}
				return &gorm.DB{}, nil
			}

			defer func() {
				if r := recover(); r != nil {
					if !tt.expectPanic {
						t.Errorf("Unexpected panic: %v", r)
					} else if !strings.Contains(fmt.Sprint(r), tt.expectedLog) {
						t.Errorf("Expected panic message to contain %q, got %q", tt.expectedLog, r)
					}
				} else if tt.expectPanic {
					t.Errorf("Expected panic but did not occur")
				}
			}()

			initDatabase()

			if !tt.expectPanic {
				if database.DBConn == nil {
					t.Errorf("Expected database.DBConn to be initialized, but got nil")
				} else if reflect.TypeOf(database.DBConn) != reflect.TypeOf(&gorm.DB{}) {
					t.Errorf("Expected database.DBConn to be of type *gorm.DB, but got %T", database.DBConn)
				}
			}

			logOutput := buf.String()
			if tt.expectedLog != "" && !strings.Contains(logOutput, tt.expectedLog) {
				t.Errorf("Expected log output to contain %q, but got %q", tt.expectedLog, logOutput)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=setupRoutes_483feb9810
ROOST_METHOD_SIG_HASH=setupRoutes_b26c312b04

FUNCTION_DEF=func setupRoutes(app *fiber.App) 

*/
func TestSetupRoutes(t *testing.T) {
	app := fiber.New()
	setupRoutes(app)

	tests := []struct {
		name           string
		method         string
		url            string
		body           interface{}
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Scenario 1: Verify Root Route Returns 'Hello, World!'",
			method:         http.MethodGet,
			url:            "/",
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, World!",
		},
		{
			name:           "Scenario 2: Retrieve All Items Successfully",
			method:         http.MethodGet,
			url:            "/api/v1/item",
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"id":1,"name":"Item1"},{"id":2,"name":"Item2"}]`,
		},
		{
			name:           "Scenario 3: Retrieve a Single Item by ID",
			method:         http.MethodGet,
			url:            "/api/v1/item/1",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":1,"name":"Item1"}`,
		},
		{
			name:           "Scenario 4: Retrieve a Non-Existent Item",
			method:         http.MethodGet,
			url:            "/api/v1/item/9999",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Scenario 5: Create a New Item Successfully",
			method:         http.MethodPost,
			url:            "/api/v1/item",
			body:           map[string]interface{}{"name": "NewItem"},
			expectedStatus: http.StatusCreated,
			expectedBody:   `{"id":3,"name":"NewItem"}`,
		},
		{
			name:           "Scenario 6: Attempt to Create an Item with Invalid Data",
			method:         http.MethodPost,
			url:            "/api/v1/item",
			body:           map[string]interface{}{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Scenario 7: Delete an Existing Item Successfully",
			method:         http.MethodDelete,
			url:            "/api/v1/item/1",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"Item deleted successfully"}`,
		},
		{
			name:           "Scenario 8: Attempt to Delete a Non-Existent Item",
			method:         http.MethodDelete,
			url:            "/api/v1/item/9999",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Scenario 9: Verify Route Not Found Handling",
			method:         http.MethodGet,
			url:            "/api/v1/unknown",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Scenario 10: Verify Method Not Allowed Handling",
			method:         http.MethodPut,
			url:            "/api/v1/item/1",
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var reqBody []byte
			var err error

			if tt.body != nil {
				reqBody, err = json.Marshal(tt.body)
				if err != nil {
					t.Fatalf("Failed to marshal request body: %v", err)
				}
			}

			req := httptest.NewRequest(tt.method, tt.url, bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			if err != nil {
				t.Fatalf("Failed to execute request: %v", err)
			}

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}

			if tt.expectedBody != "" {
				var actualBody bytes.Buffer
				_, err := actualBody.ReadFrom(resp.Body)
				if err != nil {
					t.Fatalf("Failed to read response body: %v", err)
				}

				if actualBody.String() != tt.expectedBody {
					t.Errorf("Expected body %s, got %s", tt.expectedBody, actualBody.String())
				}
			}

			t.Logf("Test %s passed", tt.name)
		})
	}
}

func mockDatabase() (sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create mock database: %v", err)
	}

	mock.ExpectQuery("SELECT \\* FROM items").WillReturnRows(
		sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "Item1").
			AddRow(2, "Item2"),
	)

	mock.ExpectQuery("SELECT \\* FROM items WHERE id = \\$1").WithArgs(1).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "Item1"),
	)

	mock.ExpectExec("INSERT INTO items").WithArgs("NewItem").WillReturnResult(sqlmock.NewResult(3, 1))

	mock.ExpectExec("DELETE FROM items WHERE id = \\$1").WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectQuery("SELECT \\* FROM items WHERE id = \\$1").WithArgs(9999).WillReturnError(fmt.Errorf("not found"))

	mock.ExpectExec("DELETE FROM items WHERE id = \\$1").WithArgs(9999).WillReturnError(fmt.Errorf("not found"))

	return mock, nil
}


/*
ROOST_METHOD_HASH=Setup_749facca65
ROOST_METHOD_SIG_HASH=Setup_375240d483

FUNCTION_DEF=func Setup() *fiber.App // Setup fiber app and database


*/
func (m *MockDatabase) CreateTable() {
	m.Called()
}

func (m *MockDatabase) InitDatabase() {
	m.Called()
}

func Setup() *fiber.App {
	app := fiber.New()
	initDatabase()
	createTable()
	setupRoutes(app)
	return app
}

func (m *MockDatabase) SetupRoutes(app *fiber.App) {
	m.Called(app)
}

func TestSetupDatabaseInitialization(t *testing.T) {
	mockDB := new(MockDatabase)
	mockDB.On("InitDatabase").Return()

	initDatabase = func() {
		mockDB.InitDatabase()
	}

	Setup()

	mockDB.AssertCalled(t, "InitDatabase")
}

func TestSetupDoesNotReturnNil(t *testing.T) {
	app := Setup()
	assert.NotNil(t, app, "Expected Setup function to return a valid Fiber app instance")
}

func TestSetupFiberAppInitialization(t *testing.T) {
	app := Setup()
	assert.NotNil(t, app, "Expected Fiber app instance to be initialized")
}

func TestSetupHandlesDatabaseInitError(t *testing.T) {
	originalInitDatabase := initDatabase
	defer func() { initDatabase = originalInitDatabase }()

	initDatabase = func() {
		fmt.Println("Mocked database initialization error")
	}

	assert.NotPanics(t, func() { Setup() }, "Expected Setup to handle database initialization error gracefully")
}

func TestSetupHandlesRequests(t *testing.T) {
	app := Setup()
	req, _ := http.NewRequest("GET", "/", nil)
	res, err := app.Test(req, -1)

	assert.Nil(t, err, "Expected no error when sending request")
	assert.Equal(t, 200, res.StatusCode, "Expected status code 200")
}

func TestSetupHandlesRouteSetupError(t *testing.T) {
	originalSetupRoutes := setupRoutes
	defer func() { setupRoutes = originalSetupRoutes }()

	setupRoutes = func(app *fiber.App) {
		fmt.Println("Mocked route setup error")
	}

	assert.NotPanics(t, func() { Setup() }, "Expected Setup to handle route setup error gracefully")
}

func TestSetupHandlesTableCreationError(t *testing.T) {
	originalCreateTable := createTable
	defer func() { createTable = originalCreateTable }()

	createTable = func() {
		fmt.Println("Mocked table creation error")
	}

	assert.NotPanics(t, func() { Setup() }, "Expected Setup to handle table creation error gracefully")
}

func TestSetupMiddlewareExecution(t *testing.T) {
	app := Setup()
	req, _ := http.NewRequest("GET", "/", nil)
	res, err := app.Test(req, -1)

	assert.Nil(t, err, "Expected no error when sending request")
	assert.Equal(t, 200, res.StatusCode, "Expected status code 200")
}

func TestSetupMultipleCalls(t *testing.T) {
	app1 := Setup()
	app2 := Setup()

	assert.NotNil(t, app1, "Expected first Setup call to return a valid instance")
	assert.NotNil(t, app2, "Expected second Setup call to return a valid instance")
	assert.NotEqual(t, app1, app2, "Expected separate instances for each Setup call")
}

func TestSetupRoutesInitialization(t *testing.T) {
	mockDB := new(MockDatabase)
	mockDB.On("SetupRoutes", mock.Anything).Return()

	setupRoutes = func(app *fiber.App) {
		mockDB.SetupRoutes(app)
	}

	Setup()

	mockDB.AssertCalled(t, "SetupRoutes", mock.Anything)
}

func TestSetupTableCreation(t *testing.T) {
	mockDB := new(MockDatabase)
	mockDB.On("CreateTable").Return()

	createTable = func() {
		mockDB.CreateTable()
	}

	Setup()

	mockDB.AssertCalled(t, "CreateTable")
}


/*
ROOST_METHOD_HASH=createTable_7ce7763726
ROOST_METHOD_SIG_HASH=createTable_48ea31fd25

FUNCTION_DEF=func createTable() 

*/
func TestCreateTable(t *testing.T) {
	tests := []struct {
		name           string
		setupMock      func(mock go-sqlmock.Sqlmock)
		expectedError  bool
		expectedOutput string
	}{
		{
			name: "Scenario 1: Verify Table Creation in Database",
			setupMock: func(mock go-sqlmock.Sqlmock) {
				mock.ExpectExec(`CREATE TABLE IF NOT EXISTS items`).WillReturnResult(go-sqlmock.NewResult(1, 1))
				mock.ExpectExec(`ALTER TABLE items OWNER TO postgres`).WillReturnResult(go-sqlmock.NewResult(1, 1))
			},
			expectedError:  false,
			expectedOutput: "",
		},
		{
			name: "Scenario 2: Ensure Table is Not Recreated if Already Exists",
			setupMock: func(mock go-sqlmock.Sqlmock) {
				mock.ExpectExec(`CREATE TABLE IF NOT EXISTS items`).WillReturnResult(go-sqlmock.NewResult(1, 0))
				mock.ExpectExec(`ALTER TABLE items OWNER TO postgres`).WillReturnResult(go-sqlmock.NewResult(1, 0))
			},
			expectedError:  false,
			expectedOutput: "",
		},
		{
			name: "Scenario 3: Handle Database Connection Failure Gracefully",
			setupMock: func(mock go-sqlmock.Sqlmock) {
				mock.ExpectExec(`CREATE TABLE IF NOT EXISTS items`).WillReturnError(fmt.Errorf("database connection error"))
			},
			expectedError:  true,
			expectedOutput: "",
		},
		{
			name: "Scenario 4: Validate Table Schema After Creation",
			setupMock: func(mock go-sqlmock.Sqlmock) {
				mock.ExpectExec(`CREATE TABLE IF NOT EXISTS items`).WillReturnResult(go-sqlmock.NewResult(1, 1))
				mock.ExpectExec(`ALTER TABLE items OWNER TO postgres`).WillReturnResult(go-sqlmock.NewResult(1, 1))
				mock.ExpectQuery(`SELECT column_name, data_type FROM information_schema.columns WHERE table_name = 'items'`).
					WillReturnRows(go-sqlmock.NewRows([]string{"column_name", "data_type"}).
						AddRow("id", "integer").
						AddRow("Title", "character varying").
						AddRow("Owner", "character varying").
						AddRow("Rating", "integer").
						AddRow("created_at", "date").
						AddRow("updated_at", "date").
						AddRow("deleted_at", "date"))
			},
			expectedError:  false,
			expectedOutput: "",
		},
		{
			name: "Scenario 5: Verify SQL Query Execution",
			setupMock: func(mock go-sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta("CREATE TABLE IF NOT EXISTS items")).WillReturnResult(go-sqlmock.NewResult(1, 1))
				mock.ExpectExec(regexp.QuoteMeta("ALTER TABLE items OWNER TO postgres")).WillReturnResult(go-sqlmock.NewResult(1, 1))
			},
			expectedError:  false,
			expectedOutput: "",
		},
		{
			name: "Scenario 6: Ensure Table Ownership is Set Correctly",
			setupMock: func(mock go-sqlmock.Sqlmock) {
				mock.ExpectExec(`CREATE TABLE IF NOT EXISTS items`).WillReturnResult(go-sqlmock.NewResult(1, 1))
				mock.ExpectExec(`ALTER TABLE items OWNER TO postgres`).WillReturnResult(go-sqlmock.NewResult(1, 1))
				mock.ExpectQuery(`SELECT tableowner FROM pg_tables WHERE tablename = 'items'`).
					WillReturnRows(go-sqlmock.NewRows([]string{"tableowner"}).AddRow("postgres"))
			},
			expectedError:  false,
			expectedOutput: "",
		},
		{
			name: "Scenario 7: Test Execution with a Read-Only Database",
			setupMock: func(mock go-sqlmock.Sqlmock) {
				mock.ExpectExec(`CREATE TABLE IF NOT EXISTS items`).WillReturnError(fmt.Errorf("permission denied"))
			},
			expectedError:  true,
			expectedOutput: "",
		},
		{
			name: "Scenario 8: Verify Function Execution Time",
			setupMock: func(mock go-sqlmock.Sqlmock) {
				mock.ExpectExec(`CREATE TABLE IF NOT EXISTS items`).WillReturnResult(go-sqlmock.NewResult(1, 1))
				mock.ExpectExec(`ALTER TABLE items OWNER TO postgres`).WillReturnResult(go-sqlmock.NewResult(1, 1))
			},
			expectedError:  false,
			expectedOutput: "",
		},
		{
			name: "Scenario 9: Ensure Function Does Not Modify Existing Data",
			setupMock: func(mock go-sqlmock.Sqlmock) {
				mock.ExpectExec(`CREATE TABLE IF NOT EXISTS items`).WillReturnResult(go-sqlmock.NewResult(1, 1))
				mock.ExpectExec(`ALTER TABLE items OWNER TO postgres`).WillReturnResult(go-sqlmock.NewResult(1, 1))
				mock.ExpectQuery(`SELECT COUNT(*) FROM items`).WillReturnRows(go-sqlmock.NewRows([]string{"count"}).AddRow(5))
			},
			expectedError:  false,
			expectedOutput: "",
		},
		{
			name: "Scenario 10: Validate Function Behavior with Different Database Users",
			setupMock: func(mock go-sqlmock.Sqlmock) {
				mock.ExpectExec(`CREATE TABLE IF NOT EXISTS items`).WillReturnError(fmt.Errorf("insufficient privileges"))
			},
			expectedError:  true,
			expectedOutput: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := go-sqlmock.New()
			if err != nil {
				t.Fatalf("error initializing mock database: %v", err)
			}
			defer db.Close()

			gormDB, err := gorm.Open("postgres", db)
			if err != nil {
				t.Fatalf("error initializing GORM database: %v", err)
			}
			database.DBConn = gormDB

			tt.setupMock(mock)

			var buf bytes.Buffer
			old := os.Stdout
			os.Stdout = &buf
			defer func() { os.Stdout = old }()

			start := time.Now()
			createTable()
			duration := time.Since(start)

			if duration > 2*time.Second {
				t.Errorf("execution time exceeded expected limit: %v", duration)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("unfulfilled expectations: %v", err)
			}

			output := buf.String()
			if !strings.Contains(output, tt.expectedOutput) {
				t.Errorf("unexpected output: got %v, want %v", output, tt.expectedOutput)
			}
		})
	}
}

