package db

import (
	"github.com/jinzhu/gorm"
	"github.com/parkmenow/PMN-api/models"
)

// DB is the database connection
var DB *gorm.DB

// Init sets the given database connection as the de-facto connection for this app
func Init(db *gorm.DB) {
	DB = db
	DB.AutoMigrate(&models.User{}, &models.Owner{}, &models.Property{})
}
