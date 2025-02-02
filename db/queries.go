package db

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

func AddBook(db *sql.DB, bookData map[string]interface{}) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	languageCode := bookData["language"].(string)
	languageID, err := getOrCreateLanguage(tx, languageCode)
	if err != nil {
		return err
	}

	var serieID sql.NullInt64
	if serieName, ok := bookData["serie"].(string); ok && serieName != "" {
		serieIDInt, err := getOrCreateSerie(tx, serieName) // Исправлено: обработка двух значений
		if err != nil {
			return fmt.Errorf("failed to get or create serie: %v", err)
		}
		serieID.Int64 = int64(serieIDInt)
		serieID.Valid = true
	}

	insertBookQuery := `
    INSERT INTO books (
        filepath, crc32, archive_name, size, format, title, sort_title, year, plot, 
        cover_path, language_id, serie_id, last_updated
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
    `
	_, err = tx.Exec(insertBookQuery,
		bookData["filepath"], bookData["crc32"], bookData["archive_name"], bookData["size"],
		bookData["format"], bookData["title"], strings.ToUpper(bookData["title"].(string)),
		bookData["year"], bookData["plot"], bookData["cover_path"], languageID, serieID,
		time.Now().Unix(),
	)
	if err != nil {
		return fmt.Errorf("failed to insert book: %v", err)
	}

	// Добавляем авторов и жанры
	// ...

	return nil
}

func getOrCreateLanguage(tx *sql.Tx, code string) (int, error) {
	var id int
	err := tx.QueryRow("SELECT id FROM languages WHERE code = ?", code).Scan(&id)
	if err == nil {
		return id, nil
	}
	if err != sql.ErrNoRows {
		return 0, fmt.Errorf("failed to check language: %v", err)
	}

	// Если язык не найден, создаем его
	res, err := tx.Exec("INSERT INTO languages (code, name) VALUES (?, ?)", code, code)
	if err != nil {
		return 0, fmt.Errorf("failed to insert language: %v", err)
	}
	id64, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get language ID: %v", err)
	}
	return int(id64), nil
}

func getOrCreateSerie(tx *sql.Tx, name string) (int, error) {
	var id int
	err := tx.QueryRow("SELECT id FROM series WHERE name = ?", name).Scan(&id)
	if err == nil {
		return id, nil
	}
	if err != sql.ErrNoRows {
		return 0, fmt.Errorf("failed to check serie: %v", err)
	}

	// Если серия не найдена, создаем ее
	res, err := tx.Exec("INSERT INTO series (name) VALUES (?)", name)
	if err != nil {
		return 0, fmt.Errorf("failed to insert serie: %v", err)
	}
	id64, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get serie ID: %v", err)
	}
	return int(id64), nil
}

func getOrCreateAuthor(tx *sql.Tx, name string) (int, error) {
	var id int
	err := tx.QueryRow("SELECT id FROM authors WHERE name = ?", name).Scan(&id)
	if err == nil {
		return id, nil
	}
	if err != sql.ErrNoRows {
		return 0, fmt.Errorf("failed to check author: %v", err)
	}

	// Если автор не найден, создаем его
	res, err := tx.Exec("INSERT INTO authors (name) VALUES (?)", name)
	if err != nil {
		return 0, fmt.Errorf("failed to insert author: %v", err)
	}
	id64, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get author ID: %v", err)
	}
	return int(id64), nil
}

func getOrCreateGenre(tx *sql.Tx, name string) (int, error) {
	var id int
	err := tx.QueryRow("SELECT id FROM genres WHERE name = ?", name).Scan(&id)
	if err == nil {
		return id, nil
	}
	if err != sql.ErrNoRows {
		return 0, fmt.Errorf("failed to check genre: %v", err)
	}

	// Если жанр не найден, создаем его
	res, err := tx.Exec("INSERT INTO genres (name) VALUES (?)", name)
	if err != nil {
		return 0, fmt.Errorf("failed to insert genre: %v", err)
	}
	id64, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get genre ID: %v", err)
	}
	return int(id64), nil
}
