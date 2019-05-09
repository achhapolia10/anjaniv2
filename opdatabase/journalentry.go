package opdatabase

import (
	"log"
	"strconv"

	"github.com/achhapolia10/anjaniv2/model"
)

//NewJournalEntry Creates a new Entry in the given products Journal
func NewJournalEntry(je model.JournalEntry) {
	id := strconv.Itoa(je.ProductID)
	query := "INSERT INTO " + id + "journal (labour,date,box,packet) VALUES(?,?,?,?);"
	_, err := db.Exec(query, je.Labour, je.Date, je.Boxes, je.Packets)
	if err != nil {
		log.Println("Error inserting in journal of Product id:", id)
		log.Println(err)
	}

}

//SelectJournalEntry Selects all entries of a Pariticular Date
func SelectJournalEntry(date string, productID int) ([]model.JournalEntry, bool) {
	var je []model.JournalEntry
	query := "SELECT * FROM " + strconv.Itoa(productID) + "journal WHERE date='" + date +
		"' ORDER BY id DESC;"
	r, err := db.Query(query)
	if err != nil {
		log.Println(err)
		log.Println("Error in Selecting Journal Entries of Product ID:", productID)
		return je, false
	}
	for r.Next() {
		var e model.JournalEntry
		err = r.Scan(&(e.ID), &(e.Labour), &(e.Date), &(e.Boxes), &(e.Packets))
		if err != nil {
			log.Println(err)
			return je, false
		}
		e.ProductID = productID
		je = append(je, e)
	}
	return je, true
}

//SelectJournalEntryByID Selects a journal entry by id
func SelectJournalEntryByID(id int, productID int) (model.JournalEntry, bool) {
	query := "SELECT * FROM " + strconv.Itoa(productID) + "journal WHERE id=?;"
	var je model.JournalEntry
	r, err := db.Query(query, id)
	if err != nil {
		log.Println(err)
		log.Println("Error in Selecting Journal Entry id")
		return je, false
	}
	if r.Next() {
		err = r.Scan(&(je.ID), &(je.Labour), &(je.Date), &(je.Boxes), &(je.Packets))
		if err != nil {
			log.Println(err)
			return je, false
		}
	}
	je.ProductID = productID
	return je, true
}

//DeleteJournalEntry Deletes a Journal entry
func DeleteJournalEntry(productID int, id int) {
	query := "DELETE FROM " + strconv.Itoa(productID) + "journal where id=?;"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Println(err)
	}
}
