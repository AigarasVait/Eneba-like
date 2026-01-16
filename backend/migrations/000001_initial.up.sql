PRAGMA foreign_keys = ON;

-- Table: gamestores
CREATE TABLE gamestores (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    image TEXT
);

-- Table: games
CREATE TABLE games (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    image TEXT
);

-- Table: game_listings
CREATE TABLE game_listings (
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
