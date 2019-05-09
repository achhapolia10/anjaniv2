package model

import "github.com/achhapolia10/anjaniv2/opdatabase"

//Product Model for the Products in Database
type Product struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	PacketQuantity int     `json:"packet"`
	BoxQuantity    int     `json:"box"`
	Price          float64 `json:"price"`
	OpeningBox     int     `json:"obox"`
	OpeningPacket  int     `json:"opacket"`
}

//GetAllProduct returns all products
func GetAllProduct() ([]opdatabase.Product, bool) {
	p, res := opdatabase.SelectProduct()
	return p, res
}

//CreateProduct Creates A new Product
func CreateProduct(product opdatabase.Product) {
	opdatabase.AddProduct(product)
}
