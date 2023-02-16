package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver     = "postgres"
	dbSourceName = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbSourceName)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	// recieves new database Connection
	testQueries = New(testDB)

	os.Exit(m.Run())
}
