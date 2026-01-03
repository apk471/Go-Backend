package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":3001"
	// TODO: Update this to match your frontend's URL
	const allowedOrigin = "http://127.0.0.1:3000"

	http.HandleFunc("/simple", func(w http.ResponseWriter, r *http.Request) {
		// Enable CORS for simple request
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)

		fmt.Println("Simple Request received")
		fmt.Fprintf(w, "Plain Simple Request!")
	})

	http.HandleFunc("/preflight", func(w http.ResponseWriter, r *http.Request) {
		// Enable CORS
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "X-Custom-Header")

		// Handle Preflight (OPTIONS request)
		if r.Method == "OPTIONS" {
			fmt.Println("Preflight (OPTIONS) Request received")
			w.WriteHeader(http.StatusOK)
			return
		}

		fmt.Println("Preflight (Actual) Request received")
		fmt.Fprintf(w, "Preflight Request Success!")
	})

	fmt.Printf("Server starting on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
