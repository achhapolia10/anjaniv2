package opdatabase

import (
	"log"
	"strconv"
)

//JournalEntry for Journal Entries
type JournalEntry struct {
	ID        int    `json:"id"`
	Labour    string `json:"labour"`
	Date      string `json:"date"`
	Box       int    `json:"box"`
	Packet    int    `json:"packet"`
	ProductID int    `json:"product"`
}

//NewJournalEntry Creates a new Entry in the given products Journal
func NewJournalEntry(je JournalEntry) {
	id := strconv.Itoa(je.ProductID)
	query := "INSERT INTO " + id + "journal (labour,date,box,packet) VALUES(?,?,?,?);"
	_, err := db.Exec(query, je.Labour, je.Date, je.Box, je.Packet)
	if err != nil {
		log.Println("Error inserting in journal of Product id:", id)
		log.Println(err)
	}

}

//SelectJournalEntry Selects all entries of a Pariticular Date
func SelectJournalEntry(date string, productID int) ([]JournalEntry, int, int, bool) {
	var je []JournalEntry
	var box, packet int
	box, packet = 0, 0
	query := "SELECT * FROM " + strconv.Itoa(productID) + "journal WHERE date='" + date +
		"' ORDER BY id DESC;"
	r, err := db.Query(query)
	defer r.Close()
	if err != nil {
		log.Println(err)
		log.Println("Error in Selecting Journal Entries of Product ID:", productID)
		return je, 0, 0, false
	}
	for r.Next() {
		var e JournalEntry
		err = r.Scan(&(e.ID), &(e.Labour), &(e.Date), &(e.Box), &(e.Packet))
		if err != nil {
			log.Println(err)
			return je, 0, 0, false
		}
		box = box + e.Box
		packet = packet + e.Packet
		e.ProductID = productID
		je = append(je, e)
	}
	return je, box, packet, true
}

//SelectJournalEntryByID Selects a journal entry by id
func SelectJournalEntryByID(id int, productID int) (JournalEntry, bool) {
	query := "SELECT * FROM " + strconv.Itoa(productID) + "journal WHERE id=?;"
	var je JournalEntry
	r, err := db.Query(query, id)
	defer r.Close()
	if err != nil {
		log.Println(err)
		log.Println("Error in Selecting Journal Entry id")
		return je, false
	}
	if r.Next() {
		err = r.Scan(&(je.ID), &(je.Labour), &(je.Date), &(je.Box), &(je.Packet))
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

//SelectJournalEntryMap Selects all entries of a Pariticular Date
func SelectJournalEntryMap(date string, productID int) (map[string]JournalEntry, bool) {
	je := make(map[string]JournalEntry)
	query := "SELECT * FROM " + strconv.Itoa(productID) + "journal WHERE date='" + date +
		"' ORDER BY id DESC;"
	r, err := db.Query(query)
	defer r.Close()
	if err != nil {
		log.Println(err)
		log.Println("Error in Selecting Journal Entries of Product ID:", productID)
		return je, false
	}
	for r.Next() {
		var e JournalEntry
		err = r.Scan(&(e.ID), &(e.Labour), &(e.Date), &(e.Box), &(e.Packet))
		if err != nil {
			log.Println(err)
			return je, false
		}

		e.ProductID = productID
		je[e.Labour] = e
	}

	return je, true
}
