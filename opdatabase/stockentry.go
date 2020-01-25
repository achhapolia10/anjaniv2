package opdatabase

import (
	"log"
	"strconv"
)

//StockEntry Structure to store StockEntries
type StockEntry struct {
	ID        int    `json:"id"`
	Date      string `json:"date"`
	BoxIn     int    `json:"boxin"`
	PacketIn  int    `json:"packetin"`
	BoxOut    int    `json:"boxout"`
	PacketOut int    `json:"packetout"`
	ProductID int    `json:"productid"`
}

//SelectStockEntry select all stockentries for a given database
func SelectStockEntry(productID int) ([]StockEntry, bool) {
	var s []StockEntry
	query := `SELECT * FROM _` + strconv.Itoa(productID) + `stock;`
	R, err := db.Query(query)
	defer R.Close()
	if err != nil {
		log.Println("Error in retrieving data from stock table of product:", productID)
		log.Println(err)
		return s, false
	}
	for R.Next() {
		var se StockEntry
		err = R.Scan(&(se.ID), &(se.Date), &(se.BoxIn), &(se.PacketIn), &(se.BoxOut), &(se.PacketOut))
		if err != nil {
			log.Println("Error in Scanning a stock entry of Product:", productID)
			log.Println(err)
			return s, false
		}
		s = append(s, se)
	}
	return s, true
}

//SelectStockEntryDate selects a stockentry of a given date
func SelectStockEntryDate(date string, productID int) (StockEntry, bool) {
	query := `SELECT * FROM _` + strconv.Itoa(productID) + `stock WHERE date=?;`
	var se StockEntry
	R, err := db.Query(query, date)
	defer R.Close()
	if err != nil {
		log.Println("Error in retrieving data from stock tabke of product:", productID)
		log.Println(err)
		return se, false
	}
	if R.Next() {
		err = R.Scan(&(se.ID), &(se.Date), &(se.BoxIn), &(se.PacketIn), &(se.BoxOut), &(se.PacketOut))
		se.ProductID = productID
		if err != nil {
			log.Println("Error in Scanning a stock entry of Product:", productID)
			log.Println(err)
			log.Println(se)
			return se, false
		}
	} else {
		id, _ := AddStockEntryDate(date, productID)
		se = StockEntry{
			id, date, 0, 0, 0, 0, productID,
		}
	}
	return se, true
}

//AddStockEntryDate Creates a stock entry at a given date for a productID
//Returns the last insert id and a bool
func AddStockEntryDate(date string, productID int) (int, bool) {
	query := `INSERT INTO _` + strconv.Itoa(productID) + "stock (date,boxIn,packetIn,boxOut,packetOut) VALUES( ? , 0,0,0,0);"
	qr, err := db.Exec(query, date)
	if err != nil {
		log.Println("Error in Creating a Stock entry of Product:", productID)
		log.Println(err)
		return -1, false
	}
	id, _ := qr.LastInsertId()

	return int(id), true
}

//UpdateStockEntry Updates the stock entry at a particular date for a productID
//Returns a bool
func UpdateStockEntry(productID int, se StockEntry) bool {
	query := "UPDATE _" + strconv.Itoa(productID) + "stock SET boxIn= ? ,packetIn= ? , boxOut= ? , packetOut= ? WHERE date= ? ;"
	_, err := db.Exec(query, se.BoxIn, se.PacketIn, se.BoxOut, se.PacketOut, se.Date)
	if err != nil {
		log.Println("Error in creating a stock entry of Product: ", productID)
		log.Println(err)
		return false
	}
	return true
}
