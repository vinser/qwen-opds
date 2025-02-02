package db

import (
	"time"
)

type Book struct {
	ID          int       `db:"id"`
	Filepath    string    `db:"filepath"`
	CRC32       string    `db:"crc32"`
	ArchiveName string    `db:"archive_name"`
	Size        int64     `db:"size"`
	Format      string    `db:"format"`
	Title       string    `db:"title"`
	SortTitle   string    `db:"sort_title"`
	Year        int       `db:"year"`
	Plot        string    `db:"plot"`
	CoverPath   string    `db:"cover_path"`
	LanguageID  int       `db:"language_id"`
	SerieID     *int      `db:"serie_id"`
	LastUpdated time.Time `db:"last_updated"`
}

type Language struct {
	ID   int    `db:"id"`
	Code string `db:"code"`
	Name string `db:"name"`
}

type Author struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Genre struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Serie struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
