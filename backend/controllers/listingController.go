package controllers

import (
	"backend/services"

	"github.com/gin-gonic/gin"
)

type ListingController struct{ Service *services.ListingService }

func NewListingController(service *services.ListingService) *ListingController {
	return &ListingController{Service: service}
}

func (lc *ListingController) ListListings(c *gin.Context) {
	search := c.Query("search")
	listingDTOs, err := lc.Service.GetAllListings(search)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, listingDTOs)
}
