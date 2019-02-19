package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
	"github.com/parkmenow/PMN-api/db"
	"github.com/parkmenow/PMN-api/server"
)

func main() {
	DATABASE := os.Getenv("DB_DRIVER")
	databaseURL := os.Getenv("DATABASE_URL")
	if DATABASE == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		DATABASE = os.Getenv("DB_DRIVER")
		databaseURL = os.Getenv("DATABASE_URL")
	}
	database, err := gorm.Open(DATABASE, databaseURL)
	if err != nil {
		panic("failed to establish database connection")
	}
	defer database.Close()
	db.Init(database)
	router := server.CreateRouter(database)
	server.StartServer(router)
}
