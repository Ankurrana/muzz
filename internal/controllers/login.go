package controllers

import (
	"fmt"
	"time"

	"github.com/ankur-toko/muzz/internal/models"
	"github.com/asaskevich/govalidator"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var SecretKey = []byte("secret-key-1")

type LoginResponse struct {
	Token string `json:"token"`
}

type JwtCustomClaims struct {
	Username string `json:"username"`
	UserId   int    `json:"user_id"`
	jwt.RegisteredClaims
}

func (uC UserController) Login(lR models.LoginRequest) (LoginResponse, error) {
	// Do the validations here and respond back
	ok, err := govalidator.ValidateStruct(lR)
	if err != nil {
		fmt.Println(err, ok)
	}

	user, _ := uC.uRepo.GetUserByEmail(lR.Email)
	if user.Password == uC.hashAlgo.Hash(lR.Password) {
		token, err := createToken(user.Email, user.Id)
		if err != nil {
			return LoginResponse{}, err
		} else {
			return LoginResponse{token}, nil
		}
	} else {
		return LoginResponse{}, fmt.Errorf("error: incorrect email or password")
	}
}

func createToken(username string, userId int) (string, error) {
	claims := &JwtCustomClaims{
		username,
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(SecretKey))
	return t, err
}

func GetJWTTokenConfig() echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtCustomClaims)
		},
		SigningKey: SecretKey,
	}
}
