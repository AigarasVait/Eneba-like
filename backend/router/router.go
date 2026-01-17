package router

import (
	"backend/controllers"
	"backend/db"
	"backend/repositories"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	db.Init()

	listingRepo := repositories.NewListingRepository(db.DB)
	listingService := services.NewListingService(listingRepo)
	listingController := controllers.NewListingController(listingService)

	r := gin.Default()

	// Serve static images
	r.Static("/images", "./images")

	// API routes
	r.GET("/list", listingController.ListListings)

	return r
}
