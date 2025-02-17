package item

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tannergabriel/learning-go/advanced-programs/FiberPostgresCRUD/database"
	"encoding/json"
	"errors"
	"io/ioutil"
	"github.com/gofiber/fiber/v2"
	"os"
	"strings"
)





type Item struct {
	gorm.Model
	Title  string `json:"Title"`
	Owner  string `json:"Owner"`
	Rating int    `json:"Rating"`
}


/*
ROOST_METHOD_HASH=GetItems_2c8f727a30
ROOST_METHOD_SIG_HASH=GetItems_5fa489a72b

FUNCTION_DEF=func GetItems(c *fiber.Ctx) 

*/
func TestGetItems(t *testing.T) {
	tests := []struct {
		name           string
		setupMock      func(sqlmock.Sqlmock)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Successfully Retrieve All Items from the Database",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "title", "owner", "rating"}).
					AddRow(1, "Item1", "Owner1", 5).
					AddRow(2, "Item2", "Owner2", 4)
				mock.ExpectQuery(`^SELECT \* FROM "items"$`).WillReturnRows(rows)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"ID":1,"Title":"Item1","Owner":"Owner1","Rating":5},{"ID":2,"Title":"Item2","Owner":"Owner2","Rating":4}]`,
		},
		{
			name: "Retrieve Items When Database is Empty",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "title", "owner", "rating"})
				mock.ExpectQuery(`^SELECT \* FROM "items"$`).WillReturnRows(rows)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "[]",
		},
		{
			name: "Database Connection Failure",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`^SELECT \* FROM "items"$`).WillReturnError(fmt.Errorf("database connection error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"database connection error"}`,
		},
		{
			name: "Database Query Failure",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`^SELECT \* FROM "items"$`).WillReturnError(fmt.Errorf("query execution failed"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"query execution failed"}`,
		},
		{
			name: "Large Dataset Retrieval",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "title", "owner", "rating"})
				for i := 1; i <= 1000; i++ {
					rows.AddRow(i, fmt.Sprintf("Item%d", i), fmt.Sprintf("Owner%d", i), i%5)
				}
				mock.ExpectQuery(`^SELECT \* FROM "items"$`).WillReturnRows(rows)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
		{
			name: "Verify JSON Response Structure",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "title", "owner", "rating"}).
					AddRow(1, "Item1", "Owner1", 5)
				mock.ExpectQuery(`^SELECT \* FROM "items"$`).WillReturnRows(rows)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"ID":1,"Title":"Item1","Owner":"Owner1","Rating":5}]`,
		},
		{
			name: "Concurrent Requests Handling",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "title", "owner", "rating"}).
					AddRow(1, "Item1", "Owner1", 5)
				mock.ExpectQuery(`^SELECT \* FROM "items"$`).WillReturnRows(rows)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"ID":1,"Title":"Item1","Owner":"Owner1","Rating":5}]`,
		},
		{
			name: "Handling of Unexpected Data in Database",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "title", "owner", "rating"}).
					AddRow(nil, "Item1", "Owner1", "invalid_rating")
				mock.ExpectQuery(`^SELECT \* FROM "items"$`).WillReturnRows(rows)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"invalid data format"}`,
		},
		{
			name: "Performance Benchmarking",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "title", "owner", "rating"})
				for i := 1; i <= 10000; i++ {
					rows.AddRow(i, fmt.Sprintf("Item%d", i), fmt.Sprintf("Owner%d", i), i%5)
				}
				mock.ExpectQuery(`^SELECT \* FROM "items"$`).WillReturnRows(rows)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Failed to create mock database: %v", err)
			}
			defer db.Close()

			gormDB, err := gorm.Open("_", db)
			if err != nil {
				t.Fatalf("Failed to open gorm database: %v", err)
			}
			database.DBConn = gormDB

			tt.setupMock(mock)

			app := fiber.New()
			app.Get("/items", GetItems)

			req := httptest.NewRequest(http.MethodGet, "/items", nil)
			resp, err := app.Test(req)
			if err != nil {
				t.Fatalf("Failed to execute request: %v", err)
			}

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}

			var body bytes.Buffer
			_, err = body.ReadFrom(resp.Body)
			if err != nil {
				t.Fatalf("Failed to read response body: %v", err)
			}

			if tt.expectedBody != "" && body.String() != tt.expectedBody {
				t.Errorf("Expected body %s, got %s", tt.expectedBody, body.String())
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Unfulfilled mock expectations: %v", err)
			}
		})
	}

	t.Run("Concurrent Requests Handling", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("Failed to create mock database: %v", err)
		}
		defer db.Close()

		gormDB, err := gorm.Open("_", db)
		if err != nil {
			t.Fatalf("Failed to open gorm database: %v", err)
		}
		database.DBConn = gormDB

		rows := sqlmock.NewRows([]string{"id", "title", "owner", "rating"}).
			AddRow(1, "Item1", "Owner1", 5)
		mock.ExpectQuery(`^SELECT \* FROM "items"$`).WillReturnRows(rows)

		app := fiber.New()
		app.Get("/items", GetItems)

		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				req := httptest.NewRequest(http.MethodGet, "/items", nil)
				app.Test(req)
			}()
		}
		wg.Wait()
	})
}


