package controllers

import (
	"testing"

	"github.com/ankur-toko/muzz/internal/models"
)

func Test_distanceBetween(t *testing.T) {
	type args struct {
		lat1 float64
		lon1 float64
		lat2 float64
		lon2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"simple", args{30.172705, 31.526725, 30.288281, 31.732326}, 23.576},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distanceBetween(tt.args.lat1, tt.args.lon1, tt.args.lat2, tt.args.lon2); !(mod(got-tt.want) < float64(0.5)) {
				t.Errorf("distanceBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAgeAndLocationScorer_CalculateMatchScore(t *testing.T) {
	var a, b, c, d, e models.User
	a, b, c, d, e = models.User{}, models.User{}, models.User{}, models.User{}, models.User{}
	a.Age = 10
	a.Lat = 10.0
	a.Lon = 10.0

	b.Age = 12
	b.Lat = 10.0
	b.Lon = 10.0

	c.Age = 14
	c.Lat = 10.0
	c.Lon = 10.0

	d.Age = 14
	d.Lat = 12.0
	d.Lon = 12.0

	e.Age = 14
	e.Lat = 14.0
	e.Lon = 14.0

	scorer := AgeAndLocationScorer{}

	ok1 := scorer.CalculateMatchScore(a, b) > scorer.CalculateMatchScore(a, c)
	if !ok1 {
		t.Errorf("people with lesser age gap must have better score ")
	}

	ok2 := scorer.CalculateMatchScore(c, d) > scorer.CalculateMatchScore(c, e)

	if !ok2 {
		t.Errorf("people living closer must have better score")
	}

}
