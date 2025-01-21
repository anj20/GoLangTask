package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// SeedAds()
	// Initialize database
	db := InitDB()
	defer db.Close()

	// Set up router
	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/ads", GetAdsHandler).Methods("GET")
	r.HandleFunc("/ads/click", PostAdClickHandler).Methods("POST")
	r.Use(IPMiddleware)

	// Start server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
