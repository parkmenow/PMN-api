package server

import (
	"github.com/gin-gonic/gin"
)

// defineRoutes defines the specification for all future endpoints
func defineRoutes(router *gin.Engine) {
	authMiddleware := JWT()
	router.GET("/", getHello)
	router.POST("/login", authMiddleware.LoginHandler)
	router.POST("/signup", userRegistration)

	user := router.Group("/dashboard/:id") //TODO: remove :id and check the api working
	{
		user.Use(authMiddleware.MiddlewareFunc())
		user.GET("/", getUserFirstName)
		user.GET("/mylistings", mylisting)
		user.PATCH("/mylistings/:property_id/modifyProperty", modifyProperty)
		user.POST("/parkmenow", fetchParkingSpots) //TODO: Aren't we need to show all free stops availalbe at that time.
		user.POST("/regparking", regParkingSpot)
		user.POST("/regparking/regSpot/:spot_id", regSpot)
		user.POST("/regparking/regSpot/:spot_id/regSlot/:slot_id", regSlot)
		user.PATCH("/listings/modifySpot", modifySpot)
		user.PATCH("/payment", payment)
		user.PATCH("/paybywallet", paymentByWallet)
		user.POST("/cancelBooking", cancelBooking)
	}
}
