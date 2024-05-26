package controllers

import (
	"reflect"
	"testing"

	"github.com/ankur-toko/muzz/internal/models"
	"github.com/ankur-toko/muzz/internal/repositories/match"
	"github.com/ankur-toko/muzz/internal/repositories/user"
)

func TestUserController_AddUser(t *testing.T) {
	uRepo := user.Instance()
	mRepo := match.Instance()
	hash := MyDummyHash{}
	type args struct {
		u models.UserApiInput
	}

	user := models.UserApiInput{
		Name:     "Ankur",
		Email:    "ankur@gmail.com",
		Password: "Asdasd",
		Age:      23,
		Gender:   "male",
		Lat:      123.123,
		Lon:      1234.123,
	}

	missingName := clone(user)
	missingName.Name = ""

	missingEmail := clone(user)
	missingEmail.Email = ""

	wrongEmail := clone(user)
	wrongEmail.Email = "notcorrectemailcom.gmail"

	wrongGender := clone(user)
	wrongGender.Gender = "somethingnotgender"

	wrongAge := clone(user)
	wrongAge.Age = 1000

	wrongAge2 := clone(user)
	wrongAge2.Age = -11

	correctGender := clone(user)
	correctGender.Gender = "lgbtqp"

	missingPassword := clone(user)
	missingPassword.Password = ""

	u := models.User{Name: user.Name, Email: user.Email, Age: user.Age, Gender: user.Gender, Lat: user.Lat, Lon: user.Lon}
	u.Gender = "lgbtqp"

	tests := []struct {
		name    string
		args    args
		want    models.User
		wantErr bool
	}{
		{"missingName", args{missingName}, models.User{}, true},
		{"missingEmail", args{missingEmail}, models.User{}, true},
		{"missingPass", args{missingPassword}, models.User{}, true},
		{"wrongEmail", args{wrongEmail}, models.User{}, true},
		{"wrongAge", args{wrongAge}, models.User{}, true},
		{"wrongAge2", args{wrongAge2}, models.User{}, true},
		{"wrongGender", args{wrongGender}, models.User{}, true},
		{"correctGender", args{correctGender}, u, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uC := UserController{
				uRepo:    uRepo,
				mRepo:    mRepo,
				hashAlgo: hash,
			}
			got, err := uC.AddUser(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserController.AddUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserController.AddUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func clone(a models.UserApiInput) models.UserApiInput {
	return models.UserApiInput{
		Name: a.Name, Email: a.Email, Password: a.Password, Age: a.Age, Gender: a.Gender, Lat: a.Lat, Lon: a.Lon,
	}
}
