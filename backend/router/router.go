package router

import (
	"backend/controllers"
	"backend/db"
	"backend/repositories"
	"backend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	db.Init()

	listingRepo := repositories.NewListingRepository(db.DB)
	listingService := services.NewListingService(listingRepo)
	listingController := controllers.NewListingController(listingService)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	// Serve static images
	r.Static("/images", "./images")

	// API routes
	r.GET("/list", listingController.ListListings)

	return r
}
