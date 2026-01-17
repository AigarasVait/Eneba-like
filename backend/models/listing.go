package models

type Listing struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	FavoritedCount int     `json:"favorited_count"`
	Region         string  `json:"region"`

	GameID int   `json:"game_id"`
	Game   *Game `json:"game,omitempty"`

	GameStoreID int        `json:"game_store_id"`
	GameStore   *GameStore `json:"game_store,omitempty"`
}

type ListingCardDTO struct {
	ID                 int     `json:"id"`
	Name               string  `json:"name"`
	Price              float64 `json:"price"`
	FavoritedCount     int     `json:"favorited_count"`
	GameStoreName      string  `json:"game_store_name"`
	GameStoreImagePath string  `json:"game_store_image"`
	GameImagePath      string  `json:"game_image"`
	BasePrice          float64 `json:"base_price"`
	Region             string  `json:"region"`
	Cashback           float64 `json:"cashback"`
	DiscountPercent    float64 `json:"discount_percent"`
}
