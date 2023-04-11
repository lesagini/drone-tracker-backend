package models

import (
	"database/sql"
	"drones/utils"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testingQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := utils.GetConfig("../..")

	if err != nil {
		log.Fatal("cannot get config")
	}

	testDb, err = sql.Open(config.DbDriver, config.DbSource)

	if err != nil {
		log.Fatal("connection to database failed", err)
	}

	testingQueries = New(testDb)

	os.Exit(m.Run())

}
