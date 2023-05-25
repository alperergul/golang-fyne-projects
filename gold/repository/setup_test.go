package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/glebarez/go-sqlite"
)

var testRepo *SQLiteRepository

func TestMain(m *testing.M) {
	fmt.Println("Alper Test!")
	_ = os.Remove("./testdata/sql.db")
	path := "./testdata/sql.db"

	db, err := sql.Open("sqlite", path)
	if err != nil {
		log.Println(err)
	}

	testRepo = NewSQLiteRepository(db)
	os.Exit(m.Run())
}
