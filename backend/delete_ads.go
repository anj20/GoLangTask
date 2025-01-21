package main

import (
	"log"
)

// DeleteAllAds deletes all ads from the ads table
func DeleteAllAds() {
	// Initialize the database
	db := InitDB()
	defer db.Close()

	// Execute the DELETE query
	_, err := db.Exec("DELETE FROM ads")
	if err != nil {
		log.Printf("Failed to delete ads: %v\n", err)
	} else {
		log.Println("All ads have been deleted successfully")
	}
}
