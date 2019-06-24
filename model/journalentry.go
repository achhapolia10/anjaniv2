package model

import (
	"fmt"

	"github.com/achhapolia10/anjaniv2/opdatabase"
)

//GetAllJournalEntry Gets all Journal entry for a productID and and date
func GetAllJournalEntry(date string, productID int) ([]opdatabase.JournalEntry, int, int, bool) {
	je, box, packet, res := opdatabase.SelectJournalEntry(date, productID)
	fmt.Println(box, packet)
	box, packet = balanceJournalEntries(box, packet, productID)

	return je, box, packet, res
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

func balanceJournalEntries(b, p, i int) (int, int) {
	pr, _ := opdatabase.SelectProductID(i)
	p += b * pr.BoxQuantity
	b = p / pr.BoxQuantity
	p = p % pr.BoxQuantity
	return b, p
}
