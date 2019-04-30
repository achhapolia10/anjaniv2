package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/achhapolia10/anjaniv2/routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

//Index Handler Method GET
func Index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	io.WriteString(w, "Hello")
}

func main() {
	router := httprouter.New()

	// db, err := sql.Open("mysql", "root:ilijksms1999@/anjani_test")
	// err1 := db.Ping()
	// if err1 != nil {
	// 	log.Fatal(err)
	// }

	//Index route is defined
	router.GET("/", routes.GetIndex)

	//Login routes are defined
	router.GET("/login", routes.GetLogin)
	router.POST("/login", routes.PostLogin)
	router.GET("/logout", routes.GetLogin)

	//Products routes are defined
	router.GET("/products", routes.GetProducts)
	router.GET("/products/new", routes.GetNewProduct)

	//Stock routes are defined
	router.GET("/stock", routes.GetStock)

	//Labour Payment routes are defined
	router.GET("/labourpayment", routes.GetLabourPayment)

	//Entry routes are defined
	router.GET("/entry", routes.GetEntry)

	fmt.Println("Starting Server on Port: 4001")
	log.Fatal(http.ListenAndServe(":4001", router))
}