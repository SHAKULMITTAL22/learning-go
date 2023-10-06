// Test generated by RoostGPT for test roost-test using AI Type Azure Open AI and AI Model roost-gpt4-32k

package item_test

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/tannergabriel/learning-go/advanced-programs/FiberPostgresCRUD/database"
	"github.com/tannergabriel/learning-go/advanced-programs/FiberPostgresCRUD/item"
	"testing"
)

func TestNewItem(t *testing.T) {
	testCases := []struct {
		name           string
		body           string
		parseError     error
		createError    error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid Item Creation",
			body:           `{"name": "testItem", "description": "this is a test"}`,
			parseError:     nil,
			createError:    nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"name": "testItem", "description": "this is a test"}`,
		},
		{
			name:           "Invalid Item Creation",
			body:           `{"name":,}`,
			parseError:     errors.New("parse error"),
			createError:    nil,
			expectedStatus: http.StatusServiceUnavailable,
			expectedBody:   `{"error": "parse error"}`,
		},
		{
			name:           "Empty Item Creation",
			body:           "{}",
			parseError:     errors.New("empty object"),
			createError:    nil,
			expectedStatus: http.StatusServiceUnavailable,
			expectedBody:   `{"error": "empty object"}`,
		},
		{
			name:           "Database Connection Failure",
			body:           `{"name": "testItem", "description": "this is a test"}`,
			parseError:     nil,
			createError:    gorm.ErrRecordNotFound,
			expectedStatus: http.StatusServiceUnavailable,
			expectedBody:   `{"error": "database connection error"}`,
		},
		{
			name:           "Database Timeout",
			body:           `{"name": "testItem", "description": "this is a test"}`,
			parseError:     nil,
			createError:    gorm.ErrQueryTimeout,
			expectedStatus: http.StatusRequestTimeout,
			expectedBody:   `{"error": "database timeout"}`,
		},
		{
			name:           "Invalid Data Do Not Align With the Item Model",
			body:           `{"name": 123}`,
			parseError:     errors.New("invalid data"),
			createError:    nil,
			expectedStatus: http.StatusServiceUnavailable,
			expectedBody:   `{"error": "invalid data"}`,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/newitem", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			c := fiber.New()

			item.DBConn = &mockDB{
				parseError:  tt.parseError,
				createError: tt.createError,
			}

			item.NewItem(c)

			res := rec.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %v, got %v", tt.expectedStatus, res.StatusCode)
			}

			body, _ := ioutil.ReadAll(res.Body)
			if string(body) != tt.expectedBody {
				t.Errorf("expected body %q, got %q", tt.expectedBody, string(body))
			}
		})
	}
}

type mockDB struct {
	parseError  error
	createError error
}

func (mdb *mockDB) Create(i interface{}) error {
	return mdb.createError
}

