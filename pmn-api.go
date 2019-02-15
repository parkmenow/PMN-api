package main

import (
	"github.com/parkmenow/PMN-api/server"
)

func main() {
	router := server.CreateRouter()
	server.StartServer(router)
}
