// Go test case

package main

import (
	"testing"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/database"
)

func TestInitDatabase(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		os.Setenv("HOST", "localhost")
		defer os.Unsetenv("HOST")

		initDatabase()

		if _, ok := database.DBConn.(*gorm.DB); !ok {
			t.Errorf("DATABASE CONNECTION error, DBConn : %v", database.DBConn)
			return
		}
		
		t.Log("Success case passed.")
	})

	t.Run("failure", func(t *testing.T) {
		os.Setenv("HOST", "wrongHost")
		defer os.Unsetenv("HOST")
		defer func() {
			if r := recover(); r == nil {
				t.Error("Panic not recovered, InitDatabase did not fail with wrong HOST.")
			}
		}()

		initDatabase()

		if _, ok := database.DBConn.(*gorm.DB); ok {
			t.Errorf("DATABASE CONNECTION error, DBConn : %v", database.DBConn)
			return
		}

		t.Log("Failure case passed.")
	})
}
