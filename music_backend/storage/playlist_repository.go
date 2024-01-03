package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Playlist struct represents a playlist
type Playlist struct {
	ID   string   `json:"id"`
	Name string   `json:"name"`
	Songs []string `json:"songs"`
}

// SavePlaylist saves the playlist as a JSON file
func SavePlaylist(playlist *Playlist, destination string) error {
	data, err := json.MarshalIndent(playlist, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(destination, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// LoadPlaylist loads a playlist from a JSON file
func LoadPlaylist(file string) (*Playlist, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var playlist Playlist
	err = json.Unmarshal(data, &playlist)
	if err != nil {
		return nil, err
	}

	return &playlist, nil
}

// GetPlaylistFilePath returns the file path of the playlist file based on the provided filename
func GetPlaylistFilePath(filename string) string {
	basePath := "C:/Music" // Replace with the actual base path for storing playlist files
	return filepath.Join(basePath, filename)
}