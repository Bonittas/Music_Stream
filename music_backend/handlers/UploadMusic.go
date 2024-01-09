// package handlers
// import (
//     "fmt"
//     "io"
//     "github.com/Bonittas/Music_Stream/music-backend/utils"

//     "encoding/json"

//     "os"
//     "net/http"
// 	"github.com/Bonittas/Music_Stream/music-backend/database"
// 	"github.com/Bonittas/Music_Stream/music-backend/models"    // Import other necessary packages/utils
// )
// // Handle the music file upload
// func UploadMusic(db *database.DB) http.HandlerFunc {
//     return func(w http.ResponseWriter, r *http.Request) {
//         // Parse the multipart form file
//         err := r.ParseMultipartForm(10 << 20) // 10MB max file size
//         if err != nil {
//             utils.RespondWithError(w, http.StatusBadRequest, err.Error())
//             return
//         }

//         file, handler, err := r.FormFile("file")
//         if err != nil {
//             utils.RespondWithError(w, http.StatusBadRequest, err.Error())
//             return
//         }
//         defer file.Close()

//         // Example: Check file type
//         if !isAllowedFileType(handler.Filename) {
//             utils.RespondWithError(w, http.StatusBadRequest, "Invalid file type")
//             return
//         }

//         // Save the file to a specified location (e.g., /path/to/music)
//         f, err := os.OpenFile("/path/to/music/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
//         if err != nil {
//             utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
//             return
//         }
//         defer f.Close()
//         io.Copy(f, file)

//         // Store metadata in the database (e.g., title, file path)
//         // Replace this with your actual database insertion logic
//         title := r.FormValue("title")
//         filePath := "/path/to/music/" + handler.Filename
//         // db.Exec("INSERT INTO songs (title, file_path) VALUES (?, ?)", title, filePath)

//         utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Music uploaded successfully"})
//     }
// }

// // Function to check allowed file types (example)
// func isAllowedFileType(filename string) bool {
//     // Implement logic to check file extensions/types here
//     // Example: return strings.HasSuffix(strings.ToLower(filename), ".mp3")
//     return true // Placeholder for demo
// }
