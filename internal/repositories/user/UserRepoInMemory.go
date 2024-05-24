package user

import (
	"fmt"

	"github.com/ankur-toko/muzz/internal/models"
)

type UserRepoInMemory struct {
	UserMap map[int]models.DBUser
}

var instance *UserRepoInMemory

func Instance() *UserRepoInMemory {
	if instance == nil {
		instance = &UserRepoInMemory{
			make(map[int]models.DBUser),
		}
	}
	return instance
}

func (u *UserRepoInMemory) SaveUser(email, name, gender, password string, age int, lat, lon float64) (models.DBUser, error) {
	id := len(u.UserMap) + 1
	user := models.DBUser{
		Id:       id,
		Password: password,
		User: models.User{
			Name:   name,
			Email:  email,
			Gender: gender,
			Age:    age,
			Lat:    lat,
			Lon:    lon,
		},
	}

	u.UserMap[id] = user
	return user, nil
}

func (u *UserRepoInMemory) GetUser(id int) (models.DBUser, error) {
	user, ok := u.UserMap[id]
	if !ok {
		return models.DBUser{}, fmt.Errorf("error: user not found")
	}

	return user, nil
}

func (u *UserRepoInMemory) GetUserByEmail(email string) (models.DBUser, error) {
	for _, u := range u.UserMap {
		if u.Email == email {
			return u, nil
		}
	}
	return models.DBUser{}, nil
}
func (u *UserRepoInMemory) GetUsers(ids []int) ([]models.DBUser, error) {
	res := []models.DBUser{}

	for i := 0; i < len(ids); i++ {
		if user, ok := u.UserMap[ids[i]]; ok {
			res = append(res, user)
		}
	}

	return res, nil
}

func (u *UserRepoInMemory) GetUsersByFilter(ids []int, fromAge, toAge int, gender string) ([]models.DBUser, error) {
	res := []models.DBUser{}
	for i := 0; i < len(ids); i++ {
		if user, ok := u.UserMap[ids[i]]; ok {
			if user.Age >= fromAge && user.Age <= toAge && (gender == "*" || user.Gender == gender) {
				res = append(res, user)
			}
		}
	}

	return res, nil
}

func (u *UserRepoInMemory) AuthenticateUser(email string, passwordHash string) (bool, error) {
	fmt.Println("authenticating user", email, passwordHash)
	return true, nil
}

func (u *UserRepoInMemory) SearchUsers(lat, lon float64, offset int, count int) ([]models.DBUser, error) {
	fmt.Println("getting users", lat, lon, offset, count)
	return []models.DBUser{}, nil
}
