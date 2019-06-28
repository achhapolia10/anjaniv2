package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/achhapolia10/anjaniv2/opdatabase"

	"github.com/achhapolia10/anjaniv2/routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

//Index Handler Method GET
func Index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	io.WriteString(w, "Hello")
}

func main() {
	var publicDir http.Dir
	publicDir = "./public/"

	router := httprouter.New()

	opdatabase.ConnectDatabase()
	opdatabase.CreateGroupTable()
	opdatabase.CreateProductTable()
	opdatabase.CreateUserTable()
	router.ServeFiles("/public/*filepath", publicDir)

	//Login routes are defined
	router.GET("/", routes.GetLogin)
	router.POST("/", routes.PostLogin)
	router.GET("/logout", routes.GetLogout)
	router.GET("/users", routes.GetUsers)

	//Groupds reoutes are  defined
	router.GET("/groups", routes.GetGroup)
	router.POST("/groups/new", routes.PostGroupNew)
	router.POST("/groups/edit/:id", routes.PostGroupEdit)
	router.POST("/groups/delete/:id", routes.PostGroupDelete)
	//Products routes are defined
	router.GET("/products", routes.GetProducts)
	router.GET("/products/new", routes.GetNewProduct)
	router.POST("/products/new", routes.PostNewProduct)
	router.GET("/products/delete/:id", routes.GetDeleteProducts)
	router.GET("/products/edit/:id", routes.GetEditProduct)
	router.POST("/products/edit/:id", routes.PostEditProduct)

	//Daily Report routes are defined
	router.GET("/report", routes.GetDailyReport)
	router.POST("/report", routes.PostDailyReport)

	//Stock routes are defined
	router.GET("/stock", routes.GetStock)
	router.POST("/stock", routes.PostStock)
	router.GET("/stock/product", routes.GetProductStock)
	router.POST("/stock/product", routes.PostProductStock)

	//Labour Payment routes are defined
	router.GET("/labourpayment", routes.GetLabourPayment)
	router.POST("/labourpayment", routes.PostLabourPayment)
	router.GET("/labourpayment/print", routes.GetPrintLabourPayment)

	//Entry routes are defined
	router.GET("/entry", routes.GetEntry)
	router.POST("/entry/new", routes.PostEntryNew)
	router.GET("/entry/getall", routes.GetJournalEntriesAll)
	router.POST("/entry/delete", routes.PostDeleteJournalEntry)

	//Dispatch routes are defined
	router.GET("/dispatch", routes.GetDispatch)
	router.GET("/dispatch/entries", routes.GetDispatchEntries)
	router.POST("/dispatch/new", routes.PostDispatchNew)
	router.GET("/dispatch/delete", routes.GetDispatchDelete)

	fmt.Println("Starting Server on Port: 4001")
	log.Println(http.ListenAndServe(":4001", router))
}
