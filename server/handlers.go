package server

import "github.com/gin-gonic/gin"

// getHello defines the endpoint for initial test
func getHello(c *gin.Context) {
	c.String(200, "Hello World")
}
