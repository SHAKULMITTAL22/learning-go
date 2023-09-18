package main

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// TestQueryAllUser_bfa4cc5989 is a test function for queryAllUser
func TestQueryAllUser_bfa4cc5989(t *testing.T) {
	// TODO: Replace with your own database details
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Test case 1: Check if the function can handle empty tables
	// TODO: Make sure the users table is empty before running this test
	err = queryAllUser(db)
	if err != nil {
		t.Errorf("Failed on an empty table: %v", err)
	} else {
		t.Log("Passed on an empty table")
	}

	// Test case 2: Check if the function can handle populated tables
	// TODO: Populate the users table before running this test
	err = queryAllUser(db)
	if err != nil {
		t.Errorf("Failed on a populated table: %v", err)
	} else {
		t.Log("Passed on a populated table")
	}
}

func queryAllUser(db *sql.DB) error {
	type user struct {
		id        int
		username  string
		password  string
		createdAt time.Time
	}

	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
	if err != nil {
		return err
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user

		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		if err != nil {
			return err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return err
	}

	fmt.Printf("%#v", users)

	return nil
}
