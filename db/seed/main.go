package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/parkmenow/PMN-api/db"
	"github.com/parkmenow/PMN-api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	var users []models.User
	var owners []models.Owner
	var properties []models.Property
	getData("data/users.json", &users)
	getData("data/owners.json", &owners)
	getData("data/properties.json", &properties)

	database, err := gorm.Open("sqlite3", "pmn.db")
	if err != nil {
		panic("failed to establish database connection")
	}
	defer database.Close()
	db.Init(database)

	for _, user := range users {
		db.DB.Create(&user)
	}
	for _, owner := range owners {
		db.DB.Create(&owner)
	}
	for _, property := range properties {
		db.DB.Create(&property)
	}

	println("Done, copy pmn.db to root folder")
}

func getData(fileName string, v interface{}) {
	file, _ := os.Open(fileName)
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, v)
}
