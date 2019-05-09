package model

//Labour type for Labour Names
type Labour string

//JournalEntry for Journal Entries
type JournalEntry struct {
	ID        int    `json:"id"`
	Labour    Labour `json:"labour"`
	Date      string `json:"date"`
	Boxes     int    `json:"box"`
	Packets   int    `json:"packet"`
	ProductID int    `json:"product"`
}
