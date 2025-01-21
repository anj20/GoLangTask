package main

import (
	"encoding/json"
	"net/http"
)

// GetAdsHandler serves a list of ads
func GetAdsHandler(w http.ResponseWriter, r *http.Request) {
	db := InitDB()
	defer db.Close()

	rows, err := db.Query("SELECT id, image_url, target_url FROM ads")
	if err != nil {
		http.Error(w, "Failed to fetch ads", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var ads []Ad
	for rows.Next() {
		var ad Ad
		if err := rows.Scan(&ad.ID, &ad.ImageURL, &ad.TargetURL); err != nil {
			http.Error(w, "Failed to parse ad data", http.StatusInternalServerError)
			return
		}
		ads = append(ads, ad)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ads)
}

// PostAdClickHandler logs ad click metadata
func PostAdClickHandler(w http.ResponseWriter, r *http.Request) {
	db := InitDB()
	defer db.Close()

	var adClick AdClick
	if err := json.NewDecoder(r.Body).Decode(&adClick); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ip := r.Context().Value("ip").(string)
	adClick.IPAddress = ip

	_, err := db.Exec(
		"INSERT INTO ad_clicks (ad_id, ip_address, video_time) VALUES (?, ?, ?)",
		adClick.AdID, adClick.IPAddress, adClick.VideoTime,
	)
	if err != nil {
		http.Error(w, "Failed to log ad click", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Ad click logged successfully"})
}
