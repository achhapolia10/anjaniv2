package model

//StockEntry Structure to store StockEntries
type StockEntry struct {
	ID         int    `json:"id"`
	Date       string `json:"date"`
	BoxesIn    int    `json:"boxes-in"`
	PacketsIn  int    `json:"packets-in"`
	BoxesOut   int    `json:"boxes-out"`
	PacketsOut int    `json:"packets-out"`
	ProductID  int    `json:"product-id"`
}
