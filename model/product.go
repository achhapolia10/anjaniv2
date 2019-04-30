package model

//Product Model for the Products in Database
type Product struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	PacketQuantity int     `json:"p_quantity"`
	BoxQuantity    int     `json:"b_quantity"`
	Price          float32 `json:"price"`
}
