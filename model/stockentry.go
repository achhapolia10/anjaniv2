package model

//StockEntry Structure to store StockEntries
type StockEntry struct {
	ID        int    `json:"id"`
	Date      string `json:"date"`
	BoxIn     int    `json:"box-in"`
	PacketIn  int    `json:"packet-in"`
	BoxOut    int    `json:"box-out"`
	PacketOut int    `json:"packet-out"`
	ProductID int    `json:"product-id"`
}
