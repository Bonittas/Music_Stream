package models
import (
	"database/sql"
)
type Playlist struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	AudioFile string `json:audioFile`
}

func GetPlaylists(db *sql.DB) ([]Playlist, error) {
	rows, err := db.Query("SELECT * FROM playlists")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var playlists []Playlist

	for rows.Next() {
		var playlist Playlist
		err := rows.Scan(&playlist.ID, &playlist.Name)
		if err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)
	}

	return playlists, nil
}