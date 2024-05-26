package main

import (
	"time"

	"github.com/ankur-toko/muzz/internal/controllers"
	"github.com/ankur-toko/muzz/internal/server"
)

func main() {
	controllers.Initialize()
	server.StartServer()
	time.Sleep(1 * time.Hour)
}
