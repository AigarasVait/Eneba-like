package models

type GameStore struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ImagePath string `json:"image"`
}
