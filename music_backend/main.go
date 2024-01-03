package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/bonittas/MUSIC_STREAM/music-backend/handlers"
	"github.com/bonittas/MUSIC_STREAM/music-backend/database"
)

func main() {
	db := database.InitMySQL() // Initialize MySQL database connection

	// Create a new instance of the Gorilla mux router
	router := mux.NewRouter()

	// Register API routes
	router.HandleFunc("/api/currentSong", handlers.GetCurrentSong(db)).Methods("GET")
	router.HandleFunc("/api/playlist", handlers.GetPlaylist(db)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}