package main

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// TestInsertUser_7cd4b6ebc2 is the test function for insertUser
func TestInsertUser_7cd4b6ebc2(t *testing.T) {
	// TODO: Replace with the actual DB connection string
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test case 1: Insert a valid user
	username1 := "testuser1"
	password1 := "password1"
	err = insertUser(db, username1, password1)
	if err != nil {
		t.Error(fmt.Sprintf("Test case 1 failed. insertUser(%s, %s): %s", username1, password1, err))
	} else {
		t.Log("Test case 1 passed.")
	}

	// Test case 2: Insert a user with empty username
	username2 := ""
	password2 := "password2"
	err = insertUser(db, username2, password2)
	if err == nil {
		t.Error(fmt.Sprintf("Test case 2 failed. insertUser(%s, %s): %s", username2, password2, "Expected error but got nil"))
	} else {
		t.Log("Test case 2 passed.")
	}
}

// insertUser is the function to be tested
func insertUser(db *sql.DB, username string, password string) (sql.Result, error) {
	createdAt := time.Now()

	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	if err != nil {
		return nil, err
	}

	return result, nil
}
