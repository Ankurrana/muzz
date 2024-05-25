package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ankur-toko/muzz/internal/controllers"
	"github.com/ankur-toko/muzz/internal/models"
	"github.com/golang-jwt/jwt/v5"
	echo "github.com/labstack/echo/v4"
)

func Discover(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)

	claims := user.Claims.(*controllers.JwtCustomClaims)
	userId := claims.UserId
	m := controllers.GetMatchController()

	fromAge := 0
	toAge := 100
	gender := "*"
	fromAgeStr := c.QueryParams().Get("min_age")
	toAgeStr := c.QueryParams().Get("max_age")
	genderStr := c.QueryParams().Get("gender")

	if fromAgeStr != "" {
		fromAge = convertToIntOrDefault(fromAgeStr, 0)
	}
	if toAgeStr != "" {
		toAge = convertToIntOrDefault(toAgeStr, 100)
	}
	if genderStr != "" {
		gender = genderStr
	}

	users, err := m.Discover(userId, fromAge, toAge, gender)
	if err != nil {
		c.String(502, err.Error())
		return err
	}
	c.JSON(200, users)
	return nil
}

func convertToIntOrDefault(str string, d int) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		return d
	}
	return v
}

func Swipe(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)

	claims := user.Claims.(*controllers.JwtCustomClaims)
	userId := claims.UserId

	m := controllers.GetMatchController()

	var sR models.SwipeApiInput
	json.NewDecoder(c.Request().Body).Decode(&sR)

	users, err := m.Swipe(userId, sR)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return err
	}
	c.JSON(http.StatusOK, users)
	return nil
}
