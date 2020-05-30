package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

// Mood is a new mood entry in the DB
type Mood struct {
	Score     int    `json:"score"`
	DeviceID  string `json:"deviceID"`
	CreatedAt time.Time
}

// SaveMood writes a new mood to the database
func SaveMood(w http.ResponseWriter, r *http.Request) {
	var mood Mood

	err := json.NewDecoder(r.Body).Decode(&mood)
	if err != nil {
		log.Println("error parsing mood data", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mood.CreatedAt = time.Now()

	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println("error establishing DB connection", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = db.Ping(context.Background())
	if err != nil {
		log.Println("error pinging DB connection", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = db.Exec(context.Background(),
		"INSERT INTO moods (score, device_id, created_at) VALUES ($1, $2, $3)", mood.Score, mood.DeviceID, mood.CreatedAt)

	if err != nil {
		log.Println("error creating mood report", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println("success! created mood report")
	w.WriteHeader(201)
}
