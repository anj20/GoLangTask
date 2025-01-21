package main

import (
	"log"
)


// SeedAds seeds the ads table with predefined ads
func SeedAds() {
	db := InitDB()
	defer db.Close()

	ads := []Ad{
		{ImageURL: "https://example.com/ad1.jpg", TargetURL: "https://example.com/product1"},
		{ImageURL: "https://example.com/ad2.jpg", TargetURL: "https://example.com/product2"},
		{ImageURL: "https://example.com/ad3.jpg", TargetURL: "https://example.com/product3"},
		{ImageURL: "https://example.com/ad4.jpg", TargetURL: "https://example.com/product4"},
	}

	for _, ad := range ads {
		_, err := db.Exec("INSERT INTO ads (image_url, target_url) VALUES (?, ?)", ad.ImageURL, ad.TargetURL)
		if err != nil {
			log.Printf("Failed to insert ad: %v\n", err)
		} else {
			log.Printf("Ad inserted: %v\n", ad)
		}
	}
}
