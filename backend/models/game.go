package models

type Game struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	ImagePath string  `json:"image"`
	BasePrice float64 `json:"basePrice"`
}
