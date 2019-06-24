package model

import "github.com/achhapolia10/anjaniv2/opdatabase"

//GetAllProduct returns all products
func GetAllProduct() ([]opdatabase.Product, bool) {
	p, res := opdatabase.SelectProduct()
	return p, res
}

//GetProduct returns product with id
func GetProduct(id int) opdatabase.Product {
	p, _ := opdatabase.SelectProductID(id)
	return p
}

//CreateProduct Creates A new Product
func CreateProduct(product opdatabase.Product) {
	opdatabase.AddProduct(product)
}
