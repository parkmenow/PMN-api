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
	// 1) Get the environment variable if any set for which database and what the database url is.
	DATABASE := os.Getenv("DB_DRIVER")
	databaseURL := os.Getenv("DATABASE_URL")

	// 2) If any of the environment variables are not set, then check for a .env file, that may contain required variables
	if DATABASE == "" || databaseURL == "" {
		// Load the environment variables from .env and set the variables
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		DATABASE = os.Getenv("DB_DRIVER")
		databaseURL = os.Getenv("DATABASE_URL")
	}
	// 3) Connect to the gorm database
	database, err := gorm.Open(DATABASE, databaseURL)
	if err != nil {
		panic("failed to establish database connection")
	}
	defer database.Close()
	db.Init(database)

	// 4) Creating a server with database parameter, so that server initalizes the database.
	router := server.CreateRouter(database)
	server.StartServer(router)
}
