package main

import (
	"playlist/sql"
	"log"
	"net/http"

	"music_back/db"
    "github.com/bonittas/music_back/handlers"

	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to the database
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/your-database-name")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize the handlers
	playlistHandler := handlers.NewPlaylistHandler(db)

	// Create a new router
	router := mux.NewRouter()

	// Define the routes
	router.HandleFunc("/playlists", playlistHandler.GetPlaylists).Methods("GET")

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}