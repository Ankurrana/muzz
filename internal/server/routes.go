package server

import (
	"github.com/ankur-toko/muzz/internal/controllers"
	echojwt "github.com/labstack/echo-jwt/v4"
	echo "github.com/labstack/echo/v4"
)

/*
curl -d '{"name":"ankur","email":"abc@gmail.com","password":"ankur123","age":"30","lat":"-23.324","lon":"23.123","gender":"male"}' -H 'Content-Type: application/json' localhost:8000/user
*/
func AddRoutes(e *echo.Echo) {
	e.POST("/user/create", AddUser)
	e.POST("/login", LoginUser)
	AddAuthenticatedRoutes(e)
}

func AddAuthenticatedRoutes(e *echo.Echo) {
	r := e.Group("")
	// Configure middleware with the custom claims type
	config := controllers.GetJWTTokenConfig()
	r.Use(echojwt.WithConfig(config))

	r.GET("/discover", Discover)
	r.POST("/swipe", Swipe)
}
