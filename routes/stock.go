package routes

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/achhapolia10/inventory-manager/opdatabase"

	"github.com/achhapolia10/inventory-manager/model"
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
	if isLoggedIn(w, req) {
		t := template.Must(template.ParseGlob("views/components/navbar.comp"))
		t.ParseFiles("views/stock.html")
		data := struct{ U User }{currentUser}
		t.ExecuteTemplate(w, "stock.html", data)
	}
}

//PostStock Handler for route / method POST query: fdate,tdate
func PostStock(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if isLoggedIn(w, req) {
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
}

//GetProductStock for route /product method GET
func GetProductStock(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if isLoggedIn(w, req) {
		products, _ := model.GetAllProduct()
		data := struct {
			Products []opdatabase.Product
			U        User
		}{products, currentUser}
		t := template.Must(template.ParseGlob("views/components/navbar.comp"))
		t.ParseFiles("views/productstock.html")
		t.ExecuteTemplate(w, "productstock.html", data)
	}
}

//PostProductStock for route /product method POST
func PostProductStock(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if isLoggedIn(w, req) {
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
}

//GetStockPrint for route /stock/print method get
func GetStockPrint(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if isLoggedIn(w, req) {
		from := req.URL.Query().Get("fdate")
		to := req.URL.Query().Get("tdate")
		f := model.ParseDate(from)
		td := model.ParseDate(to)
		stockData := model.AllStock(from, to)
		data := struct {
			Stock map[int]model.Stock
			From  string
			To    string
		}{
			stockData, f.GetReadable(), td.GetReadable(),
		}
		t := template.Must(template.ParseGlob("views/components/navbar.comp"))
		t.ParseFiles("views/stockprint.html")
		err := t.ExecuteTemplate(w, "stockprint.html", data)
		if err != nil {
			log.Println(err)
		}

	}
}
