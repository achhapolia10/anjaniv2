package opdatabase

import (
	"log"
	"strconv"

	"github.com/achhapolia10/anjaniv2/model"
)

//NewJournalEntry Creates a new Entry in the given products Journal
func NewJournalEntry(je model.JournalEntry) {
	id := strconv.Itoa(je.Product.ID)
	query := "INSERT INTO " + id + "_journal (labour,date,box,packet) VALUES(?,?,?,?)"
	_, err := db.Exec(query, je.Labour, je.Date, je.Boxes, je.Packets)
	if err != nil {
		log.Println("Error inserting in journal of Product id:", id)
		log.Println(err)
	}

}
