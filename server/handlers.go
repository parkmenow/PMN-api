package server

import (
	"github.com/parkmenow/PMN-api/constants"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// getHello defines the endpoint for initial test
func getHello(c *gin.Context) {
	c.String(200, "Hello World")
}

func getDB(c *gin.Context) *gorm.DB {
	return c.MustGet(constants.ContextDB).(*gorm.DB)
}
