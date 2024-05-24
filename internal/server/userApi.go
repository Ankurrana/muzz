package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ankur-toko/muzz/internal/controllers"
	"github.com/ankur-toko/muzz/internal/models"
	_ "github.com/asaskevich/govalidator"
	echo "github.com/labstack/echo/v4"
)

var userController controllers.UserController

func init() {
	userController = controllers.UserControllerInstance
}

func AddUser(c echo.Context) error {
	var user models.UserApiInput
	json.NewDecoder(c.Request().Body).Decode(&user)
	fmt.Printf("%+v", user)
	u, e := userController.AddUser(user)
	if e != nil {
		c.String(http.StatusBadRequest, e.Error())
		return e
	} else {
		json, _ := json.Marshal(u)
		c.String(http.StatusOK, string(json))
	}
	return nil
}

func LoginUser(c echo.Context) error {
	var loginRequest models.LoginRequest
	json.NewDecoder(c.Request().Body).Decode(&loginRequest)
	fmt.Printf("%+v", loginRequest)

	res, err := userController.Login(loginRequest)
	if err != nil {
		c.String(http.StatusBadGateway, err.Error())
		return err
	} else {
		j, _ := json.Marshal(res)
		c.String(http.StatusAccepted, string(j))
	}
	return nil
}
