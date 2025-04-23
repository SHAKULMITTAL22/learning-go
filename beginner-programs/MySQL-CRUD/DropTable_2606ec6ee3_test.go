package main

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func dropTestTable(db *sql.DB) {
	if _, err := db.Exec("DROP TABLE IF EXISTS `users`"); err != nil {
		log.Fatal(err)
	}
}

func TestDropTable_2606ec6ee3(t *testing.T) {

	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test case 1: Table `users` exists in the database
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `users` (`id` int, `name` varchar(255))")
	if err != nil {
		t.Error("Failed to create table: ", err)
		return
	}
	dropTestTable(db)
	row := db.QueryRow("SELECT COUNT(*) FROM information_schema.tables WHERE table_name = 'users'")
	var count int
	err = row.Scan(&count)
	if err != nil {
		t.Error("Failed to execute query: ", err)
		return
	}
	if count != 0 {
		t.Error("Failed to drop table `users`")
		return
	}
	t.Log("Test case 1: Passed")

	// Test case 2: Table `users` does not exist in the database
	dropTestTable(db)
	row = db.QueryRow("SELECT COUNT(*) FROM information_schema.tables WHERE table_name = 'users'")
	err = row.Scan(&count)
	if err != nil {
		t.Error("Failed to execute query: ", err)
		return
	}
	if count != 0 {
		t.Error("Table `users` exists but it should not")
		return
	}
	t.Log("Test case 2: Passed")
}
