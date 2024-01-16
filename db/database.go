package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}

	// Set maximum number of open connections
	db.SetMaxOpenConns(10)

	// Set maximum number of idle connections
	db.SetMaxIdleConns(5)

	// Set the maximum lifetime of a connection
	db.SetConnMaxLifetime(time.Hour)
}

func GetDB() *sql.DB {
	return db
}
