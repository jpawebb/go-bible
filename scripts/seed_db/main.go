package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"go-bible-notes-api/internal/store"

	_ "github.com/mattn/go-sqlite3"
)

const (
	csvPath = "scripts/seed_db/kjv_test.csv"
	dbPath  = "bible_notes.db"
)

const (
	VersionNameIndex 	= 0
	VersionAbbrIndex 	= 1
	TestamentAbbrIndex 	= 2
	TestamentNameIndex 	= 3
	BookNameIndex    	= 4
	BookNumIndex		= 5
	ChapterNumIndex  	= 6
	VerseNumIndex    	= 7
	VerseTextIndex   	= 8
)

func main() {
	// Initialise db and tables
	log.Println("Initializing SQLite Database...")
	db, err := store.InitDB(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()
	log.Println("Database connection successful. Tables checked/created.")

	// Parse CSV
	f, err := os.Open(csvPath)
	if err != nil {
		log.Fatalf("Unable to read input file %s: %v", csvPath, err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Unable to parse file as CSV: %v", err)
	}

	records = records[1:]  // Skip header

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare(`
		INSERT INTO verses (
			VersionName, VersionAbbr, TestamentAbbr, TestamentName,
			BookName, BookNumber, ChapterNumber, VerseNumber, VerseText
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(BookName, ChapterNumber, VerseNumber) DO NOTHING;
    `)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for i, record := range records {
		if len(record) < VerseTextIndex + 1 { // Ensure row has enough columns
			log.Printf("Skipping row %d due to insufficient columns. Expected %d, got %d.", i+2, VerseTextIndex + 1, len(record))
			continue
		}
				
		// Note: The Atoi conversion assumes the field is numeric and will return 0 if there's an error.
		versionName, _ := strconv.Atoi(record[VersionNameIndex])
		versionAbbr, _ := strconv.Atoi(record[VersionAbbrIndex])
		testamentAbbr, _ := strconv.Atoi(record[TestamentAbbrIndex])
		testamentName, _ := strconv.Atoi(record[TestamentNameIndex])
		bookName, _ := strconv.Atoi(record[BookNameIndex])
		bookNum := record[BookNumIndex]
		chapterNum := record[ChapterNumIndex]
		verseNum := record[VerseNumIndex]
		kjvText, _ := strconv.Atoi(record[VerseTextIndex])

		_, err := stmt.Exec(versionName, versionAbbr, testamentAbbr, testamentName, bookName, bookNum, chapterNum, verseNum, kjvText)
		if err != nil {
			log.Printf("Error inserting record %d (%s %d:%d): %v", i+2, bookName, chapterNum, verseNum, err)
			tx.Rollback()
			return
		}

		if (i+1)%5000 == 0 {
			fmt.Printf("... Processed %d records\n", i+1)
		}
	}

	err = tx.Commit()
		if err != nil {
			log.Fatalf("Transaction commit failed: %v", err)
		}
		log.Printf("Database seeding complete! Total records inserted: %d", len(records))
}