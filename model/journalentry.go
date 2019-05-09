package model

import "github.com/achhapolia10/anjaniv2/opdatabase"

//JournalEntry for Journal Entries
type JournalEntry struct {
	ID        int    `json:"id"`
	Labour    string `json:"labour"`
	Date      string `json:"date"`
	Box       int    `json:"box"`
	Packet    int    `json:"packet"`
	ProductID int    `json:"product"`
}

//GetAllJournalEntry Gets all Journal entry for a productID and and date
func GetAllJournalEntry(date string, productID int) ([]opdatabase.JournalEntry, bool) {
	je, res := opdatabase.SelectJournalEntry(date, productID)
	return je, res
}

//DeleteJournalEntry Deletes a Journal Entry of an productID with a particuar id
func DeleteJournalEntry(id int, productID int) {
	opdatabase.DeleteJournalEntry(productID, id)
}
