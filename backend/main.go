package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// SeedAds()
	// Initialize database
	db := InitDB()
	defer db.Close()

	// DeleteAllAds()
	// Set up router
	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/ads", GetAdsHandler).Methods("GET")
	r.HandleFunc("/ads/click", PostAdClickHandler).Methods("POST")
	r.HandleFunc("/ads/click", GetAdClicksHandler).Methods("GET")
	r.Use(IPMiddleware)

	// Enable CORS
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"}, // Frontend URL
        AllowedMethods:   []string{"GET", "POST", "OPTIONS"}, // Allowed HTTP methods
        AllowedHeaders:   []string{"Content-Type", "Authorization"},
        AllowCredentials: true, // If cookies or credentials are sent
    })

	// Wrap the router with the CORS middleware
    handler := c.Handler(r)


	// Start server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
