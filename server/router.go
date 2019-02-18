package server

import (
	"github.com/gin-gonic/gin"
)

// defineRoutes defines the specification for all future endpoints
func defineRoutes(router *gin.Engine) {
	authMiddleware := JWT()
	router.GET("/", getHello)
	router.POST("/login", authMiddleware.LoginHandler)
	// Initial version your API
	v1 := router.Group("/api/v1")

	v1.POST("/signup", userRegistration)
	user := router.Group("/dashboard")
	user.Use(authMiddleware.MiddlewareFunc())
	user.GET("/", getUserFirstName)
	user.Use(authMiddleware.MiddlewareFunc())
<<<<<<< Updated upstream
	user.POST("/parkmenow", fetchParkingSpots)
	user.POST("/regparking", regParkingSpot)
	user.POST("/regSpot", regSpot)
	user.POST("/regSlot", regSlot)
	user.PATCH("/modifySpot", modifySpot)
=======
	user.POST("/:id/parkmenow", fetchParkingSpots)
	user.POST("/:id/regparking", regParkingSpot)
	user.POST("/:id/regparking/regSpot/:spot_id", regSpot)
	user.POST("/:id/regparking/regSpot/:spot_id/regSlot/:slot_id", regSlot)
	user.PATCH("/:id/listings/modifySpot", modifySpot)
>>>>>>> Stashed changes
}
