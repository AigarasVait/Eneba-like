-- Add discount column back to game_listings
CREATE TABLE game_listings_new (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    price REAL NOT NULL,
    discount REAL DEFAULT 0,
    favorited_count INTEGER DEFAULT 0,
    game_id INTEGER NOT NULL,
    gamestore_id INTEGER NOT NULL,
    FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE,
    FOREIGN KEY (gamestore_id) REFERENCES gamestores(id) ON DELETE CASCADE
);

INSERT INTO game_listings_new (id, name, price, discount, favorited_count, game_id, gamestore_id)
SELECT id, name, price, 0, favorited_count, game_id, gamestore_id FROM game_listings;

DROP TABLE game_listings;
ALTER TABLE game_listings_new RENAME TO game_listings;

-- Remove base_price column from games
CREATE TABLE games_new (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    image TEXT
);

INSERT INTO games_new (id, name, image)
SELECT id, name, image FROM games;

DROP TABLE games;
ALTER TABLE games_new RENAME TO games;