package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Raihanki/articlestream/internal/helpers"
)

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status string `json:"status"`
	}

	jsonResponse, err := json.Marshal(Response{Status: "OK"})
	if err != nil {
		w.WriteHeader(500)
		log.Fatalf("error when marshaling json response: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonResponse)
}

func HandleSampleError(w http.ResponseWriter, r *http.Request) {
	helpers.JsonResponse(w, 500, nil, "")
}
