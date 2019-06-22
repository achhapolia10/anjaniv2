package opdatabase

import (
	"log"
	"strconv"
)

//MonthEntry Structure to store MonthEntries
type MonthEntry struct {
	ID        int    `json:"id"`
	Date      string `json:"date"`
	BoxIn     int    `json:"boxin"`
	PacketIn  int    `json:"packetsin"`
	BoxOut    int    `json:"boxesout"`
	PacketOut int    `json:"packetout"`
	ProductID int    `json:"productid"`
}

//SelectMonthEntry select all monthentries for a given database
func SelectMonthEntry(productID int) ([]MonthEntry, bool) {
	var s []MonthEntry
	query := `SELECT * FROM ` + strconv.Itoa(productID) + `month;`
	R, err := db.Query(query)
	if err != nil {
		log.Println("Error in retrieving data from month tabke of product:", productID)
		log.Println(err)
		return s, false
	}
	for R.Next() {
		var se MonthEntry
		err = R.Scan(&(se.ID), &(se.Date), &(se.BoxIn), &(se.PacketIn), &(se.BoxOut), &(se.PacketOut))
		if err != nil {
			log.Println("Error in Scanning a month entry of Product:", productID)
			log.Println(err)
			return s, false
		}
		s = append(s, se)
	}
	return s, true
}

//SelectMonthEntryDate selects a monthentry of a given date
func SelectMonthEntryDate(date string, productID int) (MonthEntry, bool) {
	query := `SELECT * FROM ` + strconv.Itoa(productID) + `month WHERE date=?;`
	var se MonthEntry
	R, err := db.Query(query, date)
	if err != nil {
		log.Println("Error in retrieving data from month tabke of product:", productID)
		log.Println(err)
		return se, false
	}
	if R.Next() {
		err = R.Scan(&(se.ID), &(se.Date), &(se.BoxIn), &(se.PacketIn), &(se.BoxOut), &(se.PacketOut))
		se.ProductID = productID
		if err != nil {
			log.Println("Error in Scanning a month entry of Product:", productID)
			log.Println(err)
			log.Println(se)
			return se, false
		}
	} else {
		id, _ := AddMonthEntryDate(date, productID)
		se = MonthEntry{
			id, date, 0, 0, 0, 0, productID,
		}
	}
	return se, true
}

//AddMonthEntryDate Creates a month entry at a given date for a productID
//Returns the last insert id and a bool
func AddMonthEntryDate(date string, productID int) (int, bool) {
	query := `INSERT INTO ` + strconv.Itoa(productID) + "month (date,boxIn,packetIn,boxOut,packetOut) VALUES( ? , 0,0,0,0);"
	qr, err := db.Exec(query, date)
	if err != nil {
		log.Println("Error in Creating a Month entry of Product:", productID)
		log.Println(err)
		return -1, false
	}
	id, _ := qr.LastInsertId()

	return int(id), true
}

//UpdateMonthEntry Updates the month entry at a particular date for a productID
//Returns a bool
func UpdateMonthEntry(productID int, se MonthEntry) bool {
	query := "UPDATE " + strconv.Itoa(productID) + "month SET boxIn= ? ,packetIn= ? , boxOut= ? , packetOut= ? WHERE date= ? ;"
	_, err := db.Exec(query, se.BoxIn, se.PacketIn, se.BoxOut, se.PacketOut, se.Date)
	if err != nil {
		log.Println("Error in creating a month entry of Product: ", productID)
		log.Println(err)
		return false
	}
	return true
}
