package main

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// deleteUserTest function deletes a user from the database using the user's id
func deleteUserTest(db *sql.DB, id int) {
	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, id)
	if err != nil {
		log.Fatal(err)
	}
}

// TestDeleteUser_e531756cd2 function tests the deleteUserTest function
func TestDeleteUser_e531756cd2(t *testing.T) {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test case 1: Delete a user that exists in the database
	id := 1 // TODO: Change this to the id of the user that you want to delete
	deleteUserTest(db, id)
	row := db.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", id)
	var count int
	err = row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count != 0 {
		t.Errorf("TestDeleteUser_e531756cd2 failed while deleting existing user, user with id %d still exists", id)
	} else {
		t.Logf("TestDeleteUser_e531756cd2 passed while deleting existing user, user with id %d doesn't exist", id)
	}

	// Test case 2: Delete a user that doesn't exist in the database
	id = 9999 // TODO: Change this to the id of the user that you want to delete
	deleteUserTest(db, id)
	row = db.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", id)
	err = row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count != 0 {
		t.Errorf("TestDeleteUser_e531756cd2 failed while deleting non-existing user, user with id %d still exists", id)
	} else {
		t.Logf("TestDeleteUser_e531756cd2 passed while deleting non-existing user, user with id %d doesn't exist", id)
	}
}
