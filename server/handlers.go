package server

import(
	"github.com/gin-gonic/gin"

	"github.com/parkmenow/PMN-api/data"
	)

// getHello defines the endpoint for initial test
func getHello(c *gin.Context) {
	c.String(200, "Hello World")
}

func userRegistration(c *gin.Context)  {
	buf := make([]byte, 2048)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	c.JSON(201, data.UserRegistration(reqBody))
}
