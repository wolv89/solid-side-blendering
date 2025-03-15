package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("GET /endpoint", handleEndpoint)

	serverURL := "blewb.build:8338"

	server := &http.Server{
		Handler: router,
		Addr:    serverURL,
	}

	fmt.Printf("Server started at:\nhttp://%s\n\n", serverURL)
	server.ListenAndServe()

}

type jsonDemo struct {
	Message string `json:"message"`
	Value   int    `json:"value"`
	Valid   bool   `json:"valid"`
}

func handleEndpoint(w http.ResponseWriter, req *http.Request) {

	// Development only
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	endData := jsonDemo{
		Message: "As I type this the practice session for the 2025 Aus F1 is starting...",
		Value:   15325,
		Valid:   true,
	}

	data, err := json.Marshal(endData)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		return
	}

	w.Write(data)

}
