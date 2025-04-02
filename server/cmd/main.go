package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// Basic port configuration directly from ENV
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080" // Default if not set
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid PORT value: %s", portStr)
	}

	addr := fmt.Sprintf(":%d", port)

	// Simple handler for all paths
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Simple Handler: Received %s request for %s", r.Method, r.URL.Path)
		_, err := fmt.Fprintf(w, "Server is running!\n") // Respond with simple text
		if err != nil {
			log.Printf("Simple Handler: ERROR writing response: %v", err)
		} else {
			log.Printf("Simple Handler: Successfully wrote response for %s", r.URL.Path)
		}
	})

	log.Printf("Starting SIMPLE server on %s", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Failed to start SIMPLE server: %v", err)
	}
}

// Remove other functions like loadConfig, coalescePaths, addHeaders etc. for this test
