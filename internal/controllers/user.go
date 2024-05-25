package controllers

import (
	"fmt"

	"github.com/ankur-toko/muzz/internal/models"
	"github.com/ankur-toko/muzz/internal/repositories/match"
	"github.com/ankur-toko/muzz/internal/repositories/user"
	"github.com/asaskevich/govalidator"
)

var UserControllerInstance UserController

type Hasher interface {
	Hash(str string) string
}

type MyDummyHash struct {
}

func (h MyDummyHash) Hash(str string) string {
	return str
}

type UserController struct {
	uRepo    user.UserRepository
	mRepo    match.MatcherRepository
	hashAlgo Hasher
}

func (uC UserController) AddUser(u models.UserApiInput) (models.User, error) {
	// Do the validations here and respond back
	ok, err := govalidator.ValidateStruct(u)
	if err != nil {
		fmt.Println(err, ok)
		return models.User{}, err
	}
	savedUser, err := uC.uRepo.SaveUser(u.Email, u.Name, u.Gender, uC.hashAlgo.Hash(u.Password), u.Age, u.Lat, u.Lon)
	if err != nil {
		return models.User{}, fmt.Errorf("error: unable to store user: %v", err)
	}
	err2 := uC.mRepo.AddUser(savedUser.Id)
	if err2 != nil {
		return models.User{}, fmt.Errorf("error: unable to store user: %v", err)
	}
	return savedUser.User, nil
}
