package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func JsonResponse(w http.ResponseWriter, status int, data interface{}, message string) {
	if status < 100 || status > 500 {
		w.WriteHeader(500)
		log.Fatalf("invalid status: %d", status)
		return
	}

	if status == 500 {
		message = "Internal Server Error"
	}

	response := Response{
		Status:  status,
		Data:    data,
		Message: message,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(500)
		log.Fatalf("error while marshaling json response: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonResponse)
}
