package main

import (
	"net/http"

	"github.com/Raihanki/articlestream/cmd/api/handlers"
	"github.com/Raihanki/articlestream/internal/repositories"
)

func (cfg *ApiConfig) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlers.GetHome)
	mux.HandleFunc("GET /health", handlers.HandleHealth)
	mux.HandleFunc("GET /v1/err", handlers.HandleSampleError)

	//user
	userHandler := handlers.UserHandler{
		UserRepository: repositories.UserRepository{DB: cfg.DB},
	}
	mux.HandleFunc("POST /api/v1/users", userHandler.Store)
	mux.HandleFunc("GET /api/v1/users/{userId}", userHandler.Show)

	return mux
}
