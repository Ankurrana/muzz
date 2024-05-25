package controllers

import (
	"math"

	"github.com/ankur-toko/muzz/internal/models"
)

type MatchScorer interface {
	CalculateMatchScore(a, b models.User) float64
}

type AgeAndLocationScorer struct{}

func (s AgeAndLocationScorer) CalculateMatchScore(a, b models.User) float64 {
	score := float64(0.0)
	ageDiffScore := 15.0 - float64(mod(a.Age-b.Age))
	distanceScore := 100.0 - distanceBetween(a.Lat, a.Lon, b.Lat, b.Lon)
	score = ageDiffScore * distanceScore
	return score
}

func mod(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func distanceBetween(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371.0 // Radius of the Earth in kilometers

	// Convert latitude and longitude from degrees to radians
	lat1Rad := lat1 * math.Pi / 180
	lon1Rad := lon1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	lon2Rad := lon2 * math.Pi / 180

	// Haversine formula
	dlat := lat2Rad - lat1Rad
	dlon := lon2Rad - lon1Rad

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(dlon/2)*math.Sin(dlon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Distance in kilometers
	distance := R * c

	return distance
}
