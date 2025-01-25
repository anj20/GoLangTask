package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

// GetAdsHandler serves a list of ads along with their video time
func GetAdsHandler(w http.ResponseWriter, r *http.Request) {
	db := InitDB()
	defer db.Close()

	// Fetch ads from the ads table and get video_time from the ads table
	query := `
		SELECT 
			id, 
			image_url, 
			target_url, 
			video_time 
		FROM ads
	`

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Failed to fetch ads", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var ads []Ad
	for rows.Next() {
		var ad Ad
		if err := rows.Scan(&ad.ID, &ad.ImageURL, &ad.TargetURL, &ad.VideoTime); err != nil {
			http.Error(w, "Failed to parse ad data", http.StatusInternalServerError)
			return
		}
		ads = append(ads, ad)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ads)
}

func PostAds(w http.ResponseWriter, r *http.Request) {
	db := InitDB()
	defer db.Close()

	// Parse request body
	var ads []Ad
	if err := json.NewDecoder(r.Body).Decode(&ads); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("Error decoding request body: %v\n", err)
		return
	}

	// Insert ads into the database
	for _, ad := range ads {
		_, err := db.Exec(
			"INSERT INTO ads (image_url, target_url, video_time) VALUES (?, ?, ?)",
			ad.ImageURL, ad.TargetURL, ad.VideoTime,
		)
		if err != nil {
			http.Error(w, "Failed to insert ad", http.StatusInternalServerError)
			log.Printf("Error inserting ad: %v\n", err)
			return
		}
	}

	// Respond with success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Ads inserted successfully"})
}

func PostAdClickHandler(w http.ResponseWriter, r *http.Request) {
	db := InitDB()
	defer db.Close()

	var adClick AdClick
	if err := json.NewDecoder(r.Body).Decode(&adClick); err != nil {
		http.Error(w, "Invalid request body 213", http.StatusBadRequest)
		return
	}

	// Fetch the video time from the ads table using ad_id
	var videoTime string
	err := db.QueryRow("SELECT video_time FROM ads WHERE id = ?", adClick.AdID).Scan(&videoTime)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Ad not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to fetch video time", http.StatusInternalServerError)
		return
	}

	// Add the video time to the AdClick struct
	adClick.VideoTime = videoTime

	// Log the ad click metadata into the ad_clicks table
	_, err = db.Exec(
		"INSERT INTO ad_clicks (ad_id, ip_address, video_time, timestamp) VALUES (?, ?, ?, ?)",
		adClick.AdID, adClick.IPAddress, adClick.VideoTime, adClick.TimeStamp,
	)
	if err != nil {
		http.Error(w, "Failed to log ad click", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Ad click logged successfully"})
}

// GetAdClicksHandler serves a list of ad clicks
func GetAdClicksHandler(w http.ResponseWriter, r *http.Request) {
	db := InitDB()
	defer db.Close()

	// Query all ad clicks
	rows, err := db.Query("SELECT ad_id, ip_address, video_time, timestamp FROM ad_clicks")
	if err != nil {
		http.Error(w, "Failed to fetch ad clicks", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var adClicks []AdClick
	for rows.Next() {
		var adClick AdClick
		if err := rows.Scan(&adClick.AdID, &adClick.IPAddress, &adClick.VideoTime, &adClick.TimeStamp); err != nil {
			log.Printf("Error scanning row: %v", err) // Log the error
			http.Error(w, "Failed to parse ad click data", http.StatusInternalServerError)
			return
		}
		adClicks = append(adClicks, adClick)
	}

	// Return the ad clicks in JSON format
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(adClicks); err != nil {
		http.Error(w, "Failed to encode ad clicks", http.StatusInternalServerError)
	}
}

// DeleteAllAdsClicks deletes all ads from the ads table
func DeleteAllAdsClicksHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow DELETE method
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method. Use DELETE.", http.StatusMethodNotAllowed)
		return
	}

	// Initialize the database
	db := InitDB()
	defer db.Close()

	// Execute the DELETE query
	_, err := db.Exec("DELETE FROM ad_clicks")
	if err != nil {
		log.Printf("Failed to delete ad clicks: %v\n", err)
		http.Error(w, "Failed to delete ad clicks", http.StatusInternalServerError)
		return
	}

	log.Println("All ad clicks have been deleted successfully")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("All ad clicks have been deleted successfully"))
}

// DeleteAllAds deletes all ads from the ads table
func DeleteAllAdsHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow DELETE method
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method. Use DELETE.", http.StatusMethodNotAllowed)
		return
	}

	// Initialize the database
	db := InitDB()
	defer db.Close()

	// Execute the DELETE query
	_, err := db.Exec("DELETE FROM ads")
	if err != nil {
		log.Printf("Failed to delete ads: %v\n", err)
		http.Error(w, "Failed to delete ads", http.StatusInternalServerError)
		return
	}

	log.Println("All ads have been deleted successfully")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("All ads have been deleted successfully"))
}
