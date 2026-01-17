-- Add base_price column to games table
ALTER TABLE games ADD COLUMN base_price REAL;

-- Migrate discount logic: update game base_price with max price from listings
UPDATE games
SET base_price = (
    SELECT MAX(price / (1 - discount / 100.0))
    FROM game_listings
    WHERE game_listings.game_id = games.id
);

-- For games without listings, set a default base price
UPDATE games
SET base_price = 59.99
WHERE base_price IS NULL;

-- Make base_price NOT NULL
-- SQLite doesn't support ALTER COLUMN, so we need to recreate the table
CREATE TABLE games_new (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    image TEXT,
    base_price REAL NOT NULL
);

INSERT INTO games_new (id, name, image, base_price)
SELECT id, name, image, base_price FROM games;

DROP TABLE games;
ALTER TABLE games_new RENAME TO games;

-- Remove discount column from game_listings
CREATE TABLE game_listings_new (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    price REAL NOT NULL,
    favorited_count INTEGER DEFAULT 0,
    game_id INTEGER NOT NULL,
    gamestore_id INTEGER NOT NULL,
    FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE,
    FOREIGN KEY (gamestore_id) REFERENCES gamestores(id) ON DELETE CASCADE
);

INSERT INTO game_listings_new (id, name, price, favorited_count, game_id, gamestore_id)
SELECT id, name, price, favorited_count, game_id, gamestore_id FROM game_listings;

DROP TABLE game_listings;
ALTER TABLE game_listings_new RENAME TO game_listings;