package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func sendErrorResponse(w http.ResponseWriter, code int, message string) {
	if code >= 500 {
		log.Printf("Server error (%d): %s", code, message)
	}

	response := ErrorResponse{Message: message}
	sendJSONResponse(w, code, response)
}

func sendJSONResponse(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
