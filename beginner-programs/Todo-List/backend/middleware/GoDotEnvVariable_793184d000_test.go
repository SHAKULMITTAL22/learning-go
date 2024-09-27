package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/tannergabriel/learning-go/beginner-programs/todo-list/backend/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestGoDotEnvVariable_793184d000(t *testing.T) {
	// Test case 1: Key exists and value is returned
	key := "TEST_KEY"
	expectedValue := "test_value"
	os.Setenv(key, expectedValue)
	value := goDotEnvVariable(key)
	if value != expectedValue {
		t.Errorf("Expected value %s, got %s", expectedValue, value)
	}

	// Test case 2: Key does not exist and empty string is returned
	key = "NON_EXISTENT_KEY"
	expectedValue = ""
	value = goDotEnvVariable(key)
	if value != expectedValue {
		t.Errorf("Expected value %s, got %s", expectedValue, value)
	}
}
