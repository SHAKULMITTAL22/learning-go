package main

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// TestInsertUser_7cd4b6ebc2 is a test function for insertUser function
func TestInsertUser_7cd4b6ebc2(t *testing.T) {
	// TODO: Replace with your actual database details
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		t.Error("Failed to connect to database: ", err)
	}
	defer db.Close()

	// Test Case 1: Successful Insertion
	username1 := "testUser1"
	password1 := "testPassword1"
	t.Log("Executing Test Case 1: Successful Insertion")
	err = insertUser(db, username1, password1)
	if err != nil {
		t.Error("Failed to insert user: ", username1, " Error: ", err)
	} else {
		t.Log("Test Case 1 Passed: Successful Insertion")
	}

	// Test Case 2: Insertion with existing username
	username2 := "testUser1" // same as username1
	password2 := "testPassword2"
	t.Log("Executing Test Case 2: Insertion with existing username")
	err = insertUser(db, username2, password2)
	if err == nil {
		t.Error("Test Case 2 Failed: Expected error due to duplicate username, got nil error")
	} else {
		t.Log("Test Case 2 Passed: Error received as expected for duplicate username")
	}
}

// insertUser is a function to insert a new user into the database
func insertUser(db *sql.DB, username string, password string) error {
	createdAt := time.Now()

	_, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	if err != nil {
		return err
	}

	return nil
}
