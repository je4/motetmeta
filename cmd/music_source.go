package main

type MotetMusicSource struct {
	Folio        string
	Source       string
	SourceFormat string
	MusicSource  *MusicSource
}
type MusicSource struct {
	ID                int
	Title             string
	Detail            string
	DetailFormat      string
	SourceType        string
	Description       string
	DescriptionFormat string
	Facsimile         string
	FacsimileFormat   string
}
