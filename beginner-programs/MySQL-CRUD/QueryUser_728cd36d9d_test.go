package main

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// TestQueryUser_728cd36d is a test function for queryUser
func TestQueryUser_728cd36d(t *testing.T) {
	// TODO: Replace with your database connection details
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test case 1: Existing user
	userID1 := 1 // TODO: Replace with an existing user ID
	t.Run("Existing user", func(t *testing.T) {
		if err := queryUser(db, userID1); err != nil {
			t.Errorf("queryUser(%v) returned error: %v", userID1, err)
		} else {
			t.Logf("Test for existing user passed")
		}
	})

	// Test case 2: Non-existing user
	userID2 := 9999 // TODO: Replace with a non-existing user ID
	t.Run("Non-existing user", func(t *testing.T) {
		if err := queryUser(db, userID2); err == nil {
			t.Errorf("queryUser(%v) should return error for non-existing user, got nil", userID2)
		} else {
			t.Logf("Test for non-existing user passed")
		}
	})
}

func queryUser(db *sql.DB, userID int) error {
	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)

	query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
	if err := db.QueryRow(query, userID).Scan(&id, &username, &password, &createdAt); err != nil {
		return err
	}

	fmt.Println(id, username, password, createdAt)
	return nil
}
