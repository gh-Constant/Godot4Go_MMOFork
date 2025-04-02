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
		portStr = "8081" // Default if not set
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid PORT value: %s", portStr)
	}

	addr := fmt.Sprintf(":%d", port)

	// Simple handler for all paths
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Simple Handler: Received %s request for %s", r.Method, r.URL.Path)
		fmt.Fprintf(w, "Server is running!\n") // Respond with simple text
	})

	log.Printf("Starting SIMPLE server on %s", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Failed to start SIMPLE server: %v", err)
	}
}

// Remove other functions like loadConfig, coalescePaths, addHeaders etc. for this test
