package store

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dbPath string) (*sql.DB, error) {
	// Connection
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Test Connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	// 3. Create Tables
	err = createTables(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS verses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,

		VersionName	TEXT NOT NULL,
		VersionAbbr	TEXT NOT NULL,
		TestamentAbbr TEXT NOT NULL,
		TestamentName TEXT NOT NULL,
		BookName TEXT NOT NULL,
		BookNumber INTEGER NOT NULL,
		ChapterNumber INTEGER NOT NULL,
		VerseNumber INTEGER NOT NULL,
		VerseText STRING NOT NULL,
		UNIQUE (BookName, ChapterNumber, VerseNumber)
	);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("Error creating tables: %v\n", err)
		return err
	}
	return nil
}