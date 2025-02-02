package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

// InitDatabase initializes the SQLite database
func InitDatabase(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	migrateDB(db)
	log.Println("Database initialized successfully")
	return db
}

func migrateDB(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS languages (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        code TEXT NOT NULL UNIQUE,
        name TEXT NOT NULL
    );

    CREATE TABLE IF NOT EXISTS authors (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL UNIQUE
    );

    CREATE TABLE IF NOT EXISTS genres (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL UNIQUE
    );

    CREATE TABLE IF NOT EXISTS series (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL UNIQUE
    );

    CREATE TABLE IF NOT EXISTS books (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        filepath TEXT NOT NULL,
        crc32 TEXT NOT NULL,
        archive_name TEXT,
        size INTEGER NOT NULL,
        format TEXT NOT NULL,
        title TEXT NOT NULL,
        sort_title TEXT NOT NULL,
        year INTEGER,
        plot TEXT,
        cover_path TEXT,
        language_id INTEGER,
        serie_id INTEGER,
        last_updated INTEGER NOT NULL,
        FOREIGN KEY (language_id) REFERENCES languages(id),
        FOREIGN KEY (serie_id) REFERENCES series(id)
    );

    CREATE TABLE IF NOT EXISTS book_authors (
        book_id INTEGER NOT NULL,
        author_id INTEGER NOT NULL,
        PRIMARY KEY (book_id, author_id),
        FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
        FOREIGN KEY (author_id) REFERENCES authors(id) ON DELETE CASCADE
    );

    CREATE TABLE IF NOT EXISTS book_genres (
        book_id INTEGER NOT NULL,
        genre_id INTEGER NOT NULL,
        PRIMARY KEY (book_id, genre_id),
        FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
        FOREIGN KEY (genre_id) REFERENCES genres(id) ON DELETE CASCADE
    );
    `

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