/*
ROOST_METHOD_HASH=NewItem_0c65f54868
ROOST_METHOD_SIG_HASH=NewItem_9f9fb16e38

FUNCTION_DEF=func NewItem(c *fiber.Ctx) 

*/
func TestNewItem(t *testing.T) {
	gormDB, mock := setupMockDB()
	defer gormDB.Close()

	app := fiber.New()

	tests := []struct {
		name           string
		requestBody    interface{}
		mockDBSetup    func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Successfully Creating a New Item",
			requestBody: Item{
				Title:  "Test Item",
				Owner:  "John Doe",
				Rating: 5,
			},
			mockDBSetup: func() {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO \"items\"").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `"Title":"Test Item","Owner":"John Doe","Rating":5`,
		},
		{
			name:           "Handling Invalid JSON Input",
			requestBody:    "{invalid_json}",
			mockDBSetup:    func() {},
			expectedStatus: http.StatusServiceUnavailable,
			expectedBody:   "invalid character",
		},
		{
			name: "Database Failure During Item Creation",
			requestBody: Item{
				Title:  "Test Item",
				Owner:  "John Doe",
				Rating: 5,
			},
			mockDBSetup: func() {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO \"items\"").WillReturnError(errors.New("database error"))
				mock.ExpectRollback()
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "database error",
		},
		{
			name:           "Handling Empty Request Body",
			requestBody:    "",
			mockDBSetup:    func() {},
			expectedStatus: http.StatusServiceUnavailable,
			expectedBody:   "unexpected end of JSON input",
		},
		{
			name: "Handling Missing Required Fields",
			requestBody: map[string]interface{}{
				"Title": "Missing Owner",
			},
			mockDBSetup:    func() {},
			expectedStatus: http.StatusServiceUnavailable,
			expectedBody:   "missing required fields",
		},
		{
			name: "Handling Large Request Body",
			requestBody: Item{
				Title:  string(make([]byte, 10000)),
				Owner:  "John Doe",
				Rating: 5,
			},
			mockDBSetup:    func() {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "payload too large",
		},
		{
			name: "Handling Duplicate Items",
			requestBody: Item{
				Title:  "Duplicate Item",
				Owner:  "John Doe",
				Rating: 5,
			},
			mockDBSetup: func() {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO \"items\"").WillReturnError(errors.New("duplicate key value violates unique constraint"))
				mock.ExpectRollback()
			},
			expectedStatus: http.StatusConflict,
			expectedBody:   "duplicate key value",
		},
		{
			name: "Handling Unexpected Server Errors",
			requestBody: Item{
				Title:  "Unexpected Error",
				Owner:  "John Doe",
				Rating: 5,
			},
			mockDBSetup: func() {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO \"items\"").WillReturnError(errors.New("unexpected error"))
				mock.ExpectRollback()
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "unexpected error",
		},
		{
			name: "Handling Slow Requests (Timeout)",
			requestBody: Item{
				Title:  "Slow Request",
				Owner:  "John Doe",
				Rating: 5,
			},
			mockDBSetup: func() {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO \"items\"").WillReturnError(errors.New("timeout"))
				mock.ExpectRollback()
			},
			expectedStatus: http.StatusGatewayTimeout,
			expectedBody:   "timeout",
		},
		{
			name: "Handling Unauthorized Requests",
			requestBody: Item{
				Title:  "Unauthorized",
				Owner:  "John Doe",
				Rating: 5,
			},
			mockDBSetup:    func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "unauthorized",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockDBSetup()

			var reqBody []byte
			var err error

			switch v := tt.requestBody.(type) {
			case string:
				reqBody = []byte(v)
			default:
				reqBody, err = json.Marshal(v)
				if err != nil {
					t.Fatalf("failed to marshal request body: %v", err)
				}
			}

			req := httptest.NewRequest("POST", "/new-item", bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")
			resp, err := app.Test(req, -1)
			if err != nil {
				t.Fatalf("failed to execute request: %v", err)
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("failed to read response body: %v", err)
			}

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}

			if !bytes.Contains(body, []byte(tt.expectedBody)) {
				t.Errorf("expected response body to contain %q, got %q", tt.expectedBody, string(body))
			}
		})
	}
}

func setupMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(fmt.Sprintf("failed to open sqlmock database: %v", err))
	}

	gormDB, err := gorm.Open("postgres", db)
	if err != nil {
		panic(fmt.Sprintf("failed to open gorm database: %v", err))
	}

	database.DBConn = gormDB
	return gormDB, mock
}


/*
ROOST_METHOD_HASH=GetItem_a9c09b8bd1
ROOST_METHOD_SIG_HASH=GetItem_351efb6abf

FUNCTION_DEF=func GetItem(c *fiber.Ctx) 

*/
func GetItem(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var item Item

	if err := db.First(&item, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Item not found"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	return c.JSON(item)
}

func TestGetItem(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatalf("Failed to open gorm database: %v", err)
	}
	database.DBConn = gormDB

	tests := []struct {
		name           string
		itemID         string
		mockSetup      func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:   "Successfully Retrieve an Existing Item",
			itemID: "1",
			mockSetup: func() {
				mock.ExpectQuery(`SELECT \* FROM "items" WHERE "items"."deleted_at" IS NULL AND "items"."id" = \$1`).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "title", "owner", "rating"}).
						AddRow(1, "Test Item", "John Doe", 5))
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"ID":1,"Title":"Test Item","Owner":"John Doe","Rating":5}`,
		},
		{
			name:   "Retrieve a Non-Existent Item",
			itemID: "999",
			mockSetup: func() {
				mock.ExpectQuery(`SELECT \* FROM "items" WHERE "items"."deleted_at" IS NULL AND "items"."id" = \$1`).
					WithArgs(999).
					WillReturnRows(sqlmock.NewRows([]string{"id", "title", "owner", "rating"}))
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"Item not found"}`,
		},
		{
			name:           "Invalid ID Format",
			itemID:         "abc",
			mockSetup:      func() {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Invalid ID format"}`,
		},
		{
			name:   "Database Connection Failure",
			itemID: "1",
			mockSetup: func() {
				mock.ExpectQuery(`SELECT \* FROM "items" WHERE "items"."deleted_at" IS NULL AND "items"."id" = \$1`).
					WithArgs(1).
					WillReturnError(fmt.Errorf("database connection error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"Database error"}`,
		},
		{
			name:           "Empty ID Parameter",
			itemID:         "",
			mockSetup:      func() {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"ID parameter is required"}`,
		},
		{
			name:           "Handling of Special Characters in ID",
			itemID:         "!@#$",
			mockSetup:      func() {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Invalid ID format"}`,
		},
		{
			name:   "Large Numeric ID",
			itemID: "999999999999999999",
			mockSetup: func() {
				mock.ExpectQuery(`SELECT \* FROM "items" WHERE "items"."deleted_at" IS NULL AND "items"."id" = \$1`).
					WithArgs(999999999999999999).
					WillReturnRows(sqlmock.NewRows([]string{"id", "title", "owner", "rating"}))
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"Item not found"}`,
		},
		{
			name:   "Concurrent Requests for the Same Item",
			itemID: "1",
			mockSetup: func() {
				mock.ExpectQuery(`SELECT \* FROM "items" WHERE "items"."deleted_at" IS NULL AND "items"."id" = \$1`).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "title", "owner", "rating"}).
						AddRow(1, "Test Item", "John Doe", 5))
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"ID":1,"Title":"Test Item","Owner":"John Doe","Rating":5}`,
		},
	}

	app := fiber.New()
	app.Get("/item/:id", GetItem)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/item/%s", tt.itemID), nil)
			resp, err := app.Test(req)
			if err != nil {
				t.Fatalf("Failed to execute request: %v", err)
			}

			var buf bytes.Buffer
			_, err = buf.ReadFrom(resp.Body)
			if err != nil {
				t.Fatalf("Failed to read response body: %v", err)
			}

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}

			if buf.String() != tt.expectedBody {
				t.Errorf("Expected body %s, got %s", tt.expectedBody, buf.String())
			}
		})
	}

	t.Run("Concurrent Requests for the Same Item", func(t *testing.T) {
		var wg sync.WaitGroup
		numRequests := 5
		wg.Add(numRequests)

		for i := 0; i < numRequests; i++ {
			go func() {
				defer wg.Done()
				req := httptest.NewRequest(http.MethodGet, "/item/1", nil)
				resp, err := app.Test(req)
				if err != nil {
					t.Errorf("Failed to execute request: %v", err)
					return
				}
				if resp.StatusCode != http.StatusOK {
					t.Errorf("Expected status 200, got %d", resp.StatusCode)
				}
			}()
		}

		wg.Wait()
	})
}


