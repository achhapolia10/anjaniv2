package model

import "github.com/achhapolia10/anjaniv2/opdatabase"

//GetAllProduct returns all products
func GetAllProduct() ([]opdatabase.Product, bool) {
	p, res := opdatabase.SelectProduct()
	return p, res
}

//CreateProduct Creates A new Product
func CreateProduct(product opdatabase.Product) {
	opdatabase.AddProduct(product)
}
