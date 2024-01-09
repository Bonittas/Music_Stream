package handlers

import (
    "fmt"
    "net/http"
    "os"
    "path/filepath"
)

// Handle file upload
func UploadFile(w http.ResponseWriter, r *http.Request) {
    // Parse multipart form with file and other form data
    err := r.ParseMultipartForm(10 << 20) // 10 MB file size limit
    if err != nil {
        http.Error(w, "Error parsing form", http.StatusBadRequest)
        return
    }

    // Get the file from the request
    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Error retrieving file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    // Create the uploads directory if it doesn't exist
    uploadDir := "./storage/uploads"
    if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
        err := os.MkdirAll(uploadDir, 0755)
        if err != nil {
            http.Error(w, "Error creating upload directory", http.StatusInternalServerError)
            return
        }
    }

    // Save the uploaded file to the uploads directory
    filePath := filepath.Join(uploadDir, handler.Filename)
    f, err := os.Create(filePath)
    if err != nil {
        http.Error(w, "Error saving file", http.StatusInternalServerError)
        return
    }
    defer f.Close()

    // Copy the file data to the destination
    _, err = io.Copy(f, file)
    if err != nil {
        http.Error(w, "Error copying file", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "File %s uploaded successfully", handler.Filename)
}
