package opdatabase

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/achhapolia10/anjaniv2/model"
)

var db *sql.DB

//CreateProductTable creates a product table if it's not already created
func CreateProductTable() {
	query := `create table products(
		productID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(50) NOT NULL, 
		packetQuantity SMALLINT NOT NULL,
		boxQuantity SMALLINT NOT NULL,
		price DECIMAL(4,2) NOT NULL,
		oboxes INT NOT NULL,
		opackets INT NOT NULL);`

	fmt.Print("Creating Product Table : ")
	_, err := db.Exec(query)
	if err != nil {
		fmt.Println("Product Table already exists")

	} else {
		fmt.Println("Product Table Created")
	}
}

//AddProduct adds a new Product to table products
func AddProduct(p model.Product) {
	query := "insert into products " +
		"(name,packetQuantity,boxQuantity,price,oboxes,opackets)" +
		"values	(?,?,?,?,?,?)"
	r, err := db.Exec(query, p.Name, p.PacketQuantity, p.BoxQuantity, p.Price, p.OpeningBoxes, p.OpeningPackets)
	if err != nil {
		log.Fatal(err)
	} else {
		id, _ := r.LastInsertId()
		log.Printf("Added product at index: %v", id)
	}
}

//SelectProduct get products from the database
func SelectProduct() ([]model.Product, bool) {
	var p []model.Product
	query := `SELECT * FROM products`
	r, err := db.Query(query)
	if err != nil {
		fmt.Println("Can't get the products from the product table")
		log.Fatal(err)
		return p, false
	}
	for r.Next() {
		var (
			product model.Product
		)
		r.Scan(&(product.ID), &(product.Name), &(product.PacketQuantity), &(product.BoxQuantity),
			&(product.Price), &(product.OpeningBoxes), &(product.OpeningPackets))
		p = append(p, product)
	}
	return p, true
}

//DeleteProduct deletes the product of the given id from table products
func DeleteProduct(productID int) bool {
	query := `DELETE FROM products
				WHERE productID= ?`
	_, err := db.Exec(query, productID)
	if err != nil {
		fmt.Println("Errror in Deleting a product from database")
		log.Fatal(err)
		return false
	} else {
		return true
	}
}

//ConnectDatabase connects to database Server at the start of the server
func ConnectDatabase() {
	fmt.Println("Connecting to the SQL server")
	var err error
	db, err = sql.Open("mysql", "root:ilijksms1999@/anjani_test")
	err1 := db.Ping()
	if err1 != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Server connected")

}
