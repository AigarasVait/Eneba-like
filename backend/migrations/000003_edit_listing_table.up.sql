PRAGMA foreign_keys = ON;

-- Recreate game_listings with new region column
CREATE TABLE game_listings_new (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    price REAL NOT NULL,
    favorited_count INTEGER DEFAULT 0,
    region TEXT NOT NULL DEFAULT 'GLOBAL',

    game_id INTEGER NOT NULL,
    gamestore_id INTEGER NOT NULL,

    FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE,
    FOREIGN KEY (gamestore_id) REFERENCES gamestores(id) ON DELETE CASCADE
);

-- Copy existing data, region defaults to GLOBAL
INSERT INTO game_listings_new (id, name, price, favorited_count, game_id, gamestore_id)
SELECT id, name, price, favorited_count, game_id, gamestore_id
FROM game_listings;

DROP TABLE game_listings;
ALTER TABLE game_listings_new RENAME TO game_listings;
