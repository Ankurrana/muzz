package main

import (
	"time"

	"github.com/ankur-toko/muzz/internal/server"
)

func main() {
	server.StartServer()
	time.Sleep(1 * time.Hour)
}
