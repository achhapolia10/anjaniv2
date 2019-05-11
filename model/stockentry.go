package model

import (
	"log"

	"github.com/achhapolia10/anjaniv2/opdatabase"
)

//JournalAddStock Add to Stock For a Journal Entry
func JournalAddStock(je opdatabase.JournalEntry) bool {
	s, res := opdatabase.SelectStockEntryDate(je.Date, je.ProductID)
	if !res {
		log.Printf("Error in reading Stock Entry in AddStock()")
		return res
	}
	s.BoxIn = s.BoxIn + je.Box
	s.PacketIn = s.PacketIn + je.Packet
	opdatabase.UpdateStockEntry(je.ProductID, s)
	return true
}

//JournalDeleteStock Reflects Changes in Stock for Deleting a Jounal ENtry
func JournalDeleteStock(je opdatabase.JournalEntry) bool {
	s, res := opdatabase.SelectStockEntryDate(je.Date, je.ProductID)
	if !res {
		log.Printf("Error in reading Stock Entry in Addstock() on Date: ", je.Date)
		return false
	}
	s.BoxIn = s.BoxIn - je.Box
	s.PacketIn = s.PacketIn - je.Packet
	opdatabase.UpdateStockEntry(je.ProductID, s)
	return true
}
