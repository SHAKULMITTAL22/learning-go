package main

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// TestDeleteUser_e531756cd2 tests the deleteUser function
func TestDeleteUser_e531756cd2(t *testing.T) {
	// TODO: replace with your own database connection string
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Test case 1: Delete an existing user
	id1 := 1
	deleteUser(db, id1)
	row := db.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", id1)
	var count int
	err = row.Scan(&count)
	if err != nil {
		t.Error("Failed to execute query: ", err)
	}
	if count != 0 {
		t.Error("Failed to delete user with id ", id1)
	} else {
		t.Log("Successfully deleted user with id ", id1)
	}

	// Test case 2: Delete a non-existing user
	id2 := 9999
	deleteUser(db, id2)
	row = db.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", id2)
	err = row.Scan(&count)
	if err != nil {
		t.Error("Failed to execute query: ", err)
	}
	if count != 0 {
		t.Error("User with id ", id2, " was not supposed to exist")
	} else {
		t.Log("Successfully handled non-existing user with id ", id2)
	}
}
