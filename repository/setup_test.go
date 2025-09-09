package repository

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/glebarez/go-sqlite"
)

var testRepo *SQLiteRepository

func TestMain(m *testing.M) {
	path := "./testdata/sql.db"
	_ = os.Remove(path)

	err := os.MkdirAll("./testdata", 0755)
	if err != nil {
		log.Printf("ERROR creating testdata directory: %v", err)
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		log.Printf("ERROR opening database: %v", err)
		os.Exit(1)
	}

	testRepo = NewSQLiteRepository(db)

	code := m.Run()

	db.Close()
	_ = os.Remove(path)

	os.Exit(code)
}
