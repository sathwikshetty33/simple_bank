package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/sathwikshetty33/golang_bank/db/util"
)

var testQueries *Queries
var testDB *sql.DB


func TestMain(m *testing.M) {
	
	config,err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load configurations: ", err)
	}
	testDB, err = sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}