/*
ROOST_METHOD_HASH=DeleteItem_29a4e58a35
ROOST_METHOD_SIG_HASH=DeleteItem_bc718a95a3

FUNCTION_DEF=func DeleteItem(c *fiber.Ctx) 

*/
func TestDeleteItem(t *testing.T) {
	tests := []struct {
		name           string
		setupMock      func(sqlmock.Sqlmock)
		inputID        string
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Successfully Delete an Existing Item",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT \* FROM "items" WHERE "items"."deleted_at" IS NULL AND "items"."id" = \$1 ORDER BY "items"."id" ASC LIMIT 1`).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow(1, "Test Item"))

				mock.ExpectExec(`DELETE FROM "items" WHERE "items"."id" = \$1`).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			inputID:        "1",
			expectedStatus: 200,
			expectedBody:   "Item successfully deleted",
		},
		{
			name: "Attempt to Delete a Non-Existent Item",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT \* FROM "items" WHERE "items"."deleted_at" IS NULL AND "items"."id" = \$1 ORDER BY "items"."id" ASC LIMIT 1`).
					WithArgs(999).
					WillReturnRows(sqlmock.NewRows([]string{"id", "title"}))
			},
			inputID:        "999",
			expectedStatus: 500,
			expectedBody:   "No item found with given ID",
		},
		{
			name:           "Delete an Item with an Invalid ID Format",
			setupMock:      func(mock sqlmock.Sqlmock) {},
			inputID:        "abc",
			expectedStatus: 500,
			expectedBody:   "No item found with given ID",
		},
		{
			name: "Database Connection Failure During Deletion",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT \* FROM "items" WHERE "items"."deleted_at" IS NULL AND "items"."id" = \$1 ORDER BY "items"."id" ASC LIMIT 1`).
					WithArgs(1).
					WillReturnError(errors.New("database connection error"))
			},
			inputID:        "1",
			expectedStatus: 500,
			expectedBody:   "No item found with given ID",
		},
		{
			name:           "Delete an Item with Special Characters in ID",
			setupMock:      func(mock sqlmock.Sqlmock) {},
			inputID:        "!@#",
			expectedStatus: 500,
			expectedBody:   "No item found with given ID",
		},
		{
			name: "Concurrent Deletion of the Same Item",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT \* FROM "items" WHERE "items"."deleted_at" IS NULL AND "items"."id" = \$1 ORDER BY "items"."id" ASC LIMIT 1`).
					WithArgs(2).
					WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow(2, "Concurrent Item"))

				mock.ExpectExec(`DELETE FROM "items" WHERE "items"."id" = \$1`).
					WithArgs(2).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			inputID:        "2",
			expectedStatus: 200,
			expectedBody:   "Item successfully deleted",
		},
		{
			name:           "Delete an Item with a Large ID Value",
			setupMock:      func(mock sqlmock.Sqlmock) {},
			inputID:        "999999999999",
			expectedStatus: 500,
			expectedBody:   "No item found with given ID",
		},
		{
			name: "Delete an Item When the Database is Read-Only",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT \* FROM "items" WHERE "items"."deleted_at" IS NULL AND "items"."id" = \$1 ORDER BY "items"."id" ASC LIMIT 1`).
					WithArgs(3).
					WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow(3, "Read-Only Item"))

				mock.ExpectExec(`DELETE FROM "items" WHERE "items"."id" = \$1`).
					WithArgs(3).
					WillReturnError(errors.New("database is read-only"))
			},
			inputID:        "3",
			expectedStatus: 500,
			expectedBody:   "No item found with given ID",
		},
		{
			name:           "Delete an Item with a Null ID",
			setupMock:      func(mock sqlmock.Sqlmock) {},
			inputID:        "",
			expectedStatus: 500,
			expectedBody:   "No item found with given ID",
		},
		{
			name: "Delete an Item and Verify Response Message",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT \* FROM "items" WHERE "items"."deleted_at" IS NULL AND "items"."id" = \$1 ORDER BY "items"."id" ASC LIMIT 1`).
					WithArgs(4).
					WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow(4, "Test Item"))

				mock.ExpectExec(`DELETE FROM "items" WHERE "items"."id" = \$1`).
					WithArgs(4).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			inputID:        "4",
			expectedStatus: 200,
			expectedBody:   "Item successfully deleted",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
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

			app := fiber.New()
			app.Delete("/item/:id", DeleteItem)

			req := fmt.Sprintf("DELETE /item/%s HTTP/1.1\r\nHost: localhost\r\n\r\n", tt.inputID)
			r := strings.NewReader(req)
			w := &bytes.Buffer{}

			os.Stdout = w
			app.Test(r)

			response := w.String()
			if !strings.Contains(response, tt.expectedBody) {
				t.Errorf("expected response body to contain %q, got %q", tt.expectedBody, response)
			}
		})
	}
}

