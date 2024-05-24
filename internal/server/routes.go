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
	// curl -d '{"name":"ankur","email":"abc@gmail.com","password":"ankur123","age":30,"lat":-23.324,"lon":23.123,"gender":"male"}' -H 'Content-Type: application/json' localhost:8000/user
	e.POST("/user", AddUser)

	// curl -d '{"email":"abc@gmail.com","password":"ankur123"}' -H 'Content-Type: application/json' localhost:8000/login
	e.POST("/login", LoginUser)

	AddAuthenticatedRoutes(e)

}

// curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFiY0BnbWFpbC5jb20iLCJ1c2VyX2lkIjoxLCJleHAiOjE3MTY1MTAzODB9.yKBYsYom6KY9TKpay1A8lVVNIKhaGNinoHlYE9hsiso" -d '{"user_id":2,"preference":"YES"}' -H 'Content-Type: application/json' localhost:8000/swipe
// curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImJpdHRvQGdtYWlsLmNvbSIsInVzZXJfaWQiOjIsImV4cCI6MTcxNjUxMDUyMH0.r54LvOZtLygO1bjsZ8d7cSCIJFwpJb8pDZDnVhlGZ5M" -d '{"user_id":1,"preference":"YES"}' -H 'Content-Type: application/json' localhost:8000/swipe

// curl localhost:1323/restricted -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFiY0BnbWFpbC5jb20iLCJ1c2VyX2lkIjoxLCJleHAiOjE3MTY1MDQwNzJ9.AJpQp1WgQf8LAxR3IKZFJZH55QaCV8X_JfQ62M7iPPg"
// curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFiY0BnbWFpbC5jb20iLCJ1c2VyX2lkIjoxLCJleHAiOjE3MTY1MTAwNjl9.r1PZbsoGXF5gE2CblIKSWG7VKH4EoTaddCmups1pJCw" localhost:8000/discover
func AddAuthenticatedRoutes(e *echo.Echo) {
	r := e.Group("")
	// Configure middleware with the custom claims type
	config := controllers.GetJWTTokenConfig()
	r.Use(echojwt.WithConfig(config))
	r.GET("/discover", Discover)
	r.POST("/swipe", Swipe)
}
