package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
    "github.com/bonittas/music_back/db"
    "github.com/bonittas/music_back/models"
)

type PlaylistHandler struct {
	DB *sql.DB
}

func NewPlaylistHandler(db *sql.DB) *PlaylistHandler {
	return &PlaylistHandler{
		DB: db,
	}
}

func (h *PlaylistHandler) GetPlaylists(w http.ResponseWriter, r *http.Request) {
	playlists, err := models.GetPlaylists(h.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retrieve audio file URLs or paths
	for i := range playlists {
		// Replace "audioFiles/" with the appropriate directory or URL prefix
		playlists[i].AudioFile = "audioFiles/" + playlists[i].AudioFile
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(playlists)
}