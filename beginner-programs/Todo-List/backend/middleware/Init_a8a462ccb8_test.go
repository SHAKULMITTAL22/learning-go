package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tannergabriel/learning-go/beginner-programs/todo-list/backend/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var dbName = "todolist"
var collName = "todos"

func init() {
	var connectionString string

	// Get environment variable for connection string
	host := goDotEnvVariable("HOST")
	if "" == host {
		connectionString = "mongodb://localhost:27017"
	} else {
		connectionString = "mongodb://" + host + ":27017"
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection instance created!")
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func TestInit_a8a462ccb8(t *testing.T) {
	// Test case 1: Check if the connection string is empty
	connectionString := ""
	host := goDotEnvVariable("HOST")
	if "" == host {
		connectionString = "mongodb://localhost:27017"
	} else {
		connectionString = "mongodb://" + host + ":27017"
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		t.Error(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		t.Error(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection instance created!")

	// Test case 2: Check if the connection string is not empty
	connectionString = "mongodb://localhost:27017"
	host = goDotEnvVariable("HOST")
	if "" == host {
		connectionString = "mongodb://localhost:27017"
	} else {
		connectionString = "mongodb://" + host + ":27017"
	}

	// Set client options
	clientOptions = options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		t.Error(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		t.Error(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection instance created!")
}
