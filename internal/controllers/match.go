package controllers

import (
	"strings"

	"github.com/ankur-toko/muzz/internal/models"
	"github.com/ankur-toko/muzz/internal/repositories/match"
	userRepo "github.com/ankur-toko/muzz/internal/repositories/user"
	"github.com/asaskevich/govalidator"
)

type MatchController struct {
	mRepo match.MatcherRepository
	uRepo userRepo.UserRepository
}

var matchController *MatchController = &MatchController{match.Instance(), userRepo.Instance()}

func MatchControllerInstance() *MatchController {
	return matchController
}

func (mC *MatchController) Discover(userId int, fromAge int, toAge int, gender string) ([]models.User, error) {
	users, err := mC.mRepo.Discover(userId)
	if err != nil {
		return nil, err
	}
	dbUsers, err := mC.uRepo.GetUsersByFilter(users, fromAge, toAge, gender)
	if err != nil {
		return nil, err
	}
	res := []models.User{}
	for i := 0; i < len(dbUsers); i++ {
		if dbUsers[i].Id != userId {
			res = append(res, dbUsers[i].User)
		}
	}
	return res, nil
}

func (mC *MatchController) Swipe(userId int, swipeInput models.SwipeApiInput) (models.SwipeApiResponse, error) {
	ok, err := govalidator.ValidateStruct(swipeInput)
	if !ok {
		return models.SwipeApiResponse{}, err
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
