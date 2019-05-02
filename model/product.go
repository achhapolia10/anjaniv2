package model

//Product Model for the Products in Database
type Product struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	PacketQuantity int     `json:"packet"`
	BoxQuantity    int     `json:"box"`
	Price          float64 `json:"price"`
	OpeningBoxes   int     `json:"oboxes"`
	OpeningPackets int     `json:"opackets"`
}
