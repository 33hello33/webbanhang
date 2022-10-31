package main

import (
	"database/sql"
	"log"
	"webbanhang/api"
	db "webbanhang/db/sqlc"
	"webbanhang/util"

	_ "github.com/lib/pq"
)

const (
	ADDRESS   = "0.0.0.0:8080"
	DB_NAME   = "webbanhang"
	DB_DRIVER = "postgres"
	DB_SOURCE = "postgresql://root:secret@localhost:5432/" + DB_NAME + "?sslmode=disable"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the DB. Err: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(store, config)
	if err != nil {
		log.Fatal("cannot create new server: ", err)
	}

	err = server.Start(ADDRESS)
	if err != nil {
		log.Fatal("cannot start server by address: ", err)
	}
}
