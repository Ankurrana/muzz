package user

import "github.com/ankur-toko/muzz/internal/models"

type UserRepository interface {
	SaveUser(email, name, gender, password string, age int, lat, log float64) (models.DBUser, error)
	GetUser(id int) (models.DBUser, error)
	GetUserByEmail(email string) (models.DBUser, error)
	GetUsers(id []int) ([]models.DBUser, error)
	GetUsersByFilter(id []int, fromAge, toAge int, gender string) ([]models.DBUser, error)
}
