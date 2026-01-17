package models

type Listing struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	FavoritedCount int     `json:"favoritedCount"`
	Region         string  `json:"region"`

	GameID int   `json:"gameId"`
	Game   *Game `json:"game,omitempty"`

	GameStoreID int        `json:"gameStoreId"`
	GameStore   *GameStore `json:"gameStore,omitempty"`
}

type ListingCardDTO struct {
	ID                 int     `json:"id"`
	Name               string  `json:"name"`
	Price              float64 `json:"price"`
	FavoritedCount     int     `json:"favoritedCount"`
	GameStoreName      string  `json:"gameStoreName"`
	GameStoreImagePath string  `json:"gameStoreImagePath"`
	GameImagePath      string  `json:"gameImagePath"`
	BasePrice          float64 `json:"basePrice"`
	Region             string  `json:"region"`
	Cashback           float64 `json:"cashback"`
	DiscountPercent    float64 `json:"discountPercent"`
}
