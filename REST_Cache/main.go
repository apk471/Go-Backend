package main

import (
	"log"
	"net/http"

	"task-manager/cache"
	"task-manager/db"
	"task-manager/handlers"
)

func main() {
	// Connect to MongoDB
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	// Connect to Redis
	if err := cache.Connect(); err != nil {
		log.Fatal(err)
	}

	log.Println("MongoDB & Redis connected")

	http.HandleFunc("/organizations", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.ListOrganizationsHandler(w, r)
		case http.MethodPost:
			handlers.CreateOrganizationHandler(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/organizations/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetOrganizationByIDHandler(w, r)
		case http.MethodPut:
			handlers.UpdateOrganizationHandler(w, r)
		case http.MethodDelete:
			handlers.DeleteOrganizationHandler(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
