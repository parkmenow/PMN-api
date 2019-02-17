package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/parkmenow/PMN-api/db"
	"github.com/parkmenow/PMN-api/server"
)

func main() {
	database, err := gorm.Open("sqlite3", "server/pmn.db")
	if err != nil {
		panic("failed to establish database connection")
	}
	defer database.Close()
	db.Init(database)
	router := server.CreateRouter(database)
	server.StartServer(router)
}
