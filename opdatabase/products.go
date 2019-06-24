package opdatabase

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//Product Model for the Products in Database
type Product struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	PacketQuantity int     `json:"packet"`
	BoxQuantity    int     `json:"box"`
	Price          float64 `json:"price"`
	OpeningBox     int     `json:"oboxes"`
	OpeningPacket  int     `json:"opackets"`
	Group          int     `json:"group"`
}

var db *sql.DB

//CreateProductTable creates a product table if it's not already created
func CreateProductTable() {
	query := `create table product(
		productID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(50) NOT NULL, 
		groupid INT NOT NULL,
		packetQuantity SMALLINT NOT NULL,
		boxQuantity SMALLINT NOT NULL,
		price DECIMAL(4,2) NOT NULL,
		oboxes INT NOT NULL,
		opackets INT NOT NULL);`

	fmt.Println("Creating Product Table : ")
	_, err := db.Exec(query)
	if err != nil {
		fmt.Println("Product Table already exists")

	} else {
		fmt.Println("Product Table Created")
	}
}

//AddProduct adds a new Product to table products
func AddProduct(p Product) {
	p.Name = strings.ToUpper(p.Name)
	query := "insert into product " +
		"(name,packetQuantity,boxQuantity,price,oboxes,opackets,groupid)" +
		"values	(?,?,?,?,?,?,?);"
	r, err := db.Exec(query, p.Name, p.PacketQuantity, p.BoxQuantity, p.Price, p.OpeningBox, p.OpeningPacket, p.Group)
	if err != nil {
		log.Println(err)
	} else {
		id, _ := r.LastInsertId()
		log.Printf("Added product at index: %v", id)
		r1 := CreateProductJournal(id)
		r2 := CreateProductStock(id)
		r3 := CreateProductMonth(id)
		if !(r1 && r2 && r3) {
			DeleteProduct(int(id))
		}
	}
}

//SelectProduct get products from the database
func SelectProduct() ([]Product, bool) {
	var p []Product
	query := `SELECT * FROM product;`
	r, err := db.Query(query)
	defer r.Close()
	if err != nil {
		fmt.Println("Can't get the products from the product table")
		log.Println(err)
		return p, false
	}
	for r.Next() {
		var (
			product Product
		)
		r.Scan(&(product.ID), &(product.Name), &(product.Group), &(product.PacketQuantity), &(product.BoxQuantity),
			&(product.Price), &(product.OpeningBox), &(product.OpeningPacket))
		p = append(p, product)
	}
	return p, true
}

//SelectProductMap get products from the database
func SelectProductMap() (map[int]Product, bool) {
	p := make(map[int]Product)
	query := `SELECT * FROM product;`
	r, err := db.Query(query)
	defer r.Close()
	if err != nil {
		log.Println("Can't get the products from the product table")
		log.Println(err)
		return p, false
	}
	for r.Next() {
		var (
			product Product
		)
		r.Scan(&(product.ID), &(product.Name), &(product.Group), &(product.PacketQuantity), &(product.BoxQuantity),
			&(product.Price), &(product.OpeningBox), &(product.OpeningPacket))
		p[product.ID] = product
	}
	return p, true
}

//SelectProductID get a single product form the database
func SelectProductID(id int) (Product, bool) {
	var product Product
	query := `SELECT * FROM product WHERE productID=?;`
	r, err := db.Query(query, id)
	defer r.Close()
	if err != nil {
		fmt.Println("Can't get the products from the product table")
		log.Println(err)
		return product, false
	}
	if r.Next() {
		r.Scan(&(product.ID), &(product.Name), &(product.Group), &(product.PacketQuantity), &(product.BoxQuantity),
			&(product.Price), &(product.OpeningBox), &(product.OpeningPacket))
	}
	return product, true
}

//SelectProductByGroup Selects all Product of a particular Group
func SelectProductByGroup(g Group) []Product {
	var products []Product
	query := `SELECT * FROM product WHERE groupid=?;`
	r, err := db.Query(query, g.Id)
	defer r.Close()
	if err != nil {
		fmt.Println("Can't get the products from the product table")
		log.Println(err)
	}
	for r.Next() {
		var product Product
		r.Scan(&(product.ID), &(product.Name), &(product.Group), &(product.PacketQuantity), &(product.BoxQuantity),
			&(product.Price), &(product.OpeningBox), &(product.OpeningPacket))
		products = append(products, product)
	}
	return products
}

