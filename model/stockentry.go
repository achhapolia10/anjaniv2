package model

//StockEntry Structure to store StockEntries
type StockEntry struct {
	Date      string
	BoxesIn   int
	PacketIn  int
	BoxesOut  int
	PacketOut int
	Product   *Product
}
