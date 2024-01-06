package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Bonittas/Music_Stream/music-backend/database"
	"github.com/Bonittas/Music_Stream/music-backend/models"
)

// GetCurrentlyPlayingSong retrieves the currently playing song from the database or any other data source
// Implement the logic to fetch the currently playing song
// Return the currently playing song as an instance of models.Song
func GetCurrentlyPlayingSong(db *database.DB) models.Song {
	// Implement the logic here
	return models.Song{
		ID:    1,
		Title: "Currently Playing Song",
		Audio: "http://example.com/audio/song.mp3",
	}
}

// GetCurrentSong returns the currently playing song
func GetCurrentSong(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the currently playing song from the database
		currentSong := GetCurrentlyPlayingSong(db)

		// Convert the currentSong object to JSON
		response, err := json.Marshal(currentSong)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the response headers
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Write the response
		_, err = w.Write(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}