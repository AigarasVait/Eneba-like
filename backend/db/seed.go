package db

import (
	"database/sql"
	"log"
)

func SeedDatabase(db *sql.DB) error {
	// Check if already seeded
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM gamestores").Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		log.Println("Database already seeded, skipping...")
		return nil
	}

	// Begin transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Seed gamestores
	stores := []struct {
		name  string
		image string
	}{
		{"Steam", "steam.png"},
		{"Epic Games", "epic.png"},
		{"GOG", "gog.png"},
	}

	for _, store := range stores {
		_, err := tx.Exec(
			"INSERT INTO gamestores (name, image) VALUES (?, ?)",
			store.name, store.image,
		)
		if err != nil {
			return err
		}
	}

	// Seed games with base_price
	games := []struct {
		name       string
		image      string
		base_price float64
	}{
		{"FIFA 23", "fifa23.jpg", 69.99},
		{"Red Dead Redemption 2", "rdr2.jpg", 59.99},
		{"Split Fiction", "splitfiction.jpg", 49.99},
	}

	for _, game := range games {
		_, err := tx.Exec(
			"INSERT INTO games (name, image, base_price) VALUES (?, ?, ?)",
			game.name, game.image, game.base_price,
		)
		if err != nil {
			return err
		}
	}

	// Seed game listings with actual discounted prices, regions, and favorites
	listings := []struct {
		name           string
		price          float64
		gameID         int
		gamestoreID    int
		region         string
		favoritedCount int
	}{
		{"FIFA 23", 48.99, 1, 1, "EUROPE", 142},
		{"FIFA 23", 69.99, 1, 2, "GLOBAL", 87},
		{"Red Dead Redemption 2", 29.99, 2, 1, "EUROPE", 523},
		{"Red Dead Redemption 2", 44.99, 2, 3, "GLOBAL", 298},
		{"Split Fiction", 49.99, 3, 1, "GLOBAL", 45},
		{"Split Fiction", 42.49, 3, 2, "EUROPE", 76},
	}

	for _, listing := range listings {
		_, err := tx.Exec(
			`INSERT INTO game_listings 
			(name, price, game_id, gamestore_id, region, favorited_count) 
			VALUES (?, ?, ?, ?, ?, ?)`,
			listing.name, listing.price,
			listing.gameID, listing.gamestoreID,
			listing.region, listing.favoritedCount,
		)
		if err != nil {
			return err
		}
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	log.Println("Database seeded successfully")
	return nil
}
