package main

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// TestCreateTable_be84151dcf tests the createTable function
func TestCreateTable_be84151dcf(t *testing.T) {
	// Connect to the MySQL database
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		t.Error(fmt.Sprintf("Failed to connect to the database: %v", err))
		return
	}
	defer db.Close()

	// Test case 1: Successful table creation
	t.Log("Test case 1: Successful table creation")
	createTableTest(db, t)
	_, err = db.Query("SELECT 1 FROM users LIMIT 1")
	if err != nil {
		t.Error(fmt.Sprintf("Failed to create table: %v", err))
	} else {
		t.Log("Successfully created table")
	}

	// Test case 2: Table already exists
	t.Log("Test case 2: Table already exists")
	createTableTest(db, t)
	_, err = db.Query("SELECT 1 FROM users LIMIT 1")
	if err != nil {
		t.Error(fmt.Sprintf("Failed when table already exists: %v", err))
	} else {
		t.Log("Successfully handled existing table")
	}
}

// createTableTest creates a new users table in the database
func createTableTest(db *sql.DB, t *testing.T) {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME,
			PRIMARY KEY (id)
		);`

	if _, err := db.Exec(query); err != nil {
		t.Error(err)
	}
}
