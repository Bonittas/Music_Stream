package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"os/exec"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	_ "github.com/go-sql-driver/mysql"
)

type Song struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Artist    string `json:"artist"`
	Duration  int    `json:"duration"`
	MusicFile string `json:"music_file"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/music")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/songs", getAllSongs).Methods("GET")
	router.HandleFunc("/api/songs/{id}", getSong).Methods("GET")
	router.HandleFunc("/api/songs", createSong).Methods("POST")
	router.HandleFunc("/api/songs/{id}", updateSong).Methods("PUT")
	router.HandleFunc("/api/songs/{id}", deleteSong).Methods("DELETE")

	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // Add your frontend URL here
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	handler := corsOptions.Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}

func playSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	row := db.QueryRow("SELECT music_file FROM song WHERE id = ?", id)
	var musicFile string
	err := row.Scan(&musicFile)
	if err != nil {
		handleError(w, http.StatusInternalServerError)
		return
	}

	http.ServeFile(w, r, musicFile)
}
func getAllSongs(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, artist, duration, music_file FROM song")
	if err != nil {
		handleError(w, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var songs []Song
	for rows.Next() {
		var song Song
		err := rows.Scan(&song.ID, &song.Title, &song.Artist, &song.Duration, &song.MusicFile)
		if err != nil {
			handleError(w, http.StatusInternalServerError)
			return
		}
		songs = append(songs, song)
	}

	json.NewEncoder(w).Encode(songs)
}

func getSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	row := db.QueryRow("SELECT id, title, artist, duration, music_file FROM song WHERE id = ?", id)
	var song Song
	err := row.Scan(&song.ID, &song.Title, &song.Artist, &song.Duration, &song.MusicFile)
	if err != nil {
		handleError(w, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(song)
}

func createSong(w http.ResponseWriter, r *http.Request) {
	var song Song
	err := json.NewDecoder(r.Body).Decode(&song)
	if err != nil {
		handleError(w, http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO song (title, artist, duration, music_file) VALUES (?, ?, ?, ?)",
		song.Title, song.Artist, song.Duration, song.MusicFile)
	if err != nil {
		handleError(w, http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	song.ID = int(id)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(song)
}

func updateSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var song Song
	err := json.NewDecoder(r.Body).Decode(&song)
	if err != nil {
		handleError(w, http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE song SET title=?, artist=?, duration=?, music_file=? WHERE id=?",
		song.Title, song.Artist, song.Duration, song.MusicFile, id)
	if err != nil {
		handleError(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	_, err := db.Exec("DELETE FROM song WHERE id=?", id)
	if err != nil {
		handleError(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleError(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}