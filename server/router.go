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
  
	user := router.Group("/dashboard/:id")
	{
		user.Use(authMiddleware.MiddlewareFunc())
		user.GET("/", getUserFirstName)
		user.GET("/mylistings", mylisting)
		user.POST("/parkmenow", fetchParkingSpots)
		user.POST("/regparking", regParkingSpot)
		user.POST("/regparking/regSpot/:spot_id", regSpot)
		user.POST("/regparking/regSpot/:spot_id/regSlot/:slot_id", regSlot)
		user.PATCH("/listings/modifySpot", modifySpot)
		user.PATCH("/payment", payment)
	}
}
