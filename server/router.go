package server

import (
	"github.com/gin-gonic/gin"
)

// defineRoutes defines the specification for all future endpoints
func defineRoutes(router *gin.Engine) {
	// Initial version your API

	v1 := router.Group("/api/v1")
	v1.GET("/", getHello)
	v1.POST("/signup", userRegistration)
}
