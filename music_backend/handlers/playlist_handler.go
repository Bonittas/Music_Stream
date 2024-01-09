// package handlers

// import (
//     "encoding/json"
// 	"github.com/Bonittas/Music_Stream/music-backend/database"
// 	"github.com/Bonittas/Music_Stream/music-backend/models"
//     "net/http"
// )

// // PlaylistHandlers contains HTTP handlers for playlist-related operations
// func PlaylistHandlers(db *database.DB) http.HandlerFunc {
//     return func(w http.ResponseWriter, r *http.Request) {
//         switch r.Method {
//         case http.MethodGet:
//             getPlaylist(db, w, r)
//         case http.MethodPost:
//             createPlaylist(db, w, r)
//         // Add other HTTP methods (PUT, DELETE, etc.) for playlists
//         default:
//             http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//         }
//     }
// }

// func getPlaylist(db *database.DB, w http.ResponseWriter, r *http.Request) {
//     // Retrieve playlist by ID from the database
//     playlistID := r.URL.Query().Get("id")
//     playlist, err := db.GetPlaylistByID(playlistID)
//     if err != nil {
//         http.Error(w, "Playlist not found", http.StatusNotFound)
//         return
//     }

//     jsonResponse(w, playlist, http.StatusOK)
// }

// func createPlaylist(db *database.DB, w http.ResponseWriter, r *http.Request) {
//     var playlist models.Playlist
//     if err := json.NewDecoder(r.Body).Decode(&playlist); err != nil {
//         http.Error(w, "Invalid request payload", http.StatusBadRequest)
//         return
//     }

//     // Save playlist to the database
//     if err := db.CreatePlaylist(&playlist); err != nil {
//         http.Error(w, "Error creating playlist", http.StatusInternalServerError)
//         return
//     }

//     jsonResponse(w, playlist, http.StatusCreated)
// }
