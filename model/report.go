package model

import (
	"log"
	"math"
	"time"

	"github.com/achhapolia10/inventory-manager/opdatabase"
)

// Report is a struct for daily Report
type Report struct {
	Product  opdatabase.Product `json:"product"`
	BoxIn    int                `json:"boxin"`
	PacketIn int                `json:"packetin"`
	Carry    int                `json:"carry"`
	Raw      float64            `json:"raw"`
	Plastic  float64            `json:"plastic"`
}

// GetDailyReport returns the daily report for a day
func GetDailyReport(fdate string) []Report {
	var reports []Report
	from := ParseDate(fdate)

	products, _ := opdatabase.SelectProduct()

	for _, p := range products {
		log.Println(p.ID, p.Name, p.BoxQuantity, p.PacketQuantity, p.Weight)
		se, _ := opdatabase.SelectStockEntryDate(fdate, p.ID)

		if se.BoxIn != 0 || se.PacketIn != 00 {
			date := time.Date(from.Year, from.GetMonth(), from.Day, 0, 0, 0, 0, time.Now().Location())
			s := ProductStock(date, date, time.Date(2022, time.April, 01, 0, 0, 0, 0, time.Now().Location()), p)
			s.Balance()
			se.PacketIn += s.OPacket
			BalanceStockEntries(&se)
			report := Report{
				Product:  p,
				BoxIn:    se.BoxIn,
				PacketIn: se.PacketIn,
				Carry:    s.OPacket,
				Raw:      math.Round(float64(se.BoxIn*p.BoxQuantity*p.PacketQuantity)*p.Weight) / 1000.0,
				Plastic:  math.Round(float64(se.BoxIn*p.BoxQuantity)/220.0*1000.0) / 1000.0,
			}
			reports = append(reports, report)
		}
	}
	return reports
}
