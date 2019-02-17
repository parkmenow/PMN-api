package server

import (
	"github.com/gin-gonic/gin"
)

// defineRoutes defines the specification for all future endpoints
func defineRoutes(router *gin.Engine) {
	authMiddleware := JWT()
	router.POST("/login", authMiddleware.LoginHandler)
	// Initial version your API
	v1 := router.Group("/api/v1")
	v1.GET("/", getHello)
	v1.POST("/signup", userRegistration)
	user := router.Group("/dashboard")
	user.Use(authMiddleware.MiddlewareFunc())
	user.GET("/", getUserFirstName)

}
