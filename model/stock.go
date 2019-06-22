// model file - contains the buisiness logic for calculation of the stock

package model

import (
	"fmt"
	"log"
	"time"

	"github.com/achhapolia10/anjaniv2/opdatabase"
)

//Stock struct for the stock data
type Stock struct {
	OBox    int
	OPacket int
	CBox    int
	CPacket int
}

/**
Stock Model
Calculates the stock data
Data about the products
Getting the product Details
Calculating the stocks for  a particular date
**/

//ProductStock retruns the data about the stock
func ProductStock(fDate, tDate, start time.Time, product opdatabase.Product) {
	obox, opacket := product.OpeningBox, product.OpeningPacket
	// Calculation of the opening boxes from the fiscal to fDate
	for start.Before(fDate) {
		f := parseTime(start)
		se, _ := opdatabase.SelectStockEntryDate(f.getString(), product.ID)
		obox += se.BoxIn - se.BoxOut
		opacket += se.PacketIn - se.PacketOut
		start = start.AddDate(0, 0, 1)
	}
	cbox, cpacket := obox, opacket
	//Calculation of the stock from fdate to tdate
	for start.Before(tDate) {
		f := parseTime(start)
		se, _ := opdatabase.SelectStockEntryDate(f.getString(), product.ID)
		cbox += se.BoxIn - se.BoxOut
		cpacket += se.PacketIn - se.PacketOut
		start = start.AddDate(0, 0, 1)
	}

	fmt.Println(obox, opacket, start)

}

//AllStock Returns the data for all products
func AllStock(f, t string) string {

	//Product List
	products, res := opdatabase.SelectProduct()
	if !res {
		log.Printf("Error in Getting products from the database")
	}

	//Dates parsing and whatever the fuck is here
	from, to := parseDate(f), parseDate(t)
	fromDate := time.Date(from.year, from.getMonth(), from.day, 0, 0, 0, 0, time.Now().Location())
	toDate := time.Date(to.year, to.getMonth(), to.day, 0, 0, 0, 0, time.Now().Location())
	fiscal := getFiscal(fromDate)

	for _, product := range products {
		ProductStock(fromDate, toDate, fiscal, product)
	}
	return fmt.Sprintln(from, to)
}

//getFiscal returns the date the fical year starts for the date
func getFiscal(t time.Time) time.Time {
	location := time.Now().Location()
	month := int(t.Month())
	if month < 4 {
		return time.Date(t.Year()-1, time.April, 1, 0, 0, 0, 0, location)
	}
	return time.Date(t.Year(), time.April, 1, 0, 0, 0, 0, location)
}
