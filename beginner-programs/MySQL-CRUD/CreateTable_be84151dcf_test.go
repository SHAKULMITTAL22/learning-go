package main

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// TestCreateTable_be84151dcf is a test function for createTable function
func TestCreateTable_be84151dcf(t *testing.T) {
	// TODO: replace with your database connection string
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test case 1: successful table creation
	t.Run("success", func(t *testing.T) {
		err := createTableTest(db)
		if err != nil {
			t.Errorf("createTableTest() error = %v", err)
			return
		}
		t.Log("createTableTest() passed")
	})

	// Test case 2: table already exists
	t.Run("table exists", func(t *testing.T) {
		err := createTableTest(db)
		if err == nil {
			t.Error("createTableTest() should fail when table already exists")
			return
		}
		if err.Error() != "Error 1050: Table 'users' already exists" {
			t.Errorf("createTableTest() error = %v, wantErr = %v", err, "Error 1050: Table 'users' already exists")
			return
		}
		t.Log("createTableTest() passed")
	})
}

// createTableTest creates a new users table in the database
func createTableTest(db *sql.DB) error {
	query := `
		CREATE TABLE users (
			id INT AUTO_INCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME,
			PRIMARY KEY (id)
		);`

	_, err := db.Exec(query)
	return err
}
