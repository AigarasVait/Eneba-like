package router

import (
	"github.com/gin-gonic/gin"
	// "myapp/handlers"  // if you have handlers
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Serve static images
	r.Static("/images", "./images")

	// API routes
	// r.GET("/list", handlers.ListHandler)

	return r
}
