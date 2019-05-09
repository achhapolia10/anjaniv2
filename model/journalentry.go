package model

//JournalEntry for Journal Entries
type JournalEntry struct {
	ID        int    `json:"id"`
	Labour    string `json:"labour"`
	Date      string `json:"date"`
	Boxes     int    `json:"box"`
	Packets   int    `json:"packet"`
	ProductID int    `json:"product"`
}
