package server

import (
	"fmt"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
<<<<<<< HEAD

=======
  
>>>>>>> ac56f02b8cb2a7661d1d335ca28002d5a13ad2fb
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

<<<<<<< HEAD
=======

>>>>>>> ac56f02b8cb2a7661d1d335ca28002d5a13ad2fb
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
