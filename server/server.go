package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CreateRouter creates and configures a server and returns a pointer a router with defined handles
func CreateRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(DB(db))
	defineRoutes(router)
	return router
}

// StartServer starts given server
func StartServer(router *gin.Engine) {

	// Get the port to bind server using ENV variable
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port: %s", port)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	// Starts a go routine to start the server
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
