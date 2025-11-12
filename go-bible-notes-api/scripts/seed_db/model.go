package main

type CSVVerse struct {
	VersionName			string `csv:"VersionName"`
	VersionAbbr			string `csv:"VersionAbbr"`
	TestamentAbbr		string `csv:"TestamentAbbr"`
	TestamentName		string `csv:"TestamentAbbr"`
	BookName			string `csv:"BookName"`
	BookNumber			int `csv:"BookNumber"`
	ChapterNumber		int `csv:"ChapterNumber"`
	VerseNumber			int `csv:"VerseNumber"`
	VerseText			string `csv:"VerseText"`
}