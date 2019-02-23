package server

import (
	"fmt"
	"log"
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
	db := getDB(c)
	var newuser models.User
	var checkUser models.User
	c.BindJSON(&newuser)
	db.Where("u_name = ?", newuser.UName).First(&checkUser)
	if checkUser.UName != "" {
		c.JSON(409, "User Already Exists")
	} else {
		db.Create(&newuser)
		c.JSON(201, "User added successfully!")
	}
}

func getUserFirstName(c *gin.Context) {
	db := getDB(c)
	claims := jwt.ExtractClaims(c)
	id := claims["id"]
	var user models.User
	db.Where("id = ?", id).First(&user)
	c.JSON(200, user.FName)
}

func mylisting(c *gin.Context) {
	db := getDB(c)
	claims := jwt.ExtractClaims(c)
	id := claims["id"]
	var owner models.Owner
	db.Where("user_id = ?", id).First(&owner)
	fmt.Println(owner.ID)
	var properties []models.Property
	db.Preload("Spots").Preload("Spots.Slots").Where("owner_id = ?", owner.ID).Find(&properties)
	c.JSON(200, properties)
}

func modifyProperty(c *gin.Context) {
	var property models.Property
	var modProperty models.Property
	db := getDB(c)
	c.BindJSON(&modProperty)
	fmt.Println(modProperty)
	db.Where("id = ?", modProperty.ID).First(&property).Update(&modProperty)
	c.JSON(200, "Successfully Modified Property")
}

