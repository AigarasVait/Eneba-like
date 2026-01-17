PRAGMA foreign_keys = ON;

-- Recreate the old game_listings table (without region)
CREATE TABLE game_listings_old (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    price REAL NOT NULL,
    favorited_count INTEGER DEFAULT 0,

    game_id INTEGER NOT NULL,
    gamestore_id INTEGER NOT NULL,

    FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE,
    FOREIGN KEY (gamestore_id) REFERENCES gamestores(id) ON DELETE CASCADE
);

-- Copy data back (region is dropped)
INSERT INTO game_listings_old (id, name, price, favorited_count, game_id, gamestore_id)
SELECT id, name, price, favorited_count, game_id, gamestore_id
FROM game_listings;

DROP TABLE game_listings;
ALTER TABLE game_listings_old RENAME TO game_listings;
