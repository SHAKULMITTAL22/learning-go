package main

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestQueryAllUser_bfa4cc5989(t *testing.T) {
	// TODO: Replace with your actual DB details
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Test Case 1: Check if the database connection is successful
	if err = db.Ping(); err != nil {
		t.Errorf("Failed to connect to the database: %v", err)
	} else {
		t.Log("Successfully connected to the database.")
	}

	// Test Case 2: Check if the query returns the correct data
	type user struct {
		id        int
		username  string
		password  string
		createdAt time.Time
	}

	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
	if err != nil {
		t.Errorf("Failed to execute query: %v", err)
	} else {
		t.Log("Query executed successfully.")
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user

		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		if err != nil {
			t.Errorf("Failed to scan row: %v", err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		t.Errorf("Rows contains an error: %v", err)
	} else {
		t.Log("Data fetched successfully.")
	}

	t.Logf("%#v", users)
}
