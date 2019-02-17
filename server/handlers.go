package server

import (
	"encoding/json"
	"fmt"

	jwt "github.com/appleboy/gin-jwt"
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
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])

	db := getDB(c)
	var newuser models.User

	json.Unmarshal([]byte(reqBody), &newuser)
	db.Create(&newuser)

	c.JSON(201, "User added successfully!")
}

func getUserFirstName(c *gin.Context) {
	db := getDB(c)
	claims := jwt.ExtractClaims(c)
	id := claims["id"]
	fmt.Println(id)
	var user models.User
	db.Where("id = ?", id).First(&user)
	fmt.Println(user)
	c.JSON(200, user.FName)
}

