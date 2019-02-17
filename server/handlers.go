package server

import (
	"fmt"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/parkmenow/PMN-api/constants"
	"github.com/parkmenow/PMN-api/models"
)

// getHello defines the endpoint for initial test
func getHello(c *gin.Context) {
	c.String(200, "Hello World")
}

func getDB(c *gin.Context) *gorm.DB {
	return c.MustGet(constants.ContextDB).(*gorm.DB)
}

func userRegistration(c *gin.Context) {
	db := getDB(c)
	var newuser models.User
	c.BindJSON(&newuser)
	fmt.Println(newuser)
	db.Create(&newuser)

	c.JSON(201, "User added successfully!")
}

func getUserFirstName(c *gin.Context) {
	db := getDB(c)
	claims := jwt.ExtractClaims(c)
	id := claims["id"]
	var user models.User
	db.Where("id = ?", id).First(&user)
	c.JSON(200, user.FName)
}

func regParkingSpot(c *gin.Context) {
	db := getDB(c)
	claims := jwt.ExtractClaims(c)
	id := claims["id"]

	var owner models.Owner
	db.Where("user_id = ?", id).First(&owner)

	var property models.Property
	c.BindJSON(&property)

	if owner.UserID == 0 {
		owner.Property = append(owner.Property, property)
		owner.UserID = uint(id.(float64))
		db.Create(&owner)
	} else {
		owner.Property = append(owner.Property, property)
		db.Save(&owner)
	}

	c.JSON(201, "Listed a new parking Spot Successfully!")
}
