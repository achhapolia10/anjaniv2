// model file - contains the business logic for calculation of the stock

package model

import (
	"fmt"
	"log"
	"time"

	"github.com/achhapolia10/anjaniv2/opdatabase"
)

//Stock struct for the stock data
type Stock struct {
	OBox      int                `json:"obox"`
	OPacket   int                `json:"opacket"`
	InBox     int                `json:"inbox"`
	OutBox    int                `json:"outbox"`
	InPacket  int                `json:"inpacket"`
	OutPacket int                `json:"outpacket"`
	CBox      int                `json:"cbox"`
	CPacket   int                `json:"cpacket"`
	Product   opdatabase.Product `json:"product"`
}

//StockDetails Struct for the stock details and entries
type StockDetails struct {
	Stock   Stock                   `json:"stock"`
	Entries []opdatabase.StockEntry `json:"entries"`
}

/**
Stock Model
Calculates the stock data
Data about the products
Getting the product Details
Calculating the stocks for  a particular date
**/

//ProductStock returns the data about the stock
func ProductStock(fDate, tDate, start time.Time, product opdatabase.Product) Stock {

	s := Stock{
		OBox:      product.OpeningBox,
		OPacket:   product.OpeningPacket,
		InBox:     0,
		OutBox:    0,
		CBox:      0,
		CPacket:   0,
		InPacket:  0,
		OutPacket: 0,
		Product:   product,
	}

	// Calculation of the opening boxes from the fiscal to fDate
	flag := true //Flag is for the Month Check
	for start.Before(fDate) {
		var temp time.Time
		if flag {
			temp = start.AddDate(0, 1, -1)

			if temp.Before(fDate) {
				d := ParseTime(start)
				me, _ := opdatabase.SelectMonthEntryDate(d.GetString(), product.ID)
				s.OBox = s.OBox + me.BoxIn - me.BoxOut
				s.OPacket = s.OPacket + me.PacketIn - me.PacketOut
				start = temp.AddDate(0, 0, 1)
				fmt.Println(start)
			} else {
				flag = false
			}
		} else {
			d := ParseTime(start)
			se, _ := opdatabase.SelectStockEntryDate(d.GetString(), product.ID)
			s.OBox = s.OBox + se.BoxIn - se.BoxOut
			s.OPacket = s.OPacket + se.PacketIn - se.PacketOut
			start = start.AddDate(0, 0, 1)
		}
	}

	//Calculation of the stock from fdate to tdate
	for start.Before(tDate) || start.Equal(tDate) {
		var temp time.Time
		if flag {
			temp = start.AddDate(0, 1, -1)

			if temp.Before(tDate) {
				d := ParseTime(start)
				me, _ := opdatabase.SelectMonthEntryDate(d.GetString(), product.ID)
				s.InBox = s.InBox + me.BoxIn
				s.OutBox = s.OutBox + me.BoxOut
				s.InPacket = s.InPacket + me.PacketIn
				s.OutPacket = s.OutPacket + me.PacketOut
				start = temp.AddDate(0, 0, 1)
				fmt.Println(start)
			} else {
				flag = false
			}
		} else {
			d := ParseTime(start)
			se, _ := opdatabase.SelectStockEntryDate(d.GetString(), product.ID)
			s.InBox = s.InBox + se.BoxIn
			s.OutBox = s.OutBox + se.BoxOut
			s.InPacket = s.InPacket + se.PacketIn
			s.OutPacket = s.OutPacket + se.PacketOut
			start = start.AddDate(0, 0, 1)
			fmt.Println(start)
			if start.Day() == 1 {
				flag = true
			}
		}
		s.CBox = s.OBox + s.InBox - s.OutBox
		s.CPacket = s.OPacket + s.InPacket - s.OutPacket
	}
	return s
}

//AllStock Returns the data for all products
func AllStock(f, t string) map[int]Stock {

	stocks := make(map[int]Stock)

	//Product List
	products, res := opdatabase.SelectProduct()
	if !res {
		log.Printf("Error in Getting products from the database")
	}

	//Dates parsing and whatever the fuck is here
	from, to := ParseDate(f), ParseDate(t)
	fromDate := time.Date(from.Year, from.GetMonth(), from.Day, 0, 0, 0, 0, time.Now().Location())
	toDate := time.Date(to.Year, to.GetMonth(), to.Day, 0, 0, 0, 0, time.Now().Location())
	fiscal := time.Date(2019, time.April, 1, 0, 0, 0, 0, time.Now().Location())

	for _, product := range products {
		s := ProductStock(fromDate, toDate, fiscal, product)
		s.Balance()
		stocks[product.ID] = s
	}
	return stocks
}

//ProductStockDetails returns the details of the stock
func ProductStockDetails(f, t string, id int) StockDetails {
	var entries []opdatabase.StockEntry
	product, _ := opdatabase.SelectProductID(id)
	from, to := ParseDate(f), ParseDate(t)
	fromDate := time.Date(from.Year, from.GetMonth(), from.Day, 0, 0, 0, 0, time.Now().Location())
	toDate := time.Date(to.Year, to.GetMonth(), to.Day, 0, 0, 0, 0, time.Now().Location())
	fiscal := time.Date(2019, time.April, 1, 0, 0, 0, 0, time.Now().Location())

	stock := ProductStock(fromDate, toDate, fiscal, product)
	stock.Balance()

	for fromDate.Before(toDate) || fromDate.Equal(toDate) {
		d := ParseTime(fromDate)
		entry, _ := opdatabase.SelectStockEntryDate(d.GetString(), product.ID)
		if !(entry.BoxIn == 0 && entry.BoxOut == 0 && entry.PacketIn == 0 && entry.PacketOut == 0) {
			BalanceStockEntries(&entry)
			entries = append(entries, entry)
		}
		fromDate = fromDate.AddDate(0, 0, 1)
	}
	return StockDetails{
		Stock:   stock,
		Entries: entries,
	}

}

//Balance Balances the stock details
func (s *Stock) Balance() {
	s.OPacket += s.OBox * s.Product.BoxQuantity
	s.OBox = s.OPacket / s.Product.BoxQuantity
	s.OPacket = s.OPacket % s.Product.BoxQuantity

	s.InPacket += s.InBox * s.Product.BoxQuantity
	s.InBox = s.InPacket / s.Product.BoxQuantity
	s.InPacket = s.InPacket % s.Product.BoxQuantity

	s.OutPacket += s.OutBox * s.Product.BoxQuantity
	s.OutBox = s.OutPacket / s.Product.BoxQuantity
	s.OutPacket = s.OutPacket % s.Product.BoxQuantity

	s.CPacket += s.CBox * s.Product.BoxQuantity
	s.CBox = s.CPacket / s.Product.BoxQuantity
	s.CPacket = s.CPacket % s.Product.BoxQuantity
}
