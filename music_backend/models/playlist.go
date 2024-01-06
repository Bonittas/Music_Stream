package models

type Playlist struct {
	ID    int           `json:"id"`
	Name  string        `json:"name"`
	Songs []Song        `json:"songs"`
	// Add more playlist fields if needed
}