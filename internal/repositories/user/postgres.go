package user

import (
	"github.com/ankur-toko/muzz/internal/models"
	"github.com/ankur-toko/muzz/internal/repositories/sql"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type UserPostgresRepo struct {
	Conn *gorm.DB
}

func (db *UserPostgresRepo) SaveUser(email, name, gender, password string, age int, lat, lon float64) (models.DBUser, error) {
	user := sql.User{
		Name:         name,
		Email:        email,
		Gender:       gender,
		PasswordHash: password,
		Age:          age,
		Lat:          lat,
		Lon:          lon,
	}
	res := db.Conn.Create(&user)
	if res.Error != nil {
		return models.DBUser{}, res.Error
	}
	u := mapSQLUserToDBUser(user)
	return u, nil
}

func (db *UserPostgresRepo) GetUser(id int) (models.DBUser, error) {
	var user sql.User
	res := db.Conn.First(&user, id)
	if res.Error != nil {
		return models.DBUser{}, res.Error
	}
	return mapSQLUserToDBUser(user), nil
}

func (db *UserPostgresRepo) GetUserByEmail(email string) (models.DBUser, error) {
	var user sql.User
	res := db.Conn.Where("email = ?", email).First(&user)
	if res.Error != nil {
		return models.DBUser{}, res.Error
	}
	return mapSQLUserToDBUser(user), nil
}

func (db *UserPostgresRepo) GetUsers(id []int) ([]models.DBUser, error) {
	users := []sql.User{}
	respose := db.Conn.Where(id).Find(&users)
	res := []models.DBUser{}
	if respose.Error != nil {
		return res, respose.Error
	}
	for i := 0; i < len(users); i++ {
		res = append(res, mapSQLUserToDBUser(users[i]))
	}
	return res, nil
}

func (db *UserPostgresRepo) GetUsersByFilter(id []int, fromAge, toAge int, gender string) ([]models.DBUser, error) {
	users := []sql.User{}
	inter := db.Conn

	inter = inter.Where("id in ? ", id)
	if fromAge > 0 {
		inter = inter.Where("age >= ?", fromAge)
	}
	if toAge < 100 {
		inter = inter.Where("age <= ?", toAge)
	}
	if gender != "*" {
		inter = inter.Where("gender = ?", gender)
	}

	respose := inter.Find(&users)
	res := []models.DBUser{}

	if respose.Error != nil {
		return res, respose.Error
	}
	for i := 0; i < len(users); i++ {
		res = append(res, mapSQLUserToDBUser(users[i]))
	}
	return res, nil
}

func mapSQLUserToDBUser(user sql.User) models.DBUser {
	return models.DBUser{
		Id: int(user.ID),
		User: models.User{
			Name:   user.Name,
			Email:  user.Email,
			Age:    user.Age,
			Gender: user.Gender,
			Lat:    user.Lat,
			Lon:    user.Lon,
		},
		Password: user.PasswordHash,
	}
}
