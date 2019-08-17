package model

import (
	"log"

	"github.com/achhapolia10/anjaniv2/opdatabase"
)

//JournalAddMonth Add to Month For a Journal Entry
func JournalAddMonth(je opdatabase.JournalEntry) bool {
	date := ParseDate(je.Date)
	date.Day = 1
	s, res := opdatabase.SelectMonthEntryDate(date.GetString(), je.ProductID)
	if !res {
		log.Printf("Error in reading Month Entry in JournalAddMonth()")
		return res
	}
	s.BoxIn = s.BoxIn + je.Box
	s.PacketIn = s.PacketIn + je.Packet
	opdatabase.UpdateMonthEntry(je.ProductID, s)
	return true
}

//JournalDeleteMonth Reflects Changes in Month for Deleting a Jounal ENtry
func JournalDeleteMonth(je opdatabase.JournalEntry) bool {
	date := ParseDate(je.Date)
	date.Day = 1
	s, res := opdatabase.SelectMonthEntryDate(date.GetString(), je.ProductID)
	if !res {
		log.Println("Error in reading Month Entry in JournalDeleteMonth() on Date: ", je.Date)
		return false
	}
	s.BoxIn = s.BoxIn - je.Box
	s.PacketIn = s.PacketIn - je.Packet
	opdatabase.UpdateMonthEntry(je.ProductID, s)
	return true
}

//DispatchAddMonth Reflects Changes in Month for Adding a DispatchEntry
func DispatchAddMonth(se opdatabase.StockEntry) bool {
	date := ParseDate(se.Date)
	date.Day = 1
	s, res := opdatabase.SelectMonthEntryDate(date.GetString(), se.ProductID)
	if !res {
		log.Println("Error in reading Month Entry in DispachAddMonth() on Date: ", se.Date)
		return false
	}
	s.BoxOut = s.BoxOut + se.BoxOut
	s.PacketOut = s.PacketOut + se.PacketOut
	opdatabase.UpdateMonthEntry(se.ProductID, s)
	return true
}

//DispatchDeleteMonth Reflsects Changes in Month for Deleting a DispatchEntry
func DispatchDeleteMonth(se opdatabase.StockEntry) bool { //Deprecate this
	date := ParseDate(se.Date)
	date.Day = 1
	s, res := opdatabase.SelectMonthEntryDate(date.GetString(), se.ProductID)
	if !res {
		log.Println("Error in reading Month Entry in DispachAddMonth() on Date: ", se.Date)
		return false
	}
	s.BoxOut = s.BoxOut - se.BoxOut
	s.PacketOut = s.PacketOut - se.PacketOut
	opdatabase.UpdateMonthEntry(se.ProductID, s)
	return true
}

//BalanceMonthEntries Balances the stock Entries
func BalanceMonthEntries(se *opdatabase.MonthEntry) {
	boxIn := se.BoxIn
	packetIn := se.PacketIn
	boxOut := se.BoxOut
	packetOut := se.PacketOut
	product, _ := opdatabase.SelectProductID(se.ProductID)
	box := product.BoxQuantity
	unitIn := boxIn*box + packetIn
	boxIn = unitIn / box
	packetIn = unitIn % box
	unitOut := boxOut*box + packetOut
	boxOut = unitOut / box
	packetOut = unitOut % box
	se.BoxIn = boxIn
	se.PacketIn = packetIn
	se.BoxOut = boxOut
	se.PacketOut = packetOut
}
