package models

type Listing struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	FavoritedCount int     `json:"favorited_count"`

	GameID int   `json:"game_id"`
	Game   *Game `json:"game,omitempty"`

	GameStoreID int        `json:"game_store_id"`
	GameStore   *GameStore `json:"game_store,omitempty"`
}
