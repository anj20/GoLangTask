package main

import (
	"log"
)

// SeedAds seeds the ads table with predefined ads
func SeedAds() {
	db := InitDB()
	defer db.Close()

	ads := []Ad{
		{
			ImageURL:  "https://www.coca-cola.com/content/dam/journey/us/en/private/2023/coca-cola-logo.png",
			TargetURL: "https://www.youtube.com/watch?v=2msbfN81Gm0",
			VideoTime: "00:15",
		},
		{
			ImageURL:  "https://upload.wikimedia.org/wikipedia/commons/a/a6/Logo_NIKE.svg",
			TargetURL: "https://www.youtube.com/watch?v=Q59H5C89JkA",
			VideoTime: "00:20",
		},
		{
			ImageURL:  "https://upload.wikimedia.org/wikipedia/commons/f/fa/Apple_logo_black.svg",
			TargetURL: "https://www.youtube.com/watch?v=Jb4FIuVjFOo",
			VideoTime: "00:30",
		},
		{
			ImageURL:  "https://upload.wikimedia.org/wikipedia/commons/4/4e/McDonald%27s_Golden_Arches.svg",
			TargetURL: "https://www.youtube.com/watch?v=GnPJVvvm5oQ",
			VideoTime: "00:45",
		},
		{
			ImageURL:  "https://upload.wikimedia.org/wikipedia/commons/2/24/Samsung_Logo.svg",
			TargetURL: "https://www.youtube.com/watch?v=6afINYkRxTk",
			VideoTime: "01:00",
		},
	}

	for _, ad := range ads {
		_, err := db.Exec("INSERT INTO ads (image_url, target_url, video_time) VALUES (?, ?, ?)", ad.ImageURL, ad.TargetURL,ad.VideoTime)
		if err != nil {
			log.Printf("Failed to insert ad: %v\n", err)
		} else {
			log.Printf("Ad inserted: %v\n", ad)
		}
	}
}
