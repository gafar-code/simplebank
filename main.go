package main

import (
	"database/sql"
	"log"

	"github.com/gafar-code/simplebank/api"
	db "github.com/gafar-code/simplebank/db/sqlc"
	"github.com/gafar-code/simplebank/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Tidak dapat memuat config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DbSource)
	if err != nil {
		log.Fatal("Tidak bisa terhubung ke database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	server.Start(config.ServerAddress)
}
