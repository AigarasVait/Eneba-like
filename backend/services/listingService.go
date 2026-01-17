package services

import (
	"backend/models"
	"backend/repositories"
)

type ListingService struct {
	repo *repositories.ListingRepository
}

func NewListingService(repo *repositories.ListingRepository) *ListingService {
	return &ListingService{repo: repo}
}

func (ls *ListingService) GetAllListings() ([]models.Listing, error) {
	return ls.repo.FetchAllListings()
}
