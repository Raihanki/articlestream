package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Raihanki/articlestream/internal/database"
	"github.com/joho/godotenv"
)

type ApiConfig struct {
	DB *sql.DB
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error when loading configuration file: %v", err)
	}

	db := database.LoadDB()
	configApi := ApiConfig{
		DB: db,
	}

	port := os.Getenv("APP_PORT")

	server := &http.Server{
		Addr:    ":" + port,
		Handler: configApi.Routes(),
	}

	fmt.Println("Server is running on port: ", port)
	server.ListenAndServe()
}
