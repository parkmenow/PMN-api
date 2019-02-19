package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/parkmenow/PMN-api/db"
	"github.com/parkmenow/PMN-api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	var users []models.User
	var owners []models.Owner
	var properties []models.Property
	var spots []models.Spot
	var slots []models.Slot
	var bookings []models.Booking

	getData("data/bookings.json", &bookings)
	getData("data/users.json", &users)
	getData("data/owners.json", &owners)
	getData("data/properties.json", &properties)
	getData("data/spots.json", &spots)
	getData("data/slots.json", &slots)
	//database, err := gorm.Open("sqlite3", "pmn.db")
	DATABASE := os.Getenv("DB_DRIVER")
	databaseURL := os.Getenv("DATABASE_URL")
	// var database *gorm.DB
	if DATABASE == "" {
		err := godotenv.Load("../../.env")
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

	for _, booking := range bookings {
		db.DB.Create(&booking)
	}
	for _, user := range users {
		db.DB.Create(&user)
	}
	for _, owner := range owners {
		db.DB.Create(&owner)
	}
	for _, property := range properties {
		db.DB.Create(&property)
	}
	for _, spot := range spots {
		db.DB.Create(&spot)
	}
	for _, slot := range slots {
		db.DB.Create(&slot)
	}

	println("Done, copy pmn.db to root folder")

}

func getData(fileName string, v interface{}) {
	file, _ := os.Open(fileName)
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, v)
}