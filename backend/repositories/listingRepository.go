package repositories

import (
	"backend/models"
	"database/sql"
)

type ListingRepository struct {
	db *sql.DB
}

func NewListingRepository(db *sql.DB) *ListingRepository {
	return &ListingRepository{db: db}
}

func (lr *ListingRepository) FetchListings(search string) ([]models.Listing, error) {
	baseQuery := `
        SELECT 
            l.id, l.name, l.price, l.favorited_count, l.region,
            l.game_id,
            g.id, g.name, g.image, g.base_price,
            l.gamestore_id,
            gs.id, gs.name, gs.image
        FROM game_listings l
        JOIN games g ON g.id = l.game_id
        JOIN gamestores gs ON gs.id = l.gamestore_id
    `

	var rows *sql.Rows
	var err error

	if search != "" {
		rows, err = lr.db.Query(
			baseQuery+` WHERE LOWER(l.name) LIKE '%' || LOWER(?) || '%'`,
			search,
		)
	} else {
		rows, err = lr.db.Query(baseQuery)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listings []models.Listing

	for rows.Next() {
		var l models.Listing
		var g models.Game
		var gs models.GameStore

		err := rows.Scan(
			&l.ID, &l.Name, &l.Price, &l.FavoritedCount, &l.Region,
			&l.GameID,
			&g.ID, &g.Name, &g.ImagePath, &g.BasePrice,
			&l.GameStoreID,
			&gs.ID, &gs.Name, &gs.ImagePath,
		)
		if err != nil {
			return nil, err
		}

		l.Game = &g
		l.GameStore = &gs

		listings = append(listings, l)
	}

	return listings, nil
}
