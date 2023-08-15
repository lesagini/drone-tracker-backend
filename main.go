package main

import (
	"database/sql"
	"drones/api"
	db "drones/db/models"
	"drones/utils"
	"log"

	_ "github.com/lib/pq"
)

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://postgres:secret@localhost:5000/drones?sslmode=disable"
// 	address  = "0.0.0.0:8080"
// )

func main() {
	config, err := utils.GetConfig(".")

	if err != nil {
		log.Fatal("cannot get config")
	}
	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("connection to database failed", err)
	}

	store := db.NewTransaction(conn)
	server := api.NewServer(store)

	err = server.Start(config.Address)

	if err != nil {
		log.Fatal(err)
	}

}
