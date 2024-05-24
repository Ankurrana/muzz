package server

import (
	echo "github.com/labstack/echo/v4"
)

func StartServer() {
	e := echo.New()
	AddRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
