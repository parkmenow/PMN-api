package server

import (
	"fmt"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/parkmenow/PMN-api/constants"
	"github.com/parkmenow/PMN-api/models"
	"log"
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

//fetch parking spots. We are assuming that you can book parking for 1hour only.
func fetchParkingSpots(c *gin.Context) {
	var searchInput models.SearchInput
	c.BindJSON(&searchInput)

	db := getDB(c)

	var properties []models.Property
	fmt.Println(searchInput.StartTime)
	db.Preload("Spots", "type = ?", searchInput.Type).Preload("Spots.Slots", "start_time = ?", searchInput.StartTime).Find(&properties)

	//fmt.Println(searchInput)
	// layout := "2006-01-02T15:04:05.000Z"
	// str := searchInput.StartTime[0:10] + "T" + searchInput.StartTime[11:] + ":00.000Z"
	// startTime, _ := time.Parse(layout, str)
	// str = searchInput.EndTime[0:10] + "T" + searchInput.EndTime[11:] + ":00.000Z"
	// endTime, _ := time.Parse(layout, str)
	// var results []models.Spot
	// "start_time = ?", searchInput.StartTime
	// var spots []models.Spot
	// db.Preload("Slots").Where("type = ?", searchInput.Type).Find(&spots)
	// var b bool
	// for _, sp := range spots {
	// 	b = false
	// 	var r []models.Slot
	// 	for _, s := range sp.Slots {
	// 		str := s.StartTime[0:10] + "T" + s.StartTime[11:] + ":00.000Z"
	// 		fmt.Println(str)
	// 		st, _ := time.Parse(layout, str)
	// 		if st == startTime {
	// 			b = true
	// 			r = append(r, s)
	// 		}
	// 	}
	// 	if b {
	// 		var result models.Spot
	// 		result.Type = sp.Type
	// 		result.DBModel = sp.DBModel
	// 		result.ImageURL = sp.ImageURL
	// 		result.Description = sp.Description
	// 		result.PropertyID = sp.PropertyID
	// 		result.Slots = r
	// 		results = append(results, result)
	// 	}
	// }
	// var properties []models.Property
	// for _, res := range results {
	// 	var property models.Property
	// 	db.Where("id = ?", res.PropertyID).Find(&property)
	// 	property.Spots = append(property.Spots, res)
	// 	properties = append(properties, property)
	// }
	//fmt.Println(results)
	c.JSON(200, properties)
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

func regSpot(c *gin.Context) {
	var spot models.Spot

	db := getDB(c)
	c.BindJSON(&spot)
	db.Create(&spot)
	c.JSON(200, "Successfully Added Spot")
}

func regSlot(c *gin.Context) {
	var slot models.Slot

	db := getDB(c)
	c.BindJSON(&slot)
	db.Create(&slot)
	c.JSON(200, "Successfully Added Slot")
}

func payment(c *gin.Context) {
	var input struct {
		Owner  uint
		SlotID uint
		Price  int64
		Token  string
	}
	c.BindJSON(&input)

	db := getDB(c)
	claims := jwt.ExtractClaims(c)
	id := claims["id"]
	var user models.User
	db.Where("id = ?", id).First(&user)

	var fail, failmsg = paymentHandler(input.Price, user.Email, input.Token)
	if fail == false {
		log.Print(failmsg)
		c.JSON(401, failmsg)
		return
	}

	// Since the payment is successful, Slot is no more available
	var slot models.Slot
	db.Where("id = ?", input.SlotID).First(&slot)
	slot.Availabile = false
	db.Save(slot)

	// Creating the booking record
	newBooking := models.Booking{
		Booker: user.ID,
		OwnerID:  input.Owner,
		SlotID: input.SlotID,
		Price:  input.Price,
	}
	db.Create(&newBooking)

	// Giving points to the Owner
	var owner models.Owner
	db.Where("user_id = ?",input.Owner).First(&owner)
	owner.Wallet = owner.Wallet + input.Price
	db.Save(&owner)

	c.JSON(202, "Booked Successfully!")
}
