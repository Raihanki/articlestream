package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func LoadDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("error when opening database connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	log.Println("successfully connected to database")

	return db
}
