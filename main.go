package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("GET /endpoint/", handleEndpoint)

	router.HandleFunc("GET /content/{path}/", handlePageContent)

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

func handlePageContent(w http.ResponseWriter, req *http.Request) {

	// Development only
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

	contentPath := req.PathValue("path")
	var foundContent string

	switch contentPath {
	case "home":
		foundContent = `<h2>Business Time</h2>
		<p>As in: It's</p>`
	case "services":
		foundContent = `<h3>Services</h3>
		<h5>Some services:</h5>
		<ul>
		<li>Service 1</li>
		<li>Service 2</li>
		<li>Service 3</li>
		</ul>`
	case "about":
		foundContent = `<h3>About Us</h3>
		<p>INSERT NAME HERE corporation was founded in...</p>`
	case "contact":
		foundContent = `<h3>Contact Us</h3>
		<p>Oh boy</p>
		<p>How woud I do a form here??</p>`
	}

	if len(foundContent) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(foundContent))

}
