package routes

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/achhapolia10/anjaniv2/opdatabase"

	"github.com/achhapolia10/anjaniv2/model"
	"github.com/julienschmidt/httprouter"
)

//TemplateData For Templates
type TemplateData struct {
	Product opdatabase.Product
	From    string
	To      string
	Stock   map[int]model.Stock
}

//GetStock Handler for route / method GET
func GetStock(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	t := template.Must(template.ParseGlob("views/components/navbar.comp"))
	t.ParseFiles("views/stock.html")

	t.ExecuteTemplate(w, "stock.html", "")
}

//PostStock Handler for route / method POST query: fdate,tdate
func PostStock(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fromDate := req.FormValue("fdate")
	toDate := req.FormValue("tdate")
	stockData := model.AllStock(fromDate, toDate)
	b, err := json.Marshal(stockData)
	if err != nil {
		log.Printf("Error in Marshalling stock data: %v", err)
		return
	}
	w.Write(b)
}

//GetProductStock for route /product method GET
func GetProductStock(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	products, _ := model.GetAllProduct()
	t := template.Must(template.ParseGlob("views/components/navbar.comp"))
	t.ParseFiles("views/productstock.html")
	t.ExecuteTemplate(w, "productstock.html", products)
}

//PostProductStock for route /product method POST
func PostProductStock(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	from := req.URL.Query().Get("fdate")
	to := req.URL.Query().Get("tdate")
	i := req.URL.Query().Get("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		log.Println(err)
		return
	}
	details := model.ProductStockDetails(from, to, id)
	res, e := json.Marshal(details)
	if e != nil {
		log.Println(e)
		return
	}
	w.Write(res)
}
