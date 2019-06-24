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

//GetProductStock for route /:id method GET
func GetProductStock(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	ids := params.ByName("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		log.Println(err)
	}

	t := template.Must(template.ParseGlob("views/components/navbar.comp"))
	t.ParseFiles("views/productstock.html")

	t.ExecuteTemplate(w, "productstock.html", id)
}
