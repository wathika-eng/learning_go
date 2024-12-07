package db

import (
	// prevent the import from being removed
	// just expose functionality used under the hood
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// DB Global (a pointer to the SQL DB)
var DB *sql.DB

const POOL = 10
const IDLE = 3

func InitDB() {
	var err error
	// sql.Open(driverName, path_to_the_db ending with .db)
	DB, err = sql.Open("sqlite3", "api.db")
	// if we get an error connecting to the DB, just crash the app
	if err != nil {
		log.Fatal(err)
	}
	// connection pool
	DB.SetMaxOpenConns(POOL)
	DB.SetMaxIdleConns(IDLE)
}
