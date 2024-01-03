package handlers

import (
	"encoding/json"
	"net/http"
// hellooooo
	"github.com/Bonittas/Music_Stream/music-backend/database"
	"github.com/Bonittas/Music_Stream/music-backend/models"
)

// GetPlaylist returns the user's playlist
func GetPlaylist(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the user's playlist from the database or any other data source
		// For the sake of simplicity, let's assume we have a function called GetUserPlaylist
		playlist := GetUserPlaylist(db)

		// Add the audio file URLs or paths to each song in the playlist
		for i := range playlist {
			playlist[i].Audio = "http://example.com/audio/song" + strconv.Itoa(playlist[i].ID) + ".mp3"
		}

		// Convert the playlist object to JSON
		response, err := json.Marshal(playlist)
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