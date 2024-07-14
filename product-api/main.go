package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Handles requests to the path /goodbye
	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World")
	})

	// Handles all other requests
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Running Hello Handler")

		// Read the body
		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading body", err)
			http.Error(rw, "Unable to read request body", http.StatusBadRequest)
			return
		}

		// Write the response
		fmt.Fprintf(rw, "Hello %s", b)
	})

	// Listen for connections on all IP addresses (0.0.0.0) on port 9090
	log.Println("Starting Server")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
