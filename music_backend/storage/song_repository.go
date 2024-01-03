package storage

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// SaveSongFile saves the uploaded song file to the specified destination
func SaveSongFile(file *multipart.FileHeader, destination string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}

// GetSongFilePath returns the file path of the song file based on the provided filename
func GetSongFilePath(filename string) string {
	basePath := "C:/Music" // Replace with the actual base path for storing song files
	return filepath.Join(basePath, filename)
}