package models

import (
	"github.com/Bonittas/Music_Stream/music-backend/audio"
)

type Playlist struct {
	ID    int           `json:"id"`
	Name  string        `json:"name"`
	Songs []audio.Song  `json:"songs"`
	// Add more playlist fields if needed
}