//EditProduct eidt the produt at a given id
func EditProduct(id int, p Product) bool {
	if id != p.ID {
		fmt.Println("Illegal Edit Product Operation")
		return false
	}
	p.Name = strings.ToUpper(p.Name)
	query := `UPDATE product SET
			name=? ,
			packetQuantity=?,
			boxQuantity=?,
			price=? 
			WHERE productID=?;`
	_, err := db.Exec(query, p.Name, p.PacketQuantity, p.BoxQuantity, p.Price, id)
	if err != nil {
		fmt.Println("Error in edititng product ")
		log.Println(err)
		return false
	}
	fmt.Println("Edited Product at id ", id)
	return true
}

//DeleteProduct deletes the product of the given id from table products
func DeleteProduct(productID int) bool {
	query := `DELETE FROM product
				WHERE productID= ?;`
	_, err := db.Exec(query, productID)
	if err != nil {
		fmt.Println("Errror in Deleting a product from database")
		log.Println(err)
		return false
	}
	fmt.Println("Delted Product from id ", productID)
	DeleteProductJournal(productID)
	DeleteProductStock(productID)
	DeleteProductMonth(productID)
	return true

}

//DeleteProductByGroup Deletes a Product by a group
func DeleteProductByGroup(g Group) {
	p := SelectProductByGroup(g)
	if len(p) > 0 {
		for _, product := range p {
			DeleteProduct(product.ID)
		}
	}
}

//CreateProductJournal creates a journal Table for each product [id]journal
func CreateProductJournal(id int64) bool {
	query := "CREATE TABLE " + strconv.FormatInt(id, 10) + "journal(" +
		`id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			labour VARCHAR(50) NOT NULL DEFAULT 'ANSHU',
			date VARCHAR(50) NOT NULL DEFAULT '10/01/1999',
			 box INT NOT NULL,
			packet INT NOT NULL);`
	_, err := db.Exec(query)
	if err != nil {
		log.Println(err)
		log.Println("Error creating Journal Table for the product with id:", id)
		return false
	}
	log.Println("Created a Journal Table for the product with id:", id)
	return true
}

//DeleteProductJournal deletes a journal Table for each product [id]journal
func DeleteProductJournal(id int) bool {
	query := "DROP TABLE " + strconv.Itoa(id) + "journal;"
	_, err := db.Exec(query)
	if err != nil {
		log.Println(err)
		log.Println("Error in Deleting Journal Table for the product with id:", id)
		return false
	}
	log.Println("Deleted a Journal Table for the product with id:", id)
	return true
}

//CreateProductStock creates a Stock Table for each product [id]stock
func CreateProductStock(id int64) bool {
	query := "CREATE TABLE " + strconv.FormatInt(id, 10) + "stock(" +
		`id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		date VARCHAR(50) ,
		boxIn INT NOT NULL,
		packetIn INT NOT NULL,
		boxOut INT NOT NULL,
		packetOut INT NOT NULL);`
	_, err := db.Exec(query)
	if err != nil {
		log.Println(err)
		log.Println("Error creating Stock Table for the product with id:", id)
		return false
	}
	log.Println("Created a Stock Table for the product with id:", id)
	return true
}

//DeleteProductStock deletes a Stock Table for each product [id]stock
func DeleteProductStock(id int) bool {
	query := "DROP TABLE " + strconv.Itoa(id) + "stock;"
	_, err := db.Exec(query)
	if err != nil {
		log.Println(err)
		log.Println("Error in Deleting Stock Table for the product with id:", id)
		return false
	}
	log.Println("Deleted a Stock Table for the product with id:", id)
	return true
}

//CreateProductMonth creates a Stock Table for each product [id]stock
func CreateProductMonth(id int64) bool {
	query := "CREATE TABLE " + strconv.FormatInt(id, 10) + "month(" +
		`id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		date VARCHAR(50) ,
		boxIn INT NOT NULL,
		packetIn INT NOT NULL,
		boxOut INT NOT NULL,
		packetOut INT NOT NULL);`
	_, err := db.Exec(query)
	if err != nil {
		log.Println(err)
		log.Println("Error creating Month Table for the product with id:", id)
		return false
	}
	log.Println("Created a Month Table for the product with id:", id)
	return true
}

//DeleteProductMonth deletes a Stock Table for each product [id]stock
func DeleteProductMonth(id int) bool {
	query := "DROP TABLE " + strconv.Itoa(id) + "month;"
	_, err := db.Exec(query)
	if err != nil {
		log.Println(err)
		log.Println("Error in Month Stock Table for the product with id:", id)
		return false
	}
	log.Println("Deleted Month Table for the product with id:", id)
	return true
}
