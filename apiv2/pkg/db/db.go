package db

import (
	// prevent the import from being removed
	// just expose functionality used under the hood
	"database/sql"
	"errors"
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
	// connection pool
	DB.SetMaxOpenConns(POOL)
	DB.SetMaxIdleConns(IDLE)
	err1 := createTables()
	if err != nil || err1 != nil {
		combinedErr := errors.Join(err, err1)
		log.Fatal(combinedErr)
	}
	//defer DB.Close()
}

// create tables for our DB
func createTables() error {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
	email TEXT NOT NULL unique,
	name TEXT NOT NULL,
	createdAt DATETIME NOT NULL,
	password TEXT NOT NULL
	) 
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		return err
	}

	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name TEXT NOT NULL unique,
    description TEXT NOT NULL,
	location TEXT NOT NULL,
    time DATETIME NOT NULL,
    user_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES user(id)
	);
	`
	// check if table exists else create it on start
	_, err = DB.Exec(createEventTable)
	if err != nil {
		return err
	}
	return nil
}
