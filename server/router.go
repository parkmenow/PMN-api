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
	user := v1.Group("/dashboard/")
	user.Use(authMiddleware.MiddlewareFunc())
	user.GET("/:id", getUserFirstName)
	user.Use(authMiddleware.MiddlewareFunc())
	user.POST("/:id/parkmenow", fetchParkingSpots)
	user.POST("/:id/regparking", regParkingSpot)
	user.POST("/:id/regparking/regSpot/:spot_id", regSpot)
	user.POST("/:id/regparking/regSpot/regSlot/:slot_id", regSlot)
	user.PATCH("/:id/listings/modifySpot", modifySpot)
}