//fetch parking spots. We are assuming that you can book parking for 1hour only.
func fetchParkingSpots(c *gin.Context) {
	var searchInput models.SearchInput
	c.BindJSON(&searchInput)

	db := getDB(c)
	var properties []models.Property
	fmt.Println(searchInput.StartTime)

	db.Preload("Spots", "type = ?", searchInput.Type).Preload("Spots.Slots", "start_time = ? AND available = ?", searchInput.StartTime, 1).Find(&properties)

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

//TODO: refactor function name and output string
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

// UserB pays UserA
func payment(c *gin.Context) {
	var input struct {
		SlotID uint
		Price  int64
		Token  string
	}
	c.BindJSON(&input)

	db := getDB(c)
	claims := jwt.ExtractClaims(c)
	id := claims["id"]
	var userB models.User
	db.Where("id = ?", id).First(&userB)

	// First check if the slot is available
	var slot models.Slot
	db.Where("id = ?", input.SlotID).First(&slot)
	if slot.Available == false {
		c.JSON(401, "Sorry!, Someone has taken the Slot.")
		return
	}

	var fail, failmsg = paymentHandler(input.Price, userB.Email, input.Token)
	if fail == false {
		log.Print(failmsg)
		c.JSON(401, failmsg)
		return
	}

	// Since the payment is successful, Slot is no more available
	slot.Available = false
	db.Save(slot)

	//Extracting Owner ID of the property
	var spot models.Spot
	db.Where("id = ?", slot.SpotID).First(&spot)
	var property models.Property
	db.Where("id = ?", spot.PropertyID).First(&property)

	// Creating the booking record
	newBooking := models.Booking{
		UserID:  userB.ID,
		OwnerID: property.OwnerID,
		SlotID:  input.SlotID,
		Price:   input.Price,
	}
	db.Create(&newBooking)

	// Giving points to the User A
	var owner models.Owner
	db.Where("id = ?", property.OwnerID).First(&owner)
	var userA models.User
	db.Where("id = ?", owner.UserID).First(&userA)
	userA.Wallet = userA.Wallet + input.Price
	db.Save(&userA)

	c.JSON(202, "Booked Successfully!")
}

func paymentByWallet(c *gin.Context) {
	var input struct {
		SlotID uint
		Price  int
	}
	c.BindJSON(&input)
	fmt.Println(input)
	db := getDB(c)

	claims := jwt.ExtractClaims(c)
	id := claims["id"]
	fmt.Println(id)
	// First Check if User's wallet has enough Balance in it
	var userB models.User
	db.Where("id = ?", id).First(&userB)
	if userB.Wallet < int64(input.Price) {
		c.JSON(401, "Soory!, Not enough Credit in the Wallet.")
		return
	}

	// First check if the slot is available
	var slot models.Slot
	db.Where("id = ?", input.SlotID).First(&slot)
	if slot.Available == false {
		c.JSON(401, "Sorry!, Someone has taken the Slot.")
		return
	}

	// Since the payment is successful, Slot is no more available
	slot.Available = false
	db.Save(slot)

	//Extracting User ID of the property
	var spot models.Spot
	db.Where("id = ?", slot.SpotID).First(&spot)
	var property models.Property
	db.Where("id = ?", spot.PropertyID).First(&property)
	var owner models.Owner
	db.Where("id = ?", property.OwnerID).First(&owner)
	var userA models.User
	db.Where("id = ?", owner.UserID).First(&userA)

	// Add the money to wallet of User of the Spot
	userA.Wallet = userA.Wallet + int64(input.Price)
	db.Save(userA)

	// Subtract money from wallet of User who booked the Spot
	userB.Wallet = userB.Wallet - int64(input.Price)
	db.Save(userB)
	// Creating the booking record
	newBooking := models.Booking{
		UserID:  userB.ID,
		OwnerID: property.OwnerID,
		SlotID:  input.SlotID,
		Price:   int64(input.Price),
	}
	db.Create(&newBooking)
	c.JSON(202, "Booked Successfully!")
}

func modifySpot(c *gin.Context) {
	var spot models.Spot
	var modSpot models.Spot
	db := getDB(c)
	c.BindJSON(&modSpot)
	fmt.Println(modSpot)
	db.Where("id = ?", modSpot.ID).First(&spot).Update(&modSpot)
	c.JSON(200, "Successfully Modified Spot")

}

func cancelBooking(c *gin.Context){
	// 1) Get the booking id and corresponding row of that booking
	type BID struct {
		// Booking Id
		BookingId int
	}
	var bid BID
	var booking models.Booking
	c.BindJSON(&bid)
	fmt.Println(bid.BookingId)
	db := getDB(c)
	// Assuming the id exists
	db.Where("id=?", bid.BookingId).First(&booking)
	booking.Status = "cancelled"
	db.Save(&booking)

	// 2) Make the slot available now, meaning change the status of the slot to available
	fmt.Println(booking.SlotID)
	var slot models.Slot
	db.Where("id=?", booking.SlotID).Find(&slot)
	slot.Available = true
	fmt.Println(slot.StartTime.Sub(time.Now()).Hours())
	db.Save(&slot)

	// If user cancels booking before 48 hours, cancellation charge is free
	// If before 48 hours to 24 hours, cancellation is 5% charge
	// If it is less than 24 hours, cancellation is 10 % charge
	var cancellationPercentage int64
	hoursLeft := slot.StartTime.Sub(time.Now()).Hours()
	if hoursLeft > 48{
		cancellationPercentage = 0
	}else if hoursLeft<48 && hoursLeft >24{
		cancellationPercentage = 5

	}else if hoursLeft < 24 && hoursLeft > 0{
		cancellationPercentage = 10
	}else if hoursLeft < 0{
		c.JSON(200, gin.H{"info":"Booking cannot be cancelled, as time passed the start time",})
	}

	// 3) Deduct 10% of price user has paid to his booking, so add 90% of price to his wallet, tested user gets money into his account
	var user models.User
	db.Where("id=?", booking.UserID).Find(&user)
	fmt.Println(user.Wallet)
	user.Wallet = int64((100-cancellationPercentage) * booking.Price/100)
	db.Save(&user)

	// 4) Add 10% of amount to owner of the spot
	var owner models.User
	db.Where("id=?", booking.OwnerID).Find(&owner)
	owner.Wallet = int64(cancellationPercentage * booking.Price/100)
	db.Save(&owner)

	// 5) Inform the API saying the task is done, with status code 200 and JSON that booking is cancelled
	c.JSON(200, gin.H{"info":"Booking is cancelled",})

}


func modifySlot(c *gin.Context) {
	var slot models.Slot
	var modSlot models.Slot
	db := getDB(c)
	c.BindJSON(&modSlot)
	fmt.Println(modSlot)
	db.Where("id = ?", modSlot.ID).First(&slot).Update(&modSlot)
	c.JSON(200, "Successfully Modified Spot")
}

func showBookings(c *gin.Context) {
	db := getDB(c)
	var bookings []models.Booking
	claims := jwt.ExtractClaims(c)
	id := claims["id"]
	db.Where("user_id = ? and status = ?", id, constants.StatusActive).Find(&bookings)
	c.JSON(200, bookings)
}
