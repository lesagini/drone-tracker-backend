package models

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:secret@localhost:5000/drones?sslmode=disable"
)

var testingQueries *Queries
var testDb *sql.DB
func TestMain(m *testing.M) {
	var err error
	testDb, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("connection to database failed", err)
	}

	testingQueries = New(testDb)

	os.Exit(m.Run())

}
