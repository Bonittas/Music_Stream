package models

type Song struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Audio string `json:"audio"`
	// Add more song fields if needed, such as artist, album, duration, etc.
}