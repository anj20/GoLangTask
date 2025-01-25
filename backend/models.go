package main

type Ad struct {
	ID        int    `json:"id"`
	ImageURL  string `json:"image_url"`
	TargetURL string `json:"target_url"`
	VideoTime string `json:"video_time"`
}

type AdClick struct {
	AdID      int    `json:"ad_id"`
	IPAddress string `json:"ip_address"`
	VideoTime string `json:"video_time"`
	TimeStamp string `json:"timestamp"`
}
