package model

//Labour type for Labour Names
type Labour string

//Entry Model for each entry
type Entry struct {
	ID      string  `json:"id"`
	Labour  Labour  `json:"labour"`
	Date    string  `json:"date"`
	Product Product `json:"product"`
	Packets int     `json:"packets"`
	Boxes   int     `json:"boxes"`
}

//BalanceEntry To Balance the Entry
func (entry *Entry) BalanceEntry() {
	e := *entry
	b := e.Boxes
	p := e.Packets
	bq := e.Product.BoxQuantity
	b = b + p/bq
	p = p % bq
	e.Boxes = b
	e.Packets = p
}
