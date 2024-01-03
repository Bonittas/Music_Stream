package recommendation

import (
	"math/rand"

	"github.com/bonittas/MUSIC_STREAM/music-backend/models"
)

// RecommendationService provides methods for retrieving song recommendations
type RecommendationService struct {
	songs []models.Song
}

// NewRecommendationService creates a new instance of RecommendationService
func NewRecommendationService(songs []models.Song) *RecommendationService {
	return &RecommendationService{
		songs: songs,
	}
}

// GetRecommendations returns a list of recommended songs
func (rs *RecommendationService) GetRecommendations() []models.Song {
	// Logic for retrieving recommended songs
	// Here's a simple example that returns a random subset of songs from the available list

	const numRecommendations = 5 // Number of recommended songs to return

	recommendations := make([]models.Song, 0, numRecommendations)
	totalSongs := len(rs.songs)

	// Generate random indices to select songs
	for i := 0; i < numRecommendations; i++ {
		randomIndex := rand.Intn(totalSongs)
		recommendations = append(recommendations, rs.songs[randomIndex])
	}

	return recommendations
}