package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Bonittas/Music_Stream/music-backend/database"
	"github.com/Bonittas/Music_Stream/music-backend/models"
)

// GetUserPlaylist retrieves the user's playlist from the database or any other data source
// Implement the logic to fetch the user's playlist
// Return the user's playlist as a slice of models.Playlist
func GetUserPlaylist(db *database.DB) []models.Playlist {
	// Implement the logic here
	return []models.Playlist{
		{
			ID:   1,
			Name: "My Playlist",
			Songs: []models.Song{
				{
					ID:    1,
					Title: "Song 1",
					Audio: "http://example.com/audio/song1.mp3",
				},
				{
					ID:    2,
					Title: "Song 2",
					Audio: "http://example.com/audio/song2.mp3",
				},
			},
		},
	}
}

// GetPlaylist returns the user's playlist
func GetPlaylist(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the user's playlist from the database
		playlist := GetUserPlaylist(db)

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