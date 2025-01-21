package main

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./ads.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create tables
	createAdTable := `
		CREATE TABLE IF NOT EXISTS ads (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			image_url TEXT NOT NULL,
			target_url TEXT NOT NULL,
			video_time TEXT
		);
	`
	createClickTable := `
		CREATE TABLE IF NOT EXISTS ad_clicks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			ad_id INTEGER,
			timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
			ip_address TEXT,
			video_time TEXT
		);
	`

	_, err = db.Exec(createAdTable)
	if err != nil {
		log.Fatalf("Failed to create ads table: %v", err)
	}

	_, err = db.Exec(createClickTable)
	if err != nil {
		log.Fatalf("Failed to create ad_clicks table: %v", err)
	}

	return db
}
