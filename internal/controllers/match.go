package controllers

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ankur-toko/muzz/internal/models"
	"github.com/ankur-toko/muzz/internal/repositories/match"
	userRepo "github.com/ankur-toko/muzz/internal/repositories/user"
	"github.com/asaskevich/govalidator"
)

type MatchController struct {
	mRepo  match.MatcherRepository
	uRepo  userRepo.UserRepository
	scorer MatchScorer
}

func (mC *MatchController) Discover(userId int, fromAge int, toAge int, gender string) ([]models.DiscoveredUser, error) {
	users, err := mC.mRepo.Discover(userId)
	if err != nil {
		return nil, err
	}

	user, err := mC.uRepo.GetUser(userId)
	if err != nil {
		return nil, err
	}

	dbUsers, err := mC.uRepo.GetUsersByFilter(users, fromAge, toAge, gender)
	if err != nil {
		return nil, err
	}
	res := []models.DiscoveredUser{}
	for i := 0; i < len(dbUsers); i++ {
		if dbUsers[i].Id != userId {
			distance := distanceBetween(user.Lat, user.Lon, dbUsers[i].Lat, dbUsers[i].Lon)
			if distance <= 0 {
				distance = 0.1
			}
			score := mC.scorer.CalculateMatchScore(user.User, dbUsers[i].User)
			res = append(res, models.DiscoveredUser{User: dbUsers[i].User, Id: dbUsers[i].Id, DistanceFromMe: distance, Score: score})
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].Score > res[j].Score
	})

	return res, nil
}

func (mC *MatchController) Swipe(userId int, swipeInput models.SwipeApiInput) (models.SwipeApiResponse, error) {
	ok, err := govalidator.ValidateStruct(swipeInput)
	if !ok {
		return models.SwipeApiResponse{}, err
	}

	if userId == swipeInput.UserId {
		return models.SwipeApiResponse{}, fmt.Errorf("self right swipe is unnecessary! :D ")
	}

	if strings.ToLower(swipeInput.Preference) == "yes" {
		matched, match_id, err := mC.mRepo.Swipe(userId, swipeInput.UserId)
		if err != nil {
			return models.SwipeApiResponse{}, err
		}
		if matched {
			return models.SwipeApiResponse{
				Matched: true,
				MatchId: match_id,
			}, nil
		}
	}
	return models.SwipeApiResponse{
		Matched: false,
		MatchId: -1,
	}, nil
}
