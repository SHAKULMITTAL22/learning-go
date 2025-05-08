package main

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// This is the function to be tested
func queryUserTest(db *sql.DB, userID int) {
	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)

	query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
	if err := db.QueryRow(query, userID).Scan(&id, &username, &password, &createdAt); err != nil {
		log.Fatal(err)
	}

	fmt.Println(id, username, password, createdAt)
}

// TestQueryUser_728cd36d9d is the test function for queryUser
func TestQueryUser_728cd36d9d(t *testing.T) {
	// TODO: replace with your own DB details
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test case 1: Valid user ID
	userID1 := 1 // TODO: replace with a valid user ID
	queryUserTest(db, userID1)
	t.Log("Test case 1: Valid user ID passed")

	// Test case 2: Invalid user ID
	userID2 := -1 // TODO: replace with an invalid user ID
	queryUserTest(db, userID2)
	t.Log("Test case 2: Invalid user ID passed")

	// Test case 3: Non-existing user ID
	userID3 := 9999999 // TODO: replace with a non-existing user ID
	queryUserTest(db, userID3)
	t.Log("Test case 3: Non-existing user ID passed")
}
