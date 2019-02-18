package server

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/parkmenow/PMN-api/constants"
	"github.com/parkmenow/PMN-api/models"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"log"
	"os"
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

func paymentHandler(c *gin.Context) {
	var input struct {       //to be added in ther amount
		PropertyID uint
		Price int64
		Token string
	}
	c.BindJSON(&input)

	//export SecretKey="sk_test_1pSlxntEQATjsOv5HLI49FaW"
	var sh_key = os.Getenv("SecretKey")
	stripe.Key = sh_key

	params := &stripe.ChargeParams{
		Amount:   stripe.Int64(input.Price),
		Currency: stripe.String(string(stripe.CurrencyJPY)),
	}

	//Add the token here
	params.SetSource("tok_mastercard")
	//params.SetSource(input.Token)

	ch, err := charge.New(params)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", ch.ID)
	c.JSON(200, "Payment Successful")
}
