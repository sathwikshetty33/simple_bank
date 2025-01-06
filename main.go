package main

import (
	_"github.com/lib/pq"
	"database/sql"
	"log"
	db "github.com/sathwikshetty33/golang_bank/db/sqlc"
	api "github.com/sathwikshetty33/golang_bank/api"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:7760@localhost:8000/simple_bank?sslmode=disable"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(":8080")
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}