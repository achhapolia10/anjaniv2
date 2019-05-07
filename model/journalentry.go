package model

//Labour type for Labour Names
type Labour string

//JournalEntry for Journal Entries
type JournalEntry struct {
	ID      int      `json:"id"`
	Labour  Labour   `json:"labour"`
	Date    string   `json:"date"`
	Boxes   int      `json:"box"`
	Packets int      `json:"packet"`
	Product *Product `json:"product"`
}

//BalanceEntry balances the boxes and packets
func (e *JournalEntry) BalanceEntry() {
	p := e.Product
	e.Boxes = e.Boxes + e.Packets%p.BoxQuantity
	e.Packets = e.Packets / p.BoxQuantity
}

//CalculateTotalUnit Calculate the total Unit
func (e *JournalEntry) CalculateTotalUnit() int {
	p := e.Product
	var units int
	units = (e.Boxes*p.BoxQuantity + e.Packets) * p.PacketQuantity
	return units
}
