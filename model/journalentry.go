package model

import "github.com/achhapolia10/anjaniv2/opdatabase"

//GetAllJournalEntry Gets all Journal entry for a productID and and date
func GetAllJournalEntry(date string, productID int) ([]opdatabase.JournalEntry, bool) {
	je, res := opdatabase.SelectJournalEntry(date, productID)
	return je, res
}

//DeleteJournalEntry Deletes a Journal Entry of an productID with a particuar id
func DeleteJournalEntry(id int, productID int) {
	je, res := opdatabase.SelectJournalEntryByID(id, productID)
	if res {
		opdatabase.DeleteJournalEntry(productID, id)
		JournalDeleteStock(je)
		JournalDeleteMonth(je)
	}
}

//CreateJournalEntry Creates a new Journal Entry
func CreateJournalEntry(je opdatabase.JournalEntry) {
	opdatabase.NewJournalEntry(je)
	JournalAddStock(je)
	JournalAddMonth(je)
}
