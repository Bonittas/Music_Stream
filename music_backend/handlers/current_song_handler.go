package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Bonittas/Music_Stream/music-backend/database"
	"github.com/Bonittas/Music_Stream/music-backend/models"
)

// GetCurrentSong returns the currently playing song
func GetCurrentSong(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the currently playing song from the database or any other data source
		// For the sake of simplicity, let's assume we have a function called GetCurrentlyPlayingSong
		currentSong := GetCurrentlyPlayingSong(db)

		// Add the audio file URL or path to the currentSong object
		currentSong.Audio = "http://example.com/audio/song.mp3"

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

