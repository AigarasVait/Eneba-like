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

func (ls *ListingService) GetAllListings() ([]models.ListingCardDTO, error) {
	listings, err := ls.repo.FetchAllListings()

	if err != nil {
		return nil, err
	}

	var dtos []models.ListingCardDTO

	for _, listing := range listings {
		var dto models.ListingCardDTO

		dto.ID = listing.ID
		dto.Name = listing.Name
		dto.Price = listing.Price
		dto.Region = listing.Region
		dto.FavoritedCount = listing.FavoritedCount
		dto.GameStoreName = listing.GameStore.Name
		dto.GameStoreImagePath = listing.GameStore.ImagePath
		dto.GameImagePath = listing.Game.ImagePath
		dto.BasePrice = listing.Game.BasePrice
		dto.Cashback = listing.Price * 0.11
		dto.DiscountPercent = (listing.Game.BasePrice - listing.Price) / listing.Game.BasePrice * 100

		dtos = append(dtos, dto)
	}
	return dtos, nil
}
