package server

import (
	"encoding/json"
	"fmt"
	"time"

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

	fmt.Println(num)
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

//fetch parking spots. We are assuming that you can book parking for 1hour only.
func fetchParkingSpots(c *gin.Context) {
	var searchInput models.SearchInput
	c.BindJSON(&searchInput)
	//fmt.Println(searchInput)
	layout := "2006-01-02T15:04:05.000Z"
	str := searchInput.StartTime[0:10] + "T" + searchInput.StartTime[11:] + ":00.000Z"
	startTime, _ := time.Parse(layout, str)
	// str = searchInput.EndTime[0:10] + "T" + searchInput.EndTime[11:] + ":00.000Z"
	// endTime, _ := time.Parse(layout, str)
	db := getDB(c)
	var results []models.Slot
	var slots []models.Slot
	db.Find(&slots)
	for _, s := range slots {
		// fmt.Println(s.StartTime)
		str := s.StartTime[0:10] + "T" + s.StartTime[11:] + ":00.000Z"
		st, _ := time.Parse(layout, str)
		if st == startTime {
			results = append(results, s)
		}
	}
	c.JSON(200, results)
}
