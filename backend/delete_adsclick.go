package main

import (
	"log"
)

// DeleteAllAds deletes all ads from the ads table
func DeleteAllAdsClicks() {
	// Initialize the database
	db := InitDB()
	defer db.Close()

	// Execute the DELETE query
	_, err := db.Exec("DELETE FROM ad_clicks")
	if err != nil {
		log.Printf("Failed to delete ad clicks: %v\n", err)
	} else {
		log.Println("All ad clicks have been deleted successfully")
	}
}
