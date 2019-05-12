package model

import (
	"log"

	"github.com/achhapolia10/anjaniv2/opdatabase"
)

//JournalAddStock Add to Stock For a Journal Entry
func JournalAddStock(je opdatabase.JournalEntry) bool {
	s, res := opdatabase.SelectStockEntryDate(je.Date, je.ProductID)
	if !res {
		log.Printf("Error in reading Stock Entry in JournalAddStock()")
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
		log.Println("Error in reading Stock Entry in JournalDeleteStock() on Date: ", je.Date)
		return false
	}
	s.BoxIn = s.BoxIn - je.Box
	s.PacketIn = s.PacketIn - je.Packet
	opdatabase.UpdateStockEntry(je.ProductID, s)
	return true
}

//DispatchAddStock Reflects Changes in Stock for Adding a DispatchEntry
func DispatchAddStock(se opdatabase.StockEntry) bool {
	s, res := opdatabase.SelectStockEntryDate(se.Date, se.ProductID)
	if !res {
		log.Println("Error in reading Stock Entry in DispachAddStock() on Date: ", se.Date)
		return false
	}
	s.BoxOut = s.BoxOut + se.BoxOut
	s.PacketOut = s.PacketOut + se.PacketOut
	opdatabase.UpdateStockEntry(se.ProductID, s)
	return true
}

//DispatchDeleteStock Reflsects Changes in Stock for Deleting a DispatchEntry
func DispatchDeleteStock(se opdatabase.StockEntry) bool {
	s, res := opdatabase.SelectStockEntryDate(se.Date, se.ProductID)
	if !res {
		log.Println("Error in reading Stock Entry in DispachAddStock() on Date: ", se.Date)
		return false
	}
	s.BoxOut = s.BoxOut - se.BoxOut
	s.PacketOut = s.PacketOut - se.PacketOut
	opdatabase.UpdateStockEntry(se.ProductID, s)
	return true
}
