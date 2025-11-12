package data

import "time"

type Book struct {
    ID   			int
    Name 			string
	Chapters		int
}


type Verse struct {
    ID            int
    BookID        int
    ChapterNumber int
    VerseNumber   int
    TranslationID int
    Text          string
}


type Note struct {
    ID      		int `json:"id"`
    VerseID 		int `json:"verse_id"`
    UserID  		int `json:"user_id"`
    Content 		string `json:"content"`
    CreatedAt 		time.Time `json:"created_at"`
    UpdatedAt 		time.Time `json:"updated_at"`
}