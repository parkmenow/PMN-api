package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateRouter creates and configures a server and returns a pointer a router with defined handles
func CreateRouter() *gin.Engine {
	router := gin.Default()
	defineRoutes(router)
	return router
}

// StartServer starts given server
func StartServer(router *gin.Engine) {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
