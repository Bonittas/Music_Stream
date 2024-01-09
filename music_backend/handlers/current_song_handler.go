// package handlers

// import (
//     "encoding/json"
// 	"github.com/Bonittas/Music_Stream/music-backend/database"
// 	"github.com/Bonittas/Music_Stream/music-backend/models"
//     "net/http"
// )

// // SongHandlers contains HTTP handlers for song-related operations
// func SongHandlers(db *database.DB) http.HandlerFunc {
//     return func(w http.ResponseWriter, r *http.Request) {
//         switch r.Method {
//         case http.MethodGet:
//             getSong(db, w, r)
//         case http.MethodPost:
//             createSong(db, w, r)
//         // Add other HTTP methods (PUT, DELETE, etc.) for songs
//         default:
//             http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//         }
//     }
// }

// func getSong(db *database.DB, w http.ResponseWriter, r *http.Request) {
//     // Retrieve song by ID from the database
//     songID := r.URL.Query().Get("id")
//     song, err := db.GetSongByID(songID)
//     if err != nil {
//         http.Error(w, "Song not found", http.StatusNotFound)
//         return
//     }

//     jsonResponse(w, song, http.StatusOK)
// }

// func createSong(db *database.DB, w http.ResponseWriter, r *http.Request) {
//     var song models.Song
//     if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
//         http.Error(w, "Invalid request payload", http.StatusBadRequest)
//         return
//     }

//     // Save song to the database
//     if err := db.CreateSong(&song); err != nil {
//         http.Error(w, "Error creating song", http.StatusInternalServerError)
//         return
//     }

//     jsonResponse(w, song, http.StatusCreated)
// }
