package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	// Construct MySQL configuration
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Net:                  os.Getenv("DB_NET"),
		Addr:                 os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}

	// Use the global DB variable
	var err error
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to ensure connection is established
	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	log.Println("DB connection successful")
}
