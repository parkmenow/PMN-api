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
