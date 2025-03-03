package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// format msg to JSON object
func responseWithError(w http.ResponseWriter, code int, message string) {
	if code > 500 {
		log.Println("Responding with 5xx error: ", message)
	}
	//define struct for error message
	type errorResponse struct {
		Error string `json:"error"`
	}
	//JSON error looks like {"error": "message"}
	responseWithJSON(w, code, errorResponse{Error: message})
}

// Marshall Payload to Json web
func responseWithJSON(w http.ResponseWriter, code int, playload interface{}) {
	dat, err := json.Marshal(playload)
	if err != nil {
		log.Printf("Failed to marsahl JSON response: %v", playload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	// write response body
	w.Write(dat)
}
