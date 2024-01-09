// package handlers

// import (
//     "encoding/json"
// 	"github.com/Bonittas/Music_Stream/music-backend/database"
// 	"github.com/Bonittas/Music_Stream/music-backend/models"
//     "net/http"
// )

// // UserHandlers contains HTTP handlers for user-related operations
// func UserHandlers(db *database.DB) http.HandlerFunc {
//     return func(w http.ResponseWriter, r *http.Request) {
//         switch r.Method {
//         case http.MethodGet:
//             getUser(db, w, r)
//         case http.MethodPost:
//             createUser(db, w, r)
//         // Add other HTTP methods (PUT, DELETE, etc.) for users
//         default:
//             http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//         }
//     }
// }

// func getUser(db *database.DB, w http.ResponseWriter, r *http.Request) {
//     // Retrieve user by ID from the database
//     userID := r.URL.Query().Get("id")
//     user, err := db.GetUserByID(userID)
//     if err != nil {
//         http.Error(w, "User not found", http.StatusNotFound)
//         return
//     }

//     jsonResponse(w, user, http.StatusOK)
// }

// func createUser(db *database.DB, w http.ResponseWriter, r *http.Request) {
//     var user models.User
//     if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
//         http.Error(w, "Invalid request payload", http.StatusBadRequest)
//         return
//     }

//     // Save user to the database
//     if err := db.CreateUser(&user); err != nil {
//         http.Error(w, "Error creating user", http.StatusInternalServerError)
//         return
//     }

//     jsonResponse(w, user, http.StatusCreated)
// }
