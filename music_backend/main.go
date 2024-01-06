package main

import (
	"log"
	"net/http"

	"github.com/Bonittas/Music_Stream/music-backend/database"
	"github.com/Bonittas/Music_Stream/music-backend/handlers"
)

func main() {
	// Initialize the database
	db := database.InitMySQL()

	// Set up routes and handlers
	http.HandleFunc("/playlist", handlers.GetPlaylist(db))
	http.HandleFunc("/current-song", handlers.GetCurrentSong(db))

	// Start the server
	log.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}