package main

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestQueryAllUser_bfa4cc5989(t *testing.T) {
	// TODO: Replace with your database connection details
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test case 1: Check if the function can successfully query all users
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Error("Failed to query all users, args: db")
			}
		}()

		QueryAllUser(db)
		t.Log("Successfully queried all users")
	}()

	// Test case 2: Check if the function can handle a nil db
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Log("Successfully handled a nil db")
			} else {
				t.Error("Failed to handle a nil db, args: nil")
			}
		}()

		QueryAllUser(nil)
	}()
}

func QueryAllUser(db *sql.DB) {
	type user struct {
		id        int
		username  string
		password  string
		createdAt time.Time
	}

	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user

		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", users)
}
